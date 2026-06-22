//go:build integration

package integration

import "testing"

func TestLocateList(t *testing.T) {
	t.Parallel()
	data, ok := alpacaJSONOrStructuredError(t, "locate", "list", "--limit", "5")
	if !ok {
		return
	}
	requireFields(t, data, "locates", "next_page_token")
}

func TestLocateQuotes(t *testing.T) {
	t.Parallel()
	data, ok := alpacaJSONOrStructuredError(t, "locate", "quotes", "--symbols", "TSLA,AAPL")
	if !ok {
		return
	}
	requireFields(t, data, "quotes")
}
