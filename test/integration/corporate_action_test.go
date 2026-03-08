//go:build integration

package integration

import (
	"testing"
)

func TestCorporateActionList(t *testing.T) {
	out := alpaca(t, "corporate-action", "list",
		"--ca-types", "dividend",
		"--since", daysAgo(180),
		"--until", daysAgo(90),
		"--json",
	)
	_ = parseJSONArray(t, out)
}

func TestCorporateActionGet(t *testing.T) {
	out := alpaca(t, "corporate-action", "list",
		"--ca-types", "dividend",
		"--since", daysAgo(180),
		"--until", daysAgo(90),
		"--json",
	)
	actions := parseJSONArray(t, out)
	if len(actions) == 0 {
		t.Skip("no corporate actions found to test get")
	}

	id, ok := actions[0]["id"].(string)
	if !ok || id == "" {
		t.Fatal("corporate action missing id")
	}

	out = alpaca(t, "corporate-action", "get", id, "--json")
	action := parseJSONMap(t, out)
	requireFields(t, action, "id")
}
