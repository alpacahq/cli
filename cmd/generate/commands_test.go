package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestBuildPostBody(t *testing.T) {
	// Set up global schema maps for $ref resolution.
	schemaByOASName = map[string]*schemaInfo{
		"OrderSide":   {goName: "OrderSide", kind: "enum"},
		"TimeInForce": {goName: "TimeInForce", kind: "enum"},
		"AdvancedInstructions": {
			goName: "AdvancedInstructions", kind: "struct",
			props: map[string]map[string]any{"mpid": {"type": "string"}},
		},
		"MLegOrderLeg": {
			goName: "MLegOrderLeg", kind: "struct",
			props: map[string]map[string]any{"symbol": {"type": "string"}},
		},
	}
	schemaByGoName = map[string]*schemaInfo{}
	for _, s := range schemaByOASName {
		schemaByGoName[s.goName] = s
	}

	tests := []struct {
		name       string
		typeName   string
		props      map[string]map[string]any
		skipFields []string
		wantOK     bool
		wantSnips  []string // substrings the output must contain
		wantAbsent []string // substrings the output must NOT contain
	}{
		{
			name:   "empty props",
			props:  nil,
			wantOK: false,
		},
		{
			name:     "string fields only",
			typeName: "Req",
			props: map[string]map[string]any{
				"symbol": {"type": "string"},
				"qty":    {"type": "string"},
			},
			wantOK: true,
			wantSnips: []string{
				`body := &api.Req{`,
				`Qty: cmdutil.Str(cmd, "qty")`,
				`Symbol: cmdutil.Str(cmd, "symbol")`,
			},
		},
		{
			name:     "boolean field",
			typeName: "Req",
			props: map[string]map[string]any{
				"extended_hours": {"type": "boolean"},
			},
			wantOK: true,
			wantSnips: []string{
				`ExtendedHours: cmdutil.Bool(cmd, "extended-hours")`,
			},
		},
		{
			name:     "integer field",
			typeName: "Req",
			props: map[string]map[string]any{
				"count": {"type": "integer"},
			},
			wantOK: true,
			wantSnips: []string{
				`Count: cmdutil.Int(cmd, "count")`,
			},
		},
		{
			name:     "enum ref",
			typeName: "Req",
			props: map[string]map[string]any{
				"side": {"$ref": "#/components/schemas/OrderSide"},
			},
			wantOK: true,
			wantSnips: []string{
				`Side: api.OrderSide(cmdutil.Str(cmd, "side"))`,
			},
		},
		{
			name:     "string array comma split",
			typeName: "Req",
			props: map[string]map[string]any{
				"symbols": {"type": "array", "items": map[string]any{"type": "string"}},
			},
			wantOK: true,
			wantSnips: []string{
				`body := &api.Req{}`,
				`cmdutil.Str(cmd, "symbols")`,
				`strings.Split(s, ",")`,
			},
		},
		{
			name:     "struct ref uses json unmarshal",
			typeName: "Req",
			props: map[string]map[string]any{
				"advanced_instructions": {"$ref": "#/components/schemas/AdvancedInstructions"},
			},
			wantOK: true,
			wantSnips: []string{
				`cmdutil.Changed(cmd, "advanced-instructions")`,
				`json.Unmarshal`,
				`&body.AdvancedInstructions`,
				`"--advanced-instructions:`,
			},
		},
		{
			name:     "nested object uses json unmarshal",
			typeName: "Req",
			props: map[string]map[string]any{
				"take_profit": {"type": "object", "properties": map[string]any{
					"limit_price": map[string]any{"type": "string"},
				}},
			},
			wantOK: true,
			wantSnips: []string{
				`cmdutil.Changed(cmd, "take-profit")`,
				`json.Unmarshal`,
				`&body.TakeProfit`,
			},
		},
		{
			name:     "array with ref items uses json unmarshal",
			typeName: "Req",
			props: map[string]map[string]any{
				"legs": {"type": "array", "items": map[string]any{
					"$ref": "#/components/schemas/MLegOrderLeg",
				}},
			},
			wantOK: true,
			wantSnips: []string{
				`cmdutil.Changed(cmd, "legs")`,
				`json.Unmarshal`,
				`&body.Legs`,
			},
		},
		{
			name:     "skip fields excluded",
			typeName: "Req",
			props: map[string]map[string]any{
				"symbol":      {"type": "string"},
				"take_profit": {"type": "object"},
				"stop_loss":   {"type": "object"},
			},
			skipFields: []string{"take-profit", "stop-loss"},
			wantOK:     true,
			wantSnips: []string{
				`Symbol: cmdutil.Str(cmd, "symbol")`,
			},
			wantAbsent: []string{
				"take-profit",
				"stop-loss",
			},
		},
		{
			name:     "mixed types",
			typeName: "PostOrderRequest",
			props: map[string]map[string]any{
				"symbol":                {"type": "string"},
				"extended_hours":        {"type": "boolean"},
				"side":                  {"$ref": "#/components/schemas/OrderSide"},
				"time_in_force":         {"$ref": "#/components/schemas/TimeInForce"},
				"advanced_instructions": {"$ref": "#/components/schemas/AdvancedInstructions"},
			},
			wantOK: true,
			wantSnips: []string{
				`body := &api.PostOrderRequest{`,
				`Symbol: cmdutil.Str(cmd, "symbol")`,
				`ExtendedHours: cmdutil.Bool(cmd, "extended-hours")`,
				`Side: api.OrderSide(cmdutil.Str(cmd, "side"))`,
				`TimeInForce: api.TimeInForce(cmdutil.Str(cmd, "time-in-force"))`,
				`json.Unmarshal`,
				`&body.AdvancedInstructions`,
			},
		},
		{
			name:     "unknown ref fails",
			typeName: "Req",
			props: map[string]map[string]any{
				"foo": {"$ref": "#/components/schemas/Unknown"},
			},
			wantOK: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := buildPostBody(tt.typeName, tt.props, tt.skipFields)
			if ok != tt.wantOK {
				t.Fatalf("buildPostBody() ok = %v, want %v\noutput:\n%s", ok, tt.wantOK, got)
			}
			if !ok {
				return
			}
			for _, snip := range tt.wantSnips {
				if !strings.Contains(got, snip) {
					t.Errorf("output missing expected snippet %q\nfull output:\n%s", snip, got)
				}
			}
			for _, snip := range tt.wantAbsent {
				if strings.Contains(got, snip) {
					t.Errorf("output unexpectedly contains %q\nfull output:\n%s", snip, got)
				}
			}
		})
	}
}

