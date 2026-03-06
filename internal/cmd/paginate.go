package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

const maxPages = 100

func fetchAllDataPages(
	fetch func(pageToken string) (json.RawMessage, error),
	extract func(json.RawMessage) json.RawMessage,
	maxItems int,
) (json.RawMessage, error) {
	var allItems []json.RawMessage
	pageToken := ""

	for page := 0; page < maxPages; page++ {
		raw, err := fetch(pageToken)
		if err != nil {
			return nil, err
		}

		pageData := extract(raw)
		var items []json.RawMessage
		if err := json.Unmarshal(pageData, &items); err == nil {
			allItems = append(allItems, items...)
		}

		if maxItems > 0 && len(allItems) >= maxItems {
			allItems = allItems[:maxItems]
			break
		}

		npt := nextPageToken(raw)
		if npt == "" {
			break
		}
		pageToken = npt

		if verboseFlag {
			fmt.Fprintf(os.Stderr, "  fetching page %d (%d items so far)\n", page+2, len(allItems))
		}
	}

	result, err := json.Marshal(allItems)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func nextPageToken(raw json.RawMessage) string {
	var envelope map[string]json.RawMessage
	if json.Unmarshal(raw, &envelope) != nil {
		return ""
	}
	tokenRaw, ok := envelope["next_page_token"]
	if !ok {
		return ""
	}
	var token string
	if json.Unmarshal(tokenRaw, &token) != nil {
		return ""
	}
	return token
}
