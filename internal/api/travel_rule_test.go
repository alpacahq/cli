package api

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestTravelRuleInfoOmitsUnsetDestination(t *testing.T) {
	cases := []string{
		`{"beneficiary_is_self_hosted":true}`,
		`{"beneficiary_vasp_id":"did:ethr:0xabc"}`,
	}
	for _, in := range cases {
		var info TravelRuleInfo
		if err := json.Unmarshal([]byte(in), &info); err != nil {
			t.Fatalf("unmarshal %s: %v", in, err)
		}
		out, err := json.Marshal(info)
		if err != nil {
			t.Fatal(err)
		}
		if bytes.Contains(out, []byte("beneficiary_manual_entry")) {
			t.Fatalf("marshaled %s as %s", in, out)
		}
	}
}

func TestTravelRuleInfoKeepsSuppliedManualEntry(t *testing.T) {
	in := `{"beneficiary_manual_entry":{"vasp_name":"Alpaca","vasp_website":"https://alpaca.markets"}}`
	var info TravelRuleInfo
	if err := json.Unmarshal([]byte(in), &info); err != nil {
		t.Fatal(err)
	}
	out, err := json.Marshal(info)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(out, []byte(`"beneficiary_manual_entry":{"vasp_name":"Alpaca","vasp_website":"https://alpaca.markets"}`)) {
		t.Fatalf("manual entry was dropped: %s", out)
	}
}

func TestUpdateTravelRuleRequestOmitsUnsetManualEntry(t *testing.T) {
	var body UpdateWhitelistedAddressTravelRuleInfoRequest
	if err := json.Unmarshal([]byte(`{"beneficiary_is_self_hosted":true}`), &body.TravelRuleInfo); err != nil {
		t.Fatal(err)
	}
	out, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]any
	if err := json.Unmarshal(out, &got); err != nil {
		t.Fatal(err)
	}
	info, ok := got["travel_rule_info"].(map[string]any)
	if !ok {
		t.Fatalf("travel_rule_info missing: %s", out)
	}
	if info["beneficiary_is_self_hosted"] != true {
		t.Fatalf("beneficiary_is_self_hosted = %v", info["beneficiary_is_self_hosted"])
	}
	if _, present := info["beneficiary_manual_entry"]; present {
		t.Fatalf("unsupplied beneficiary_manual_entry was sent: %s", out)
	}
}
