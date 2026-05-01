// Code generated from api/specs/trading-api.json; DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/alpacahq/cli/internal/client"
)

// TradingClient provides typed methods for the Trading API.
type TradingClient struct {
	Raw     *client.Client
	baseURL string
}

func NewTradingClient(raw *client.Client) *TradingClient {
	return &TradingClient{Raw: raw, baseURL: raw.BaseURL}
}

// GetAccount — Get Account
func (c *TradingClient) GetAccount() (*Account, error) {
	return unmarshal[Account](c.Raw.Do("GET", c.baseURL, "/v2/account", nil, nil))
}

// GetAccountActivities — Retrieve Account Activities
func (c *TradingClient) GetAccountActivities(params url.Values) (json.RawMessage, error) {
	return c.Raw.Do("GET", c.baseURL, "/v2/account/activities", params, nil)
}

// GetAccountActivitiesByActivityType — Retrieve Account Activities of Specific Type
func (c *TradingClient) GetAccountActivitiesByActivityType(ActivityType string, params url.Values) (json.RawMessage, error) {
	return c.Raw.Do("GET", c.baseURL, fmt.Sprintf("/v2/account/activities/%s", url.PathEscape(ActivityType)), params, nil)
}

// GetAccountConfig — Get Account Configurations
func (c *TradingClient) GetAccountConfig() (*AccountConfigurations, error) {
	return unmarshal[AccountConfigurations](c.Raw.Do("GET", c.baseURL, "/v2/account/configurations", nil, nil))
}

// PatchAccountConfig — Account Configurations
func (c *TradingClient) PatchAccountConfig(body *AccountConfigurations) (*AccountConfigurations, error) {
	return unmarshal[AccountConfigurations](c.Raw.Do("PATCH", c.baseURL, "/v2/account/configurations", nil, body))
}

// GetAccountPortfolioHistory — Get Account Portfolio History
func (c *TradingClient) GetAccountPortfolioHistory(params url.Values) (*PortfolioHistory, error) {
	return unmarshal[PortfolioHistory](c.Raw.Do("GET", c.baseURL, "/v2/account/portfolio/history", params, nil))
}

// GetV2Assets — Get Assets
func (c *TradingClient) GetV2Assets(params url.Values) ([]Assets, error) {
	return unmarshalSlice[Assets](c.Raw.Do("GET", c.baseURL, "/v2/assets", params, nil))
}

// UsCorporates — Get US corporates
func (c *TradingClient) UsCorporates(params url.Values) (*UsCorporatesResp, error) {
	return unmarshal[UsCorporatesResp](c.Raw.Do("GET", c.baseURL, "/v2/assets/fixed_income/us_corporates", params, nil))
}

// UsTreasuries — Get US treasuries
func (c *TradingClient) UsTreasuries(params url.Values) (*UsTreasuriesResp, error) {
	return unmarshal[UsTreasuriesResp](c.Raw.Do("GET", c.baseURL, "/v2/assets/fixed_income/us_treasuries", params, nil))
}

// GetV2AssetsSymbolOrAssetID — Get an Asset by ID or Symbol
func (c *TradingClient) GetV2AssetsSymbolOrAssetID(SymbolOrAssetID string) (*Assets, error) {
	return unmarshal[Assets](c.Raw.Do("GET", c.baseURL, fmt.Sprintf("/v2/assets/%s", url.PathEscape(SymbolOrAssetID)), nil, nil))
}

// LegacyCalendar — Get US Market Calendar
func (c *TradingClient) LegacyCalendar(params url.Values) (json.RawMessage, error) {
	return c.Raw.Do("GET", c.baseURL, "/v2/calendar", params, nil)
}

// LegacyClock — Get US Market Clock
func (c *TradingClient) LegacyClock() (*LegacyClock, error) {
	return unmarshal[LegacyClock](c.Raw.Do("GET", c.baseURL, "/v2/clock", nil, nil))
}

