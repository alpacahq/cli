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

// CorporateActions — Corporate actions
func (c *MarketDataClient) CorporateActions(params url.Values) (*CorporateActionsResp, error) {
	return unmarshal[CorporateActionsResp](c.Raw.GetData("/v1/corporate-actions", params))
}

// CryptoPerpLatestBars — Latest bars
func (c *MarketDataClient) CryptoPerpLatestBars(Loc string, params url.Values) (*CryptoLatestBarsResp, error) {
	return unmarshal[CryptoLatestBarsResp](c.Raw.GetData(fmt.Sprintf("/v1beta1/crypto-perps/%s/latest/bars", url.PathEscape(Loc)), params))
}

// CryptoPerpLatestOrderbooks — Latest orderbook
func (c *MarketDataClient) CryptoPerpLatestOrderbooks(Loc string, params url.Values) (*CryptoLatestOrderbooksResp, error) {
	return unmarshal[CryptoLatestOrderbooksResp](c.Raw.GetData(fmt.Sprintf("/v1beta1/crypto-perps/%s/latest/orderbooks", url.PathEscape(Loc)), params))
}

// CryptoPerpLatestFuturesPricing — Latest pricing
func (c *MarketDataClient) CryptoPerpLatestFuturesPricing(Loc string, params url.Values) (*CryptoPerpLatestFuturesPricingResp, error) {
	return unmarshal[CryptoPerpLatestFuturesPricingResp](c.Raw.GetData(fmt.Sprintf("/v1beta1/crypto-perps/%s/latest/pricing", url.PathEscape(Loc)), params))
}

// CryptoPerpLatestQuotes — Latest quotes
func (c *MarketDataClient) CryptoPerpLatestQuotes(Loc string, params url.Values) (*CryptoLatestQuotesResp, error) {
	return unmarshal[CryptoLatestQuotesResp](c.Raw.GetData(fmt.Sprintf("/v1beta1/crypto-perps/%s/latest/quotes", url.PathEscape(Loc)), params))
}

// CryptoPerpLatestTrades — Latest trades
func (c *MarketDataClient) CryptoPerpLatestTrades(Loc string, params url.Values) (*CryptoLatestTradesResp, error) {
	return unmarshal[CryptoLatestTradesResp](c.Raw.GetData(fmt.Sprintf("/v1beta1/crypto-perps/%s/latest/trades", url.PathEscape(Loc)), params))
}

// FixedIncomeLatestPrices — Latest prices
func (c *MarketDataClient) FixedIncomeLatestPrices(params url.Values) (*FixedIncomeLatestPricesResp, error) {
	return unmarshal[FixedIncomeLatestPricesResp](c.Raw.GetData("/v1beta1/fixed_income/latest/prices", params))
}

// LatestRates — Latest rates for currency pairs
func (c *MarketDataClient) LatestRates(params url.Values) (*ForexLatestRatesResp, error) {
	return unmarshal[ForexLatestRatesResp](c.Raw.GetData("/v1beta1/forex/latest/rates", params))
}

// Rates — Historical rates for currency pairs
func (c *MarketDataClient) Rates(params url.Values) (*ForexRatesResp, error) {
	return unmarshal[ForexRatesResp](c.Raw.GetData("/v1beta1/forex/rates", params))
}

// Logos — Logos
func (c *MarketDataClient) Logos(Symbol string, params url.Values) (json.RawMessage, error) {
	return c.Raw.GetData(fmt.Sprintf("/v1beta1/logos/%s", url.PathEscape(Symbol)), params)
}

// News — News articles
func (c *MarketDataClient) News(params url.Values) (*NewsResp, error) {
	return unmarshal[NewsResp](c.Raw.GetData("/v1beta1/news", params))
}

// OptionBars — Historical bars
func (c *MarketDataClient) OptionBars(params url.Values) (*OptionBarsResp, error) {
	return unmarshal[OptionBarsResp](c.Raw.GetData("/v1beta1/options/bars", params))
}

// OptionMetaConditions — Condition codes
func (c *MarketDataClient) OptionMetaConditions(Ticktype string) (json.RawMessage, error) {
	return c.Raw.GetData(fmt.Sprintf("/v1beta1/options/meta/conditions/%s", url.PathEscape(Ticktype)), nil)
}

// OptionMetaExchanges — Exchange codes
func (c *MarketDataClient) OptionMetaExchanges() (json.RawMessage, error) {
	return c.Raw.GetData("/v1beta1/options/meta/exchanges", nil)
}

// OptionLatestQuotes — Latest quotes
func (c *MarketDataClient) OptionLatestQuotes(params url.Values) (*OptionLatestQuotesResp, error) {
	return unmarshal[OptionLatestQuotesResp](c.Raw.GetData("/v1beta1/options/quotes/latest", params))
}

// OptionSnapshots — Snapshots
func (c *MarketDataClient) OptionSnapshots(params url.Values) (*OptionSnapshotsResp, error) {
	return unmarshal[OptionSnapshotsResp](c.Raw.GetData("/v1beta1/options/snapshots", params))
}

