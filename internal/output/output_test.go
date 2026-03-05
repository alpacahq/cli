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
	if got := colorPL("0"); got != "0" {
		t.Errorf("zero: got %q", got)
	}
	if got := colorPL("0.00"); got != "0.00" {
		t.Errorf("0.00: got %q", got)
	}
	if got := colorPL(""); got != "" {
		t.Errorf("empty: got %q", got)
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

func TestPercentPL(t *testing.T) {
	if got := PercentPL(""); got != "0.00%" {
		t.Errorf("empty: got %q", got)
	}
	if got := PercentPL("0"); got != "0%" {
		t.Errorf("zero: got %q", got)
	}
}

func TestRender_JSON(t *testing.T) {
	var buf bytes.Buffer
	cols := []Column{{Header: "NAME", Field: "name"}}
	data := json.RawMessage(`[{"name":"AAPL"}]`)
	if err := Render(&buf, "json", cols, data); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "AAPL") {
		t.Errorf("expected AAPL in JSON output, got: %s", buf.String())
	}
}

func TestRender_Table(t *testing.T) {
	var buf bytes.Buffer
	cols := []Column{{Header: "SYMBOL", Field: "symbol"}}
	data := json.RawMessage(`[{"symbol":"MSFT"}]`)
	if err := Render(&buf, "table", cols, data); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "SYMBOL") || !strings.Contains(out, "MSFT") {
		t.Errorf("expected table with SYMBOL/MSFT, got: %s", out)
	}
}

func TestRender_CSV(t *testing.T) {
	var buf bytes.Buffer
	cols := []Column{{Header: "ID", Field: "id"}}
	data := json.RawMessage(`[{"id":"123"}]`)
	if err := Render(&buf, "csv", cols, data); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "123") {
		t.Errorf("expected 123, got: %s", buf.String())
	}
}

func TestRenderWithHint_EmptyTable(t *testing.T) {
	var buf bytes.Buffer
	cols := []Column{{Header: "NAME", Field: "name"}}
	data := json.RawMessage(`[]`)
	if err := RenderWithHint(&buf, "table", cols, data, "No items found."); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "No items found.") {
		t.Errorf("expected hint, got: %s", buf.String())
	}
}

func TestRenderWithHint_EmptyJSON(t *testing.T) {
	var buf bytes.Buffer
	cols := []Column{{Header: "NAME", Field: "name"}}
	data := json.RawMessage(`[]`)
	if err := RenderWithHint(&buf, "json", cols, data, "No items found."); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "[]") {
		t.Errorf("expected empty JSON array, got: %s", buf.String())
	}
}

func TestRenderWithHint_NonEmpty(t *testing.T) {
	var buf bytes.Buffer
	cols := []Column{{Header: "NAME", Field: "name"}}
	data := json.RawMessage(`[{"name":"test"}]`)
	if err := RenderWithHint(&buf, "table", cols, data, "Should not appear"); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(buf.String(), "Should not appear") {
		t.Error("hint should not appear for non-empty data")
	}
}

func TestPrintSingle_Table(t *testing.T) {
	var buf bytes.Buffer
	cols := []Column{
		{Header: "Symbol", Field: "symbol"},
		{Header: "Price", Field: "price"},
	}
	data := json.RawMessage(`{"symbol":"AAPL","price":150.5}`)
	if err := PrintSingle(&buf, "table", cols, data); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "Symbol:") || !strings.Contains(out, "AAPL") {
		t.Errorf("expected key-value format, got: %s", out)
	}
}

func TestPrintSingle_JSON(t *testing.T) {
	var buf bytes.Buffer
	cols := []Column{{Header: "Symbol", Field: "symbol"}}
	data := json.RawMessage(`{"symbol":"AAPL"}`)
	if err := PrintSingle(&buf, "json", cols, data); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "AAPL") {
		t.Errorf("expected AAPL, got: %s", buf.String())
	}
}

func TestPrintSingle_NoData(t *testing.T) {
	var buf bytes.Buffer
	cols := []Column{{Header: "Name", Field: "name"}}
	if err := PrintSingle(&buf, "table", cols, nil); err == nil {
		t.Error("expected error for nil data")
	}
}
