//go:build integration

package integration

import (
	"testing"
)

func TestAssetList_Filter(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "asset", "list", "--asset-class", "us_equity", "--status", "active", "--json")
	assets := requireArrayNonEmpty(t, out)
	for _, a := range assets[:min(len(assets), 5)] {
		if a["class"] != "us_equity" {
			t.Errorf("expected class us_equity, got %v", a["class"])
		}
	}
}

func TestAssetList_Exchange(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "asset", "list", "--exchange", "NASDAQ", "--json")
	_ = requireArrayNonEmpty(t, out)
}

func TestAssetList_CryptoClass(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "asset", "list", "--asset-class", "crypto", "--status", "active", "--json")
	assets := requireArrayNonEmpty(t, out)
	for _, a := range assets[:min(len(assets), 5)] {
		if a["class"] != "crypto" {
			t.Errorf("expected class crypto, got %v", a["class"])
		}
	}
}
