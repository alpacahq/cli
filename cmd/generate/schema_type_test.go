package main

import "testing"

func TestSchemaType(t *testing.T) {
	tests := []struct {
		name         string
		schema       map[string]any
		wantType     string
		wantNullable bool
	}{
		{
			name:     "oas30 string",
			schema:   map[string]any{"type": "string"},
			wantType: "string",
		},
		{
			name:         "oas30 nullable string",
			schema:       map[string]any{"type": "string", "nullable": true},
			wantType:     "string",
			wantNullable: true,
		},
		{
			name:         "oas31 string or null",
			schema:       map[string]any{"type": []any{"string", "null"}},
			wantType:     "string",
			wantNullable: true,
		},
		{
			name:         "oas31 null first",
			schema:       map[string]any{"type": []any{"null", "integer"}},
			wantType:     "integer",
			wantNullable: true,
		},
		{
			name:         "oas31 with legacy nullable flag",
			schema:       map[string]any{"type": []any{"number", "null"}, "nullable": true},
			wantType:     "number",
			wantNullable: true,
		},
		{
			name:   "missing type",
			schema: map[string]any{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotType, gotNullable := schemaType(tt.schema)
			if gotType != tt.wantType || gotNullable != tt.wantNullable {
				t.Fatalf("schemaType() = (%q, %v), want (%q, %v)",
					gotType, gotNullable, tt.wantType, tt.wantNullable)
			}
			if isNullable(tt.schema) != tt.wantNullable {
				t.Fatalf("isNullable() = %v, want %v", isNullable(tt.schema), tt.wantNullable)
			}
		})
	}
}

func TestResolveGoTypeOAS31Nullable(t *testing.T) {
	got := resolveGoType(map[string]any{"type": []any{"string", "null"}})
	if got != "string" {
		t.Fatalf("resolveGoType() = %q, want %q", got, "string")
	}
}

func TestClassifyFieldsOAS31Nullable(t *testing.T) {
	fields, ok := classifyFields(map[string]map[string]any{
		"limit_price": {"type": []any{"string", "null"}},
		"qty":         {"type": "integer"},
		"symbol":      {"type": "string"},
	}, nil, nil)
	if !ok {
		t.Fatal("classifyFields() failed")
	}
	kinds := map[string]string{}
	for _, f := range fields {
		kinds[f.flagName] = f.kind
	}
	if kinds["limit-price"] != "ptrString" {
		t.Fatalf("limit-price kind = %q, want ptrString", kinds["limit-price"])
	}
	if kinds["qty"] != "int" || kinds["symbol"] != "string" {
		t.Fatalf("unexpected kinds: %#v", kinds)
	}
}
