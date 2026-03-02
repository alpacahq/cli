package output

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestJSON(t *testing.T) {
	var buf bytes.Buffer
	data := map[string]string{"key": "value"}
	if err := JSON(&buf, data); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), `"key": "value"`) {
		t.Errorf("expected JSON output, got: %s", buf.String())
	}
}

func TestTable_EmptyData(t *testing.T) {
	var buf bytes.Buffer
	cols := []Column{{Header: "NAME", Field: "name"}}
	_ = Table(&buf, cols, json.RawMessage(`[]`))
	if !strings.Contains(buf.String(), "No results.") {
		t.Errorf("expected 'No results.', got: %s", buf.String())
	}
}

func TestTable_WithRows(t *testing.T) {
	var buf bytes.Buffer
	cols := []Column{
		{Header: "SYMBOL", Field: "symbol"},
		{Header: "PRICE", Field: "price"},
	}
	data := json.RawMessage(`[{"symbol":"AAPL","price":150.25},{"symbol":"MSFT","price":400.50}]`)
	if err := Table(&buf, cols, data); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "SYMBOL") || !strings.Contains(out, "AAPL") {
		t.Errorf("expected table output with AAPL, got: %s", out)
	}
}

func TestCSV_Output(t *testing.T) {
	var buf bytes.Buffer
	cols := []Column{
		{Header: "ID", Field: "id"},
		{Header: "NAME", Field: "name"},
	}
	data := json.RawMessage(`[{"id":"1","name":"Test"}]`)
	if err := CSV(&buf, cols, data); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "id,name") {
		t.Errorf("expected CSV header, got: %s", out)
	}
	if !strings.Contains(out, "1,Test") {
		t.Errorf("expected CSV row, got: %s", out)
	}
}

func TestRawField_Types(t *testing.T) {
	row := map[string]any{
		"str":   "hello",
		"num":   150.25,
		"int":   float64(42),
		"bool":  true,
		"empty": nil,
	}

	if got := rawField(row, "str"); got != "hello" {
		t.Errorf("string field: got %q", got)
	}
	if got := rawField(row, "num"); got != "150.25" {
		t.Errorf("float field: got %q", got)
	}
	if got := rawField(row, "int"); got != "42" {
		t.Errorf("int field: got %q", got)
	}
	if got := rawField(row, "bool"); got != "true" {
		t.Errorf("bool field: got %q", got)
	}
	if got := rawField(row, "empty"); got != "" {
		t.Errorf("nil field: got %q", got)
	}
	if got := rawField(row, "missing"); got != "" {
		t.Errorf("missing field: got %q", got)
	}
}

func TestColorPL(t *testing.T) {
	if got := ColorPL("0"); got != "0" {
		t.Errorf("zero: got %q", got)
	}
}

func TestDollarPL(t *testing.T) {
	if got := DollarPL(""); got != "$0.00" {
		t.Errorf("empty: got %q", got)
	}
	if got := DollarPL("0"); got != "$0" {
		t.Errorf("zero: got %q", got)
	}
}
