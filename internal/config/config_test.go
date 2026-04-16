package config

import "testing"

func TestResolveBaseURL(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"", "https://paper-api.alpaca.markets"},
		{"paper", "https://paper-api.alpaca.markets"},
		{"PAPER", "https://paper-api.alpaca.markets"},
		{"live", "https://api.alpaca.markets"},
		{"LIVE", "https://api.alpaca.markets"},
		{"Live", "https://api.alpaca.markets"},
		{"https://custom.example.com/", "https://custom.example.com"},
		{"https://custom.example.com", "https://custom.example.com"},
	}
	for _, tc := range cases {
		got := ResolveBaseURL(tc.input)
		if got != tc.want {
			t.Errorf("ResolveBaseURL(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

func TestResolve(t *testing.T) {
	cases := []struct {
		values []string
		want   string
	}{
		{[]string{"", "", ""}, ""},
		{[]string{"", "b", "c"}, "b"},
		{[]string{"a", "b", "c"}, "a"},
		{[]string{"", "", "c"}, "c"},
	}
	for _, tc := range cases {
		got := resolve(tc.values...)
		if got != tc.want {
			t.Errorf("resolve(%v) = %q, want %q", tc.values, got, tc.want)
		}
	}
}

func TestResolved_HasCredentials(t *testing.T) {
	cases := []struct {
		name   string
		source Source
		want   bool
	}{
		{"env-apikey", SourceEnvAPIKey, true},
		{"profile-oauth", SourceProfileOAuth, true},
		{"profile-apikey", SourceProfileAPIKey, true},
		{"none", SourceNone, false},
	}
	for _, tc := range cases {
		r := &Resolved{Source: tc.source}
		if got := r.HasCredentials(); got != tc.want {
			t.Errorf("Source=%q HasCredentials()=%v, want %v", tc.source, got, tc.want)
		}
	}
}

func TestResolved_Validate(t *testing.T) {
	r := &Resolved{Source: SourceEnvAPIKey}
	if err := r.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	r2 := &Resolved{Source: SourceNone}
	if err := r2.Validate(); err == nil {
		t.Error("expected error for empty credentials")
	}
}

func TestResolved_IsOAuth(t *testing.T) {
	if !(&Resolved{Source: SourceProfileOAuth}).IsOAuth() {
		t.Error("profile-oauth should be IsOAuth")
	}
	if (&Resolved{Source: SourceEnvAPIKey}).IsOAuth() {
		t.Error("env-apikey should not be IsOAuth")
	}
	if (&Resolved{Source: SourceProfileAPIKey}).IsOAuth() {
		t.Error("profile-apikey should not be IsOAuth")
	}
}
