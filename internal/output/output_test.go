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

func TestJSON_NilSlice(t *testing.T) {
	var buf bytes.Buffer
	var data []string
	if err := JSON(&buf, data); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(buf.String()) != "[]" {
		t.Errorf("expected [], got: %s", buf.String())
	}
}

func TestCSV(t *testing.T) {
	var buf bytes.Buffer
	data := json.RawMessage(`[{"id":"1","name":"Test"}]`)
	if err := CSV(&buf, data); err != nil {
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

func TestCSV_SingleObject(t *testing.T) {
	var buf bytes.Buffer
	data := json.RawMessage(`{"symbol":"AAPL","price":150.5}`)
	if err := CSV(&buf, data); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "AAPL") {
		t.Errorf("expected AAPL in CSV, got: %s", out)
	}
}

func TestCSV_Empty(t *testing.T) {
	var buf bytes.Buffer
	if err := CSV(&buf, json.RawMessage(`[]`)); err != nil {
		t.Fatal(err)
	}
	if buf.Len() != 0 {
		t.Errorf("expected empty output for empty data, got: %s", buf.String())
	}
}

func TestCSVWithHeaders_Empty(t *testing.T) {
	var buf bytes.Buffer
	if err := CSVWithHeaders(&buf, json.RawMessage(`[]`), []string{"id", "symbol"}); err != nil {
		t.Fatal(err)
	}
	if got := strings.TrimSpace(buf.String()); got != "id,symbol" {
		t.Errorf("expected header-only CSV, got: %s", got)
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

func TestRender_JSON(t *testing.T) {
	var buf bytes.Buffer
	data := json.RawMessage(`[{"name":"AAPL"}]`)
	if err := Render(&buf, FormatJSON, data); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "AAPL") {
		t.Errorf("expected AAPL in JSON output, got: %s", buf.String())
	}
}

func TestRender_CSV(t *testing.T) {
	var buf bytes.Buffer
	data := json.RawMessage(`[{"id":"123"}]`)
	if err := Render(&buf, FormatCSV, data); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "123") {
		t.Errorf("expected 123, got: %s", buf.String())
	}
}

func TestRender_DefaultIsJSON(t *testing.T) {
	var buf bytes.Buffer
	data := json.RawMessage(`[{"symbol":"MSFT"}]`)
	if err := Render(&buf, "", data); err != nil {
		t.Fatal(err)
	}
	var parsed []map[string]any
	if err := json.Unmarshal(buf.Bytes(), &parsed); err != nil {
		t.Fatalf("default output should be valid JSON: %v", err)
	}
	if parsed[0]["symbol"] != "MSFT" {
		t.Errorf("expected MSFT, got: %v", parsed[0]["symbol"])
	}
}
