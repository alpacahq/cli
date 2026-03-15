package output

import (
	"encoding/json"
	"fmt"

	"github.com/itchyny/gojq"
)

// ApplyJQ runs a jq expression against data and returns the result.
// If the expression produces a single value, it is returned unwrapped.
// Multiple values are collected into a slice.
func ApplyJQ(data any, expr string) (any, error) {
	query, err := gojq.Parse(expr)
	if err != nil {
		return nil, fmt.Errorf("--jq: %w", err)
	}
	code, err := gojq.Compile(query)
	if err != nil {
		return nil, fmt.Errorf("--jq: %w", err)
	}

	input, err := normalize(data)
	if err != nil {
		return nil, fmt.Errorf("--jq: %w", err)
	}

	var results []any
	iter := code.Run(input)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if e, isErr := v.(error); isErr {
			return nil, fmt.Errorf("--jq: %w", e)
		}
		results = append(results, v)
	}

	if len(results) == 1 {
		return results[0], nil
	}
	return results, nil
}

// normalize converts any Go value into the untyped representation
// (map[string]any / []any / scalar) that gojq expects.
func normalize(data any) (any, error) {
	switch v := data.(type) {
	case json.RawMessage:
		var out any
		if err := json.Unmarshal(v, &out); err != nil {
			return nil, err
		}
		return out, nil
	default:
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		var out any
		if err := json.Unmarshal(b, &out); err != nil {
			return nil, err
		}
		return out, nil
	}
}
