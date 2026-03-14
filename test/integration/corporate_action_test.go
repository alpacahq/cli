//go:build integration

package integration

import (
	"testing"
)

func TestCorporateAction(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "corporate-action", "list",
		"--ca-types", "dividend",
		"--since", daysAgo(180),
		"--until", daysAgo(90),
	)
	actions := parseJSONArray(t, out)

	t.Run("get", func(t *testing.T) {
		if len(actions) == 0 {
			t.Skip("no corporate actions found to test get")
		}

		id, ok := actions[0]["id"].(string)
		if !ok || id == "" {
			t.Fatal("corporate action missing id")
		}

		out := alpaca(t, "corporate-action", "get", "--id", id)
		action := parseJSONMap(t, out)
		requireFields(t, action, "id")
	})
}
