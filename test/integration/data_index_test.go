//go:build integration

package integration

import "testing"

func TestDataIndexLatestValues(t *testing.T) {
	t.Parallel()
	data, ok := alpacaJSONOrStructuredError(t, "data", "index", "latest-values", "--symbols", "SPX,VIX")
	if !ok {
		return
	}
	requireFields(t, data, "values")
}

func TestDataIndexValues(t *testing.T) {
	t.Parallel()
	data, ok := alpacaJSONOrStructuredError(t,
		"data", "index", "values",
		"--symbols", "SPX,VIX",
		"--start", daysAgo(10),
		"--limit", "5",
	)
	if !ok {
		return
	}
	requireFields(t, data, "values")
}
