// Code generated from api/specs; DO NOT EDIT.

package api

// Op is satisfied by every generated operation variable (e.g. GetAccountOp).
// Use it to pass operations type-safely instead of raw strings.
type Op interface {
	Summary() string
	Flags() []FlagDef
	RequiredFlags() []string
	ResponseFields() []ResponseField
}

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
}

type calendarOp struct{}

var CalendarOp = calendarOp{}

func (o calendarOp) Summary() string {
	return "Get market calendar"
}

func (o calendarOp) ResponseFields() []ResponseField {
	return ResponseSchemas["Calendar"]
}

func (o calendarOp) RequiredFlags() []string {
	return nil
}

func (o calendarOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "last date to retrieve data for (inclusive). Default: one week from the start date", OpName: "Calendar"},
		{Name: "start", OASName: "start", Type: "string", Description: "first date to retrieve data for (inclusive). Default: today", OpName: "Calendar"},
		{Name: "timezone", OASName: "timezone", Type: "string", Description: "timezone of the times. Default: the timezone of the market", Completions: []string{"UTC"}, OpName: "Calendar"},
	}
}

type clockOp struct{}

var ClockOp = clockOp{}

func (o clockOp) Summary() string {
	return "Get market clock"
}

func (o clockOp) ResponseFields() []ResponseField {
	return ResponseSchemas["Clock"]
}

func (o clockOp) RequiredFlags() []string {
	return nil
}

func (o clockOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "markets", OASName: "markets", Type: "string", Description: "comma-separated list of markets", OpName: "Clock"},
		{Name: "time", OASName: "time", Type: "string", Description: "instead of the current time, use this time for the clock", OpName: "Clock"},
	}
}

type corporateActionsOp struct{}

var CorporateActionsOp = corporateActionsOp{}

func (o corporateActionsOp) Summary() string {
	return "Get corporate actions"
}

func (o corporateActionsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["CorporateActions"]
}

func (o corporateActionsOp) RequiredFlags() []string {
	return nil
}

func (o corporateActionsOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "cusips", OASName: "cusips", Type: "string", Description: "A comma-separated list of CUSIPs", OpName: "CorporateActions"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "CorporateActions"},
		{Name: "ids", OASName: "ids", Type: "string", Description: "A comma-separated list of corporate action IDs", OpName: "CorporateActions"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "100", Description: "maximum number of corporate actions to return in a response.", OpName: "CorporateActions"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "CorporateActions"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "CorporateActions"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "CorporateActions"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of symbols", OpName: "CorporateActions"},
		{Name: "types", OASName: "types", Type: "string", Description: "A comma-separated list of types", OpName: "CorporateActions"},
	}
}

type cryptoBarsOp struct{}

var CryptoBarsOp = cryptoBarsOp{}

func (o cryptoBarsOp) Summary() string {
	return "Get historical bars"
}

func (o cryptoBarsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["CryptoBars"]
}

func (o cryptoBarsOp) RequiredFlags() []string {
	return []string{"symbols", "timeframe"}
}

func (o cryptoBarsOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "CryptoBars"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "CryptoBars"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "CryptoBars"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "CryptoBars"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "CryptoBars"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoBars", Required: true},
		{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g", OpName: "CryptoBars", Required: true},
	}
}

type cryptoLatestBarsOp struct{}

var CryptoLatestBarsOp = cryptoLatestBarsOp{}

func (o cryptoLatestBarsOp) Summary() string {
	return "Get latest bars"
}

func (o cryptoLatestBarsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["CryptoLatestBars"]
}

func (o cryptoLatestBarsOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o cryptoLatestBarsOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoLatestBars", Required: true},
	}
}

type cryptoLatestOrderbooksOp struct{}

var CryptoLatestOrderbooksOp = cryptoLatestOrderbooksOp{}

func (o cryptoLatestOrderbooksOp) Summary() string {
	return "Get latest orderbook"
}

func (o cryptoLatestOrderbooksOp) ResponseFields() []ResponseField {
	return ResponseSchemas["CryptoLatestOrderbooks"]
}

func (o cryptoLatestOrderbooksOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o cryptoLatestOrderbooksOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoLatestOrderbooks", Required: true},
	}
}

type cryptoLatestQuotesOp struct{}

var CryptoLatestQuotesOp = cryptoLatestQuotesOp{}

func (o cryptoLatestQuotesOp) Summary() string {
	return "Get latest quotes"
}

func (o cryptoLatestQuotesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["CryptoLatestQuotes"]
}

func (o cryptoLatestQuotesOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o cryptoLatestQuotesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoLatestQuotes", Required: true},
	}
}

type cryptoLatestTradesOp struct{}

var CryptoLatestTradesOp = cryptoLatestTradesOp{}

func (o cryptoLatestTradesOp) Summary() string {
	return "Get latest trades"
}

func (o cryptoLatestTradesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["CryptoLatestTrades"]
}

func (o cryptoLatestTradesOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o cryptoLatestTradesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoLatestTrades", Required: true},
	}
}

type cryptoPerpLatestBarsOp struct{}

var CryptoPerpLatestBarsOp = cryptoPerpLatestBarsOp{}

func (o cryptoPerpLatestBarsOp) Summary() string {
	return "Get latest bars"
}

func (o cryptoPerpLatestBarsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["CryptoPerpLatestBars"]
}

func (o cryptoPerpLatestBarsOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o cryptoPerpLatestBarsOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoPerpLatestBars", Required: true},
	}
}

type cryptoPerpLatestFuturesPricingOp struct{}

var CryptoPerpLatestFuturesPricingOp = cryptoPerpLatestFuturesPricingOp{}

func (o cryptoPerpLatestFuturesPricingOp) Summary() string {
	return "Get latest pricing"
}

func (o cryptoPerpLatestFuturesPricingOp) ResponseFields() []ResponseField {
	return ResponseSchemas["CryptoPerpLatestFuturesPricing"]
}

func (o cryptoPerpLatestFuturesPricingOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o cryptoPerpLatestFuturesPricingOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoPerpLatestFuturesPricing", Required: true},
	}
}

type cryptoPerpLatestOrderbooksOp struct{}

var CryptoPerpLatestOrderbooksOp = cryptoPerpLatestOrderbooksOp{}

func (o cryptoPerpLatestOrderbooksOp) Summary() string {
	return "Get latest orderbook"
}

func (o cryptoPerpLatestOrderbooksOp) ResponseFields() []ResponseField {
	return ResponseSchemas["CryptoPerpLatestOrderbooks"]
}

func (o cryptoPerpLatestOrderbooksOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o cryptoPerpLatestOrderbooksOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoPerpLatestOrderbooks", Required: true},
	}
}

type cryptoPerpLatestQuotesOp struct{}

var CryptoPerpLatestQuotesOp = cryptoPerpLatestQuotesOp{}

func (o cryptoPerpLatestQuotesOp) Summary() string {
	return "Get latest quotes"
}

func (o cryptoPerpLatestQuotesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["CryptoPerpLatestQuotes"]
}

func (o cryptoPerpLatestQuotesOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o cryptoPerpLatestQuotesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoPerpLatestQuotes", Required: true},
	}
}

type cryptoPerpLatestTradesOp struct{}

var CryptoPerpLatestTradesOp = cryptoPerpLatestTradesOp{}

func (o cryptoPerpLatestTradesOp) Summary() string {
	return "Get latest trades"
}

func (o cryptoPerpLatestTradesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["CryptoPerpLatestTrades"]
}

func (o cryptoPerpLatestTradesOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o cryptoPerpLatestTradesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoPerpLatestTrades", Required: true},
	}
}

type cryptoQuotesOp struct{}

var CryptoQuotesOp = cryptoQuotesOp{}

func (o cryptoQuotesOp) Summary() string {
	return "Get historical quotes"
}

func (o cryptoQuotesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["CryptoQuotes"]
}

func (o cryptoQuotesOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o cryptoQuotesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "CryptoQuotes"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "CryptoQuotes"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "CryptoQuotes"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "CryptoQuotes"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "CryptoQuotes"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoQuotes", Required: true},
	}
}

type cryptoSnapshotsOp struct{}

var CryptoSnapshotsOp = cryptoSnapshotsOp{}

func (o cryptoSnapshotsOp) Summary() string {
	return "Get snapshots"
}

func (o cryptoSnapshotsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["CryptoSnapshots"]
}

func (o cryptoSnapshotsOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o cryptoSnapshotsOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoSnapshots", Required: true},
	}
}

type cryptoTradesOp struct{}

var CryptoTradesOp = cryptoTradesOp{}

func (o cryptoTradesOp) Summary() string {
	return "Get historical trades"
}

func (o cryptoTradesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["CryptoTrades"]
}

func (o cryptoTradesOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o cryptoTradesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "CryptoTrades"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "CryptoTrades"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "CryptoTrades"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "CryptoTrades"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "CryptoTrades"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", OpName: "CryptoTrades", Required: true},
	}
}

type fixedIncomeLatestPricesOp struct{}

var FixedIncomeLatestPricesOp = fixedIncomeLatestPricesOp{}

func (o fixedIncomeLatestPricesOp) Summary() string {
	return "Get latest prices"
}

func (o fixedIncomeLatestPricesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["FixedIncomeLatestPrices"]
}

func (o fixedIncomeLatestPricesOp) RequiredFlags() []string {
	return []string{"isins"}
}

func (o fixedIncomeLatestPricesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "isins", OASName: "isins", Type: "string", Description: "A comma-separated list of ISINs with a limit of 1000", OpName: "FixedIncomeLatestPrices", Required: true},
	}
}

type latestRatesOp struct{}

var LatestRatesOp = latestRatesOp{}

func (o latestRatesOp) Summary() string {
	return "Get latest rates for currency pairs"
}

func (o latestRatesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["LatestRates"]
}

func (o latestRatesOp) RequiredFlags() []string {
	return []string{"currency-pairs"}
}

func (o latestRatesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "currency-pairs", OASName: "currency_pairs", Type: "string", Description: "A comma-separated string with currency pairs", OpName: "LatestRates", Required: true},
	}
}

type legacyCalendarOp struct{}

var LegacyCalendarOp = legacyCalendarOp{}

func (o legacyCalendarOp) Summary() string {
	return "Get US market calendar"
}

func (o legacyCalendarOp) ResponseFields() []ResponseField {
	return ResponseSchemas["LegacyCalendar"]
}

func (o legacyCalendarOp) RequiredFlags() []string {
	return nil
}

func (o legacyCalendarOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "date-type", OASName: "date_type", Type: "string", Description: "indicates what start and end mean", Completions: []string{"SETTLEMENT", "TRADING"}, OpName: "LegacyCalendar"},
		{Name: "end", OASName: "end", Type: "string", Description: "last date to retrieve data for (inclusive)", OpName: "LegacyCalendar"},
		{Name: "start", OASName: "start", Type: "string", Description: "first date to retrieve data for (inclusive)", OpName: "LegacyCalendar"},
	}
}

