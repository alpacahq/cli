//go:build integration

package integration

import (
	"testing"
)

func TestTokenization_List(t *testing.T) {
	t.Parallel()
	stdout, stderr, code := alpacaWithStderr(t, "tokenization", "list")
	if code == 0 {
		_ = parseJSON[[]any](t, stdout)
	} else {
		errMap := parseJSONMap(t, stderr)
		if errMap["error"] == nil || errMap["error"] == "" {
			t.Error("expected structured JSON error with 'error' field")
		}
	}
}
