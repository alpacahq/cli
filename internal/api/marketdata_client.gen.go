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
	path := "/v1/corporate-actions"
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result CorporateActionsResp
	return &result, json.Unmarshal(data, &result)
}

// CryptoPerpLatestBars — Latest bars
func (c *MarketDataClient) CryptoPerpLatestBars(Loc string, params url.Values) (*CryptoLatestBarsResp, error) {
	path := fmt.Sprintf("/v1beta1/crypto-perps/%s/latest/bars", url.PathEscape(Loc))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result CryptoLatestBarsResp
	return &result, json.Unmarshal(data, &result)
}

// CryptoPerpLatestOrderbooks — Latest orderbook
func (c *MarketDataClient) CryptoPerpLatestOrderbooks(Loc string, params url.Values) (*CryptoLatestOrderbooksResp, error) {
	path := fmt.Sprintf("/v1beta1/crypto-perps/%s/latest/orderbooks", url.PathEscape(Loc))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result CryptoLatestOrderbooksResp
	return &result, json.Unmarshal(data, &result)
}

// CryptoPerpLatestFuturesPricing — Latest pricing
func (c *MarketDataClient) CryptoPerpLatestFuturesPricing(Loc string, params url.Values) (*CryptoPerpLatestFuturesPricingResp, error) {
	path := fmt.Sprintf("/v1beta1/crypto-perps/%s/latest/pricing", url.PathEscape(Loc))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result CryptoPerpLatestFuturesPricingResp
	return &result, json.Unmarshal(data, &result)
}

// CryptoPerpLatestQuotes — Latest quotes
func (c *MarketDataClient) CryptoPerpLatestQuotes(Loc string, params url.Values) (*CryptoLatestQuotesResp, error) {
	path := fmt.Sprintf("/v1beta1/crypto-perps/%s/latest/quotes", url.PathEscape(Loc))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result CryptoLatestQuotesResp
	return &result, json.Unmarshal(data, &result)
}

// CryptoPerpLatestTrades — Latest trades
func (c *MarketDataClient) CryptoPerpLatestTrades(Loc string, params url.Values) (*CryptoLatestTradesResp, error) {
	path := fmt.Sprintf("/v1beta1/crypto-perps/%s/latest/trades", url.PathEscape(Loc))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result CryptoLatestTradesResp
	return &result, json.Unmarshal(data, &result)
}

// FixedIncomeLatestPrices — Latest prices
func (c *MarketDataClient) FixedIncomeLatestPrices(params url.Values) (*FixedIncomeLatestPricesResp, error) {
	path := "/v1beta1/fixed_income/latest/prices"
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result FixedIncomeLatestPricesResp
	return &result, json.Unmarshal(data, &result)
}

// LatestRates — Latest rates for currency pairs
func (c *MarketDataClient) LatestRates(params url.Values) (*ForexLatestRatesResp, error) {
	path := "/v1beta1/forex/latest/rates"
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result ForexLatestRatesResp
	return &result, json.Unmarshal(data, &result)
}

// Rates — Historical rates for currency pairs
func (c *MarketDataClient) Rates(params url.Values) (*ForexRatesResp, error) {
	path := "/v1beta1/forex/rates"
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result ForexRatesResp
	return &result, json.Unmarshal(data, &result)
}