type legacyClockOp struct{}

var LegacyClockOp = legacyClockOp{}

func (o legacyClockOp) Summary() string {
	return "Get US market clock"
}

func (o legacyClockOp) ResponseFields() []ResponseField {
	return ResponseSchemas["LegacyClock"]
}

func (o legacyClockOp) RequiredFlags() []string {
	return nil
}

func (o legacyClockOp) Flags() []FlagDef {
	return nil
}

type logosOp struct{}

var LogosOp = logosOp{}

func (o logosOp) Summary() string {
	return "Get logos"
}

func (o logosOp) ResponseFields() []ResponseField {
	return ResponseSchemas["Logos"]
}

func (o logosOp) RequiredFlags() []string {
	return nil
}

func (o logosOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "placeholder", OASName: "placeholder", Type: "bool", Default: "true", Description: "placeholder", OpName: "Logos"},
	}
}

type mostActivesOp struct{}

var MostActivesOp = mostActivesOp{}

func (o mostActivesOp) Summary() string {
	return "Get most active stocks"
}

func (o mostActivesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["MostActives"]
}

func (o mostActivesOp) RequiredFlags() []string {
	return nil
}

func (o mostActivesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "by", OASName: "by", Type: "string", Default: "volume", Description: "metric used for ranking the most active stocks", Completions: []string{"trades", "volume"}, OpName: "MostActives"},
		{Name: "top", OASName: "top", Type: "int", Default: "10", Description: "number of top most active stocks to fetch per day", OpName: "MostActives"},
	}
}

type moversOp struct{}

var MoversOp = moversOp{}

func (o moversOp) Summary() string {
	return "Get top market movers"
}

func (o moversOp) ResponseFields() []ResponseField {
	return ResponseSchemas["Movers"]
}

func (o moversOp) RequiredFlags() []string {
	return nil
}

func (o moversOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "top", OASName: "top", Type: "int", Default: "10", Description: "number of top market movers to fetch (gainers and losers)", OpName: "Movers"},
	}
}

type newsOp struct{}

var NewsOp = newsOp{}

func (o newsOp) Summary() string {
	return "Get news articles"
}

func (o newsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["News"]
}

func (o newsOp) RequiredFlags() []string {
	return nil
}

func (o newsOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "News"},
		{Name: "exclude-contentless", OASName: "exclude_contentless", Type: "bool", Description: "boolean indicator to exclude news articles that do not contain content", OpName: "News"},
		{Name: "include-content", OASName: "include_content", Type: "bool", Description: "boolean indicator to include content for news articles (if available)", OpName: "News"},
		{Name: "limit", OASName: "limit", Type: "int", Description: "limit of news items to be returned for a result page", OpName: "News"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "News"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "desc", Description: "sort articles by updated date", Completions: []string{"asc", "desc"}, OpName: "News"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "News"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of symbols for which to query news", OpName: "News"},
	}
}

type optionChainOp struct{}

var OptionChainOp = optionChainOp{}

func (o optionChainOp) Summary() string {
	return "Get option chain"
}

func (o optionChainOp) ResponseFields() []ResponseField {
	return ResponseSchemas["OptionChain"]
}

func (o optionChainOp) RequiredFlags() []string {
	return nil
}

func (o optionChainOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "expiration-date", OASName: "expiration_date", Type: "string", Description: "filter contracts by the exact expiration date (format: YYYY-MM-DD)", OpName: "OptionChain"},
		{Name: "expiration-date-gte", OASName: "expiration_date_gte", Type: "string", Description: "filter contracts with expiration date greater than or equal to the specified date", OpName: "OptionChain"},
		{Name: "expiration-date-lte", OASName: "expiration_date_lte", Type: "string", Description: "filter contracts with expiration date less than or equal to the specified date", OpName: "OptionChain"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "opra", Description: "source feed of the data", Completions: []string{"indicative", "opra"}, OpName: "OptionChain"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "100", Description: "number of maximum snapshots to return in a response.", OpName: "OptionChain"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "OptionChain"},
		{Name: "root-symbol", OASName: "root_symbol", Type: "string", Description: "filter contracts by the root symbol", OpName: "OptionChain"},
		{Name: "strike-price-gte", OASName: "strike_price_gte", Type: "string", Description: "filter contracts with strike price greater than or equal to the specified value", OpName: "OptionChain"},
		{Name: "strike-price-lte", OASName: "strike_price_lte", Type: "string", Description: "filter contracts with strike price less than or equal to the specified value", OpName: "OptionChain"},
		{Name: "type", OASName: "type", Type: "string", Description: "filter contracts by the type (call or put)", Completions: []string{"call", "put"}, OpName: "OptionChain"},
		{Name: "updated-since", OASName: "updated_since", Type: "string", Description: "filter to snapshots that were updated since this timestamp, meaning that the timestamp of the trade or the quote is g...", OpName: "OptionChain"},
	}
}

type optionLatestQuotesOp struct{}

var OptionLatestQuotesOp = optionLatestQuotesOp{}

func (o optionLatestQuotesOp) Summary() string {
	return "Get latest quotes"
}

func (o optionLatestQuotesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["OptionLatestQuotes"]
}

func (o optionLatestQuotesOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o optionLatestQuotesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "feed", OASName: "feed", Type: "string", Default: "opra", Description: "source feed of the data", Completions: []string{"indicative", "opra"}, OpName: "OptionLatestQuotes"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100", OpName: "OptionLatestQuotes", Required: true},
	}
}

type optionLatestTradesOp struct{}

var OptionLatestTradesOp = optionLatestTradesOp{}

func (o optionLatestTradesOp) Summary() string {
	return "Get latest trades"
}

func (o optionLatestTradesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["OptionLatestTrades"]
}

func (o optionLatestTradesOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o optionLatestTradesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "feed", OASName: "feed", Type: "string", Default: "opra", Description: "source feed of the data", Completions: []string{"indicative", "opra"}, OpName: "OptionLatestTrades"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100", OpName: "OptionLatestTrades", Required: true},
	}
}

type optionMetaConditionsOp struct{}

var OptionMetaConditionsOp = optionMetaConditionsOp{}

func (o optionMetaConditionsOp) Summary() string {
	return "Get condition codes"
}

func (o optionMetaConditionsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["OptionMetaConditions"]
}

func (o optionMetaConditionsOp) RequiredFlags() []string {
	return nil
}

func (o optionMetaConditionsOp) Flags() []FlagDef {
	return nil
}

type optionMetaExchangesOp struct{}

var OptionMetaExchangesOp = optionMetaExchangesOp{}

func (o optionMetaExchangesOp) Summary() string {
	return "Get exchange codes"
}

func (o optionMetaExchangesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["OptionMetaExchanges"]
}

func (o optionMetaExchangesOp) RequiredFlags() []string {
	return nil
}

func (o optionMetaExchangesOp) Flags() []FlagDef {
	return nil
}

type optionSnapshotsOp struct{}

var OptionSnapshotsOp = optionSnapshotsOp{}

func (o optionSnapshotsOp) Summary() string {
	return "Get snapshots"
}

func (o optionSnapshotsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["OptionSnapshots"]
}

func (o optionSnapshotsOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o optionSnapshotsOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "feed", OASName: "feed", Type: "string", Default: "opra", Description: "source feed of the data", Completions: []string{"indicative", "opra"}, OpName: "OptionSnapshots"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "100", Description: "number of maximum snapshots to return in a response.", OpName: "OptionSnapshots"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "OptionSnapshots"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100", OpName: "OptionSnapshots", Required: true},
		{Name: "updated-since", OASName: "updated_since", Type: "string", Description: "filter to snapshots that were updated since this timestamp, meaning that the timestamp of the trade or the quote is g...", OpName: "OptionSnapshots"},
	}
}

type optionTradesOp struct{}

var OptionTradesOp = optionTradesOp{}

func (o optionTradesOp) Summary() string {
	return "Get historical trades"
}

func (o optionTradesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["OptionTrades"]
}

func (o optionTradesOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o optionTradesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "OptionTrades"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "OptionTrades"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "OptionTrades"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "OptionTrades"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "OptionTrades"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100", OpName: "OptionTrades", Required: true},
	}
}

type ratesOp struct{}

var RatesOp = ratesOp{}

func (o ratesOp) Summary() string {
	return "Get historical rates for currency pairs"
}

func (o ratesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["Rates"]
}

func (o ratesOp) RequiredFlags() []string {
	return []string{"currency-pairs"}
}

func (o ratesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "currency-pairs", OASName: "currency_pairs", Type: "string", Description: "A comma-separated string with currency pairs", OpName: "Rates", Required: true},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "Rates"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "Rates"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "Rates"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "Rates"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "Rates"},
		{Name: "timeframe", OASName: "timeframe", Type: "string", Default: "1Min", Description: "sampling interval of the currency rates", OpName: "Rates"},
	}
}

type stockAuctionSingleOp struct{}

var StockAuctionSingleOp = stockAuctionSingleOp{}

func (o stockAuctionSingleOp) Summary() string {
	return "Get historical auctions (single)"
}

func (o stockAuctionSingleOp) ResponseFields() []ResponseField {
	return ResponseSchemas["StockAuctionSingle"]
}

func (o stockAuctionSingleOp) RequiredFlags() []string {
	return nil
}

func (o stockAuctionSingleOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", OpName: "StockAuctionSingle"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockAuctionSingle"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "StockAuctionSingle"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "only sip is valid for auctions", OpName: "StockAuctionSingle"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "StockAuctionSingle"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "StockAuctionSingle"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "StockAuctionSingle"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "StockAuctionSingle"},
	}
}

type stockAuctionsOp struct{}

var StockAuctionsOp = stockAuctionsOp{}

func (o stockAuctionsOp) Summary() string {
	return "Get historical auctions"
}

func (o stockAuctionsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["StockAuctions"]
}

func (o stockAuctionsOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o stockAuctionsOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", OpName: "StockAuctions"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockAuctions"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "StockAuctions"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "only sip is valid for auctions", OpName: "StockAuctions"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "StockAuctions"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "StockAuctions"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "StockAuctions"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "StockAuctions"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", OpName: "StockAuctions", Required: true},
	}
}

type stockBarSingleOp struct{}

var StockBarSingleOp = stockBarSingleOp{}

func (o stockBarSingleOp) Summary() string {
	return "Get historical bars (single symbol)"
}

func (o stockBarSingleOp) ResponseFields() []ResponseField {
	return ResponseSchemas["StockBarSingle"]
}

