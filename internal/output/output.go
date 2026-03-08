package output

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/fatih/color"
)

type Format string

const (
	FormatJSON  Format = "json"
	FormatCSV   Format = "csv"
	FormatTable Format = "table"
)

type Column struct {
	Header string
	Field  string
	Format func(any) string
}

func Render(w io.Writer, format Format, columns []Column, data any) error {
	return dispatch(w, format, columns, data, func() error {
		return Table(w, columns, data)
	})
}

func dispatch(w io.Writer, format Format, columns []Column, data any, tableFn func() error) error {
	switch format {
	case FormatJSON:
		return JSON(w, data)
	case FormatCSV:
		return CSV(w, columns, data)
	default:
		return tableFn()
	}
}

func JSON(w io.Writer, data any) error {
	if data != nil {
		v := reflect.ValueOf(data)
		if v.Kind() == reflect.Slice && v.IsNil() {
			_, err := io.WriteString(w, "[]\n")
			return err
		}
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}

func Table(w io.Writer, columns []Column, data any) error {
	rows := toRows(data)
	if len(rows) == 0 {
		_, _ = fmt.Fprintln(w, "No results.")
		return nil
	}

	if len(columns) == 0 {
		columns = columnsFromData(rows[0])
	}

	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)

	headers := make([]string, len(columns))
	for i, c := range columns {
		headers[i] = c.Header
	}
	_, _ = fmt.Fprintln(tw, strings.Join(headers, "\t"))

	for _, row := range rows {
		vals := make([]string, len(columns))
		for i, c := range columns {
			vals[i] = formatField(row, c)
		}
		_, _ = fmt.Fprintln(tw, strings.Join(vals, "\t"))
	}

	return tw.Flush()
}

func CSV(w io.Writer, columns []Column, data any) error {
	rows := toRows(data)
	if len(columns) == 0 && len(rows) > 0 {
		columns = columnsFromData(rows[0])
	}

	cw := csv.NewWriter(w)

	headers := make([]string, len(columns))
	for i, c := range columns {
		headers[i] = c.Field
	}
	_ = cw.Write(headers)

	for _, row := range rows {
		vals := make([]string, len(columns))
		for i, c := range columns {
			vals[i] = rawField(row, c.Field)
		}
		_ = cw.Write(vals)
	}

	cw.Flush()
	return cw.Error()
}

func PrintSingle(w io.Writer, format Format, columns []Column, data any) error {
	return dispatch(w, format, columns, data, func() error {
		row := toMap(data)
		if row == nil {
			return fmt.Errorf("no data")
		}

		if len(columns) == 0 {
			columns = columnsFromData(row)
		}

		tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
		for _, c := range columns {
			val := formatField(row, c)
			_, _ = fmt.Fprintf(tw, "%s:\t%s\n", c.Header, val)
		}
		return tw.Flush()
	})
}

func colorPL(val string) string {
	if strings.HasPrefix(val, "-") {
		return color.RedString(val)
	}
	if val != "0" && val != "0.00" && val != "" {
		return color.GreenString("+" + val)
	}
	return val
}

func DollarPL(val string) string {
	if val == "" {
		return "$0.00"
	}
	if strings.HasPrefix(val, "-") {
		return color.RedString("-$" + val[1:])
	}
	if val != "0" && val != "0.00" {
		return color.GreenString("+$" + val)
	}
	return "$" + val
}

func PercentPL(val string) string {
	if val == "" {
		return "0.00%"
	}
	return colorPL(val) + "%"
}

func FormatTimestamp(v any) string {
	s := fmt.Sprintf("%v", v)
	if s == "" || s == "<nil>" {
		return ""
	}
	for _, layout := range []string{time.RFC3339Nano, time.RFC3339, "2006-01-02T15:04:05Z", "2006-01-02T15:04:05"} {
		if t, err := time.Parse(layout, s); err == nil {
			return t.Local().Format("Jan 02 15:04")
		}
	}
	if _, err := time.Parse("2006-01-02", s); err == nil {
		return s
	}
	return s
}

func toRows(data any) []map[string]any {
	switch v := data.(type) {
	case []map[string]any:
		return v
	case json.RawMessage:
		var arr []map[string]any
		if json.Unmarshal(v, &arr) == nil {
			return arr
		}
		var single map[string]any
		if json.Unmarshal(v, &single) == nil {
			return []map[string]any{single}
		}
		return nil
	default:
		b, _ := json.Marshal(v)
		var arr []map[string]any
		if json.Unmarshal(b, &arr) == nil {
			return arr
		}
		return nil
	}
}

func toMap(data any) map[string]any {
	switch v := data.(type) {
	case map[string]any:
		return v
	case json.RawMessage:
		var m map[string]any
		_ = json.Unmarshal(v, &m)
		return m
	default:
		b, _ := json.Marshal(v)
		var m map[string]any
		_ = json.Unmarshal(b, &m)
		return m
	}
}

// columnsFromData auto-discovers columns from a data row's keys.
// Keys are sorted alphabetically. Timestamp fields get auto-formatting.
func columnsFromData(row map[string]any) []Column {
	keys := make([]string, 0, len(row))
	for k := range row {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	cols := make([]Column, len(keys))
	for i, k := range keys {
		cols[i] = Column{
			Header: strings.ToUpper(strings.ReplaceAll(k, "_", " ")),
			Field:  k,
		}
		if IsTimestampField(k) {
			cols[i].Format = FormatTimestamp
		}
	}
	return cols
}

// IsTimestampField returns true for JSON field names that typically contain timestamps.
func IsTimestampField(name string) bool {
	switch {
	case strings.HasSuffix(name, "_at"):
		return true
	case strings.HasSuffix(name, "_time"):
		return true
	case name == "t" || name == "transaction_time":
		return true
	default:
		return false
	}
}

func formatField(row map[string]any, col Column) string {
	if col.Format != nil {
		return col.Format(row[col.Field])
	}
	return rawField(row, col.Field)
}

func rawField(row map[string]any, field string) string {
	v, ok := row[field]
	if !ok {
		return ""
	}
	switch val := v.(type) {
	case string:
		return val
	case float64:
		if val == float64(int64(val)) {
			return fmt.Sprintf("%d", int64(val))
		}
		return fmt.Sprintf("%.2f", val)
	case bool:
		if val {
			return "true"
		}
		return "false"
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", val)
	}
}
