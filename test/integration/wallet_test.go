//go:build integration

package integration

import "testing"

func TestWalletVASPSearch(t *testing.T) {
	t.Parallel()
	data, ok := alpacaJSONOrStructuredError(t,
		"wallet", "vasp", "search",
		"--q", "Coinbase",
		"--email-domain", "coinbase.com",
	)
	if ok {
		requireFields(t, data, "vasps", "page", "pages", "total")
	}
}

func TestWalletWhitelistUpdateTravelRule(t *testing.T) {
	data, ok := alpacaJSONOrStructuredError(t,
		"wallet", "whitelist", "update-travel-rule",
		"--whitelisted-address-id", "00000000-0000-0000-0000-000000000000",
		"--travel-rule-info", `{"beneficiary_is_self_hosted":true}`,
	)
	if ok && len(data) != 0 {
		t.Fatalf("expected an empty success response, got %v", data)
	}
}
