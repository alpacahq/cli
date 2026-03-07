package oauth

const (
	AuthorizeURL = "https://app.alpaca.markets/oauth/authorize"
	TokenURL     = "https://api.alpaca.markets/oauth/token"

	DefaultScopes = "account:write trading data"
)

// Ports the callback server tries, in order. Each must be whitelisted
// as http://localhost:{port}/callback on the OAuth client.
var CallbackPorts = []string{"41920", "41921", "41922", "41923", "41924"}

const (
	// ClientID and ClientSecret are public OAuth app identifiers, NOT
	// security credentials. Exposing them in source code is safe and
	// standard practice for native CLI applications. They cannot be used
	// to access any account without explicit user consent via the browser.
	//
	// Evidence:
	//  - RFC 8252 §8.5: static secrets in distributed apps "should not be
	//    treated as confidential secrets" — https://datatracker.ietf.org/doc/html/rfc8252#section-8.5
	//  - RFC 9700 (BCP 240): mandates PKCE, deprecates reliance on static
	//    client secrets — https://datatracker.ietf.org/doc/html/rfc9700
	//  - GitHub CLI embeds its secret identically — https://github.com/cli/oauth/issues/1
	//  - Google OAuth docs: "installed apps cannot keep secrets" —
	//    https://developers.google.com/identity/protocols/oauth2/native-app
	ClientID     = "3d2427aa1cf0863412d54e185c374d21"
	ClientSecret = "ff01f503caaaaf9a1e769576b7b6129a5a83d5ff"
)