func (o stockBarSingleOp) RequiredFlags() []string {
	return []string{"timeframe"}
}

func (o stockBarSingleOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "adjustment", OASName: "adjustment", Type: "string", Default: "raw", Description: "specifies the adjustments for the bars.\n\n - raw: no adjustments\n - split: adjust price and volume for forward and rev...", OpName: "StockBarSingle"},
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", OpName: "StockBarSingle"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockBarSingle"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "StockBarSingle"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "source feed of the data.", Completions: []string{"boats", "iex", "otc", "sip"}, OpName: "StockBarSingle"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "StockBarSingle"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "StockBarSingle"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "StockBarSingle"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "StockBarSingle"},
		{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g", OpName: "StockBarSingle", Required: true},
	}
}

type stockBarsOp struct{}

var StockBarsOp = stockBarsOp{}

func (o stockBarsOp) Summary() string {
	return "Get historical bars"
}

func (o stockBarsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["StockBars"]
}

func (o stockBarsOp) RequiredFlags() []string {
	return []string{"symbols", "timeframe"}
}

func (o stockBarsOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "adjustment", OASName: "adjustment", Type: "string", Default: "raw", Description: "specifies the adjustments for the bars.\n\n - raw: no adjustments\n - split: adjust price and volume for forward and rev...", OpName: "StockBars"},
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", OpName: "StockBars"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockBars"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "StockBars"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "source feed of the data.", Completions: []string{"boats", "iex", "otc", "sip"}, OpName: "StockBars"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "StockBars"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "StockBars"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "StockBars"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "StockBars"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", OpName: "StockBars", Required: true},
		{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g", OpName: "StockBars", Required: true},
	}
}

type stockLatestBarSingleOp struct{}

var StockLatestBarSingleOp = stockLatestBarSingleOp{}

func (o stockLatestBarSingleOp) Summary() string {
	return "Get latest bar (single symbol)"
}

func (o stockLatestBarSingleOp) ResponseFields() []ResponseField {
	return ResponseSchemas["StockLatestBarSingle"]
}

func (o stockLatestBarSingleOp) RequiredFlags() []string {
	return nil
}

func (o stockLatestBarSingleOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockLatestBarSingle"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, OpName: "StockLatestBarSingle"},
	}
}

type stockLatestBarsOp struct{}

var StockLatestBarsOp = stockLatestBarsOp{}

func (o stockLatestBarsOp) Summary() string {
	return "Get latest bars"
}

func (o stockLatestBarsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["StockLatestBars"]
}

func (o stockLatestBarsOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o stockLatestBarsOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockLatestBars"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, OpName: "StockLatestBars"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", OpName: "StockLatestBars", Required: true},
	}
}

type stockLatestQuoteSingleOp struct{}

var StockLatestQuoteSingleOp = stockLatestQuoteSingleOp{}

func (o stockLatestQuoteSingleOp) Summary() string {
	return "Get latest quote (single symbol)"
}

func (o stockLatestQuoteSingleOp) ResponseFields() []ResponseField {
	return ResponseSchemas["StockLatestQuoteSingle"]
}

func (o stockLatestQuoteSingleOp) RequiredFlags() []string {
	return nil
}

func (o stockLatestQuoteSingleOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockLatestQuoteSingle"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, OpName: "StockLatestQuoteSingle"},
	}
}

type stockLatestQuotesOp struct{}

var StockLatestQuotesOp = stockLatestQuotesOp{}

func (o stockLatestQuotesOp) Summary() string {
	return "Get latest quotes"
}

func (o stockLatestQuotesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["StockLatestQuotes"]
}

func (o stockLatestQuotesOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o stockLatestQuotesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockLatestQuotes"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, OpName: "StockLatestQuotes"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", OpName: "StockLatestQuotes", Required: true},
	}
}

type stockLatestTradeSingleOp struct{}

var StockLatestTradeSingleOp = stockLatestTradeSingleOp{}

func (o stockLatestTradeSingleOp) Summary() string {
	return "Get latest trade (single symbol)"
}

func (o stockLatestTradeSingleOp) ResponseFields() []ResponseField {
	return ResponseSchemas["StockLatestTradeSingle"]
}

func (o stockLatestTradeSingleOp) RequiredFlags() []string {
	return nil
}

func (o stockLatestTradeSingleOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockLatestTradeSingle"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, OpName: "StockLatestTradeSingle"},
	}
}

type stockLatestTradesOp struct{}

var StockLatestTradesOp = stockLatestTradesOp{}

func (o stockLatestTradesOp) Summary() string {
	return "Get latest trades"
}

func (o stockLatestTradesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["StockLatestTrades"]
}

func (o stockLatestTradesOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o stockLatestTradesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockLatestTrades"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, OpName: "StockLatestTrades"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", OpName: "StockLatestTrades", Required: true},
	}
}

type stockMetaConditionsOp struct{}

var StockMetaConditionsOp = stockMetaConditionsOp{}

func (o stockMetaConditionsOp) Summary() string {
	return "Get condition codes"
}

func (o stockMetaConditionsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["StockMetaConditions"]
}

func (o stockMetaConditionsOp) RequiredFlags() []string {
	return []string{"tape"}
}

func (o stockMetaConditionsOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "tape", OASName: "tape", Type: "string", Description: "one character name of the tape", Completions: []string{"A", "B", "C"}, OpName: "StockMetaConditions", Required: true},
	}
}

type stockMetaExchangesOp struct{}

var StockMetaExchangesOp = stockMetaExchangesOp{}

func (o stockMetaExchangesOp) Summary() string {
	return "Get exchange codes"
}

func (o stockMetaExchangesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["StockMetaExchanges"]
}

func (o stockMetaExchangesOp) RequiredFlags() []string {
	return nil
}

func (o stockMetaExchangesOp) Flags() []FlagDef {
	return nil
}

type stockQuoteSingleOp struct{}

var StockQuoteSingleOp = stockQuoteSingleOp{}

func (o stockQuoteSingleOp) Summary() string {
	return "Get historical quotes (single symbol)"
}

func (o stockQuoteSingleOp) ResponseFields() []ResponseField {
	return ResponseSchemas["StockQuoteSingle"]
}

func (o stockQuoteSingleOp) RequiredFlags() []string {
	return nil
}

func (o stockQuoteSingleOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", OpName: "StockQuoteSingle"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockQuoteSingle"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "StockQuoteSingle"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "source feed of the data.", Completions: []string{"boats", "iex", "otc", "sip"}, OpName: "StockQuoteSingle"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "StockQuoteSingle"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "StockQuoteSingle"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "StockQuoteSingle"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "StockQuoteSingle"},
	}
}

type stockQuotesOp struct{}

var StockQuotesOp = stockQuotesOp{}

func (o stockQuotesOp) Summary() string {
	return "Get historical quotes"
}

func (o stockQuotesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["StockQuotes"]
}

func (o stockQuotesOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o stockQuotesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", OpName: "StockQuotes"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockQuotes"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "StockQuotes"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "source feed of the data.", Completions: []string{"boats", "iex", "otc", "sip"}, OpName: "StockQuotes"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "StockQuotes"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "StockQuotes"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "StockQuotes"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "StockQuotes"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", OpName: "StockQuotes", Required: true},
	}
}

type stockSnapshotSingleOp struct{}

var StockSnapshotSingleOp = stockSnapshotSingleOp{}

func (o stockSnapshotSingleOp) Summary() string {
	return "Get snapshot (single symbol)"
}

func (o stockSnapshotSingleOp) ResponseFields() []ResponseField {
	return ResponseSchemas["StockSnapshotSingle"]
}

func (o stockSnapshotSingleOp) RequiredFlags() []string {
	return nil
}

func (o stockSnapshotSingleOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockSnapshotSingle"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, OpName: "StockSnapshotSingle"},
	}
}

type stockSnapshotsOp struct{}

var StockSnapshotsOp = stockSnapshotsOp{}

func (o stockSnapshotsOp) Summary() string {
	return "Get snapshots"
}

func (o stockSnapshotsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["StockSnapshots"]
}

func (o stockSnapshotsOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o stockSnapshotsOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockSnapshots"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, OpName: "StockSnapshots"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", OpName: "StockSnapshots", Required: true},
	}
}

type stockTradeSingleOp struct{}

var StockTradeSingleOp = stockTradeSingleOp{}

func (o stockTradeSingleOp) Summary() string {
	return "Get historical trades (single symbol)"
}

func (o stockTradeSingleOp) ResponseFields() []ResponseField {
	return ResponseSchemas["StockTradeSingle"]
}

func (o stockTradeSingleOp) RequiredFlags() []string {
	return nil
}

func (o stockTradeSingleOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", OpName: "StockTradeSingle"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockTradeSingle"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "StockTradeSingle"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "source feed of the data.", Completions: []string{"boats", "iex", "otc", "sip"}, OpName: "StockTradeSingle"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "StockTradeSingle"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "StockTradeSingle"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "StockTradeSingle"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "StockTradeSingle"},
	}
}

type stockTradesOp struct{}

var StockTradesOp = stockTradesOp{}

func (o stockTradesOp) Summary() string {
	return "Get historical trades"
}

func (o stockTradesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["StockTrades"]
}

func (o stockTradesOp) RequiredFlags() []string {
	return []string{"symbols"}
}

func (o stockTradesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", OpName: "StockTrades"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", OpName: "StockTrades"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "StockTrades"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "source feed of the data.", Completions: []string{"boats", "iex", "otc", "sip"}, OpName: "StockTrades"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "StockTrades"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "StockTrades"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "StockTrades"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "StockTrades"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", OpName: "StockTrades", Required: true},
	}
}

type usCorporatesOp struct{}

var UsCorporatesOp = usCorporatesOp{}

func (o usCorporatesOp) Summary() string {
	return "Get US corporates"
}

func (o usCorporatesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["UsCorporates"]
}

func (o usCorporatesOp) RequiredFlags() []string {
	return nil
}

func (o usCorporatesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "bond-status", OASName: "bond_status", Type: "string", Description: "status of the bond", Completions: []string{"matured", "outstanding", "pre_issuance"}, OpName: "UsCorporates"},
		{Name: "cusips", OASName: "cusips", Type: "string", Description: "A comma-separated list of CUSIPs with a limit of 1000", OpName: "UsCorporates"},
		{Name: "isins", OASName: "isins", Type: "string", Description: "A comma-separated list of ISINs with a limit of 1000", OpName: "UsCorporates"},
		{Name: "tickers", OASName: "tickers", Type: "string", Description: "A comma-separated list of tickers with a limit of 1000", OpName: "UsCorporates"},
	}
}

type usTreasuriesOp struct{}