// GetV2CorporateActionsAnnouncements — Retrieve Announcements
func (c *TradingClient) GetV2CorporateActionsAnnouncements(params url.Values) (json.RawMessage, error) {
	return c.Raw.Do("GET", c.baseURL, "/v2/corporate_actions/announcements", params, nil)
}

// GetV2CorporateActionsAnnouncementsID — Retrieve a Specific Announcement
func (c *TradingClient) GetV2CorporateActionsAnnouncementsID(ID string) (json.RawMessage, error) {
	return c.Raw.Do("GET", c.baseURL, fmt.Sprintf("/v2/corporate_actions/announcements/%s", url.PathEscape(ID)), nil, nil)
}

// GetOptionsContracts — Get Option Contracts
func (c *TradingClient) GetOptionsContracts(params url.Values) (json.RawMessage, error) {
	return c.Raw.Do("GET", c.baseURL, "/v2/options/contracts", params, nil)
}

// GetOptionContractSymbolOrID — Get an option contract by ID or Symbol
func (c *TradingClient) GetOptionContractSymbolOrID(SymbolOrID string) (*OptionContract, error) {
	return unmarshal[OptionContract](c.Raw.Do("GET", c.baseURL, fmt.Sprintf("/v2/options/contracts/%s", url.PathEscape(SymbolOrID)), nil, nil))
}

// GetAllOrders — Get All Orders
func (c *TradingClient) GetAllOrders(params url.Values) ([]Order, error) {
	return unmarshalSlice[Order](c.Raw.Do("GET", c.baseURL, "/v2/orders", params, nil))
}

type PostOrderRequest struct {
	AdvancedInstructions AdvancedInstructions `json:"advanced_instructions,omitempty"`
	ClientOrderID        string               `json:"client_order_id,omitempty"`
	ExtendedHours        bool                 `json:"extended_hours,omitempty"`
	Legs                 []MLegOrderLeg       `json:"legs,omitempty"`
	LimitPrice           string               `json:"limit_price,omitempty"`
	Notional             string               `json:"notional,omitempty"`
	OrderClass           OrderClass           `json:"order_class,omitempty"`
	PositionIntent       PositionIntent       `json:"position_intent,omitempty"`
	Qty                  string               `json:"qty,omitempty"`
	Side                 OrderSide            `json:"side,omitempty"`
	StopLoss             map[string]any       `json:"stop_loss,omitempty"`
	StopPrice            string               `json:"stop_price,omitempty"`
	Symbol               string               `json:"symbol,omitempty"`
	TakeProfit           map[string]any       `json:"take_profit,omitempty"`
	TimeInForce          TimeInForce          `json:"time_in_force"`
	TrailPercent         string               `json:"trail_percent,omitempty"`
	TrailPrice           string               `json:"trail_price,omitempty"`
	Type                 OrderType            `json:"type"`
}

// PostOrder — Create an Order
func (c *TradingClient) PostOrder(body *PostOrderRequest) (*Order, error) {
	return unmarshal[Order](c.Raw.Do("POST", c.baseURL, "/v2/orders", nil, body))
}

// DeleteAllOrders — Delete All Orders
func (c *TradingClient) DeleteAllOrders() ([]CanceledOrderResponse, error) {
	return unmarshalSlice[CanceledOrderResponse](c.Raw.Do("DELETE", c.baseURL, "/v2/orders", nil, nil))
}

// GetOrderByOrderID — Get Order by ID
func (c *TradingClient) GetOrderByOrderID(OrderID string, params url.Values) (*Order, error) {
	return unmarshal[Order](c.Raw.Do("GET", c.baseURL, fmt.Sprintf("/v2/orders/%s", url.PathEscape(OrderID)), params, nil))
}

// PatchOrderByOrderID — Replace Order by ID
func (c *TradingClient) PatchOrderByOrderID(OrderID string, body *PatchOrderRequest) (*Order, error) {
	return unmarshal[Order](c.Raw.Do("PATCH", c.baseURL, fmt.Sprintf("/v2/orders/%s", url.PathEscape(OrderID)), nil, body))
}

