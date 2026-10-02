//go:build integration

package integration

import (
	"encoding/json"
	"net/url"
	"strings"
	"testing"
)

func TestWalletVASPSearch(t *testing.T) {
	t.Parallel()
	stdout, stderr, code := alpacaWithStderr(t,
		"--debug",
		"wallet", "vasp", "search",
		"--q", "Coinbase",
		"--email-domain", "coinbase.com",
	)
	rawURL, _ := debugRequest(t, stderr, "GET")
	u, err := url.Parse(rawURL)
	if err != nil {
		t.Fatalf("parse request URL %q: %v", rawURL, err)
	}
	if !strings.HasSuffix(u.Path, "/v2/wallets/travel-rule/vasps") {
		t.Fatalf("request path = %q", u.Path)
	}
	q := u.Query()
	if q.Get("q") != "Coinbase" {
		t.Fatalf("q = %q, request URL %s", q.Get("q"), rawURL)
	}
	if q.Get("emailDomain") != "coinbase.com" {
		t.Fatalf("emailDomain = %q, request URL %s", q.Get("emailDomain"), rawURL)
	}
	if _, ok := q["email_domain"]; ok {
		t.Fatalf("request used email_domain: %s", rawURL)
	}
	if _, ok := q["email-domain"]; ok {
		t.Fatalf("request used email-domain: %s", rawURL)
	}

	if code == 0 {
		data := parseJSONMap(t, stdout)
		requireFields(t, data, "vasps", "page", "pages", "total")
		if _, ok := data["vasps"].([]any); !ok {
			t.Fatalf("vasps = %#v, want an array", data["vasps"])
		}
		return
	}
	allowWalletUnavailable(t, stderr)
}

func TestWalletWhitelistUpdateTravelRule(t *testing.T) {
	stdout, stderr, code := alpacaWithStderr(t,
		"--debug",
		"wallet", "whitelist", "update-travel-rule",
		"--whitelisted-address-id", "00000000-0000-0000-0000-000000000000",
		"--travel-rule-info", `{"beneficiary_is_self_hosted":true}`,
	)
	rawURL, body := debugRequest(t, stderr, "PATCH")
	u, err := url.Parse(rawURL)
	if err != nil {
		t.Fatalf("parse request URL %q: %v", rawURL, err)
	}
	if !strings.Contains(u.Path, "/v2/wallets/whitelists/00000000-0000-0000-0000-000000000000/travel-rule-info") {
		t.Fatalf("request path = %q", u.Path)
	}
	assertSelfHostedTravelRuleBody(t, body)

	if code == 0 {
		data := parseJSONMap(t, stdout)
		if len(data) != 0 {
			t.Fatalf("expected an empty success response, got %v", data)
		}
		return
	}
	errObj := cliErrorObject(t, stderr)
	requireFields(t, errObj, "error", "status")
	status, _ := errObj["status"].(float64)
	switch int(status) {
	case 403, 404:
		return
	case 400:
		errorCode, _ := errObj["error_code"].(string)
		actions, ok := errObj["next_actions"].([]any)
		if errorCode == "" || !ok || len(actions) == 0 {
			t.Fatalf("travel-rule validation response dropped error_code or next_actions: %s", stderr)
		}
	default:
		t.Fatalf("unexpected wallet error status %v: %s", status, stderr)
	}
}

func assertSelfHostedTravelRuleBody(t *testing.T, body string) {
	t.Helper()
	var payload map[string]any
	if err := json.Unmarshal([]byte(body), &payload); err != nil {
		t.Fatalf("request body is not JSON: %s", body)
	}
	raw, ok := payload["travel_rule_info"]
	if !ok {
		t.Fatalf("request omitted travel_rule_info: %s", body)
	}
	info, ok := raw.(map[string]any)
	if !ok {
		t.Fatalf("travel_rule_info = %#v", raw)
	}
	if info["beneficiary_is_self_hosted"] != true {
		t.Fatalf("beneficiary_is_self_hosted = %v", info["beneficiary_is_self_hosted"])
	}
	if _, present := info["beneficiary_manual_entry"]; present {
		t.Fatalf("unsupplied beneficiary_manual_entry was sent: %s", body)
	}
	if _, present := info["beneficiary_vasp_id"]; present {
		t.Fatalf("unsupplied beneficiary_vasp_id was sent: %s", body)
	}
}

func debugRequest(t *testing.T, stderr []byte, method string) (rawURL, body string) {
	t.Helper()
	prefix := "→ " + method + " "
	var found bool
	for _, line := range strings.Split(string(stderr), "\n") {
		if !found && strings.HasPrefix(line, prefix) {
			rawURL = strings.TrimPrefix(line, prefix)
			found = true
			continue
		}
		if found && body == "" && strings.HasPrefix(line, "→ {") {
			body = strings.TrimPrefix(line, "→ ")
		}
	}
	if !found {
		t.Fatalf("debug output missing %s request:\n%s", method, stderr)
	}
	if method != "GET" && body == "" {
		t.Fatalf("debug output missing %s body:\n%s", method, stderr)
	}
	return rawURL, body
}

func cliErrorObject(t *testing.T, stderr []byte) map[string]any {
	t.Helper()
	return parseJSONMap(t, []byte(stripCLIStderr(string(stderr))))
}

func allowWalletUnavailable(t *testing.T, stderr []byte) {
	t.Helper()
	errObj := cliErrorObject(t, stderr)
	requireFields(t, errObj, "error", "status")
	status, _ := errObj["status"].(float64)
	switch int(status) {
	case 403, 404:
		return
	default:
		t.Fatalf("wallet call failed with status %v: %s", status, stderr)
	}
}