// OptionChain — Option chain
func (c *MarketDataClient) OptionChain(UnderlyingSymbol string, params url.Values) (*OptionSnapshotsResp, error) {
	return unmarshal[OptionSnapshotsResp](c.Raw.GetData(fmt.Sprintf("/v1beta1/options/snapshots/%s", url.PathEscape(UnderlyingSymbol)), params))
}

// OptionTrades — Historical trades
func (c *MarketDataClient) OptionTrades(params url.Values) (*OptionTradesResp, error) {
	return unmarshal[OptionTradesResp](c.Raw.GetData("/v1beta1/options/trades", params))
}

// OptionLatestTrades — Latest trades
func (c *MarketDataClient) OptionLatestTrades(params url.Values) (*OptionLatestTradesResp, error) {
	return unmarshal[OptionLatestTradesResp](c.Raw.GetData("/v1beta1/options/trades/latest", params))
}

// MostActives — Most active stocks
func (c *MarketDataClient) MostActives(params url.Values) (*MostActivesResp, error) {
	return unmarshal[MostActivesResp](c.Raw.GetData("/v1beta1/screener/stocks/most-actives", params))
}

// Movers — Top market movers
func (c *MarketDataClient) Movers(MarketType string, params url.Values) (*MoversResp, error) {
	return unmarshal[MoversResp](c.Raw.GetData(fmt.Sprintf("/v1beta1/screener/%s/movers", url.PathEscape(MarketType)), params))
}

// CryptoBars — Historical bars
func (c *MarketDataClient) CryptoBars(Loc string, params url.Values) (*CryptoBarsResp, error) {
	return unmarshal[CryptoBarsResp](c.Raw.GetData(fmt.Sprintf("/v1beta3/crypto/%s/bars", url.PathEscape(Loc)), params))
}

// CryptoLatestBars — Latest bars
func (c *MarketDataClient) CryptoLatestBars(Loc string, params url.Values) (*CryptoLatestBarsResp, error) {
	return unmarshal[CryptoLatestBarsResp](c.Raw.GetData(fmt.Sprintf("/v1beta3/crypto/%s/latest/bars", url.PathEscape(Loc)), params))
}

// CryptoLatestOrderbooks — Latest orderbook
func (c *MarketDataClient) CryptoLatestOrderbooks(Loc string, params url.Values) (*CryptoLatestOrderbooksResp, error) {
	return unmarshal[CryptoLatestOrderbooksResp](c.Raw.GetData(fmt.Sprintf("/v1beta3/crypto/%s/latest/orderbooks", url.PathEscape(Loc)), params))
}

// CryptoLatestQuotes — Latest quotes
func (c *MarketDataClient) CryptoLatestQuotes(Loc string, params url.Values) (*CryptoLatestQuotesResp, error) {
	return unmarshal[CryptoLatestQuotesResp](c.Raw.GetData(fmt.Sprintf("/v1beta3/crypto/%s/latest/quotes", url.PathEscape(Loc)), params))
}

// CryptoLatestTrades — Latest trades
func (c *MarketDataClient) CryptoLatestTrades(Loc string, params url.Values) (*CryptoLatestTradesResp, error) {
	return unmarshal[CryptoLatestTradesResp](c.Raw.GetData(fmt.Sprintf("/v1beta3/crypto/%s/latest/trades", url.PathEscape(Loc)), params))
}

// CryptoQuotes — Historical quotes
func (c *MarketDataClient) CryptoQuotes(Loc string, params url.Values) (*CryptoQuotesResp, error) {
	return unmarshal[CryptoQuotesResp](c.Raw.GetData(fmt.Sprintf("/v1beta3/crypto/%s/quotes", url.PathEscape(Loc)), params))
}

// CryptoSnapshots — Snapshots
func (c *MarketDataClient) CryptoSnapshots(Loc string, params url.Values) (*CryptoSnapshotsResp, error) {
	return unmarshal[CryptoSnapshotsResp](c.Raw.GetData(fmt.Sprintf("/v1beta3/crypto/%s/snapshots", url.PathEscape(Loc)), params))
}

// CryptoTrades — Historical trades
func (c *MarketDataClient) CryptoTrades(Loc string, params url.Values) (*CryptoTradesResp, error) {
	return unmarshal[CryptoTradesResp](c.Raw.GetData(fmt.Sprintf("/v1beta3/crypto/%s/trades", url.PathEscape(Loc)), params))
}

// StockAuctions — Historical auctions
func (c *MarketDataClient) StockAuctions(params url.Values) (*StockAuctionsResp, error) {
	return unmarshal[StockAuctionsResp](c.Raw.GetData("/v2/stocks/auctions", params))
}

// StockBars — Historical bars
func (c *MarketDataClient) StockBars(params url.Values) (*StockBarsResp, error) {
	return unmarshal[StockBarsResp](c.Raw.GetData("/v2/stocks/bars", params))
}