// DeleteOrderByOrderID — Delete Order by ID
func (c *TradingClient) DeleteOrderByOrderID(OrderID string) (json.RawMessage, error) {
	return c.Raw.Do("DELETE", c.baseURL, fmt.Sprintf("/v2/orders/%s", url.PathEscape(OrderID)), nil, nil)
}

// GetOrderByClientOrderID — Get Order by Client Order ID
func (c *TradingClient) GetOrderByClientOrderID(params url.Values) (*Order, error) {
	return unmarshal[Order](c.Raw.Do("GET", c.baseURL, "/v2/orders:by_client_order_id", params, nil))
}

// GetCryptoPerpAccountVitals — Retrieve Account Vitals
func (c *TradingClient) GetCryptoPerpAccountVitals() (json.RawMessage, error) {
	return c.Raw.Do("GET", c.baseURL, "/v2/perpetuals/account_vitals", nil, nil)
}

// GetCryptoPerpAccountLeverage — Get Account Leverage for an Asset
func (c *TradingClient) GetCryptoPerpAccountLeverage(params url.Values) (json.RawMessage, error) {
	return c.Raw.Do("GET", c.baseURL, "/v2/perpetuals/leverage", params, nil)
}

// SetCryptoPerpAccountLeverage — Set Account Leverage for an Asset
func (c *TradingClient) SetCryptoPerpAccountLeverage(params url.Values) (json.RawMessage, error) {
	return c.Raw.Do("POST", c.baseURL, "/v2/perpetuals/leverage", params, nil)
}

// ListCryptoPerpFundingWallets — Retrieve Crypto Funding Wallets
func (c *TradingClient) ListCryptoPerpFundingWallets(params url.Values) (*CryptoWallet, error) {
	return unmarshal[CryptoWallet](c.Raw.Do("GET", c.baseURL, "/v2/perpetuals/wallets", params, nil))
}

// GetCryptoPerpTransferEstimate — Returns the estimated gas fee for a proposed transaction
func (c *TradingClient) GetCryptoPerpTransferEstimate(params url.Values) (json.RawMessage, error) {
	return c.Raw.Do("GET", c.baseURL, "/v2/perpetuals/wallets/fees/estimate", params, nil)
}

// ListCryptoPerpFundingTransfers — Retrieve Crypto Funding Transfers
func (c *TradingClient) ListCryptoPerpFundingTransfers() (*CryptoTransfer, error) {
	return unmarshal[CryptoTransfer](c.Raw.Do("GET", c.baseURL, "/v2/perpetuals/wallets/transfers", nil, nil))
}

// CreateCryptoPerpTransferForAccount — Request a New Withdrawal
func (c *TradingClient) CreateCryptoPerpTransferForAccount(body *CreateCryptoTransferRequest) (*CryptoTransfer, error) {
	return unmarshal[CryptoTransfer](c.Raw.Do("POST", c.baseURL, "/v2/perpetuals/wallets/transfers", nil, body))
}

// GetCryptoPerpFundingTransfer — Retrieve a Crypto Funding Transfer
func (c *TradingClient) GetCryptoPerpFundingTransfer(TransferID string) (*CryptoTransfer, error) {
	return unmarshal[CryptoTransfer](c.Raw.Do("GET", c.baseURL, fmt.Sprintf("/v2/perpetuals/wallets/transfers/%s", url.PathEscape(TransferID)), nil, nil))
}

// ListWhitelistedPerpAddress — An array of whitelisted addresses
func (c *TradingClient) ListWhitelistedPerpAddress() (*WhitelistedAddress, error) {
	return unmarshal[WhitelistedAddress](c.Raw.Do("GET", c.baseURL, "/v2/perpetuals/wallets/whitelists", nil, nil))
}

type CreateWhitelistedPerpAddressRequest struct {
	Address string `json:"address,omitempty"`
	Asset   string `json:"asset,omitempty"`
}

