// Code generated from api/specs; DO NOT EDIT.

package api

// Op describes a generated API operation. Passed to fetchCmd/attachCmd
// for automatic flag registration, help text, and required-flag validation.
type Op struct {
	Name          string
	summary       string
	flags         []FlagDef
	requiredFlags []string
}

func (o Op) Summary() string         { return o.summary }
func (o Op) Flags() []FlagDef        { return o.flags }
func (o Op) RequiredFlags() []string { return o.requiredFlags }

// FlagDef describes a CLI flag derived from the OpenAPI spec.
type FlagDef struct {
	Name        string // kebab-case CLI flag name
	OASName     string // original OAS property/parameter name
	Type        string // "string", "bool", "int"
	Default     string
	Description string
	Completions []string // enum values for shell completion
	OpName      string   // operation name for schema lookup
	Required    bool     // true if OAS marks this parameter as required
	Source      string   // "path", "query", or "body"
}

var CalendarOp = Op{
	Name: "Calendar", summary: "Get market calendar",
	requiredFlags: []string{"market"},
	flags: []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "last date to retrieve data for (inclusive). Default: one week from the start date", OpName: "Calendar", Source: "query"},
		{Name: "market", OASName: "market", Type: "string", Description: "market identifier (MIC, BIC, or acronym)", OpName: "Calendar", Required: true, Source: "path"},
		{Name: "start", OASName: "start", Type: "string", Description: "first date to retrieve data for (inclusive). Default: today", OpName: "Calendar", Source: "query"},
		{Name: "timezone", OASName: "timezone", Type: "string", Description: "timezone of the times. Default: the timezone of the market", Completions: []string{"UTC"}, OpName: "Calendar", Source: "query"},
	},
}

var ClockOp = Op{
	Name: "Clock", summary: "Get market clock",
	flags: []FlagDef{
		{Name: "markets", OASName: "markets", Type: "string", Description: "comma-separated list of markets", OpName: "Clock", Source: "query"},
		{Name: "time", OASName: "time", Type: "string", Description: "instead of the current time, use this time for the clock", OpName: "Clock", Source: "query"},
	},
}

var CorporateActionsOp = Op{
	Name: "CorporateActions", summary: "Get corporate actions",
	flags: []FlagDef{
		{Name: "cusips", OASName: "cusips", Type: "string", Description: "A comma-separated list of CUSIPs", OpName: "CorporateActions", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "CorporateActions", Source: "query"},
		{Name: "ids", OASName: "ids", Type: "string", Description: "A comma-separated list of corporate action IDs", OpName: "CorporateActions", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "100", Description: "maximum number of corporate actions to return in a response.", OpName: "CorporateActions", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "CorporateActions", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "CorporateActions", Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "CorporateActions", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of symbols", OpName: "CorporateActions", Source: "query"},
		{Name: "types", OASName: "types", Type: "string", Description: "A comma-separated list of types", OpName: "CorporateActions", Source: "query"},
	},
}

var CryptoBarsOp = Op{
	Name: "CryptoBars", summary: "Get historical bars",
	requiredFlags: []string{"loc", "symbols", "timeframe"},
	flags: []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "CryptoBars", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "CryptoBars", Source: "query"},
		{Name: "loc", OASName: "loc", Type: "string", Description: "crypto location from where the historical market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU", OpName: "CryptoBars", Required: true, Source: "path"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "CryptoBars", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "CryptoBars", Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "CryptoBars", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoBars", Required: true, Source: "query"},
		{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g", OpName: "CryptoBars", Required: true, Source: "query"},
	},
}

var CryptoLatestBarsOp = Op{
	Name: "CryptoLatestBars", summary: "Get latest bars",
	requiredFlags: []string{"loc", "symbols"},
	flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Description: "crypto location from where the latest market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU", OpName: "CryptoLatestBars", Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoLatestBars", Required: true, Source: "query"},
	},
}

var CryptoLatestOrderbooksOp = Op{
	Name: "CryptoLatestOrderbooks", summary: "Get latest orderbook",
	requiredFlags: []string{"loc", "symbols"},
	flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Description: "crypto location from where the latest market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU", OpName: "CryptoLatestOrderbooks", Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoLatestOrderbooks", Required: true, Source: "query"},
	},
}

var CryptoLatestQuotesOp = Op{
	Name: "CryptoLatestQuotes", summary: "Get latest quotes",
	requiredFlags: []string{"loc", "symbols"},
	flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Description: "crypto location from where the latest market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU", OpName: "CryptoLatestQuotes", Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoLatestQuotes", Required: true, Source: "query"},
	},
}

var CryptoLatestTradesOp = Op{
	Name: "CryptoLatestTrades", summary: "Get latest trades",
	requiredFlags: []string{"loc", "symbols"},
	flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Description: "crypto location from where the latest market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU", OpName: "CryptoLatestTrades", Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoLatestTrades", Required: true, Source: "query"},
	},
}

var CryptoPerpLatestBarsOp = Op{
	Name: "CryptoPerpLatestBars", summary: "Get latest bars",
	requiredFlags: []string{"loc", "symbols"},
	flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Description: "crypto perpetual location", OpName: "CryptoPerpLatestBars", Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoPerpLatestBars", Required: true, Source: "query"},
	},
}

var CryptoPerpLatestFuturesPricingOp = Op{
	Name: "CryptoPerpLatestFuturesPricing", summary: "Get latest pricing",
	requiredFlags: []string{"loc", "symbols"},
	flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Description: "crypto perpetual location", OpName: "CryptoPerpLatestFuturesPricing", Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoPerpLatestFuturesPricing", Required: true, Source: "query"},
	},
}

var CryptoPerpLatestOrderbooksOp = Op{
	Name: "CryptoPerpLatestOrderbooks", summary: "Get latest orderbook",
	requiredFlags: []string{"loc", "symbols"},
	flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Description: "crypto perpetual location", OpName: "CryptoPerpLatestOrderbooks", Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoPerpLatestOrderbooks", Required: true, Source: "query"},
	},
}

var CryptoPerpLatestQuotesOp = Op{
	Name: "CryptoPerpLatestQuotes", summary: "Get latest quotes",
	requiredFlags: []string{"loc", "symbols"},
	flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Description: "crypto perpetual location", OpName: "CryptoPerpLatestQuotes", Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoPerpLatestQuotes", Required: true, Source: "query"},
	},
}

var CryptoPerpLatestTradesOp = Op{
	Name: "CryptoPerpLatestTrades", summary: "Get latest trades",
	requiredFlags: []string{"loc", "symbols"},
	flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Description: "crypto perpetual location", OpName: "CryptoPerpLatestTrades", Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoPerpLatestTrades", Required: true, Source: "query"},
	},
}

var CryptoQuotesOp = Op{
	Name: "CryptoQuotes", summary: "Get historical quotes",
	requiredFlags: []string{"loc", "symbols"},
	flags: []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "CryptoQuotes", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "CryptoQuotes", Source: "query"},
		{Name: "loc", OASName: "loc", Type: "string", Description: "crypto location from where the historical market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU", OpName: "CryptoQuotes", Required: true, Source: "path"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "CryptoQuotes", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "CryptoQuotes", Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "CryptoQuotes", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoQuotes", Required: true, Source: "query"},
	},
}

var CryptoSnapshotsOp = Op{
	Name: "CryptoSnapshots", summary: "Get snapshots",
	requiredFlags: []string{"loc", "symbols"},
	flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Description: "crypto location from where the latest market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU", OpName: "CryptoSnapshots", Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoSnapshots", Required: true, Source: "query"},
	},
}

var CryptoTradesOp = Op{
	Name: "CryptoTrades", summary: "Get historical trades",
	requiredFlags: []string{"loc", "symbols"},
	flags: []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "CryptoTrades", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "CryptoTrades", Source: "query"},
		{Name: "loc", OASName: "loc", Type: "string", Description: "crypto location from where the historical market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU", OpName: "CryptoTrades", Required: true, Source: "path"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "CryptoTrades", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "CryptoTrades", Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "CryptoTrades", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoTrades", Required: true, Source: "query"},
	},
}

var FixedIncomeLatestPricesOp = Op{
	Name: "FixedIncomeLatestPrices", summary: "Get latest prices",
	requiredFlags: []string{"isins"},
	flags: []FlagDef{
		{Name: "isins", OASName: "isins", Type: "string", Description: "A comma-separated list of ISINs with a limit of 1000", OpName: "FixedIncomeLatestPrices", Required: true, Source: "query"},
	},
}

var LatestRatesOp = Op{
	Name: "LatestRates", summary: "Get latest rates for currency pairs",
	requiredFlags: []string{"currency-pairs"},
	flags: []FlagDef{
		{Name: "currency-pairs", OASName: "currency_pairs", Type: "string", Description: "A comma-separated string with currency pairs", OpName: "LatestRates", Required: true, Source: "query"},
	},
}

var LegacyCalendarOp = Op{
	Name: "LegacyCalendar", summary: "Get US market calendar",
	flags: []FlagDef{
		{Name: "date-type", OASName: "date_type", Type: "string", Description: "indicates what start and end mean", Completions: []string{"SETTLEMENT", "TRADING"}, OpName: "LegacyCalendar", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "last date to retrieve data for (inclusive)", OpName: "LegacyCalendar", Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "first date to retrieve data for (inclusive)", OpName: "LegacyCalendar", Source: "query"},
	},
}

var LegacyClockOp = Op{
	Name: "LegacyClock", summary: "Get US market clock",
}

var LogosOp = Op{
	Name: "Logos", summary: "Get logos",
	requiredFlags: []string{"symbol"},
	flags: []FlagDef{
		{Name: "placeholder", OASName: "placeholder", Type: "bool", Default: "true", Description: "if true, returns a placeholder image when no logo is available. Defaults to true", OpName: "Logos", Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "A unique series of letters assigned to a security for trading purposes", OpName: "Logos", Required: true, Source: "path"},
	},
}

var MostActivesOp = Op{
	Name: "MostActives", summary: "Get most active stocks",
	flags: []FlagDef{
		{Name: "by", OASName: "by", Type: "string", Default: "volume", Description: "metric used for ranking the most active stocks", Completions: []string{"trades", "volume"}, OpName: "MostActives", Source: "query"},
		{Name: "top", OASName: "top", Type: "int", Default: "10", Description: "number of top most active stocks to fetch per day", OpName: "MostActives", Source: "query"},
	},
}

var MoversOp = Op{
	Name: "Movers", summary: "Get top market movers",
	requiredFlags: []string{"market-type"},
	flags: []FlagDef{
		{Name: "market-type", OASName: "market_type", Type: "string", Description: "screen-specific market (stocks or crypto)", OpName: "Movers", Required: true, Source: "path"},
		{Name: "top", OASName: "top", Type: "int", Default: "10", Description: "number of top market movers to fetch (gainers and losers)", OpName: "Movers", Source: "query"},
	},
}

var NewsOp = Op{
	Name: "News", summary: "Get news articles",
	flags: []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "News", Source: "query"},
		{Name: "exclude-contentless", OASName: "exclude_contentless", Type: "bool", Description: "boolean indicator to exclude news articles that do not contain content", OpName: "News", Source: "query"},
		{Name: "include-content", OASName: "include_content", Type: "bool", Description: "boolean indicator to include content for news articles (if available)", OpName: "News", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Description: "limit of news items to be returned for a result page", OpName: "News", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "News", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "desc", Description: "sort articles by updated date", Completions: []string{"asc", "desc"}, OpName: "News", Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "News", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of symbols for which to query news", OpName: "News", Source: "query"},
	},
}

var OptionChainOp = Op{
	Name: "OptionChain", summary: "Get option chain",
	requiredFlags: []string{"underlying-symbol"},
	flags: []FlagDef{
		{Name: "expiration-date", OASName: "expiration_date", Type: "string", Description: "filter contracts by the exact expiration date (format: YYYY-MM-DD)", OpName: "OptionChain", Source: "query"},
		{Name: "expiration-date-gte", OASName: "expiration_date_gte", Type: "string", Description: "filter contracts with expiration date greater than or equal to the specified date", OpName: "OptionChain", Source: "query"},
		{Name: "expiration-date-lte", OASName: "expiration_date_lte", Type: "string", Description: "filter contracts with expiration date less than or equal to the specified date", OpName: "OptionChain", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "opra", Description: "source feed of the data", Completions: []string{"indicative", "opra"}, OpName: "OptionChain", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "100", Description: "number of maximum snapshots to return in a response.", OpName: "OptionChain", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "OptionChain", Source: "query"},
		{Name: "root-symbol", OASName: "root_symbol", Type: "string", Description: "filter contracts by the root symbol", OpName: "OptionChain", Source: "query"},
		{Name: "strike-price-gte", OASName: "strike_price_gte", Type: "string", Description: "filter contracts with strike price greater than or equal to the specified value", OpName: "OptionChain", Source: "query"},
		{Name: "strike-price-lte", OASName: "strike_price_lte", Type: "string", Description: "filter contracts with strike price less than or equal to the specified value", OpName: "OptionChain", Source: "query"},
		{Name: "type", OASName: "type", Type: "string", Description: "filter contracts by the type (call or put)", Completions: []string{"call", "put"}, OpName: "OptionChain", Source: "query"},
		{Name: "underlying-symbol", OASName: "underlying_symbol", Type: "string", Description: "financial instrument on which an option contract is based or derived", OpName: "OptionChain", Required: true, Source: "path"},
		{Name: "updated-since", OASName: "updated_since", Type: "string", Description: "filter to snapshots that were updated since this timestamp, meaning that the timestamp of the trade or the quote is g...", OpName: "OptionChain", Source: "query"},
	},
}

var OptionLatestQuotesOp = Op{
	Name: "OptionLatestQuotes", summary: "Get latest quotes",
	requiredFlags: []string{"symbols"},
	flags: []FlagDef{
		{Name: "feed", OASName: "feed", Type: "string", Default: "opra", Description: "source feed of the data", Completions: []string{"indicative", "opra"}, OpName: "OptionLatestQuotes", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100", OpName: "OptionLatestQuotes", Required: true, Source: "query"},
	},
}

var OptionLatestTradesOp = Op{
	Name: "OptionLatestTrades", summary: "Get latest trades",
	requiredFlags: []string{"symbols"},
	flags: []FlagDef{
		{Name: "feed", OASName: "feed", Type: "string", Default: "opra", Description: "source feed of the data", Completions: []string{"indicative", "opra"}, OpName: "OptionLatestTrades", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100", OpName: "OptionLatestTrades", Required: true, Source: "query"},
	},
}

var OptionMetaConditionsOp = Op{
	Name: "OptionMetaConditions", summary: "Get condition codes",
	requiredFlags: []string{"ticktype"},
	flags: []FlagDef{
		{Name: "ticktype", OASName: "ticktype", Type: "string", Description: "type of ticks", OpName: "OptionMetaConditions", Required: true, Source: "path"},
	},
}

var OptionMetaExchangesOp = Op{
	Name: "OptionMetaExchanges", summary: "Get exchange codes",
}

var OptionSnapshotsOp = Op{
	Name: "OptionSnapshots", summary: "Get snapshots",
	requiredFlags: []string{"symbols"},
	flags: []FlagDef{
		{Name: "feed", OASName: "feed", Type: "string", Default: "opra", Description: "source feed of the data", Completions: []string{"indicative", "opra"}, OpName: "OptionSnapshots", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "100", Description: "number of maximum snapshots to return in a response.", OpName: "OptionSnapshots", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "OptionSnapshots", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100", OpName: "OptionSnapshots", Required: true, Source: "query"},
		{Name: "updated-since", OASName: "updated_since", Type: "string", Description: "filter to snapshots that were updated since this timestamp, meaning that the timestamp of the trade or the quote is g...", OpName: "OptionSnapshots", Source: "query"},
	},
}

var OptionTradesOp = Op{
	Name: "OptionTrades", summary: "Get historical trades",
	requiredFlags: []string{"symbols"},
	flags: []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "OptionTrades", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "OptionTrades", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "OptionTrades", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "OptionTrades", Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "OptionTrades", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100", OpName: "OptionTrades", Required: true, Source: "query"},
	},
}

var RatesOp = Op{
	Name: "Rates", summary: "Get historical rates for currency pairs",
	requiredFlags: []string{"currency-pairs"},
	flags: []FlagDef{
		{Name: "currency-pairs", OASName: "currency_pairs", Type: "string", Description: "A comma-separated string with currency pairs", OpName: "Rates", Required: true, Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "Rates", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "Rates", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "Rates", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "Rates", Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "Rates", Source: "query"},
		{Name: "timeframe", OASName: "timeframe", Type: "string", Default: "1Min", Description: "sampling interval of the currency rates", OpName: "Rates", Source: "query"},
	},
}

var StockAuctionSingleOp = Op{
	Name: "StockAuctionSingle", summary: "Get historical auctions (single)",
	requiredFlags: []string{"symbol"},
	flags: []FlagDef{
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", OpName: "StockAuctionSingle", Source: "query"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockAuctionSingle", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "StockAuctionSingle", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "only sip is valid for auctions", OpName: "StockAuctionSingle", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "StockAuctionSingle", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "StockAuctionSingle", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "StockAuctionSingle", Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "StockAuctionSingle", Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol to query", OpName: "StockAuctionSingle", Required: true, Source: "path"},
	},
}

var StockAuctionsOp = Op{
	Name: "StockAuctions", summary: "Get historical auctions",
	requiredFlags: []string{"symbols"},
	flags: []FlagDef{
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", OpName: "StockAuctions", Source: "query"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockAuctions", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "StockAuctions", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "only sip is valid for auctions", OpName: "StockAuctions", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "StockAuctions", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "StockAuctions", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "StockAuctions", Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "StockAuctions", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", OpName: "StockAuctions", Required: true, Source: "query"},
	},
}

var StockBarSingleOp = Op{
	Name: "StockBarSingle", summary: "Get historical bars (single symbol)",
	requiredFlags: []string{"symbol", "timeframe"},
	flags: []FlagDef{
		{Name: "adjustment", OASName: "adjustment", Type: "string", Default: "raw", Description: "specifies the adjustments for the bars.\n\n - raw: no adjustments\n - split: adjust price and volume for forward and rev...", OpName: "StockBarSingle", Source: "query"},
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", OpName: "StockBarSingle", Source: "query"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockBarSingle", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "StockBarSingle", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "source feed of the data.", Completions: []string{"boats", "iex", "otc", "sip"}, OpName: "StockBarSingle", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "StockBarSingle", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "StockBarSingle", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "StockBarSingle", Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "StockBarSingle", Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol to query", OpName: "StockBarSingle", Required: true, Source: "path"},
		{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g", OpName: "StockBarSingle", Required: true, Source: "query"},
	},
}

var StockBarsOp = Op{
	Name: "StockBars", summary: "Get historical bars",
	requiredFlags: []string{"symbols", "timeframe"},
	flags: []FlagDef{
		{Name: "adjustment", OASName: "adjustment", Type: "string", Default: "raw", Description: "specifies the adjustments for the bars.\n\n - raw: no adjustments\n - split: adjust price and volume for forward and rev...", OpName: "StockBars", Source: "query"},
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", OpName: "StockBars", Source: "query"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockBars", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "StockBars", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "source feed of the data.", Completions: []string{"boats", "iex", "otc", "sip"}, OpName: "StockBars", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "StockBars", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "StockBars", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "StockBars", Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "StockBars", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", OpName: "StockBars", Required: true, Source: "query"},
		{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g", OpName: "StockBars", Required: true, Source: "query"},
	},
}

var StockLatestBarSingleOp = Op{
	Name: "StockLatestBarSingle", summary: "Get latest bar (single symbol)",
	requiredFlags: []string{"symbol"},
	flags: []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockLatestBarSingle", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, OpName: "StockLatestBarSingle", Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol to query", OpName: "StockLatestBarSingle", Required: true, Source: "path"},
	},
}

var StockLatestBarsOp = Op{
	Name: "StockLatestBars", summary: "Get latest bars",
	requiredFlags: []string{"symbols"},
	flags: []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockLatestBars", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, OpName: "StockLatestBars", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", OpName: "StockLatestBars", Required: true, Source: "query"},
	},
}

