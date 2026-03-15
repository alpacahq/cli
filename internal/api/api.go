package api

import "encoding/json"

func unmarshal[T any](data json.RawMessage, err error) (*T, error) {
	if err != nil {
		return nil, err
	}
	var result T
	return &result, json.Unmarshal(data, &result)
}

func unmarshalSlice[T any](data json.RawMessage, err error) ([]T, error) {
	if err != nil {
		return nil, err
	}
	var result []T
	return result, json.Unmarshal(data, &result)
}