// CreateWhitelistedPerpAddress — Request a new whitelisted address
func (c *TradingClient) CreateWhitelistedPerpAddress(body *CreateWhitelistedPerpAddressRequest) (*WhitelistedAddress, error) {
	return unmarshal[WhitelistedAddress](c.Raw.Do("POST", c.baseURL, "/v2/perpetuals/wallets/whitelists", nil, body))
}

// DeleteWhitelistedPerpAddress — Delete a whitelisted address
func (c *TradingClient) DeleteWhitelistedPerpAddress(WhitelistedAddressID string) (json.RawMessage, error) {
	return c.Raw.Do("DELETE", c.baseURL, fmt.Sprintf("/v2/perpetuals/wallets/whitelists/%s", url.PathEscape(WhitelistedAddressID)), nil, nil)
}

// GetAllOpenPositions — All Open Positions
func (c *TradingClient) GetAllOpenPositions() ([]Position, error) {
	return unmarshalSlice[Position](c.Raw.Do("GET", c.baseURL, "/v2/positions", nil, nil))
}

// DeleteAllOpenPositions — Close All Positions
func (c *TradingClient) DeleteAllOpenPositions(params url.Values) ([]PositionClosedReponse, error) {
	return unmarshalSlice[PositionClosedReponse](c.Raw.Do("DELETE", c.baseURL, "/v2/positions", params, nil))
}

// GetOpenPosition — Get an Open Position
func (c *TradingClient) GetOpenPosition(SymbolOrAssetID string) (*Position, error) {
	return unmarshal[Position](c.Raw.Do("GET", c.baseURL, fmt.Sprintf("/v2/positions/%s", url.PathEscape(SymbolOrAssetID)), nil, nil))
}

// DeleteOpenPosition — Close a Position
func (c *TradingClient) DeleteOpenPosition(SymbolOrAssetID string, params url.Values) (*Order, error) {
	return unmarshal[Order](c.Raw.Do("DELETE", c.baseURL, fmt.Sprintf("/v2/positions/%s", url.PathEscape(SymbolOrAssetID)), params, nil))
}

// OptionDoNotExercise — Do Not Exercise an Options Position
func (c *TradingClient) OptionDoNotExercise(SymbolOrContractID string) (json.RawMessage, error) {
	return c.Raw.Do("POST", c.baseURL, fmt.Sprintf("/v2/positions/%s/do-not-exercise", url.PathEscape(SymbolOrContractID)), nil, nil)
}

// OptionExercise — Exercise an Options Position
func (c *TradingClient) OptionExercise(SymbolOrContractID string) (json.RawMessage, error) {
	return c.Raw.Do("POST", c.baseURL, fmt.Sprintf("/v2/positions/%s/exercise", url.PathEscape(SymbolOrContractID)), nil, nil)
}

// PostTokenizationMint — Mint a Tokenized Asset
func (c *TradingClient) PostTokenizationMint(body *TokenizationMintRequest) (*TokenizationMintResponse, error) {
	return unmarshal[TokenizationMintResponse](c.Raw.Do("POST", c.baseURL, "/v2/tokenization/mint", nil, body))
}

// GetTokenizationRequests — List Tokenization Requests
func (c *TradingClient) GetTokenizationRequests(params url.Values) ([]TokenizationRequest, error) {
	return unmarshalSlice[TokenizationRequest](c.Raw.Do("GET", c.baseURL, "/v2/tokenization/requests", params, nil))
}

// ListCryptoFundingWallets — Retrieve Crypto Funding Wallets
func (c *TradingClient) ListCryptoFundingWallets(params url.Values) (*CryptoWallet, error) {
	return unmarshal[CryptoWallet](c.Raw.Do("GET", c.baseURL, "/v2/wallets", params, nil))
}

// GetCryptoTransferEstimate — Returns the estimated gas fee for a proposed transaction.
func (c *TradingClient) GetCryptoTransferEstimate(params url.Values) (json.RawMessage, error) {
	return c.Raw.Do("GET", c.baseURL, "/v2/wallets/fees/estimate", params, nil)
}

