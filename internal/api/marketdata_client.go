// Code generated from api/specs/market-data-api.json; DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/alpacahq/cli/internal/client"
)

// MarketDataClient provides typed methods for the MarketData API.
type MarketDataClient struct {
	Raw *client.Client
}

func NewMarketDataClient(raw *client.Client) *MarketDataClient {
	return &MarketDataClient{Raw: raw}
}

type CorporateActionsParams struct {
	Symbols   string
	Cusips    string
	Types     string
	Start     string
	End       string
	Ids       string
	Limit     int
	PageToken string
	Sort      string
}

func (p *CorporateActionsParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	if p.Cusips != "" {
		v.Set("cusips", p.Cusips)
	}
	if p.Types != "" {
		v.Set("types", p.Types)
	}
	if p.Start != "" {
		v.Set("start", p.Start)
	}
	if p.End != "" {
		v.Set("end", p.End)
	}
	if p.Ids != "" {
		v.Set("ids", p.Ids)
	}
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	if p.Sort != "" {
		v.Set("sort", p.Sort)
	}
	return v
}

// CorporateActions — Corporate actions
func (c *MarketDataClient) CorporateActions(params *CorporateActionsParams) (*CorporateActionsResp, error) {
	path := "/v1/corporate-actions"
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result CorporateActionsResp
	return &result, json.Unmarshal(data, &result)
}

type CryptoPerpLatestBarsParams struct {
	Symbols string
}

func (p *CryptoPerpLatestBarsParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	return v
}

// CryptoPerpLatestBars — Latest bars
func (c *MarketDataClient) CryptoPerpLatestBars(Loc string, params *CryptoPerpLatestBarsParams) (*CryptoLatestBarsResp, error) {
	path := fmt.Sprintf("/v1beta1/crypto-perps/%s/latest/bars", Loc)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result CryptoLatestBarsResp
	return &result, json.Unmarshal(data, &result)
}

type CryptoPerpLatestOrderbooksParams struct {
	Symbols string
}

func (p *CryptoPerpLatestOrderbooksParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	return v
}

// CryptoPerpLatestOrderbooks — Latest orderbook
func (c *MarketDataClient) CryptoPerpLatestOrderbooks(Loc string, params *CryptoPerpLatestOrderbooksParams) (*CryptoLatestOrderbooksResp, error) {
	path := fmt.Sprintf("/v1beta1/crypto-perps/%s/latest/orderbooks", Loc)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result CryptoLatestOrderbooksResp
	return &result, json.Unmarshal(data, &result)
}

type CryptoPerpLatestFuturesPricingParams struct {
	Symbols string
}

func (p *CryptoPerpLatestFuturesPricingParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	return v
}

// CryptoPerpLatestFuturesPricing — Latest pricing
func (c *MarketDataClient) CryptoPerpLatestFuturesPricing(Loc string, params *CryptoPerpLatestFuturesPricingParams) (*CryptoPerpLatestFuturesPricingResp, error) {
	path := fmt.Sprintf("/v1beta1/crypto-perps/%s/latest/pricing", Loc)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result CryptoPerpLatestFuturesPricingResp
	return &result, json.Unmarshal(data, &result)
}

type CryptoPerpLatestQuotesParams struct {
	Symbols string
}

func (p *CryptoPerpLatestQuotesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	return v
}

// CryptoPerpLatestQuotes — Latest quotes
func (c *MarketDataClient) CryptoPerpLatestQuotes(Loc string, params *CryptoPerpLatestQuotesParams) (*CryptoLatestQuotesResp, error) {
	path := fmt.Sprintf("/v1beta1/crypto-perps/%s/latest/quotes", Loc)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result CryptoLatestQuotesResp
	return &result, json.Unmarshal(data, &result)
}

type CryptoPerpLatestTradesParams struct {
	Symbols string
}

func (p *CryptoPerpLatestTradesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	return v
}