var StockLatestQuoteSingleOp = Op{
	Name: "StockLatestQuoteSingle", summary: "Get latest quote (single symbol)",
	requiredFlags: []string{"symbol"},
	flags: []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockLatestQuoteSingle", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, OpName: "StockLatestQuoteSingle", Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol to query", OpName: "StockLatestQuoteSingle", Required: true, Source: "path"},
	},
}

var StockLatestQuotesOp = Op{
	Name: "StockLatestQuotes", summary: "Get latest quotes",
	requiredFlags: []string{"symbols"},
	flags: []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockLatestQuotes", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, OpName: "StockLatestQuotes", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", OpName: "StockLatestQuotes", Required: true, Source: "query"},
	},
}

var StockLatestTradeSingleOp = Op{
	Name: "StockLatestTradeSingle", summary: "Get latest trade (single symbol)",
	requiredFlags: []string{"symbol"},
	flags: []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockLatestTradeSingle", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, OpName: "StockLatestTradeSingle", Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol to query", OpName: "StockLatestTradeSingle", Required: true, Source: "path"},
	},
}

var StockLatestTradesOp = Op{
	Name: "StockLatestTrades", summary: "Get latest trades",
	requiredFlags: []string{"symbols"},
	flags: []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockLatestTrades", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, OpName: "StockLatestTrades", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", OpName: "StockLatestTrades", Required: true, Source: "query"},
	},
}

var StockMetaConditionsOp = Op{
	Name: "StockMetaConditions", summary: "Get condition codes",
	requiredFlags: []string{"tape", "ticktype"},
	flags: []FlagDef{
		{Name: "tape", OASName: "tape", Type: "string", Description: "one character name of the tape", Completions: []string{"A", "B", "C"}, OpName: "StockMetaConditions", Required: true, Source: "query"},
		{Name: "ticktype", OASName: "ticktype", Type: "string", Description: "type of ticks", OpName: "StockMetaConditions", Required: true, Source: "path"},
	},
}

var StockMetaExchangesOp = Op{
	Name: "StockMetaExchanges", summary: "Get exchange codes",
}

var StockQuoteSingleOp = Op{
	Name: "StockQuoteSingle", summary: "Get historical quotes (single symbol)",
	requiredFlags: []string{"symbol"},
	flags: []FlagDef{
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", OpName: "StockQuoteSingle", Source: "query"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockQuoteSingle", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "StockQuoteSingle", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "source feed of the data.", Completions: []string{"boats", "iex", "otc", "sip"}, OpName: "StockQuoteSingle", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "StockQuoteSingle", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "StockQuoteSingle", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "StockQuoteSingle", Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "StockQuoteSingle", Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol to query", OpName: "StockQuoteSingle", Required: true, Source: "path"},
	},
}

var StockQuotesOp = Op{
	Name: "StockQuotes", summary: "Get historical quotes",
	requiredFlags: []string{"symbols"},
	flags: []FlagDef{
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", OpName: "StockQuotes", Source: "query"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockQuotes", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "StockQuotes", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "source feed of the data.", Completions: []string{"boats", "iex", "otc", "sip"}, OpName: "StockQuotes", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "StockQuotes", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "StockQuotes", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "StockQuotes", Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "StockQuotes", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", OpName: "StockQuotes", Required: true, Source: "query"},
	},
}

var StockSnapshotSingleOp = Op{
	Name: "StockSnapshotSingle", summary: "Get snapshot (single symbol)",
	requiredFlags: []string{"symbol"},
	flags: []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockSnapshotSingle", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, OpName: "StockSnapshotSingle", Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol to query", OpName: "StockSnapshotSingle", Required: true, Source: "path"},
	},
}

var StockSnapshotsOp = Op{
	Name: "StockSnapshots", summary: "Get snapshots",
	requiredFlags: []string{"symbols"},
	flags: []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockSnapshots", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, OpName: "StockSnapshots", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", OpName: "StockSnapshots", Required: true, Source: "query"},
	},
}

var StockTradeSingleOp = Op{
	Name: "StockTradeSingle", summary: "Get historical trades (single symbol)",
	requiredFlags: []string{"symbol"},
	flags: []FlagDef{
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", OpName: "StockTradeSingle", Source: "query"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockTradeSingle", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "StockTradeSingle", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "source feed of the data.", Completions: []string{"boats", "iex", "otc", "sip"}, OpName: "StockTradeSingle", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "StockTradeSingle", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "StockTradeSingle", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "StockTradeSingle", Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "StockTradeSingle", Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol to query", OpName: "StockTradeSingle", Required: true, Source: "path"},
	},
}

var StockTradesOp = Op{
	Name: "StockTrades", summary: "Get historical trades",
	requiredFlags: []string{"symbols"},
	flags: []FlagDef{
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", OpName: "StockTrades", Source: "query"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockTrades", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "StockTrades", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "source feed of the data.", Completions: []string{"boats", "iex", "otc", "sip"}, OpName: "StockTrades", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "StockTrades", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "StockTrades", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "StockTrades", Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "StockTrades", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", OpName: "StockTrades", Required: true, Source: "query"},
	},
}

var UsCorporatesOp = Op{
	Name: "UsCorporates", summary: "Get US corporates",
	flags: []FlagDef{
		{Name: "bond-status", OASName: "bond_status", Type: "string", Description: "status of the bond", Completions: []string{"matured", "outstanding", "pre_issuance"}, OpName: "UsCorporates", Source: "query"},
		{Name: "cusips", OASName: "cusips", Type: "string", Description: "A comma-separated list of CUSIPs with a limit of 1000", OpName: "UsCorporates", Source: "query"},
		{Name: "isins", OASName: "isins", Type: "string", Description: "A comma-separated list of ISINs with a limit of 1000", OpName: "UsCorporates", Source: "query"},
		{Name: "tickers", OASName: "tickers", Type: "string", Description: "A comma-separated list of tickers with a limit of 1000", OpName: "UsCorporates", Source: "query"},
	},
}

var UsTreasuriesOp = Op{
	Name: "UsTreasuries", summary: "Get US treasuries",
	flags: []FlagDef{
		{Name: "bond-status", OASName: "bond_status", Type: "string", Description: "status of the bond", Completions: []string{"matured", "outstanding", "pre_issuance"}, OpName: "UsTreasuries", Source: "query"},
		{Name: "cusips", OASName: "cusips", Type: "string", Description: "A comma-separated list of CUSIPs with a limit of 1000", OpName: "UsTreasuries", Source: "query"},
		{Name: "isins", OASName: "isins", Type: "string", Description: "A comma-separated list of ISINs with a limit of 1000", OpName: "UsTreasuries", Source: "query"},
		{Name: "subtype", OASName: "subtype", Type: "string", Description: "subtype of the treasury", Completions: []string{"bill", "bond", "floating", "note", "strips", "tips"}, OpName: "UsTreasuries", Source: "query"},
	},
}

var AddAssetToWatchlistOp = Op{
	Name: "AddAssetToWatchlist", summary: "Add asset to watchlist",
	requiredFlags: []string{"watchlist-id"},
	flags: []FlagDef{
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "the symbol name to add to the watchlist", OpName: "AddAssetToWatchlist", Source: "body"},
		{Name: "watchlist-id", OASName: "watchlist_id", Type: "string", Description: "watchlist id", OpName: "AddAssetToWatchlist", Required: true, Source: "path"},
	},
}

var AddAssetToWatchlistByNameOp = Op{
	Name: "AddAssetToWatchlistByName", summary: "Add asset to watchlist by name",
	requiredFlags: []string{"name"},
	flags: []FlagDef{
		{Name: "name", OASName: "name", Type: "string", Description: "name of the watchlist", OpName: "AddAssetToWatchlistByName", Required: true, Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "the symbol name to add to the watchlist", OpName: "AddAssetToWatchlistByName", Source: "body"},
	},
}

var CreateCryptoPerpTransferForAccountOp = Op{
	Name: "CreateCryptoPerpTransferForAccount", summary: "Request a new withdrawal",
	flags: []FlagDef{
		{Name: "address", OASName: "address", Type: "string", Description: "destination wallet address", OpName: "CreateCryptoPerpTransferForAccount", Source: "body"},
		{Name: "amount", OASName: "amount", Type: "string", Description: "amount, denoted in the specified asset, to be withdrawn from the user’s wallet", OpName: "CreateCryptoPerpTransferForAccount", Source: "body"},
		{Name: "asset", OASName: "asset", Type: "string", Description: "crypto asset symbol, e.g. BTC, ETH, USDT", OpName: "CreateCryptoPerpTransferForAccount", Source: "body"},
	},
}

var CreateCryptoTransferForAccountOp = Op{
	Name: "CreateCryptoTransferForAccount", summary: "Request a new withdrawal",
	flags: []FlagDef{
		{Name: "address", OASName: "address", Type: "string", Description: "destination wallet address", OpName: "CreateCryptoTransferForAccount", Source: "body"},
		{Name: "amount", OASName: "amount", Type: "string", Description: "amount, denoted in the specified asset, to be withdrawn from the user’s wallet", OpName: "CreateCryptoTransferForAccount", Source: "body"},
		{Name: "asset", OASName: "asset", Type: "string", Description: "crypto asset symbol, e.g. BTC, ETH, USDT", OpName: "CreateCryptoTransferForAccount", Source: "body"},
	},
}

var CreateWhitelistedAddressOp = Op{
	Name: "CreateWhitelistedAddress", summary: "Request a new whitelisted address",
	flags: []FlagDef{
		{Name: "address", OASName: "address", Type: "string", Description: "address to be whitelisted", OpName: "CreateWhitelistedAddress", Source: "body"},
		{Name: "asset", OASName: "asset", Type: "string", Description: "symbol of underlying asset for the whitelisted address", OpName: "CreateWhitelistedAddress", Source: "body"},
	},
}

var CreateWhitelistedPerpAddressOp = Op{
	Name: "CreateWhitelistedPerpAddress", summary: "Request a new whitelisted address",
	flags: []FlagDef{
		{Name: "address", OASName: "address", Type: "string", Description: "address to be whitelisted", OpName: "CreateWhitelistedPerpAddress", Source: "body"},
		{Name: "asset", OASName: "asset", Type: "string", Description: "symbol of underlying asset for the whitelisted address", OpName: "CreateWhitelistedPerpAddress", Source: "body"},
	},
}

var DeleteAllOpenPositionsOp = Op{
	Name: "DeleteAllOpenPositions", summary: "Close all positions",
	flags: []FlagDef{
		{Name: "cancel-orders", OASName: "cancel_orders", Type: "bool", Description: "if true is specified, cancel all open orders before liquidating all positions", OpName: "DeleteAllOpenPositions", Source: "query"},
	},
}

var DeleteAllOrdersOp = Op{
	Name: "DeleteAllOrders", summary: "Delete all orders",
}

var DeleteOpenPositionOp = Op{
	Name: "DeleteOpenPosition", summary: "Close a position",
	requiredFlags: []string{"symbol-or-asset-id"},
	flags: []FlagDef{
		{Name: "percentage", OASName: "percentage", Type: "string", Description: "percentage of position to liquidate", OpName: "DeleteOpenPosition", Source: "query"},
		{Name: "qty", OASName: "qty", Type: "string", Description: "the number of shares to liquidate. Can accept up to 9 decimal points. Cannot work with percentage", OpName: "DeleteOpenPosition", Source: "query"},
		{Name: "symbol-or-asset-id", OASName: "symbol_or_asset_id", Type: "string", Description: "symbol or assetId", OpName: "DeleteOpenPosition", Required: true, Source: "path"},
	},
}

var DeleteOrderByOrderIDOp = Op{
	Name: "DeleteOrderByOrderID", summary: "Delete order by ID",
	requiredFlags: []string{"order-id"},
	flags: []FlagDef{
		{Name: "order-id", OASName: "order_id", Type: "string", Description: "order id", OpName: "DeleteOrderByOrderID", Required: true, Source: "path"},
	},
}

var DeleteWatchlistByIDOp = Op{
	Name: "DeleteWatchlistByID", summary: "Delete watchlist by id",
	requiredFlags: []string{"watchlist-id"},
	flags: []FlagDef{
		{Name: "watchlist-id", OASName: "watchlist_id", Type: "string", Description: "watchlist id", OpName: "DeleteWatchlistByID", Required: true, Source: "path"},
	},
}

var DeleteWatchlistByNameOp = Op{
	Name: "DeleteWatchlistByName", summary: "Delete watchlist by name",
	requiredFlags: []string{"name"},
	flags: []FlagDef{
		{Name: "name", OASName: "name", Type: "string", Description: "name of the watchlist", OpName: "DeleteWatchlistByName", Required: true, Source: "query"},
	},
}

var DeleteWhitelistedAddressOp = Op{
	Name: "DeleteWhitelistedAddress", summary: "Delete a whitelisted address",
	requiredFlags: []string{"whitelisted-address-id"},
	flags: []FlagDef{
		{Name: "whitelisted-address-id", OASName: "whitelisted_address_id", Type: "string", Description: "whitelisted address to delete", OpName: "DeleteWhitelistedAddress", Required: true, Source: "path"},
	},
}

var DeleteWhitelistedPerpAddressOp = Op{
	Name: "DeleteWhitelistedPerpAddress", summary: "Delete a whitelisted address",
	requiredFlags: []string{"whitelisted-address-id"},
	flags: []FlagDef{
		{Name: "whitelisted-address-id", OASName: "whitelisted_address_id", Type: "string", Description: "whitelisted address to delete", OpName: "DeleteWhitelistedPerpAddress", Required: true, Source: "path"},
	},
}

var GetOptionContractSymbolOrIDOp = Op{
	Name: "GetOptionContractSymbolOrID", summary: "Get an option contract by ID or symbol",
	requiredFlags: []string{"symbol-or-id"},
	flags: []FlagDef{
		{Name: "symbol-or-id", OASName: "symbol_or_id", Type: "string", Description: "symbol or contract ID", OpName: "GetOptionContractSymbolOrID", Required: true, Source: "path"},
	},
}

