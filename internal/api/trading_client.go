// Code generated from api/specs/trading-api.json; DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
	"net/url"

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

type GetAccountActivitiesParams struct {
	ActivityTypes string
	Category      string
	Date          string
	Until         string
	After         string
	Direction     string // default: desc
	PageSize      int    // default: 100
	PageToken     string
}

func (p *GetAccountActivitiesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.ActivityTypes != "" {
		v.Set("activity_types", p.ActivityTypes)
	}
	if p.Category != "" {
		v.Set("category", p.Category)
	}
	if p.Date != "" {
		v.Set("date", p.Date)
	}
	if p.Until != "" {
		v.Set("until", p.Until)
	}
	if p.After != "" {
		v.Set("after", p.After)
	}
	if p.Direction != "" {
		v.Set("direction", p.Direction)
	}
	if p.PageSize != 0 {
		v.Set("page_size", fmt.Sprint(p.PageSize))
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	return v
}

var GetAccountActivitiesParamsCategoryValues = []string{"non_trade_activity", "trade_activity"}

var GetAccountActivitiesParamsDirectionValues = []string{"asc", "desc"}

var GetAccountActivitiesParamsDefaults = map[string]string{
	"direction": "desc",
	"page_size": "100",
}

// GetAccountActivities — Retrieve Account Activities
func (c *TradingClient) GetAccountActivities(params *GetAccountActivitiesParams) (json.RawMessage, error) {
	path := "/v2/account/activities"
	data, err := c.Raw.Get(path, params.Values())
	if err != nil {
		return nil, err
	}
	return data, nil
}

type GetAccountActivitiesByActivityTypeParams struct {
	Date      string
	Until     string
	After     string
	Direction string // default: desc
	PageSize  int    // default: 100
	PageToken string
}

func (p *GetAccountActivitiesByActivityTypeParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Date != "" {
		v.Set("date", p.Date)
	}
	if p.Until != "" {
		v.Set("until", p.Until)
	}
	if p.After != "" {
		v.Set("after", p.After)
	}
	if p.Direction != "" {
		v.Set("direction", p.Direction)
	}
	if p.PageSize != 0 {
		v.Set("page_size", fmt.Sprint(p.PageSize))
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	return v
}

var GetAccountActivitiesByActivityTypeParamsDirectionValues = []string{"asc", "desc"}

var GetAccountActivitiesByActivityTypeParamsDefaults = map[string]string{
	"direction": "desc",
	"page_size": "100",
}

// GetAccountActivitiesByActivityType — Retrieve Account Activities of Specific Type
func (c *TradingClient) GetAccountActivitiesByActivityType(ActivityType string, params *GetAccountActivitiesByActivityTypeParams) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/account/activities/%s", ActivityType)
	data, err := c.Raw.Get(path, params.Values())
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
	data, err := c.Raw.Patch(path, body)
	if err != nil {
		return nil, err
	}
	var result AccountConfigurations
	return &result, json.Unmarshal(data, &result)
}

type GetAccountPortfolioHistoryParams struct {
	Period            string
	Timeframe         string
	IntradayReporting string // default: market_hours
	Start             string
	PNLReset          string // default: per_day
	End               string
	ExtendedHours     string
	CashflowTypes     string
}

func (p *GetAccountPortfolioHistoryParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Period != "" {
		v.Set("period", p.Period)
	}
	if p.Timeframe != "" {
		v.Set("timeframe", p.Timeframe)
	}
	if p.IntradayReporting != "" {
		v.Set("intraday_reporting", p.IntradayReporting)
	}
	if p.Start != "" {
		v.Set("start", p.Start)
	}
	if p.PNLReset != "" {
		v.Set("pnl_reset", p.PNLReset)
	}
	if p.End != "" {
		v.Set("end", p.End)
	}
	if p.ExtendedHours != "" {
		v.Set("extended_hours", p.ExtendedHours)
	}
	if p.CashflowTypes != "" {
		v.Set("cashflow_types", p.CashflowTypes)
	}
	return v
}

var GetAccountPortfolioHistoryParamsIntradayReportingValues = []string{"continuous", "extended_hours", "market_hours"}

var GetAccountPortfolioHistoryParamsPNLResetValues = []string{"no_reset", "per_day"}

var GetAccountPortfolioHistoryParamsDefaults = map[string]string{
	"intraday_reporting": "market_hours",
	"pnl_reset":          "per_day",
}

// GetAccountPortfolioHistory — Get Account Portfolio History
func (c *TradingClient) GetAccountPortfolioHistory(params *GetAccountPortfolioHistoryParams) (*PortfolioHistory, error) {
	path := "/v2/account/portfolio/history"
	data, err := c.Raw.Get(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result PortfolioHistory
	return &result, json.Unmarshal(data, &result)
}

type GetV2AssetsParams struct {
	Status     string
	AssetClass string
	Exchange   string
	Attributes string // default: []
}

func (p *GetV2AssetsParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Status != "" {
		v.Set("status", p.Status)
	}
	if p.AssetClass != "" {
		v.Set("asset_class", p.AssetClass)
	}
	if p.Exchange != "" {
		v.Set("exchange", p.Exchange)
	}
	if p.Attributes != "" {
		v.Set("attributes", p.Attributes)
	}
	return v
}

var GetV2AssetsParamsDefaults = map[string]string{
	"attributes": "[]",
}

// GetV2Assets — Get Assets
func (c *TradingClient) GetV2Assets(params *GetV2AssetsParams) ([]Assets, error) {
	path := "/v2/assets"
	data, err := c.Raw.Get(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result []Assets
	return result, json.Unmarshal(data, &result)
}

type UsCorporatesParams struct {
	BondStatus string
	Isins      string
	Cusips     string
	Tickers    string
}

func (p *UsCorporatesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.BondStatus != "" {
		v.Set("bond_status", p.BondStatus)
	}
	if p.Isins != "" {
		v.Set("isins", p.Isins)
	}
	if p.Cusips != "" {
		v.Set("cusips", p.Cusips)
	}
	if p.Tickers != "" {
		v.Set("tickers", p.Tickers)
	}
	return v
}

// UsCorporates — Get US corporates
func (c *TradingClient) UsCorporates(params *UsCorporatesParams) (*UsCorporatesResp, error) {
	path := "/v2/assets/fixed_income/us_corporates"
	data, err := c.Raw.Get(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result UsCorporatesResp
	return &result, json.Unmarshal(data, &result)
}

type UsTreasuriesParams struct {
	Subtype    string
	BondStatus string
	Cusips     string
	Isins      string
}

func (p *UsTreasuriesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Subtype != "" {
		v.Set("subtype", p.Subtype)
	}
	if p.BondStatus != "" {
		v.Set("bond_status", p.BondStatus)
	}
	if p.Cusips != "" {
		v.Set("cusips", p.Cusips)
	}
	if p.Isins != "" {
		v.Set("isins", p.Isins)
	}
	return v
}

// UsTreasuries — Get US treasuries
func (c *TradingClient) UsTreasuries(params *UsTreasuriesParams) (*UsTreasuriesResp, error) {
	path := "/v2/assets/fixed_income/us_treasuries"
	data, err := c.Raw.Get(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result UsTreasuriesResp
	return &result, json.Unmarshal(data, &result)
}

// GetV2AssetsSymbolOrAssetID — Get an Asset by ID or Symbol
func (c *TradingClient) GetV2AssetsSymbolOrAssetID(SymbolOrAssetID string) (*Assets, error) {
	path := fmt.Sprintf("/v2/assets/%s", SymbolOrAssetID)
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var result Assets
	return &result, json.Unmarshal(data, &result)
}

type LegacyCalendarParams struct {
	Start    string
	End      string
	DateType string
}

func (p *LegacyCalendarParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Start != "" {
		v.Set("start", p.Start)
	}
	if p.End != "" {
		v.Set("end", p.End)
	}
	if p.DateType != "" {
		v.Set("date_type", p.DateType)
	}
	return v
}

var LegacyCalendarParamsDateTypeValues = []string{"SETTLEMENT", "TRADING"}

// LegacyCalendar — Get US Market Calendar
func (c *TradingClient) LegacyCalendar(params *LegacyCalendarParams) (json.RawMessage, error) {
	path := "/v2/calendar"
	data, err := c.Raw.Get(path, params.Values())
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

type GetV2CorporateActionsAnnouncementsParams struct {
	CaTypes  string
	Since    string
	Until    string
	Symbol   string
	Cusip    string
	DateType string
}

func (p *GetV2CorporateActionsAnnouncementsParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.CaTypes != "" {
		v.Set("ca_types", p.CaTypes)
	}
	if p.Since != "" {
		v.Set("since", p.Since)
	}
	if p.Until != "" {
		v.Set("until", p.Until)
	}
	if p.Symbol != "" {
		v.Set("symbol", p.Symbol)
	}
	if p.Cusip != "" {
		v.Set("cusip", p.Cusip)
	}
	if p.DateType != "" {
		v.Set("date_type", p.DateType)
	}
	return v
}

// GetV2CorporateActionsAnnouncements — Retrieve Announcements
func (c *TradingClient) GetV2CorporateActionsAnnouncements(params *GetV2CorporateActionsAnnouncementsParams) (json.RawMessage, error) {
	path := "/v2/corporate_actions/announcements"
	data, err := c.Raw.Get(path, params.Values())
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetV2CorporateActionsAnnouncementsID — Retrieve a Specific Announcement
func (c *TradingClient) GetV2CorporateActionsAnnouncementsID(ID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/corporate_actions/announcements/%s", ID)
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type GetOptionsContractsParams struct {
	UnderlyingSymbols string
	ShowDeliverables  bool
	Status            string
	ExpirationDate    string
	ExpirationDateGte string
	ExpirationDateLte string
	RootSymbol        string
	Type              string
	Style             string
	StrikePriceGte    float64
	StrikePriceLte    float64
	PageToken         string
	Limit             int
	Ppind             bool
}

func (p *GetOptionsContractsParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.UnderlyingSymbols != "" {
		v.Set("underlying_symbols", p.UnderlyingSymbols)
	}
	if p.ShowDeliverables {
		v.Set("show_deliverables", "true")
	}
	if p.Status != "" {
		v.Set("status", p.Status)
	}
	if p.ExpirationDate != "" {
		v.Set("expiration_date", p.ExpirationDate)
	}
	if p.ExpirationDateGte != "" {
		v.Set("expiration_date_gte", p.ExpirationDateGte)
	}
	if p.ExpirationDateLte != "" {
		v.Set("expiration_date_lte", p.ExpirationDateLte)
	}
	if p.RootSymbol != "" {
		v.Set("root_symbol", p.RootSymbol)
	}
	if p.Type != "" {
		v.Set("type", p.Type)
	}
	if p.Style != "" {
		v.Set("style", p.Style)
	}
	if p.StrikePriceGte != 0 {
		v.Set("strike_price_gte", fmt.Sprintf("%g", p.StrikePriceGte))
	}
	if p.StrikePriceLte != 0 {
		v.Set("strike_price_lte", fmt.Sprintf("%g", p.StrikePriceLte))
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.Ppind {
		v.Set("ppind", "true")
	}
	return v
}

var GetOptionsContractsParamsStatusValues = []string{"active", "inactive"}

var GetOptionsContractsParamsTypeValues = []string{"call", "put"}

var GetOptionsContractsParamsStyleValues = []string{"american", "european"}

// GetOptionsContracts — Get Option Contracts
func (c *TradingClient) GetOptionsContracts(params *GetOptionsContractsParams) (json.RawMessage, error) {
	path := "/v2/options/contracts"
	data, err := c.Raw.Get(path, params.Values())
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetOptionContractSymbolOrID — Get an option contract by ID or Symbol
func (c *TradingClient) GetOptionContractSymbolOrID(SymbolOrID string) (*OptionContract, error) {
	path := fmt.Sprintf("/v2/options/contracts/%s", SymbolOrID)
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var result OptionContract
	return &result, json.Unmarshal(data, &result)
}

type GetAllOrdersParams struct {
	Status        string
	Limit         int
	After         string
	Until         string
	Direction     string
	Nested        bool
	Symbols       string
	Side          string
	AssetClass    string
	BeforeOrderID string
	AfterOrderID  string
}

func (p *GetAllOrdersParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Status != "" {
		v.Set("status", p.Status)
	}
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.After != "" {
		v.Set("after", p.After)
	}
	if p.Until != "" {
		v.Set("until", p.Until)
	}
	if p.Direction != "" {
		v.Set("direction", p.Direction)
	}
	if p.Nested {
		v.Set("nested", "true")
	}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	if p.Side != "" {
		v.Set("side", p.Side)
	}
	if p.AssetClass != "" {
		v.Set("asset_class", p.AssetClass)
	}
	if p.BeforeOrderID != "" {
		v.Set("before_order_id", p.BeforeOrderID)
	}
	if p.AfterOrderID != "" {
		v.Set("after_order_id", p.AfterOrderID)
	}
	return v
}

var GetAllOrdersParamsStatusValues = []string{"all", "closed", "open"}

var GetAllOrdersParamsDirectionValues = []string{"asc", "desc"}

// GetAllOrders — Get All Orders
func (c *TradingClient) GetAllOrders(params *GetAllOrdersParams) ([]Order, error) {
	path := "/v2/orders"
	data, err := c.Raw.Get(path, params.Values())
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
	data, err := c.Raw.Post(path, body)
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

type GetOrderByOrderIDParams struct {
	Nested bool
}

func (p *GetOrderByOrderIDParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Nested {
		v.Set("nested", "true")
	}
	return v
}

// GetOrderByOrderID — Get Order by ID
func (c *TradingClient) GetOrderByOrderID(OrderID string, params *GetOrderByOrderIDParams) (*Order, error) {
	path := fmt.Sprintf("/v2/orders/%s", OrderID)
	data, err := c.Raw.Get(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result Order
	return &result, json.Unmarshal(data, &result)
}

// PatchOrderByOrderID — Replace Order by ID
func (c *TradingClient) PatchOrderByOrderID(OrderID string, body *PatchOrderRequest) (*Order, error) {
	path := fmt.Sprintf("/v2/orders/%s", OrderID)
	data, err := c.Raw.Patch(path, body)
	if err != nil {
		return nil, err
	}
	var result Order
	return &result, json.Unmarshal(data, &result)
}

// DeleteOrderByOrderID — Delete Order by ID
func (c *TradingClient) DeleteOrderByOrderID(OrderID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/orders/%s", OrderID)
	data, err := c.Raw.Delete(path, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type GetOrderByClientOrderIDParams struct {
	ClientOrderID string
}

func (p *GetOrderByClientOrderIDParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.ClientOrderID != "" {
		v.Set("client_order_id", p.ClientOrderID)
	}
	return v
}

// GetOrderByClientOrderID — Get Order by Client Order ID
func (c *TradingClient) GetOrderByClientOrderID(params *GetOrderByClientOrderIDParams) (*Order, error) {
	path := "/v2/orders:by_client_order_id"
	data, err := c.Raw.Get(path, params.Values())
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

type GetCryptoPerpAccountLeverageParams struct {
	Symbol string
}

func (p *GetCryptoPerpAccountLeverageParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbol != "" {
		v.Set("symbol", p.Symbol)
	}
	return v
}

// GetCryptoPerpAccountLeverage — Get Account Leverage for an Asset
func (c *TradingClient) GetCryptoPerpAccountLeverage(params *GetCryptoPerpAccountLeverageParams) (json.RawMessage, error) {
	path := "/v2/perpetuals/leverage"
	data, err := c.Raw.Get(path, params.Values())
	if err != nil {
		return nil, err
	}
	return data, nil
}

type SetCryptoPerpAccountLeverageParams struct {
	Symbol   string
	Leverage int
}

func (p *SetCryptoPerpAccountLeverageParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbol != "" {
		v.Set("symbol", p.Symbol)
	}
	if p.Leverage != 0 {
		v.Set("leverage", fmt.Sprint(p.Leverage))
	}
	return v
}

// SetCryptoPerpAccountLeverage — Set Account Leverage for an Asset
func (c *TradingClient) SetCryptoPerpAccountLeverage(params *SetCryptoPerpAccountLeverageParams) (json.RawMessage, error) {
	path := "/v2/perpetuals/leverage"
	data, err := c.Raw.Post(path, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type ListCryptoPerpFundingWalletsParams struct {
	Asset string
}

func (p *ListCryptoPerpFundingWalletsParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Asset != "" {
		v.Set("asset", p.Asset)
	}
	return v
}

// ListCryptoPerpFundingWallets — Retrieve Crypto Funding Wallets
func (c *TradingClient) ListCryptoPerpFundingWallets(params *ListCryptoPerpFundingWalletsParams) (*CryptoWallet, error) {
	path := "/v2/perpetuals/wallets"
	data, err := c.Raw.Get(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result CryptoWallet
	return &result, json.Unmarshal(data, &result)
}

type GetCryptoPerpTransferEstimateParams struct {
	Asset       string
	FromAddress string
	ToAddress   string
	Amount      string
}

func (p *GetCryptoPerpTransferEstimateParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Asset != "" {
		v.Set("asset", p.Asset)
	}
	if p.FromAddress != "" {
		v.Set("from_address", p.FromAddress)
	}
	if p.ToAddress != "" {
		v.Set("to_address", p.ToAddress)
	}
	if p.Amount != "" {
		v.Set("amount", p.Amount)
	}
	return v
}

// GetCryptoPerpTransferEstimate — Returns the estimated gas fee for a proposed transaction
func (c *TradingClient) GetCryptoPerpTransferEstimate(params *GetCryptoPerpTransferEstimateParams) (json.RawMessage, error) {
	path := "/v2/perpetuals/wallets/fees/estimate"
	data, err := c.Raw.Get(path, params.Values())
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
	data, err := c.Raw.Post(path, body)
	if err != nil {
		return nil, err
	}
	var result CryptoTransfer
	return &result, json.Unmarshal(data, &result)
}

// GetCryptoPerpFundingTransfer — Retrieve a Crypto Funding Transfer
func (c *TradingClient) GetCryptoPerpFundingTransfer(TransferID string) (*CryptoTransfer, error) {
	path := fmt.Sprintf("/v2/perpetuals/wallets/transfers/%s", TransferID)
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
	data, err := c.Raw.Post(path, body)
	if err != nil {
		return nil, err
	}
	var result WhitelistedAddress
	return &result, json.Unmarshal(data, &result)
}

// DeleteWhitelistedPerpAddress — Delete a whitelisted address
func (c *TradingClient) DeleteWhitelistedPerpAddress(WhitelistedAddressID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/perpetuals/wallets/whitelists/%s", WhitelistedAddressID)
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

type DeleteAllOpenPositionsParams struct {
	CancelOrders bool
}

func (p *DeleteAllOpenPositionsParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.CancelOrders {
		v.Set("cancel_orders", "true")
	}
	return v
}

// DeleteAllOpenPositions — Close All Positions
func (c *TradingClient) DeleteAllOpenPositions(params *DeleteAllOpenPositionsParams) ([]PositionClosedReponse, error) {
	path := "/v2/positions"
	data, err := c.Raw.Delete(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result []PositionClosedReponse
	return result, json.Unmarshal(data, &result)
}

// GetOpenPosition — Get an Open Position
func (c *TradingClient) GetOpenPosition(SymbolOrAssetID string) (*Position, error) {
	path := fmt.Sprintf("/v2/positions/%s", SymbolOrAssetID)
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var result Position
	return &result, json.Unmarshal(data, &result)
}

type DeleteOpenPositionParams struct {
	Qty        float64
	Percentage float64
}

func (p *DeleteOpenPositionParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Qty != 0 {
		v.Set("qty", fmt.Sprintf("%g", p.Qty))
	}
	if p.Percentage != 0 {
		v.Set("percentage", fmt.Sprintf("%g", p.Percentage))
	}
	return v
}

// DeleteOpenPosition — Close a Position
func (c *TradingClient) DeleteOpenPosition(SymbolOrAssetID string, params *DeleteOpenPositionParams) (*Order, error) {
	path := fmt.Sprintf("/v2/positions/%s", SymbolOrAssetID)
	data, err := c.Raw.Delete(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result Order
	return &result, json.Unmarshal(data, &result)
}

// OptionDoNotExercise — Do Not Exercise an Options Position
func (c *TradingClient) OptionDoNotExercise(SymbolOrContractID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/positions/%s/do-not-exercise", SymbolOrContractID)
	data, err := c.Raw.Post(path, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// OptionExercise — Exercise an Options Position
func (c *TradingClient) OptionExercise(SymbolOrContractID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/positions/%s/exercise", SymbolOrContractID)
	data, err := c.Raw.Post(path, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type ListCryptoFundingWalletsParams struct {
	Asset   string
	Network string
}

func (p *ListCryptoFundingWalletsParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Asset != "" {
		v.Set("asset", p.Asset)
	}
	if p.Network != "" {
		v.Set("network", p.Network)
	}
	return v
}

var ListCryptoFundingWalletsParamsNetworkValues = []string{"ethereum", "solana"}

// ListCryptoFundingWallets — Retrieve Crypto Funding Wallets
func (c *TradingClient) ListCryptoFundingWallets(params *ListCryptoFundingWalletsParams) (*CryptoWallet, error) {
	path := "/v2/wallets"
	data, err := c.Raw.Get(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result CryptoWallet
	return &result, json.Unmarshal(data, &result)
}

type GetCryptoTransferEstimateParams struct {
	Asset       string
	FromAddress string
	ToAddress   string
	Amount      string
}

func (p *GetCryptoTransferEstimateParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Asset != "" {
		v.Set("asset", p.Asset)
	}
	if p.FromAddress != "" {
		v.Set("from_address", p.FromAddress)
	}
	if p.ToAddress != "" {
		v.Set("to_address", p.ToAddress)
	}
	if p.Amount != "" {
		v.Set("amount", p.Amount)
	}
	return v
}

// GetCryptoTransferEstimate — Returns the estimated gas fee for a proposed transaction.
func (c *TradingClient) GetCryptoTransferEstimate(params *GetCryptoTransferEstimateParams) (json.RawMessage, error) {
	path := "/v2/wallets/fees/estimate"
	data, err := c.Raw.Get(path, params.Values())
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
	data, err := c.Raw.Post(path, body)
	if err != nil {
		return nil, err
	}
	var result CryptoTransfer
	return &result, json.Unmarshal(data, &result)
}

// GetCryptoFundingTransfer — Retrieve a Crypto Funding Transfer
func (c *TradingClient) GetCryptoFundingTransfer(TransferID string) (*CryptoTransfer, error) {
	path := fmt.Sprintf("/v2/wallets/transfers/%s", TransferID)
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
	data, err := c.Raw.Post(path, body)
	if err != nil {
		return nil, err
	}
	var result WhitelistedAddress
	return &result, json.Unmarshal(data, &result)
}

// DeleteWhitelistedAddress — Delete a whitelisted address
func (c *TradingClient) DeleteWhitelistedAddress(WhitelistedAddressID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/wallets/whitelists/%s", WhitelistedAddressID)
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
	data, err := c.Raw.Post(path, body)
	if err != nil {
		return nil, err
	}
	var result Watchlist
	return &result, json.Unmarshal(data, &result)
}

// GetWatchlistByID — Get Watchlist by ID
func (c *TradingClient) GetWatchlistByID(WatchlistID string) (*Watchlist, error) {
	path := fmt.Sprintf("/v2/watchlists/%s", WatchlistID)
	data, err := c.Raw.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var result Watchlist
	return &result, json.Unmarshal(data, &result)
}

// AddAssetToWatchlist — Add Asset to Watchlist
func (c *TradingClient) AddAssetToWatchlist(WatchlistID string, body *AddAssetToWatchlistRequest) (*Watchlist, error) {
	path := fmt.Sprintf("/v2/watchlists/%s", WatchlistID)
	data, err := c.Raw.Post(path, body)
	if err != nil {
		return nil, err
	}
	var result Watchlist
	return &result, json.Unmarshal(data, &result)
}

// UpdateWatchlistByID — Update Watchlist By Id
func (c *TradingClient) UpdateWatchlistByID(WatchlistID string, body *UpdateWatchlistRequest) (*Watchlist, error) {
	path := fmt.Sprintf("/v2/watchlists/%s", WatchlistID)
	data, err := c.Raw.Put(path, body)
	if err != nil {
		return nil, err
	}
	var result Watchlist
	return &result, json.Unmarshal(data, &result)
}

// DeleteWatchlistByID — Delete Watchlist By Id
func (c *TradingClient) DeleteWatchlistByID(WatchlistID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/watchlists/%s", WatchlistID)
	data, err := c.Raw.Delete(path, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// RemoveAssetFromWatchlist — Delete Symbol from Watchlist
func (c *TradingClient) RemoveAssetFromWatchlist(WatchlistID string, Symbol string) (*Watchlist, error) {
	path := fmt.Sprintf("/v2/watchlists/%s/%s", WatchlistID, Symbol)
	data, err := c.Raw.Delete(path, nil)
	if err != nil {
		return nil, err
	}
	var result Watchlist
	return &result, json.Unmarshal(data, &result)
}

type GetWatchlistByNameParams struct {
	Name string
}

func (p *GetWatchlistByNameParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Name != "" {
		v.Set("name", p.Name)
	}
	return v
}

// GetWatchlistByName — Get Watchlist by Name
func (c *TradingClient) GetWatchlistByName(params *GetWatchlistByNameParams) (*Watchlist, error) {
	path := "/v2/watchlists:by_name"
	data, err := c.Raw.Get(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result Watchlist
	return &result, json.Unmarshal(data, &result)
}

type AddAssetToWatchlistByNameParams struct {
	Name string
}

func (p *AddAssetToWatchlistByNameParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Name != "" {
		v.Set("name", p.Name)
	}
	return v
}

type AddAssetToWatchlistByNameRequest struct {
	Symbol string `json:"symbol,omitempty"`
}

// AddAssetToWatchlistByName — Add Asset to Watchlist By Name
func (c *TradingClient) AddAssetToWatchlistByName(params *AddAssetToWatchlistByNameParams, body *AddAssetToWatchlistByNameRequest) (*Watchlist, error) {
	path := "/v2/watchlists:by_name"
	data, err := c.Raw.Post(path, body)
	if err != nil {
		return nil, err
	}
	var result Watchlist
	return &result, json.Unmarshal(data, &result)
}

type UpdateWatchlistByNameParams struct {
	Name string
}

func (p *UpdateWatchlistByNameParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Name != "" {
		v.Set("name", p.Name)
	}
	return v
}

// UpdateWatchlistByName — Update Watchlist By Name
func (c *TradingClient) UpdateWatchlistByName(params *UpdateWatchlistByNameParams, body *UpdateWatchlistRequest) (*Watchlist, error) {
	path := "/v2/watchlists:by_name"
	data, err := c.Raw.Put(path, body)
	if err != nil {
		return nil, err
	}
	var result Watchlist
	return &result, json.Unmarshal(data, &result)
}

type DeleteWatchlistByNameParams struct {
	Name string
}

func (p *DeleteWatchlistByNameParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Name != "" {
		v.Set("name", p.Name)
	}
	return v
}

// DeleteWatchlistByName — Delete Watchlist By Name
func (c *TradingClient) DeleteWatchlistByName(params *DeleteWatchlistByNameParams) (json.RawMessage, error) {
	path := "/v2/watchlists:by_name"
	data, err := c.Raw.Delete(path, params.Values())
	if err != nil {
		return nil, err
	}
	return data, nil
}

type CalendarParams struct {
	Start    string
	End      string
	Timezone string
}

func (p *CalendarParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Start != "" {
		v.Set("start", p.Start)
	}
	if p.End != "" {
		v.Set("end", p.End)
	}
	if p.Timezone != "" {
		v.Set("timezone", p.Timezone)
	}
	return v
}

var CalendarParamsTimezoneValues = []string{"UTC"}

// Calendar — Get Market Calendar
func (c *TradingClient) Calendar(Market string, params *CalendarParams) (*PublicCalendarResp, error) {
	path := fmt.Sprintf("/v3/calendar/%s", Market)
	data, err := c.Raw.Get(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result PublicCalendarResp
	return &result, json.Unmarshal(data, &result)
}

type ClockParams struct {
	Markets string
	Time    string
}

func (p *ClockParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Markets != "" {
		v.Set("markets", p.Markets)
	}
	if p.Time != "" {
		v.Set("time", p.Time)
	}
	return v
}

// Clock — Get Market Clock
func (c *TradingClient) Clock(params *ClockParams) (*ClockResp, error) {
	path := "/v3/clock"
	data, err := c.Raw.Get(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result ClockResp
	return &result, json.Unmarshal(data, &result)
}