var UsTreasuriesOp = usTreasuriesOp{}

func (o usTreasuriesOp) Summary() string {
	return "Get US treasuries"
}

func (o usTreasuriesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["UsTreasuries"]
}

func (o usTreasuriesOp) RequiredFlags() []string {
	return nil
}

func (o usTreasuriesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "bond-status", OASName: "bond_status", Type: "string", Description: "status of the bond", Completions: []string{"matured", "outstanding", "pre_issuance"}, OpName: "UsTreasuries"},
		{Name: "cusips", OASName: "cusips", Type: "string", Description: "A comma-separated list of CUSIPs with a limit of 1000", OpName: "UsTreasuries"},
		{Name: "isins", OASName: "isins", Type: "string", Description: "A comma-separated list of ISINs with a limit of 1000", OpName: "UsTreasuries"},
		{Name: "subtype", OASName: "subtype", Type: "string", Description: "subtype of the treasury", Completions: []string{"bill", "bond", "floating", "note", "strips", "tips"}, OpName: "UsTreasuries"},
	}
}

type addAssetToWatchlistOp struct{}

var AddAssetToWatchlistOp = addAssetToWatchlistOp{}

func (o addAssetToWatchlistOp) Summary() string {
	return "Add asset to watchlist"
}

func (o addAssetToWatchlistOp) ResponseFields() []ResponseField {
	return ResponseSchemas["AddAssetToWatchlist"]
}

func (o addAssetToWatchlistOp) RequiredFlags() []string {
	return nil
}

func (o addAssetToWatchlistOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "the symbol name to add to the watchlist", OpName: "AddAssetToWatchlist"},
	}
}

type addAssetToWatchlistByNameOp struct{}

var AddAssetToWatchlistByNameOp = addAssetToWatchlistByNameOp{}

func (o addAssetToWatchlistByNameOp) Summary() string {
	return "Add asset to watchlist by name"
}

func (o addAssetToWatchlistByNameOp) ResponseFields() []ResponseField {
	return ResponseSchemas["AddAssetToWatchlistByName"]
}

func (o addAssetToWatchlistByNameOp) RequiredFlags() []string {
	return []string{"name"}
}

func (o addAssetToWatchlistByNameOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "name", OASName: "name", Type: "string", Description: "name of the watchlist", OpName: "AddAssetToWatchlistByName", Required: true},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "the symbol name to add to the watchlist", OpName: "AddAssetToWatchlistByName"},
	}
}

type createCryptoPerpTransferForAccountOp struct{}

var CreateCryptoPerpTransferForAccountOp = createCryptoPerpTransferForAccountOp{}

func (o createCryptoPerpTransferForAccountOp) Summary() string {
	return "Request a new withdrawal"
}

func (o createCryptoPerpTransferForAccountOp) ResponseFields() []ResponseField {
	return ResponseSchemas["CreateCryptoPerpTransferForAccount"]
}

func (o createCryptoPerpTransferForAccountOp) RequiredFlags() []string {
	return nil
}

func (o createCryptoPerpTransferForAccountOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "address", OASName: "address", Type: "string", Description: "destination wallet address", OpName: "CreateCryptoPerpTransferForAccount"},
		{Name: "amount", OASName: "amount", Type: "string", Description: "amount, denoted in the specified asset, to be withdrawn from the user’s wallet", OpName: "CreateCryptoPerpTransferForAccount"},
		{Name: "asset", OASName: "asset", Type: "string", Description: "asset", OpName: "CreateCryptoPerpTransferForAccount"},
	}
}

type createCryptoTransferForAccountOp struct{}

var CreateCryptoTransferForAccountOp = createCryptoTransferForAccountOp{}

func (o createCryptoTransferForAccountOp) Summary() string {
	return "Request a new withdrawal"
}

func (o createCryptoTransferForAccountOp) ResponseFields() []ResponseField {
	return ResponseSchemas["CreateCryptoTransferForAccount"]
}

func (o createCryptoTransferForAccountOp) RequiredFlags() []string {
	return nil
}

func (o createCryptoTransferForAccountOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "address", OASName: "address", Type: "string", Description: "destination wallet address", OpName: "CreateCryptoTransferForAccount"},
		{Name: "amount", OASName: "amount", Type: "string", Description: "amount, denoted in the specified asset, to be withdrawn from the user’s wallet", OpName: "CreateCryptoTransferForAccount"},
		{Name: "asset", OASName: "asset", Type: "string", Description: "asset", OpName: "CreateCryptoTransferForAccount"},
	}
}

type createWhitelistedAddressOp struct{}

var CreateWhitelistedAddressOp = createWhitelistedAddressOp{}

func (o createWhitelistedAddressOp) Summary() string {
	return "Request a new whitelisted address"
}

func (o createWhitelistedAddressOp) ResponseFields() []ResponseField {
	return ResponseSchemas["CreateWhitelistedAddress"]
}

func (o createWhitelistedAddressOp) RequiredFlags() []string {
	return nil
}

func (o createWhitelistedAddressOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "address", OASName: "address", Type: "string", Description: "address to be whitelisted", OpName: "CreateWhitelistedAddress"},
		{Name: "asset", OASName: "asset", Type: "string", Description: "symbol of underlying asset for the whitelisted address", OpName: "CreateWhitelistedAddress"},
	}
}

type createWhitelistedPerpAddressOp struct{}

var CreateWhitelistedPerpAddressOp = createWhitelistedPerpAddressOp{}

func (o createWhitelistedPerpAddressOp) Summary() string {
	return "Request a new whitelisted address"
}

func (o createWhitelistedPerpAddressOp) ResponseFields() []ResponseField {
	return ResponseSchemas["CreateWhitelistedPerpAddress"]
}

func (o createWhitelistedPerpAddressOp) RequiredFlags() []string {
	return nil
}

func (o createWhitelistedPerpAddressOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "address", OASName: "address", Type: "string", Description: "address to be whitelisted", OpName: "CreateWhitelistedPerpAddress"},
		{Name: "asset", OASName: "asset", Type: "string", Description: "symbol of underlying asset for the whitelisted address", OpName: "CreateWhitelistedPerpAddress"},
	}
}

type deleteAllOpenPositionsOp struct{}

var DeleteAllOpenPositionsOp = deleteAllOpenPositionsOp{}

func (o deleteAllOpenPositionsOp) Summary() string {
	return "Close all positions"
}

func (o deleteAllOpenPositionsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["DeleteAllOpenPositions"]
}

func (o deleteAllOpenPositionsOp) RequiredFlags() []string {
	return nil
}

func (o deleteAllOpenPositionsOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "cancel-orders", OASName: "cancel_orders", Type: "bool", Description: "if true is specified, cancel all open orders before liquidating all positions", OpName: "DeleteAllOpenPositions"},
	}
}

type deleteAllOrdersOp struct{}

var DeleteAllOrdersOp = deleteAllOrdersOp{}

func (o deleteAllOrdersOp) Summary() string {
	return "Delete all orders"
}

func (o deleteAllOrdersOp) ResponseFields() []ResponseField {
	return ResponseSchemas["DeleteAllOrders"]
}

func (o deleteAllOrdersOp) RequiredFlags() []string {
	return nil
}

func (o deleteAllOrdersOp) Flags() []FlagDef {
	return nil
}

type deleteOpenPositionOp struct{}

var DeleteOpenPositionOp = deleteOpenPositionOp{}

func (o deleteOpenPositionOp) Summary() string {
	return "Close a position"
}

func (o deleteOpenPositionOp) ResponseFields() []ResponseField {
	return ResponseSchemas["DeleteOpenPosition"]
}

func (o deleteOpenPositionOp) RequiredFlags() []string {
	return nil
}

func (o deleteOpenPositionOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "percentage", OASName: "percentage", Type: "string", Description: "percentage of position to liquidate", OpName: "DeleteOpenPosition"},
		{Name: "qty", OASName: "qty", Type: "string", Description: "the number of shares to liquidate. Can accept up to 9 decimal points. Cannot work with percentage", OpName: "DeleteOpenPosition"},
	}
}

type deleteOrderByOrderIDOp struct{}

var DeleteOrderByOrderIDOp = deleteOrderByOrderIDOp{}

func (o deleteOrderByOrderIDOp) Summary() string {
	return "Delete order by ID"
}

func (o deleteOrderByOrderIDOp) ResponseFields() []ResponseField {
	return ResponseSchemas["DeleteOrderByOrderID"]
}

func (o deleteOrderByOrderIDOp) RequiredFlags() []string {
	return nil
}

func (o deleteOrderByOrderIDOp) Flags() []FlagDef {
	return nil
}

type deleteWatchlistByIDOp struct{}

var DeleteWatchlistByIDOp = deleteWatchlistByIDOp{}

func (o deleteWatchlistByIDOp) Summary() string {
	return "Delete watchlist by id"
}

func (o deleteWatchlistByIDOp) ResponseFields() []ResponseField {
	return ResponseSchemas["DeleteWatchlistByID"]
}

func (o deleteWatchlistByIDOp) RequiredFlags() []string {
	return nil
}

func (o deleteWatchlistByIDOp) Flags() []FlagDef {
	return nil
}

type deleteWatchlistByNameOp struct{}

var DeleteWatchlistByNameOp = deleteWatchlistByNameOp{}

func (o deleteWatchlistByNameOp) Summary() string {
	return "Delete watchlist by name"
}

func (o deleteWatchlistByNameOp) ResponseFields() []ResponseField {
	return ResponseSchemas["DeleteWatchlistByName"]
}

func (o deleteWatchlistByNameOp) RequiredFlags() []string {
	return []string{"name"}
}

func (o deleteWatchlistByNameOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "name", OASName: "name", Type: "string", Description: "name of the watchlist", OpName: "DeleteWatchlistByName", Required: true},
	}
}

type deleteWhitelistedAddressOp struct{}

var DeleteWhitelistedAddressOp = deleteWhitelistedAddressOp{}

func (o deleteWhitelistedAddressOp) Summary() string {
	return "Delete a whitelisted address"
}

func (o deleteWhitelistedAddressOp) ResponseFields() []ResponseField {
	return ResponseSchemas["DeleteWhitelistedAddress"]
}

func (o deleteWhitelistedAddressOp) RequiredFlags() []string {
	return nil
}

func (o deleteWhitelistedAddressOp) Flags() []FlagDef {
	return nil
}

type deleteWhitelistedPerpAddressOp struct{}

var DeleteWhitelistedPerpAddressOp = deleteWhitelistedPerpAddressOp{}

func (o deleteWhitelistedPerpAddressOp) Summary() string {
	return "Delete a whitelisted address"
}

func (o deleteWhitelistedPerpAddressOp) ResponseFields() []ResponseField {
	return ResponseSchemas["DeleteWhitelistedPerpAddress"]
}

