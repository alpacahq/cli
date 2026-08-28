//go:build integration

package integration

import (
	"testing"
)

func TestDataFixedIncomeLatestPrices(t *testing.T) {
	t.Parallel()
	data, ok := alpacaJSONOrStructuredError(t,
		"data", "fixed-income", "latest-prices",
		"--isins", "US912797KJ59,US912797KS58",
	)
	if !ok {
		return
	}
	requireFields(t, data, "prices")
}

func TestDataFixedIncomeLatestQuotes(t *testing.T) {
	t.Parallel()
	data, ok := alpacaJSONOrStructuredError(t,
		"data", "fixed-income", "latest-quotes",
		"--isins", "US912797SX61,US912810SK51",
	)
	if !ok {
		return
	}
	requireFields(t, data, "quotes")
}