// ListCryptoFundingTransfers — Retrieve Crypto Funding Transfers
func (c *TradingClient) ListCryptoFundingTransfers() (*CryptoTransfer, error) {
	return unmarshal[CryptoTransfer](c.Raw.Do("GET", c.baseURL, "/v2/wallets/transfers", nil, nil))
}

// CreateCryptoTransferForAccount — Request a New Withdrawal
func (c *TradingClient) CreateCryptoTransferForAccount(body *CreateCryptoTransferRequest) (*CryptoTransfer, error) {
	return unmarshal[CryptoTransfer](c.Raw.Do("POST", c.baseURL, "/v2/wallets/transfers", nil, body))
}

// GetCryptoFundingTransfer — Retrieve a Crypto Funding Transfer
func (c *TradingClient) GetCryptoFundingTransfer(TransferID string) (*CryptoTransfer, error) {
	return unmarshal[CryptoTransfer](c.Raw.Do("GET", c.baseURL, fmt.Sprintf("/v2/wallets/transfers/%s", url.PathEscape(TransferID)), nil, nil))
}

// ListWhitelistedAddress — An array of whitelisted addresses
func (c *TradingClient) ListWhitelistedAddress() (*WhitelistedAddress, error) {
	return unmarshal[WhitelistedAddress](c.Raw.Do("GET", c.baseURL, "/v2/wallets/whitelists", nil, nil))
}

type CreateWhitelistedAddressRequest struct {
	Address string `json:"address,omitempty"`
	Asset   string `json:"asset,omitempty"`
}

// CreateWhitelistedAddress — Request a new whitelisted address
func (c *TradingClient) CreateWhitelistedAddress(body *CreateWhitelistedAddressRequest) (*WhitelistedAddress, error) {
	return unmarshal[WhitelistedAddress](c.Raw.Do("POST", c.baseURL, "/v2/wallets/whitelists", nil, body))
}

// DeleteWhitelistedAddress — Delete a whitelisted address
func (c *TradingClient) DeleteWhitelistedAddress(WhitelistedAddressID string) (json.RawMessage, error) {
	return c.Raw.Do("DELETE", c.baseURL, fmt.Sprintf("/v2/wallets/whitelists/%s", url.PathEscape(WhitelistedAddressID)), nil, nil)
}

// GetWatchlists — Get All Watchlists
func (c *TradingClient) GetWatchlists() ([]WatchlistWithoutAsset, error) {
	return unmarshalSlice[WatchlistWithoutAsset](c.Raw.Do("GET", c.baseURL, "/v2/watchlists", nil, nil))
}

// PostWatchlist — Create Watchlist
func (c *TradingClient) PostWatchlist(body *UpdateWatchlistRequest) (*Watchlist, error) {
	return unmarshal[Watchlist](c.Raw.Do("POST", c.baseURL, "/v2/watchlists", nil, body))
}

// GetWatchlistByID — Get Watchlist by ID
func (c *TradingClient) GetWatchlistByID(WatchlistID string) (*Watchlist, error) {
	return unmarshal[Watchlist](c.Raw.Do("GET", c.baseURL, fmt.Sprintf("/v2/watchlists/%s", url.PathEscape(WatchlistID)), nil, nil))
}

// AddAssetToWatchlist — Add Asset to Watchlist
func (c *TradingClient) AddAssetToWatchlist(WatchlistID string, body *AddAssetToWatchlistRequest) (*Watchlist, error) {
	return unmarshal[Watchlist](c.Raw.Do("POST", c.baseURL, fmt.Sprintf("/v2/watchlists/%s", url.PathEscape(WatchlistID)), nil, body))
}

// UpdateWatchlistByID — Update Watchlist By Id
func (c *TradingClient) UpdateWatchlistByID(WatchlistID string, body *UpdateWatchlistRequest) (*Watchlist, error) {
	return unmarshal[Watchlist](c.Raw.Do("PUT", c.baseURL, fmt.Sprintf("/v2/watchlists/%s", url.PathEscape(WatchlistID)), nil, body))
}