// StockLatestBars — Latest bars
func (c *MarketDataClient) StockLatestBars(params url.Values) (*StockLatestBarsResp, error) {
	return unmarshal[StockLatestBarsResp](c.Raw.GetData("/v2/stocks/bars/latest", params))
}

// StockMetaConditions — Condition codes
func (c *MarketDataClient) StockMetaConditions(Ticktype string, params url.Values) (json.RawMessage, error) {
	return c.Raw.GetData(fmt.Sprintf("/v2/stocks/meta/conditions/%s", url.PathEscape(Ticktype)), params)
}

// StockMetaExchanges — Exchange codes
func (c *MarketDataClient) StockMetaExchanges() (json.RawMessage, error) {
	return c.Raw.GetData("/v2/stocks/meta/exchanges", nil)
}

// StockQuotes — Historical quotes
func (c *MarketDataClient) StockQuotes(params url.Values) (*StockQuotesResp, error) {
	return unmarshal[StockQuotesResp](c.Raw.GetData("/v2/stocks/quotes", params))
}

// StockLatestQuotes — Latest quotes
func (c *MarketDataClient) StockLatestQuotes(params url.Values) (*StockLatestQuotesResp, error) {
	return unmarshal[StockLatestQuotesResp](c.Raw.GetData("/v2/stocks/quotes/latest", params))
}

// StockSnapshots — Snapshots
func (c *MarketDataClient) StockSnapshots(params url.Values) (json.RawMessage, error) {
	return c.Raw.GetData("/v2/stocks/snapshots", params)
}

// StockTrades — Historical trades
func (c *MarketDataClient) StockTrades(params url.Values) (*StockTradesResp, error) {
	return unmarshal[StockTradesResp](c.Raw.GetData("/v2/stocks/trades", params))
}

// StockLatestTrades — Latest trades
func (c *MarketDataClient) StockLatestTrades(params url.Values) (*StockLatestTradesResp, error) {
	return unmarshal[StockLatestTradesResp](c.Raw.GetData("/v2/stocks/trades/latest", params))
}

// StockAuctionSingle — Historical auctions (single)
func (c *MarketDataClient) StockAuctionSingle(Symbol string, params url.Values) (*StockAuctionsRespSingle, error) {
	return unmarshal[StockAuctionsRespSingle](c.Raw.GetData(fmt.Sprintf("/v2/stocks/%s/auctions", url.PathEscape(Symbol)), params))
}

// StockBarSingle — Historical bars (single symbol)
func (c *MarketDataClient) StockBarSingle(Symbol string, params url.Values) (*StockBarsRespSingle, error) {
	return unmarshal[StockBarsRespSingle](c.Raw.GetData(fmt.Sprintf("/v2/stocks/%s/bars", url.PathEscape(Symbol)), params))
}

// StockLatestBarSingle — Latest bar (single symbol)
func (c *MarketDataClient) StockLatestBarSingle(Symbol string, params url.Values) (*StockLatestBarsRespSingle, error) {
	return unmarshal[StockLatestBarsRespSingle](c.Raw.GetData(fmt.Sprintf("/v2/stocks/%s/bars/latest", url.PathEscape(Symbol)), params))
}

// StockQuoteSingle — Historical quotes (single symbol)
func (c *MarketDataClient) StockQuoteSingle(Symbol string, params url.Values) (*StockQuotesRespSingle, error) {
	return unmarshal[StockQuotesRespSingle](c.Raw.GetData(fmt.Sprintf("/v2/stocks/%s/quotes", url.PathEscape(Symbol)), params))
}

// StockLatestQuoteSingle — Latest quote (single symbol)
func (c *MarketDataClient) StockLatestQuoteSingle(Symbol string, params url.Values) (*StockLatestQuotesRespSingle, error) {
	return unmarshal[StockLatestQuotesRespSingle](c.Raw.GetData(fmt.Sprintf("/v2/stocks/%s/quotes/latest", url.PathEscape(Symbol)), params))
}

// StockSnapshotSingle — Snapshot (single symbol)
func (c *MarketDataClient) StockSnapshotSingle(Symbol string, params url.Values) (json.RawMessage, error) {
	return c.Raw.GetData(fmt.Sprintf("/v2/stocks/%s/snapshot", url.PathEscape(Symbol)), params)
}

// StockTradeSingle — Historical trades (single symbol)
func (c *MarketDataClient) StockTradeSingle(Symbol string, params url.Values) (*StockTradesRespSingle, error) {
	return unmarshal[StockTradesRespSingle](c.Raw.GetData(fmt.Sprintf("/v2/stocks/%s/trades", url.PathEscape(Symbol)), params))
}

// StockLatestTradeSingle — Latest trade (single symbol)
func (c *MarketDataClient) StockLatestTradeSingle(Symbol string, params url.Values) (*StockLatestTradesRespSingle, error) {
	return unmarshal[StockLatestTradesRespSingle](c.Raw.GetData(fmt.Sprintf("/v2/stocks/%s/trades/latest", url.PathEscape(Symbol)), params))
}