func (o deleteWhitelistedPerpAddressOp) RequiredFlags() []string {
	return nil
}

func (o deleteWhitelistedPerpAddressOp) Flags() []FlagDef {
	return nil
}

type getOptionContractSymbolOrIDOp struct{}

var GetOptionContractSymbolOrIDOp = getOptionContractSymbolOrIDOp{}

func (o getOptionContractSymbolOrIDOp) Summary() string {
	return "Get an option contract by ID or symbol"
}

func (o getOptionContractSymbolOrIDOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetOptionContractSymbolOrID"]
}

func (o getOptionContractSymbolOrIDOp) RequiredFlags() []string {
	return nil
}

func (o getOptionContractSymbolOrIDOp) Flags() []FlagDef {
	return nil
}

type getOptionsContractsOp struct{}

var GetOptionsContractsOp = getOptionsContractsOp{}

func (o getOptionsContractsOp) Summary() string {
	return "Get option contracts"
}

func (o getOptionsContractsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetOptionsContracts"]
}

func (o getOptionsContractsOp) RequiredFlags() []string {
	return nil
}

func (o getOptionsContractsOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "expiration-date", OASName: "expiration_date", Type: "string", Description: "filter contracts by the exact expiration date (format: YYYY-MM-DD)", OpName: "GetOptionsContracts"},
		{Name: "expiration-date-gte", OASName: "expiration_date_gte", Type: "string", Description: "filter contracts with expiration date greater than or equal to the specified date", OpName: "GetOptionsContracts"},
		{Name: "expiration-date-lte", OASName: "expiration_date_lte", Type: "string", Description: "filter contracts with expiration date less than or equal to the specified date", OpName: "GetOptionsContracts"},
		{Name: "limit", OASName: "limit", Type: "int", Description: "number of contracts to limit per page (default=100, max=10000)", OpName: "GetOptionsContracts"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "used for pagination, this token retrieves the next page of results", OpName: "GetOptionsContracts"},
		{Name: "ppind", OASName: "ppind", Type: "bool", Description: "ppind(Penny Program Indicator) field indicates whether an option contract is eligible for penny price increments,", OpName: "GetOptionsContracts"},
		{Name: "root-symbol", OASName: "root_symbol", Type: "string", Description: "filter contracts by the root symbol", OpName: "GetOptionsContracts"},
		{Name: "show-deliverables", OASName: "show_deliverables", Type: "bool", Description: "include deliverables array in the response", OpName: "GetOptionsContracts"},
		{Name: "status", OASName: "status", Type: "string", Description: "filter contracts by status (active/inactive). By default only active contracts are returned", Completions: []string{"active", "inactive"}, OpName: "GetOptionsContracts"},
		{Name: "strike-price-gte", OASName: "strike_price_gte", Type: "string", Description: "filter contracts with strike price greater than or equal to the specified value", OpName: "GetOptionsContracts"},
		{Name: "strike-price-lte", OASName: "strike_price_lte", Type: "string", Description: "filter contracts with strike price less than or equal to the specified value", OpName: "GetOptionsContracts"},
		{Name: "style", OASName: "style", Type: "string", Description: "filter contracts by the style (american/european)", Completions: []string{"american", "european"}, OpName: "GetOptionsContracts"},
		{Name: "type", OASName: "type", Type: "string", Description: "filter contracts by the type (call/put)", Completions: []string{"call", "put"}, OpName: "GetOptionsContracts"},
		{Name: "underlying-symbols", OASName: "underlying_symbols", Type: "string", Description: "filter contracts by one or more underlying symbols", OpName: "GetOptionsContracts"},
	}
}

type getV2AssetsOp struct{}

var GetV2AssetsOp = getV2AssetsOp{}

func (o getV2AssetsOp) Summary() string {
	return "Get assets"
}

func (o getV2AssetsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetV2Assets"]
}

func (o getV2AssetsOp) RequiredFlags() []string {
	return nil
}

func (o getV2AssetsOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "asset-class", OASName: "asset_class", Type: "string", Description: "defaults to us_equity", OpName: "GetV2Assets"},
		{Name: "attributes", OASName: "attributes", Type: "string", Default: "[]", Description: "comma separated values to query for more than one attribute", OpName: "GetV2Assets"},
		{Name: "exchange", OASName: "exchange", Type: "string", Description: "optional AMEX, ARCA, BATS, NYSE, NASDAQ, NYSEARCA or OTC", OpName: "GetV2Assets"},
		{Name: "status", OASName: "status", Type: "string", Description: "e.g. “active”. By default, all statuses are included", OpName: "GetV2Assets"},
	}
}

type getV2AssetsSymbolOrAssetIDOp struct{}

var GetV2AssetsSymbolOrAssetIDOp = getV2AssetsSymbolOrAssetIDOp{}

func (o getV2AssetsSymbolOrAssetIDOp) Summary() string {
	return "Get an asset by ID or symbol"
}

func (o getV2AssetsSymbolOrAssetIDOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetV2AssetsSymbolOrAssetID"]
}

func (o getV2AssetsSymbolOrAssetIDOp) RequiredFlags() []string {
	return nil
}

func (o getV2AssetsSymbolOrAssetIDOp) Flags() []FlagDef {
	return nil
}

type getV2CorporateActionsAnnouncementsOp struct{}

var GetV2CorporateActionsAnnouncementsOp = getV2CorporateActionsAnnouncementsOp{}

func (o getV2CorporateActionsAnnouncementsOp) Summary() string {
	return "Retrieve announcements"
}

func (o getV2CorporateActionsAnnouncementsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetV2CorporateActionsAnnouncements"]
}

func (o getV2CorporateActionsAnnouncementsOp) RequiredFlags() []string {
	return []string{"ca-types", "since", "until"}
}

func (o getV2CorporateActionsAnnouncementsOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "ca-types", OASName: "ca_types", Type: "string", Description: "A comma-delimited list of Dividend, Merger, Spinoff, or Split", OpName: "GetV2CorporateActionsAnnouncements", Required: true},
		{Name: "cusip", OASName: "cusip", Type: "string", Description: "CUSIP of the company initiating the announcement", OpName: "GetV2CorporateActionsAnnouncements"},
		{Name: "date-type", OASName: "date_type", Type: "string", Description: "declaration_date, ex_date, record_date, or payable_date", OpName: "GetV2CorporateActionsAnnouncements"},
		{Name: "since", OASName: "since", Type: "string", Description: "start (inclusive) of the date range when searching corporate action announcements", OpName: "GetV2CorporateActionsAnnouncements", Required: true},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol of the company initiating the announcement", OpName: "GetV2CorporateActionsAnnouncements"},
		{Name: "until", OASName: "until", Type: "string", Description: "end (inclusive) of the date range when searching corporate action announcements", OpName: "GetV2CorporateActionsAnnouncements", Required: true},
	}
}

type getV2CorporateActionsAnnouncementsIDOp struct{}

var GetV2CorporateActionsAnnouncementsIDOp = getV2CorporateActionsAnnouncementsIDOp{}

func (o getV2CorporateActionsAnnouncementsIDOp) Summary() string {
	return "Retrieve a specific announcement"
}

func (o getV2CorporateActionsAnnouncementsIDOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetV2CorporateActionsAnnouncementsID"]
}

func (o getV2CorporateActionsAnnouncementsIDOp) RequiredFlags() []string {
	return nil
}

func (o getV2CorporateActionsAnnouncementsIDOp) Flags() []FlagDef {
	return nil
}

type getAccountOp struct{}

var GetAccountOp = getAccountOp{}

func (o getAccountOp) Summary() string {
	return "Get account"
}

func (o getAccountOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetAccount"]
}

func (o getAccountOp) RequiredFlags() []string {
	return nil
}

func (o getAccountOp) Flags() []FlagDef {
	return nil
}

type getAccountActivitiesOp struct{}

var GetAccountActivitiesOp = getAccountActivitiesOp{}

func (o getAccountActivitiesOp) Summary() string {
	return "Retrieve account activities"
}

func (o getAccountActivitiesOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetAccountActivities"]
}

func (o getAccountActivitiesOp) RequiredFlags() []string {
	return nil
}

func (o getAccountActivitiesOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "activity-types", OASName: "activity_types", Type: "string", Description: "A comma-separated list of activity types used to filter the results", OpName: "GetAccountActivities"},
		{Name: "after", OASName: "after", Type: "string", Description: "get activities created after this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported", OpName: "GetAccountActivities"},
		{Name: "category", OASName: "category", Type: "string", Description: "activity category. Cannot be used with \"activity_types\" parameter", Completions: []string{"non_trade_activity", "trade_activity"}, OpName: "GetAccountActivities"},
		{Name: "date", OASName: "date", Type: "string", Description: "filter activities by the activity date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported", OpName: "GetAccountActivities"},
		{Name: "direction", OASName: "direction", Type: "string", Default: "desc", Description: "chronological order of response based on the activity datetime", Completions: []string{"asc", "desc"}, OpName: "GetAccountActivities"},
		{Name: "page-size", OASName: "page_size", Type: "int", Default: "100", Description: "maximum number of entries to return in the response", OpName: "GetAccountActivities"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "token used for pagination. Provide the ID of the last activity from the last page to retrieve the next set of results", OpName: "GetAccountActivities"},
		{Name: "until", OASName: "until", Type: "string", Description: "get activities created before this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported", OpName: "GetAccountActivities"},
	}
}

type getAccountActivitiesByActivityTypeOp struct{}

var GetAccountActivitiesByActivityTypeOp = getAccountActivitiesByActivityTypeOp{}

func (o getAccountActivitiesByActivityTypeOp) Summary() string {
	return "Retrieve account activities of specific type"
}

func (o getAccountActivitiesByActivityTypeOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetAccountActivitiesByActivityType"]
}

func (o getAccountActivitiesByActivityTypeOp) RequiredFlags() []string {
	return nil
}

func (o getAccountActivitiesByActivityTypeOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "after", OASName: "after", Type: "string", Description: "get activities created after this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported", OpName: "GetAccountActivitiesByActivityType"},
		{Name: "date", OASName: "date", Type: "string", Description: "filter activities by the activity date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported", OpName: "GetAccountActivitiesByActivityType"},
		{Name: "direction", OASName: "direction", Type: "string", Default: "desc", Description: "chronological order of response based on the activity datetime", Completions: []string{"asc", "desc"}, OpName: "GetAccountActivitiesByActivityType"},
		{Name: "page-size", OASName: "page_size", Type: "int", Default: "100", Description: "maximum number of entries to return in the response", OpName: "GetAccountActivitiesByActivityType"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "token used for pagination. Provide the ID of the last activity from the last page to retrieve the next set of results", OpName: "GetAccountActivitiesByActivityType"},
		{Name: "until", OASName: "until", Type: "string", Description: "get activities created before this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported", OpName: "GetAccountActivitiesByActivityType"},
	}
}