// CryptoPerpLatestTrades — Latest trades
func (c *MarketDataClient) CryptoPerpLatestTrades(Loc string, params *CryptoPerpLatestTradesParams) (*CryptoLatestTradesResp, error) {
	path := fmt.Sprintf("/v1beta1/crypto-perps/%s/latest/trades", Loc)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result CryptoLatestTradesResp
	return &result, json.Unmarshal(data, &result)
}

type FixedIncomeLatestPricesParams struct {
	Isins string
}

func (p *FixedIncomeLatestPricesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Isins != "" {
		v.Set("isins", p.Isins)
	}
	return v
}

// FixedIncomeLatestPrices — Latest prices
func (c *MarketDataClient) FixedIncomeLatestPrices(params *FixedIncomeLatestPricesParams) (*FixedIncomeLatestPricesResp, error) {
	path := "/v1beta1/fixed_income/latest/prices"
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result FixedIncomeLatestPricesResp
	return &result, json.Unmarshal(data, &result)
}

type LatestRatesParams struct {
	CurrencyPairs string
}

func (p *LatestRatesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.CurrencyPairs != "" {
		v.Set("currency_pairs", p.CurrencyPairs)
	}
	return v
}

// LatestRates — Latest rates for currency pairs
func (c *MarketDataClient) LatestRates(params *LatestRatesParams) (*ForexLatestRatesResp, error) {
	path := "/v1beta1/forex/latest/rates"
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result ForexLatestRatesResp
	return &result, json.Unmarshal(data, &result)
}

type RatesParams struct {
	CurrencyPairs string
	Timeframe     string
	Start         string
	End           string
	Limit         int
	Sort          string
	PageToken     string
}