var GetOptionsContractsOp = Op{
	Name: "GetOptionsContracts", summary: "Get option contracts",
	flags: []FlagDef{
		{Name: "expiration-date", OASName: "expiration_date", Type: "string", Description: "filter contracts by the exact expiration date (format: YYYY-MM-DD)", OpName: "GetOptionsContracts", Source: "query"},
		{Name: "expiration-date-gte", OASName: "expiration_date_gte", Type: "string", Description: "filter contracts with expiration date greater than or equal to the specified date", OpName: "GetOptionsContracts", Source: "query"},
		{Name: "expiration-date-lte", OASName: "expiration_date_lte", Type: "string", Description: "filter contracts with expiration date less than or equal to the specified date", OpName: "GetOptionsContracts", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Description: "number of contracts to limit per page (default=100, max=10000)", OpName: "GetOptionsContracts", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "used for pagination, this token retrieves the next page of results", OpName: "GetOptionsContracts", Source: "query"},
		{Name: "ppind", OASName: "ppind", Type: "bool", Description: "ppind(Penny Program Indicator) field indicates whether an option contract is eligible for penny price increments,", OpName: "GetOptionsContracts", Source: "query"},
		{Name: "root-symbol", OASName: "root_symbol", Type: "string", Description: "filter contracts by the root symbol", OpName: "GetOptionsContracts", Source: "query"},
		{Name: "show-deliverables", OASName: "show_deliverables", Type: "bool", Description: "include deliverables array in the response", OpName: "GetOptionsContracts", Source: "query"},
		{Name: "status", OASName: "status", Type: "string", Description: "filter contracts by status (active/inactive). By default only active contracts are returned", Completions: []string{"active", "inactive"}, OpName: "GetOptionsContracts", Source: "query"},
		{Name: "strike-price-gte", OASName: "strike_price_gte", Type: "string", Description: "filter contracts with strike price greater than or equal to the specified value", OpName: "GetOptionsContracts", Source: "query"},
		{Name: "strike-price-lte", OASName: "strike_price_lte", Type: "string", Description: "filter contracts with strike price less than or equal to the specified value", OpName: "GetOptionsContracts", Source: "query"},
		{Name: "style", OASName: "style", Type: "string", Description: "filter contracts by the style (american/european)", Completions: []string{"american", "european"}, OpName: "GetOptionsContracts", Source: "query"},
		{Name: "type", OASName: "type", Type: "string", Description: "filter contracts by the type (call/put)", Completions: []string{"call", "put"}, OpName: "GetOptionsContracts", Source: "query"},
		{Name: "underlying-symbols", OASName: "underlying_symbols", Type: "string", Description: "filter contracts by one or more underlying symbols", OpName: "GetOptionsContracts", Source: "query"},
	},
}

var GetV2AssetsOp = Op{
	Name: "GetV2Assets", summary: "Get assets",
	flags: []FlagDef{
		{Name: "asset-class", OASName: "asset_class", Type: "string", Description: "defaults to us_equity", OpName: "GetV2Assets", Source: "query"},
		{Name: "attributes", OASName: "attributes", Type: "string", Description: "comma separated values to query for more than one attribute", OpName: "GetV2Assets", Source: "query"},
		{Name: "exchange", OASName: "exchange", Type: "string", Description: "optional AMEX, ARCA, BATS, NYSE, NASDAQ, NYSEARCA or OTC", OpName: "GetV2Assets", Source: "query"},
		{Name: "status", OASName: "status", Type: "string", Description: "e.g. “active”. By default, all statuses are included", OpName: "GetV2Assets", Source: "query"},
	},
}

var GetV2AssetsSymbolOrAssetIDOp = Op{
	Name: "GetV2AssetsSymbolOrAssetID", summary: "Get an asset by ID or symbol",
	requiredFlags: []string{"symbol-or-asset-id"},
	flags: []FlagDef{
		{Name: "symbol-or-asset-id", OASName: "symbol_or_asset_id", Type: "string", Description: "symbol or assetId. CUSIP is also accepted for US equities", OpName: "GetV2AssetsSymbolOrAssetID", Required: true, Source: "path"},
	},
}

