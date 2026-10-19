package main

import (
	"strings"
	"testing"
)

func TestToFlagName(t *testing.T) {
	tests := map[string]string{
		"page_token":             "page-token",
		"emailDomain":            "email-domain",
		"chainalysisName":        "chainalysis-name",
		"includeSubsidiaryVASPs": "include-subsidiary-vasps",
		"Last-Event-Id":          "last-event-id",
	}
	for input, want := range tests {
		if got := toFlagName(input); got != want {
			t.Errorf("toFlagName(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestNormalizeSummaryPreservesSearchVerb(t *testing.T) {
	if got := normalizeSummary("GET", "Search for VASPs", false); got != "Search for vasps" {
		t.Fatalf("normalizeSummary() = %q, want %q", got, "Search for vasps")
	}
}

func TestNormalizeOASDescription(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want string
	}{
		{name: "empty", raw: "", want: ""},
		{name: "whitespace only", raw: "   \n\t  ", want: ""},
		{
			name: "simple one-liner",
			raw:  "Returns your account details.",
			want: "Returns your account details",
		},
		{
			name: "first paragraph only",
			raw:  "First paragraph here.\n\nSecond paragraph ignored.",
			want: "First paragraph here",
		},
		{
			name: "multi-line first paragraph joined",
			raw:  "Line one of the description.\nLine two continues.\n\nSecond paragraph.",
			want: "Line one of the description. Line two continues",
		},
		{
			name: "blockquote markers stripped",
			raw:  "> Warning\n>\n> This is a warning block.",
			want: "Warning This is a warning block",
		},
		{
			name: "blockquote with content prefix",
			raw:  "> ⚠️ Warning\n>\n> Do not use this.",
			want: "⚠️ Warning Do not use this",
		},
		{
			name: "markdown links converted to text",
			raw:  "Use [the docs](https://example.com) for details.",
			want: "Use the docs for details",
		},
		{
			name: "multiple markdown links",
			raw:  "See [foo](https://a.com) and [bar](https://b.com).",
			want: "See foo and bar",
		},
		{
			name: "unclosed bracket not mangled",
			raw:  "Array items like [0] are fine.",
			want: "Array items like [0] are fine",
		},
		{
			name: "trailing periods stripped",
			raw:  "Multiple trailing periods...",
			want: "Multiple trailing periods",
		},
		{
			name: "long text truncated at sentence boundary",
			raw:  "First sentence here. " + string(make([]byte, 300)),
			want: "First sentence here",
		},
		{
			name: "long text truncated at word boundary when no sentence break",
			raw:  "Word " + repeatWord("word", 80),
			want: "Word " + repeatWord("word", 59),
		},
		{
			name: "exactly 300 chars not truncated",
			raw:  repeat('A', 300),
			want: repeat('A', 300),
		},
		{
			name: "301 chars truncated",
			raw:  repeat('A', 301),
			want: repeat('A', 300),
		},
		{
			name: "real OAS: corporate actions",
			raw:  "This endpoint provides data about the corporate actions for each given symbol over a specified time period.\n\n> ⚠️ Warning\n>\n> Currently Alpaca has no guarantees...",
			want: "This endpoint provides data about the corporate actions for each given symbol over a specified time period",
		},
		{
			name: "real OAS: deprecated with markdown link",
			raw:  "This endpoint is deprecated, please use [the new corporate actions endpoint](https://docs.alpaca.markets/reference/corporateactions-1) instead.",
			want: "This endpoint is deprecated, please use the new corporate actions endpoint instead",
		},
		{
			name: "real OAS: account details",
			raw:  "Returns your account details.",
			want: "Returns your account details",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeOASDescription(tt.raw)
			if got != tt.want {
				t.Errorf("normalizeOASDescription(%q)\n  got:  %q\n  want: %q", tt.raw, got, tt.want)
			}
		})
	}
}

func TestFlagDescriptionKeepsFormatGuidance(t *testing.T) {
	period := "The total duration of history being requested defined as `number` + `unit`.\nValid units are: `D`, `W`, `M`, and `A` (annum).\n\nThe following restrictions apply:\n- MUST NOT be included if both `start` and `end` are provided.\n"
	got := flagDescription(period, []string{"7D", "4W", "1M", "1A"})
	for _, want := range []string{
		"number + unit",
		"Valid units are: D, W, M, and A (annum)",
		"e.g. 7D, 4W, 1M, 1A",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("period help %q missing %q", got, want)
		}
	}
	if strings.Contains(got, "MUST NOT") {
		t.Errorf("period help included restrictions: %q", got)
	}

	start := "The first timestamp of the timeseries.\n\nDefaults to `end - period`.\n\nWhen `timeframe=1D`, if `start` occurs any time on a trading day (including outside of market hours), `start` is set to the closing time of the current trading day. Otherwise, `start` is rolled forward.\n"
	got = flagDescription(start, []string{
		"2021-03-16T18:38:01Z",
		"2021-03-16T18:38:01-04:00",
		"2021-03-16",
		"2021-03-16T14:38:01-04:00",
		"1615919881",
	})
	want := "first timestamp of the timeseries. e.g. 2021-03-16T18:38:01Z, 2021-03-16T18:38:01-04:00, 2021-03-16, 1615919881"
	if got != want {
		t.Errorf("start help\n got: %q\nwant: %q", got, want)
	}

	dated := "The inclusive start of the interval. Format: RFC-3339 or YYYY-MM-DD.\nDefault: the beginning of the current day, but at least 15 minutes ago if the user doesn't have real-time access for the feed.\n"
	got = flagDescription(dated, []string{"2024-01-03T01:02:03.123456789Z", "2024-01-03"})
	want = "inclusive start of the interval. Format: RFC-3339 or YYYY-MM-DD"
	if got != want {
		t.Errorf("format sentence should replace the example list\n got: %q\nwant: %q", got, want)
	}

	long := "Filter to snapshots that were updated since this timestamp, meaning that the timestamp of the trade or the quote is greater than or equal to this value.\nFormat: RFC-3339 or YYYY-MM-DD. If missing, all values are returned.\n"
	got = flagDescription(long, nil)
	want = "filter to snapshots that were updated since this timestamp. Format: RFC-3339 or YYYY-MM-DD"
	if got != want {
		t.Errorf("truncated help\n got: %q\nwant: %q", got, want)
	}

	locate := "Filter locates with locate trading date on or after this date. Format: YYYY-MM-DD. The locate trading date uses America/New_York and rolls at 8pm ET."
	got = flagDescription(locate, nil)
	want = "filter locates with locate trading date on or after this date. Format: YYYY-MM-DD"
	if got != want {
		t.Errorf("locate help\n got: %q\nwant: %q", got, want)
	}

	already := "Get activities created after this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported."
	if got, want := flagDescription(already, nil), normalizeDesc(already); got != want {
		t.Errorf("inlined format was rewritten\n got: %q\nwant: %q", got, want)
	}

	sampled := "One of 'ALL', 'NONE', or a comma-separated list of activity types."
	got = flagDescription(sampled, []string{"ALL", "NONE", "JNLC,CSD,CSW"})
	if strings.Contains(got, "e.g.") {
		t.Errorf("examples already described were appended: %q", got)
	}
	if got = flagDescription("Filter by stock symbol.", []string{"TSLA"}); strings.Contains(got, "e.g.") || strings.Contains(got, "TSLA") {
		t.Errorf("single sample value was appended: %q", got)
	}
	if got = flagDescription("Tokenization request status to be queried.", []string{"completed", "pending"}); strings.Contains(got, "e.g.") {
		t.Errorf("enum samples were appended: %q", got)
	}
}

func TestStripMarkdownLinks(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"no links here", "no links here"},
		{"[text](url)", "text"},
		{"before [text](url) after", "before text after"},
		{"[a](x) and [b](y)", "a and b"},
		{"unclosed [bracket", "unclosed [bracket"},
		{"[text](", "[text]("},
		{"[text](url) [no-url]", "text [no-url]"},
	}
	for _, tt := range tests {
		got := stripMarkdownLinks(tt.in)
		if got != tt.want {
			t.Errorf("stripMarkdownLinks(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func repeat(ch byte, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = ch
	}
	return string(b)
}

func repeatWord(w string, n int) string {
	s := ""
	for i := 0; i < n; i++ {
		if i > 0 {
			s += " "
		}
		s += w
	}
	return s
}