func (p *RatesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.CurrencyPairs != "" {
		v.Set("currency_pairs", p.CurrencyPairs)
	}
	if p.Timeframe != "" {
		v.Set("timeframe", p.Timeframe)
	}
	if p.Start != "" {
		v.Set("start", p.Start)
	}
	if p.End != "" {
		v.Set("end", p.End)
	}
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.Sort != "" {
		v.Set("sort", p.Sort)
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	return v
}

// Rates — Historical rates for currency pairs
func (c *MarketDataClient) Rates(params *RatesParams) (*ForexRatesResp, error) {
	path := "/v1beta1/forex/rates"
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result ForexRatesResp
	return &result, json.Unmarshal(data, &result)
}

type LogosParams struct {
	Placeholder bool
}

func (p *LogosParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Placeholder {
		v.Set("placeholder", "true")
	}
	return v
}

// Logos — Logos
func (c *MarketDataClient) Logos(Symbol string, params *LogosParams) (json.RawMessage, error) {
	path := fmt.Sprintf("/v1beta1/logos/%s", Symbol)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	return data, nil
}

type NewsParams struct {
	Start              string
	End                string
	Sort               string
	Symbols            string
	Limit              int
	IncludeContent     bool
	ExcludeContentless bool
	PageToken          string
}

func (p *NewsParams) Values() url.Values {
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
	if p.Sort != "" {
		v.Set("sort", p.Sort)
	}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.IncludeContent {
		v.Set("include_content", "true")
	}
	if p.ExcludeContentless {
		v.Set("exclude_contentless", "true")
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	return v
}

// News — News articles
func (c *MarketDataClient) News(params *NewsParams) (*NewsResp, error) {
	path := "/v1beta1/news"
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result NewsResp
	return &result, json.Unmarshal(data, &result)
}

type OptionBarsParams struct {
	Symbols   string
	Timeframe string
	Start     string
	End       string
	Limit     int
	PageToken string
	Sort      string
}

func (p *OptionBarsParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	if p.Timeframe != "" {
		v.Set("timeframe", p.Timeframe)
	}
	if p.Start != "" {
		v.Set("start", p.Start)
	}
	if p.End != "" {
		v.Set("end", p.End)
	}
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	if p.Sort != "" {
		v.Set("sort", p.Sort)
	}
	return v
}

// OptionBars — Historical bars
func (c *MarketDataClient) OptionBars(params *OptionBarsParams) (*OptionBarsResp, error) {
	path := "/v1beta1/options/bars"
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result OptionBarsResp
	return &result, json.Unmarshal(data, &result)
}

// OptionMetaConditions — Condition codes
func (c *MarketDataClient) OptionMetaConditions(Ticktype string) (json.RawMessage, error) {
	path := fmt.Sprintf("/v1beta1/options/meta/conditions/%s", Ticktype)
	data, err := c.Raw.GetData(path, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// OptionMetaExchanges — Exchange codes
func (c *MarketDataClient) OptionMetaExchanges() (json.RawMessage, error) {
	path := "/v1beta1/options/meta/exchanges"
	data, err := c.Raw.GetData(path, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type OptionLatestQuotesParams struct {
	Symbols string
	Feed    string
}

func (p *OptionLatestQuotesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	return v
}

// OptionLatestQuotes — Latest quotes
func (c *MarketDataClient) OptionLatestQuotes(params *OptionLatestQuotesParams) (*OptionLatestQuotesResp, error) {
	path := "/v1beta1/options/quotes/latest"
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result OptionLatestQuotesResp
	return &result, json.Unmarshal(data, &result)
}

type OptionSnapshotsParams struct {
	Symbols      string
	Feed         string
	UpdatedSince string
	Limit        int
	PageToken    string
}

func (p *OptionSnapshotsParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	if p.UpdatedSince != "" {
		v.Set("updated_since", p.UpdatedSince)
	}
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	return v
}

// OptionSnapshots — Snapshots
func (c *MarketDataClient) OptionSnapshots(params *OptionSnapshotsParams) (*OptionSnapshotsResp, error) {
	path := "/v1beta1/options/snapshots"
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result OptionSnapshotsResp
	return &result, json.Unmarshal(data, &result)
}

type OptionChainParams struct {
	Feed              string
	Limit             int
	UpdatedSince      string
	PageToken         string
	Type              string
	StrikePriceGte    float64
	StrikePriceLte    float64
	ExpirationDate    string
	ExpirationDateGte string
	ExpirationDateLte string
	RootSymbol        string
}

func (p *OptionChainParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.UpdatedSince != "" {
		v.Set("updated_since", p.UpdatedSince)
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	if p.Type != "" {
		v.Set("type", p.Type)
	}
	if p.StrikePriceGte != 0 {
		v.Set("strike_price_gte", fmt.Sprintf("%g", p.StrikePriceGte))
	}
	if p.StrikePriceLte != 0 {
		v.Set("strike_price_lte", fmt.Sprintf("%g", p.StrikePriceLte))
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
	return v
}

// OptionChain — Option chain
func (c *MarketDataClient) OptionChain(UnderlyingSymbol string, params *OptionChainParams) (*OptionSnapshotsResp, error) {
	path := fmt.Sprintf("/v1beta1/options/snapshots/%s", UnderlyingSymbol)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result OptionSnapshotsResp
	return &result, json.Unmarshal(data, &result)
}

type OptionTradesParams struct {
	Symbols   string
	Start     string
	End       string
	Limit     int
	PageToken string
	Sort      string
}

func (p *OptionTradesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	if p.Start != "" {
		v.Set("start", p.Start)
	}
	if p.End != "" {
		v.Set("end", p.End)
	}
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	if p.Sort != "" {
		v.Set("sort", p.Sort)
	}
	return v
}

// OptionTrades — Historical trades
func (c *MarketDataClient) OptionTrades(params *OptionTradesParams) (*OptionTradesResp, error) {
	path := "/v1beta1/options/trades"
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result OptionTradesResp
	return &result, json.Unmarshal(data, &result)
}

type OptionLatestTradesParams struct {
	Symbols string
	Feed    string
}

func (p *OptionLatestTradesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	return v
}

// OptionLatestTrades — Latest trades
func (c *MarketDataClient) OptionLatestTrades(params *OptionLatestTradesParams) (*OptionLatestTradesResp, error) {
	path := "/v1beta1/options/trades/latest"
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result OptionLatestTradesResp
	return &result, json.Unmarshal(data, &result)
}

type MostActivesParams struct {
	By  string
	Top int
}

func (p *MostActivesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.By != "" {
		v.Set("by", p.By)
	}
	if p.Top != 0 {
		v.Set("top", fmt.Sprint(p.Top))
	}
	return v
}

// MostActives — Most active stocks
func (c *MarketDataClient) MostActives(params *MostActivesParams) (*MostActivesResp, error) {
	path := "/v1beta1/screener/stocks/most-actives"
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result MostActivesResp
	return &result, json.Unmarshal(data, &result)
}

type MoversParams struct {
	Top int
}

func (p *MoversParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Top != 0 {
		v.Set("top", fmt.Sprint(p.Top))
	}
	return v
}

// Movers — Top market movers
func (c *MarketDataClient) Movers(MarketType string, params *MoversParams) (*MoversResp, error) {
	path := fmt.Sprintf("/v1beta1/screener/%s/movers", MarketType)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result MoversResp
	return &result, json.Unmarshal(data, &result)
}

type CryptoBarsParams struct {
	Symbols   string
	Timeframe string
	Start     string
	End       string
	Limit     int
	PageToken string
	Sort      string
}

func (p *CryptoBarsParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	if p.Timeframe != "" {
		v.Set("timeframe", p.Timeframe)
	}
	if p.Start != "" {
		v.Set("start", p.Start)
	}
	if p.End != "" {
		v.Set("end", p.End)
	}
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	if p.Sort != "" {
		v.Set("sort", p.Sort)
	}
	return v
}

// CryptoBars — Historical bars
func (c *MarketDataClient) CryptoBars(Loc string, params *CryptoBarsParams) (*CryptoBarsResp, error) {
	path := fmt.Sprintf("/v1beta3/crypto/%s/bars", Loc)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result CryptoBarsResp
	return &result, json.Unmarshal(data, &result)
}

type CryptoLatestBarsParams struct {
	Symbols string
}

func (p *CryptoLatestBarsParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	return v
}

// CryptoLatestBars — Latest bars
func (c *MarketDataClient) CryptoLatestBars(Loc string, params *CryptoLatestBarsParams) (*CryptoLatestBarsResp, error) {
	path := fmt.Sprintf("/v1beta3/crypto/%s/latest/bars", Loc)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result CryptoLatestBarsResp
	return &result, json.Unmarshal(data, &result)
}

type CryptoLatestOrderbooksParams struct {
	Symbols string
}

func (p *CryptoLatestOrderbooksParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	return v
}

// CryptoLatestOrderbooks — Latest orderbook
func (c *MarketDataClient) CryptoLatestOrderbooks(Loc string, params *CryptoLatestOrderbooksParams) (*CryptoLatestOrderbooksResp, error) {
	path := fmt.Sprintf("/v1beta3/crypto/%s/latest/orderbooks", Loc)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result CryptoLatestOrderbooksResp
	return &result, json.Unmarshal(data, &result)
}

type CryptoLatestQuotesParams struct {
	Symbols string
}

func (p *CryptoLatestQuotesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	return v
}

// CryptoLatestQuotes — Latest quotes
func (c *MarketDataClient) CryptoLatestQuotes(Loc string, params *CryptoLatestQuotesParams) (*CryptoLatestQuotesResp, error) {
	path := fmt.Sprintf("/v1beta3/crypto/%s/latest/quotes", Loc)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result CryptoLatestQuotesResp
	return &result, json.Unmarshal(data, &result)
}

type CryptoLatestTradesParams struct {
	Symbols string
}

func (p *CryptoLatestTradesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	return v
}

// CryptoLatestTrades — Latest trades
func (c *MarketDataClient) CryptoLatestTrades(Loc string, params *CryptoLatestTradesParams) (*CryptoLatestTradesResp, error) {
	path := fmt.Sprintf("/v1beta3/crypto/%s/latest/trades", Loc)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result CryptoLatestTradesResp
	return &result, json.Unmarshal(data, &result)
}

type CryptoQuotesParams struct {
	Symbols   string
	Start     string
	End       string
	Limit     int
	PageToken string
	Sort      string
}

func (p *CryptoQuotesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	if p.Start != "" {
		v.Set("start", p.Start)
	}
	if p.End != "" {
		v.Set("end", p.End)
	}
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	if p.Sort != "" {
		v.Set("sort", p.Sort)
	}
	return v
}

// CryptoQuotes — Historical quotes
func (c *MarketDataClient) CryptoQuotes(Loc string, params *CryptoQuotesParams) (*CryptoQuotesResp, error) {
	path := fmt.Sprintf("/v1beta3/crypto/%s/quotes", Loc)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result CryptoQuotesResp
	return &result, json.Unmarshal(data, &result)
}

type CryptoSnapshotsParams struct {
	Symbols string
}

func (p *CryptoSnapshotsParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	return v
}

// CryptoSnapshots — Snapshots
func (c *MarketDataClient) CryptoSnapshots(Loc string, params *CryptoSnapshotsParams) (*CryptoSnapshotsResp, error) {
	path := fmt.Sprintf("/v1beta3/crypto/%s/snapshots", Loc)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result CryptoSnapshotsResp
	return &result, json.Unmarshal(data, &result)
}

type CryptoTradesParams struct {
	Symbols   string
	Start     string
	End       string
	Limit     int
	PageToken string
	Sort      string
}

func (p *CryptoTradesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	if p.Start != "" {
		v.Set("start", p.Start)
	}
	if p.End != "" {
		v.Set("end", p.End)
	}
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	if p.Sort != "" {
		v.Set("sort", p.Sort)
	}
	return v
}

// CryptoTrades — Historical trades
func (c *MarketDataClient) CryptoTrades(Loc string, params *CryptoTradesParams) (*CryptoTradesResp, error) {
	path := fmt.Sprintf("/v1beta3/crypto/%s/trades", Loc)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result CryptoTradesResp
	return &result, json.Unmarshal(data, &result)
}

type StockAuctionsParams struct {
	Symbols   string
	Start     string
	End       string
	Limit     int
	Asof      string
	Feed      string
	Currency  string
	PageToken string
	Sort      string
}

func (p *StockAuctionsParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	if p.Start != "" {
		v.Set("start", p.Start)
	}
	if p.End != "" {
		v.Set("end", p.End)
	}
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.Asof != "" {
		v.Set("asof", p.Asof)
	}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	if p.Currency != "" {
		v.Set("currency", p.Currency)
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	if p.Sort != "" {
		v.Set("sort", p.Sort)
	}
	return v
}

// StockAuctions — Historical auctions
func (c *MarketDataClient) StockAuctions(params *StockAuctionsParams) (*StockAuctionsResp, error) {
	path := "/v2/stocks/auctions"
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result StockAuctionsResp
	return &result, json.Unmarshal(data, &result)
}

type StockBarsParams struct {
	Symbols    string
	Timeframe  string
	Start      string
	End        string
	Limit      int
	Adjustment string
	Asof       string
	Feed       string
	Currency   string
	PageToken  string
	Sort       string
}

func (p *StockBarsParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	if p.Timeframe != "" {
		v.Set("timeframe", p.Timeframe)
	}
	if p.Start != "" {
		v.Set("start", p.Start)
	}
	if p.End != "" {
		v.Set("end", p.End)
	}
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.Adjustment != "" {
		v.Set("adjustment", p.Adjustment)
	}
	if p.Asof != "" {
		v.Set("asof", p.Asof)
	}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	if p.Currency != "" {
		v.Set("currency", p.Currency)
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	if p.Sort != "" {
		v.Set("sort", p.Sort)
	}
	return v
}

// StockBars — Historical bars
func (c *MarketDataClient) StockBars(params *StockBarsParams) (*StockBarsResp, error) {
	path := "/v2/stocks/bars"
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result StockBarsResp
	return &result, json.Unmarshal(data, &result)
}

type StockLatestBarsParams struct {
	Symbols  string
	Feed     string
	Currency string
}

func (p *StockLatestBarsParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	if p.Currency != "" {
		v.Set("currency", p.Currency)
	}
	return v
}

// StockLatestBars — Latest bars
func (c *MarketDataClient) StockLatestBars(params *StockLatestBarsParams) (*StockLatestBarsResp, error) {
	path := "/v2/stocks/bars/latest"
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result StockLatestBarsResp
	return &result, json.Unmarshal(data, &result)
}

type StockMetaConditionsParams struct {
	Tape string
}

func (p *StockMetaConditionsParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Tape != "" {
		v.Set("tape", p.Tape)
	}
	return v
}

// StockMetaConditions — Condition codes
func (c *MarketDataClient) StockMetaConditions(Ticktype string, params *StockMetaConditionsParams) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/stocks/meta/conditions/%s", Ticktype)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	return data, nil
}

// StockMetaExchanges — Exchange codes
func (c *MarketDataClient) StockMetaExchanges() (json.RawMessage, error) {
	path := "/v2/stocks/meta/exchanges"
	data, err := c.Raw.GetData(path, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type StockQuotesParams struct {
	Symbols   string
	Start     string
	End       string
	Limit     int
	Asof      string
	Feed      string
	Currency  string
	PageToken string
	Sort      string
}

func (p *StockQuotesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	if p.Start != "" {
		v.Set("start", p.Start)
	}
	if p.End != "" {
		v.Set("end", p.End)
	}
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.Asof != "" {
		v.Set("asof", p.Asof)
	}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	if p.Currency != "" {
		v.Set("currency", p.Currency)
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	if p.Sort != "" {
		v.Set("sort", p.Sort)
	}
	return v
}

// StockQuotes — Historical quotes
func (c *MarketDataClient) StockQuotes(params *StockQuotesParams) (*StockQuotesResp, error) {
	path := "/v2/stocks/quotes"
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result StockQuotesResp
	return &result, json.Unmarshal(data, &result)
}

type StockLatestQuotesParams struct {
	Symbols  string
	Feed     string
	Currency string
}

func (p *StockLatestQuotesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	if p.Currency != "" {
		v.Set("currency", p.Currency)
	}
	return v
}

// StockLatestQuotes — Latest quotes
func (c *MarketDataClient) StockLatestQuotes(params *StockLatestQuotesParams) (*StockLatestQuotesResp, error) {
	path := "/v2/stocks/quotes/latest"
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result StockLatestQuotesResp
	return &result, json.Unmarshal(data, &result)
}

type StockSnapshotsParams struct {
	Symbols  string
	Feed     string
	Currency string
}

func (p *StockSnapshotsParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	if p.Currency != "" {
		v.Set("currency", p.Currency)
	}
	return v
}

// StockSnapshots — Snapshots
func (c *MarketDataClient) StockSnapshots(params *StockSnapshotsParams) (json.RawMessage, error) {
	path := "/v2/stocks/snapshots"
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	return data, nil
}

type StockTradesParams struct {
	Symbols   string
	Start     string
	End       string
	Limit     int
	Asof      string
	Feed      string
	Currency  string
	PageToken string
	Sort      string
}

func (p *StockTradesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	if p.Start != "" {
		v.Set("start", p.Start)
	}
	if p.End != "" {
		v.Set("end", p.End)
	}
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.Asof != "" {
		v.Set("asof", p.Asof)
	}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	if p.Currency != "" {
		v.Set("currency", p.Currency)
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	if p.Sort != "" {
		v.Set("sort", p.Sort)
	}
	return v
}

// StockTrades — Historical trades
func (c *MarketDataClient) StockTrades(params *StockTradesParams) (*StockTradesResp, error) {
	path := "/v2/stocks/trades"
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result StockTradesResp
	return &result, json.Unmarshal(data, &result)
}

type StockLatestTradesParams struct {
	Symbols  string
	Feed     string
	Currency string
}

func (p *StockLatestTradesParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Symbols != "" {
		v.Set("symbols", p.Symbols)
	}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	if p.Currency != "" {
		v.Set("currency", p.Currency)
	}
	return v
}

// StockLatestTrades — Latest trades
func (c *MarketDataClient) StockLatestTrades(params *StockLatestTradesParams) (*StockLatestTradesResp, error) {
	path := "/v2/stocks/trades/latest"
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result StockLatestTradesResp
	return &result, json.Unmarshal(data, &result)
}

type StockAuctionSingleParams struct {
	Start     string
	End       string
	Limit     int
	Asof      string
	Feed      string
	Currency  string
	PageToken string
	Sort      string
}

func (p *StockAuctionSingleParams) Values() url.Values {
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
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.Asof != "" {
		v.Set("asof", p.Asof)
	}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	if p.Currency != "" {
		v.Set("currency", p.Currency)
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	if p.Sort != "" {
		v.Set("sort", p.Sort)
	}
	return v
}

// StockAuctionSingle — Historical auctions (single)
func (c *MarketDataClient) StockAuctionSingle(Symbol string, params *StockAuctionSingleParams) (*StockAuctionsRespSingle, error) {
	path := fmt.Sprintf("/v2/stocks/%s/auctions", Symbol)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result StockAuctionsRespSingle
	return &result, json.Unmarshal(data, &result)
}

type StockBarSingleParams struct {
	Timeframe  string
	Start      string
	End        string
	Limit      int
	Adjustment string
	Asof       string
	Feed       string
	Currency   string
	PageToken  string
	Sort       string
}

func (p *StockBarSingleParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Timeframe != "" {
		v.Set("timeframe", p.Timeframe)
	}
	if p.Start != "" {
		v.Set("start", p.Start)
	}
	if p.End != "" {
		v.Set("end", p.End)
	}
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.Adjustment != "" {
		v.Set("adjustment", p.Adjustment)
	}
	if p.Asof != "" {
		v.Set("asof", p.Asof)
	}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	if p.Currency != "" {
		v.Set("currency", p.Currency)
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	if p.Sort != "" {
		v.Set("sort", p.Sort)
	}
	return v
}

// StockBarSingle — Historical bars (single symbol)
func (c *MarketDataClient) StockBarSingle(Symbol string, params *StockBarSingleParams) (*StockBarsRespSingle, error) {
	path := fmt.Sprintf("/v2/stocks/%s/bars", Symbol)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result StockBarsRespSingle
	return &result, json.Unmarshal(data, &result)
}

type StockLatestBarSingleParams struct {
	Feed     string
	Currency string
}

func (p *StockLatestBarSingleParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	if p.Currency != "" {
		v.Set("currency", p.Currency)
	}
	return v
}

// StockLatestBarSingle — Latest bar (single symbol)
func (c *MarketDataClient) StockLatestBarSingle(Symbol string, params *StockLatestBarSingleParams) (*StockLatestBarsRespSingle, error) {
	path := fmt.Sprintf("/v2/stocks/%s/bars/latest", Symbol)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result StockLatestBarsRespSingle
	return &result, json.Unmarshal(data, &result)
}

type StockQuoteSingleParams struct {
	Start     string
	End       string
	Limit     int
	Asof      string
	Feed      string
	Currency  string
	PageToken string
	Sort      string
}

func (p *StockQuoteSingleParams) Values() url.Values {
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
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.Asof != "" {
		v.Set("asof", p.Asof)
	}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	if p.Currency != "" {
		v.Set("currency", p.Currency)
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	if p.Sort != "" {
		v.Set("sort", p.Sort)
	}
	return v
}

// StockQuoteSingle — Historical quotes (single symbol)
func (c *MarketDataClient) StockQuoteSingle(Symbol string, params *StockQuoteSingleParams) (*StockQuotesRespSingle, error) {
	path := fmt.Sprintf("/v2/stocks/%s/quotes", Symbol)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result StockQuotesRespSingle
	return &result, json.Unmarshal(data, &result)
}

type StockLatestQuoteSingleParams struct {
	Feed     string
	Currency string
}

func (p *StockLatestQuoteSingleParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	if p.Currency != "" {
		v.Set("currency", p.Currency)
	}
	return v
}

// StockLatestQuoteSingle — Latest quote (single symbol)
func (c *MarketDataClient) StockLatestQuoteSingle(Symbol string, params *StockLatestQuoteSingleParams) (*StockLatestQuotesRespSingle, error) {
	path := fmt.Sprintf("/v2/stocks/%s/quotes/latest", Symbol)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result StockLatestQuotesRespSingle
	return &result, json.Unmarshal(data, &result)
}

type StockSnapshotSingleParams struct {
	Feed     string
	Currency string
}

func (p *StockSnapshotSingleParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	if p.Currency != "" {
		v.Set("currency", p.Currency)
	}
	return v
}

// StockSnapshotSingle — Snapshot (single symbol)
func (c *MarketDataClient) StockSnapshotSingle(Symbol string, params *StockSnapshotSingleParams) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/stocks/%s/snapshot", Symbol)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	return data, nil
}

type StockTradeSingleParams struct {
	Start     string
	End       string
	Limit     int
	Asof      string
	Feed      string
	Currency  string
	PageToken string
	Sort      string
}

func (p *StockTradeSingleParams) Values() url.Values {
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
	if p.Limit != 0 {
		v.Set("limit", fmt.Sprint(p.Limit))
	}
	if p.Asof != "" {
		v.Set("asof", p.Asof)
	}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	if p.Currency != "" {
		v.Set("currency", p.Currency)
	}
	if p.PageToken != "" {
		v.Set("page_token", p.PageToken)
	}
	if p.Sort != "" {
		v.Set("sort", p.Sort)
	}
	return v
}

// StockTradeSingle — Historical trades (single symbol)
func (c *MarketDataClient) StockTradeSingle(Symbol string, params *StockTradeSingleParams) (*StockTradesRespSingle, error) {
	path := fmt.Sprintf("/v2/stocks/%s/trades", Symbol)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result StockTradesRespSingle
	return &result, json.Unmarshal(data, &result)
}

type StockLatestTradeSingleParams struct {
	Feed     string
	Currency string
}

func (p *StockLatestTradeSingleParams) Values() url.Values {
	if p == nil {
		return nil
	}
	v := url.Values{}
	if p.Feed != "" {
		v.Set("feed", p.Feed)
	}
	if p.Currency != "" {
		v.Set("currency", p.Currency)
	}
	return v
}

// StockLatestTradeSingle — Latest trade (single symbol)
func (c *MarketDataClient) StockLatestTradeSingle(Symbol string, params *StockLatestTradeSingleParams) (*StockLatestTradesRespSingle, error) {
	path := fmt.Sprintf("/v2/stocks/%s/trades/latest", Symbol)
	data, err := c.Raw.GetData(path, params.Values())
	if err != nil {
		return nil, err
	}
	var result StockLatestTradesRespSingle
	return &result, json.Unmarshal(data, &result)
}