type getAccountConfigOp struct{}

var GetAccountConfigOp = getAccountConfigOp{}

func (o getAccountConfigOp) Summary() string {
	return "Get account configurations"
}

func (o getAccountConfigOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetAccountConfig"]
}

func (o getAccountConfigOp) RequiredFlags() []string {
	return nil
}

func (o getAccountConfigOp) Flags() []FlagDef {
	return nil
}

type getAccountPortfolioHistoryOp struct{}

var GetAccountPortfolioHistoryOp = getAccountPortfolioHistoryOp{}

func (o getAccountPortfolioHistoryOp) Summary() string {
	return "Get account portfolio history"
}

func (o getAccountPortfolioHistoryOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetAccountPortfolioHistory"]
}

func (o getAccountPortfolioHistoryOp) RequiredFlags() []string {
	return nil
}

func (o getAccountPortfolioHistoryOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "cashflow-types", OASName: "cashflow_types", Type: "string", Description: "cashflow activities to include in the report. One of 'ALL', 'NONE', or a comma-separated list of activity types", OpName: "GetAccountPortfolioHistory"},
		{Name: "end", OASName: "end", Type: "string", Description: "timestamp the data is returned up to in RFC3339 format (including timezone specification)", OpName: "GetAccountPortfolioHistory"},
		{Name: "extended-hours", OASName: "extended_hours", Type: "string", Description: "**deprecated**: Users are strongly advised to **rely on the intraday_reporting query parameter** for better control\no...", OpName: "GetAccountPortfolioHistory"},
		{Name: "intraday-reporting", OASName: "intraday_reporting", Type: "string", Default: "market_hours", Description: "for intraday resolutions (<1D) this specfies which timestamps to return data points for:\n\nAllowed values are:\n- **mar...", Completions: []string{"continuous", "extended_hours", "market_hours"}, OpName: "GetAccountPortfolioHistory"},
		{Name: "period", OASName: "period", Type: "string", Description: "duration of the data in number + unit format, such as 1D, where unit can be D for day, W for week, M for month and A ...", OpName: "GetAccountPortfolioHistory"},
		{Name: "pnl-reset", OASName: "pnl_reset", Type: "string", Default: "per_day", Description: "pnl_reset defines how we are calculating the baseline values for Profit And Loss (pnl) for queries with timeframe les...", Completions: []string{"no_reset", "per_day"}, OpName: "GetAccountPortfolioHistory"},
		{Name: "start", OASName: "start", Type: "string", Description: "timestamp the data is returned starting from in RFC3339 format (including timezone specification)", OpName: "GetAccountPortfolioHistory"},
		{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "resolution of time window", OpName: "GetAccountPortfolioHistory"},
	}
}

type getAllOpenPositionsOp struct{}

var GetAllOpenPositionsOp = getAllOpenPositionsOp{}

func (o getAllOpenPositionsOp) Summary() string {
	return "List all open positions"
}

func (o getAllOpenPositionsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetAllOpenPositions"]
}

func (o getAllOpenPositionsOp) RequiredFlags() []string {
	return nil
}

func (o getAllOpenPositionsOp) Flags() []FlagDef {
	return nil
}

type getAllOrdersOp struct{}

var GetAllOrdersOp = getAllOrdersOp{}

func (o getAllOrdersOp) Summary() string {
	return "Get all orders"
}

func (o getAllOrdersOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetAllOrders"]
}

func (o getAllOrdersOp) RequiredFlags() []string {
	return nil
}

func (o getAllOrdersOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "after", OASName: "after", Type: "string", Description: "response will include only ones submitted after this timestamp (exclusive.)", OpName: "GetAllOrders"},
		{Name: "after-order-id", OASName: "after_order_id", Type: "string", Description: "return orders submitted after the order with this ID (exclusive).\nMutually exclusive with before_order_id", OpName: "GetAllOrders"},
		{Name: "asset-class", OASName: "asset_class", Type: "string", Description: "A comma seperated list of asset classes, the response will include only orders in the specified asset classes", OpName: "GetAllOrders"},
		{Name: "before-order-id", OASName: "before_order_id", Type: "string", Description: "return orders submitted before the order with this ID (exclusive).\nMutually exclusive with after_order_id", OpName: "GetAllOrders"},
		{Name: "direction", OASName: "direction", Type: "string", Description: "chronological order of response based on the submission time. asc or desc. Defaults to desc", Completions: []string{"asc", "desc"}, OpName: "GetAllOrders"},
		{Name: "limit", OASName: "limit", Type: "int", Description: "maximum number of orders in response. Defaults to 50 and max is 500", OpName: "GetAllOrders"},
		{Name: "nested", OASName: "nested", Type: "bool", Description: "if true, the result will roll up multi-leg orders under the legs field of primary order", OpName: "GetAllOrders"},
		{Name: "side", OASName: "side", Type: "string", Description: "filters down to orders that have a matching side field set", OpName: "GetAllOrders"},
		{Name: "status", OASName: "status", Type: "string", Description: "order status to be queried. open, closed or all. Defaults to open", Completions: []string{"all", "closed", "open"}, OpName: "GetAllOrders"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of symbols to filter by (ex", OpName: "GetAllOrders"},
		{Name: "until", OASName: "until", Type: "string", Description: "response will include only ones submitted until this timestamp (exclusive.)", OpName: "GetAllOrders"},
	}
}

type getCryptoFundingTransferOp struct{}

var GetCryptoFundingTransferOp = getCryptoFundingTransferOp{}

func (o getCryptoFundingTransferOp) Summary() string {
	return "Retrieve a crypto funding transfer"
}

func (o getCryptoFundingTransferOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetCryptoFundingTransfer"]
}

func (o getCryptoFundingTransferOp) RequiredFlags() []string {
	return nil
}

func (o getCryptoFundingTransferOp) Flags() []FlagDef {
	return nil
}

type getCryptoPerpAccountLeverageOp struct{}

var GetCryptoPerpAccountLeverageOp = getCryptoPerpAccountLeverageOp{}

func (o getCryptoPerpAccountLeverageOp) Summary() string {
	return "Get account leverage for an asset"
}

func (o getCryptoPerpAccountLeverageOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetCryptoPerpAccountLeverage"]
}

func (o getCryptoPerpAccountLeverageOp) RequiredFlags() []string {
	return nil
}

func (o getCryptoPerpAccountLeverageOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol of underlying asset", OpName: "GetCryptoPerpAccountLeverage"},
	}
}

type getCryptoPerpAccountVitalsOp struct{}

var GetCryptoPerpAccountVitalsOp = getCryptoPerpAccountVitalsOp{}

func (o getCryptoPerpAccountVitalsOp) Summary() string {
	return "Retrieve account vitals"
}

func (o getCryptoPerpAccountVitalsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetCryptoPerpAccountVitals"]
}

func (o getCryptoPerpAccountVitalsOp) RequiredFlags() []string {
	return nil
}

func (o getCryptoPerpAccountVitalsOp) Flags() []FlagDef {
	return nil
}

type getCryptoPerpFundingTransferOp struct{}

var GetCryptoPerpFundingTransferOp = getCryptoPerpFundingTransferOp{}

func (o getCryptoPerpFundingTransferOp) Summary() string {
	return "Retrieve a crypto funding transfer"
}

func (o getCryptoPerpFundingTransferOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetCryptoPerpFundingTransfer"]
}

func (o getCryptoPerpFundingTransferOp) RequiredFlags() []string {
	return nil
}

func (o getCryptoPerpFundingTransferOp) Flags() []FlagDef {
	return nil
}

type getCryptoPerpTransferEstimateOp struct{}

var GetCryptoPerpTransferEstimateOp = getCryptoPerpTransferEstimateOp{}

func (o getCryptoPerpTransferEstimateOp) Summary() string {
	return "Returns the estimated gas fee for a proposed transaction"
}

func (o getCryptoPerpTransferEstimateOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetCryptoPerpTransferEstimate"]
}

func (o getCryptoPerpTransferEstimateOp) RequiredFlags() []string {
	return nil
}

func (o getCryptoPerpTransferEstimateOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "amount", OASName: "amount", Type: "string", Description: "amount, denoted in the specified asset, of the proposed transaction", OpName: "GetCryptoPerpTransferEstimate"},
		{Name: "asset", OASName: "asset", Type: "string", Description: "asset for the proposed transaction", OpName: "GetCryptoPerpTransferEstimate"},
		{Name: "from-address", OASName: "from_address", Type: "string", Description: "originating address of the proposed transaction", OpName: "GetCryptoPerpTransferEstimate"},
		{Name: "to-address", OASName: "to_address", Type: "string", Description: "destination address of the proposed transaction", OpName: "GetCryptoPerpTransferEstimate"},
	}
}

type getCryptoTransferEstimateOp struct{}

var GetCryptoTransferEstimateOp = getCryptoTransferEstimateOp{}

func (o getCryptoTransferEstimateOp) Summary() string {
	return "Returns the estimated gas fee for a proposed transaction"
}

func (o getCryptoTransferEstimateOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetCryptoTransferEstimate"]
}

func (o getCryptoTransferEstimateOp) RequiredFlags() []string {
	return nil
}

func (o getCryptoTransferEstimateOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "amount", OASName: "amount", Type: "string", Description: "amount, denoted in the specified asset, of the proposed transaction", OpName: "GetCryptoTransferEstimate"},
		{Name: "asset", OASName: "asset", Type: "string", Description: "asset for the proposed transaction", OpName: "GetCryptoTransferEstimate"},
		{Name: "from-address", OASName: "from_address", Type: "string", Description: "originating address of the proposed transaction", OpName: "GetCryptoTransferEstimate"},
		{Name: "to-address", OASName: "to_address", Type: "string", Description: "destination address of the proposed transaction", OpName: "GetCryptoTransferEstimate"},
	}
}

type getOpenPositionOp struct{}

var GetOpenPositionOp = getOpenPositionOp{}

func (o getOpenPositionOp) Summary() string {
	return "Get an open position"
}

func (o getOpenPositionOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetOpenPosition"]
}

func (o getOpenPositionOp) RequiredFlags() []string {
	return nil
}

func (o getOpenPositionOp) Flags() []FlagDef {
	return nil
}

type getOrderByClientOrderIDOp struct{}

var GetOrderByClientOrderIDOp = getOrderByClientOrderIDOp{}

func (o getOrderByClientOrderIDOp) Summary() string {
	return "Get order by client order ID"
}

func (o getOrderByClientOrderIDOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetOrderByClientOrderID"]
}

func (o getOrderByClientOrderIDOp) RequiredFlags() []string {
	return []string{"client-order-id"}
}

func (o getOrderByClientOrderIDOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "client-order-id", OASName: "client_order_id", Type: "string", Description: "client-assigned order ID", OpName: "GetOrderByClientOrderID", Required: true},
	}
}

type getOrderByOrderIDOp struct{}

var GetOrderByOrderIDOp = getOrderByOrderIDOp{}

func (o getOrderByOrderIDOp) Summary() string {
	return "Get order by ID"
}

func (o getOrderByOrderIDOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetOrderByOrderID"]
}

func (o getOrderByOrderIDOp) RequiredFlags() []string {
	return nil
}

func (o getOrderByOrderIDOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "nested", OASName: "nested", Type: "bool", Description: "if true, the result will roll up multi-leg orders under the legs field of primary order", OpName: "GetOrderByOrderID"},
	}
}

type getWatchlistByIDOp struct{}

var GetWatchlistByIDOp = getWatchlistByIDOp{}

func (o getWatchlistByIDOp) Summary() string {
	return "Get watchlist by ID"
}

func (o getWatchlistByIDOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetWatchlistByID"]
}

func (o getWatchlistByIDOp) RequiredFlags() []string {
	return nil
}

func (o getWatchlistByIDOp) Flags() []FlagDef {
	return nil
}

type getWatchlistByNameOp struct{}

var GetWatchlistByNameOp = getWatchlistByNameOp{}

func (o getWatchlistByNameOp) Summary() string {
	return "Get watchlist by name"
}

func (o getWatchlistByNameOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetWatchlistByName"]
}

func (o getWatchlistByNameOp) RequiredFlags() []string {
	return []string{"name"}
}

func (o getWatchlistByNameOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "name", OASName: "name", Type: "string", Description: "name of the watchlist", OpName: "GetWatchlistByName", Required: true},
	}
}

type getWatchlistsOp struct{}

var GetWatchlistsOp = getWatchlistsOp{}

func (o getWatchlistsOp) Summary() string {
	return "Get all watchlists"
}

func (o getWatchlistsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["GetWatchlists"]
}

func (o getWatchlistsOp) RequiredFlags() []string {
	return nil
}

func (o getWatchlistsOp) Flags() []FlagDef {
	return nil
}

type listCryptoFundingTransfersOp struct{}

var ListCryptoFundingTransfersOp = listCryptoFundingTransfersOp{}

func (o listCryptoFundingTransfersOp) Summary() string {
	return "Retrieve crypto funding transfers"
}

func (o listCryptoFundingTransfersOp) ResponseFields() []ResponseField {
	return ResponseSchemas["ListCryptoFundingTransfers"]
}

func (o listCryptoFundingTransfersOp) RequiredFlags() []string {
	return nil
}

func (o listCryptoFundingTransfersOp) Flags() []FlagDef {
	return nil
}

type listCryptoFundingWalletsOp struct{}

var ListCryptoFundingWalletsOp = listCryptoFundingWalletsOp{}

func (o listCryptoFundingWalletsOp) Summary() string {
	return "Retrieve crypto funding wallets"
}

func (o listCryptoFundingWalletsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["ListCryptoFundingWallets"]
}

func (o listCryptoFundingWalletsOp) RequiredFlags() []string {
	return nil
}

func (o listCryptoFundingWalletsOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "asset", OASName: "asset", Type: "string", Description: "asset", OpName: "ListCryptoFundingWallets"},
		{Name: "network", OASName: "network", Type: "string", Description: "optional network identifier", Completions: []string{"ethereum", "solana"}, OpName: "ListCryptoFundingWallets"},
	}
}

type listCryptoPerpFundingTransfersOp struct{}

var ListCryptoPerpFundingTransfersOp = listCryptoPerpFundingTransfersOp{}

func (o listCryptoPerpFundingTransfersOp) Summary() string {
	return "Retrieve crypto funding transfers"
}

func (o listCryptoPerpFundingTransfersOp) ResponseFields() []ResponseField {
	return ResponseSchemas["ListCryptoPerpFundingTransfers"]
}

func (o listCryptoPerpFundingTransfersOp) RequiredFlags() []string {
	return nil
}

func (o listCryptoPerpFundingTransfersOp) Flags() []FlagDef {
	return nil
}

type listCryptoPerpFundingWalletsOp struct{}

var ListCryptoPerpFundingWalletsOp = listCryptoPerpFundingWalletsOp{}

func (o listCryptoPerpFundingWalletsOp) Summary() string {
	return "Retrieve crypto funding wallets"
}

func (o listCryptoPerpFundingWalletsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["ListCryptoPerpFundingWallets"]
}

func (o listCryptoPerpFundingWalletsOp) RequiredFlags() []string {
	return nil
}

func (o listCryptoPerpFundingWalletsOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "asset", OASName: "asset", Type: "string", Description: "asset", OpName: "ListCryptoPerpFundingWallets"},
	}
}

type listWhitelistedAddressOp struct{}

var ListWhitelistedAddressOp = listWhitelistedAddressOp{}

func (o listWhitelistedAddressOp) Summary() string {
	return "List an array of whitelisted addresses"
}

func (o listWhitelistedAddressOp) ResponseFields() []ResponseField {
	return ResponseSchemas["ListWhitelistedAddress"]
}

func (o listWhitelistedAddressOp) RequiredFlags() []string {
	return nil
}

func (o listWhitelistedAddressOp) Flags() []FlagDef {
	return nil
}

type listWhitelistedPerpAddressOp struct{}

var ListWhitelistedPerpAddressOp = listWhitelistedPerpAddressOp{}

func (o listWhitelistedPerpAddressOp) Summary() string {
	return "List an array of whitelisted addresses"
}

func (o listWhitelistedPerpAddressOp) ResponseFields() []ResponseField {
	return ResponseSchemas["ListWhitelistedPerpAddress"]
}

func (o listWhitelistedPerpAddressOp) RequiredFlags() []string {
	return nil
}

func (o listWhitelistedPerpAddressOp) Flags() []FlagDef {
	return nil
}

type optionBarsOp struct{}

var OptionBarsOp = optionBarsOp{}

func (o optionBarsOp) Summary() string {
	return "Get historical bars"
}

func (o optionBarsOp) ResponseFields() []ResponseField {
	return ResponseSchemas["OptionBars"]
}

func (o optionBarsOp) RequiredFlags() []string {
	return []string{"symbols", "timeframe"}
}

func (o optionBarsOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", OpName: "OptionBars"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", OpName: "OptionBars"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", OpName: "OptionBars"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, OpName: "OptionBars"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", OpName: "OptionBars"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100", OpName: "OptionBars", Required: true},
		{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g", OpName: "OptionBars", Required: true},
	}
}

type optionDoNotExerciseOp struct{}

var OptionDoNotExerciseOp = optionDoNotExerciseOp{}

func (o optionDoNotExerciseOp) Summary() string {
	return "Do not exercise an options position"
}

func (o optionDoNotExerciseOp) ResponseFields() []ResponseField {
	return ResponseSchemas["OptionDoNotExercise"]
}

func (o optionDoNotExerciseOp) RequiredFlags() []string {
	return nil
}

func (o optionDoNotExerciseOp) Flags() []FlagDef {
	return nil
}

type optionExerciseOp struct{}

var OptionExerciseOp = optionExerciseOp{}

func (o optionExerciseOp) Summary() string {
	return "Exercise an options position"
}

func (o optionExerciseOp) ResponseFields() []ResponseField {
	return ResponseSchemas["OptionExercise"]
}

func (o optionExerciseOp) RequiredFlags() []string {
	return nil
}

func (o optionExerciseOp) Flags() []FlagDef {
	return nil
}

type patchAccountConfigOp struct{}

var PatchAccountConfigOp = patchAccountConfigOp{}

func (o patchAccountConfigOp) Summary() string {
	return "Update account configurations"
}

func (o patchAccountConfigOp) ResponseFields() []ResponseField {
	return ResponseSchemas["PatchAccountConfig"]
}

func (o patchAccountConfigOp) RequiredFlags() []string {
	return nil
}

func (o patchAccountConfigOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "disable-overnight-trading", OASName: "disable_overnight_trading", Type: "bool", Description: "if true, overnight trading is disabled", OpName: "PatchAccountConfig"},
		{Name: "dtbp-check", OASName: "dtbp_check", Type: "string", Description: "both, entry, or exit. Controls Day Trading Margin Call (DTMC) checks", Completions: []string{"both", "entry", "exit"}, OpName: "PatchAccountConfig"},
		{Name: "fractional-trading", OASName: "fractional_trading", Type: "bool", Description: "if true, account is able to participate in fractional trading", OpName: "PatchAccountConfig"},
		{Name: "max-margin-multiplier", OASName: "max_margin_multiplier", Type: "string", Description: "can be \"1\", \"2\", or \"4\"", OpName: "PatchAccountConfig"},
		{Name: "max-options-trading-level", OASName: "max_options_trading_level", Type: "int", Description: "desired maximum options trading level. 0=disabled, 1=Covered Call/Cash-Secured Put, 2=Long Call/Put, 3=Spreads/Straddles", Completions: []string{"0", "1", "2", "3"}, OpName: "PatchAccountConfig"},
		{Name: "no-shorting", OASName: "no_shorting", Type: "bool", Description: "if true, account becomes long-only mode", OpName: "PatchAccountConfig"},
		{Name: "pdt-check", OASName: "pdt_check", Type: "string", Description: "both, entry, or exit", OpName: "PatchAccountConfig"},
		{Name: "ptp-no-exception-entry", OASName: "ptp_no_exception_entry", Type: "bool", Description: "if set to true then Alpaca will accept orders for PTP symbols with no exception. Default is false", OpName: "PatchAccountConfig"},
		{Name: "suspend-trade", OASName: "suspend_trade", Type: "bool", Description: "if true, new orders are blocked", OpName: "PatchAccountConfig"},
		{Name: "trade-confirm-email", OASName: "trade_confirm_email", Type: "string", Description: "all or none. If none, emails for order fills are not sent", OpName: "PatchAccountConfig"},
	}
}

type patchOrderByOrderIDOp struct{}

var PatchOrderByOrderIDOp = patchOrderByOrderIDOp{}

func (o patchOrderByOrderIDOp) Summary() string {
	return "Replace order by ID"
}

func (o patchOrderByOrderIDOp) ResponseFields() []ResponseField {
	return ResponseSchemas["PatchOrderByOrderID"]
}

func (o patchOrderByOrderIDOp) RequiredFlags() []string {
	return nil
}