// DeleteWatchlistByID — Delete Watchlist By Id
func (c *TradingClient) DeleteWatchlistByID(WatchlistID string) (json.RawMessage, error) {
	return c.Raw.Do("DELETE", c.baseURL, fmt.Sprintf("/v2/watchlists/%s", url.PathEscape(WatchlistID)), nil, nil)
}

// RemoveAssetFromWatchlist — Delete Symbol from Watchlist
func (c *TradingClient) RemoveAssetFromWatchlist(WatchlistID string, Symbol string) (*Watchlist, error) {
	return unmarshal[Watchlist](c.Raw.Do("DELETE", c.baseURL, fmt.Sprintf("/v2/watchlists/%s/%s", url.PathEscape(WatchlistID), url.PathEscape(Symbol)), nil, nil))
}

// GetWatchlistByName — Get Watchlist by Name
func (c *TradingClient) GetWatchlistByName(params url.Values) (*Watchlist, error) {
	return unmarshal[Watchlist](c.Raw.Do("GET", c.baseURL, "/v2/watchlists:by_name", params, nil))
}

type AddAssetToWatchlistByNameRequest struct {
	Symbol string `json:"symbol,omitempty"`
}

// AddAssetToWatchlistByName — Add Asset to Watchlist By Name
func (c *TradingClient) AddAssetToWatchlistByName(params url.Values, body *AddAssetToWatchlistByNameRequest) (*Watchlist, error) {
	return unmarshal[Watchlist](c.Raw.Do("POST", c.baseURL, "/v2/watchlists:by_name", params, body))
}

// UpdateWatchlistByName — Update Watchlist By Name
func (c *TradingClient) UpdateWatchlistByName(params url.Values, body *UpdateWatchlistRequest) (*Watchlist, error) {
	return unmarshal[Watchlist](c.Raw.Do("PUT", c.baseURL, "/v2/watchlists:by_name", params, body))
}

// DeleteWatchlistByName — Delete Watchlist By Name
func (c *TradingClient) DeleteWatchlistByName(params url.Values) (json.RawMessage, error) {
	return c.Raw.Do("DELETE", c.baseURL, "/v2/watchlists:by_name", params, nil)
}

// SubscribeToActivitiesSSE — Subscribe to Activity Events (SSE)
func (c *TradingClient) SubscribeToActivitiesSSE(params url.Values) (json.RawMessage, error) {
	return c.Raw.Do("GET", c.baseURL, "/v2beta1/events/activities", params, nil)
}

// Calendar — Get Market Calendar
func (c *TradingClient) Calendar(Market string, params url.Values) (*PublicCalendarResp, error) {
	return unmarshal[PublicCalendarResp](c.Raw.Do("GET", c.baseURL, fmt.Sprintf("/v3/calendar/%s", url.PathEscape(Market)), params, nil))
}

// Clock — Get Market Clock
func (c *TradingClient) Clock(params url.Values) (*ClockResp, error) {
	return unmarshal[ClockResp](c.Raw.Do("GET", c.baseURL, "/v3/clock", params, nil))
}

func (r *CreateCryptoTransferRequest) Validate() error {
	var missing []string
	if r.Address == "" {
		missing = append(missing, "address")
	}
	if r.Amount == "" {
		missing = append(missing, "amount")
	}
	if r.Asset == "" {
		missing = append(missing, "asset")
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing required fields: %s", strings.Join(missing, ", "))
	}
	return nil
}

func (r *TokenizationMintRequest) Validate() error {
	var missing []string
	if r.Qty == "" {
		missing = append(missing, "qty")
	}
	if r.UnderlyingSymbol == "" {
		missing = append(missing, "underlying_symbol")
	}
	if r.WalletAddress == "" {
		missing = append(missing, "wallet_address")
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing required fields: %s", strings.Join(missing, ", "))
	}
	return nil
}

func (r *UpdateWatchlistRequest) Validate() error {
	var missing []string
	if r.Name == "" {
		missing = append(missing, "name")
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing required fields: %s", strings.Join(missing, ", "))
	}
	return nil
}
