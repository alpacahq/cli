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
	// ClientID and ClientSecret identify this CLI as Alpaca's first-party
	// OAuth app. They are embedded in the distributed binary and must be
	// treated as PUBLIC — any user can extract them, and any app could
	// reuse them to initiate an OAuth flow posing as this CLI.
	//
	// This is standard for native CLI apps ("public clients" per RFC 8252)
	// because the secret is not a security boundary — access is gated by:
	//  1. User consent via Alpaca's browser UI.
	//  2. Server-side redirect URI validation (only pre-registered
	//     localhost callback ports receive tokens).
	//  3. The state parameter (CSRF protection).
	//
	// OAuth login is restricted to paper trading until the flow is
	// hardened with PKCE or Device Authorization Grant (RFC 8628).
	// Live trading requires API keys. See loginWithOAuth in auth.go.
	//
	// References:
	//  - RFC 8252 §8.5 — https://datatracker.ietf.org/doc/html/rfc8252#section-8.5
	//  - GitHub CLI embeds its secret identically — https://github.com/cli/oauth/issues/1
	ClientID     = "3d2427aa1cf0863412d54e185c374d21"
	ClientSecret = "ff01f503caaaaf9a1e769576b7b6129a5a83d5ff"
)