// Logos — Logos
func (c *MarketDataClient) Logos(Symbol string, params url.Values) (json.RawMessage, error) {
	path := fmt.Sprintf("/v1beta1/logos/%s", url.PathEscape(Symbol))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// News — News articles
func (c *MarketDataClient) News(params url.Values) (*NewsResp, error) {
	path := "/v1beta1/news"
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result NewsResp
	return &result, json.Unmarshal(data, &result)
}

// OptionBars — Historical bars
func (c *MarketDataClient) OptionBars(params url.Values) (*OptionBarsResp, error) {
	path := "/v1beta1/options/bars"
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result OptionBarsResp
	return &result, json.Unmarshal(data, &result)
}

// OptionMetaConditions — Condition codes
func (c *MarketDataClient) OptionMetaConditions(Ticktype string) (json.RawMessage, error) {
	path := fmt.Sprintf("/v1beta1/options/meta/conditions/%s", url.PathEscape(Ticktype))
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

// OptionLatestQuotes — Latest quotes
func (c *MarketDataClient) OptionLatestQuotes(params url.Values) (*OptionLatestQuotesResp, error) {
	path := "/v1beta1/options/quotes/latest"
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result OptionLatestQuotesResp
	return &result, json.Unmarshal(data, &result)
}

// OptionSnapshots — Snapshots
func (c *MarketDataClient) OptionSnapshots(params url.Values) (*OptionSnapshotsResp, error) {
	path := "/v1beta1/options/snapshots"
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result OptionSnapshotsResp
	return &result, json.Unmarshal(data, &result)
}

// OptionChain — Option chain
func (c *MarketDataClient) OptionChain(UnderlyingSymbol string, params url.Values) (*OptionSnapshotsResp, error) {
	path := fmt.Sprintf("/v1beta1/options/snapshots/%s", url.PathEscape(UnderlyingSymbol))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result OptionSnapshotsResp
	return &result, json.Unmarshal(data, &result)
}

// OptionTrades — Historical trades
func (c *MarketDataClient) OptionTrades(params url.Values) (*OptionTradesResp, error) {
	path := "/v1beta1/options/trades"
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result OptionTradesResp
	return &result, json.Unmarshal(data, &result)
}

// OptionLatestTrades — Latest trades
func (c *MarketDataClient) OptionLatestTrades(params url.Values) (*OptionLatestTradesResp, error) {
	path := "/v1beta1/options/trades/latest"
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result OptionLatestTradesResp
	return &result, json.Unmarshal(data, &result)
}

// MostActives — Most active stocks
func (c *MarketDataClient) MostActives(params url.Values) (*MostActivesResp, error) {
	path := "/v1beta1/screener/stocks/most-actives"
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result MostActivesResp
	return &result, json.Unmarshal(data, &result)
}

// Movers — Top market movers
func (c *MarketDataClient) Movers(MarketType string, params url.Values) (*MoversResp, error) {
	path := fmt.Sprintf("/v1beta1/screener/%s/movers", url.PathEscape(MarketType))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result MoversResp
	return &result, json.Unmarshal(data, &result)
}

// CryptoBars — Historical bars
func (c *MarketDataClient) CryptoBars(Loc string, params url.Values) (*CryptoBarsResp, error) {
	path := fmt.Sprintf("/v1beta3/crypto/%s/bars", url.PathEscape(Loc))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result CryptoBarsResp
	return &result, json.Unmarshal(data, &result)
}

// CryptoLatestBars — Latest bars
func (c *MarketDataClient) CryptoLatestBars(Loc string, params url.Values) (*CryptoLatestBarsResp, error) {
	path := fmt.Sprintf("/v1beta3/crypto/%s/latest/bars", url.PathEscape(Loc))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result CryptoLatestBarsResp
	return &result, json.Unmarshal(data, &result)
}

// CryptoLatestOrderbooks — Latest orderbook
func (c *MarketDataClient) CryptoLatestOrderbooks(Loc string, params url.Values) (*CryptoLatestOrderbooksResp, error) {
	path := fmt.Sprintf("/v1beta3/crypto/%s/latest/orderbooks", url.PathEscape(Loc))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result CryptoLatestOrderbooksResp
	return &result, json.Unmarshal(data, &result)
}

// CryptoLatestQuotes — Latest quotes
func (c *MarketDataClient) CryptoLatestQuotes(Loc string, params url.Values) (*CryptoLatestQuotesResp, error) {
	path := fmt.Sprintf("/v1beta3/crypto/%s/latest/quotes", url.PathEscape(Loc))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result CryptoLatestQuotesResp
	return &result, json.Unmarshal(data, &result)
}

// CryptoLatestTrades — Latest trades
func (c *MarketDataClient) CryptoLatestTrades(Loc string, params url.Values) (*CryptoLatestTradesResp, error) {
	path := fmt.Sprintf("/v1beta3/crypto/%s/latest/trades", url.PathEscape(Loc))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result CryptoLatestTradesResp
	return &result, json.Unmarshal(data, &result)
}

// CryptoQuotes — Historical quotes
func (c *MarketDataClient) CryptoQuotes(Loc string, params url.Values) (*CryptoQuotesResp, error) {
	path := fmt.Sprintf("/v1beta3/crypto/%s/quotes", url.PathEscape(Loc))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result CryptoQuotesResp
	return &result, json.Unmarshal(data, &result)
}

// CryptoSnapshots — Snapshots
func (c *MarketDataClient) CryptoSnapshots(Loc string, params url.Values) (*CryptoSnapshotsResp, error) {
	path := fmt.Sprintf("/v1beta3/crypto/%s/snapshots", url.PathEscape(Loc))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result CryptoSnapshotsResp
	return &result, json.Unmarshal(data, &result)
}

// CryptoTrades — Historical trades
func (c *MarketDataClient) CryptoTrades(Loc string, params url.Values) (*CryptoTradesResp, error) {
	path := fmt.Sprintf("/v1beta3/crypto/%s/trades", url.PathEscape(Loc))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result CryptoTradesResp
	return &result, json.Unmarshal(data, &result)
}

// StockAuctions — Historical auctions
func (c *MarketDataClient) StockAuctions(params url.Values) (*StockAuctionsResp, error) {
	path := "/v2/stocks/auctions"
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result StockAuctionsResp
	return &result, json.Unmarshal(data, &result)
}

// StockBars — Historical bars
func (c *MarketDataClient) StockBars(params url.Values) (*StockBarsResp, error) {
	path := "/v2/stocks/bars"
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result StockBarsResp
	return &result, json.Unmarshal(data, &result)
}

// StockLatestBars — Latest bars
func (c *MarketDataClient) StockLatestBars(params url.Values) (*StockLatestBarsResp, error) {
	path := "/v2/stocks/bars/latest"
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result StockLatestBarsResp
	return &result, json.Unmarshal(data, &result)
}

// StockMetaConditions — Condition codes
func (c *MarketDataClient) StockMetaConditions(Ticktype string, params url.Values) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/stocks/meta/conditions/%s", url.PathEscape(Ticktype))
	data, err := c.Raw.GetData(path, params)
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

// StockQuotes — Historical quotes
func (c *MarketDataClient) StockQuotes(params url.Values) (*StockQuotesResp, error) {
	path := "/v2/stocks/quotes"
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result StockQuotesResp
	return &result, json.Unmarshal(data, &result)
}

// StockLatestQuotes — Latest quotes
func (c *MarketDataClient) StockLatestQuotes(params url.Values) (*StockLatestQuotesResp, error) {
	path := "/v2/stocks/quotes/latest"
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result StockLatestQuotesResp
	return &result, json.Unmarshal(data, &result)
}

// StockSnapshots — Snapshots
func (c *MarketDataClient) StockSnapshots(params url.Values) (json.RawMessage, error) {
	path := "/v2/stocks/snapshots"
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// StockTrades — Historical trades
func (c *MarketDataClient) StockTrades(params url.Values) (*StockTradesResp, error) {
	path := "/v2/stocks/trades"
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result StockTradesResp
	return &result, json.Unmarshal(data, &result)
}

// StockLatestTrades — Latest trades
func (c *MarketDataClient) StockLatestTrades(params url.Values) (*StockLatestTradesResp, error) {
	path := "/v2/stocks/trades/latest"
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result StockLatestTradesResp
	return &result, json.Unmarshal(data, &result)
}

// StockAuctionSingle — Historical auctions (single)
func (c *MarketDataClient) StockAuctionSingle(Symbol string, params url.Values) (*StockAuctionsRespSingle, error) {
	path := fmt.Sprintf("/v2/stocks/%s/auctions", url.PathEscape(Symbol))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result StockAuctionsRespSingle
	return &result, json.Unmarshal(data, &result)
}

// StockBarSingle — Historical bars (single symbol)
func (c *MarketDataClient) StockBarSingle(Symbol string, params url.Values) (*StockBarsRespSingle, error) {
	path := fmt.Sprintf("/v2/stocks/%s/bars", url.PathEscape(Symbol))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result StockBarsRespSingle
	return &result, json.Unmarshal(data, &result)
}

// StockLatestBarSingle — Latest bar (single symbol)
func (c *MarketDataClient) StockLatestBarSingle(Symbol string, params url.Values) (*StockLatestBarsRespSingle, error) {
	path := fmt.Sprintf("/v2/stocks/%s/bars/latest", url.PathEscape(Symbol))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result StockLatestBarsRespSingle
	return &result, json.Unmarshal(data, &result)
}

// StockQuoteSingle — Historical quotes (single symbol)
func (c *MarketDataClient) StockQuoteSingle(Symbol string, params url.Values) (*StockQuotesRespSingle, error) {
	path := fmt.Sprintf("/v2/stocks/%s/quotes", url.PathEscape(Symbol))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result StockQuotesRespSingle
	return &result, json.Unmarshal(data, &result)
}

// StockLatestQuoteSingle — Latest quote (single symbol)
func (c *MarketDataClient) StockLatestQuoteSingle(Symbol string, params url.Values) (*StockLatestQuotesRespSingle, error) {
	path := fmt.Sprintf("/v2/stocks/%s/quotes/latest", url.PathEscape(Symbol))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result StockLatestQuotesRespSingle
	return &result, json.Unmarshal(data, &result)
}

// StockSnapshotSingle — Snapshot (single symbol)
func (c *MarketDataClient) StockSnapshotSingle(Symbol string, params url.Values) (json.RawMessage, error) {
	path := fmt.Sprintf("/v2/stocks/%s/snapshot", url.PathEscape(Symbol))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// StockTradeSingle — Historical trades (single symbol)
func (c *MarketDataClient) StockTradeSingle(Symbol string, params url.Values) (*StockTradesRespSingle, error) {
	path := fmt.Sprintf("/v2/stocks/%s/trades", url.PathEscape(Symbol))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result StockTradesRespSingle
	return &result, json.Unmarshal(data, &result)
}

// StockLatestTradeSingle — Latest trade (single symbol)
func (c *MarketDataClient) StockLatestTradeSingle(Symbol string, params url.Values) (*StockLatestTradesRespSingle, error) {
	path := fmt.Sprintf("/v2/stocks/%s/trades/latest", url.PathEscape(Symbol))
	data, err := c.Raw.GetData(path, params)
	if err != nil {
		return nil, err
	}
	var result StockLatestTradesRespSingle
	return &result, json.Unmarshal(data, &result)
}