var GetV2CorporateActionsAnnouncementsOp = Op{
	Name: "GetV2CorporateActionsAnnouncements", summary: "Retrieve announcements",
	requiredFlags: []string{"ca-types", "since", "until"},
	flags: []FlagDef{
		{Name: "ca-types", OASName: "ca_types", Type: "string", Description: "A comma-delimited list of Dividend, Merger, Spinoff, or Split", OpName: "GetV2CorporateActionsAnnouncements", Required: true, Source: "query"},
		{Name: "cusip", OASName: "cusip", Type: "string", Description: "CUSIP of the company initiating the announcement", OpName: "GetV2CorporateActionsAnnouncements", Source: "query"},
		{Name: "date-type", OASName: "date_type", Type: "string", Description: "declaration_date, ex_date, record_date, or payable_date", OpName: "GetV2CorporateActionsAnnouncements", Source: "query"},
		{Name: "since", OASName: "since", Type: "string", Description: "start (inclusive) of the date range when searching corporate action announcements", OpName: "GetV2CorporateActionsAnnouncements", Required: true, Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol of the company initiating the announcement", OpName: "GetV2CorporateActionsAnnouncements", Source: "query"},
		{Name: "until", OASName: "until", Type: "string", Description: "end (inclusive) of the date range when searching corporate action announcements", OpName: "GetV2CorporateActionsAnnouncements", Required: true, Source: "query"},
	},
}

var GetV2CorporateActionsAnnouncementsIDOp = Op{
	Name: "GetV2CorporateActionsAnnouncementsID", summary: "Retrieve a specific announcement",
	requiredFlags: []string{"id"},
	flags: []FlagDef{
		{Name: "id", OASName: "id", Type: "string", Description: "corporate announcement’s id", OpName: "GetV2CorporateActionsAnnouncementsID", Required: true, Source: "path"},
	},
}

var GetAccountOp = Op{
	Name: "GetAccount", summary: "Get account",
}

var GetAccountActivitiesOp = Op{
	Name: "GetAccountActivities", summary: "Retrieve account activities",
	flags: []FlagDef{
		{Name: "activity-types", OASName: "activity_types", Type: "string", Description: "A comma-separated list of activity types used to filter the results", OpName: "GetAccountActivities", Source: "query"},
		{Name: "after", OASName: "after", Type: "string", Description: "get activities created after this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported", OpName: "GetAccountActivities", Source: "query"},
		{Name: "category", OASName: "category", Type: "string", Description: "activity category. Cannot be used with \"activity_types\" parameter", Completions: []string{"non_trade_activity", "trade_activity"}, OpName: "GetAccountActivities", Source: "query"},
		{Name: "date", OASName: "date", Type: "string", Description: "filter activities by the activity date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported", OpName: "GetAccountActivities", Source: "query"},
		{Name: "direction", OASName: "direction", Type: "string", Default: "desc", Description: "chronological order of response based on the activity datetime", Completions: []string{"asc", "desc"}, OpName: "GetAccountActivities", Source: "query"},
		{Name: "page-size", OASName: "page_size", Type: "int", Default: "100", Description: "maximum number of entries to return in the response", OpName: "GetAccountActivities", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "token used for pagination. Provide the ID of the last activity from the last page to retrieve the next set of results", OpName: "GetAccountActivities", Source: "query"},
		{Name: "until", OASName: "until", Type: "string", Description: "get activities created before this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported", OpName: "GetAccountActivities", Source: "query"},
	},
}

var GetAccountActivitiesByActivityTypeOp = Op{
	Name: "GetAccountActivitiesByActivityType", summary: "Retrieve account activities of specific type",
	requiredFlags: []string{"activity-type"},
	flags: []FlagDef{
		{Name: "activity-type", OASName: "activity_type", Type: "string", Description: "activity type you want to view entries for. A list of valid activity types can be found at the bottom of this page", OpName: "GetAccountActivitiesByActivityType", Required: true, Source: "path"},
		{Name: "after", OASName: "after", Type: "string", Description: "get activities created after this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported", OpName: "GetAccountActivitiesByActivityType", Source: "query"},
		{Name: "date", OASName: "date", Type: "string", Description: "filter activities by the activity date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported", OpName: "GetAccountActivitiesByActivityType", Source: "query"},
		{Name: "direction", OASName: "direction", Type: "string", Default: "desc", Description: "chronological order of response based on the activity datetime", Completions: []string{"asc", "desc"}, OpName: "GetAccountActivitiesByActivityType", Source: "query"},
		{Name: "page-size", OASName: "page_size", Type: "int", Default: "100", Description: "maximum number of entries to return in the response", OpName: "GetAccountActivitiesByActivityType", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "token used for pagination. Provide the ID of the last activity from the last page to retrieve the next set of results", OpName: "GetAccountActivitiesByActivityType", Source: "query"},
		{Name: "until", OASName: "until", Type: "string", Description: "get activities created before this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported", OpName: "GetAccountActivitiesByActivityType", Source: "query"},
	},
}

var GetAccountConfigOp = Op{
	Name: "GetAccountConfig", summary: "Get account configurations",
}

var GetAccountPortfolioHistoryOp = Op{
	Name: "GetAccountPortfolioHistory", summary: "Get account portfolio history",
	flags: []FlagDef{
		{Name: "cashflow-types", OASName: "cashflow_types", Type: "string", Description: "cashflow activities to include in the report. One of 'ALL', 'NONE', or a comma-separated list of activity types", OpName: "GetAccountPortfolioHistory", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "timestamp the data is returned up to in RFC3339 format (including timezone specification)", OpName: "GetAccountPortfolioHistory", Source: "query"},
		{Name: "extended-hours", OASName: "extended_hours", Type: "string", Description: "**deprecated**: Users are strongly advised to **rely on the intraday_reporting query parameter** for better control\no...", OpName: "GetAccountPortfolioHistory", Source: "query"},
		{Name: "intraday-reporting", OASName: "intraday_reporting", Type: "string", Default: "market_hours", Description: "for intraday resolutions (<1D) this specifies which timestamps to return data points for:\n\nAllowed values are:\n- **ma...", Completions: []string{"continuous", "extended_hours", "market_hours"}, OpName: "GetAccountPortfolioHistory", Source: "query"},
		{Name: "period", OASName: "period", Type: "string", Description: "duration of the data in number + unit format, such as 1D, where unit can be D for day, W for week, M for month and A ...", OpName: "GetAccountPortfolioHistory", Source: "query"},
		{Name: "pnl-reset", OASName: "pnl_reset", Type: "string", Default: "per_day", Description: "pnl_reset defines how we are calculating the baseline values for Profit And Loss (pnl) for queries with timeframe les...", Completions: []string{"no_reset", "per_day"}, OpName: "GetAccountPortfolioHistory", Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "timestamp the data is returned starting from in RFC3339 format (including timezone specification)", OpName: "GetAccountPortfolioHistory", Source: "query"},
		{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "resolution of time window", OpName: "GetAccountPortfolioHistory", Source: "query"},
	},
}

var GetAllOpenPositionsOp = Op{
	Name: "GetAllOpenPositions", summary: "List all open positions",
}

var GetAllOrdersOp = Op{
	Name: "GetAllOrders", summary: "Get all orders",
	flags: []FlagDef{
		{Name: "after", OASName: "after", Type: "string", Description: "response will include only ones submitted after this timestamp (exclusive.)", OpName: "GetAllOrders", Source: "query"},
		{Name: "after-order-id", OASName: "after_order_id", Type: "string", Description: "return orders submitted after the order with this ID (exclusive).\nMutually exclusive with before_order_id", OpName: "GetAllOrders", Source: "query"},
		{Name: "asset-class", OASName: "asset_class", Type: "string", Description: "A comma-separated list of asset classes, the response will include only orders in the specified asset classes", OpName: "GetAllOrders", Source: "query"},
		{Name: "before-order-id", OASName: "before_order_id", Type: "string", Description: "return orders submitted before the order with this ID (exclusive).\nMutually exclusive with after_order_id", OpName: "GetAllOrders", Source: "query"},
		{Name: "direction", OASName: "direction", Type: "string", Description: "chronological order of response based on the submission time. asc or desc. Defaults to desc", Completions: []string{"asc", "desc"}, OpName: "GetAllOrders", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Description: "maximum number of orders in response. Defaults to 50 and max is 500", OpName: "GetAllOrders", Source: "query"},
		{Name: "nested", OASName: "nested", Type: "bool", Description: "if true, the result will roll up multi-leg orders under the legs field of primary order", OpName: "GetAllOrders", Source: "query"},
		{Name: "side", OASName: "side", Type: "string", Description: "filters down to orders that have a matching side field set", OpName: "GetAllOrders", Source: "query"},
		{Name: "status", OASName: "status", Type: "string", Description: "order status to be queried. open, closed or all. Defaults to open", Completions: []string{"all", "closed", "open"}, OpName: "GetAllOrders", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of symbols to filter by (ex", OpName: "GetAllOrders", Source: "query"},
		{Name: "until", OASName: "until", Type: "string", Description: "response will include only ones submitted until this timestamp (exclusive.)", OpName: "GetAllOrders", Source: "query"},
	},
}

var GetCryptoFundingTransferOp = Op{
	Name: "GetCryptoFundingTransfer", summary: "Retrieve a crypto funding transfer",
	requiredFlags: []string{"transfer-id"},
	flags: []FlagDef{
		{Name: "transfer-id", OASName: "transfer_id", Type: "string", Description: "crypto transfer to retrieve", OpName: "GetCryptoFundingTransfer", Required: true, Source: "path"},
	},
}

var GetCryptoPerpAccountLeverageOp = Op{
	Name: "GetCryptoPerpAccountLeverage", summary: "Get account leverage for an asset",
	flags: []FlagDef{
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol of underlying asset", OpName: "GetCryptoPerpAccountLeverage", Source: "query"},
	},
}

var GetCryptoPerpAccountVitalsOp = Op{
	Name: "GetCryptoPerpAccountVitals", summary: "Retrieve account vitals",
}

var GetCryptoPerpFundingTransferOp = Op{
	Name: "GetCryptoPerpFundingTransfer", summary: "Retrieve a crypto funding transfer",
	requiredFlags: []string{"transfer-id"},
	flags: []FlagDef{
		{Name: "transfer-id", OASName: "transfer_id", Type: "string", Description: "crypto transfer to retrieve", OpName: "GetCryptoPerpFundingTransfer", Required: true, Source: "path"},
	},
}

var GetCryptoPerpTransferEstimateOp = Op{
	Name: "GetCryptoPerpTransferEstimate", summary: "Returns the estimated gas fee for a proposed transaction",
	flags: []FlagDef{
		{Name: "amount", OASName: "amount", Type: "string", Description: "amount, denoted in the specified asset, of the proposed transaction", OpName: "GetCryptoPerpTransferEstimate", Source: "query"},
		{Name: "asset", OASName: "asset", Type: "string", Description: "asset for the proposed transaction", OpName: "GetCryptoPerpTransferEstimate", Source: "query"},
		{Name: "from-address", OASName: "from_address", Type: "string", Description: "originating address of the proposed transaction", OpName: "GetCryptoPerpTransferEstimate", Source: "query"},
		{Name: "to-address", OASName: "to_address", Type: "string", Description: "destination address of the proposed transaction", OpName: "GetCryptoPerpTransferEstimate", Source: "query"},
	},
}

var GetCryptoTransferEstimateOp = Op{
	Name: "GetCryptoTransferEstimate", summary: "Returns the estimated gas fee for a proposed transaction",
	flags: []FlagDef{
		{Name: "amount", OASName: "amount", Type: "string", Description: "amount, denoted in the specified asset, of the proposed transaction", OpName: "GetCryptoTransferEstimate", Source: "query"},
		{Name: "asset", OASName: "asset", Type: "string", Description: "asset for the proposed transaction", OpName: "GetCryptoTransferEstimate", Source: "query"},
		{Name: "from-address", OASName: "from_address", Type: "string", Description: "originating address of the proposed transaction", OpName: "GetCryptoTransferEstimate", Source: "query"},
		{Name: "to-address", OASName: "to_address", Type: "string", Description: "destination address of the proposed transaction", OpName: "GetCryptoTransferEstimate", Source: "query"},
	},
}

var GetOpenPositionOp = Op{
	Name: "GetOpenPosition", summary: "Get an open position",
	requiredFlags: []string{"symbol-or-asset-id"},
	flags: []FlagDef{
		{Name: "symbol-or-asset-id", OASName: "symbol_or_asset_id", Type: "string", Description: "symbol or assetId", OpName: "GetOpenPosition", Required: true, Source: "path"},
	},
}

var GetOrderByClientOrderIDOp = Op{
	Name: "GetOrderByClientOrderID", summary: "Get order by client order ID",
	requiredFlags: []string{"client-order-id"},
	flags: []FlagDef{
		{Name: "client-order-id", OASName: "client_order_id", Type: "string", Description: "client-assigned order ID", OpName: "GetOrderByClientOrderID", Required: true, Source: "query"},
	},
}

var GetOrderByOrderIDOp = Op{
	Name: "GetOrderByOrderID", summary: "Get order by ID",
	requiredFlags: []string{"order-id"},
	flags: []FlagDef{
		{Name: "nested", OASName: "nested", Type: "bool", Description: "if true, the result will roll up multi-leg orders under the legs field of primary order", OpName: "GetOrderByOrderID", Source: "query"},
		{Name: "order-id", OASName: "order_id", Type: "string", Description: "order id", OpName: "GetOrderByOrderID", Required: true, Source: "path"},
	},
}

var GetWatchlistByIDOp = Op{
	Name: "GetWatchlistByID", summary: "Get watchlist by ID",
	requiredFlags: []string{"watchlist-id"},
	flags: []FlagDef{
		{Name: "watchlist-id", OASName: "watchlist_id", Type: "string", Description: "watchlist id", OpName: "GetWatchlistByID", Required: true, Source: "path"},
	},
}

var GetWatchlistByNameOp = Op{
	Name: "GetWatchlistByName", summary: "Get watchlist by name",
	requiredFlags: []string{"name"},
	flags: []FlagDef{
		{Name: "name", OASName: "name", Type: "string", Description: "name of the watchlist", OpName: "GetWatchlistByName", Required: true, Source: "query"},
	},
}

var GetWatchlistsOp = Op{
	Name: "GetWatchlists", summary: "Get all watchlists",
}

var ListCryptoFundingTransfersOp = Op{
	Name: "ListCryptoFundingTransfers", summary: "Retrieve crypto funding transfers",
}

var ListCryptoFundingWalletsOp = Op{
	Name: "ListCryptoFundingWallets", summary: "Retrieve crypto funding wallets",
	flags: []FlagDef{
		{Name: "asset", OASName: "asset", Type: "string", Description: "filter by crypto asset symbol, e.g. BTC, ETH, USDT. If specified and no wallet exists, one will be created", OpName: "ListCryptoFundingWallets", Source: "query"},
		{Name: "network", OASName: "network", Type: "string", Description: "optional network identifier", Completions: []string{"ethereum", "solana"}, OpName: "ListCryptoFundingWallets", Source: "query"},
	},
}

var ListCryptoPerpFundingTransfersOp = Op{
	Name: "ListCryptoPerpFundingTransfers", summary: "Retrieve crypto funding transfers",
}

var ListCryptoPerpFundingWalletsOp = Op{
	Name: "ListCryptoPerpFundingWallets", summary: "Retrieve crypto funding wallets",
	flags: []FlagDef{
		{Name: "asset", OASName: "asset", Type: "string", Description: "asset", OpName: "ListCryptoPerpFundingWallets", Source: "query"},
	},
}

var ListWhitelistedAddressOp = Op{
	Name: "ListWhitelistedAddress", summary: "Get an array of whitelisted addresses",
}

var ListWhitelistedPerpAddressOp = Op{
	Name: "ListWhitelistedPerpAddress", summary: "Get an array of whitelisted addresses",
}

var OptionBarsOp = Op{
	Name: "OptionBars", summary: "Get historical bars",
	requiredFlags: []string{"symbols", "timeframe"},
	flags: []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "OptionBars", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "OptionBars", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "OptionBars", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "OptionBars", Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "OptionBars", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100", OpName: "OptionBars", Required: true, Source: "query"},
		{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g", OpName: "OptionBars", Required: true, Source: "query"},
	},
}

var OptionDoNotExerciseOp = Op{
	Name: "OptionDoNotExercise", summary: "Do not exercise an options position",
	requiredFlags: []string{"symbol-or-contract-id"},
	flags: []FlagDef{
		{Name: "symbol-or-contract-id", OASName: "symbol_or_contract_id", Type: "string", Description: "option contract symbol or ID", OpName: "OptionDoNotExercise", Required: true, Source: "path"},
	},
}

var OptionExerciseOp = Op{
	Name: "OptionExercise", summary: "Exercise an options position",
	requiredFlags: []string{"symbol-or-contract-id"},
	flags: []FlagDef{
		{Name: "symbol-or-contract-id", OASName: "symbol_or_contract_id", Type: "string", Description: "option contract symbol or ID", OpName: "OptionExercise", Required: true, Source: "path"},
	},
}

var PatchAccountConfigOp = Op{
	Name: "PatchAccountConfig", summary: "Update account configurations",
	flags: []FlagDef{
		{Name: "disable-overnight-trading", OASName: "disable_overnight_trading", Type: "bool", Description: "if true, overnight trading is disabled", OpName: "PatchAccountConfig", Source: "body"},
		{Name: "dtbp-check", OASName: "dtbp_check", Type: "string", Description: "both, entry, or exit. Controls Day Trading Margin Call (DTMC) checks", Completions: []string{"both", "entry", "exit"}, OpName: "PatchAccountConfig", Source: "body"},
		{Name: "fractional-trading", OASName: "fractional_trading", Type: "bool", Description: "if true, account is able to participate in fractional trading", OpName: "PatchAccountConfig", Source: "body"},
		{Name: "max-margin-multiplier", OASName: "max_margin_multiplier", Type: "string", Description: "can be \"1\", \"2\", or \"4\"", OpName: "PatchAccountConfig", Source: "body"},
		{Name: "max-options-trading-level", OASName: "max_options_trading_level", Type: "int", Description: "desired maximum options trading level. 0=disabled, 1=Covered Call/Cash-Secured Put, 2=Long Call/Put, 3=Spreads/Straddles", Completions: []string{"0", "1", "2", "3"}, OpName: "PatchAccountConfig", Source: "body"},
		{Name: "no-shorting", OASName: "no_shorting", Type: "bool", Description: "if true, account becomes long-only mode", OpName: "PatchAccountConfig", Source: "body"},
		{Name: "pdt-check", OASName: "pdt_check", Type: "string", Description: "both, entry, or exit", OpName: "PatchAccountConfig", Source: "body"},
		{Name: "ptp-no-exception-entry", OASName: "ptp_no_exception_entry", Type: "bool", Description: "if set to true then Alpaca will accept orders for PTP symbols with no exception. Default is false", OpName: "PatchAccountConfig", Source: "body"},
		{Name: "suspend-trade", OASName: "suspend_trade", Type: "bool", Description: "if true, new orders are blocked", OpName: "PatchAccountConfig", Source: "body"},
		{Name: "trade-confirm-email", OASName: "trade_confirm_email", Type: "string", Description: "all or none. If none, emails for order fills are not sent", OpName: "PatchAccountConfig", Source: "body"},
	},
}

var PatchOrderByOrderIDOp = Op{
	Name: "PatchOrderByOrderID", summary: "Replace order by ID",
	requiredFlags: []string{"order-id"},
	flags: []FlagDef{
		{Name: "advanced-instructions", OASName: "advanced_instructions", Type: "string", Description: "advanced instructions for Elite Smart Router: https://docs.alpaca.markets/docs/alpaca-elite-smart-router", OpName: "PatchOrderByOrderID", Source: "body"},
		{Name: "client-order-id", OASName: "client_order_id", Type: "string", Description: "A unique identifier for the new order. Automatically generated if not sent. (<= 128 characters)", OpName: "PatchOrderByOrderID", Source: "body"},
		{Name: "limit-price", OASName: "limit_price", Type: "string", Description: "required if original order's type field was limit or stop_limit.", OpName: "PatchOrderByOrderID", Source: "body"},
		{Name: "order-id", OASName: "order_id", Type: "string", Description: "order id", OpName: "PatchOrderByOrderID", Required: true, Source: "path"},
		{Name: "qty", OASName: "qty", Type: "string", Description: "number of shares to trade.", OpName: "PatchOrderByOrderID", Source: "body"},
		{Name: "stop-price", OASName: "stop_price", Type: "string", Description: "required if original order type is limit or stop_limit", OpName: "PatchOrderByOrderID", Source: "body"},
		{Name: "time-in-force", OASName: "time_in_force", Type: "string", Description: "time-In-Force values supported by Alpaca vary based on the order's security type", Completions: []string{"cls", "day", "fok", "gtc", "ioc", "opg"}, OpName: "PatchOrderByOrderID", Source: "body"},
		{Name: "trail", OASName: "trail", Type: "string", Description: "the new value of the trail_price or trail_percent value (works only for type=“trailing_stop”)", OpName: "PatchOrderByOrderID", Source: "body"},
	},
}

var PostOrderOp = Op{
	Name: "PostOrder", summary: "Create an order",
	flags: []FlagDef{
		{Name: "advanced-instructions", OASName: "advanced_instructions", Type: "string", Description: "advanced instructions for Elite Smart Router: https://docs.alpaca.markets/docs/alpaca-elite-smart-router", OpName: "PostOrder", Source: "body"},
		{Name: "client-order-id", OASName: "client_order_id", Type: "string", Description: "A unique identifier for the order. Automatically generated if not sent. (<= 128 characters)", OpName: "PostOrder", Source: "body"},
		{Name: "extended-hours", OASName: "extended_hours", Type: "bool", Description: "(default) false", OpName: "PostOrder", Source: "body"},
		{Name: "legs", OASName: "legs", Type: "string", Description: "list of order legs (<= 4)", OpName: "PostOrder", Source: "body"},
		{Name: "limit-price", OASName: "limit_price", Type: "string", Description: "required if type is limit or stop_limit.", OpName: "PostOrder", Source: "body"},
		{Name: "notional", OASName: "notional", Type: "string", Description: "dollar amount to trade. Cannot work with qty. Can only work for market order types and day for time in force", OpName: "PostOrder", Source: "body"},
		{Name: "order-class", OASName: "order_class", Type: "string", Description: "order classes supported by Alpaca vary based on the order's security type", Completions: []string{"bracket", "mleg", "oco", "oto", "simple"}, OpName: "PostOrder", Source: "body"},
		{Name: "position-intent", OASName: "position_intent", Type: "string", Description: "represents the desired position strategy", Completions: []string{"buy_to_close", "buy_to_open", "sell_to_close", "sell_to_open"}, OpName: "PostOrder", Source: "body"},
		{Name: "qty", OASName: "qty", Type: "string", Description: "number of shares to trade", OpName: "PostOrder", Source: "body"},
		{Name: "side", OASName: "side", Type: "string", Description: "represents which side this order was on:\n- buy\n- sell\nRequired for all order classes except for mleg", Completions: []string{"buy", "sell"}, OpName: "PostOrder", Source: "body"},
		{Name: "stop-loss", OASName: "stop_loss", Type: "string", Description: "takes in string/number values for stop_price and limit_price", OpName: "PostOrder", Source: "body"},
		{Name: "stop-price", OASName: "stop_price", Type: "string", Description: "required if type is stop or stop_limit", OpName: "PostOrder", Source: "body"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol, asset ID, or currency pair to identify the asset to trade, required for all order classes except for mleg", OpName: "PostOrder", Source: "body"},
		{Name: "take-profit", OASName: "take_profit", Type: "string", Description: "takes in a string/number value for limit_price", OpName: "PostOrder", Source: "body"},
		{Name: "time-in-force", OASName: "time_in_force", Type: "string", Description: "time-In-Force values supported by Alpaca vary based on the order's security type", Completions: []string{"cls", "day", "fok", "gtc", "ioc", "opg"}, OpName: "PostOrder", Source: "body"},
		{Name: "trail-percent", OASName: "trail_percent", Type: "string", Description: "this or trail_price is required if type is trailing_stop", OpName: "PostOrder", Source: "body"},
		{Name: "trail-price", OASName: "trail_price", Type: "string", Description: "this or trail_percent is required if type is trailing_stop", OpName: "PostOrder", Source: "body"},
		{Name: "type", OASName: "type", Type: "string", Description: "order types supported by Alpaca vary based on the order's security type", Completions: []string{"limit", "market", "stop", "stop_limit", "trailing_stop"}, OpName: "PostOrder", Source: "body"},
	},
}

var PostWatchlistOp = Op{
	Name: "PostWatchlist", summary: "Create watchlist",
	flags: []FlagDef{
		{Name: "name", OASName: "name", Type: "string", Description: "watchlist name", OpName: "PostWatchlist", Source: "body"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "list of asset symbols to include in the watchlist", OpName: "PostWatchlist", Source: "body"},
	},
}

var RemoveAssetFromWatchlistOp = Op{
	Name: "RemoveAssetFromWatchlist", summary: "Delete symbol from watchlist",
	requiredFlags: []string{"symbol", "watchlist-id"},
	flags: []FlagDef{
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol name to remove from the watchlist content", OpName: "RemoveAssetFromWatchlist", Required: true, Source: "path"},
		{Name: "watchlist-id", OASName: "watchlist_id", Type: "string", Description: "watchlist ID", OpName: "RemoveAssetFromWatchlist", Required: true, Source: "path"},
	},
}

var SetCryptoPerpAccountLeverageOp = Op{
	Name: "SetCryptoPerpAccountLeverage", summary: "Set account leverage for an asset",
	flags: []FlagDef{
		{Name: "leverage", OASName: "leverage", Type: "int", Description: "leverage for the underlying asset", OpName: "SetCryptoPerpAccountLeverage", Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol of underlying asset", OpName: "SetCryptoPerpAccountLeverage", Source: "query"},
	},
}

var UpdateWatchlistByIDOp = Op{
	Name: "UpdateWatchlistByID", summary: "Update watchlist by id",
	requiredFlags: []string{"watchlist-id"},
	flags: []FlagDef{
		{Name: "name", OASName: "name", Type: "string", Description: "watchlist name", OpName: "UpdateWatchlistByID", Source: "body"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "list of asset symbols to include in the watchlist", OpName: "UpdateWatchlistByID", Source: "body"},
		{Name: "watchlist-id", OASName: "watchlist_id", Type: "string", Description: "watchlist id", OpName: "UpdateWatchlistByID", Required: true, Source: "path"},
	},
}

var UpdateWatchlistByNameOp = Op{
	Name: "UpdateWatchlistByName", summary: "Update watchlist by name",
	requiredFlags: []string{"name"},
	flags: []FlagDef{
		{Name: "name", OASName: "name", Type: "string", Description: "name of the watchlist", OpName: "UpdateWatchlistByName", Required: true, Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "list of asset symbols to include in the watchlist", OpName: "UpdateWatchlistByName", Source: "body"},
	},
}

// ResponseField describes a field in an API response.
type ResponseField struct {
	Name        string
	Type        string
	Description string
	EnumValues  []string
}

// ResponseSchemas maps operation names to their response fields.
var ResponseSchemas = map[string][]ResponseField{
	"Calendar": {
		{Name: "calendar", Type: "[]object", Description: "market calendar"},
		{Name: "market", Type: "object", Description: "A market"},
	},
	"Clock": {
		{Name: "clocks", Type: "[]object", Description: "clocks"},
	},
	"CorporateActions": {
		{Name: "corporate_actions", Type: "object", Description: "corporate actions"},
		{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
	},
	"CryptoBars": {
		{Name: "bars", Type: "map[string][]object", Description: "bars"},
		{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
	},
	"CryptoLatestBars": {
		{Name: "bars", Type: "map[string]object", Description: "bars"},
	},
	"CryptoLatestOrderbooks": {
		{Name: "orderbooks", Type: "map[string]object", Description: "orderbooks"},
	},
	"CryptoLatestQuotes": {
		{Name: "quotes", Type: "map[string]object", Description: "quotes"},
	},
	"CryptoLatestTrades": {
		{Name: "trades", Type: "map[string]object", Description: "trades"},
	},
	"CryptoPerpLatestBars": {
		{Name: "bars", Type: "map[string]object", Description: "bars"},
	},
	"CryptoPerpLatestFuturesPricing": {
		{Name: "pricing", Type: "map[string]object", Description: "pricing"},
	},
	"CryptoPerpLatestOrderbooks": {
		{Name: "orderbooks", Type: "map[string]object", Description: "orderbooks"},
	},
	"CryptoPerpLatestQuotes": {
		{Name: "quotes", Type: "map[string]object", Description: "quotes"},
	},
	"CryptoPerpLatestTrades": {
		{Name: "trades", Type: "map[string]object", Description: "trades"},
	},
	"CryptoQuotes": {
		{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
		{Name: "quotes", Type: "map[string][]object", Description: "quotes"},
	},
	"CryptoSnapshots": {
		{Name: "snapshots", Type: "map[string]object", Description: "snapshots"},
	},
	"CryptoTrades": {
		{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
		{Name: "trades", Type: "map[string][]object", Description: "trades"},
	},
	"FixedIncomeLatestPrices": {
		{Name: "prices", Type: "map[string]object", Description: "prices"},
	},
	"LatestRates": {
		{Name: "rates", Type: "map[string]object", Description: "rates"},
	},
	"LegacyClock": {
		{Name: "is_open", Type: "boolean", Description: "whether or not the market is open"},
		{Name: "next_close", Type: "string", Description: "next market close timestamp"},
		{Name: "next_open", Type: "string", Description: "next market open timestamp"},
		{Name: "timestamp", Type: "string", Description: "current timestamp"},
	},
	"MostActives": {
		{Name: "last_updated", Type: "string", Description: "time when the most actives were last computed. Formatted as a RFC-3339 date-time with nanosecond precision"},
		{Name: "most_actives", Type: "[]object", Description: "list of top N most active symbols"},
	},
	"Movers": {
		{Name: "gainers", Type: "[]object", Description: "list of top N gainers"},
		{Name: "last_updated", Type: "string", Description: "time when the movers were last computed. Formatted as a RFC-3339 date-time with nanosecond precision"},
		{Name: "losers", Type: "[]object", Description: "list of top N losers"},
		{Name: "market_type", Type: "enum", Description: "market type (stocks or crypto)", EnumValues: []string{"crypto", "stocks"}},
	},
	"News": {
		{Name: "news", Type: "[]object", Description: "news"},
		{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
	},
	"OptionChain": {
		{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
		{Name: "snapshots", Type: "map[string]object", Description: "snapshots"},
	},
	"OptionLatestQuotes": {
		{Name: "quotes", Type: "map[string]object", Description: "quotes"},
	},
	"OptionLatestTrades": {
		{Name: "trades", Type: "map[string]object", Description: "trades"},
	},
	"OptionSnapshots": {
		{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
		{Name: "snapshots", Type: "map[string]object", Description: "snapshots"},
	},
	"OptionTrades": {
		{Name: "currency", Type: "string", Description: "currency"},
		{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
		{Name: "trades", Type: "map[string][]object", Description: "trades"},
	},
	"Rates": {
		{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
		{Name: "rates", Type: "map[string][]object", Description: "rates"},
	},
	"StockAuctionSingle": {
		{Name: "auctions", Type: "[]object", Description: "auctions"},
		{Name: "currency", Type: "string", Description: "currency"},
		{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
		{Name: "symbol", Type: "string", Description: "symbol"},
	},
	"StockAuctions": {
		{Name: "auctions", Type: "map[string][]object", Description: "auctions"},
		{Name: "currency", Type: "string", Description: "currency"},
		{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
	},
	"StockBarSingle": {
		{Name: "bars", Type: "[]object", Description: "bars"},
		{Name: "currency", Type: "string", Description: "currency"},
		{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
		{Name: "symbol", Type: "string", Description: "symbol"},
	},
	"StockBars": {
		{Name: "bars", Type: "map[string][]object", Description: "bars"},
		{Name: "currency", Type: "string", Description: "currency"},
		{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
	},
	"StockLatestBarSingle": {
		{Name: "bar", Type: "object", Description: "OHLC aggregate of all the trades in a given interval"},
		{Name: "currency", Type: "string", Description: "currency"},
		{Name: "symbol", Type: "string", Description: "symbol"},
	},
	"StockLatestBars": {
		{Name: "bars", Type: "map[string]object", Description: "bars"},
		{Name: "currency", Type: "string", Description: "currency"},
	},
	"StockLatestQuoteSingle": {
		{Name: "currency", Type: "string", Description: "currency"},
		{Name: "quote", Type: "object", Description: "best bid and ask information for a given security"},
		{Name: "symbol", Type: "string", Description: "symbol"},
	},
	"StockLatestQuotes": {
		{Name: "currency", Type: "string", Description: "currency"},
		{Name: "quotes", Type: "map[string]object", Description: "quotes"},
	},
	"StockLatestTradeSingle": {
		{Name: "currency", Type: "string", Description: "currency"},
		{Name: "symbol", Type: "string", Description: "symbol"},
		{Name: "trade", Type: "object", Description: "A stock trade"},
	},
	"StockLatestTrades": {
		{Name: "currency", Type: "string", Description: "currency"},
		{Name: "trades", Type: "map[string]object", Description: "trades"},
	},
	"StockQuoteSingle": {
		{Name: "currency", Type: "string", Description: "currency"},
		{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
		{Name: "quotes", Type: "[]object", Description: "quotes"},
		{Name: "symbol", Type: "string", Description: "symbol"},
	},
	"StockQuotes": {
		{Name: "currency", Type: "string", Description: "currency"},
		{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
		{Name: "quotes", Type: "map[string][]object", Description: "quotes"},
	},
	"StockTradeSingle": {
		{Name: "currency", Type: "string", Description: "currency"},
		{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
		{Name: "symbol", Type: "string", Description: "symbol"},
		{Name: "trades", Type: "[]object", Description: "trades"},
	},
	"StockTrades": {
		{Name: "currency", Type: "string", Description: "currency"},
		{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
		{Name: "trades", Type: "map[string][]object", Description: "trades"},
	},
	"UsCorporates": {
		{Name: "us_corporates", Type: "[]object", Description: "us corporates"},
	},
	"UsTreasuries": {
		{Name: "us_treasuries", Type: "[]object", Description: "us treasuries"},
	},
	"AddAssetToWatchlist": {
		{Name: "account_id", Type: "string", Description: "account ID"},
		{Name: "assets", Type: "[]object", Description: "the content of this watchlist, in the order as registered by the client"},
		{Name: "created_at", Type: "string", Description: "created at"},
		{Name: "id", Type: "string", Description: "watchlist id"},
		{Name: "name", Type: "string", Description: "user-defined watchlist name (up to 64 characters)"},
		{Name: "updated_at", Type: "string", Description: "updated at"},
	},
	"AddAssetToWatchlistByName": {
		{Name: "account_id", Type: "string", Description: "account ID"},
		{Name: "assets", Type: "[]object", Description: "the content of this watchlist, in the order as registered by the client"},
		{Name: "created_at", Type: "string", Description: "created at"},
		{Name: "id", Type: "string", Description: "watchlist id"},
		{Name: "name", Type: "string", Description: "user-defined watchlist name (up to 64 characters)"},
		{Name: "updated_at", Type: "string", Description: "updated at"},
	},
	"CreateCryptoPerpTransferForAccount": {
		{Name: "amount", Type: "string", Description: "amount of transfer denominated in the underlying crypto asset"},
		{Name: "asset", Type: "string", Description: "symbol of crypto asset for given transfer (e.g. BTC)"},
		{Name: "chain", Type: "string", Description: "underlying network for given transfer"},
		{Name: "created_at", Type: "string", Description: "timestamp when transfer was created"},
		{Name: "direction", Type: "enum", Description: "direction", EnumValues: []string{"INCOMING", "OUTGOING"}},
		{Name: "fees", Type: "string", Description: "fees"},
		{Name: "from_address", Type: "string", Description: "originating address of the transfer"},
		{Name: "id", Type: "string", Description: "crypto transfer ID"},
		{Name: "network_fee", Type: "string", Description: "network fee"},
		{Name: "status", Type: "enum", Description: "status", EnumValues: []string{"COMPLETE", "FAILED", "PROCESSING"}},
		{Name: "to_address", Type: "string", Description: "destination address of the transfer"},
		{Name: "tx_hash", Type: "string", Description: "on-chain transaction hash (e.g. 0xabc...xyz)"},
		{Name: "usd_value", Type: "string", Description: "equivalent USD value at time of transfer"},
	},
	"CreateCryptoTransferForAccount": {
		{Name: "amount", Type: "string", Description: "amount of transfer denominated in the underlying crypto asset"},
		{Name: "asset", Type: "string", Description: "symbol of crypto asset for given transfer (e.g. BTC)"},
		{Name: "chain", Type: "string", Description: "underlying network for given transfer"},
		{Name: "created_at", Type: "string", Description: "timestamp when transfer was created"},
		{Name: "direction", Type: "enum", Description: "direction", EnumValues: []string{"INCOMING", "OUTGOING"}},
		{Name: "fees", Type: "string", Description: "fees"},
		{Name: "from_address", Type: "string", Description: "originating address of the transfer"},
		{Name: "id", Type: "string", Description: "crypto transfer ID"},
		{Name: "network_fee", Type: "string", Description: "network fee"},
		{Name: "status", Type: "enum", Description: "status", EnumValues: []string{"COMPLETE", "FAILED", "PROCESSING"}},
		{Name: "to_address", Type: "string", Description: "destination address of the transfer"},
		{Name: "tx_hash", Type: "string", Description: "on-chain transaction hash (e.g. 0xabc...xyz)"},
		{Name: "usd_value", Type: "string", Description: "equivalent USD value at time of transfer"},
	},
	"CreateWhitelistedAddress": {
		{Name: "address", Type: "string", Description: "whitelisted address"},
		{Name: "asset", Type: "string", Description: "symbol of underlying asset for the whitelisted address"},
		{Name: "chain", Type: "string", Description: "underlying network this address represents"},
		{Name: "created_at", Type: "string", Description: "timestamp (RFC3339) of account creation"},
		{Name: "id", Type: "string", Description: "unique ID for whitelisted address"},
		{Name: "status", Type: "enum", Description: "status of whitelisted address which is either ACTIVE or PENDING", EnumValues: []string{"APPROVED", "PENDING"}},
	},
	"CreateWhitelistedPerpAddress": {
		{Name: "address", Type: "string", Description: "whitelisted address"},
		{Name: "asset", Type: "string", Description: "symbol of underlying asset for the whitelisted address"},
		{Name: "chain", Type: "string", Description: "underlying network this address represents"},
		{Name: "created_at", Type: "string", Description: "timestamp (RFC3339) of account creation"},
		{Name: "id", Type: "string", Description: "unique ID for whitelisted address"},
		{Name: "status", Type: "enum", Description: "status of whitelisted address which is either ACTIVE or PENDING", EnumValues: []string{"APPROVED", "PENDING"}},
	},
	"DeleteAllOpenPositions": {
		{Name: "body", Type: "object", Description: "orders API allows a user to monitor, place and cancel their orders with Alpaca.\n\nEach order has a unique identifier p..."},
		{Name: "status", Type: "integer", Description: "HTTP status code for the attempt to close this position"},
		{Name: "symbol", Type: "string", Description: "symbol name of the asset"},
	},
	"DeleteAllOrders": {
		{Name: "id", Type: "string", Description: "orderId"},
		{Name: "status", Type: "integer", Description: "http response code"},
	},
	"DeleteOpenPosition": {
		{Name: "asset_class", Type: "enum", Description: "this represents the category to which the asset belongs to", EnumValues: []string{"crypto", "us_equity", "us_option"}},
		{Name: "asset_id", Type: "string", Description: "asset ID (For options this represents the option contract ID)"},
		{Name: "canceled_at", Type: "string", Description: "canceled at"},
		{Name: "client_order_id", Type: "string", Description: "client unique order ID"},
		{Name: "created_at", Type: "string", Description: "created at"},
		{Name: "expired_at", Type: "string", Description: "expired at"},
		{Name: "extended_hours", Type: "boolean", Description: "if true, eligible for execution outside regular trading hours"},
		{Name: "failed_at", Type: "string", Description: "failed at"},
		{Name: "filled_at", Type: "string", Description: "filled at"},
		{Name: "filled_avg_price", Type: "string", Description: "filled average price"},
		{Name: "filled_qty", Type: "string", Description: "filled quantity"},
		{Name: "hwm", Type: "string", Description: "highest (lowest) market price seen since the trailing stop order was submitted"},
		{Name: "id", Type: "string", Description: "order ID"},
		{Name: "legs", Type: "[]object", Description: "when querying non-simple order_class orders in a nested style, an array of Order entities associated with this order"},
		{Name: "limit_price", Type: "string", Description: "limit price"},
		{Name: "notional", Type: "string", Description: "ordered notional amount. If entered, qty will be null. Can take up to 9 decimal points"},
		{Name: "order_class", Type: "enum", Description: "order classes supported by Alpaca vary based on the order's security type", EnumValues: []string{"bracket", "mleg", "oco", "oto", "simple"}},
		{Name: "order_type", Type: "string", Description: "deprecated in favour of the field \"type\""},
		{Name: "position_intent", Type: "enum", Description: "represents the desired position strategy", EnumValues: []string{"buy_to_close", "buy_to_open", "sell_to_close", "sell_to_open"}},
		{Name: "qty", Type: "string", Description: "ordered quantity. If entered, notional will be null. Can take up to 9 decimal points. Required if order class is mleg"},
		{Name: "replaced_at", Type: "string", Description: "replaced at"},
		{Name: "replaced_by", Type: "string", Description: "order ID that this order was replaced by"},
		{Name: "replaces", Type: "string", Description: "order ID that this order replaces"},
		{Name: "side", Type: "enum", Description: "represents which side this order was on:\n- buy\n- sell\nRequired for all order classes except for mleg", EnumValues: []string{"buy", "sell"}},
		{Name: "status", Type: "enum", Description: "an order executed through Alpaca can experience several status changes during its lifecycle", EnumValues: []string{"accepted", "accepted_for_bidding", "calculated", "canceled", "done_for_day", "expired", "filled", "new", "partially_filled", "pending_cancel", "pending_new", "pending_replace", "rejected", "replaced", "stopped", "suspended"}},
		{Name: "stop_price", Type: "string", Description: "stop price"},
		{Name: "submitted_at", Type: "string", Description: "submitted at"},
		{Name: "symbol", Type: "string", Description: "asset symbol, required for all order classes except for mleg"},
		{Name: "time_in_force", Type: "enum", Description: "time-In-Force values supported by Alpaca vary based on the order's security type", EnumValues: []string{"cls", "day", "fok", "gtc", "ioc", "opg"}},
		{Name: "trail_percent", Type: "string", Description: "percent value away from the high water mark for trailing stop orders"},
		{Name: "trail_price", Type: "string", Description: "dollar value away from the high water mark for trailing stop orders"},
		{Name: "type", Type: "enum", Description: "order types supported by Alpaca vary based on the order's security type", EnumValues: []string{"limit", "market", "stop", "stop_limit", "trailing_stop"}},
		{Name: "updated_at", Type: "string", Description: "updated at"},
	},
	"GetOptionContractSymbolOrID": {
		{Name: "close_price", Type: "string", Description: "close price of the option contract"},
		{Name: "close_price_date", Type: "string", Description: "date of the close price data"},
		{Name: "deliverables", Type: "[]object", Description: "represents the deliverables tied to the option contract"},
		{Name: "expiration_date", Type: "string", Description: "expiration date of the option contract"},
		{Name: "id", Type: "string", Description: "unique identifier of the option contract"},
		{Name: "multiplier", Type: "string", Description: "multiplier of the option contract is crucial for calculating both the trade premium and the extended strike price"},
		{Name: "name", Type: "string", Description: "name of the option contract"},
		{Name: "open_interest", Type: "string", Description: "open interest of the option contract"},
		{Name: "open_interest_date", Type: "string", Description: "date of the open interest data"},
		{Name: "root_symbol", Type: "string", Description: "root symbol of the option contract"},
		{Name: "size", Type: "string", Description: "represents the number of underlying shares to be delivered in case the contract is exercised/assigned"},
		{Name: "status", Type: "enum", Description: "status of the option contract", EnumValues: []string{"active", "inactive"}},
		{Name: "strike_price", Type: "string", Description: "strike price of the option contract"},
		{Name: "style", Type: "enum", Description: "style of the option contract", EnumValues: []string{"american", "european"}},
		{Name: "symbol", Type: "string", Description: "symbol representing the option contract"},
		{Name: "tradable", Type: "boolean", Description: "indicates whether the option contract is tradable"},
		{Name: "type", Type: "enum", Description: "type of the option contract", EnumValues: []string{"call", "put"}},
		{Name: "underlying_asset_id", Type: "string", Description: "unique identifier of the underlying asset"},
		{Name: "underlying_symbol", Type: "string", Description: "underlying symbol of the option contract"},
	},
	"GetV2Assets": {
		{Name: "attributes", Type: "[]enum", Description: "unique characteristics of the asset", EnumValues: []string{"fractional_eh_enabled", "has_options", "ipo", "options_late_close", "overnight_halted", "overnight_tradable", "ptp_no_exception", "ptp_with_exception"}},
		{Name: "class", Type: "enum", Description: "this represents the category to which the asset belongs to", EnumValues: []string{"crypto", "us_equity", "us_option"}},
		{Name: "cusip", Type: "string", Description: "CUSIP identifier for the asset (US Equities only).\nTo request a specific CUSIP, please reach out to Alpaca support"},
		{Name: "easy_to_borrow", Type: "boolean", Description: "asset is easy-to-borrow or not (filtering for easy_to_borrow = True is the best way to check whether the name is curr..."},
		{Name: "exchange", Type: "enum", Description: "represents the current exchanges Alpaca supports", EnumValues: []string{"AMEX", "ARCA", "BATS", "NASDAQ", "NYSE", "NYSEARCA", "OTC"}},
		{Name: "fractionable", Type: "boolean", Description: "asset is fractionable or not"},
		{Name: "id", Type: "string", Description: "asset ID"},
		{Name: "maintenance_margin_requirement", Type: "number", Description: "**deprecated**: Please use margin_requirement_long or margin_requirement_short instead"},
		{Name: "margin_requirement_long", Type: "string", Description: "margin requirement percentage for the asset's long positions (equities only)"},
		{Name: "margin_requirement_short", Type: "string", Description: "margin requirement percentage for the asset's short positions (equities only)"},
		{Name: "marginable", Type: "boolean", Description: "asset is marginable or not"},
		{Name: "name", Type: "string", Description: "official name of the asset"},
		{Name: "shortable", Type: "boolean", Description: "asset is shortable or not"},
		{Name: "status", Type: "enum", Description: "active or inactive", EnumValues: []string{"active", "inactive"}},
		{Name: "symbol", Type: "string", Description: "symbol of the asset"},
		{Name: "tradable", Type: "boolean", Description: "asset is tradable on Alpaca or not"},
	},
	"GetV2AssetsSymbolOrAssetID": {
		{Name: "attributes", Type: "[]enum", Description: "unique characteristics of the asset", EnumValues: []string{"fractional_eh_enabled", "has_options", "ipo", "options_late_close", "overnight_halted", "overnight_tradable", "ptp_no_exception", "ptp_with_exception"}},
		{Name: "class", Type: "enum", Description: "this represents the category to which the asset belongs to", EnumValues: []string{"crypto", "us_equity", "us_option"}},
		{Name: "cusip", Type: "string", Description: "CUSIP identifier for the asset (US Equities only).\nTo request a specific CUSIP, please reach out to Alpaca support"},
		{Name: "easy_to_borrow", Type: "boolean", Description: "asset is easy-to-borrow or not (filtering for easy_to_borrow = True is the best way to check whether the name is curr..."},
		{Name: "exchange", Type: "enum", Description: "represents the current exchanges Alpaca supports", EnumValues: []string{"AMEX", "ARCA", "BATS", "NASDAQ", "NYSE", "NYSEARCA", "OTC"}},
		{Name: "fractionable", Type: "boolean", Description: "asset is fractionable or not"},
		{Name: "id", Type: "string", Description: "asset ID"},
		{Name: "maintenance_margin_requirement", Type: "number", Description: "**deprecated**: Please use margin_requirement_long or margin_requirement_short instead"},
		{Name: "margin_requirement_long", Type: "string", Description: "margin requirement percentage for the asset's long positions (equities only)"},
		{Name: "margin_requirement_short", Type: "string", Description: "margin requirement percentage for the asset's short positions (equities only)"},
		{Name: "marginable", Type: "boolean", Description: "asset is marginable or not"},
		{Name: "name", Type: "string", Description: "official name of the asset"},
		{Name: "shortable", Type: "boolean", Description: "asset is shortable or not"},
		{Name: "status", Type: "enum", Description: "active or inactive", EnumValues: []string{"active", "inactive"}},
		{Name: "symbol", Type: "string", Description: "symbol of the asset"},
		{Name: "tradable", Type: "boolean", Description: "asset is tradable on Alpaca or not"},
	},
	"GetAccount": {
		{Name: "account_blocked", Type: "boolean", Description: "if true, the account activity by user is prohibited"},
		{Name: "account_number", Type: "string", Description: "account number"},
		{Name: "accrued_fees", Type: "string", Description: "fees collected"},
		{Name: "balance_asof", Type: "string", Description: "date of the snapshot for last_* fields"},
		{Name: "buying_power", Type: "string", Description: "current available $ buying power; If multiplier = 4, this is your daytrade buying power which is calculated as (last_..."},
		{Name: "cash", Type: "string", Description: "cash Balance"},
		{Name: "created_at", Type: "string", Description: "timestamp this account was created at"},
		{Name: "currency", Type: "string", Description: "USD"},
		{Name: "daytrade_count", Type: "integer", Description: "current number of daytrades that have been made in the last 5 trading days (inclusive of today)"},
		{Name: "daytrading_buying_power", Type: "string", Description: "your buying power for day trades (continuously updated value)"},
		{Name: "equity", Type: "string", Description: "cash + long_market_value + short_market_value"},
		{Name: "id", Type: "string", Description: "account Id"},
		{Name: "initial_margin", Type: "string", Description: "reg T initial margin requirement (continuously updated value)"},
		{Name: "intraday_adjustments", Type: "string", Description: "intraday adjustment by non_trade_activities such as fund deposit/withdraw"},
		{Name: "last_equity", Type: "string", Description: "equity as of previous trading day at 16:00:00 ET"},
		{Name: "last_maintenance_margin", Type: "string", Description: "your maintenance margin requirement on the previous trading day"},
		{Name: "long_market_value", Type: "string", Description: "real-time MtM value of all long positions held in the account"},
		{Name: "maintenance_margin", Type: "string", Description: "maintenance margin requirement (continuously updated value)"},
		{Name: "multiplier", Type: "string", Description: "buying power multiplier that represents account margin classification; valid values 1 (standard limited margin accoun..."},
		{Name: "non_marginable_buying_power", Type: "string", Description: "current available non-margin dollar buying power"},
		{Name: "options_approved_level", Type: "integer", Description: "options trading level that was approved for this account.", EnumValues: []string{"0", "1", "2", "3"}},
		{Name: "options_buying_power", Type: "string", Description: "your buying power for options trading"},
		{Name: "options_trading_level", Type: "integer", Description: "effective options trading level of the account.", EnumValues: []string{"0", "1", "2", "3"}},
		{Name: "pattern_day_trader", Type: "boolean", Description: "whether or not the account has been flagged as a pattern day trader"},
		{Name: "pending_reg_taf_fees", Type: "string", Description: "pending regulatory fees for the account"},
		{Name: "pending_transfer_in", Type: "string", Description: "cash pending transfer in"},
		{Name: "pending_transfer_out", Type: "string", Description: "cash pending transfer out"},
		{Name: "portfolio_value", Type: "string", Description: "total value of cash + holding positions (This field is deprecated. It is equivalent to the equity field.)"},
		{Name: "regt_buying_power", Type: "string", Description: "your buying power under Regulation T (your excess equity - equity minus margin value - times your margin multiplier)"},
		{Name: "short_market_value", Type: "string", Description: "real-time MtM value of all short positions held in the account"},
		{Name: "shorting_enabled", Type: "boolean", Description: "flag to denote whether or not the account is permitted to short"},
		{Name: "sma", Type: "string", Description: "value of special memorandum account (will be used at a later date to provide additional buying_power)"},
		{Name: "status", Type: "enum", Description: "an enum representing the various possible account status values.\n\nMost likely, the account status is ACTIVE unless th...", EnumValues: []string{"ACCOUNT_UPDATED", "ACTIVE", "APPROVAL_PENDING", "ONBOARDING", "REJECTED", "SUBMISSION_FAILED", "SUBMITTED"}},
		{Name: "trade_suspended_by_user", Type: "boolean", Description: "user setting. If true, the account is not allowed to place orders"},
		{Name: "trading_blocked", Type: "boolean", Description: "if true, the account is not allowed to place orders"},
		{Name: "transfers_blocked", Type: "boolean", Description: "if true, the account is not allowed to request money transfers"},
	},
	"GetAccountConfig": {
		{Name: "disable_overnight_trading", Type: "boolean", Description: "if true, overnight trading is disabled"},
		{Name: "dtbp_check", Type: "enum", Description: "both, entry, or exit. Controls Day Trading Margin Call (DTMC) checks", EnumValues: []string{"both", "entry", "exit"}},
		{Name: "fractional_trading", Type: "boolean", Description: "if true, account is able to participate in fractional trading"},
		{Name: "max_margin_multiplier", Type: "string", Description: "can be \"1\", \"2\", or \"4\""},
		{Name: "max_options_trading_level", Type: "integer", Description: "desired maximum options trading level. 0=disabled, 1=Covered Call/Cash-Secured Put, 2=Long Call/Put, 3=Spreads/Straddles", EnumValues: []string{"0", "1", "2", "3"}},
		{Name: "no_shorting", Type: "boolean", Description: "if true, account becomes long-only mode"},
		{Name: "pdt_check", Type: "string", Description: "both, entry, or exit"},
		{Name: "ptp_no_exception_entry", Type: "boolean", Description: "if set to true then Alpaca will accept orders for PTP symbols with no exception. Default is false"},
		{Name: "suspend_trade", Type: "boolean", Description: "if true, new orders are blocked"},
		{Name: "trade_confirm_email", Type: "string", Description: "all or none. If none, emails for order fills are not sent"},
	},
	"GetAccountPortfolioHistory": {
		{Name: "base_value", Type: "number", Description: "basis in dollar of the profit loss calculation"},
		{Name: "base_value_asof", Type: "string", Description: "if included, then it indicates that the base_value is the account's closing\nequity value at this trading date.\n\nIf no..."},
		{Name: "cashflow", Type: "object", Description: "accumulated value in dollar amount as of the end of each time window"},
		{Name: "equity", Type: "[]number", Description: "equity value of the account in dollar amount as of the end of each time window"},
		{Name: "profit_loss", Type: "[]number", Description: "profit/loss in dollar from the base value"},
		{Name: "profit_loss_pct", Type: "[]number", Description: "profit/loss in percentage from the base value"},
		{Name: "timeframe", Type: "string", Description: "time window size of each data element"},
		{Name: "timestamp", Type: "[]integer", Description: "time of each data element, left-labeled (the beginning of time window)."},
	},
	"GetAllOpenPositions": {
		{Name: "asset_class", Type: "enum", Description: "this represents the category to which the asset belongs to", EnumValues: []string{"crypto", "us_equity", "us_option"}},
		{Name: "asset_id", Type: "string", Description: "asset ID (For options this represents the option contract ID)"},
		{Name: "asset_marginable", Type: "boolean", Description: "asset marginable"},
		{Name: "avg_entry_price", Type: "string", Description: "average entry price of the position"},
		{Name: "change_today", Type: "string", Description: "percent change from last day price (by a factor of 1)"},
		{Name: "cost_basis", Type: "string", Description: "total cost basis in dollar"},
		{Name: "current_price", Type: "string", Description: "current asset price per share"},
		{Name: "exchange", Type: "enum", Description: "represents the current exchanges Alpaca supports", EnumValues: []string{"AMEX", "ARCA", "BATS", "NASDAQ", "NYSE", "NYSEARCA", "OTC"}},
		{Name: "lastday_price", Type: "string", Description: "last day’s asset price per share based on the closing value of the last trading day"},
		{Name: "market_value", Type: "string", Description: "total dollar amount of the position"},
		{Name: "qty", Type: "string", Description: "number of shares"},
		{Name: "qty_available", Type: "string", Description: "total number of shares available minus open orders / locked for options covered call"},
		{Name: "side", Type: "string", Description: "“long”"},
		{Name: "symbol", Type: "string", Description: "symbol name of the asset"},
		{Name: "unrealized_intraday_pl", Type: "string", Description: "unrealized profit/loss in dollars for the day"},
		{Name: "unrealized_intraday_plpc", Type: "string", Description: "unrealized profit/loss percent (by a factor of 1)"},
		{Name: "unrealized_pl", Type: "string", Description: "unrealized profit/loss in dollars"},
		{Name: "unrealized_plpc", Type: "string", Description: "unrealized profit/loss percent (by a factor of 1)"},
	},
	"GetAllOrders": {
		{Name: "asset_class", Type: "enum", Description: "this represents the category to which the asset belongs to", EnumValues: []string{"crypto", "us_equity", "us_option"}},
		{Name: "asset_id", Type: "string", Description: "asset ID (For options this represents the option contract ID)"},
		{Name: "canceled_at", Type: "string", Description: "canceled at"},
		{Name: "client_order_id", Type: "string", Description: "client unique order ID"},
		{Name: "created_at", Type: "string", Description: "created at"},
		{Name: "expired_at", Type: "string", Description: "expired at"},
		{Name: "extended_hours", Type: "boolean", Description: "if true, eligible for execution outside regular trading hours"},
		{Name: "failed_at", Type: "string", Description: "failed at"},
		{Name: "filled_at", Type: "string", Description: "filled at"},
		{Name: "filled_avg_price", Type: "string", Description: "filled average price"},
		{Name: "filled_qty", Type: "string", Description: "filled quantity"},
		{Name: "hwm", Type: "string", Description: "highest (lowest) market price seen since the trailing stop order was submitted"},
		{Name: "id", Type: "string", Description: "order ID"},
		{Name: "legs", Type: "[]object", Description: "when querying non-simple order_class orders in a nested style, an array of Order entities associated with this order"},
		{Name: "limit_price", Type: "string", Description: "limit price"},
		{Name: "notional", Type: "string", Description: "ordered notional amount. If entered, qty will be null. Can take up to 9 decimal points"},
		{Name: "order_class", Type: "enum", Description: "order classes supported by Alpaca vary based on the order's security type", EnumValues: []string{"bracket", "mleg", "oco", "oto", "simple"}},
		{Name: "order_type", Type: "string", Description: "deprecated in favour of the field \"type\""},
		{Name: "position_intent", Type: "enum", Description: "represents the desired position strategy", EnumValues: []string{"buy_to_close", "buy_to_open", "sell_to_close", "sell_to_open"}},
		{Name: "qty", Type: "string", Description: "ordered quantity. If entered, notional will be null. Can take up to 9 decimal points. Required if order class is mleg"},
		{Name: "replaced_at", Type: "string", Description: "replaced at"},
		{Name: "replaced_by", Type: "string", Description: "order ID that this order was replaced by"},
		{Name: "replaces", Type: "string", Description: "order ID that this order replaces"},
		{Name: "side", Type: "enum", Description: "represents which side this order was on:\n- buy\n- sell\nRequired for all order classes except for mleg", EnumValues: []string{"buy", "sell"}},
		{Name: "status", Type: "enum", Description: "an order executed through Alpaca can experience several status changes during its lifecycle", EnumValues: []string{"accepted", "accepted_for_bidding", "calculated", "canceled", "done_for_day", "expired", "filled", "new", "partially_filled", "pending_cancel", "pending_new", "pending_replace", "rejected", "replaced", "stopped", "suspended"}},
		{Name: "stop_price", Type: "string", Description: "stop price"},
		{Name: "submitted_at", Type: "string", Description: "submitted at"},
		{Name: "symbol", Type: "string", Description: "asset symbol, required for all order classes except for mleg"},
		{Name: "time_in_force", Type: "enum", Description: "time-In-Force values supported by Alpaca vary based on the order's security type", EnumValues: []string{"cls", "day", "fok", "gtc", "ioc", "opg"}},
		{Name: "trail_percent", Type: "string", Description: "percent value away from the high water mark for trailing stop orders"},
		{Name: "trail_price", Type: "string", Description: "dollar value away from the high water mark for trailing stop orders"},
		{Name: "type", Type: "enum", Description: "order types supported by Alpaca vary based on the order's security type", EnumValues: []string{"limit", "market", "stop", "stop_limit", "trailing_stop"}},
		{Name: "updated_at", Type: "string", Description: "updated at"},
	},
	"GetCryptoFundingTransfer": {
		{Name: "amount", Type: "string", Description: "amount of transfer denominated in the underlying crypto asset"},
		{Name: "asset", Type: "string", Description: "symbol of crypto asset for given transfer (e.g. BTC)"},
		{Name: "chain", Type: "string", Description: "underlying network for given transfer"},
		{Name: "created_at", Type: "string", Description: "timestamp when transfer was created"},
		{Name: "direction", Type: "enum", Description: "direction", EnumValues: []string{"INCOMING", "OUTGOING"}},
		{Name: "fees", Type: "string", Description: "fees"},
		{Name: "from_address", Type: "string", Description: "originating address of the transfer"},
		{Name: "id", Type: "string", Description: "crypto transfer ID"},
		{Name: "network_fee", Type: "string", Description: "network fee"},
		{Name: "status", Type: "enum", Description: "status", EnumValues: []string{"COMPLETE", "FAILED", "PROCESSING"}},
		{Name: "to_address", Type: "string", Description: "destination address of the transfer"},
		{Name: "tx_hash", Type: "string", Description: "on-chain transaction hash (e.g. 0xabc...xyz)"},
		{Name: "usd_value", Type: "string", Description: "equivalent USD value at time of transfer"},
	},
	"GetCryptoPerpFundingTransfer": {
		{Name: "amount", Type: "string", Description: "amount of transfer denominated in the underlying crypto asset"},
		{Name: "asset", Type: "string", Description: "symbol of crypto asset for given transfer (e.g. BTC)"},
		{Name: "chain", Type: "string", Description: "underlying network for given transfer"},
		{Name: "created_at", Type: "string", Description: "timestamp when transfer was created"},
		{Name: "direction", Type: "enum", Description: "direction", EnumValues: []string{"INCOMING", "OUTGOING"}},
		{Name: "fees", Type: "string", Description: "fees"},
		{Name: "from_address", Type: "string", Description: "originating address of the transfer"},
		{Name: "id", Type: "string", Description: "crypto transfer ID"},
		{Name: "network_fee", Type: "string", Description: "network fee"},
		{Name: "status", Type: "enum", Description: "status", EnumValues: []string{"COMPLETE", "FAILED", "PROCESSING"}},
		{Name: "to_address", Type: "string", Description: "destination address of the transfer"},
		{Name: "tx_hash", Type: "string", Description: "on-chain transaction hash (e.g. 0xabc...xyz)"},
		{Name: "usd_value", Type: "string", Description: "equivalent USD value at time of transfer"},
	},
	"GetOpenPosition": {
		{Name: "asset_class", Type: "enum", Description: "this represents the category to which the asset belongs to", EnumValues: []string{"crypto", "us_equity", "us_option"}},
		{Name: "asset_id", Type: "string", Description: "asset ID (For options this represents the option contract ID)"},
		{Name: "asset_marginable", Type: "boolean", Description: "asset marginable"},
		{Name: "avg_entry_price", Type: "string", Description: "average entry price of the position"},
		{Name: "change_today", Type: "string", Description: "percent change from last day price (by a factor of 1)"},
		{Name: "cost_basis", Type: "string", Description: "total cost basis in dollar"},
		{Name: "current_price", Type: "string", Description: "current asset price per share"},
		{Name: "exchange", Type: "enum", Description: "represents the current exchanges Alpaca supports", EnumValues: []string{"AMEX", "ARCA", "BATS", "NASDAQ", "NYSE", "NYSEARCA", "OTC"}},
		{Name: "lastday_price", Type: "string", Description: "last day’s asset price per share based on the closing value of the last trading day"},
		{Name: "market_value", Type: "string", Description: "total dollar amount of the position"},
		{Name: "qty", Type: "string", Description: "number of shares"},
		{Name: "qty_available", Type: "string", Description: "total number of shares available minus open orders / locked for options covered call"},
		{Name: "side", Type: "string", Description: "“long”"},
		{Name: "symbol", Type: "string", Description: "symbol name of the asset"},
		{Name: "unrealized_intraday_pl", Type: "string", Description: "unrealized profit/loss in dollars for the day"},
		{Name: "unrealized_intraday_plpc", Type: "string", Description: "unrealized profit/loss percent (by a factor of 1)"},
		{Name: "unrealized_pl", Type: "string", Description: "unrealized profit/loss in dollars"},
		{Name: "unrealized_plpc", Type: "string", Description: "unrealized profit/loss percent (by a factor of 1)"},
	},
	"GetOrderByClientOrderID": {
		{Name: "asset_class", Type: "enum", Description: "this represents the category to which the asset belongs to", EnumValues: []string{"crypto", "us_equity", "us_option"}},
		{Name: "asset_id", Type: "string", Description: "asset ID (For options this represents the option contract ID)"},
		{Name: "canceled_at", Type: "string", Description: "canceled at"},
		{Name: "client_order_id", Type: "string", Description: "client unique order ID"},
		{Name: "created_at", Type: "string", Description: "created at"},
		{Name: "expired_at", Type: "string", Description: "expired at"},
		{Name: "extended_hours", Type: "boolean", Description: "if true, eligible for execution outside regular trading hours"},
		{Name: "failed_at", Type: "string", Description: "failed at"},
		{Name: "filled_at", Type: "string", Description: "filled at"},
		{Name: "filled_avg_price", Type: "string", Description: "filled average price"},
		{Name: "filled_qty", Type: "string", Description: "filled quantity"},
		{Name: "hwm", Type: "string", Description: "highest (lowest) market price seen since the trailing stop order was submitted"},
		{Name: "id", Type: "string", Description: "order ID"},
		{Name: "legs", Type: "[]object", Description: "when querying non-simple order_class orders in a nested style, an array of Order entities associated with this order"},
		{Name: "limit_price", Type: "string", Description: "limit price"},
		{Name: "notional", Type: "string", Description: "ordered notional amount. If entered, qty will be null. Can take up to 9 decimal points"},
		{Name: "order_class", Type: "enum", Description: "order classes supported by Alpaca vary based on the order's security type", EnumValues: []string{"bracket", "mleg", "oco", "oto", "simple"}},
		{Name: "order_type", Type: "string", Description: "deprecated in favour of the field \"type\""},
		{Name: "position_intent", Type: "enum", Description: "represents the desired position strategy", EnumValues: []string{"buy_to_close", "buy_to_open", "sell_to_close", "sell_to_open"}},
		{Name: "qty", Type: "string", Description: "ordered quantity. If entered, notional will be null. Can take up to 9 decimal points. Required if order class is mleg"},
		{Name: "replaced_at", Type: "string", Description: "replaced at"},
		{Name: "replaced_by", Type: "string", Description: "order ID that this order was replaced by"},
		{Name: "replaces", Type: "string", Description: "order ID that this order replaces"},
		{Name: "side", Type: "enum", Description: "represents which side this order was on:\n- buy\n- sell\nRequired for all order classes except for mleg", EnumValues: []string{"buy", "sell"}},
		{Name: "status", Type: "enum", Description: "an order executed through Alpaca can experience several status changes during its lifecycle", EnumValues: []string{"accepted", "accepted_for_bidding", "calculated", "canceled", "done_for_day", "expired", "filled", "new", "partially_filled", "pending_cancel", "pending_new", "pending_replace", "rejected", "replaced", "stopped", "suspended"}},
		{Name: "stop_price", Type: "string", Description: "stop price"},
		{Name: "submitted_at", Type: "string", Description: "submitted at"},
		{Name: "symbol", Type: "string", Description: "asset symbol, required for all order classes except for mleg"},
		{Name: "time_in_force", Type: "enum", Description: "time-In-Force values supported by Alpaca vary based on the order's security type", EnumValues: []string{"cls", "day", "fok", "gtc", "ioc", "opg"}},
		{Name: "trail_percent", Type: "string", Description: "percent value away from the high water mark for trailing stop orders"},
		{Name: "trail_price", Type: "string", Description: "dollar value away from the high water mark for trailing stop orders"},
		{Name: "type", Type: "enum", Description: "order types supported by Alpaca vary based on the order's security type", EnumValues: []string{"limit", "market", "stop", "stop_limit", "trailing_stop"}},
		{Name: "updated_at", Type: "string", Description: "updated at"},
	},
	"GetOrderByOrderID": {
		{Name: "asset_class", Type: "enum", Description: "this represents the category to which the asset belongs to", EnumValues: []string{"crypto", "us_equity", "us_option"}},
		{Name: "asset_id", Type: "string", Description: "asset ID (For options this represents the option contract ID)"},
		{Name: "canceled_at", Type: "string", Description: "canceled at"},
		{Name: "client_order_id", Type: "string", Description: "client unique order ID"},
		{Name: "created_at", Type: "string", Description: "created at"},
		{Name: "expired_at", Type: "string", Description: "expired at"},
		{Name: "extended_hours", Type: "boolean", Description: "if true, eligible for execution outside regular trading hours"},
		{Name: "failed_at", Type: "string", Description: "failed at"},
		{Name: "filled_at", Type: "string", Description: "filled at"},
		{Name: "filled_avg_price", Type: "string", Description: "filled average price"},
		{Name: "filled_qty", Type: "string", Description: "filled quantity"},
		{Name: "hwm", Type: "string", Description: "highest (lowest) market price seen since the trailing stop order was submitted"},
		{Name: "id", Type: "string", Description: "order ID"},
		{Name: "legs", Type: "[]object", Description: "when querying non-simple order_class orders in a nested style, an array of Order entities associated with this order"},
		{Name: "limit_price", Type: "string", Description: "limit price"},
		{Name: "notional", Type: "string", Description: "ordered notional amount. If entered, qty will be null. Can take up to 9 decimal points"},
		{Name: "order_class", Type: "enum", Description: "order classes supported by Alpaca vary based on the order's security type", EnumValues: []string{"bracket", "mleg", "oco", "oto", "simple"}},
		{Name: "order_type", Type: "string", Description: "deprecated in favour of the field \"type\""},
		{Name: "position_intent", Type: "enum", Description: "represents the desired position strategy", EnumValues: []string{"buy_to_close", "buy_to_open", "sell_to_close", "sell_to_open"}},
		{Name: "qty", Type: "string", Description: "ordered quantity. If entered, notional will be null. Can take up to 9 decimal points. Required if order class is mleg"},
		{Name: "replaced_at", Type: "string", Description: "replaced at"},
		{Name: "replaced_by", Type: "string", Description: "order ID that this order was replaced by"},
		{Name: "replaces", Type: "string", Description: "order ID that this order replaces"},
		{Name: "side", Type: "enum", Description: "represents which side this order was on:\n- buy\n- sell\nRequired for all order classes except for mleg", EnumValues: []string{"buy", "sell"}},
		{Name: "status", Type: "enum", Description: "an order executed through Alpaca can experience several status changes during its lifecycle", EnumValues: []string{"accepted", "accepted_for_bidding", "calculated", "canceled", "done_for_day", "expired", "filled", "new", "partially_filled", "pending_cancel", "pending_new", "pending_replace", "rejected", "replaced", "stopped", "suspended"}},
		{Name: "stop_price", Type: "string", Description: "stop price"},
		{Name: "submitted_at", Type: "string", Description: "submitted at"},
		{Name: "symbol", Type: "string", Description: "asset symbol, required for all order classes except for mleg"},
		{Name: "time_in_force", Type: "enum", Description: "time-In-Force values supported by Alpaca vary based on the order's security type", EnumValues: []string{"cls", "day", "fok", "gtc", "ioc", "opg"}},
		{Name: "trail_percent", Type: "string", Description: "percent value away from the high water mark for trailing stop orders"},
		{Name: "trail_price", Type: "string", Description: "dollar value away from the high water mark for trailing stop orders"},
		{Name: "type", Type: "enum", Description: "order types supported by Alpaca vary based on the order's security type", EnumValues: []string{"limit", "market", "stop", "stop_limit", "trailing_stop"}},
		{Name: "updated_at", Type: "string", Description: "updated at"},
	},
	"GetWatchlistByID": {
		{Name: "account_id", Type: "string", Description: "account ID"},
		{Name: "assets", Type: "[]object", Description: "the content of this watchlist, in the order as registered by the client"},
		{Name: "created_at", Type: "string", Description: "created at"},
		{Name: "id", Type: "string", Description: "watchlist id"},
		{Name: "name", Type: "string", Description: "user-defined watchlist name (up to 64 characters)"},
		{Name: "updated_at", Type: "string", Description: "updated at"},
	},
	"GetWatchlistByName": {
		{Name: "account_id", Type: "string", Description: "account ID"},
		{Name: "assets", Type: "[]object", Description: "the content of this watchlist, in the order as registered by the client"},
		{Name: "created_at", Type: "string", Description: "created at"},
		{Name: "id", Type: "string", Description: "watchlist id"},
		{Name: "name", Type: "string", Description: "user-defined watchlist name (up to 64 characters)"},
		{Name: "updated_at", Type: "string", Description: "updated at"},
	},
	"GetWatchlists": {
		{Name: "account_id", Type: "string", Description: "account ID"},
		{Name: "created_at", Type: "string", Description: "created at"},
		{Name: "id", Type: "string", Description: "watchlist id"},
		{Name: "name", Type: "string", Description: "user-defined watchlist name (up to 64 characters)"},
		{Name: "updated_at", Type: "string", Description: "updated at"},
	},
	"ListCryptoFundingTransfers": {
		{Name: "amount", Type: "string", Description: "amount of transfer denominated in the underlying crypto asset"},
		{Name: "asset", Type: "string", Description: "symbol of crypto asset for given transfer (e.g. BTC)"},
		{Name: "chain", Type: "string", Description: "underlying network for given transfer"},
		{Name: "created_at", Type: "string", Description: "timestamp when transfer was created"},
		{Name: "direction", Type: "enum", Description: "direction", EnumValues: []string{"INCOMING", "OUTGOING"}},
		{Name: "fees", Type: "string", Description: "fees"},
		{Name: "from_address", Type: "string", Description: "originating address of the transfer"},
		{Name: "id", Type: "string", Description: "crypto transfer ID"},
		{Name: "network_fee", Type: "string", Description: "network fee"},
		{Name: "status", Type: "enum", Description: "status", EnumValues: []string{"COMPLETE", "FAILED", "PROCESSING"}},
		{Name: "to_address", Type: "string", Description: "destination address of the transfer"},
		{Name: "tx_hash", Type: "string", Description: "on-chain transaction hash (e.g. 0xabc...xyz)"},
		{Name: "usd_value", Type: "string", Description: "equivalent USD value at time of transfer"},
	},
	"ListCryptoFundingWallets": {
		{Name: "address", Type: "string", Description: "address"},
		{Name: "chain", Type: "string", Description: "chain"},
		{Name: "created_at", Type: "string", Description: "timestamp (RFC3339) of account creation"},
	},
	"ListCryptoPerpFundingTransfers": {
		{Name: "amount", Type: "string", Description: "amount of transfer denominated in the underlying crypto asset"},
		{Name: "asset", Type: "string", Description: "symbol of crypto asset for given transfer (e.g. BTC)"},
		{Name: "chain", Type: "string", Description: "underlying network for given transfer"},
		{Name: "created_at", Type: "string", Description: "timestamp when transfer was created"},
		{Name: "direction", Type: "enum", Description: "direction", EnumValues: []string{"INCOMING", "OUTGOING"}},
		{Name: "fees", Type: "string", Description: "fees"},
		{Name: "from_address", Type: "string", Description: "originating address of the transfer"},
		{Name: "id", Type: "string", Description: "crypto transfer ID"},
		{Name: "network_fee", Type: "string", Description: "network fee"},
		{Name: "status", Type: "enum", Description: "status", EnumValues: []string{"COMPLETE", "FAILED", "PROCESSING"}},
		{Name: "to_address", Type: "string", Description: "destination address of the transfer"},
		{Name: "tx_hash", Type: "string", Description: "on-chain transaction hash (e.g. 0xabc...xyz)"},
		{Name: "usd_value", Type: "string", Description: "equivalent USD value at time of transfer"},
	},
	"ListCryptoPerpFundingWallets": {
		{Name: "address", Type: "string", Description: "address"},
		{Name: "chain", Type: "string", Description: "chain"},
		{Name: "created_at", Type: "string", Description: "timestamp (RFC3339) of account creation"},
	},
	"ListWhitelistedAddress": {
		{Name: "address", Type: "string", Description: "whitelisted address"},
		{Name: "asset", Type: "string", Description: "symbol of underlying asset for the whitelisted address"},
		{Name: "chain", Type: "string", Description: "underlying network this address represents"},
		{Name: "created_at", Type: "string", Description: "timestamp (RFC3339) of account creation"},
		{Name: "id", Type: "string", Description: "unique ID for whitelisted address"},
		{Name: "status", Type: "enum", Description: "status of whitelisted address which is either ACTIVE or PENDING", EnumValues: []string{"APPROVED", "PENDING"}},
	},
	"ListWhitelistedPerpAddress": {
		{Name: "address", Type: "string", Description: "whitelisted address"},
		{Name: "asset", Type: "string", Description: "symbol of underlying asset for the whitelisted address"},
		{Name: "chain", Type: "string", Description: "underlying network this address represents"},
		{Name: "created_at", Type: "string", Description: "timestamp (RFC3339) of account creation"},
		{Name: "id", Type: "string", Description: "unique ID for whitelisted address"},
		{Name: "status", Type: "enum", Description: "status of whitelisted address which is either ACTIVE or PENDING", EnumValues: []string{"APPROVED", "PENDING"}},
	},
	"OptionBars": {
		{Name: "bars", Type: "map[string][]object", Description: "bars"},
		{Name: "currency", Type: "string", Description: "currency"},
		{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
	},
	"PatchAccountConfig": {
		{Name: "disable_overnight_trading", Type: "boolean", Description: "if true, overnight trading is disabled"},
		{Name: "dtbp_check", Type: "enum", Description: "both, entry, or exit. Controls Day Trading Margin Call (DTMC) checks", EnumValues: []string{"both", "entry", "exit"}},
		{Name: "fractional_trading", Type: "boolean", Description: "if true, account is able to participate in fractional trading"},
		{Name: "max_margin_multiplier", Type: "string", Description: "can be \"1\", \"2\", or \"4\""},
		{Name: "max_options_trading_level", Type: "integer", Description: "desired maximum options trading level. 0=disabled, 1=Covered Call/Cash-Secured Put, 2=Long Call/Put, 3=Spreads/Straddles", EnumValues: []string{"0", "1", "2", "3"}},
		{Name: "no_shorting", Type: "boolean", Description: "if true, account becomes long-only mode"},
		{Name: "pdt_check", Type: "string", Description: "both, entry, or exit"},
		{Name: "ptp_no_exception_entry", Type: "boolean", Description: "if set to true then Alpaca will accept orders for PTP symbols with no exception. Default is false"},
		{Name: "suspend_trade", Type: "boolean", Description: "if true, new orders are blocked"},
		{Name: "trade_confirm_email", Type: "string", Description: "all or none. If none, emails for order fills are not sent"},
	},
	"PatchOrderByOrderID": {
		{Name: "asset_class", Type: "enum", Description: "this represents the category to which the asset belongs to", EnumValues: []string{"crypto", "us_equity", "us_option"}},
		{Name: "asset_id", Type: "string", Description: "asset ID (For options this represents the option contract ID)"},
		{Name: "canceled_at", Type: "string", Description: "canceled at"},
		{Name: "client_order_id", Type: "string", Description: "client unique order ID"},
		{Name: "created_at", Type: "string", Description: "created at"},
		{Name: "expired_at", Type: "string", Description: "expired at"},
		{Name: "extended_hours", Type: "boolean", Description: "if true, eligible for execution outside regular trading hours"},
		{Name: "failed_at", Type: "string", Description: "failed at"},
		{Name: "filled_at", Type: "string", Description: "filled at"},
		{Name: "filled_avg_price", Type: "string", Description: "filled average price"},
		{Name: "filled_qty", Type: "string", Description: "filled quantity"},
		{Name: "hwm", Type: "string", Description: "highest (lowest) market price seen since the trailing stop order was submitted"},
		{Name: "id", Type: "string", Description: "order ID"},
		{Name: "legs", Type: "[]object", Description: "when querying non-simple order_class orders in a nested style, an array of Order entities associated with this order"},
		{Name: "limit_price", Type: "string", Description: "limit price"},
		{Name: "notional", Type: "string", Description: "ordered notional amount. If entered, qty will be null. Can take up to 9 decimal points"},
		{Name: "order_class", Type: "enum", Description: "order classes supported by Alpaca vary based on the order's security type", EnumValues: []string{"bracket", "mleg", "oco", "oto", "simple"}},
		{Name: "order_type", Type: "string", Description: "deprecated in favour of the field \"type\""},
		{Name: "position_intent", Type: "enum", Description: "represents the desired position strategy", EnumValues: []string{"buy_to_close", "buy_to_open", "sell_to_close", "sell_to_open"}},
		{Name: "qty", Type: "string", Description: "ordered quantity. If entered, notional will be null. Can take up to 9 decimal points. Required if order class is mleg"},
		{Name: "replaced_at", Type: "string", Description: "replaced at"},
		{Name: "replaced_by", Type: "string", Description: "order ID that this order was replaced by"},
		{Name: "replaces", Type: "string", Description: "order ID that this order replaces"},
		{Name: "side", Type: "enum", Description: "represents which side this order was on:\n- buy\n- sell\nRequired for all order classes except for mleg", EnumValues: []string{"buy", "sell"}},
		{Name: "status", Type: "enum", Description: "an order executed through Alpaca can experience several status changes during its lifecycle", EnumValues: []string{"accepted", "accepted_for_bidding", "calculated", "canceled", "done_for_day", "expired", "filled", "new", "partially_filled", "pending_cancel", "pending_new", "pending_replace", "rejected", "replaced", "stopped", "suspended"}},
		{Name: "stop_price", Type: "string", Description: "stop price"},
		{Name: "submitted_at", Type: "string", Description: "submitted at"},
		{Name: "symbol", Type: "string", Description: "asset symbol, required for all order classes except for mleg"},
		{Name: "time_in_force", Type: "enum", Description: "time-In-Force values supported by Alpaca vary based on the order's security type", EnumValues: []string{"cls", "day", "fok", "gtc", "ioc", "opg"}},
		{Name: "trail_percent", Type: "string", Description: "percent value away from the high water mark for trailing stop orders"},
		{Name: "trail_price", Type: "string", Description: "dollar value away from the high water mark for trailing stop orders"},
		{Name: "type", Type: "enum", Description: "order types supported by Alpaca vary based on the order's security type", EnumValues: []string{"limit", "market", "stop", "stop_limit", "trailing_stop"}},
		{Name: "updated_at", Type: "string", Description: "updated at"},
	},
	"PostOrder": {
		{Name: "asset_class", Type: "enum", Description: "this represents the category to which the asset belongs to", EnumValues: []string{"crypto", "us_equity", "us_option"}},
		{Name: "asset_id", Type: "string", Description: "asset ID (For options this represents the option contract ID)"},
		{Name: "canceled_at", Type: "string", Description: "canceled at"},
		{Name: "client_order_id", Type: "string", Description: "client unique order ID"},
		{Name: "created_at", Type: "string", Description: "created at"},
		{Name: "expired_at", Type: "string", Description: "expired at"},
		{Name: "extended_hours", Type: "boolean", Description: "if true, eligible for execution outside regular trading hours"},
		{Name: "failed_at", Type: "string", Description: "failed at"},
		{Name: "filled_at", Type: "string", Description: "filled at"},
		{Name: "filled_avg_price", Type: "string", Description: "filled average price"},
		{Name: "filled_qty", Type: "string", Description: "filled quantity"},
		{Name: "hwm", Type: "string", Description: "highest (lowest) market price seen since the trailing stop order was submitted"},
		{Name: "id", Type: "string", Description: "order ID"},
		{Name: "legs", Type: "[]object", Description: "when querying non-simple order_class orders in a nested style, an array of Order entities associated with this order"},
		{Name: "limit_price", Type: "string", Description: "limit price"},
		{Name: "notional", Type: "string", Description: "ordered notional amount. If entered, qty will be null. Can take up to 9 decimal points"},
		{Name: "order_class", Type: "enum", Description: "order classes supported by Alpaca vary based on the order's security type", EnumValues: []string{"bracket", "mleg", "oco", "oto", "simple"}},
		{Name: "order_type", Type: "string", Description: "deprecated in favour of the field \"type\""},
		{Name: "position_intent", Type: "enum", Description: "represents the desired position strategy", EnumValues: []string{"buy_to_close", "buy_to_open", "sell_to_close", "sell_to_open"}},
		{Name: "qty", Type: "string", Description: "ordered quantity. If entered, notional will be null. Can take up to 9 decimal points. Required if order class is mleg"},
		{Name: "replaced_at", Type: "string", Description: "replaced at"},
		{Name: "replaced_by", Type: "string", Description: "order ID that this order was replaced by"},
		{Name: "replaces", Type: "string", Description: "order ID that this order replaces"},
		{Name: "side", Type: "enum", Description: "represents which side this order was on:\n- buy\n- sell\nRequired for all order classes except for mleg", EnumValues: []string{"buy", "sell"}},
		{Name: "status", Type: "enum", Description: "an order executed through Alpaca can experience several status changes during its lifecycle", EnumValues: []string{"accepted", "accepted_for_bidding", "calculated", "canceled", "done_for_day", "expired", "filled", "new", "partially_filled", "pending_cancel", "pending_new", "pending_replace", "rejected", "replaced", "stopped", "suspended"}},
		{Name: "stop_price", Type: "string", Description: "stop price"},
		{Name: "submitted_at", Type: "string", Description: "submitted at"},
		{Name: "symbol", Type: "string", Description: "asset symbol, required for all order classes except for mleg"},
		{Name: "time_in_force", Type: "enum", Description: "time-In-Force values supported by Alpaca vary based on the order's security type", EnumValues: []string{"cls", "day", "fok", "gtc", "ioc", "opg"}},
		{Name: "trail_percent", Type: "string", Description: "percent value away from the high water mark for trailing stop orders"},
		{Name: "trail_price", Type: "string", Description: "dollar value away from the high water mark for trailing stop orders"},
		{Name: "type", Type: "enum", Description: "order types supported by Alpaca vary based on the order's security type", EnumValues: []string{"limit", "market", "stop", "stop_limit", "trailing_stop"}},
		{Name: "updated_at", Type: "string", Description: "updated at"},
	},
	"PostWatchlist": {
		{Name: "account_id", Type: "string", Description: "account ID"},
		{Name: "assets", Type: "[]object", Description: "the content of this watchlist, in the order as registered by the client"},
		{Name: "created_at", Type: "string", Description: "created at"},
		{Name: "id", Type: "string", Description: "watchlist id"},
		{Name: "name", Type: "string", Description: "user-defined watchlist name (up to 64 characters)"},
		{Name: "updated_at", Type: "string", Description: "updated at"},
	},
	"RemoveAssetFromWatchlist": {
		{Name: "account_id", Type: "string", Description: "account ID"},
		{Name: "assets", Type: "[]object", Description: "the content of this watchlist, in the order as registered by the client"},
		{Name: "created_at", Type: "string", Description: "created at"},
		{Name: "id", Type: "string", Description: "watchlist id"},
		{Name: "name", Type: "string", Description: "user-defined watchlist name (up to 64 characters)"},
		{Name: "updated_at", Type: "string", Description: "updated at"},
	},
	"UpdateWatchlistByID": {
		{Name: "account_id", Type: "string", Description: "account ID"},
		{Name: "assets", Type: "[]object", Description: "the content of this watchlist, in the order as registered by the client"},
		{Name: "created_at", Type: "string", Description: "created at"},
		{Name: "id", Type: "string", Description: "watchlist id"},
		{Name: "name", Type: "string", Description: "user-defined watchlist name (up to 64 characters)"},
		{Name: "updated_at", Type: "string", Description: "updated at"},
	},
	"UpdateWatchlistByName": {
		{Name: "account_id", Type: "string", Description: "account ID"},
		{Name: "assets", Type: "[]object", Description: "the content of this watchlist, in the order as registered by the client"},
		{Name: "created_at", Type: "string", Description: "created at"},
		{Name: "id", Type: "string", Description: "watchlist id"},
		{Name: "name", Type: "string", Description: "user-defined watchlist name (up to 64 characters)"},
		{Name: "updated_at", Type: "string", Description: "updated at"},
	},
}

// OperationSummaries maps operation names to their summaries.
var OperationSummaries = map[string]string{
	"Calendar":                           "Get market calendar",
	"Clock":                              "Get market clock",
	"CorporateActions":                   "Get corporate actions",
	"CryptoBars":                         "Get historical bars",
	"CryptoLatestBars":                   "Get latest bars",
	"CryptoLatestOrderbooks":             "Get latest orderbook",
	"CryptoLatestQuotes":                 "Get latest quotes",
	"CryptoLatestTrades":                 "Get latest trades",
	"CryptoPerpLatestBars":               "Get latest bars",
	"CryptoPerpLatestFuturesPricing":     "Get latest pricing",
	"CryptoPerpLatestOrderbooks":         "Get latest orderbook",
	"CryptoPerpLatestQuotes":             "Get latest quotes",
	"CryptoPerpLatestTrades":             "Get latest trades",
	"CryptoQuotes":                       "Get historical quotes",
	"CryptoSnapshots":                    "Get snapshots",
	"CryptoTrades":                       "Get historical trades",
	"FixedIncomeLatestPrices":            "Get latest prices",
	"LatestRates":                        "Get latest rates for currency pairs",
	"LegacyClock":                        "Get US market clock",
	"MostActives":                        "Get most active stocks",
	"Movers":                             "Get top market movers",
	"News":                               "Get news articles",
	"OptionChain":                        "Get option chain",
	"OptionLatestQuotes":                 "Get latest quotes",
	"OptionLatestTrades":                 "Get latest trades",
	"OptionSnapshots":                    "Get snapshots",
	"OptionTrades":                       "Get historical trades",
	"Rates":                              "Get historical rates for currency pairs",
	"StockAuctionSingle":                 "Get historical auctions (single)",
	"StockAuctions":                      "Get historical auctions",
	"StockBarSingle":                     "Get historical bars (single symbol)",
	"StockBars":                          "Get historical bars",
	"StockLatestBarSingle":               "Get latest bar (single symbol)",
	"StockLatestBars":                    "Get latest bars",
	"StockLatestQuoteSingle":             "Get latest quote (single symbol)",
	"StockLatestQuotes":                  "Get latest quotes",
	"StockLatestTradeSingle":             "Get latest trade (single symbol)",
	"StockLatestTrades":                  "Get latest trades",
	"StockQuoteSingle":                   "Get historical quotes (single symbol)",
	"StockQuotes":                        "Get historical quotes",
	"StockTradeSingle":                   "Get historical trades (single symbol)",
	"StockTrades":                        "Get historical trades",
	"UsCorporates":                       "Get US corporates",
	"UsTreasuries":                       "Get US treasuries",
	"AddAssetToWatchlist":                "Add asset to watchlist",
	"AddAssetToWatchlistByName":          "Add asset to watchlist by name",
	"CreateCryptoPerpTransferForAccount": "Request a new withdrawal",
	"CreateCryptoTransferForAccount":     "Request a new withdrawal",
	"CreateWhitelistedAddress":           "Request a new whitelisted address",
	"CreateWhitelistedPerpAddress":       "Request a new whitelisted address",
	"DeleteAllOpenPositions":             "Close all positions",
	"DeleteAllOrders":                    "Delete all orders",
	"DeleteOpenPosition":                 "Close a position",
	"GetOptionContractSymbolOrID":        "Get an option contract by ID or symbol",
	"GetV2Assets":                        "Get assets",
	"GetV2AssetsSymbolOrAssetID":         "Get an asset by ID or symbol",
	"GetAccount":                         "Get account",
	"GetAccountConfig":                   "Get account configurations",
	"GetAccountPortfolioHistory":         "Get account portfolio history",
	"GetAllOpenPositions":                "List all open positions",
	"GetAllOrders":                       "Get all orders",
	"GetCryptoFundingTransfer":           "Retrieve a crypto funding transfer",
	"GetCryptoPerpFundingTransfer":       "Retrieve a crypto funding transfer",
	"GetOpenPosition":                    "Get an open position",
	"GetOrderByClientOrderID":            "Get order by client order ID",
	"GetOrderByOrderID":                  "Get order by ID",
	"GetWatchlistByID":                   "Get watchlist by ID",
	"GetWatchlistByName":                 "Get watchlist by name",
	"GetWatchlists":                      "Get all watchlists",
	"ListCryptoFundingTransfers":         "Retrieve crypto funding transfers",
	"ListCryptoFundingWallets":           "Retrieve crypto funding wallets",
	"ListCryptoPerpFundingTransfers":     "Retrieve crypto funding transfers",
	"ListCryptoPerpFundingWallets":       "Retrieve crypto funding wallets",
	"ListWhitelistedAddress":             "Get an array of whitelisted addresses",
	"ListWhitelistedPerpAddress":         "Get an array of whitelisted addresses",
	"OptionBars":                         "Get historical bars",
	"PatchAccountConfig":                 "Update account configurations",
	"PatchOrderByOrderID":                "Replace order by ID",
	"PostOrder":                          "Create an order",
	"PostWatchlist":                      "Create watchlist",
	"RemoveAssetFromWatchlist":           "Delete symbol from watchlist",
	"UpdateWatchlistByID":                "Update watchlist by id",
	"UpdateWatchlistByName":              "Update watchlist by name",
}

// ArrayResponses tracks which operations return arrays vs single objects.
var ArrayResponses = map[string]bool{
	"DeleteAllOpenPositions": true,
	"DeleteAllOrders":        true,
	"GetV2Assets":            true,
	"GetAllOpenPositions":    true,
	"GetAllOrders":           true,
	"GetWatchlists":          true,
}

// AllOps lists every generated Op for iteration in tests and tooling.
var AllOps = []Op{
	CalendarOp,
	ClockOp,
	CorporateActionsOp,
	CryptoBarsOp,
	CryptoLatestBarsOp,
	CryptoLatestOrderbooksOp,
	CryptoLatestQuotesOp,
	CryptoLatestTradesOp,
	CryptoPerpLatestBarsOp,
	CryptoPerpLatestFuturesPricingOp,
	CryptoPerpLatestOrderbooksOp,
	CryptoPerpLatestQuotesOp,
	CryptoPerpLatestTradesOp,
	CryptoQuotesOp,
	CryptoSnapshotsOp,
	CryptoTradesOp,
	FixedIncomeLatestPricesOp,
	LatestRatesOp,
	LegacyCalendarOp,
	LegacyClockOp,
	LogosOp,
	MostActivesOp,
	MoversOp,
	NewsOp,
	OptionChainOp,
	OptionLatestQuotesOp,
	OptionLatestTradesOp,
	OptionMetaConditionsOp,
	OptionMetaExchangesOp,
	OptionSnapshotsOp,
	OptionTradesOp,
	RatesOp,
	StockAuctionSingleOp,
	StockAuctionsOp,
	StockBarSingleOp,
	StockBarsOp,
	StockLatestBarSingleOp,
	StockLatestBarsOp,
	StockLatestQuoteSingleOp,
	StockLatestQuotesOp,
	StockLatestTradeSingleOp,
	StockLatestTradesOp,
	StockMetaConditionsOp,
	StockMetaExchangesOp,
	StockQuoteSingleOp,
	StockQuotesOp,
	StockSnapshotSingleOp,
	StockSnapshotsOp,
	StockTradeSingleOp,
	StockTradesOp,
	UsCorporatesOp,
	UsTreasuriesOp,
	AddAssetToWatchlistOp,
	AddAssetToWatchlistByNameOp,
	CreateCryptoPerpTransferForAccountOp,
	CreateCryptoTransferForAccountOp,
	CreateWhitelistedAddressOp,
	CreateWhitelistedPerpAddressOp,
	DeleteAllOpenPositionsOp,
	DeleteAllOrdersOp,
	DeleteOpenPositionOp,
	DeleteOrderByOrderIDOp,
	DeleteWatchlistByIDOp,
	DeleteWatchlistByNameOp,
	DeleteWhitelistedAddressOp,
	DeleteWhitelistedPerpAddressOp,
	GetOptionContractSymbolOrIDOp,
	GetOptionsContractsOp,
	GetV2AssetsOp,
	GetV2AssetsSymbolOrAssetIDOp,
	GetV2CorporateActionsAnnouncementsOp,
	GetV2CorporateActionsAnnouncementsIDOp,
	GetAccountOp,
	GetAccountActivitiesOp,
	GetAccountActivitiesByActivityTypeOp,
	GetAccountConfigOp,
	GetAccountPortfolioHistoryOp,
	GetAllOpenPositionsOp,
	GetAllOrdersOp,
	GetCryptoFundingTransferOp,
	GetCryptoPerpAccountLeverageOp,
	GetCryptoPerpAccountVitalsOp,
	GetCryptoPerpFundingTransferOp,
	GetCryptoPerpTransferEstimateOp,
	GetCryptoTransferEstimateOp,
	GetOpenPositionOp,
	GetOrderByClientOrderIDOp,
	GetOrderByOrderIDOp,
	GetWatchlistByIDOp,
	GetWatchlistByNameOp,
	GetWatchlistsOp,
	ListCryptoFundingTransfersOp,
	ListCryptoFundingWalletsOp,
	ListCryptoPerpFundingTransfersOp,
	ListCryptoPerpFundingWalletsOp,
	ListWhitelistedAddressOp,
	ListWhitelistedPerpAddressOp,
	OptionBarsOp,
	OptionDoNotExerciseOp,
	OptionExerciseOp,
	PatchAccountConfigOp,
	PatchOrderByOrderIDOp,
	PostOrderOp,
	PostWatchlistOp,
	RemoveAssetFromWatchlistOp,
	SetCryptoPerpAccountLeverageOp,
	UpdateWatchlistByIDOp,
	UpdateWatchlistByNameOp,
}
