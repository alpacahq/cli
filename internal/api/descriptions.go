// Code generated from api/specs; DO NOT EDIT.

package api

// FlagDef describes a CLI flag derived from the OpenAPI spec.
type FlagDef struct {
	Name        string // kebab-case CLI flag name
	OASName     string // original OAS property/parameter name
	Type        string // "string", "bool", "int", "float64"
	Default     string
	Description string
	Completions []string // enum values for shell completion
}

type calendarOp struct {
	Summary  string
	End      string
	Market   string
	Start    string
	Timezone string
}

var CalendarOp = calendarOp{
	Summary:  "Get market calendar",
	End:      "last date to retrieve data for (inclusive). Default: one week from the start date",
	Market:   "market",
	Start:    "first date to retrieve data for (inclusive). Default: today",
	Timezone: "timezone of the times. Default: the timezone of the market",
}

var CalendarFlags = []FlagDef{
	{Name: "end", OASName: "end", Type: "string", Description: "last date to retrieve data for (inclusive). Default: one week from the start date"},
	{Name: "start", OASName: "start", Type: "string", Description: "first date to retrieve data for (inclusive). Default: today"},
	{Name: "timezone", OASName: "timezone", Type: "string", Description: "timezone of the times. Default: the timezone of the market", Completions: []string{"UTC"}},
}

type clockOp struct {
	Summary string
	Markets string
	Time    string
}

var ClockOp = clockOp{
	Summary: "Get market clock",
	Markets: "comma-separated list of markets",
	Time:    "instead of the current time, use this time for the clock",
}

var ClockFlags = []FlagDef{
	{Name: "markets", OASName: "markets", Type: "string", Description: "comma-separated list of markets"},
	{Name: "time", OASName: "time", Type: "string", Description: "instead of the current time, use this time for the clock"},
}

type corporateActionsOp struct {
	Summary   string
	Cusips    string
	End       string
	Ids       string
	Limit     string
	PageToken string
	Sort      string
	Start     string
	Symbols   string
	Types     string
}

var CorporateActionsOp = corporateActionsOp{
	Summary:   "Get corporate actions",
	Cusips:    "A comma-separated list of CUSIPs",
	End:       "inclusive end of the interval",
	Ids:       "A comma-separated list of corporate action IDs",
	Limit:     "maximum number of corporate actions to return in a response.",
	PageToken: "pagination token from which to continue",
	Sort:      "sort data in ascending or descending order",
	Start:     "inclusive start of the interval",
	Symbols:   "A comma-separated list of symbols",
	Types:     "A comma-separated list of types",
}

