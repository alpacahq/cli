//go:build integration

package integration

import "testing"

func TestDataEventContractSeries(t *testing.T) {
	t.Parallel()
	data, ok := alpacaJSONOrStructuredError(t,
		"data", "event-contract", "series",
		"--limit", "10",
	)
	if !ok {
		return
	}
	requireFields(t, data, "series")
}

func TestDataEventContractEvents(t *testing.T) {
	t.Parallel()
	data, ok := alpacaJSONOrStructuredError(t,
		"data", "event-contract", "events",
		"--limit", "10",
	)
	if !ok {
		return
	}
	requireFields(t, data, "events")
}

func TestDataEventContractContracts(t *testing.T) {
	t.Parallel()
	data, ok := alpacaJSONOrStructuredError(t,
		"data", "event-contract", "contracts",
		"--limit", "10",
	)
	if !ok {
		return
	}
	requireFields(t, data, "contracts")
}

func TestDataEventContractCategories(t *testing.T) {
	t.Parallel()
	// The categories endpoint returns a map keyed by category name, so there
	// are no fixed top-level fields to assert beyond valid JSON.
	if _, ok := alpacaJSONOrStructuredError(t, "data", "event-contract", "categories"); !ok {
		return
	}
}
