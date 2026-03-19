package output

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"sort"
)

type Format string

const (
	FormatJSON Format = "json"
	FormatCSV  Format = "csv"
)

func Render(w io.Writer, format Format, data any) error {
	switch format {
	case FormatCSV:
		return CSV(w, data)
	default:
		return JSON(w, data)
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

func CSV(w io.Writer, data any) error {
	return CSVWithHeaders(w, data, nil)
}

func CSVWithHeaders(w io.Writer, data any, headers []string) error {
	rows := toRows(data)
	if len(rows) == 0 {
		if len(headers) == 0 {
			return nil
		}
		cw := csv.NewWriter(w)
		_ = cw.Write(headers)
		cw.Flush()
		return cw.Error()
	}

	keys := sortedKeys(rows[0])
	cw := csv.NewWriter(w)
	_ = cw.Write(keys)

	for _, row := range rows {
		vals := make([]string, len(keys))
		for i, k := range keys {
			vals[i] = rawField(row, k)
		}
		_ = cw.Write(vals)
	}

	cw.Flush()
	return cw.Error()
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
		var single map[string]any
		if json.Unmarshal(b, &single) == nil {
			return []map[string]any{single}
		}
		return nil
	}
}

func sortedKeys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
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