func (o patchOrderByOrderIDOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "advanced-instructions", OASName: "advanced_instructions", Type: "string", Description: "advanced instructions for Elite Smart Router: https://docs.alpaca.markets/docs/alpaca-elite-smart-router", OpName: "PatchOrderByOrderID"},
		{Name: "client-order-id", OASName: "client_order_id", Type: "string", Description: "A unique identifier for the new order. Automatically generated if not sent. (<= 128 characters)", OpName: "PatchOrderByOrderID"},
		{Name: "limit-price", OASName: "limit_price", Type: "string", Description: "required if original order's type field was limit or stop_limit.", OpName: "PatchOrderByOrderID"},
		{Name: "qty", OASName: "qty", Type: "string", Description: "number of shares to trade.", OpName: "PatchOrderByOrderID"},
		{Name: "stop-price", OASName: "stop_price", Type: "string", Description: "required if original order type is limit or stop_limit", OpName: "PatchOrderByOrderID"},
		{Name: "time-in-force", OASName: "time_in_force", Type: "string", Description: "time-In-Force values supported by Alpaca vary based on the order's security type", Completions: []string{"cls", "day", "fok", "gtc", "ioc", "opg"}, OpName: "PatchOrderByOrderID"},
		{Name: "trail", OASName: "trail", Type: "string", Description: "the new value of the trail_price or trail_percent value (works only for type=“trailing_stop”)", OpName: "PatchOrderByOrderID"},
	}
}

type postOrderOp struct{}

var PostOrderOp = postOrderOp{}

func (o postOrderOp) Summary() string {
	return "Create an order"
}

func (o postOrderOp) ResponseFields() []ResponseField {
	return ResponseSchemas["PostOrder"]
}

func (o postOrderOp) RequiredFlags() []string {
	return nil
}

func (o postOrderOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "advanced-instructions", OASName: "advanced_instructions", Type: "string", Description: "advanced instructions for Elite Smart Router: https://docs.alpaca.markets/docs/alpaca-elite-smart-router", OpName: "PostOrder"},
		{Name: "client-order-id", OASName: "client_order_id", Type: "string", Description: "A unique identifier for the order. Automatically generated if not sent. (<= 128 characters)", OpName: "PostOrder"},
		{Name: "extended-hours", OASName: "extended_hours", Type: "bool", Description: "(default) false", OpName: "PostOrder"},
		{Name: "legs", OASName: "legs", Type: "string", Description: "list of order legs (<= 4)", OpName: "PostOrder"},
		{Name: "limit-price", OASName: "limit_price", Type: "string", Description: "required if type is limit or stop_limit.", OpName: "PostOrder"},
		{Name: "notional", OASName: "notional", Type: "string", Description: "dollar amount to trade. Cannot work with qty. Can only work for market order types and day for time in force", OpName: "PostOrder"},
		{Name: "order-class", OASName: "order_class", Type: "string", Description: "order classes supported by Alpaca vary based on the order's security type", Completions: []string{"bracket", "mleg", "oco", "oto", "simple"}, OpName: "PostOrder"},
		{Name: "position-intent", OASName: "position_intent", Type: "string", Description: "represents the desired position strategy", Completions: []string{"buy_to_close", "buy_to_open", "sell_to_close", "sell_to_open"}, OpName: "PostOrder"},
		{Name: "qty", OASName: "qty", Type: "string", Description: "number of shares to trade", OpName: "PostOrder"},
		{Name: "side", OASName: "side", Type: "string", Description: "represents which side this order was on:\n- buy\n- sell\nRequired for all order classes except for mleg", Completions: []string{"buy", "sell"}, OpName: "PostOrder"},
		{Name: "stop-loss", OASName: "stop_loss", Type: "string", Description: "takes in string/number values for stop_price and limit_price", OpName: "PostOrder"},
		{Name: "stop-price", OASName: "stop_price", Type: "string", Description: "required if type is stop or stop_limit", OpName: "PostOrder"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol, asset ID, or currency pair to identify the asset to trade, required for all order classes except for mleg", OpName: "PostOrder"},
		{Name: "take-profit", OASName: "take_profit", Type: "string", Description: "takes in a string/number value for limit_price", OpName: "PostOrder"},
		{Name: "time-in-force", OASName: "time_in_force", Type: "string", Description: "time-In-Force values supported by Alpaca vary based on the order's security type", Completions: []string{"cls", "day", "fok", "gtc", "ioc", "opg"}, OpName: "PostOrder"},
		{Name: "trail-percent", OASName: "trail_percent", Type: "string", Description: "this or trail_price is required if type is trailing_stop", OpName: "PostOrder"},
		{Name: "trail-price", OASName: "trail_price", Type: "string", Description: "this or trail_percent is required if type is trailing_stop", OpName: "PostOrder"},
		{Name: "type", OASName: "type", Type: "string", Description: "order types supported by Alpaca vary based on the order's security type", Completions: []string{"limit", "market", "stop", "stop_limit", "trailing_stop"}, OpName: "PostOrder"},
	}
}

type postWatchlistOp struct{}

var PostWatchlistOp = postWatchlistOp{}

func (o postWatchlistOp) Summary() string {
	return "Create watchlist"
}

func (o postWatchlistOp) ResponseFields() []ResponseField {
	return ResponseSchemas["PostWatchlist"]
}

func (o postWatchlistOp) RequiredFlags() []string {
	return nil
}

func (o postWatchlistOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "name", OASName: "name", Type: "string", Description: "name", OpName: "PostWatchlist"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "symbols", OpName: "PostWatchlist"},
	}
}

type removeAssetFromWatchlistOp struct{}

var RemoveAssetFromWatchlistOp = removeAssetFromWatchlistOp{}

func (o removeAssetFromWatchlistOp) Summary() string {
	return "Delete symbol from watchlist"
}

func (o removeAssetFromWatchlistOp) ResponseFields() []ResponseField {
	return ResponseSchemas["RemoveAssetFromWatchlist"]
}

func (o removeAssetFromWatchlistOp) RequiredFlags() []string {
	return nil
}

func (o removeAssetFromWatchlistOp) Flags() []FlagDef {
	return nil
}

type setCryptoPerpAccountLeverageOp struct{}

var SetCryptoPerpAccountLeverageOp = setCryptoPerpAccountLeverageOp{}

func (o setCryptoPerpAccountLeverageOp) Summary() string {
	return "Set account leverage for an asset"
}

func (o setCryptoPerpAccountLeverageOp) ResponseFields() []ResponseField {
	return ResponseSchemas["SetCryptoPerpAccountLeverage"]
}

func (o setCryptoPerpAccountLeverageOp) RequiredFlags() []string {
	return nil
}

func (o setCryptoPerpAccountLeverageOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "leverage", OASName: "leverage", Type: "int", Description: "leverage for the underlying asset", OpName: "SetCryptoPerpAccountLeverage"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol of underlying asset", OpName: "SetCryptoPerpAccountLeverage"},
	}
}

type updateWatchlistByIDOp struct{}

var UpdateWatchlistByIDOp = updateWatchlistByIDOp{}

func (o updateWatchlistByIDOp) Summary() string {
	return "Update watchlist by id"
}

func (o updateWatchlistByIDOp) ResponseFields() []ResponseField {
	return ResponseSchemas["UpdateWatchlistByID"]
}

func (o updateWatchlistByIDOp) RequiredFlags() []string {
	return nil
}

func (o updateWatchlistByIDOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "name", OASName: "name", Type: "string", Description: "name", OpName: "UpdateWatchlistByID"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "symbols", OpName: "UpdateWatchlistByID"},
	}
}

type updateWatchlistByNameOp struct{}

var UpdateWatchlistByNameOp = updateWatchlistByNameOp{}

func (o updateWatchlistByNameOp) Summary() string {
	return "Update watchlist by name"
}

func (o updateWatchlistByNameOp) ResponseFields() []ResponseField {
	return ResponseSchemas["UpdateWatchlistByName"]
}

func (o updateWatchlistByNameOp) RequiredFlags() []string {
	return []string{"name"}
}

func (o updateWatchlistByNameOp) Flags() []FlagDef {
	return []FlagDef{
		{Name: "name", OASName: "name", Type: "string", Description: "name of the watchlist", OpName: "UpdateWatchlistByName", Required: true},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "symbols", OpName: "UpdateWatchlistByName"},
	}
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
		{Name: "asset", Type: "string", Description: "symbol of crypto asset for given transfer (e.g. BTC )"},
		{Name: "chain", Type: "string", Description: "underlying network for given transfer"},
		{Name: "created_at", Type: "string", Description: "timedate when transfer was created"},
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
		{Name: "asset", Type: "string", Description: "symbol of crypto asset for given transfer (e.g. BTC )"},
		{Name: "chain", Type: "string", Description: "underlying network for given transfer"},
		{Name: "created_at", Type: "string", Description: "timedate when transfer was created"},
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
		{Name: "status", Type: "string", Description: "http status code for the attempt to close this position"},
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
		{Name: "asset", Type: "string", Description: "symbol of crypto asset for given transfer (e.g. BTC )"},
		{Name: "chain", Type: "string", Description: "underlying network for given transfer"},
		{Name: "created_at", Type: "string", Description: "timedate when transfer was created"},
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
		{Name: "asset", Type: "string", Description: "symbol of crypto asset for given transfer (e.g. BTC )"},
		{Name: "chain", Type: "string", Description: "underlying network for given transfer"},
		{Name: "created_at", Type: "string", Description: "timedate when transfer was created"},
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
		{Name: "asset", Type: "string", Description: "symbol of crypto asset for given transfer (e.g. BTC )"},
		{Name: "chain", Type: "string", Description: "underlying network for given transfer"},
		{Name: "created_at", Type: "string", Description: "timedate when transfer was created"},
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
		{Name: "asset", Type: "string", Description: "symbol of crypto asset for given transfer (e.g. BTC )"},
		{Name: "chain", Type: "string", Description: "underlying network for given transfer"},
		{Name: "created_at", Type: "string", Description: "timedate when transfer was created"},
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
	"ListWhitelistedAddress":             "List an array of whitelisted addresses",
	"ListWhitelistedPerpAddress":         "List an array of whitelisted addresses",
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
	"DeleteAllOpenPositions":         true,
	"DeleteAllOrders":                true,
	"GetV2Assets":                    true,
	"GetAllOpenPositions":            true,
	"GetAllOrders":                   true,
	"GetWatchlists":                  true,
	"ListCryptoFundingTransfers":     true,
	"ListCryptoFundingWallets":       true,
	"ListCryptoPerpFundingTransfers": true,
	"ListCryptoPerpFundingWallets":   true,
	"ListWhitelistedAddress":         true,
	"ListWhitelistedPerpAddress":     true,
}
