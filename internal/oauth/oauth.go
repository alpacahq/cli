package oauth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

// Login performs the OAuth 2.0 authorization code flow:
// 1. Starts a localhost callback server
// 2. Opens the browser to Alpaca's authorize endpoint
// 3. Receives the authorization code via callback
// 4. Exchanges the code for an access token
func Login(env, scope string) (*Token, error) {
	state, err := randomState()
	if err != nil {
		return nil, fmt.Errorf("generating state: %w", err)
	}

	listener, port, err := listenOnAvailablePort()
	if err != nil {
		return nil, err
	}
	redirectURI := fmt.Sprintf("http://localhost:%s/callback", port)

	codeCh := make(chan callbackResult, 1)
	mux := http.NewServeMux()
	mux.HandleFunc("/callback", callbackHandler(state, codeCh))
	srv := &http.Server{Handler: mux}

	go func() { _ = srv.Serve(listener) }()
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		_ = srv.Shutdown(ctx)
	}()

	authURL := buildAuthorizeURL(redirectURI, state, env, scope)

	fmt.Printf("Opening browser to authorize...\n")
	if err := openBrowser(authURL); err != nil {
		fmt.Printf("Could not open browser automatically.\n")
		fmt.Printf("Open this URL in your browser:\n\n  %s\n\n", authURL)
	}
	fmt.Printf("Waiting for authorization (timeout: 2 minutes)...\n")

	select {
	case result := <-codeCh:
		if result.err != nil {
			return nil, result.err
		}
		return exchangeCode(result.code, redirectURI)
	case <-time.After(2 * time.Minute):
		return nil, fmt.Errorf("timed out waiting for authorization\nHint: make sure you completed the authorization in your browser")
	}
}

type callbackResult struct {
	code string
	err  error
}

func callbackHandler(expectedState string, ch chan<- callbackResult) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if errParam := r.URL.Query().Get("error"); errParam != "" {
			desc := r.URL.Query().Get("error_description")
			if desc == "" {
				desc = errParam
			}
			ch <- callbackResult{err: fmt.Errorf("authorization denied: %s", desc)}
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprint(w, errorPage(desc))
			return
		}

		code := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")

		if state != expectedState {
			ch <- callbackResult{err: fmt.Errorf("state mismatch (possible CSRF attack)")}
			http.Error(w, "state mismatch", http.StatusBadRequest)
			return
		}

		if code == "" {
			ch <- callbackResult{err: fmt.Errorf("no authorization code received")}
			http.Error(w, "missing code", http.StatusBadRequest)
			return
		}

		ch <- callbackResult{code: code}
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, successPage)
	}
}

func exchangeCode(code, redirectURI string) (*Token, error) {
	data := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"client_id":     {ClientID},
		"client_secret": {ClientSecret},
		"redirect_uri":  {redirectURI},
	}

	resp, err := (&http.Client{Timeout: 15 * time.Second}).Post(
		TokenURL,
		"application/x-www-form-urlencoded",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return nil, fmt.Errorf("token exchange failed: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading token response: %w", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("token exchange failed (HTTP %d): %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	var token Token
	if err := json.Unmarshal(body, &token); err != nil {
		return nil, fmt.Errorf("parsing token response: %w", err)
	}

	if token.AccessToken == "" {
		return nil, fmt.Errorf("token exchange returned empty access token")
	}

	return &token, nil
}

func buildAuthorizeURL(redirectURI, state, env, scope string) string {
	params := url.Values{
		"response_type": {"code"},
		"client_id":     {ClientID},
		"redirect_uri":  {redirectURI},
		"state":         {state},
	}
	if scope != "" {
		params.Set("scope", scope)
	}
	if env != "" {
		params.Set("env", env)
	}
	return AuthorizeURL + "?" + params.Encode()
}

func randomState() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func listenOnAvailablePort() (net.Listener, string, error) {
	for _, port := range CallbackPorts {
		l, err := net.Listen("tcp", "127.0.0.1:"+port)
		if err == nil {
			return l, port, nil
		}
	}
	return nil, "", fmt.Errorf("could not bind to any callback port (%s–%s)\nHint: free one of these ports or close other instances of `alpaca profile login`",
		CallbackPorts[0], CallbackPorts[len(CallbackPorts)-1])
}

func openBrowser(url string) error {
	switch runtime.GOOS {
	case "darwin":
		return exec.Command("open", url).Start()
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	default:
		return fmt.Errorf("unsupported platform %s", runtime.GOOS)
	}
}

const successPage = `<!DOCTYPE html>
<html>
<head><title>Alpaca CLI</title></head>
<body style="font-family: system-ui, sans-serif; display: flex; justify-content: center; align-items: center; height: 100vh; margin: 0; background: #f8f9fa;">
<div style="text-align: center;">
<h1 style="color: #1a1a2e;">&#10003; Authorization successful</h1>
<p style="color: #666;">You can close this tab and return to your terminal.</p>
</div>
</body>
</html>`

func errorPage(desc string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html>
<head><title>Alpaca CLI</title></head>
<body style="font-family: system-ui, sans-serif; display: flex; justify-content: center; align-items: center; height: 100vh; margin: 0; background: #f8f9fa;">
<div style="text-align: center;">
<h1 style="color: #c0392b;">&#10007; Authorization failed</h1>
<p style="color: #666;">%s</p>
<p style="color: #999;">Return to your terminal for details.</p>
</div>
</body>
</html>`, desc)
}