var CorporateActionsFlags = []FlagDef{
	{Name: "cusips", OASName: "cusips", Type: "string", Description: "A comma-separated list of CUSIPs"},
	{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval"},
	{Name: "ids", OASName: "ids", Type: "string", Description: "A comma-separated list of corporate action IDs"},
	{Name: "limit", OASName: "limit", Type: "int", Default: "100", Description: "maximum number of corporate actions to return in a response."},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue"},
	{Name: "sort", OASName: "sort", Type: "string", Description: "sort data in ascending or descending order"},
	{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval"},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of symbols"},
	{Name: "types", OASName: "types", Type: "string", Description: "A comma-separated list of types"},
}

type cryptoBarsOp struct {
	Summary   string
	End       string
	Limit     string
	Loc       string
	PageToken string
	Sort      string
	Start     string
	Symbols   string
	Timeframe string
}

var CryptoBarsOp = cryptoBarsOp{
	Summary:   "Get historical bars",
	End:       "inclusive end of the interval",
	Limit:     "maximum number of data points to return in the response page.",
	Loc:       "crypto location from where the historical market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU",
	PageToken: "pagination token from which to continue",
	Sort:      "sort data in ascending or descending order",
	Start:     "inclusive start of the interval",
	Symbols:   "A comma-separated list of crypto symbols",
	Timeframe: "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g",
}

var CryptoBarsFlags = []FlagDef{
	{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval"},
	{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page."},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue"},
	{Name: "sort", OASName: "sort", Type: "string", Description: "sort data in ascending or descending order"},
	{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval"},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols"},
	{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g"},
}

type cryptoLatestBarsOp struct {
	Summary string
	Loc     string
	Symbols string
}

var CryptoLatestBarsOp = cryptoLatestBarsOp{
	Summary: "Get latest bars",
	Loc:     "crypto location from where the latest market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU",
	Symbols: "A comma-separated list of crypto symbols",
}

var CryptoLatestBarsFlags = []FlagDef{
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols"},
}

type cryptoLatestOrderbooksOp struct {
	Summary string
	Loc     string
	Symbols string
}

var CryptoLatestOrderbooksOp = cryptoLatestOrderbooksOp{
	Summary: "Get latest orderbook",
	Loc:     "crypto location from where the latest market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU",
	Symbols: "A comma-separated list of crypto symbols",
}

var CryptoLatestOrderbooksFlags = []FlagDef{
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols"},
}

type cryptoLatestQuotesOp struct {
	Summary string
	Loc     string
	Symbols string
}

var CryptoLatestQuotesOp = cryptoLatestQuotesOp{
	Summary: "Get latest quotes",
	Loc:     "crypto location from where the latest market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU",
	Symbols: "A comma-separated list of crypto symbols",
}

var CryptoLatestQuotesFlags = []FlagDef{
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols"},
}

type cryptoLatestTradesOp struct {
	Summary string
	Loc     string
	Symbols string
}

var CryptoLatestTradesOp = cryptoLatestTradesOp{
	Summary: "Get latest trades",
	Loc:     "crypto location from where the latest market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU",
	Symbols: "A comma-separated list of crypto symbols",
}

var CryptoLatestTradesFlags = []FlagDef{
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols"},
}

type cryptoPerpLatestBarsOp struct {
	Summary string
	Loc     string
	Symbols string
}

var CryptoPerpLatestBarsOp = cryptoPerpLatestBarsOp{
	Summary: "Get latest bars",
	Loc:     "crypto perpetual location",
	Symbols: "A comma-separated list of crypto symbols",
}

var CryptoPerpLatestBarsFlags = []FlagDef{
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols"},
}

type cryptoPerpLatestFuturesPricingOp struct {
	Summary string
	Loc     string
	Symbols string
}

var CryptoPerpLatestFuturesPricingOp = cryptoPerpLatestFuturesPricingOp{
	Summary: "Get latest pricing",
	Loc:     "crypto perpetual location",
	Symbols: "A comma-separated list of crypto symbols",
}

var CryptoPerpLatestFuturesPricingFlags = []FlagDef{
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols"},
}

type cryptoPerpLatestOrderbooksOp struct {
	Summary string
	Loc     string
	Symbols string
}

var CryptoPerpLatestOrderbooksOp = cryptoPerpLatestOrderbooksOp{
	Summary: "Get latest orderbook",
	Loc:     "crypto perpetual location",
	Symbols: "A comma-separated list of crypto symbols",
}

var CryptoPerpLatestOrderbooksFlags = []FlagDef{
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols"},
}

type cryptoPerpLatestQuotesOp struct {
	Summary string
	Loc     string
	Symbols string
}

var CryptoPerpLatestQuotesOp = cryptoPerpLatestQuotesOp{
	Summary: "Get latest quotes",
	Loc:     "crypto perpetual location",
	Symbols: "A comma-separated list of crypto symbols",
}

var CryptoPerpLatestQuotesFlags = []FlagDef{
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols"},
}

type cryptoPerpLatestTradesOp struct {
	Summary string
	Loc     string
	Symbols string
}

var CryptoPerpLatestTradesOp = cryptoPerpLatestTradesOp{
	Summary: "Get latest trades",
	Loc:     "crypto perpetual location",
	Symbols: "A comma-separated list of crypto symbols",
}

var CryptoPerpLatestTradesFlags = []FlagDef{
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols"},
}

type cryptoQuotesOp struct {
	Summary   string
	End       string
	Limit     string
	Loc       string
	PageToken string
	Sort      string
	Start     string
	Symbols   string
}

var CryptoQuotesOp = cryptoQuotesOp{
	Summary:   "Get historical quotes",
	End:       "inclusive end of the interval",
	Limit:     "maximum number of data points to return in the response page.",
	Loc:       "crypto location from where the historical market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU",
	PageToken: "pagination token from which to continue",
	Sort:      "sort data in ascending or descending order",
	Start:     "inclusive start of the interval",
	Symbols:   "A comma-separated list of crypto symbols",
}

var CryptoQuotesFlags = []FlagDef{
	{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval"},
	{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page."},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue"},
	{Name: "sort", OASName: "sort", Type: "string", Description: "sort data in ascending or descending order"},
	{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval"},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols"},
}

type cryptoSnapshotsOp struct {
	Summary string
	Loc     string
	Symbols string
}

var CryptoSnapshotsOp = cryptoSnapshotsOp{
	Summary: "Get snapshots",
	Loc:     "crypto location from where the latest market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU",
	Symbols: "A comma-separated list of crypto symbols",
}

var CryptoSnapshotsFlags = []FlagDef{
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols"},
}

type cryptoTradesOp struct {
	Summary   string
	End       string
	Limit     string
	Loc       string
	PageToken string
	Sort      string
	Start     string
	Symbols   string
}

var CryptoTradesOp = cryptoTradesOp{
	Summary:   "Get historical trades",
	End:       "inclusive end of the interval",
	Limit:     "maximum number of data points to return in the response page.",
	Loc:       "crypto location from where the historical market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU",
	PageToken: "pagination token from which to continue",
	Sort:      "sort data in ascending or descending order",
	Start:     "inclusive start of the interval",
	Symbols:   "A comma-separated list of crypto symbols",
}

var CryptoTradesFlags = []FlagDef{
	{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval"},
	{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page."},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue"},
	{Name: "sort", OASName: "sort", Type: "string", Description: "sort data in ascending or descending order"},
	{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval"},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols"},
}

type fixedIncomeLatestPricesOp struct {
	Summary string
	Isins   string
}

var FixedIncomeLatestPricesOp = fixedIncomeLatestPricesOp{
	Summary: "Get latest prices",
	Isins:   "A comma-separated list of ISINs with a limit of 1000",
}

var FixedIncomeLatestPricesFlags = []FlagDef{
	{Name: "isins", OASName: "isins", Type: "string", Description: "A comma-separated list of ISINs with a limit of 1000"},
}

type latestRatesOp struct {
	Summary       string
	CurrencyPairs string
}

var LatestRatesOp = latestRatesOp{
	Summary:       "Get latest rates for currency pairs",
	CurrencyPairs: "currency pairs",
}

var LatestRatesFlags = []FlagDef{
	{Name: "currency-pairs", OASName: "currency_pairs", Type: "string", Description: "currency pairs"},
}

type legacyCalendarOp struct {
	Summary  string
	DateType string
	End      string
	Start    string
}

var LegacyCalendarOp = legacyCalendarOp{
	Summary:  "Get US market calendar",
	DateType: "indicates what start and end mean",
	End:      "last date to retrieve data for (inclusive)",
	Start:    "first date to retrieve data for (inclusive)",
}

var LegacyCalendarFlags = []FlagDef{
	{Name: "date-type", OASName: "date_type", Type: "string", Description: "indicates what start and end mean", Completions: []string{"SETTLEMENT", "TRADING"}},
	{Name: "end", OASName: "end", Type: "string", Description: "last date to retrieve data for (inclusive)"},
	{Name: "start", OASName: "start", Type: "string", Description: "first date to retrieve data for (inclusive)"},
}

type legacyClockOp struct {
	Summary string
}

var LegacyClockOp = legacyClockOp{
	Summary: "Get US market clock",
}

type logosOp struct {
	Summary     string
	Placeholder string
	Symbol      string
}

var LogosOp = logosOp{
	Summary:     "Get logos",
	Placeholder: "placeholder",
	Symbol:      "A unique series of letters assigned to a security for trading purposes",
}

var LogosFlags = []FlagDef{
	{Name: "placeholder", OASName: "placeholder", Type: "bool", Default: "true", Description: "placeholder"},
}

type mostActivesOp struct {
	Summary string
	By      string
	Top     string
}

var MostActivesOp = mostActivesOp{
	Summary: "Get most active stocks",
	By:      "metric used for ranking the most active stocks",
	Top:     "number of top most active stocks to fetch per day",
}

var MostActivesFlags = []FlagDef{
	{Name: "by", OASName: "by", Type: "string", Default: "volume", Description: "metric used for ranking the most active stocks", Completions: []string{"trades", "volume"}},
	{Name: "top", OASName: "top", Type: "int", Default: "10", Description: "number of top most active stocks to fetch per day"},
}

type moversOp struct {
	Summary    string
	MarketType string
	Top        string
}

var MoversOp = moversOp{
	Summary:    "Get top market movers",
	MarketType: "screen-specific market (stocks or crypto)",
	Top:        "number of top market movers to fetch (gainers and losers)",
}

var MoversFlags = []FlagDef{
	{Name: "top", OASName: "top", Type: "int", Default: "10", Description: "number of top market movers to fetch (gainers and losers)"},
}

type newsOp struct {
	Summary            string
	End                string
	ExcludeContentless string
	IncludeContent     string
	Limit              string
	PageToken          string
	Sort               string
	Start              string
	Symbols            string
}

var NewsOp = newsOp{
	Summary:            "Get news articles",
	End:                "inclusive end of the interval",
	ExcludeContentless: "boolean indicator to exclude news articles that do not contain content",
	IncludeContent:     "boolean indicator to include content for news articles (if available)",
	Limit:              "limit of news items to be returned for a result page",
	PageToken:          "pagination token from which to continue",
	Sort:               "sort articles by updated date",
	Start:              "inclusive start of the interval",
	Symbols:            "A comma-separated list of symbols for which to query news",
}

var NewsFlags = []FlagDef{
	{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval"},
	{Name: "exclude-contentless", OASName: "exclude_contentless", Type: "bool", Description: "boolean indicator to exclude news articles that do not contain content"},
	{Name: "include-content", OASName: "include_content", Type: "bool", Description: "boolean indicator to include content for news articles (if available)"},
	{Name: "limit", OASName: "limit", Type: "int", Description: "limit of news items to be returned for a result page"},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue"},
	{Name: "sort", OASName: "sort", Type: "string", Default: "desc", Description: "sort articles by updated date", Completions: []string{"asc", "desc"}},
	{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval"},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of symbols for which to query news"},
}

type optionChainOp struct {
	Summary           string
	ExpirationDate    string
	ExpirationDateGte string
	ExpirationDateLte string
	Feed              string
	Limit             string
	PageToken         string
	RootSymbol        string
	StrikePriceGte    string
	StrikePriceLte    string
	Type              string
	UnderlyingSymbol  string
	UpdatedSince      string
}

var OptionChainOp = optionChainOp{
	Summary:           "Get option chain",
	ExpirationDate:    "filter contracts by the exact expiration date (format: YYYY-MM-DD)",
	ExpirationDateGte: "filter contracts with expiration date greater than or equal to the specified date",
	ExpirationDateLte: "filter contracts with expiration date less than or equal to the specified date",
	Feed:              "source feed of the data",
	Limit:             "number of maximum snapshots to return in a response.",
	PageToken:         "pagination token from which to continue",
	RootSymbol:        "filter contracts by the root symbol",
	StrikePriceGte:    "filter contracts with strike price greater than or equal to the specified value",
	StrikePriceLte:    "filter contracts with strike price less than or equal to the specified value",
	Type:              "filter contracts by the type (call or put)",
	UnderlyingSymbol:  "financial instrument on which an option contract is based or derived",
	UpdatedSince:      "filter to snapshots that were updated since this timestamp, meaning that the timestamp of the trade or the quote is g...",
}

var OptionChainFlags = []FlagDef{
	{Name: "expiration-date", OASName: "expiration_date", Type: "string", Description: "filter contracts by the exact expiration date (format: YYYY-MM-DD)"},
	{Name: "expiration-date-gte", OASName: "expiration_date_gte", Type: "string", Description: "filter contracts with expiration date greater than or equal to the specified date"},
	{Name: "expiration-date-lte", OASName: "expiration_date_lte", Type: "string", Description: "filter contracts with expiration date less than or equal to the specified date"},
	{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data"},
	{Name: "limit", OASName: "limit", Type: "int", Default: "100", Description: "number of maximum snapshots to return in a response."},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue"},
	{Name: "root-symbol", OASName: "root_symbol", Type: "string", Description: "filter contracts by the root symbol"},
	{Name: "strike-price-gte", OASName: "strike_price_gte", Type: "float64", Description: "filter contracts with strike price greater than or equal to the specified value"},
	{Name: "strike-price-lte", OASName: "strike_price_lte", Type: "float64", Description: "filter contracts with strike price less than or equal to the specified value"},
	{Name: "type", OASName: "type", Type: "string", Description: "filter contracts by the type (call or put)", Completions: []string{"call", "put"}},
	{Name: "updated-since", OASName: "updated_since", Type: "string", Description: "filter to snapshots that were updated since this timestamp, meaning that the timestamp of the trade or the quote is g..."},
}

type optionLatestQuotesOp struct {
	Summary string
	Feed    string
	Symbols string
}

var OptionLatestQuotesOp = optionLatestQuotesOp{
	Summary: "Get latest quotes",
	Feed:    "source feed of the data",
	Symbols: "A comma-separated list of contract symbols with a limit of 100",
}

var OptionLatestQuotesFlags = []FlagDef{
	{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data"},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100"},
}

type optionLatestTradesOp struct {
	Summary string
	Feed    string
	Symbols string
}

var OptionLatestTradesOp = optionLatestTradesOp{
	Summary: "Get latest trades",
	Feed:    "source feed of the data",
	Symbols: "A comma-separated list of contract symbols with a limit of 100",
}

var OptionLatestTradesFlags = []FlagDef{
	{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data"},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100"},
}

type optionMetaConditionsOp struct {
	Summary  string
	Ticktype string
}

var OptionMetaConditionsOp = optionMetaConditionsOp{
	Summary:  "Get condition codes",
	Ticktype: "type of ticks",
}

type optionMetaExchangesOp struct {
	Summary string
}

var OptionMetaExchangesOp = optionMetaExchangesOp{
	Summary: "Get exchange codes",
}

type optionSnapshotsOp struct {
	Summary      string
	Feed         string
	Limit        string
	PageToken    string
	Symbols      string
	UpdatedSince string
}

var OptionSnapshotsOp = optionSnapshotsOp{
	Summary:      "Get snapshots",
	Feed:         "source feed of the data",
	Limit:        "number of maximum snapshots to return in a response.",
	PageToken:    "pagination token from which to continue",
	Symbols:      "A comma-separated list of contract symbols with a limit of 100",
	UpdatedSince: "filter to snapshots that were updated since this timestamp, meaning that the timestamp of the trade or the quote is g...",
}

var OptionSnapshotsFlags = []FlagDef{
	{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data"},
	{Name: "limit", OASName: "limit", Type: "int", Default: "100", Description: "number of maximum snapshots to return in a response."},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue"},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100"},
	{Name: "updated-since", OASName: "updated_since", Type: "string", Description: "filter to snapshots that were updated since this timestamp, meaning that the timestamp of the trade or the quote is g..."},
}

type optionTradesOp struct {
	Summary   string
	End       string
	Limit     string
	PageToken string
	Sort      string
	Start     string
	Symbols   string
}

var OptionTradesOp = optionTradesOp{
	Summary:   "Get historical trades",
	End:       "inclusive end of the interval",
	Limit:     "maximum number of data points to return in the response page.",
	PageToken: "pagination token from which to continue",
	Sort:      "sort data in ascending or descending order",
	Start:     "inclusive start of the interval",
	Symbols:   "A comma-separated list of contract symbols with a limit of 100",
}

var OptionTradesFlags = []FlagDef{
	{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval"},
	{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page."},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue"},
	{Name: "sort", OASName: "sort", Type: "string", Description: "sort data in ascending or descending order"},
	{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval"},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100"},
}

type ratesOp struct {
	Summary       string
	CurrencyPairs string
	End           string
	Limit         string
	PageToken     string
	Sort          string
	Start         string
	Timeframe     string
}

var RatesOp = ratesOp{
	Summary:       "Get historical rates for currency pairs",
	CurrencyPairs: "currency pairs",
	End:           "inclusive end of the interval",
	Limit:         "maximum number of data points to return in the response page.",
	PageToken:     "pagination token from which to continue",
	Sort:          "sort data in ascending or descending order",
	Start:         "inclusive start of the interval",
	Timeframe:     "timeframe",
}

var RatesFlags = []FlagDef{
	{Name: "currency-pairs", OASName: "currency_pairs", Type: "string", Description: "currency pairs"},
	{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval"},
	{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page."},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue"},
	{Name: "sort", OASName: "sort", Type: "string", Description: "sort data in ascending or descending order"},
	{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval"},
	{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "timeframe"},
}

type stockAuctionSingleOp struct {
	Summary   string
	Asof      string
	Currency  string
	End       string
	Feed      string
	Limit     string
	PageToken string
	Sort      string
	Start     string
	Symbol    string
}

var StockAuctionSingleOp = stockAuctionSingleOp{
	Summary:   "Get historical auctions (single)",
	Asof:      "as-of date of the queried stock symbol(s)",
	Currency:  "currency of all prices in ISO 4217 format. Default: USD",
	End:       "inclusive end of the interval",
	Feed:      "only sip is valid for auctions",
	Limit:     "maximum number of data points to return in the response page.",
	PageToken: "pagination token from which to continue",
	Sort:      "sort data in ascending or descending order",
	Start:     "inclusive start of the interval",
	Symbol:    "symbol to query",
}

var StockAuctionSingleFlags = []FlagDef{
	{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)"},
	{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD"},
	{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval"},
	{Name: "feed", OASName: "feed", Type: "string", Description: "only sip is valid for auctions"},
	{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page."},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue"},
	{Name: "sort", OASName: "sort", Type: "string", Description: "sort data in ascending or descending order"},
	{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval"},
}

type stockAuctionsOp struct {
	Summary   string
	Asof      string
	Currency  string
	End       string
	Feed      string
	Limit     string
	PageToken string
	Sort      string
	Start     string
	Symbols   string
}

var StockAuctionsOp = stockAuctionsOp{
	Summary:   "Get historical auctions",
	Asof:      "as-of date of the queried stock symbol(s)",
	Currency:  "currency of all prices in ISO 4217 format. Default: USD",
	End:       "inclusive end of the interval",
	Feed:      "only sip is valid for auctions",
	Limit:     "maximum number of data points to return in the response page.",
	PageToken: "pagination token from which to continue",
	Sort:      "sort data in ascending or descending order",
	Start:     "inclusive start of the interval",
	Symbols:   "A comma-separated list of stock symbols",
}

var StockAuctionsFlags = []FlagDef{
	{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)"},
	{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD"},
	{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval"},
	{Name: "feed", OASName: "feed", Type: "string", Description: "only sip is valid for auctions"},
	{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page."},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue"},
	{Name: "sort", OASName: "sort", Type: "string", Description: "sort data in ascending or descending order"},
	{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval"},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols"},
}

type stockBarSingleOp struct {
	Summary    string
	Adjustment string
	Asof       string
	Currency   string
	End        string
	Feed       string
	Limit      string
	PageToken  string
	Sort       string
	Start      string
	Symbol     string
	Timeframe  string
}

var StockBarSingleOp = stockBarSingleOp{
	Summary:    "Get historical bars (single symbol)",
	Adjustment: "specifies the adjustments for the bars.\n\n - raw: no adjustments\n - split: adjust price and volume for forward and rev...",
	Asof:       "as-of date of the queried stock symbol(s)",
	Currency:   "currency of all prices in ISO 4217 format. Default: USD",
	End:        "inclusive end of the interval",
	Feed:       "source feed of the data.",
	Limit:      "maximum number of data points to return in the response page.",
	PageToken:  "pagination token from which to continue",
	Sort:       "sort data in ascending or descending order",
	Start:      "inclusive start of the interval",
	Symbol:     "symbol to query",
	Timeframe:  "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g",
}

var StockBarSingleFlags = []FlagDef{
	{Name: "adjustment", OASName: "adjustment", Type: "string", Default: "raw", Description: "specifies the adjustments for the bars.\n\n - raw: no adjustments\n - split: adjust price and volume for forward and rev..."},
	{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)"},
	{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD"},
	{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval"},
	{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data."},
	{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page."},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue"},
	{Name: "sort", OASName: "sort", Type: "string", Description: "sort data in ascending or descending order"},
	{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval"},
	{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g"},
}

type stockBarsOp struct {
	Summary    string
	Adjustment string
	Asof       string
	Currency   string
	End        string
	Feed       string
	Limit      string
	PageToken  string
	Sort       string
	Start      string
	Symbols    string
	Timeframe  string
}

var StockBarsOp = stockBarsOp{
	Summary:    "Get historical bars",
	Adjustment: "specifies the adjustments for the bars.\n\n - raw: no adjustments\n - split: adjust price and volume for forward and rev...",
	Asof:       "as-of date of the queried stock symbol(s)",
	Currency:   "currency of all prices in ISO 4217 format. Default: USD",
	End:        "inclusive end of the interval",
	Feed:       "source feed of the data.",
	Limit:      "maximum number of data points to return in the response page.",
	PageToken:  "pagination token from which to continue",
	Sort:       "sort data in ascending or descending order",
	Start:      "inclusive start of the interval",
	Symbols:    "A comma-separated list of stock symbols",
	Timeframe:  "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g",
}

var StockBarsFlags = []FlagDef{
	{Name: "adjustment", OASName: "adjustment", Type: "string", Default: "raw", Description: "specifies the adjustments for the bars.\n\n - raw: no adjustments\n - split: adjust price and volume for forward and rev..."},
	{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)"},
	{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD"},
	{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval"},
	{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data."},
	{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page."},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue"},
	{Name: "sort", OASName: "sort", Type: "string", Description: "sort data in ascending or descending order"},
	{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval"},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols"},
	{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g"},
}

type stockLatestBarSingleOp struct {
	Summary  string
	Currency string
	Feed     string
	Symbol   string
}

var StockLatestBarSingleOp = stockLatestBarSingleOp{
	Summary:  "Get latest bar (single symbol)",
	Currency: "currency of all prices in ISO 4217 format. Default: USD",
	Feed:     "source feed of the data.",
	Symbol:   "symbol to query",
}

var StockLatestBarSingleFlags = []FlagDef{
	{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD"},
	{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data."},
}

type stockLatestBarsOp struct {
	Summary  string
	Currency string
	Feed     string
	Symbols  string
}

var StockLatestBarsOp = stockLatestBarsOp{
	Summary:  "Get latest bars",
	Currency: "currency of all prices in ISO 4217 format. Default: USD",
	Feed:     "source feed of the data.",
	Symbols:  "A comma-separated list of stock symbols",
}

var StockLatestBarsFlags = []FlagDef{
	{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD"},
	{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data."},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols"},
}

type stockLatestQuoteSingleOp struct {
	Summary  string
	Currency string
	Feed     string
	Symbol   string
}

var StockLatestQuoteSingleOp = stockLatestQuoteSingleOp{
	Summary:  "Get latest quote (single symbol)",
	Currency: "currency of all prices in ISO 4217 format. Default: USD",
	Feed:     "source feed of the data.",
	Symbol:   "symbol to query",
}

var StockLatestQuoteSingleFlags = []FlagDef{
	{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD"},
	{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data."},
}

type stockLatestQuotesOp struct {
	Summary  string
	Currency string
	Feed     string
	Symbols  string
}

var StockLatestQuotesOp = stockLatestQuotesOp{
	Summary:  "Get latest quotes",
	Currency: "currency of all prices in ISO 4217 format. Default: USD",
	Feed:     "source feed of the data.",
	Symbols:  "A comma-separated list of stock symbols",
}

var StockLatestQuotesFlags = []FlagDef{
	{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD"},
	{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data."},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols"},
}

type stockLatestTradeSingleOp struct {
	Summary  string
	Currency string
	Feed     string
	Symbol   string
}

var StockLatestTradeSingleOp = stockLatestTradeSingleOp{
	Summary:  "Get latest trade (single symbol)",
	Currency: "currency of all prices in ISO 4217 format. Default: USD",
	Feed:     "source feed of the data.",
	Symbol:   "symbol to query",
}

var StockLatestTradeSingleFlags = []FlagDef{
	{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD"},
	{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data."},
}

type stockLatestTradesOp struct {
	Summary  string
	Currency string
	Feed     string
	Symbols  string
}

var StockLatestTradesOp = stockLatestTradesOp{
	Summary:  "Get latest trades",
	Currency: "currency of all prices in ISO 4217 format. Default: USD",
	Feed:     "source feed of the data.",
	Symbols:  "A comma-separated list of stock symbols",
}

var StockLatestTradesFlags = []FlagDef{
	{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD"},
	{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data."},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols"},
}

type stockMetaConditionsOp struct {
	Summary  string
	Tape     string
	Ticktype string
}

var StockMetaConditionsOp = stockMetaConditionsOp{
	Summary:  "Get condition codes",
	Tape:     "one character name of the tape",
	Ticktype: "type of ticks",
}

var StockMetaConditionsFlags = []FlagDef{
	{Name: "tape", OASName: "tape", Type: "string", Description: "one character name of the tape", Completions: []string{"A", "B", "C"}},
}

type stockMetaExchangesOp struct {
	Summary string
}

var StockMetaExchangesOp = stockMetaExchangesOp{
	Summary: "Get exchange codes",
}

type stockQuoteSingleOp struct {
	Summary   string
	Asof      string
	Currency  string
	End       string
	Feed      string
	Limit     string
	PageToken string
	Sort      string
	Start     string
	Symbol    string
}

var StockQuoteSingleOp = stockQuoteSingleOp{
	Summary:   "Get historical quotes (single symbol)",
	Asof:      "as-of date of the queried stock symbol(s)",
	Currency:  "currency of all prices in ISO 4217 format. Default: USD",
	End:       "inclusive end of the interval",
	Feed:      "source feed of the data.",
	Limit:     "maximum number of data points to return in the response page.",
	PageToken: "pagination token from which to continue",
	Sort:      "sort data in ascending or descending order",
	Start:     "inclusive start of the interval",
	Symbol:    "symbol to query",
}

var StockQuoteSingleFlags = []FlagDef{
	{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)"},
	{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD"},
	{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval"},
	{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data."},
	{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page."},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue"},
	{Name: "sort", OASName: "sort", Type: "string", Description: "sort data in ascending or descending order"},
	{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval"},
}

type stockQuotesOp struct {
	Summary   string
	Asof      string
	Currency  string
	End       string
	Feed      string
	Limit     string
	PageToken string
	Sort      string
	Start     string
	Symbols   string
}

var StockQuotesOp = stockQuotesOp{
	Summary:   "Get historical quotes",
	Asof:      "as-of date of the queried stock symbol(s)",
	Currency:  "currency of all prices in ISO 4217 format. Default: USD",
	End:       "inclusive end of the interval",
	Feed:      "source feed of the data.",
	Limit:     "maximum number of data points to return in the response page.",
	PageToken: "pagination token from which to continue",
	Sort:      "sort data in ascending or descending order",
	Start:     "inclusive start of the interval",
	Symbols:   "A comma-separated list of stock symbols",
}

var StockQuotesFlags = []FlagDef{
	{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)"},
	{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD"},
	{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval"},
	{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data."},
	{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page."},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue"},
	{Name: "sort", OASName: "sort", Type: "string", Description: "sort data in ascending or descending order"},
	{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval"},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols"},
}

type stockSnapshotSingleOp struct {
	Summary  string
	Currency string
	Feed     string
	Symbol   string
}

var StockSnapshotSingleOp = stockSnapshotSingleOp{
	Summary:  "Get snapshot (single symbol)",
	Currency: "currency of all prices in ISO 4217 format. Default: USD",
	Feed:     "source feed of the data.",
	Symbol:   "symbol to query",
}

var StockSnapshotSingleFlags = []FlagDef{
	{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD"},
	{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data."},
}

type stockSnapshotsOp struct {
	Summary  string
	Currency string
	Feed     string
	Symbols  string
}

var StockSnapshotsOp = stockSnapshotsOp{
	Summary:  "Get snapshots",
	Currency: "currency of all prices in ISO 4217 format. Default: USD",
	Feed:     "source feed of the data.",
	Symbols:  "A comma-separated list of stock symbols",
}

var StockSnapshotsFlags = []FlagDef{
	{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD"},
	{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data."},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols"},
}

type stockTradeSingleOp struct {
	Summary   string
	Asof      string
	Currency  string
	End       string
	Feed      string
	Limit     string
	PageToken string
	Sort      string
	Start     string
	Symbol    string
}

var StockTradeSingleOp = stockTradeSingleOp{
	Summary:   "Get historical trades (single symbol)",
	Asof:      "as-of date of the queried stock symbol(s)",
	Currency:  "currency of all prices in ISO 4217 format. Default: USD",
	End:       "inclusive end of the interval",
	Feed:      "source feed of the data.",
	Limit:     "maximum number of data points to return in the response page.",
	PageToken: "pagination token from which to continue",
	Sort:      "sort data in ascending or descending order",
	Start:     "inclusive start of the interval",
	Symbol:    "symbol to query",
}

var StockTradeSingleFlags = []FlagDef{
	{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)"},
	{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD"},
	{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval"},
	{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data."},
	{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page."},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue"},
	{Name: "sort", OASName: "sort", Type: "string", Description: "sort data in ascending or descending order"},
	{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval"},
}

type stockTradesOp struct {
	Summary   string
	Asof      string
	Currency  string
	End       string
	Feed      string
	Limit     string
	PageToken string
	Sort      string
	Start     string
	Symbols   string
}

var StockTradesOp = stockTradesOp{
	Summary:   "Get historical trades",
	Asof:      "as-of date of the queried stock symbol(s)",
	Currency:  "currency of all prices in ISO 4217 format. Default: USD",
	End:       "inclusive end of the interval",
	Feed:      "source feed of the data.",
	Limit:     "maximum number of data points to return in the response page.",
	PageToken: "pagination token from which to continue",
	Sort:      "sort data in ascending or descending order",
	Start:     "inclusive start of the interval",
	Symbols:   "A comma-separated list of stock symbols",
}

var StockTradesFlags = []FlagDef{
	{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)"},
	{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD"},
	{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval"},
	{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data."},
	{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page."},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue"},
	{Name: "sort", OASName: "sort", Type: "string", Description: "sort data in ascending or descending order"},
	{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval"},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols"},
}

type usCorporatesOp struct {
	Summary    string
	BondStatus string
	Cusips     string
	Isins      string
	Tickers    string
}

var UsCorporatesOp = usCorporatesOp{
	Summary:    "Get US corporates",
	BondStatus: "bond status",
	Cusips:     "A comma-separated list of CUSIPs with a limit of 1000",
	Isins:      "A comma-separated list of ISINs with a limit of 1000",
	Tickers:    "A comma-separated list of tickers with a limit of 1000",
}

var UsCorporatesFlags = []FlagDef{
	{Name: "bond-status", OASName: "bond_status", Type: "string", Description: "bond status"},
	{Name: "cusips", OASName: "cusips", Type: "string", Description: "A comma-separated list of CUSIPs with a limit of 1000"},
	{Name: "isins", OASName: "isins", Type: "string", Description: "A comma-separated list of ISINs with a limit of 1000"},
	{Name: "tickers", OASName: "tickers", Type: "string", Description: "A comma-separated list of tickers with a limit of 1000"},
}

type usTreasuriesOp struct {
	Summary    string
	BondStatus string
	Cusips     string
	Isins      string
	Subtype    string
}

var UsTreasuriesOp = usTreasuriesOp{
	Summary:    "Get US treasuries",
	BondStatus: "bond status",
	Cusips:     "A comma-separated list of CUSIPs with a limit of 1000",
	Isins:      "A comma-separated list of ISINs with a limit of 1000",
	Subtype:    "subtype",
}

var UsTreasuriesFlags = []FlagDef{
	{Name: "bond-status", OASName: "bond_status", Type: "string", Description: "bond status"},
	{Name: "cusips", OASName: "cusips", Type: "string", Description: "A comma-separated list of CUSIPs with a limit of 1000"},
	{Name: "isins", OASName: "isins", Type: "string", Description: "A comma-separated list of ISINs with a limit of 1000"},
	{Name: "subtype", OASName: "subtype", Type: "string", Description: "subtype"},
}

type addAssetToWatchlistOp struct {
	Summary     string
	Symbol      string
	WatchlistID string
}

var AddAssetToWatchlistOp = addAssetToWatchlistOp{
	Summary:     "Add asset to watchlist",
	Symbol:      "the symbol name to add to the watchlist",
	WatchlistID: "watchlist id",
}

var AddAssetToWatchlistFlags = []FlagDef{
	{Name: "symbol", OASName: "symbol", Type: "string", Description: "the symbol name to add to the watchlist"},
}

type addAssetToWatchlistByNameOp struct {
	Summary string
	Name    string
	Symbol  string
}

var AddAssetToWatchlistByNameOp = addAssetToWatchlistByNameOp{
	Summary: "Add asset to watchlist by name",
	Name:    "name of the watchlist",
	Symbol:  "the symbol name to add to the watchlist",
}

var AddAssetToWatchlistByNameFlags = []FlagDef{
	{Name: "name", OASName: "name", Type: "string", Description: "name of the watchlist"},
	{Name: "symbol", OASName: "symbol", Type: "string", Description: "the symbol name to add to the watchlist"},
}

type createCryptoPerpTransferForAccountOp struct {
	Summary string
	Address string
	Amount  string
	Asset   string
}

var CreateCryptoPerpTransferForAccountOp = createCryptoPerpTransferForAccountOp{
	Summary: "Request a new withdrawal",
	Address: "destination wallet address",
	Amount:  "amount, denoted in the specified asset, to be withdrawn from the user’s wallet",
	Asset:   "asset",
}

var CreateCryptoPerpTransferForAccountFlags = []FlagDef{
	{Name: "address", OASName: "address", Type: "string", Description: "destination wallet address"},
	{Name: "amount", OASName: "amount", Type: "string", Description: "amount, denoted in the specified asset, to be withdrawn from the user’s wallet"},
	{Name: "asset", OASName: "asset", Type: "string", Description: "asset"},
}

type createCryptoTransferForAccountOp struct {
	Summary string
	Address string
	Amount  string
	Asset   string
}

var CreateCryptoTransferForAccountOp = createCryptoTransferForAccountOp{
	Summary: "Request a new withdrawal",
	Address: "destination wallet address",
	Amount:  "amount, denoted in the specified asset, to be withdrawn from the user’s wallet",
	Asset:   "asset",
}

var CreateCryptoTransferForAccountFlags = []FlagDef{
	{Name: "address", OASName: "address", Type: "string", Description: "destination wallet address"},
	{Name: "amount", OASName: "amount", Type: "string", Description: "amount, denoted in the specified asset, to be withdrawn from the user’s wallet"},
	{Name: "asset", OASName: "asset", Type: "string", Description: "asset"},
}

type createWhitelistedAddressOp struct {
	Summary string
	Address string
	Asset   string
}

var CreateWhitelistedAddressOp = createWhitelistedAddressOp{
	Summary: "Request a new whitelisted address",
	Address: "address to be whitelisted",
	Asset:   "symbol of underlying asset for the whitelisted address",
}

var CreateWhitelistedAddressFlags = []FlagDef{
	{Name: "address", OASName: "address", Type: "string", Description: "address to be whitelisted"},
	{Name: "asset", OASName: "asset", Type: "string", Description: "symbol of underlying asset for the whitelisted address"},
}

type createWhitelistedPerpAddressOp struct {
	Summary string
	Address string
	Asset   string
}

var CreateWhitelistedPerpAddressOp = createWhitelistedPerpAddressOp{
	Summary: "Request a new whitelisted address",
	Address: "address to be whitelisted",
	Asset:   "symbol of underlying asset for the whitelisted address",
}

var CreateWhitelistedPerpAddressFlags = []FlagDef{
	{Name: "address", OASName: "address", Type: "string", Description: "address to be whitelisted"},
	{Name: "asset", OASName: "asset", Type: "string", Description: "symbol of underlying asset for the whitelisted address"},
}

type deleteAllOpenPositionsOp struct {
	Summary      string
	CancelOrders string
}

var DeleteAllOpenPositionsOp = deleteAllOpenPositionsOp{
	Summary:      "Close all positions",
	CancelOrders: "if true is specified, cancel all open orders before liquidating all positions",
}

var DeleteAllOpenPositionsFlags = []FlagDef{
	{Name: "cancel-orders", OASName: "cancel_orders", Type: "bool", Description: "if true is specified, cancel all open orders before liquidating all positions"},
}

type deleteAllOrdersOp struct {
	Summary string
}

var DeleteAllOrdersOp = deleteAllOrdersOp{
	Summary: "Delete all orders",
}

type deleteOpenPositionOp struct {
	Summary         string
	Percentage      string
	Qty             string
	SymbolOrAssetID string
}

var DeleteOpenPositionOp = deleteOpenPositionOp{
	Summary:         "Close a position",
	Percentage:      "percentage of position to liquidate",
	Qty:             "the number of shares to liquidate. Can accept up to 9 decimal points. Cannot work with percentage",
	SymbolOrAssetID: "symbol or assetId",
}

var DeleteOpenPositionFlags = []FlagDef{
	{Name: "percentage", OASName: "percentage", Type: "float64", Description: "percentage of position to liquidate"},
	{Name: "qty", OASName: "qty", Type: "float64", Description: "the number of shares to liquidate. Can accept up to 9 decimal points. Cannot work with percentage"},
}

type deleteOrderByOrderIDOp struct {
	Summary string
	OrderID string
}

var DeleteOrderByOrderIDOp = deleteOrderByOrderIDOp{
	Summary: "Delete order by ID",
	OrderID: "order id",
}

type deleteWatchlistByIDOp struct {
	Summary     string
	WatchlistID string
}

var DeleteWatchlistByIDOp = deleteWatchlistByIDOp{
	Summary:     "Delete watchlist by id",
	WatchlistID: "watchlist id",
}

type deleteWatchlistByNameOp struct {
	Summary string
	Name    string
}

var DeleteWatchlistByNameOp = deleteWatchlistByNameOp{
	Summary: "Delete watchlist by name",
	Name:    "name of the watchlist",
}

var DeleteWatchlistByNameFlags = []FlagDef{
	{Name: "name", OASName: "name", Type: "string", Description: "name of the watchlist"},
}

type deleteWhitelistedAddressOp struct {
	Summary              string
	WhitelistedAddressID string
}

var DeleteWhitelistedAddressOp = deleteWhitelistedAddressOp{
	Summary:              "Delete a whitelisted address",
	WhitelistedAddressID: "whitelisted address to delete",
}

type deleteWhitelistedPerpAddressOp struct {
	Summary              string
	WhitelistedAddressID string
}

var DeleteWhitelistedPerpAddressOp = deleteWhitelistedPerpAddressOp{
	Summary:              "Delete a whitelisted address",
	WhitelistedAddressID: "whitelisted address to delete",
}

type getOptionContractSymbolOrIDOp struct {
	Summary    string
	SymbolOrID string
}

var GetOptionContractSymbolOrIDOp = getOptionContractSymbolOrIDOp{
	Summary:    "Get an option contract by ID or symbol",
	SymbolOrID: "symbol or contract ID",
}

type getOptionsContractsOp struct {
	Summary           string
	ExpirationDate    string
	ExpirationDateGte string
	ExpirationDateLte string
	Limit             string
	PageToken         string
	Ppind             string
	RootSymbol        string
	ShowDeliverables  string
	Status            string
	StrikePriceGte    string
	StrikePriceLte    string
	Style             string
	Type              string
	UnderlyingSymbols string
}

var GetOptionsContractsOp = getOptionsContractsOp{
	Summary:           "Get option contracts",
	ExpirationDate:    "filter contracts by the exact expiration date (format: YYYY-MM-DD)",
	ExpirationDateGte: "filter contracts with expiration date greater than or equal to the specified date",
	ExpirationDateLte: "filter contracts with expiration date less than or equal to the specified date",
	Limit:             "number of contracts to limit per page (default=100, max=10000)",
	PageToken:         "used for pagination, this token retrieves the next page of results",
	Ppind:             "ppind(Penny Program Indicator) field indicates whether an option contract is eligible for penny price increments,",
	RootSymbol:        "filter contracts by the root symbol",
	ShowDeliverables:  "include deliverables array in the response",
	Status:            "filter contracts by status (active/inactive). By default only active contracts are returned",
	StrikePriceGte:    "filter contracts with strike price greater than or equal to the specified value",
	StrikePriceLte:    "filter contracts with strike price less than or equal to the specified value",
	Style:             "filter contracts by the style (american/european)",
	Type:              "filter contracts by the type (call/put)",
	UnderlyingSymbols: "filter contracts by one or more underlying symbols",
}

var GetOptionsContractsFlags = []FlagDef{
	{Name: "expiration-date", OASName: "expiration_date", Type: "string", Description: "filter contracts by the exact expiration date (format: YYYY-MM-DD)"},
	{Name: "expiration-date-gte", OASName: "expiration_date_gte", Type: "string", Description: "filter contracts with expiration date greater than or equal to the specified date"},
	{Name: "expiration-date-lte", OASName: "expiration_date_lte", Type: "string", Description: "filter contracts with expiration date less than or equal to the specified date"},
	{Name: "limit", OASName: "limit", Type: "int", Description: "number of contracts to limit per page (default=100, max=10000)"},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "used for pagination, this token retrieves the next page of results"},
	{Name: "ppind", OASName: "ppind", Type: "bool", Description: "ppind(Penny Program Indicator) field indicates whether an option contract is eligible for penny price increments,"},
	{Name: "root-symbol", OASName: "root_symbol", Type: "string", Description: "filter contracts by the root symbol"},
	{Name: "show-deliverables", OASName: "show_deliverables", Type: "bool", Description: "include deliverables array in the response"},
	{Name: "status", OASName: "status", Type: "string", Description: "filter contracts by status (active/inactive). By default only active contracts are returned", Completions: []string{"active", "inactive"}},
	{Name: "strike-price-gte", OASName: "strike_price_gte", Type: "float64", Description: "filter contracts with strike price greater than or equal to the specified value"},
	{Name: "strike-price-lte", OASName: "strike_price_lte", Type: "float64", Description: "filter contracts with strike price less than or equal to the specified value"},
	{Name: "style", OASName: "style", Type: "string", Description: "filter contracts by the style (american/european)", Completions: []string{"american", "european"}},
	{Name: "type", OASName: "type", Type: "string", Description: "filter contracts by the type (call/put)", Completions: []string{"call", "put"}},
	{Name: "underlying-symbols", OASName: "underlying_symbols", Type: "string", Description: "filter contracts by one or more underlying symbols"},
}

type getV2AssetsOp struct {
	Summary    string
	AssetClass string
	Attributes string
	Exchange   string
	Status     string
}

var GetV2AssetsOp = getV2AssetsOp{
	Summary:    "Get assets",
	AssetClass: "defaults to us_equity",
	Attributes: "comma separated values to query for more than one attribute",
	Exchange:   "optional AMEX, ARCA, BATS, NYSE, NASDAQ, NYSEARCA or OTC",
	Status:     "e.g. “active”. By default, all statuses are included",
}

var GetV2AssetsFlags = []FlagDef{
	{Name: "asset-class", OASName: "asset_class", Type: "string", Description: "defaults to us_equity"},
	{Name: "attributes", OASName: "attributes", Type: "string", Default: "[]", Description: "comma separated values to query for more than one attribute"},
	{Name: "exchange", OASName: "exchange", Type: "string", Description: "optional AMEX, ARCA, BATS, NYSE, NASDAQ, NYSEARCA or OTC"},
	{Name: "status", OASName: "status", Type: "string", Description: "e.g. “active”. By default, all statuses are included"},
}

type getV2AssetsSymbolOrAssetIDOp struct {
	Summary         string
	SymbolOrAssetID string
}

var GetV2AssetsSymbolOrAssetIDOp = getV2AssetsSymbolOrAssetIDOp{
	Summary:         "Get an asset by ID or symbol",
	SymbolOrAssetID: "symbol or assetId. CUSIP is also accepted for US equities",
}

type getV2CorporateActionsAnnouncementsOp struct {
	Summary  string
	CaTypes  string
	Cusip    string
	DateType string
	Since    string
	Symbol   string
	Until    string
}

var GetV2CorporateActionsAnnouncementsOp = getV2CorporateActionsAnnouncementsOp{
	Summary:  "Retrieve announcements",
	CaTypes:  "A comma-delimited list of Dividend, Merger, Spinoff, or Split",
	Cusip:    "CUSIP of the company initiating the announcement",
	DateType: "declaration_date, ex_date, record_date, or payable_date",
	Since:    "start (inclusive) of the date range when searching corporate action announcements",
	Symbol:   "symbol of the company initiating the announcement",
	Until:    "end (inclusive) of the date range when searching corporate action announcements",
}

var GetV2CorporateActionsAnnouncementsFlags = []FlagDef{
	{Name: "ca-types", OASName: "ca_types", Type: "string", Description: "A comma-delimited list of Dividend, Merger, Spinoff, or Split"},
	{Name: "cusip", OASName: "cusip", Type: "string", Description: "CUSIP of the company initiating the announcement"},
	{Name: "date-type", OASName: "date_type", Type: "string", Description: "declaration_date, ex_date, record_date, or payable_date"},
	{Name: "since", OASName: "since", Type: "string", Description: "start (inclusive) of the date range when searching corporate action announcements"},
	{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol of the company initiating the announcement"},
	{Name: "until", OASName: "until", Type: "string", Description: "end (inclusive) of the date range when searching corporate action announcements"},
}

type getV2CorporateActionsAnnouncementsIDOp struct {
	Summary string
	ID      string
}

var GetV2CorporateActionsAnnouncementsIDOp = getV2CorporateActionsAnnouncementsIDOp{
	Summary: "Retrieve a specific announcement",
	ID:      "corporate announcement’s id",
}

type getAccountOp struct {
	Summary string
}

var GetAccountOp = getAccountOp{
	Summary: "Get account",
}

type getAccountActivitiesOp struct {
	Summary       string
	ActivityTypes string
	After         string
	Category      string
	Date          string
	Direction     string
	PageSize      string
	PageToken     string
	Until         string
}

var GetAccountActivitiesOp = getAccountActivitiesOp{
	Summary:       "Retrieve account activities",
	ActivityTypes: "A comma-separated list of activity types used to filter the results",
	After:         "get activities created after this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported",
	Category:      "activity category. Cannot be used with \"activity_types\" parameter",
	Date:          "filter activities by the activity date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported",
	Direction:     "chronological order of response based on the activity datetime",
	PageSize:      "maximum number of entries to return in the response",
	PageToken:     "token used for pagination. Provide the ID of the last activity from the last page to retrieve the next set of results",
	Until:         "get activities created before this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported",
}

var GetAccountActivitiesFlags = []FlagDef{
	{Name: "activity-types", OASName: "activity_types", Type: "string", Description: "A comma-separated list of activity types used to filter the results"},
	{Name: "after", OASName: "after", Type: "string", Description: "get activities created after this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported"},
	{Name: "category", OASName: "category", Type: "string", Description: "activity category. Cannot be used with \"activity_types\" parameter", Completions: []string{"non_trade_activity", "trade_activity"}},
	{Name: "date", OASName: "date", Type: "string", Description: "filter activities by the activity date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported"},
	{Name: "direction", OASName: "direction", Type: "string", Default: "desc", Description: "chronological order of response based on the activity datetime", Completions: []string{"asc", "desc"}},
	{Name: "page-size", OASName: "page_size", Type: "int", Default: "100", Description: "maximum number of entries to return in the response"},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "token used for pagination. Provide the ID of the last activity from the last page to retrieve the next set of results"},
	{Name: "until", OASName: "until", Type: "string", Description: "get activities created before this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported"},
}

type getAccountActivitiesByActivityTypeOp struct {
	Summary      string
	ActivityType string
	After        string
	Date         string
	Direction    string
	PageSize     string
	PageToken    string
	Until        string
}

var GetAccountActivitiesByActivityTypeOp = getAccountActivitiesByActivityTypeOp{
	Summary:      "Retrieve account activities of specific type",
	ActivityType: "activity type you want to view entries for. A list of valid activity types can be found at the bottom of this page",
	After:        "get activities created after this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported",
	Date:         "filter activities by the activity date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported",
	Direction:    "chronological order of response based on the activity datetime",
	PageSize:     "maximum number of entries to return in the response",
	PageToken:    "token used for pagination. Provide the ID of the last activity from the last page to retrieve the next set of results",
	Until:        "get activities created before this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported",
}

var GetAccountActivitiesByActivityTypeFlags = []FlagDef{
	{Name: "after", OASName: "after", Type: "string", Description: "get activities created after this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported"},
	{Name: "date", OASName: "date", Type: "string", Description: "filter activities by the activity date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported"},
	{Name: "direction", OASName: "direction", Type: "string", Default: "desc", Description: "chronological order of response based on the activity datetime", Completions: []string{"asc", "desc"}},
	{Name: "page-size", OASName: "page_size", Type: "int", Default: "100", Description: "maximum number of entries to return in the response"},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "token used for pagination. Provide the ID of the last activity from the last page to retrieve the next set of results"},
	{Name: "until", OASName: "until", Type: "string", Description: "get activities created before this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported"},
}

type getAccountConfigOp struct {
	Summary string
}

var GetAccountConfigOp = getAccountConfigOp{
	Summary: "Get account configurations",
}

type getAccountPortfolioHistoryOp struct {
	Summary           string
	CashflowTypes     string
	End               string
	ExtendedHours     string
	IntradayReporting string
	Period            string
	PNLReset          string
	Start             string
	Timeframe         string
}

var GetAccountPortfolioHistoryOp = getAccountPortfolioHistoryOp{
	Summary:           "Get account portfolio history",
	CashflowTypes:     "cashflow activities to include in the report. One of 'ALL', 'NONE', or a comma-separated list of activity types",
	End:               "timestamp the data is returned up to in RFC3339 format (including timezone specification)",
	ExtendedHours:     "**deprecated**: Users are strongly advised to **rely on the intraday_reporting query parameter** for better control\no...",
	IntradayReporting: "for intraday resolutions (<1D) this specfies which timestamps to return data points for:\n\nAllowed values are:\n- **mar...",
	Period:            "duration of the data in number + unit format, such as 1D, where unit can be D for day, W for week, M for month and A ...",
	PNLReset:          "pnl_reset defines how we are calculating the baseline values for Profit And Loss (pnl) for queries with timeframe les...",
	Start:             "timestamp the data is returned starting from in RFC3339 format (including timezone specification)",
	Timeframe:         "resolution of time window",
}

var GetAccountPortfolioHistoryFlags = []FlagDef{
	{Name: "cashflow-types", OASName: "cashflow_types", Type: "string", Description: "cashflow activities to include in the report. One of 'ALL', 'NONE', or a comma-separated list of activity types"},
	{Name: "end", OASName: "end", Type: "string", Description: "timestamp the data is returned up to in RFC3339 format (including timezone specification)"},
	{Name: "extended-hours", OASName: "extended_hours", Type: "string", Description: "**deprecated**: Users are strongly advised to **rely on the intraday_reporting query parameter** for better control\no..."},
	{Name: "intraday-reporting", OASName: "intraday_reporting", Type: "string", Default: "market_hours", Description: "for intraday resolutions (<1D) this specfies which timestamps to return data points for:\n\nAllowed values are:\n- **mar...", Completions: []string{"continuous", "extended_hours", "market_hours"}},
	{Name: "period", OASName: "period", Type: "string", Description: "duration of the data in number + unit format, such as 1D, where unit can be D for day, W for week, M for month and A ..."},
	{Name: "pnl-reset", OASName: "pnl_reset", Type: "string", Default: "per_day", Description: "pnl_reset defines how we are calculating the baseline values for Profit And Loss (pnl) for queries with timeframe les...", Completions: []string{"no_reset", "per_day"}},
	{Name: "start", OASName: "start", Type: "string", Description: "timestamp the data is returned starting from in RFC3339 format (including timezone specification)"},
	{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "resolution of time window"},
}

type getAllOpenPositionsOp struct {
	Summary string
}

var GetAllOpenPositionsOp = getAllOpenPositionsOp{
	Summary: "List all open positions",
}

type getAllOrdersOp struct {
	Summary       string
	After         string
	AfterOrderID  string
	AssetClass    string
	BeforeOrderID string
	Direction     string
	Limit         string
	Nested        string
	Side          string
	Status        string
	Symbols       string
	Until         string
}

var GetAllOrdersOp = getAllOrdersOp{
	Summary:       "Get all orders",
	After:         "response will include only ones submitted after this timestamp (exclusive.)",
	AfterOrderID:  "return orders submitted after the order with this ID (exclusive).\nMutually exclusive with before_order_id",
	AssetClass:    "A comma seperated list of asset classes, the response will include only orders in the specified asset classes",
	BeforeOrderID: "return orders submitted before the order with this ID (exclusive).\nMutually exclusive with after_order_id",
	Direction:     "chronological order of response based on the submission time. asc or desc. Defaults to desc",
	Limit:         "maximum number of orders in response. Defaults to 50 and max is 500",
	Nested:        "if true, the result will roll up multi-leg orders under the legs field of primary order",
	Side:          "filters down to orders that have a matching side field set",
	Status:        "order status to be queried. open, closed or all. Defaults to open",
	Symbols:       "A comma-separated list of symbols to filter by (ex",
	Until:         "response will include only ones submitted until this timestamp (exclusive.)",
}

var GetAllOrdersFlags = []FlagDef{
	{Name: "after", OASName: "after", Type: "string", Description: "response will include only ones submitted after this timestamp (exclusive.)"},
	{Name: "after-order-id", OASName: "after_order_id", Type: "string", Description: "return orders submitted after the order with this ID (exclusive).\nMutually exclusive with before_order_id"},
	{Name: "asset-class", OASName: "asset_class", Type: "string", Description: "A comma seperated list of asset classes, the response will include only orders in the specified asset classes"},
	{Name: "before-order-id", OASName: "before_order_id", Type: "string", Description: "return orders submitted before the order with this ID (exclusive).\nMutually exclusive with after_order_id"},
	{Name: "direction", OASName: "direction", Type: "string", Description: "chronological order of response based on the submission time. asc or desc. Defaults to desc", Completions: []string{"asc", "desc"}},
	{Name: "limit", OASName: "limit", Type: "int", Description: "maximum number of orders in response. Defaults to 50 and max is 500"},
	{Name: "nested", OASName: "nested", Type: "bool", Description: "if true, the result will roll up multi-leg orders under the legs field of primary order"},
	{Name: "side", OASName: "side", Type: "string", Description: "filters down to orders that have a matching side field set"},
	{Name: "status", OASName: "status", Type: "string", Description: "order status to be queried. open, closed or all. Defaults to open", Completions: []string{"all", "closed", "open"}},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of symbols to filter by (ex"},
	{Name: "until", OASName: "until", Type: "string", Description: "response will include only ones submitted until this timestamp (exclusive.)"},
}

type getCryptoFundingTransferOp struct {
	Summary    string
	TransferID string
}

var GetCryptoFundingTransferOp = getCryptoFundingTransferOp{
	Summary:    "Retrieve a crypto funding transfer",
	TransferID: "crypto transfer to retrieve",
}

type getCryptoPerpAccountLeverageOp struct {
	Summary string
	Symbol  string
}

var GetCryptoPerpAccountLeverageOp = getCryptoPerpAccountLeverageOp{
	Summary: "Get account leverage for an asset",
	Symbol:  "symbol of underlying asset",
}

var GetCryptoPerpAccountLeverageFlags = []FlagDef{
	{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol of underlying asset"},
}

type getCryptoPerpAccountVitalsOp struct {
	Summary string
}

var GetCryptoPerpAccountVitalsOp = getCryptoPerpAccountVitalsOp{
	Summary: "Retrieve account vitals",
}

type getCryptoPerpFundingTransferOp struct {
	Summary    string
	TransferID string
}

var GetCryptoPerpFundingTransferOp = getCryptoPerpFundingTransferOp{
	Summary:    "Retrieve a crypto funding transfer",
	TransferID: "crypto transfer to retrieve",
}

type getCryptoPerpTransferEstimateOp struct {
	Summary     string
	Amount      string
	Asset       string
	FromAddress string
	ToAddress   string
}

var GetCryptoPerpTransferEstimateOp = getCryptoPerpTransferEstimateOp{
	Summary:     "Returns the estimated gas fee for a proposed transaction",
	Amount:      "amount, denoted in the specified asset, of the proposed transaction",
	Asset:       "asset for the proposed transaction",
	FromAddress: "originating address of the proposed transaction",
	ToAddress:   "destination address of the proposed transaction",
}

var GetCryptoPerpTransferEstimateFlags = []FlagDef{
	{Name: "amount", OASName: "amount", Type: "string", Description: "amount, denoted in the specified asset, of the proposed transaction"},
	{Name: "asset", OASName: "asset", Type: "string", Description: "asset for the proposed transaction"},
	{Name: "from-address", OASName: "from_address", Type: "string", Description: "originating address of the proposed transaction"},
	{Name: "to-address", OASName: "to_address", Type: "string", Description: "destination address of the proposed transaction"},
}

type getCryptoTransferEstimateOp struct {
	Summary     string
	Amount      string
	Asset       string
	FromAddress string
	ToAddress   string
}

var GetCryptoTransferEstimateOp = getCryptoTransferEstimateOp{
	Summary:     "Returns the estimated gas fee for a proposed transaction",
	Amount:      "amount, denoted in the specified asset, of the proposed transaction",
	Asset:       "asset for the proposed transaction",
	FromAddress: "originating address of the proposed transaction",
	ToAddress:   "destination address of the proposed transaction",
}

var GetCryptoTransferEstimateFlags = []FlagDef{
	{Name: "amount", OASName: "amount", Type: "string", Description: "amount, denoted in the specified asset, of the proposed transaction"},
	{Name: "asset", OASName: "asset", Type: "string", Description: "asset for the proposed transaction"},
	{Name: "from-address", OASName: "from_address", Type: "string", Description: "originating address of the proposed transaction"},
	{Name: "to-address", OASName: "to_address", Type: "string", Description: "destination address of the proposed transaction"},
}

type getOpenPositionOp struct {
	Summary         string
	SymbolOrAssetID string
}

var GetOpenPositionOp = getOpenPositionOp{
	Summary:         "Get an open position",
	SymbolOrAssetID: "symbol or assetId",
}

type getOrderByClientOrderIDOp struct {
	Summary       string
	ClientOrderID string
}

var GetOrderByClientOrderIDOp = getOrderByClientOrderIDOp{
	Summary:       "Get order by client order ID",
	ClientOrderID: "client-assigned order ID",
}

var GetOrderByClientOrderIDFlags = []FlagDef{
	{Name: "client-order-id", OASName: "client_order_id", Type: "string", Description: "client-assigned order ID"},
}

type getOrderByOrderIDOp struct {
	Summary string
	Nested  string
	OrderID string
}

var GetOrderByOrderIDOp = getOrderByOrderIDOp{
	Summary: "Get order by ID",
	Nested:  "if true, the result will roll up multi-leg orders under the legs field of primary order",
	OrderID: "order id",
}

var GetOrderByOrderIDFlags = []FlagDef{
	{Name: "nested", OASName: "nested", Type: "bool", Description: "if true, the result will roll up multi-leg orders under the legs field of primary order"},
}

type getWatchlistByIDOp struct {
	Summary     string
	WatchlistID string
}

var GetWatchlistByIDOp = getWatchlistByIDOp{
	Summary:     "Get watchlist by ID",
	WatchlistID: "watchlist id",
}

type getWatchlistByNameOp struct {
	Summary string
	Name    string
}

var GetWatchlistByNameOp = getWatchlistByNameOp{
	Summary: "Get watchlist by name",
	Name:    "name of the watchlist",
}

var GetWatchlistByNameFlags = []FlagDef{
	{Name: "name", OASName: "name", Type: "string", Description: "name of the watchlist"},
}

type getWatchlistsOp struct {
	Summary string
}

var GetWatchlistsOp = getWatchlistsOp{
	Summary: "Get all watchlists",
}

type listCryptoFundingTransfersOp struct {
	Summary string
}

var ListCryptoFundingTransfersOp = listCryptoFundingTransfersOp{
	Summary: "Retrieve crypto funding transfers",
}

type listCryptoFundingWalletsOp struct {
	Summary string
	Asset   string
	Network string
}

var ListCryptoFundingWalletsOp = listCryptoFundingWalletsOp{
	Summary: "Retrieve crypto funding wallets",
	Asset:   "asset",
	Network: "optional network identifier",
}

var ListCryptoFundingWalletsFlags = []FlagDef{
	{Name: "asset", OASName: "asset", Type: "string", Description: "asset"},
	{Name: "network", OASName: "network", Type: "string", Description: "optional network identifier", Completions: []string{"ethereum", "solana"}},
}

type listCryptoPerpFundingTransfersOp struct {
	Summary string
}

var ListCryptoPerpFundingTransfersOp = listCryptoPerpFundingTransfersOp{
	Summary: "Retrieve crypto funding transfers",
}

type listCryptoPerpFundingWalletsOp struct {
	Summary string
	Asset   string
}

var ListCryptoPerpFundingWalletsOp = listCryptoPerpFundingWalletsOp{
	Summary: "Retrieve crypto funding wallets",
	Asset:   "asset",
}

var ListCryptoPerpFundingWalletsFlags = []FlagDef{
	{Name: "asset", OASName: "asset", Type: "string", Description: "asset"},
}

type listWhitelistedAddressOp struct {
	Summary string
}

var ListWhitelistedAddressOp = listWhitelistedAddressOp{
	Summary: "List an array of whitelisted addresses",
}

type listWhitelistedPerpAddressOp struct {
	Summary string
}

var ListWhitelistedPerpAddressOp = listWhitelistedPerpAddressOp{
	Summary: "List an array of whitelisted addresses",
}

type optionBarsOp struct {
	Summary   string
	End       string
	Limit     string
	PageToken string
	Sort      string
	Start     string
	Symbols   string
	Timeframe string
}

var OptionBarsOp = optionBarsOp{
	Summary:   "Get historical bars",
	End:       "inclusive end of the interval",
	Limit:     "maximum number of data points to return in the response page.",
	PageToken: "pagination token from which to continue",
	Sort:      "sort data in ascending or descending order",
	Start:     "inclusive start of the interval",
	Symbols:   "A comma-separated list of contract symbols with a limit of 100",
	Timeframe: "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g",
}

var OptionBarsFlags = []FlagDef{
	{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval"},
	{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page."},
	{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue"},
	{Name: "sort", OASName: "sort", Type: "string", Description: "sort data in ascending or descending order"},
	{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval"},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100"},
	{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g"},
}

type optionDoNotExerciseOp struct {
	Summary            string
	SymbolOrContractID string
}

var OptionDoNotExerciseOp = optionDoNotExerciseOp{
	Summary:            "Do not exercise an options position",
	SymbolOrContractID: "option contract symbol or ID",
}

type optionExerciseOp struct {
	Summary            string
	SymbolOrContractID string
}

var OptionExerciseOp = optionExerciseOp{
	Summary:            "Exercise an options position",
	SymbolOrContractID: "option contract symbol or ID",
}

type patchAccountConfigOp struct {
	Summary                 string
	DisableOvernightTrading string
	DTBPCheck               string
	FractionalTrading       string
	MaxMarginMultiplier     string
	MaxOptionsTradingLevel  string
	NoShorting              string
	PDTCheck                string
	PtpNoExceptionEntry     string
	SuspendTrade            string
	TradeConfirmEmail       string
}

var PatchAccountConfigOp = patchAccountConfigOp{
	Summary:                 "Update account configurations",
	DisableOvernightTrading: "if true, overnight trading is disabled",
	DTBPCheck:               "both, entry, or exit. Controls Day Trading Margin Call (DTMC) checks",
	FractionalTrading:       "if true, account is able to participate in fractional trading",
	MaxMarginMultiplier:     "can be \"1\", \"2\", or \"4\"",
	MaxOptionsTradingLevel:  "desired maximum options trading level. 0=disabled, 1=Covered Call/Cash-Secured Put, 2=Long Call/Put, 3=Spreads/Straddles",
	NoShorting:              "if true, account becomes long-only mode",
	PDTCheck:                "both, entry, or exit",
	PtpNoExceptionEntry:     "if set to true then Alpaca will accept orders for PTP symbols with no exception. Default is false",
	SuspendTrade:            "if true, new orders are blocked",
	TradeConfirmEmail:       "all or none. If none, emails for order fills are not sent",
}

var PatchAccountConfigFlags = []FlagDef{
	{Name: "disable-overnight-trading", OASName: "disable_overnight_trading", Type: "bool", Description: "if true, overnight trading is disabled"},
	{Name: "dtbp-check", OASName: "dtbp_check", Type: "string", Description: "both, entry, or exit. Controls Day Trading Margin Call (DTMC) checks", Completions: []string{"both", "entry", "exit"}},
	{Name: "fractional-trading", OASName: "fractional_trading", Type: "bool", Description: "if true, account is able to participate in fractional trading"},
	{Name: "max-margin-multiplier", OASName: "max_margin_multiplier", Type: "string", Description: "can be \"1\", \"2\", or \"4\""},
	{Name: "max-options-trading-level", OASName: "max_options_trading_level", Type: "int", Description: "desired maximum options trading level. 0=disabled, 1=Covered Call/Cash-Secured Put, 2=Long Call/Put, 3=Spreads/Straddles", Completions: []string{"0", "1", "2", "3"}},
	{Name: "no-shorting", OASName: "no_shorting", Type: "bool", Description: "if true, account becomes long-only mode"},
	{Name: "pdt-check", OASName: "pdt_check", Type: "string", Description: "both, entry, or exit"},
	{Name: "ptp-no-exception-entry", OASName: "ptp_no_exception_entry", Type: "bool", Description: "if set to true then Alpaca will accept orders for PTP symbols with no exception. Default is false"},
	{Name: "suspend-trade", OASName: "suspend_trade", Type: "bool", Description: "if true, new orders are blocked"},
	{Name: "trade-confirm-email", OASName: "trade_confirm_email", Type: "string", Description: "all or none. If none, emails for order fills are not sent"},
}

type patchOrderByOrderIDOp struct {
	Summary              string
	AdvancedInstructions string
	ClientOrderID        string
	LimitPrice           string
	OrderID              string
	Qty                  string
	StopPrice            string
	TimeInForce          string
	Trail                string
}

var PatchOrderByOrderIDOp = patchOrderByOrderIDOp{
	Summary:              "Replace order by ID",
	AdvancedInstructions: "advanced instructions for Elite Smart Router: https://docs.alpaca.markets/docs/alpaca-elite-smart-router",
	ClientOrderID:        "A unique identifier for the new order. Automatically generated if not sent. (<= 128 characters)",
	LimitPrice:           "required if original order's type field was limit or stop_limit.",
	OrderID:              "order id",
	Qty:                  "number of shares to trade.",
	StopPrice:            "required if original order type is limit or stop_limit",
	TimeInForce:          "time-In-Force values supported by Alpaca vary based on the order's security type",
	Trail:                "the new value of the trail_price or trail_percent value (works only for type=“trailing_stop”)",
}

var PatchOrderByOrderIDFlags = []FlagDef{
	{Name: "advanced-instructions", OASName: "advanced_instructions", Type: "string", Description: "advanced instructions for Elite Smart Router: https://docs.alpaca.markets/docs/alpaca-elite-smart-router"},
	{Name: "client-order-id", OASName: "client_order_id", Type: "string", Description: "A unique identifier for the new order. Automatically generated if not sent. (<= 128 characters)"},
	{Name: "limit-price", OASName: "limit_price", Type: "string", Description: "required if original order's type field was limit or stop_limit."},
	{Name: "qty", OASName: "qty", Type: "string", Description: "number of shares to trade."},
	{Name: "stop-price", OASName: "stop_price", Type: "string", Description: "required if original order type is limit or stop_limit"},
	{Name: "time-in-force", OASName: "time_in_force", Type: "string", Description: "time-In-Force values supported by Alpaca vary based on the order's security type", Completions: []string{"cls", "day", "fok", "gtc", "ioc", "opg"}},
	{Name: "trail", OASName: "trail", Type: "string", Description: "the new value of the trail_price or trail_percent value (works only for type=“trailing_stop”)"},
}

type postOrderOp struct {
	Summary              string
	AdvancedInstructions string
	ClientOrderID        string
	ExtendedHours        string
	Legs                 string
	LimitPrice           string
	Notional             string
	OrderClass           string
	PositionIntent       string
	Qty                  string
	Side                 string
	StopLoss             string
	StopPrice            string
	Symbol               string
	TakeProfit           string
	TimeInForce          string
	TrailPercent         string
	TrailPrice           string
	Type                 string
}

var PostOrderOp = postOrderOp{
	Summary:              "Create an order",
	AdvancedInstructions: "advanced instructions for Elite Smart Router: https://docs.alpaca.markets/docs/alpaca-elite-smart-router",
	ClientOrderID:        "A unique identifier for the order. Automatically generated if not sent. (<= 128 characters)",
	ExtendedHours:        "(default) false",
	Legs:                 "list of order legs (<= 4)",
	LimitPrice:           "required if type is limit or stop_limit.",
	Notional:             "dollar amount to trade. Cannot work with qty. Can only work for market order types and day for time in force",
	OrderClass:           "order classes supported by Alpaca vary based on the order's security type",
	PositionIntent:       "represents the desired position strategy",
	Qty:                  "number of shares to trade",
	Side:                 "represents which side this order was on:\n- buy\n- sell\nRequired for all order classes except for mleg",
	StopLoss:             "takes in string/number values for stop_price and limit_price",
	StopPrice:            "required if type is stop or stop_limit",
	Symbol:               "symbol, asset ID, or currency pair to identify the asset to trade, required for all order classes except for mleg",
	TakeProfit:           "takes in a string/number value for limit_price",
	TimeInForce:          "time-In-Force values supported by Alpaca vary based on the order's security type",
	TrailPercent:         "this or trail_price is required if type is trailing_stop",
	TrailPrice:           "this or trail_percent is required if type is trailing_stop",
	Type:                 "order types supported by Alpaca vary based on the order's security type",
}

var PostOrderFlags = []FlagDef{
	{Name: "advanced-instructions", OASName: "advanced_instructions", Type: "string", Description: "advanced instructions for Elite Smart Router: https://docs.alpaca.markets/docs/alpaca-elite-smart-router"},
	{Name: "client-order-id", OASName: "client_order_id", Type: "string", Description: "A unique identifier for the order. Automatically generated if not sent. (<= 128 characters)"},
	{Name: "extended-hours", OASName: "extended_hours", Type: "bool", Description: "(default) false"},
	{Name: "legs", OASName: "legs", Type: "string", Description: "list of order legs (<= 4)"},
	{Name: "limit-price", OASName: "limit_price", Type: "string", Description: "required if type is limit or stop_limit."},
	{Name: "notional", OASName: "notional", Type: "string", Description: "dollar amount to trade. Cannot work with qty. Can only work for market order types and day for time in force"},
	{Name: "order-class", OASName: "order_class", Type: "string", Description: "order classes supported by Alpaca vary based on the order's security type", Completions: []string{"bracket", "mleg", "oco", "oto", "simple"}},
	{Name: "position-intent", OASName: "position_intent", Type: "string", Description: "represents the desired position strategy", Completions: []string{"buy_to_close", "buy_to_open", "sell_to_close", "sell_to_open"}},
	{Name: "qty", OASName: "qty", Type: "string", Description: "number of shares to trade"},
	{Name: "side", OASName: "side", Type: "string", Description: "represents which side this order was on:\n- buy\n- sell\nRequired for all order classes except for mleg", Completions: []string{"buy", "sell"}},
	{Name: "stop-loss", OASName: "stop_loss", Type: "string", Description: "takes in string/number values for stop_price and limit_price"},
	{Name: "stop-price", OASName: "stop_price", Type: "string", Description: "required if type is stop or stop_limit"},
	{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol, asset ID, or currency pair to identify the asset to trade, required for all order classes except for mleg"},
	{Name: "take-profit", OASName: "take_profit", Type: "string", Description: "takes in a string/number value for limit_price"},
	{Name: "time-in-force", OASName: "time_in_force", Type: "string", Description: "time-In-Force values supported by Alpaca vary based on the order's security type", Completions: []string{"cls", "day", "fok", "gtc", "ioc", "opg"}},
	{Name: "trail-percent", OASName: "trail_percent", Type: "string", Description: "this or trail_price is required if type is trailing_stop"},
	{Name: "trail-price", OASName: "trail_price", Type: "string", Description: "this or trail_percent is required if type is trailing_stop"},
	{Name: "type", OASName: "type", Type: "string", Description: "order types supported by Alpaca vary based on the order's security type", Completions: []string{"limit", "market", "stop", "stop_limit", "trailing_stop"}},
}

type postWatchlistOp struct {
	Summary string
	Name    string
	Symbols string
}

var PostWatchlistOp = postWatchlistOp{
	Summary: "Create watchlist",
	Name:    "name",
	Symbols: "symbols",
}

var PostWatchlistFlags = []FlagDef{
	{Name: "name", OASName: "name", Type: "string", Description: "name"},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "symbols"},
}

type removeAssetFromWatchlistOp struct {
	Summary     string
	Symbol      string
	WatchlistID string
}

var RemoveAssetFromWatchlistOp = removeAssetFromWatchlistOp{
	Summary:     "Delete symbol from watchlist",
	Symbol:      "symbol name to remove from the watchlist content",
	WatchlistID: "watchlist ID",
}

type setCryptoPerpAccountLeverageOp struct {
	Summary  string
	Leverage string
	Symbol   string
}

var SetCryptoPerpAccountLeverageOp = setCryptoPerpAccountLeverageOp{
	Summary:  "Set account leverage for an asset",
	Leverage: "leverage for the underlying asset",
	Symbol:   "symbol of underlying asset",
}

var SetCryptoPerpAccountLeverageFlags = []FlagDef{
	{Name: "leverage", OASName: "leverage", Type: "int", Description: "leverage for the underlying asset"},
	{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol of underlying asset"},
}

type updateWatchlistByIDOp struct {
	Summary     string
	Name        string
	Symbols     string
	WatchlistID string
}

var UpdateWatchlistByIDOp = updateWatchlistByIDOp{
	Summary:     "Update watchlist by id",
	Name:        "name",
	Symbols:     "symbols",
	WatchlistID: "watchlist id",
}

var UpdateWatchlistByIDFlags = []FlagDef{
	{Name: "name", OASName: "name", Type: "string", Description: "name"},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "symbols"},
}

type updateWatchlistByNameOp struct {
	Summary string
	Name    string
	Symbols string
}

var UpdateWatchlistByNameOp = updateWatchlistByNameOp{
	Summary: "Update watchlist by name",
	Name:    "name of the watchlist",
	Symbols: "symbols",
}

var UpdateWatchlistByNameFlags = []FlagDef{
	{Name: "name", OASName: "name", Type: "string", Description: "name of the watchlist"},
	{Name: "symbols", OASName: "symbols", Type: "string", Description: "symbols"},
}
