package api

// Op describes a generated API operation. Passed to fetchCmd/attachCmd
// for automatic flag registration, help text, and required-flag validation.
type Op struct {
	Name         string
	Summary      string
	Long         string
	Example      string
	ReturnsArray bool
	Flags        []FlagDef
}

// FlagDef describes a CLI flag derived from the OpenAPI spec.
type FlagDef struct {
	Name        string   // kebab-case CLI flag name
	OASName     string   // original OAS property/parameter name
	Type        string   // "string", "bool", "int"
	Default     string
	Description string
	Completions []string // enum values for shell completion
	Required    bool     // true if OAS marks this parameter as required
	Source      string   // "path", "query", or "body"
}

// ResponseField describes a field in an API response.
type ResponseField struct {
	Name        string
	Type        string
	Description string
	EnumValues  []string
}
