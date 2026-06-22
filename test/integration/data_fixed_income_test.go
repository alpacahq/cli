//go:build integration

package integration

import "testing"

func TestDataFixedIncomeLatestQuotes(t *testing.T) {
	t.Parallel()
	data, ok := alpacaJSONOrStructuredError(t,
		"data", "fixed-income-quotes",
		"--isins", "US912797SX61,US912810SK51",
		"--trade-size", "1000",
	)
	if !ok {
		return
	}
	requireFields(t, data, "quotes")
}