func TestExtractEndpointsHeaderParameter(t *testing.T) {
	spec := map[string]any{
		"paths": map[string]any{
			"/v1/locates": map[string]any{
				"post": map[string]any{
					"operationId": "createLocates",
					"parameters": []any{
						map[string]any{
							"name":        "Idempotency-Key",
							"in":          "header",
							"required":    false,
							"description": "Idempotency key for safe retries.",
							"schema":      map[string]any{"type": "string"},
						},
					},
				},
			},
		},
	}

	endpoints := extractEndpoints(spec)
	if len(endpoints) != 1 {
		t.Fatalf("extractEndpoints() returned %d endpoints, want 1", len(endpoints))
	}
	headers := endpoints[0].headerParams
	if len(headers) != 1 {
		t.Fatalf("headerParams has %d entries, want 1", len(headers))
	}
	if headers[0].name != "Idempotency-Key" {
		t.Errorf("header name = %q, want %q", headers[0].name, "Idempotency-Key")
	}
}

func TestHeaderParameterGeneration(t *testing.T) {
	ep := &endpointInfo{
		method:       "POST",
		path:         "/v1/locates",
		goName:       "CreateLocates",
		bodyRef:      "CreateLocateRequest",
		headerParams: []paramInfo{{name: "Idempotency-Key", goType: "string"}},
	}

	var method bytes.Buffer
	genEndpointMethod(&method, ep, "Trading")
	for _, want := range []string{
		"headers http.Header",
		`c.Raw.DoWithHeaders("POST", c.baseURL, "/v1/locates", nil, headers, body)`,
	} {
		if !strings.Contains(method.String(), want) {
			t.Errorf("generated method missing %q:\n%s", want, method.String())
		}
	}

	call := buildCallExpr(ep, cmdDef{}, "tradingClient", "body")
	wantCall := "tradingClient.CreateLocates(headersFromFlags(cmd, api.CreateLocatesOp), body)"
	if call != wantCall {
		t.Errorf("buildCallExpr() = %q, want %q", call, wantCall)
	}

	flag := paramToFlag(ep.headerParams[0], "header")
	if flag.flagName != "idempotency-key" || flag.oasName != "Idempotency-Key" || flag.source != "header" {
		t.Errorf("unexpected header flag: %#v", flag)
	}
}
