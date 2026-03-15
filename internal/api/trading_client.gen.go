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
	Raw *client.Client
}

func NewTradingClient(raw *client.Client) *TradingClient {
	return &TradingClient{Raw: raw}
}

// GetAccount — Get Account
func (c *TradingClient) GetAccount() (*Account, error) {
	path := "/v2/account"
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var result Account
	return &result, json.Unmarshal(data, &result)
}

// GetAccountActivities — Retrieve Account Activities
func (c *TradingClient) GetAccountActivities(params url.Values) (json.RawMessage, error) {
	path := "/v2/account/activities"
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetAccountActivitiesByActivityType — Retrieve Account Activities of Specific Type
func (c *TradingClient) GetAccountActivitiesByActivityType(ActivityType string, params url.Values) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/account/activities/%s", url.PathEscape(ActivityType))
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetAccountConfig — Get Account Configurations
func (c *TradingClient) GetAccountConfig() (*AccountConfigurations, error) {
	path := "/v2/account/configurations"
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var result AccountConfigurations
	return &result, json.Unmarshal(data, &result)
}

// PatchAccountConfig — Account Configurations
func (c *TradingClient) PatchAccountConfig(body *AccountConfigurations) (*AccountConfigurations, error) {
	path := "/v2/account/configurations"
	data, err := c.Raw.Patch(path, nil, body)
	if err != nil {
		return nil, err
	}
	var result AccountConfigurations
	return &result, json.Unmarshal(data, &result)
}

// GetAccountPortfolioHistory — Get Account Portfolio History
func (c *TradingClient) GetAccountPortfolioHistory(params url.Values) (*PortfolioHistory, error) {
	path := "/v2/account/portfolio/history"
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	var result PortfolioHistory
	return &result, json.Unmarshal(data, &result)
}

// GetV2Assets — Get Assets
func (c *TradingClient) GetV2Assets(params url.Values) ([]Assets, error) {
	path := "/v2/assets"
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	var result []Assets
	return result, json.Unmarshal(data, &result)
}

// UsCorporates — Get US corporates
func (c *TradingClient) UsCorporates(params url.Values) (*UsCorporatesResp, error) {
	path := "/v2/assets/fixed_income/us_corporates"
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	var result UsCorporatesResp
	return &result, json.Unmarshal(data, &result)
}

// UsTreasuries — Get US treasuries
func (c *TradingClient) UsTreasuries(params url.Values) (*UsTreasuriesResp, error) {
	path := "/v2/assets/fixed_income/us_treasuries"
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	var result UsTreasuriesResp
	return &result, json.Unmarshal(data, &result)
}

// GetV2AssetsSymbolOrAssetID — Get an Asset by ID or Symbol
func (c *TradingClient) GetV2AssetsSymbolOrAssetID(SymbolOrAssetID string) (*Assets, error) {
	path := fmt.Sprintf("/v2/assets/%s", url.PathEscape(SymbolOrAssetID))
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var result Assets
	return &result, json.Unmarshal(data, &result)
}

// LegacyCalendar — Get US Market Calendar
func (c *TradingClient) LegacyCalendar(params url.Values) (json.RawMessage, error) {
	path := "/v2/calendar"
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// LegacyClock — Get US Market Clock
func (c *TradingClient) LegacyClock() (*LegacyClock, error) {
	path := "/v2/clock"
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var result LegacyClock
	return &result, json.Unmarshal(data, &result)
}

// GetV2CorporateActionsAnnouncements — Retrieve Announcements
func (c *TradingClient) GetV2CorporateActionsAnnouncements(params url.Values) (json.RawMessage, error) {
	path := "/v2/corporate_actions/announcements"
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetV2CorporateActionsAnnouncementsID — Retrieve a Specific Announcement
func (c *TradingClient) GetV2CorporateActionsAnnouncementsID(ID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/corporate_actions/announcements/%s", url.PathEscape(ID))
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetOptionsContracts — Get Option Contracts
func (c *TradingClient) GetOptionsContracts(params url.Values) (json.RawMessage, error) {
	path := "/v2/options/contracts"
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetOptionContractSymbolOrID — Get an option contract by ID or Symbol
func (c *TradingClient) GetOptionContractSymbolOrID(SymbolOrID string) (*OptionContract, error) {
	path := fmt.Sprintf("/v2/options/contracts/%s", url.PathEscape(SymbolOrID))
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var result OptionContract
	return &result, json.Unmarshal(data, &result)
}

// GetAllOrders — Get All Orders
func (c *TradingClient) GetAllOrders(params url.Values) ([]Order, error) {
	path := "/v2/orders"
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	var result []Order
	return result, json.Unmarshal(data, &result)
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
	path := "/v2/orders"
	data, err := c.Raw.Post(path, nil, body)
	if err != nil {
		return nil, err
	}
	var result Order
	return &result, json.Unmarshal(data, &result)
}

// DeleteAllOrders — Delete All Orders
func (c *TradingClient) DeleteAllOrders() ([]CanceledOrderResponse, error) {
	path := "/v2/orders"
	data, err := c.Raw.Delete(path, nil)
	if err != nil {
		return nil, err
	}
	var result []CanceledOrderResponse
	return result, json.Unmarshal(data, &result)
}

// GetOrderByOrderID — Get Order by ID
func (c *TradingClient) GetOrderByOrderID(OrderID string, params url.Values) (*Order, error) {
	path := fmt.Sprintf("/v2/orders/%s", url.PathEscape(OrderID))
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	var result Order
	return &result, json.Unmarshal(data, &result)
}

// PatchOrderByOrderID — Replace Order by ID
func (c *TradingClient) PatchOrderByOrderID(OrderID string, body *PatchOrderRequest) (*Order, error) {
	path := fmt.Sprintf("/v2/orders/%s", url.PathEscape(OrderID))
	data, err := c.Raw.Patch(path, nil, body)
	if err != nil {
		return nil, err
	}
	var result Order
	return &result, json.Unmarshal(data, &result)
}

// DeleteOrderByOrderID — Delete Order by ID
func (c *TradingClient) DeleteOrderByOrderID(OrderID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/orders/%s", url.PathEscape(OrderID))
	data, err := c.Raw.Delete(path, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetOrderByClientOrderID — Get Order by Client Order ID
func (c *TradingClient) GetOrderByClientOrderID(params url.Values) (*Order, error) {
	path := "/v2/orders:by_client_order_id"
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	var result Order
	return &result, json.Unmarshal(data, &result)
}

// GetCryptoPerpAccountVitals — Retrieve Account Vitals
func (c *TradingClient) GetCryptoPerpAccountVitals() (json.RawMessage, error) {
	path := "/v2/perpetuals/account_vitals"
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetCryptoPerpAccountLeverage — Get Account Leverage for an Asset
func (c *TradingClient) GetCryptoPerpAccountLeverage(params url.Values) (json.RawMessage, error) {
	path := "/v2/perpetuals/leverage"
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// SetCryptoPerpAccountLeverage — Set Account Leverage for an Asset
func (c *TradingClient) SetCryptoPerpAccountLeverage(params url.Values) (json.RawMessage, error) {
	path := "/v2/perpetuals/leverage"
	data, err := c.Raw.Post(path, params, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ListCryptoPerpFundingWallets — Retrieve Crypto Funding Wallets
func (c *TradingClient) ListCryptoPerpFundingWallets(params url.Values) (*CryptoWallet, error) {
	path := "/v2/perpetuals/wallets"
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	var result CryptoWallet
	return &result, json.Unmarshal(data, &result)
}

// GetCryptoPerpTransferEstimate — Returns the estimated gas fee for a proposed transaction
func (c *TradingClient) GetCryptoPerpTransferEstimate(params url.Values) (json.RawMessage, error) {
	path := "/v2/perpetuals/wallets/fees/estimate"
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ListCryptoPerpFundingTransfers — Retrieve Crypto Funding Transfers
func (c *TradingClient) ListCryptoPerpFundingTransfers() (*CryptoTransfer, error) {
	path := "/v2/perpetuals/wallets/transfers"
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var result CryptoTransfer
	return &result, json.Unmarshal(data, &result)
}

// CreateCryptoPerpTransferForAccount — Request a New Withdrawal
func (c *TradingClient) CreateCryptoPerpTransferForAccount(body *CreateCryptoTransferRequest) (*CryptoTransfer, error) {
	path := "/v2/perpetuals/wallets/transfers"
	data, err := c.Raw.Post(path, nil, body)
	if err != nil {
		return nil, err
	}
	var result CryptoTransfer
	return &result, json.Unmarshal(data, &result)
}

// GetCryptoPerpFundingTransfer — Retrieve a Crypto Funding Transfer
func (c *TradingClient) GetCryptoPerpFundingTransfer(TransferID string) (*CryptoTransfer, error) {
	path := fmt.Sprintf("/v2/perpetuals/wallets/transfers/%s", url.PathEscape(TransferID))
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var result CryptoTransfer
	return &result, json.Unmarshal(data, &result)
}

// ListWhitelistedPerpAddress — An array of whitelisted addresses
func (c *TradingClient) ListWhitelistedPerpAddress() (*WhitelistedAddress, error) {
	path := "/v2/perpetuals/wallets/whitelists"
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var result WhitelistedAddress
	return &result, json.Unmarshal(data, &result)
}

type CreateWhitelistedPerpAddressRequest struct {
	Address string `json:"address,omitempty"`
	Asset   string `json:"asset,omitempty"`
}

// CreateWhitelistedPerpAddress — Request a new whitelisted address
func (c *TradingClient) CreateWhitelistedPerpAddress(body *CreateWhitelistedPerpAddressRequest) (*WhitelistedAddress, error) {
	path := "/v2/perpetuals/wallets/whitelists"
	data, err := c.Raw.Post(path, nil, body)
	if err != nil {
		return nil, err
	}
	var result WhitelistedAddress
	return &result, json.Unmarshal(data, &result)
}

// DeleteWhitelistedPerpAddress — Delete a whitelisted address
func (c *TradingClient) DeleteWhitelistedPerpAddress(WhitelistedAddressID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/perpetuals/wallets/whitelists/%s", url.PathEscape(WhitelistedAddressID))
	data, err := c.Raw.Delete(path, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetAllOpenPositions — All Open Positions
func (c *TradingClient) GetAllOpenPositions() ([]Position, error) {
	path := "/v2/positions"
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var result []Position
	return result, json.Unmarshal(data, &result)
}

// DeleteAllOpenPositions — Close All Positions
func (c *TradingClient) DeleteAllOpenPositions(params url.Values) ([]PositionClosedReponse, error) {
	path := "/v2/positions"
	data, err := c.Raw.Delete(path, params)
	if err != nil {
		return nil, err
	}
	var result []PositionClosedReponse
	return result, json.Unmarshal(data, &result)
}

// GetOpenPosition — Get an Open Position
func (c *TradingClient) GetOpenPosition(SymbolOrAssetID string) (*Position, error) {
	path := fmt.Sprintf("/v2/positions/%s", url.PathEscape(SymbolOrAssetID))
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var result Position
	return &result, json.Unmarshal(data, &result)
}

// DeleteOpenPosition — Close a Position
func (c *TradingClient) DeleteOpenPosition(SymbolOrAssetID string, params url.Values) (*Order, error) {
	path := fmt.Sprintf("/v2/positions/%s", url.PathEscape(SymbolOrAssetID))
	data, err := c.Raw.Delete(path, params)
	if err != nil {
		return nil, err
	}
	var result Order
	return &result, json.Unmarshal(data, &result)
}

// OptionDoNotExercise — Do Not Exercise an Options Position
func (c *TradingClient) OptionDoNotExercise(SymbolOrContractID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/positions/%s/do-not-exercise", url.PathEscape(SymbolOrContractID))
	data, err := c.Raw.Post(path, nil, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// OptionExercise — Exercise an Options Position
func (c *TradingClient) OptionExercise(SymbolOrContractID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/positions/%s/exercise", url.PathEscape(SymbolOrContractID))
	data, err := c.Raw.Post(path, nil, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ListCryptoFundingWallets — Retrieve Crypto Funding Wallets
func (c *TradingClient) ListCryptoFundingWallets(params url.Values) (*CryptoWallet, error) {
	path := "/v2/wallets"
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	var result CryptoWallet
	return &result, json.Unmarshal(data, &result)
}

// GetCryptoTransferEstimate — Returns the estimated gas fee for a proposed transaction.
func (c *TradingClient) GetCryptoTransferEstimate(params url.Values) (json.RawMessage, error) {
	path := "/v2/wallets/fees/estimate"
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ListCryptoFundingTransfers — Retrieve Crypto Funding Transfers
func (c *TradingClient) ListCryptoFundingTransfers() (*CryptoTransfer, error) {
	path := "/v2/wallets/transfers"
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var result CryptoTransfer
	return &result, json.Unmarshal(data, &result)
}

// CreateCryptoTransferForAccount — Request a New Withdrawal
func (c *TradingClient) CreateCryptoTransferForAccount(body *CreateCryptoTransferRequest) (*CryptoTransfer, error) {
	path := "/v2/wallets/transfers"
	data, err := c.Raw.Post(path, nil, body)
	if err != nil {
		return nil, err
	}
	var result CryptoTransfer
	return &result, json.Unmarshal(data, &result)
}

// GetCryptoFundingTransfer — Retrieve a Crypto Funding Transfer
func (c *TradingClient) GetCryptoFundingTransfer(TransferID string) (*CryptoTransfer, error) {
	path := fmt.Sprintf("/v2/wallets/transfers/%s", url.PathEscape(TransferID))
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var result CryptoTransfer
	return &result, json.Unmarshal(data, &result)
}

// ListWhitelistedAddress — An array of whitelisted addresses
func (c *TradingClient) ListWhitelistedAddress() (*WhitelistedAddress, error) {
	path := "/v2/wallets/whitelists"
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var result WhitelistedAddress
	return &result, json.Unmarshal(data, &result)
}

type CreateWhitelistedAddressRequest struct {
	Address string `json:"address,omitempty"`
	Asset   string `json:"asset,omitempty"`
}

// CreateWhitelistedAddress — Request a new whitelisted address
func (c *TradingClient) CreateWhitelistedAddress(body *CreateWhitelistedAddressRequest) (*WhitelistedAddress, error) {
	path := "/v2/wallets/whitelists"
	data, err := c.Raw.Post(path, nil, body)
	if err != nil {
		return nil, err
	}
	var result WhitelistedAddress
	return &result, json.Unmarshal(data, &result)
}

// DeleteWhitelistedAddress — Delete a whitelisted address
func (c *TradingClient) DeleteWhitelistedAddress(WhitelistedAddressID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/wallets/whitelists/%s", url.PathEscape(WhitelistedAddressID))
	data, err := c.Raw.Delete(path, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetWatchlists — Get All Watchlists
func (c *TradingClient) GetWatchlists() ([]WatchlistWithoutAsset, error) {
	path := "/v2/watchlists"
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var result []WatchlistWithoutAsset
	return result, json.Unmarshal(data, &result)
}

// PostWatchlist — Create Watchlist
func (c *TradingClient) PostWatchlist(body *UpdateWatchlistRequest) (*Watchlist, error) {
	path := "/v2/watchlists"
	data, err := c.Raw.Post(path, nil, body)
	if err != nil {
		return nil, err
	}
	var result Watchlist
	return &result, json.Unmarshal(data, &result)
}

// GetWatchlistByID — Get Watchlist by ID
func (c *TradingClient) GetWatchlistByID(WatchlistID string) (*Watchlist, error) {
	path := fmt.Sprintf("/v2/watchlists/%s", url.PathEscape(WatchlistID))
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var result Watchlist
	return &result, json.Unmarshal(data, &result)
}

// AddAssetToWatchlist — Add Asset to Watchlist
func (c *TradingClient) AddAssetToWatchlist(WatchlistID string, body *AddAssetToWatchlistRequest) (*Watchlist, error) {
	path := fmt.Sprintf("/v2/watchlists/%s", url.PathEscape(WatchlistID))
	data, err := c.Raw.Post(path, nil, body)
	if err != nil {
		return nil, err
	}
	var result Watchlist
	return &result, json.Unmarshal(data, &result)
}

// UpdateWatchlistByID — Update Watchlist By Id
func (c *TradingClient) UpdateWatchlistByID(WatchlistID string, body *UpdateWatchlistRequest) (*Watchlist, error) {
	path := fmt.Sprintf("/v2/watchlists/%s", url.PathEscape(WatchlistID))
	data, err := c.Raw.Put(path, nil, body)
	if err != nil {
		return nil, err
	}
	var result Watchlist
	return &result, json.Unmarshal(data, &result)
}

// DeleteWatchlistByID — Delete Watchlist By Id
func (c *TradingClient) DeleteWatchlistByID(WatchlistID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/watchlists/%s", url.PathEscape(WatchlistID))
	data, err := c.Raw.Delete(path, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// RemoveAssetFromWatchlist — Delete Symbol from Watchlist
func (c *TradingClient) RemoveAssetFromWatchlist(WatchlistID string, Symbol string) (*Watchlist, error) {
	path := fmt.Sprintf("/v2/watchlists/%s/%s", url.PathEscape(WatchlistID), url.PathEscape(Symbol))
	data, err := c.Raw.Delete(path, nil)
	if err != nil {
		return nil, err
	}
	var result Watchlist
	return &result, json.Unmarshal(data, &result)
}

// GetWatchlistByName — Get Watchlist by Name
func (c *TradingClient) GetWatchlistByName(params url.Values) (*Watchlist, error) {
	path := "/v2/watchlists:by_name"
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	var result Watchlist
	return &result, json.Unmarshal(data, &result)
}

type AddAssetToWatchlistByNameRequest struct {
	Symbol string `json:"symbol,omitempty"`
}

// AddAssetToWatchlistByName — Add Asset to Watchlist By Name
func (c *TradingClient) AddAssetToWatchlistByName(params url.Values, body *AddAssetToWatchlistByNameRequest) (*Watchlist, error) {
	path := "/v2/watchlists:by_name"
	data, err := c.Raw.Post(path, params, body)
	if err != nil {
		return nil, err
	}
	var result Watchlist
	return &result, json.Unmarshal(data, &result)
}

// UpdateWatchlistByName — Update Watchlist By Name
func (c *TradingClient) UpdateWatchlistByName(params url.Values, body *UpdateWatchlistRequest) (*Watchlist, error) {
	path := "/v2/watchlists:by_name"
	data, err := c.Raw.Put(path, params, body)
	if err != nil {
		return nil, err
	}
	var result Watchlist
	return &result, json.Unmarshal(data, &result)
}

// DeleteWatchlistByName — Delete Watchlist By Name
func (c *TradingClient) DeleteWatchlistByName(params url.Values) (json.RawMessage, error) {
	path := "/v2/watchlists:by_name"
	data, err := c.Raw.Delete(path, params)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Calendar — Get Market Calendar
func (c *TradingClient) Calendar(Market string, params url.Values) (*PublicCalendarResp, error) {
	path := fmt.Sprintf("/v3/calendar/%s", url.PathEscape(Market))
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	var result PublicCalendarResp
	return &result, json.Unmarshal(data, &result)
}

// Clock — Get Market Clock
func (c *TradingClient) Clock(params url.Values) (*ClockResp, error) {
	path := "/v3/clock"
	data, err := c.Raw.Get(path, params)
	if err != nil {
		return nil, err
	}
	var result ClockResp
	return &result, json.Unmarshal(data, &result)
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
