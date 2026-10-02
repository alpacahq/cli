package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/alpacahq/cli/internal/client"
)

func TestAPIErrorJSONIncludesTravelRuleFields(t *testing.T) {
	obj := apiErrorJSON(&client.APIError{
		StatusCode:  400,
		Message:     "Please provide the beneficiary's name.",
		ErrorCode:   "BENEFICIARY_NAME_REQUIRED",
		NextActions: json.RawMessage(`[{"name":"supply_beneficiaryName_naturalPerson","type":"ONE_OF"}]`),
	})
	if obj["code"] != 0 {
		t.Fatalf("code = %v, want 0", obj["code"])
	}
	if obj["error_code"] != "BENEFICIARY_NAME_REQUIRED" {
		t.Fatalf("error_code = %v", obj["error_code"])
	}
	encoded, err := json.Marshal(obj)
	if err != nil {
		t.Fatal(err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatal(err)
	}
	actions, ok := decoded["next_actions"].([]any)
	if !ok || len(actions) != 1 {
		t.Fatalf("next_actions = %#v", decoded["next_actions"])
	}
	first, ok := actions[0].(map[string]any)
	if !ok || first["name"] != "supply_beneficiaryName_naturalPerson" {
		t.Fatalf("next_actions[0] = %#v", actions[0])
	}
}

func TestAPIErrorJSONOmitsAbsentTravelRuleFields(t *testing.T) {
	obj := apiErrorJSON(&client.APIError{StatusCode: 404, Code: 40410000, Message: "not found"})
	if _, ok := obj["error_code"]; ok {
		t.Fatalf("error_code = %v", obj["error_code"])
	}
	if _, ok := obj["next_actions"]; ok {
		t.Fatalf("next_actions = %v", obj["next_actions"])
	}
}

func TestUpdateTravelRuleOmitsUnsetManualEntry(t *testing.T) {
	const resp = `{"error_code":"BENEFICIARY_NAME_REQUIRED","message":"Please provide the beneficiary's name.","next_actions":[{"name":"supply_beneficiaryName_naturalPerson","description":"Provide the recipient's first and last name.","type":"ONE_OF"}]}`
	var gotBody []byte
	cleanup := setupMockClients(t, func(w http.ResponseWriter, r *http.Request) {
		var err error
		gotBody, err = io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("read body: %v", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(resp))
	})
	defer cleanup()

	root := Root()
	root.SetOut(new(bytes.Buffer))
	root.SetErr(new(bytes.Buffer))
	root.SetArgs([]string{
		"wallet", "whitelist", "update-travel-rule",
		"--whitelisted-address-id", "00000000-0000-0000-0000-000000000000",
		"--travel-rule-info", `{"beneficiary_is_self_hosted":true}`,
	})
	err := root.Execute()
	var apiErr *client.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected API error, got %v", err)
	}

	var sent map[string]any
	if err := json.Unmarshal(gotBody, &sent); err != nil {
		t.Fatalf("request body %s: %v", gotBody, err)
	}
	info, ok := sent["travel_rule_info"].(map[string]any)
	if !ok {
		t.Fatalf("request body missing travel_rule_info: %s", gotBody)
	}
	if _, present := info["beneficiary_manual_entry"]; present {
		t.Fatalf("unsupplied beneficiary_manual_entry was sent: %s", gotBody)
	}
	if info["beneficiary_is_self_hosted"] != true {
		t.Fatalf("beneficiary_is_self_hosted = %v", info["beneficiary_is_self_hosted"])
	}

	out := apiErrorJSON(apiErr)
	if out["error_code"] != "BENEFICIARY_NAME_REQUIRED" {
		t.Fatalf("error_code = %v", out["error_code"])
	}
	actions, ok := out["next_actions"].(json.RawMessage)
	if !ok {
		t.Fatalf("next_actions type = %T", out["next_actions"])
	}
	var decoded []map[string]any
	if err := json.Unmarshal(actions, &decoded); err != nil {
		t.Fatal(err)
	}
	if len(decoded) != 1 || decoded[0]["name"] != "supply_beneficiaryName_naturalPerson" {
		t.Fatalf("next_actions = %#v", decoded)
	}
}
