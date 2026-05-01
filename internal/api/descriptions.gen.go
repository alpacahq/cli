// Code generated from api/specs; DO NOT EDIT.

package api

import "sync"

var CalendarOp = Op{
	Name: "Calendar", Summary: "Get market calendar",
	Long:    "This endpoint returns the market calendar",
	Example: `  alpaca calendar market --market XNYS --start 2025-01-01`,
	Flags: []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "last date to retrieve data for (inclusive). Default: one week from the start date", Source: "query"},
		{Name: "market", OASName: "market", Type: "string", Description: "market identifier (MIC, BIC, or acronym)", Completions: []string{"BMO", "BNYM", "BOATS", "CEUX", "CHIX", "HKEX", "IEX", "IEXG", "ISE", "LSE", "MTA", "MTAA", "NASDAQ", "NYSE", "OCEA", "OPRA", "OTC", "OTCM", "SIFMA", "TADAWUL", "XAMS", "XBRU", "XDUB", "XETR", "XETRA", "XHKG", "XLIS", "XLON", "XNAS", "XNYS", "XPAR", "XSAU"}, Required: true, Source: "path"},
		{Name: "start", OASName: "start", Type: "string", Description: "first date to retrieve data for (inclusive). Default: today", Source: "query"},
		{Name: "timezone", OASName: "timezone", Type: "string", Description: "timezone of the times. Default: the timezone of the market", Completions: []string{"UTC"}, Source: "query"},
	},
}

var ClockOp = Op{
	Name: "Clock", Summary: "Get market clock",
	Long:    "This API serves information about multiple markets: the current time, if it's a market day, the current phase of the market, etc",
	Example: `  alpaca clock markets --markets XNYS,XNAS`,
	Flags: []FlagDef{
		{Name: "markets", OASName: "markets", Type: "string", Description: "comma-separated list of markets", Source: "query"},
		{Name: "time", OASName: "time", Type: "string", Description: "instead of the current time, use this time for the clock", Source: "query"},
	},
}

var CorporateActionsOp = Op{
	Name: "CorporateActions", Summary: "Get corporate actions",
	Long:    "This endpoint provides data about the corporate actions for each given symbol over a specified time period",
	Example: `  alpaca data corporate-actions --symbols AAPL --types forward_split --start 2025-01-01`,
	Flags: []FlagDef{
		{Name: "cusips", OASName: "cusips", Type: "string", Description: "A comma-separated list of CUSIPs", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", Source: "query"},
		{Name: "ids", OASName: "ids", Type: "string", Description: "A comma-separated list of corporate action IDs", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "100", Description: "maximum number of corporate actions to return in a response.", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of symbols", Source: "query"},
		{Name: "types", OASName: "types", Type: "string", Description: "A comma-separated list of types", Source: "query"},
	},
}

var CryptoBarsOp = Op{
	Name: "CryptoBars", Summary: "Get historical bars",
	Long:    "The crypto bars API provides historical aggregates for a list of crypto symbols between the specified dates",
	Example: `  alpaca data crypto bars --symbols BTC/USD --start 2025-01-01 --timeframe 1Day`,
	Flags: []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", Source: "query"},
		{Name: "loc", OASName: "loc", Type: "string", Default: "us", Description: "crypto location from where the historical market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU", Completions: []string{"bs-1", "eu-1", "us", "us-1", "us-2"}, Required: true, Source: "path"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", Required: true, Source: "query"},
		{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g", Required: true, Source: "query"},
	},
}

var CryptoLatestBarsOp = Op{
	Name: "CryptoLatestBars", Summary: "Get latest bars",
	Long:    "The latest multi-bars endpoint returns the latest minute-aggregated historical bar data for each of the crypto symbols provided",
	Example: `  alpaca data crypto latest-bars --symbols BTC/USD`,
	Flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Default: "us", Description: "crypto location from where the latest market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU", Completions: []string{"bs-1", "eu-1", "us", "us-1", "us-2"}, Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", Required: true, Source: "query"},
	},
}

var CryptoLatestOrderbooksOp = Op{
	Name: "CryptoLatestOrderbooks", Summary: "Get latest orderbook",
	Long:    "The latest orderbook endpoint returns the latest bid and ask orderbook for the crypto symbols provided",
	Example: `  alpaca data crypto-orderbook --symbols BTC/USD,ETH/USD`,
	Flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Default: "us", Description: "crypto location from where the latest market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU", Completions: []string{"bs-1", "eu-1", "us", "us-1", "us-2"}, Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", Required: true, Source: "query"},
	},
}

var CryptoLatestQuotesOp = Op{
	Name: "CryptoLatestQuotes", Summary: "Get latest quotes",
	Long:    "The latest quotes endpoint returns the latest bid and ask prices for the crypto symbols provided",
	Example: `  alpaca data crypto latest-quotes --symbols BTC/USD`,
	Flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Default: "us", Description: "crypto location from where the latest market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU", Completions: []string{"bs-1", "eu-1", "us", "us-1", "us-2"}, Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", Required: true, Source: "query"},
	},
}

var CryptoLatestTradesOp = Op{
	Name: "CryptoLatestTrades", Summary: "Get latest trades",
	Long:    "The latest trades endpoint returns the latest trade data for the crypto symbols provided",
	Example: `  alpaca data crypto latest-trades --symbols BTC/USD`,
	Flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Default: "us", Description: "crypto location from where the latest market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU", Completions: []string{"bs-1", "eu-1", "us", "us-1", "us-2"}, Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", Required: true, Source: "query"},
	},
}

var CryptoPerpLatestBarsOp = Op{
	Name: "CryptoPerpLatestBars", Summary: "Get latest bars",
	Long:    "The latest bars endpoint returns the latest bar data for the crypto perpetual futures symbols provided",
	Example: `  alpaca crypto-perp data latest-bars --symbols BTC/USD`,
	Flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Default: "us", Description: "crypto perpetual location", Completions: []string{"global"}, Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", Required: true, Source: "query"},
	},
}

var CryptoPerpLatestFuturesPricingOp = Op{
	Name: "CryptoPerpLatestFuturesPricing", Summary: "Get latest pricing",
	Long:    "The latest futures pricing endpoint returns the latest pricing data for the crypto perpetual futures symbols provided",
	Example: `  alpaca crypto-perp data latest-futures-pricing --symbols BTC/USD`,
	Flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Default: "us", Description: "crypto perpetual location", Completions: []string{"global"}, Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", Required: true, Source: "query"},
	},
}

var CryptoPerpLatestOrderbooksOp = Op{
	Name: "CryptoPerpLatestOrderbooks", Summary: "Get latest orderbook",
	Long:    "The latest orderbook endpoint returns the latest bid and ask orderbook for the crypto perpetual futures symbols provided",
	Example: `  alpaca crypto-perp data latest-orderbooks --symbols BTC/USD`,
	Flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Default: "us", Description: "crypto perpetual location", Completions: []string{"global"}, Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", Required: true, Source: "query"},
	},
}

var CryptoPerpLatestQuotesOp = Op{
	Name: "CryptoPerpLatestQuotes", Summary: "Get latest quotes",
	Long:    "The latest quotes endpoint returns the latest bid and ask prices for the crypto perpetual futures symbols provided",
	Example: `  alpaca crypto-perp data latest-quotes --symbols BTC/USD`,
	Flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Default: "us", Description: "crypto perpetual location", Completions: []string{"global"}, Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", Required: true, Source: "query"},
	},
}

var CryptoPerpLatestTradesOp = Op{
	Name: "CryptoPerpLatestTrades", Summary: "Get latest trades",
	Long:    "The latest trades endpoint returns the latest trade data for the crypto perpetual futures symbols provided",
	Example: `  alpaca crypto-perp data latest-trades --symbols BTC/USD`,
	Flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Default: "us", Description: "crypto perpetual location", Completions: []string{"global"}, Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", Required: true, Source: "query"},
	},
}

var CryptoQuotesOp = Op{
	Name: "CryptoQuotes", Summary: "Get historical quotes",
	Long:    "The crypto quotes API provides historical quote data for a list of crypto symbols between the specified dates. The oldest date to retrieve historical quotes of us-1 location is 14th October, 2025 12AM UTC",
	Example: `  alpaca data crypto quotes --symbols BTC/USD --start 2025-01-01`,
	Flags: []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", Source: "query"},
		{Name: "loc", OASName: "loc", Type: "string", Default: "us", Description: "crypto location from where the historical market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU", Completions: []string{"bs-1", "eu-1", "us", "us-1", "us-2"}, Required: true, Source: "path"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", Required: true, Source: "query"},
	},
}

var CryptoSnapshotsOp = Op{
	Name: "CryptoSnapshots", Summary: "Get snapshots",
	Long:    "The snapshots endpoint returns the latest trade, latest quote, latest minute bar, latest daily bar, and previous daily bar data for crypto symbols",
	Example: `  alpaca data crypto snapshots --symbols BTC/USD`,
	Flags: []FlagDef{
		{Name: "loc", OASName: "loc", Type: "string", Default: "us", Description: "crypto location from where the latest market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU", Completions: []string{"bs-1", "eu-1", "us", "us-1", "us-2"}, Required: true, Source: "path"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", Required: true, Source: "query"},
	},
}

var CryptoTradesOp = Op{
	Name: "CryptoTrades", Summary: "Get historical trades",
	Long:    "The crypto trades API provides historical trade data for a list of crypto symbols between the specified dates",
	Example: `  alpaca data crypto trades --symbols BTC/USD --start 2025-01-01`,
	Flags: []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", Source: "query"},
		{Name: "loc", OASName: "loc", Type: "string", Default: "us", Description: "crypto location from where the historical market data is retrieved.\n- us: Alpaca US\n- us-1: Kraken US\n- eu-1: Kraken EU", Completions: []string{"bs-1", "eu-1", "us", "us-1", "us-2"}, Required: true, Source: "path"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of crypto symbols", Required: true, Source: "query"},
	},
}

var FixedIncomeLatestPricesOp = Op{
	Name: "FixedIncomeLatestPrices", Summary: "Get latest prices",
	Long:    "This endpoint returns the latest prices for the given fixed income securities",
	Example: `  alpaca data fixed-income --isins 912797KR1,912797LB5`,
	Flags: []FlagDef{
		{Name: "isins", OASName: "isins", Type: "string", Description: "A comma-separated list of ISINs with a limit of 1000", Required: true, Source: "query"},
	},
}

var LatestRatesOp = Op{
	Name: "LatestRates", Summary: "Get latest rates for currency pairs",
	Long:    "Get the latest forex rates for the given currency pairs",
	Example: `  alpaca data forex latest --currency-pairs EUR/USD,GBP/USD`,
	Flags: []FlagDef{
		{Name: "currency-pairs", OASName: "currency_pairs", Type: "string", Description: "A comma-separated string with currency pairs", Required: true, Source: "query"},
	},
}

var LegacyCalendarOp = Op{
	Name: "LegacyCalendar", Summary: "Get US market calendar",
	Long: "The calendar API serves the full list of market days from 1970 to 2029. It can also be queried by specifying a start and/or end time to narrow down the results",
	Example: `  alpaca calendar
  alpaca calendar --start 2025-01-01 --end 2025-12-31`,
	Flags: []FlagDef{
		{Name: "date-type", OASName: "date_type", Type: "string", Description: "indicates what start and end mean", Completions: []string{"SETTLEMENT", "TRADING"}, Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "last date to retrieve data for (inclusive)", Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "first date to retrieve data for (inclusive)", Source: "query"},
	},
}

var LegacyClockOp = Op{
	Name: "LegacyClock", Summary: "Get US market clock",
	Long:    "The clock API serves the current market timestamp, whether or not the market is currently open, as well as the times of the next market open and close",
	Example: `  alpaca clock`,
}

var LogosOp = Op{
	Name: "Logos", Summary: "Get logos",
	Long:    "Get the image of the company logo for the given symbol",
	Example: `  alpaca data logo --symbol AAPL`,
	Flags: []FlagDef{
		{Name: "placeholder", OASName: "placeholder", Type: "bool", Default: "true", Description: "if true, returns a placeholder image when no logo is available. Defaults to true", Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "A unique series of letters assigned to a security for trading purposes", Required: true, Source: "path"},
	},
}

var MostActivesOp = Op{
	Name: "MostActives", Summary: "Get most active stocks",
	Long: "Returns the most active stocks by volume or trade count based on real time SIP data. By default, returns the top 10 symbols by volume",
	Example: `  alpaca data screener most-actives
  alpaca data screener most-actives --by trades --top 10`,
	Flags: []FlagDef{
		{Name: "by", OASName: "by", Type: "string", Default: "volume", Description: "metric used for ranking the most active stocks", Completions: []string{"trades", "volume"}, Source: "query"},
		{Name: "top", OASName: "top", Type: "int", Default: "10", Description: "number of top most active stocks to fetch per day", Source: "query"},
	},
}

var MoversOp = Op{
	Name: "Movers", Summary: "Get top market movers",
	Long: "Returns the top market movers (gainers and losers) based on real time SIP data. The change for each symbol is calculated from the previous closing price and the latest closing price",
	Example: `  alpaca data screener movers
  alpaca data screener movers --market-type crypto --top 5`,
	Flags: []FlagDef{
		{Name: "market-type", OASName: "market_type", Type: "string", Default: "stocks", Description: "screen-specific market (stocks or crypto)", Completions: []string{"crypto", "stocks"}, Required: true, Source: "path"},
		{Name: "top", OASName: "top", Type: "int", Default: "10", Description: "number of top market movers to fetch (gainers and losers)", Source: "query"},
	},
}

var NewsOp = Op{
	Name: "News", Summary: "Get news articles",
	Long: "Returns the latest news articles across stocks and crypto. By default, returns the latest 10 news articles",
	Example: `  alpaca data news
  alpaca data news --symbols AAPL,MSFT --limit 10`,
	Flags: []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", Source: "query"},
		{Name: "exclude-contentless", OASName: "exclude_contentless", Type: "bool", Description: "boolean indicator to exclude news articles that do not contain content", Source: "query"},
		{Name: "include-content", OASName: "include_content", Type: "bool", Description: "boolean indicator to include content for news articles (if available)", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Description: "limit of news items to be returned for a result page", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "desc", Description: "sort articles by updated date", Completions: []string{"asc", "desc"}, Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of symbols for which to query news", Source: "query"},
	},
}

var OptionChainOp = Op{
	Name: "OptionChain", Summary: "Get option chain",
	Long: "The option chain endpoint provides the latest trade, latest quote, and greeks for each contract symbol of the underlying symbol",
	Example: `  alpaca data option chain --underlying-symbol AAPL
  alpaca data option chain --underlying-symbol SPY --expiration-date 2025-06-20 --type call`,
	Flags: []FlagDef{
		{Name: "expiration-date", OASName: "expiration_date", Type: "string", Description: "filter contracts by the exact expiration date (format: YYYY-MM-DD)", Source: "query"},
		{Name: "expiration-date-gte", OASName: "expiration_date_gte", Type: "string", Description: "filter contracts with expiration date greater than or equal to the specified date", Source: "query"},
		{Name: "expiration-date-lte", OASName: "expiration_date_lte", Type: "string", Description: "filter contracts with expiration date less than or equal to the specified date", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "opra", Description: "source feed of the data", Completions: []string{"indicative", "opra"}, Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "100", Description: "number of maximum snapshots to return in a response.", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", Source: "query"},
		{Name: "root-symbol", OASName: "root_symbol", Type: "string", Description: "filter contracts by the root symbol", Source: "query"},
		{Name: "strike-price-gte", OASName: "strike_price_gte", Type: "string", Description: "filter contracts with strike price greater than or equal to the specified value", Source: "query"},
		{Name: "strike-price-lte", OASName: "strike_price_lte", Type: "string", Description: "filter contracts with strike price less than or equal to the specified value", Source: "query"},
		{Name: "type", OASName: "type", Type: "string", Description: "filter contracts by the type (call or put)", Completions: []string{"call", "put"}, Source: "query"},
		{Name: "underlying-symbol", OASName: "underlying_symbol", Type: "string", Description: "financial instrument on which an option contract is based or derived", Required: true, Source: "path"},
		{Name: "updated-since", OASName: "updated_since", Type: "string", Description: "filter to snapshots that were updated since this timestamp, meaning that the timestamp of the trade or the quote is g...", Source: "query"},
	},
}

var OptionLatestQuotesOp = Op{
	Name: "OptionLatestQuotes", Summary: "Get latest quotes",
	Long:    "The latest multi-quotes endpoint provides the latest bid and ask prices for each given contract symbol",
	Example: `  alpaca data option latest-quotes --symbols AAPL250620C00200000`,
	Flags: []FlagDef{
		{Name: "feed", OASName: "feed", Type: "string", Default: "opra", Description: "source feed of the data", Completions: []string{"indicative", "opra"}, Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100", Required: true, Source: "query"},
	},
}

var OptionLatestTradesOp = Op{
	Name: "OptionLatestTrades", Summary: "Get latest trades",
	Long:    "The latest multi-trades endpoint provides the latest historical trade data for multiple given contract symbols",
	Example: `  alpaca data option latest-trades --symbols AAPL250620C00200000`,
	Flags: []FlagDef{
		{Name: "feed", OASName: "feed", Type: "string", Default: "opra", Description: "source feed of the data", Completions: []string{"indicative", "opra"}, Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100", Required: true, Source: "query"},
	},
}

var OptionMetaConditionsOp = Op{
	Name: "OptionMetaConditions", Summary: "Get condition codes",
	Long:    "Returns the mapping between the condition codes and names",
	Example: `  alpaca data option conditions --ticktype trade`,
	Flags: []FlagDef{
		{Name: "ticktype", OASName: "ticktype", Type: "string", Description: "type of ticks", Completions: []string{"quote", "trade"}, Required: true, Source: "path"},
	},
}

var OptionMetaExchangesOp = Op{
	Name: "OptionMetaExchanges", Summary: "Get exchange codes",
	Long:    "Returns the mapping between the option exchange codes and the corresponding exchange names",
	Example: `  alpaca data option exchanges`,
}

var OptionSnapshotsOp = Op{
	Name: "OptionSnapshots", Summary: "Get snapshots",
	Long: "The snapshots endpoint provides the latest trade, latest quote and greeks for each given contract symbol",
	Example: `  alpaca data option snapshot --symbols AAPL250620C00200000
  alpaca data option snapshot --symbols AAPL250620C00200000,AAPL250620P00200000`,
	Flags: []FlagDef{
		{Name: "feed", OASName: "feed", Type: "string", Default: "opra", Description: "source feed of the data", Completions: []string{"indicative", "opra"}, Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "100", Description: "number of maximum snapshots to return in a response.", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100", Required: true, Source: "query"},
		{Name: "updated-since", OASName: "updated_since", Type: "string", Description: "filter to snapshots that were updated since this timestamp, meaning that the timestamp of the trade or the quote is g...", Source: "query"},
	},
}

var OptionTradesOp = Op{
	Name: "OptionTrades", Summary: "Get historical trades",
	Long:    "The historical option trades API provides trade data for a list of contract symbols between the specified dates",
	Example: `  alpaca data option trades --symbols AAPL250620C00200000 --start 2025-01-01`,
	Flags: []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100", Required: true, Source: "query"},
	},
}

var RatesOp = Op{
	Name: "Rates", Summary: "Get historical rates for currency pairs",
	Long: "Get historical forex rates for the given currency pairs in the given time interval and at the given timeframe (snapshot frequency)",
	Example: `  alpaca data forex rates --currency-pairs EUR/USD,GBP/USD --start 2025-01-01
  alpaca data forex rates --currency-pairs USD/JPY --timeframe 1Hour`,
	Flags: []FlagDef{
		{Name: "currency-pairs", OASName: "currency_pairs", Type: "string", Description: "A comma-separated string with currency pairs", Required: true, Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", Source: "query"},
		{Name: "timeframe", OASName: "timeframe", Type: "string", Default: "1Min", Description: "sampling interval of the currency rates", Source: "query"},
	},
}

var StockAuctionSingleOp = Op{
	Name: "StockAuctionSingle", Summary: "Get historical auctions (single)",
	Long:    "The historical auctions endpoint provides auction prices for the given stock symbol between the specified dates",
	Example: `  alpaca data auction --symbol AAPL --start 2025-01-01`,
	Flags: []FlagDef{
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", Source: "query"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "only sip is valid for auctions", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol to query", Required: true, Source: "path"},
	},
}

var StockAuctionsOp = Op{
	Name: "StockAuctions", Summary: "Get historical auctions",
	Long: "The historical auctions endpoint provides auction prices for a list of stock symbols between the specified dates",
	Example: `  alpaca data auctions --symbols AAPL --start 2025-01-01
  alpaca data auctions --symbols AAPL,MSFT --limit 10`,
	Flags: []FlagDef{
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", Source: "query"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "only sip is valid for auctions", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", Required: true, Source: "query"},
	},
}

var StockBarSingleOp = Op{
	Name: "StockBarSingle", Summary: "Get historical bars (single symbol)",
	Long: "The historical stock bars API provides aggregates for the stock symbol between the specified dates",
	Example: `  alpaca data bars --symbol AAPL --start 2025-01-01 --timeframe 1Day
  alpaca data bars --symbol AAPL --start 2025-01-01 --end 2025-06-01 --limit 100`,
	Flags: []FlagDef{
		{Name: "adjustment", OASName: "adjustment", Type: "string", Default: "raw", Description: "specifies the adjustments for the bars.\n\n - raw: no adjustments\n - split: adjust price and volume for forward and rev...", Source: "query"},
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", Source: "query"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "source feed of the data.", Completions: []string{"boats", "iex", "otc", "sip"}, Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol to query", Required: true, Source: "path"},
		{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g", Required: true, Source: "query"},
	},
}

var StockBarsOp = Op{
	Name: "StockBars", Summary: "Get historical bars",
	Long:    "The historical stock bars API provides aggregates for a list of stock symbols between the specified dates",
	Example: `  alpaca data multi-bars --symbols AAPL,MSFT --start 2025-01-01 --timeframe 1Day`,
	Flags: []FlagDef{
		{Name: "adjustment", OASName: "adjustment", Type: "string", Default: "raw", Description: "specifies the adjustments for the bars.\n\n - raw: no adjustments\n - split: adjust price and volume for forward and rev...", Source: "query"},
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", Source: "query"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "source feed of the data.", Completions: []string{"boats", "iex", "otc", "sip"}, Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", Required: true, Source: "query"},
		{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g", Required: true, Source: "query"},
	},
}

var StockLatestBarSingleOp = Op{
	Name: "StockLatestBarSingle", Summary: "Get latest bar (single symbol)",
	Long:    "The latest bar endpoint returns the latest minute bar for the given ticker symbol",
	Example: `  alpaca data latest-bar --symbol AAPL`,
	Flags: []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol to query", Required: true, Source: "path"},
	},
}

var StockLatestBarsOp = Op{
	Name: "StockLatestBars", Summary: "Get latest bars",
	Long:    "The latest bars endpoint provides the latest minute bar for the given ticker symbols",
	Example: `  alpaca data latest-bars --symbols AAPL,MSFT`,
	Flags: []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", Required: true, Source: "query"},
	},
}

var StockLatestQuoteSingleOp = Op{
	Name: "StockLatestQuoteSingle", Summary: "Get latest quote (single symbol)",
	Long:    "The latest quote endpoint provides the latest best bid and ask prices for a given ticker symbol",
	Example: `  alpaca data latest-quote --symbol AAPL`,
	Flags: []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol to query", Required: true, Source: "path"},
	},
}

var StockLatestQuotesOp = Op{
	Name: "StockLatestQuotes", Summary: "Get latest quotes",
	Long:    "The latest quotes endpoint provides the latest best bid and ask prices for the given ticker symbols",
	Example: `  alpaca data latest-quotes --symbols AAPL,MSFT`,
	Flags: []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", Required: true, Source: "query"},
	},
}

var StockLatestTradeSingleOp = Op{
	Name: "StockLatestTradeSingle", Summary: "Get latest trade (single symbol)",
	Long: "The latest trade endpoint provides the latest trade for the given ticker symbol",
	Example: `  alpaca data latest-trade --symbol AAPL
  alpaca data latest-trade --symbol AAPL --feed sip`,
	Flags: []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol to query", Required: true, Source: "path"},
	},
}

var StockLatestTradesOp = Op{
	Name: "StockLatestTrades", Summary: "Get latest trades",
	Long:    "The latest trades endpoint provides the latest trades for the given ticker symbols",
	Example: `  alpaca data latest-trades --symbols AAPL,MSFT`,
	Flags: []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", Required: true, Source: "query"},
	},
}

var StockMetaConditionsOp = Op{
	Name: "StockMetaConditions", Summary: "Get condition codes",
	Long:    "Returns the mapping between the condition codes and names",
	Example: `  alpaca data meta conditions --ticktype trade`,
	Flags: []FlagDef{
		{Name: "tape", OASName: "tape", Type: "string", Description: "one character name of the tape", Completions: []string{"A", "B", "C"}, Required: true, Source: "query"},
		{Name: "ticktype", OASName: "ticktype", Type: "string", Description: "type of ticks", Completions: []string{"quote", "trade"}, Required: true, Source: "path"},
	},
}

var StockMetaExchangesOp = Op{
	Name: "StockMetaExchanges", Summary: "Get exchange codes",
	Long:    "Returns the mapping between the stock exchange codes and the corresponding exchange names",
	Example: `  alpaca data meta exchanges`,
}

var StockQuoteSingleOp = Op{
	Name: "StockQuoteSingle", Summary: "Get historical quotes (single symbol)",
	Long: "The historical stock quotes API provides quote data for a stock symbol between the specified dates",
	Example: `  alpaca data quotes --symbol AAPL --start 2025-01-01
  alpaca data quotes --symbol AAPL --start 2025-01-01 --end 2025-01-31 --limit 50`,
	Flags: []FlagDef{
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", Source: "query"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "source feed of the data.", Completions: []string{"boats", "iex", "otc", "sip"}, Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol to query", Required: true, Source: "path"},
	},
}

var StockQuotesOp = Op{
	Name: "StockQuotes", Summary: "Get historical quotes",
	Long:    "The historical stock quotes API provides quote data for a list of stock symbols between the specified dates",
	Example: `  alpaca data multi-quotes --symbols AAPL,MSFT --start 2025-01-01`,
	Flags: []FlagDef{
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", Source: "query"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "source feed of the data.", Completions: []string{"boats", "iex", "otc", "sip"}, Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", Required: true, Source: "query"},
	},
}

var StockSnapshotSingleOp = Op{
	Name: "StockSnapshotSingle", Summary: "Get snapshot (single symbol)",
	Long: "Returns the latest snapshot for a stock symbol. Use --jq to flatten for CSV.",
	Example: `  alpaca data snapshot --symbol AAPL
  alpaca data snapshot --symbol AAPL --feed sip`,
	Flags: []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol to query", Required: true, Source: "path"},
	},
}

var StockSnapshotsOp = Op{
	Name: "StockSnapshots", Summary: "Get snapshots",
	Long:    "The snapshot endpoint for multiple tickers provides the latest trade, latest quote, minute bar, daily bar, and previous daily bar data for each given ticker symbol",
	Example: `  alpaca data multi-snapshots --symbols AAPL,MSFT`,
	Flags: []FlagDef{
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Description: "source feed of the data.", Completions: []string{"boats", "delayed_sip", "iex", "otc", "overnight", "sip"}, Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", Required: true, Source: "query"},
	},
}

var StockTradeSingleOp = Op{
	Name: "StockTradeSingle", Summary: "Get historical trades (single symbol)",
	Long: "The historical stock trades API provides trade data for a stock symbol between the specified dates",
	Example: `  alpaca data trades --symbol AAPL --start 2025-01-01
  alpaca data trades --symbol AAPL --start 2025-01-01 --limit 100`,
	Flags: []FlagDef{
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", Source: "query"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "source feed of the data.", Completions: []string{"boats", "iex", "otc", "sip"}, Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol to query", Required: true, Source: "path"},
	},
}

var StockTradesOp = Op{
	Name: "StockTrades", Summary: "Get historical trades",
	Long:    "The historical stock trades API provides trade data for a list of stock symbols between the specified dates",
	Example: `  alpaca data multi-trades --symbols AAPL,MSFT --start 2025-01-01`,
	Flags: []FlagDef{
		{Name: "asof", OASName: "asof", Type: "string", Description: "as-of date of the queried stock symbol(s)", Source: "query"},
		{Name: "currency", OASName: "currency", Type: "string", Description: "currency of all prices in ISO 4217 format. Default: USD", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", Source: "query"},
		{Name: "feed", OASName: "feed", Type: "string", Default: "sip", Description: "source feed of the data.", Completions: []string{"boats", "iex", "otc", "sip"}, Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of stock symbols", Required: true, Source: "query"},
	},
}

var UsCorporatesOp = Op{
	Name: "UsCorporates", Summary: "Get US corporates",
	Long: "Serves the list of US corporates available at Alpaca. The response is sorted by ISIN",
	Example: `  alpaca asset bond
  alpaca asset bond --bond-status active`,
	Flags: []FlagDef{
		{Name: "bond-status", OASName: "bond_status", Type: "string", Description: "status of the bond", Completions: []string{"matured", "outstanding", "pre_issuance"}, Source: "query"},
		{Name: "cusips", OASName: "cusips", Type: "string", Description: "A comma-separated list of CUSIPs with a limit of 1000", Source: "query"},
		{Name: "isins", OASName: "isins", Type: "string", Description: "A comma-separated list of ISINs with a limit of 1000", Source: "query"},
		{Name: "tickers", OASName: "tickers", Type: "string", Description: "A comma-separated list of tickers with a limit of 1000", Source: "query"},
	},
}

var UsTreasuriesOp = Op{
	Name: "UsTreasuries", Summary: "Get US treasuries",
	Long: "Serves the list of US treasuries available at Alpaca. The response is sorted by ISIN",
	Example: `  alpaca asset treasury
  alpaca asset treasury --bond-status active`,
	Flags: []FlagDef{
		{Name: "bond-status", OASName: "bond_status", Type: "string", Description: "status of the bond", Completions: []string{"matured", "outstanding", "pre_issuance"}, Source: "query"},
		{Name: "cusips", OASName: "cusips", Type: "string", Description: "A comma-separated list of CUSIPs with a limit of 1000", Source: "query"},
		{Name: "isins", OASName: "isins", Type: "string", Description: "A comma-separated list of ISINs with a limit of 1000", Source: "query"},
		{Name: "subtype", OASName: "subtype", Type: "string", Description: "subtype of the treasury", Completions: []string{"bill", "bond", "floating", "note", "strips", "tips"}, Source: "query"},
	},
}

var AddAssetToWatchlistOp = Op{
	Name: "AddAssetToWatchlist", Summary: "Add asset to watchlist",
	Long:    "Append an asset for the symbol to the end of watchlist asset list",
	Example: `  alpaca watchlist add --watchlist-id <id> --symbol AAPL`,
	Flags: []FlagDef{
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "the symbol name to add to the watchlist", Source: "body"},
		{Name: "watchlist-id", OASName: "watchlist_id", Type: "string", Description: "watchlist id", Required: true, Source: "path"},
	},
}

var AddAssetToWatchlistByNameOp = Op{
	Name: "AddAssetToWatchlistByName", Summary: "Add asset to watchlist by name",
	Long:    "Append an asset for the symbol to the end of watchlist asset list",
	Example: `  alpaca watchlist add-by-name --name "Tech Stocks" --symbol NVDA`,
	Flags: []FlagDef{
		{Name: "name", OASName: "name", Type: "string", Description: "name of the watchlist", Required: true, Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "the symbol name to add to the watchlist", Source: "body"},
	},
}

var CreateCryptoPerpTransferForAccountOp = Op{
	Name: "CreateCryptoPerpTransferForAccount", Summary: "Request a new withdrawal",
	Long:    "Creates a withdrawal request. Note that outgoing withdrawals must be sent to a whitelisted address and you must whitelist addresses at least 24 hours in advance. If you attempt to withdraw funds to a non-whitelisted address then the transfer will be rejected",
	Example: `  alpaca crypto-perp wallet transfer create --amount 0.5 --address 0xabc... --asset BTC`,
	Flags: []FlagDef{
		{Name: "address", OASName: "address", Type: "string", Description: "destination wallet address", Source: "body"},
		{Name: "amount", OASName: "amount", Type: "string", Description: "amount, denoted in the specified asset, to be withdrawn from the user’s wallet", Source: "body"},
		{Name: "asset", OASName: "asset", Type: "string", Description: "crypto asset symbol, e.g. BTC, ETH, USDT", Source: "body"},
	},
}

var CreateCryptoTransferForAccountOp = Op{
	Name: "CreateCryptoTransferForAccount", Summary: "Request a new withdrawal",
	Long:    "Creates a withdrawal request. Note that outgoing withdrawals must be sent to a whitelisted address and you must whitelist addresses at least 24 hours in advance. If you attempt to withdraw funds to a non-whitelisted address then the transfer will be rejected",
	Example: `  alpaca wallet transfer create --amount 0.5 --address 0xabc... --asset BTC`,
	Flags: []FlagDef{
		{Name: "address", OASName: "address", Type: "string", Description: "destination wallet address", Source: "body"},
		{Name: "amount", OASName: "amount", Type: "string", Description: "amount, denoted in the specified asset, to be withdrawn from the user’s wallet", Source: "body"},
		{Name: "asset", OASName: "asset", Type: "string", Description: "crypto asset symbol, e.g. BTC, ETH, USDT", Source: "body"},
	},
}

var CreateWhitelistedAddressOp = Op{
	Name: "CreateWhitelistedAddress", Summary: "Request a new whitelisted address",
	Example: `  alpaca wallet whitelist add --address 0xabc... --asset ETH`,
	Flags: []FlagDef{
		{Name: "address", OASName: "address", Type: "string", Description: "address to be whitelisted", Source: "body"},
		{Name: "asset", OASName: "asset", Type: "string", Description: "symbol of underlying asset for the whitelisted address", Source: "body"},
	},
}

var CreateWhitelistedPerpAddressOp = Op{
	Name: "CreateWhitelistedPerpAddress", Summary: "Request a new whitelisted address",
	Example: `  alpaca crypto-perp wallet whitelist add --address 0xabc... --asset ETH`,
	Flags: []FlagDef{
		{Name: "address", OASName: "address", Type: "string", Description: "address to be whitelisted", Source: "body"},
		{Name: "asset", OASName: "asset", Type: "string", Description: "symbol of underlying asset for the whitelisted address", Source: "body"},
	},
}

var DeleteAllOpenPositionsOp = Op{
	Name: "DeleteAllOpenPositions", Summary: "Close all positions", ReturnsArray: true,
	Long:    "Closes (liquidates) all of the account’s open long and short positions. A response will be provided for each order that is attempted to be cancelled. If an order is no longer cancelable, the server will respond with status 500 and reject the request",
	Example: `  alpaca position close-all`,
	Flags: []FlagDef{
		{Name: "cancel-orders", OASName: "cancel_orders", Type: "bool", Description: "if true is specified, cancel all open orders before liquidating all positions", Source: "query"},
	},
}

var DeleteAllOrdersOp = Op{
	Name: "DeleteAllOrders", Summary: "Delete all orders", ReturnsArray: true,
	Long:    "Attempts to cancel all open orders. A response will be provided for each order that is attempted to be cancelled. If an order is no longer cancelable, the server will respond with status 500 and reject the request",
	Example: `  alpaca order cancel-all`,
}

var DeleteOpenPositionOp = Op{
	Name: "DeleteOpenPosition", Summary: "Close a position",
	Long: "Closes (liquidates) the account’s open position for the given symbol. Works for both long and short positions",
	Example: `  alpaca position close --symbol-or-asset-id AAPL
  alpaca position close --symbol-or-asset-id AAPL --qty 5
  alpaca position close --symbol-or-asset-id AAPL --percentage 50`,
	Flags: []FlagDef{
		{Name: "percentage", OASName: "percentage", Type: "string", Description: "percentage of position to liquidate", Source: "query"},
		{Name: "qty", OASName: "qty", Type: "string", Description: "the number of shares to liquidate. Can accept up to 9 decimal points. Cannot work with percentage", Source: "query"},
		{Name: "symbol-or-asset-id", OASName: "symbol_or_asset_id", Type: "string", Description: "symbol or assetId", Required: true, Source: "path"},
	},
}

var DeleteOrderByOrderIDOp = Op{
	Name: "DeleteOrderByOrderID", Summary: "Delete order by ID",
	Long:    "Attempts to cancel an Open Order. If the order is no longer cancelable, the request will be rejected with status 422; otherwise accepted with return status 204",
	Example: `  alpaca order cancel --order-id <id>`,
	Flags: []FlagDef{
		{Name: "order-id", OASName: "order_id", Type: "string", Description: "order id", Required: true, Source: "path"},
	},
}

var DeleteWatchlistByIDOp = Op{
	Name: "DeleteWatchlistByID", Summary: "Delete watchlist by id",
	Long:    "Delete a watchlist. This is a permanent deletion",
	Example: `  alpaca watchlist delete --watchlist-id <id>`,
	Flags: []FlagDef{
		{Name: "watchlist-id", OASName: "watchlist_id", Type: "string", Description: "watchlist id", Required: true, Source: "path"},
	},
}

var DeleteWatchlistByNameOp = Op{
	Name: "DeleteWatchlistByName", Summary: "Delete watchlist by name",
	Long:    "Delete a watchlist. This is a permanent deletion",
	Example: `  alpaca watchlist delete-by-name --name "Tech Stocks"`,
	Flags: []FlagDef{
		{Name: "name", OASName: "name", Type: "string", Description: "name of the watchlist", Required: true, Source: "query"},
	},
}

var DeleteWhitelistedAddressOp = Op{
	Name: "DeleteWhitelistedAddress", Summary: "Delete a whitelisted address",
	Example: `  alpaca wallet whitelist delete --whitelisted-address-id <id>`,
	Flags: []FlagDef{
		{Name: "whitelisted-address-id", OASName: "whitelisted_address_id", Type: "string", Description: "whitelisted address to delete", Required: true, Source: "path"},
	},
}

var DeleteWhitelistedPerpAddressOp = Op{
	Name: "DeleteWhitelistedPerpAddress", Summary: "Delete a whitelisted address",
	Example: `  alpaca crypto-perp wallet whitelist delete --whitelisted-address-id <id>`,
	Flags: []FlagDef{
		{Name: "whitelisted-address-id", OASName: "whitelisted_address_id", Type: "string", Description: "whitelisted address to delete", Required: true, Source: "path"},
	},
}

var GetOptionContractSymbolOrIDOp = Op{
	Name: "GetOptionContractSymbolOrID", Summary: "Get an option contract by ID or symbol",
	Long:    "Get an option contract by symbol or contract ID. The symbol or id should be passed in as a path parameter",
	Example: `  alpaca option get --symbol-or-id AAPL250620C00200000`,
	Flags: []FlagDef{
		{Name: "symbol-or-id", OASName: "symbol_or_id", Type: "string", Description: "symbol or contract ID", Required: true, Source: "path"},
	},
}

var GetOptionsContractsOp = Op{
	Name: "GetOptionsContracts", Summary: "Get option contracts",
	Long: "List option contracts for an underlying symbol. For market data (greeks, pricing), use `data option chain`.",
	Example: `  alpaca option contracts --underlying-symbols AAPL
  alpaca option contracts --underlying-symbols AAPL --expiration-date 2025-06-20 --type call
  alpaca option contracts --underlying-symbols SPY --strike-price-gte 400 --strike-price-lte 450`,
	Flags: []FlagDef{
		{Name: "expiration-date", OASName: "expiration_date", Type: "string", Description: "filter contracts by the exact expiration date (format: YYYY-MM-DD)", Source: "query"},
		{Name: "expiration-date-gte", OASName: "expiration_date_gte", Type: "string", Description: "filter contracts with expiration date greater than or equal to the specified date", Source: "query"},
		{Name: "expiration-date-lte", OASName: "expiration_date_lte", Type: "string", Description: "filter contracts with expiration date less than or equal to the specified date", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Description: "number of contracts to limit per page (default=100, max=10000)", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "used for pagination, this token retrieves the next page of results", Source: "query"},
		{Name: "ppind", OASName: "ppind", Type: "bool", Description: "ppind(Penny Program Indicator) field indicates whether an option contract is eligible for penny price increments,", Source: "query"},
		{Name: "root-symbol", OASName: "root_symbol", Type: "string", Description: "filter contracts by the root symbol", Source: "query"},
		{Name: "show-deliverables", OASName: "show_deliverables", Type: "bool", Description: "include deliverables array in the response", Source: "query"},
		{Name: "status", OASName: "status", Type: "string", Description: "filter contracts by status (active/inactive). By default only active contracts are returned", Completions: []string{"active", "inactive"}, Source: "query"},
		{Name: "strike-price-gte", OASName: "strike_price_gte", Type: "string", Description: "filter contracts with strike price greater than or equal to the specified value", Source: "query"},
		{Name: "strike-price-lte", OASName: "strike_price_lte", Type: "string", Description: "filter contracts with strike price less than or equal to the specified value", Source: "query"},
		{Name: "style", OASName: "style", Type: "string", Description: "filter contracts by the style (american/european)", Completions: []string{"american", "european"}, Source: "query"},
		{Name: "type", OASName: "type", Type: "string", Description: "filter contracts by the type (call/put)", Completions: []string{"call", "put"}, Source: "query"},
		{Name: "underlying-symbols", OASName: "underlying_symbols", Type: "string", Description: "filter contracts by one or more underlying symbols", Source: "query"},
	},
}

var GetV2AssetsOp = Op{
	Name: "GetV2Assets", Summary: "Get assets", ReturnsArray: true,
	Long: "The assets API serves as the master list of assets available for trade and data consumption from Alpaca. Assets are sorted by asset class, exchange and symbol",
	Example: `  alpaca asset list
  alpaca asset list --asset-class us_equity --status active
  alpaca asset list --exchange NYSE`,
	Flags: []FlagDef{
		{Name: "asset-class", OASName: "asset_class", Type: "string", Description: "defaults to us_equity", Source: "query"},
		{Name: "attributes", OASName: "attributes", Type: "string", Description: "comma separated values to query for more than one attribute", Source: "query"},
		{Name: "exchange", OASName: "exchange", Type: "string", Description: "optional AMEX, ARCA, BATS, NYSE, NASDAQ, NYSEARCA or OTC", Source: "query"},
		{Name: "status", OASName: "status", Type: "string", Description: "e.g. “active”. By default, all statuses are included", Source: "query"},
	},
}

var GetV2AssetsSymbolOrAssetIDOp = Op{
	Name: "GetV2AssetsSymbolOrAssetID", Summary: "Get an asset by ID or symbol",
	Long: "Get the asset model for a given symbol or asset_id. The symbol or asset_id should be passed in as a path parameter",
	Example: `  alpaca asset get --symbol-or-asset-id AAPL
  alpaca asset get --symbol-or-asset-id BTC/USD`,
	Flags: []FlagDef{
		{Name: "symbol-or-asset-id", OASName: "symbol_or_asset_id", Type: "string", Description: "symbol or assetId. CUSIP is also accepted for US equities", Required: true, Source: "path"},
	},
}

var GetV2CorporateActionsAnnouncementsOp = Op{
	Name: "GetV2CorporateActionsAnnouncements", Summary: "Retrieve announcements",
	Long: "This endpoint is deprecated, please use the new corporate actions endpoint instead",
	Example: `  alpaca corporate-action list --ca-types reverse_split --since 2025-01-01 --until 2025-12-31
  alpaca corporate-action list --ca-types cash_dividend --symbol AAPL --since 2025-01-01 --until 2025-06-30`,
	Flags: []FlagDef{
		{Name: "ca-types", OASName: "ca_types", Type: "string", Description: "A comma-delimited list of Dividend, Merger, Spinoff, or Split", Required: true, Source: "query"},
		{Name: "cusip", OASName: "cusip", Type: "string", Description: "CUSIP of the company initiating the announcement", Source: "query"},
		{Name: "date-type", OASName: "date_type", Type: "string", Description: "declaration_date, ex_date, record_date, or payable_date", Source: "query"},
		{Name: "since", OASName: "since", Type: "string", Description: "start (inclusive) of the date range when searching corporate action announcements", Required: true, Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol of the company initiating the announcement", Source: "query"},
		{Name: "until", OASName: "until", Type: "string", Description: "end (inclusive) of the date range when searching corporate action announcements", Required: true, Source: "query"},
	},
}

var GetV2CorporateActionsAnnouncementsIDOp = Op{
	Name: "GetV2CorporateActionsAnnouncementsID", Summary: "Retrieve a specific announcement",
	Long:    "This endpoint is deprecated, please use the new corporate actions endpoint instead",
	Example: `  alpaca corporate-action get --id <announcement-id>`,
	Flags: []FlagDef{
		{Name: "id", OASName: "id", Type: "string", Description: "corporate announcement’s id", Required: true, Source: "path"},
	},
}

var GetAccountOp = Op{
	Name: "GetAccount", Summary: "Get account",
	Long:    "Returns your account details",
	Example: `  alpaca account get`,
}

var GetAccountActivitiesOp = Op{
	Name: "GetAccountActivities", Summary: "Retrieve account activities",
	Long: "Returns a list of activities",
	Example: `  alpaca account activity list
  alpaca account activity list --activity-types FILL,TRANS --page-size 20
  alpaca account activity list --direction desc`,
	Flags: []FlagDef{
		{Name: "activity-types", OASName: "activity_types", Type: "string", Description: "A comma-separated list of activity types used to filter the results", Source: "query"},
		{Name: "after", OASName: "after", Type: "string", Description: "get activities created after this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported", Source: "query"},
		{Name: "category", OASName: "category", Type: "string", Description: "activity category. Cannot be used with \"activity_types\" parameter", Completions: []string{"non_trade_activity", "trade_activity"}, Source: "query"},
		{Name: "date", OASName: "date", Type: "string", Description: "filter activities by their creation date (created_at), not the activity's settlement date", Source: "query"},
		{Name: "direction", OASName: "direction", Type: "string", Default: "desc", Description: "chronological order of response based on the activity datetime", Completions: []string{"asc", "desc"}, Source: "query"},
		{Name: "page-size", OASName: "page_size", Type: "int", Default: "100", Description: "maximum number of entries to return in the response", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "token used for pagination. Provide the ID of the last activity from the last page to retrieve the next set of results", Source: "query"},
		{Name: "until", OASName: "until", Type: "string", Description: "get activities created before this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported", Source: "query"},
	},
}

var GetAccountActivitiesByActivityTypeOp = Op{
	Name: "GetAccountActivitiesByActivityType", Summary: "Retrieve account activities of specific type",
	Long: "Returns account activity entries for a specific type of activity",
	Example: `  alpaca account activity list-by-type --activity-type FILL --page-size 20
  alpaca account activity list-by-type --activity-type DIV --after 2025-01-01`,
	Flags: []FlagDef{
		{Name: "activity-type", OASName: "activity_type", Type: "string", Description: "activity type you want to view entries for. A list of valid activity types can be found at the bottom of this page", Required: true, Source: "path"},
		{Name: "after", OASName: "after", Type: "string", Description: "get activities created after this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported", Source: "query"},
		{Name: "date", OASName: "date", Type: "string", Description: "filter activities by their creation date (created_at), not the activity's settlement date", Source: "query"},
		{Name: "direction", OASName: "direction", Type: "string", Default: "desc", Description: "chronological order of response based on the activity datetime", Completions: []string{"asc", "desc"}, Source: "query"},
		{Name: "page-size", OASName: "page_size", Type: "int", Default: "100", Description: "maximum number of entries to return in the response", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "token used for pagination. Provide the ID of the last activity from the last page to retrieve the next set of results", Source: "query"},
		{Name: "until", OASName: "until", Type: "string", Description: "get activities created before this date. Both formats YYYY-MM-DD and YYYY-MM-DDTHH:MM:SSZ are supported", Source: "query"},
	},
}

var GetAccountConfigOp = Op{
	Name: "GetAccountConfig", Summary: "Get account configurations",
	Long:    "gets the current account configuration values",
	Example: `  alpaca account config get`,
}

var GetAccountPortfolioHistoryOp = Op{
	Name: "GetAccountPortfolioHistory", Summary: "Get account portfolio history",
	Long: "Returns portfolio equity and P&L history. Use --jq to flatten for CSV.",
	Example: `  alpaca account portfolio
  alpaca account portfolio --period 1M --timeframe 1D`,
	Flags: []FlagDef{
		{Name: "cashflow-types", OASName: "cashflow_types", Type: "string", Description: "cashflow activities to include in the report. One of 'ALL', 'NONE', or a comma-separated list of activity types", Source: "query"},
		{Name: "end", OASName: "end", Type: "string", Description: "timestamp the data is returned up to in RFC3339 format (including timezone specification)", Source: "query"},
		{Name: "extended-hours", OASName: "extended_hours", Type: "string", Description: "**deprecated**: Users are strongly advised to **rely on the intraday_reporting query parameter** for better control\no...", Source: "query"},
		{Name: "intraday-reporting", OASName: "intraday_reporting", Type: "string", Default: "market_hours", Description: "for intraday resolutions (<1D) this specifies which timestamps to return data points for:\n\nAllowed values are:\n- **ma...", Completions: []string{"continuous", "extended_hours", "market_hours"}, Source: "query"},
		{Name: "period", OASName: "period", Type: "string", Description: "duration of the data in number + unit format, such as 1D, where unit can be D for day, W for week, M for month and A ...", Source: "query"},
		{Name: "pnl-reset", OASName: "pnl_reset", Type: "string", Default: "per_day", Description: "pnl_reset defines how we are calculating the baseline values for Profit And Loss (pnl) for queries with timeframe les...", Completions: []string{"no_reset", "per_day"}, Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "timestamp the data is returned starting from in RFC3339 format (including timezone specification)", Source: "query"},
		{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "resolution of time window", Source: "query"},
	},
}

var GetAllOpenPositionsOp = Op{
	Name: "GetAllOpenPositions", Summary: "List all open positions", ReturnsArray: true,
	Long: "The positions API provides information about an account’s current open positions. The response will include information such as cost basis, shares traded, and market value, which will be updated live as price information is updated",
	Example: `  alpaca position list
  alpaca position list --csv`,
}

var GetAllOrdersOp = Op{
	Name: "GetAllOrders", Summary: "Get all orders", ReturnsArray: true,
	Long: "Retrieves a list of orders for the account, filtered by the supplied query parameters",
	Example: `  alpaca order list
  alpaca order list --status closed --limit 20
  alpaca order list --symbols AAPL,MSFT --after 2025-01-01`,
	Flags: []FlagDef{
		{Name: "after", OASName: "after", Type: "string", Description: "response will include only ones submitted after this timestamp (exclusive.)", Source: "query"},
		{Name: "after-order-id", OASName: "after_order_id", Type: "string", Description: "return orders submitted after the order with this ID (exclusive).\nMutually exclusive with before_order_id", Source: "query"},
		{Name: "asset-class", OASName: "asset_class", Type: "string", Description: "A comma-separated list of asset classes, the response will include only orders in the specified asset classes", Source: "query"},
		{Name: "before-order-id", OASName: "before_order_id", Type: "string", Description: "return orders submitted before the order with this ID (exclusive).\nMutually exclusive with after_order_id", Source: "query"},
		{Name: "direction", OASName: "direction", Type: "string", Description: "chronological order of response based on the submission time. asc or desc. Defaults to desc", Completions: []string{"asc", "desc"}, Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Description: "maximum number of orders in response. Defaults to 50 and max is 500", Source: "query"},
		{Name: "nested", OASName: "nested", Type: "bool", Description: "if true, the result will roll up multi-leg orders under the legs field of primary order", Source: "query"},
		{Name: "side", OASName: "side", Type: "string", Description: "filters down to orders that have a matching side field set", Source: "query"},
		{Name: "status", OASName: "status", Type: "string", Default: "open", Description: "order status to be queried. open, closed or all. Defaults to open", Completions: []string{"all", "closed", "open"}, Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of symbols to filter by (ex", Source: "query"},
		{Name: "until", OASName: "until", Type: "string", Description: "response will include only ones submitted until this timestamp (exclusive.)", Source: "query"},
	},
}

var GetCryptoFundingTransferOp = Op{
	Name: "GetCryptoFundingTransfer", Summary: "Retrieve a crypto funding transfer",
	Long:    "Returns a specific wallet transfer by passing into the query the transfer_id",
	Example: `  alpaca wallet transfer get --transfer-id <id>`,
	Flags: []FlagDef{
		{Name: "transfer-id", OASName: "transfer_id", Type: "string", Description: "crypto transfer to retrieve", Required: true, Source: "path"},
	},
}

var GetCryptoPerpAccountLeverageOp = Op{
	Name: "GetCryptoPerpAccountLeverage", Summary: "Get account leverage for an asset",
	Long:    "Retrieves the current leverage setting for the crypto perpetuals account, specific to a given underlying asset. To use this endpoint, provide the 'symbol' of the asset as a query parameter",
	Example: `  alpaca crypto-perp leverage`,
	Flags: []FlagDef{
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol of underlying asset", Source: "query"},
	},
}

var GetCryptoPerpAccountVitalsOp = Op{
	Name: "GetCryptoPerpAccountVitals", Summary: "Retrieve account vitals",
	Long:    "Fetches key financial metrics for the crypto perpetuals account, providing a snapshot of its current status by detailing the relationship between the user's positions and their collateral",
	Example: `  alpaca crypto-perp vitals`,
}

var GetCryptoPerpFundingTransferOp = Op{
	Name: "GetCryptoPerpFundingTransfer", Summary: "Retrieve a crypto funding transfer",
	Long:    "Returns a specific wallet transfer by passing into the query the transfer_id",
	Example: `  alpaca crypto-perp wallet transfer get --transfer-id <id>`,
	Flags: []FlagDef{
		{Name: "transfer-id", OASName: "transfer_id", Type: "string", Description: "crypto transfer to retrieve", Required: true, Source: "path"},
	},
}

var GetCryptoPerpTransferEstimateOp = Op{
	Name: "GetCryptoPerpTransferEstimate", Summary: "Returns the estimated gas fee for a proposed transaction",
	Example: `  alpaca crypto-perp wallet transfer estimate --asset BTC --amount 0.5`,
	Flags: []FlagDef{
		{Name: "amount", OASName: "amount", Type: "string", Description: "amount, denoted in the specified asset, of the proposed transaction", Source: "query"},
		{Name: "asset", OASName: "asset", Type: "string", Description: "asset for the proposed transaction", Source: "query"},
		{Name: "from-address", OASName: "from_address", Type: "string", Description: "originating address of the proposed transaction", Source: "query"},
		{Name: "to-address", OASName: "to_address", Type: "string", Description: "destination address of the proposed transaction", Source: "query"},
	},
}

var GetCryptoTransferEstimateOp = Op{
	Name: "GetCryptoTransferEstimate", Summary: "Returns the estimated gas fee for a proposed transaction",
	Example: `  alpaca wallet transfer estimate --asset BTC --amount 0.5 \
    --from-address 0xabc... --to-address 0xdef...`,
	Flags: []FlagDef{
		{Name: "amount", OASName: "amount", Type: "string", Description: "amount, denoted in the specified asset, of the proposed transaction", Source: "query"},
		{Name: "asset", OASName: "asset", Type: "string", Description: "asset for the proposed transaction", Source: "query"},
		{Name: "from-address", OASName: "from_address", Type: "string", Description: "originating address of the proposed transaction", Source: "query"},
		{Name: "to-address", OASName: "to_address", Type: "string", Description: "destination address of the proposed transaction", Source: "query"},
	},
}

var GetOpenPositionOp = Op{
	Name: "GetOpenPosition", Summary: "Get an open position",
	Long: "Retrieves the account’s open position for the given symbol or assetId",
	Example: `  alpaca position get --symbol-or-asset-id AAPL
  alpaca position get --symbol-or-asset-id BTC/USD`,
	Flags: []FlagDef{
		{Name: "symbol-or-asset-id", OASName: "symbol_or_asset_id", Type: "string", Description: "symbol or assetId", Required: true, Source: "path"},
	},
}

var GetOrderByClientOrderIDOp = Op{
	Name: "GetOrderByClientOrderID", Summary: "Get order by client order ID",
	Long:    "Retrieves a single order specified by the client order ID",
	Example: `  alpaca order get-by-client-id --client-order-id my-order-123`,
	Flags: []FlagDef{
		{Name: "client-order-id", OASName: "client_order_id", Type: "string", Description: "client-assigned order ID", Required: true, Source: "query"},
	},
}

var GetOrderByOrderIDOp = Op{
	Name: "GetOrderByOrderID", Summary: "Get order by ID",
	Long:    "Retrieves a single order for the given order_id",
	Example: `  alpaca order get --order-id 61e69015-8549-4baf-b96f-9c4f3e8d0c35`,
	Flags: []FlagDef{
		{Name: "nested", OASName: "nested", Type: "bool", Description: "if true, the result will roll up multi-leg orders under the legs field of primary order", Source: "query"},
		{Name: "order-id", OASName: "order_id", Type: "string", Description: "order id", Required: true, Source: "path"},
	},
}

var GetTokenizationRequestsOp = Op{
	Name: "GetTokenizationRequests", Summary: "List tokenization requests", ReturnsArray: true,
	Long: "An Authorized Participant can use this endpoint to list the tokenization requests performed on the Instant Tokenization Network (ITN)",
	Flags: []FlagDef{
		{Name: "after", OASName: "after", Type: "string", Description: "response will include only requests created after this timestamp (exclusive)", Source: "query"},
		{Name: "before", OASName: "before", Type: "string", Description: "response will include only requests created before this timestamp (exclusive)", Source: "query"},
		{Name: "issuer", OASName: "issuer", Type: "string", Description: "issuer of the tokenization requests to be queried", Completions: []string{"xstocks"}, Source: "query"},
		{Name: "network", OASName: "network", Type: "string", Description: "network of the tokenization requests to be queried", Completions: []string{"arbitrum", "base", "binance", "ethereum", "mantle", "solana", "ton", "tron"}, Source: "query"},
		{Name: "status", OASName: "status", Type: "string", Description: "tokenization request status to be queried", Completions: []string{"completed", "pending", "rejected"}, Source: "query"},
		{Name: "type", OASName: "type", Type: "string", Description: "tokenization request type to be queried", Completions: []string{"mint", "redeem"}, Source: "query"},
		{Name: "underlying-symbol", OASName: "underlying_symbol", Type: "string", Description: "underlying symbol of the tokenization requests to be queried", Source: "query"},
	},
}

var GetWatchlistByIDOp = Op{
	Name: "GetWatchlistByID", Summary: "Get watchlist by ID",
	Long:    "Returns a watchlist identified by the ID",
	Example: `  alpaca watchlist get --watchlist-id <id>`,
	Flags: []FlagDef{
		{Name: "watchlist-id", OASName: "watchlist_id", Type: "string", Description: "watchlist id", Required: true, Source: "path"},
	},
}

var GetWatchlistByNameOp = Op{
	Name: "GetWatchlistByName", Summary: "Get watchlist by name",
	Long:    "You can also call GET, PUT, POST and DELETE with watchlist name with another endpoint /v2/watchlists:by_name and query parameter name=<watchlist_name>, instead of /v2/watchlists/{watchlist_id} endpoints",
	Example: `  alpaca watchlist get-by-name --name "Tech Stocks"`,
	Flags: []FlagDef{
		{Name: "name", OASName: "name", Type: "string", Description: "name of the watchlist", Required: true, Source: "query"},
	},
}

var GetWatchlistsOp = Op{
	Name: "GetWatchlists", Summary: "Get all watchlists", ReturnsArray: true,
	Long:    "Returns the list of watchlists registered under the account",
	Example: `  alpaca watchlist list`,
}

var ListCryptoFundingTransfersOp = Op{
	Name: "ListCryptoFundingTransfers", Summary: "Retrieve crypto funding transfers",
	Long:    "Returns an array of all transfers associated with the given account across all wallets",
	Example: `  alpaca wallet transfer list`,
}

var ListCryptoFundingWalletsOp = Op{
	Name: "ListCryptoFundingWallets", Summary: "Retrieve crypto funding wallets",
	Long:    "Lists wallets for the account given in the path parameter. If an asset is specified and no wallet for the account and asset pair exists one will be created. If no asset is specified only existing wallets will be listed for the account. An account may have at most one wallet per asset",
	Example: `  alpaca wallet list`,
	Flags: []FlagDef{
		{Name: "asset", OASName: "asset", Type: "string", Description: "filter by crypto asset symbol, e.g. BTC, ETH, USDT. If specified and no wallet exists, one will be created", Source: "query"},
		{Name: "network", OASName: "network", Type: "string", Description: "optional network identifier", Completions: []string{"ethereum", "solana"}, Source: "query"},
	},
}

var ListCryptoPerpFundingTransfersOp = Op{
	Name: "ListCryptoPerpFundingTransfers", Summary: "Retrieve crypto funding transfers",
	Long:    "Returns an array of all transfers associated with the given account across all wallets",
	Example: `  alpaca crypto-perp wallet transfer list`,
}

var ListCryptoPerpFundingWalletsOp = Op{
	Name: "ListCryptoPerpFundingWallets", Summary: "Retrieve crypto funding wallets",
	Long:    "Lists wallets for the account given in the path parameter. If an asset is specified and no wallet for the account and asset pair exists one will be created. If no asset is specified only existing wallets will be listed for the account. An account may have at most one wallet per asset",
	Example: `  alpaca crypto-perp wallet list`,
	Flags: []FlagDef{
		{Name: "asset", OASName: "asset", Type: "string", Description: "asset", Source: "query"},
	},
}

var ListWhitelistedAddressOp = Op{
	Name: "ListWhitelistedAddress", Summary: "Get an array of whitelisted addresses",
	Example: `  alpaca wallet whitelist list`,
}

var ListWhitelistedPerpAddressOp = Op{
	Name: "ListWhitelistedPerpAddress", Summary: "Get an array of whitelisted addresses",
	Example: `  alpaca crypto-perp wallet whitelist list`,
}

var OptionBarsOp = Op{
	Name: "OptionBars", Summary: "Get historical bars",
	Long: "The historical option bars API provides aggregates for a list of option symbols between the specified dates",
	Example: `  alpaca data option bars --symbols AAPL250620C00200000 --start 2025-01-01
  alpaca data option bars --symbols AAPL250620C00200000,AAPL250620P00200000 --timeframe 1Day`,
	Flags: []FlagDef{
		{Name: "end", OASName: "end", Type: "string", Description: "inclusive end of the interval", Source: "query"},
		{Name: "limit", OASName: "limit", Type: "int", Default: "1000", Description: "maximum number of data points to return in the response page.", Source: "query"},
		{Name: "page-token", OASName: "page_token", Type: "string", Description: "pagination token from which to continue", Source: "query"},
		{Name: "sort", OASName: "sort", Type: "string", Default: "asc", Description: "sort data in ascending or descending order", Completions: []string{"asc", "desc"}, Source: "query"},
		{Name: "start", OASName: "start", Type: "string", Description: "inclusive start of the interval", Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "A comma-separated list of contract symbols with a limit of 100", Required: true, Source: "query"},
		{Name: "timeframe", OASName: "timeframe", Type: "string", Description: "timeframe represented by each bar in aggregation.\nYou can use any of the following values:\n - [1-59]Min or [1-59]T, e.g", Required: true, Source: "query"},
	},
}

var OptionDoNotExerciseOp = Op{
	Name: "OptionDoNotExercise", Summary: "Do not exercise an options position",
	Long:    "This endpoint enables users to submit a do-not-exercise (DNE) instruction for a held option contract, preventing automatic exercise at expiry. By default, Alpaca will automatically exercise in-the-money (ITM) contracts at expiry. This endpoint allows users to override that behavior",
	Example: `  alpaca option do-not-exercise --symbol-or-contract-id AAPL250620C00200000`,
	Flags: []FlagDef{
		{Name: "symbol-or-contract-id", OASName: "symbol_or_contract_id", Type: "string", Description: "option contract symbol or ID", Required: true, Source: "path"},
	},
}

var OptionExerciseOp = Op{
	Name: "OptionExercise", Summary: "Exercise an options position",
	Long:    "This endpoint enables users to exercise a held option contract, converting it into the underlying asset based on the specified terms. All available held shares of this option contract will be exercised. By default, Alpaca will automatically exercise in-the-money (ITM) contracts at expiry",
	Example: `  alpaca option exercise --symbol-or-contract-id AAPL250620C00200000`,
	Flags: []FlagDef{
		{Name: "symbol-or-contract-id", OASName: "symbol_or_contract_id", Type: "string", Description: "option contract symbol or ID", Required: true, Source: "path"},
	},
}

var PatchAccountConfigOp = Op{
	Name: "PatchAccountConfig", Summary: "Update account configurations",
	Long: "Updates and returns the current account configuration values",
	Example: `  alpaca account config set --no-shorting true
  alpaca account config set --dtbp-check entry`,
	Flags: []FlagDef{
		{Name: "disable-overnight-trading", OASName: "disable_overnight_trading", Type: "bool", Description: "if true, overnight trading is disabled", Source: "body"},
		{Name: "dtbp-check", OASName: "dtbp_check", Type: "string", Description: "both, entry, or exit. Controls Day Trading Margin Call (DTMC) checks", Completions: []string{"both", "entry", "exit"}, Source: "body"},
		{Name: "fractional-trading", OASName: "fractional_trading", Type: "bool", Description: "if true, account is able to participate in fractional trading", Source: "body"},
		{Name: "max-margin-multiplier", OASName: "max_margin_multiplier", Type: "string", Description: "can be \"1\", \"2\", or \"4\"", Source: "body"},
		{Name: "max-options-trading-level", OASName: "max_options_trading_level", Type: "int", Description: "desired maximum options trading level. 0=disabled, 1=Covered Call/Cash-Secured Put, 2=Long Call/Put, 3=Spreads/Straddles", Completions: []string{"0", "1", "2", "3"}, Source: "body"},
		{Name: "no-shorting", OASName: "no_shorting", Type: "bool", Description: "if true, account becomes long-only mode", Source: "body"},
		{Name: "pdt-check", OASName: "pdt_check", Type: "string", Description: "both, entry, or exit", Source: "body"},
		{Name: "ptp-no-exception-entry", OASName: "ptp_no_exception_entry", Type: "bool", Description: "if set to true then Alpaca will accept orders for PTP symbols with no exception. Default is false", Source: "body"},
		{Name: "suspend-trade", OASName: "suspend_trade", Type: "bool", Description: "if true, new orders are blocked", Source: "body"},
		{Name: "trade-confirm-email", OASName: "trade_confirm_email", Type: "string", Description: "all or none. If none, emails for order fills are not sent", Source: "body"},
	},
}

var PatchOrderByOrderIDOp = Op{
	Name: "PatchOrderByOrderID", Summary: "Replace order by ID",
	Long:    "Replaces a single order with updated parameters. Each parameter overrides the corresponding attribute of the existing order. The other attributes remain the same as the existing order",
	Example: `  alpaca order replace --order-id <id> --qty 20 --limit-price 190.00`,
	Flags: []FlagDef{
		{Name: "advanced-instructions", OASName: "advanced_instructions", Type: "string", Description: "advanced instructions for Elite Smart Router: https://docs.alpaca.markets/docs/alpaca-elite-smart-router", Source: "body"},
		{Name: "client-order-id", OASName: "client_order_id", Type: "string", Description: "A unique identifier for the new order. Automatically generated if not sent. (<= 128 characters)", Source: "body"},
		{Name: "limit-price", OASName: "limit_price", Type: "string", Description: "required if original order's type field was limit or stop_limit.", Source: "body"},
		{Name: "order-id", OASName: "order_id", Type: "string", Description: "order id", Required: true, Source: "path"},
		{Name: "qty", OASName: "qty", Type: "string", Description: "number of shares to trade.\n\nYou can only patch full shares for now.\n\nQty of equity fractional orders are not allowed ...", Source: "body"},
		{Name: "stop-price", OASName: "stop_price", Type: "string", Description: "required if original order type is limit or stop_limit", Source: "body"},
		{Name: "time-in-force", OASName: "time_in_force", Type: "string", Description: "time-In-Force values supported by Alpaca vary based on the order's security type", Completions: []string{"cls", "day", "fok", "gtc", "ioc", "opg"}, Source: "body"},
		{Name: "trail", OASName: "trail", Type: "string", Description: "the new value of the trail_price or trail_percent value (works only for type=“trailing_stop”)", Source: "body"},
	},
}

var PostOrderOp = Op{
	Name: "PostOrder", Summary: "Create an order",
	Long: "Places a new order for the given account. An order request may be rejected if the account is not authorized for trading, or if the tradable balance is insufficient to fill the order",
	Example: `  alpaca order submit --symbol AAPL --qty 10 --side buy --type market
  alpaca order submit --symbol AAPL --qty 5 --side buy --type limit --limit-price 185.00
  alpaca order submit --symbol AAPL --qty 10 --side sell --type stop --stop-price 175.00
  alpaca order submit --symbol AAPL --notional 1000 --side buy --type market`,
	Flags: []FlagDef{
		{Name: "advanced-instructions", OASName: "advanced_instructions", Type: "string", Description: "advanced instructions for Elite Smart Router: https://docs.alpaca.markets/docs/alpaca-elite-smart-router", Source: "body"},
		{Name: "client-order-id", OASName: "client_order_id", Type: "string", Description: "A unique identifier for the order. Automatically generated if not sent. (<= 128 characters)", Source: "body"},
		{Name: "extended-hours", OASName: "extended_hours", Type: "bool", Description: "(default) false", Source: "body"},
		{Name: "legs", OASName: "legs", Type: "string", Description: "list of order legs (<= 4)", Source: "body"},
		{Name: "limit-price", OASName: "limit_price", Type: "string", Description: "required if type is limit or stop_limit.", Source: "body"},
		{Name: "notional", OASName: "notional", Type: "string", Description: "dollar amount to trade. Cannot work with qty. Can only work for market order types and day for time in force", Source: "body"},
		{Name: "order-class", OASName: "order_class", Type: "string", Description: "order classes supported by Alpaca vary based on the order's security type", Completions: []string{"bracket", "mleg", "oco", "oto", "simple"}, Source: "body"},
		{Name: "position-intent", OASName: "position_intent", Type: "string", Description: "represents the desired position strategy", Completions: []string{"buy_to_close", "buy_to_open", "sell_to_close", "sell_to_open"}, Source: "body"},
		{Name: "qty", OASName: "qty", Type: "string", Description: "number of shares to trade", Source: "body"},
		{Name: "side", OASName: "side", Type: "string", Description: "represents which side this order was on:\n- buy\n- sell\nRequired for all order classes except for mleg", Completions: []string{"buy", "sell"}, Source: "body"},
		{Name: "stop-loss", OASName: "stop_loss", Type: "string", Description: "takes in string/number values for stop_price and limit_price", Source: "body"},
		{Name: "stop-price", OASName: "stop_price", Type: "string", Description: "required if type is stop or stop_limit", Source: "body"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol, asset ID, or currency pair to identify the asset to trade, required for all order classes except for mleg", Source: "body"},
		{Name: "take-profit", OASName: "take_profit", Type: "string", Description: "takes in a string/number value for limit_price", Source: "body"},
		{Name: "time-in-force", OASName: "time_in_force", Type: "string", Description: "time-In-Force values supported by Alpaca vary based on the order's security type", Completions: []string{"cls", "day", "fok", "gtc", "ioc", "opg"}, Source: "body"},
		{Name: "trail-percent", OASName: "trail_percent", Type: "string", Description: "this or trail_price is required if type is trailing_stop", Source: "body"},
		{Name: "trail-price", OASName: "trail_price", Type: "string", Description: "this or trail_percent is required if type is trailing_stop", Source: "body"},
		{Name: "type", OASName: "type", Type: "string", Default: "market", Description: "order types supported by Alpaca vary based on the order's security type", Completions: []string{"limit", "market", "stop", "stop_limit", "trailing_stop"}, Source: "body"},
	},
}

var PostTokenizationMintOp = Op{
	Name: "PostTokenizationMint", Summary: "Create mint a tokenized asset",
	Long: "This endpoint is used by an Authorized Participant to request the minting of a tokenized asset",
	Flags: []FlagDef{
		{Name: "issuer", OASName: "issuer", Type: "string", Description: "tokenized asset's issuer", Completions: []string{"st0x", "xstocks"}, Source: "body"},
		{Name: "network", OASName: "network", Type: "string", Description: "token's blockchain network", Completions: []string{"arbitrum", "base", "binance", "ethereum", "mantle", "solana", "ton", "tron"}, Source: "body"},
		{Name: "qty", OASName: "qty", Type: "string", Description: "underlying quantity to convert into the tokenized asset. It can be fractional", Source: "body"},
		{Name: "underlying-symbol", OASName: "underlying_symbol", Type: "string", Description: "underlying asset symbol", Source: "body"},
		{Name: "wallet-address", OASName: "wallet_address", Type: "string", Description: "wallet address to receive the tokenized asset", Source: "body"},
	},
}

var PostWatchlistOp = Op{
	Name: "PostWatchlist", Summary: "Create watchlist",
	Long:    "Create a new watchlist with initial set of assets",
	Example: `  alpaca watchlist create --name "Tech Stocks" --symbols AAPL,MSFT,GOOG`,
	Flags: []FlagDef{
		{Name: "name", OASName: "name", Type: "string", Description: "watchlist name", Source: "body"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "list of asset symbols to include in the watchlist", Source: "body"},
	},
}

var RemoveAssetFromWatchlistOp = Op{
	Name: "RemoveAssetFromWatchlist", Summary: "Delete symbol from watchlist",
	Long:    "Delete one entry for an asset by symbol name",
	Example: `  alpaca watchlist remove --watchlist-id <id> --symbol AAPL`,
	Flags: []FlagDef{
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol name to remove from the watchlist content", Required: true, Source: "path"},
		{Name: "watchlist-id", OASName: "watchlist_id", Type: "string", Description: "watchlist ID", Required: true, Source: "path"},
	},
}

var SetCryptoPerpAccountLeverageOp = Op{
	Name: "SetCryptoPerpAccountLeverage", Summary: "Set account leverage for an asset",
	Long:    "Updates the leverage for the crypto perpetuals account for a specific underlying asset. Provide the 'symbol' of the asset and the desired 'leverage' (as an integer) using query parameters. The system will return the asset's symbol and the newly set leverage value upon successful update",
	Example: `  alpaca crypto-perp set-leverage --asset BTC --leverage 5`,
	Flags: []FlagDef{
		{Name: "leverage", OASName: "leverage", Type: "int", Description: "leverage for the underlying asset", Source: "query"},
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "symbol of underlying asset", Source: "query"},
	},
}

var SubscribeToActivitiesSSEOp = Op{
	Name: "SubscribeToActivitiesSSE", Summary: "Get subscribe to activity events (SSE)",
	Long: "The Events API sends the real-time events and provides historical queries with SSE (Server Sent Events)",
	Flags: []FlagDef{
		{Name: "since", OASName: "since", Type: "string", Description: "format: RFC3339 or YYYY-MM-DD", Source: "query"},
		{Name: "since-id", OASName: "since_id", Type: "string", Description: "since id", Source: "query"},
		{Name: "until", OASName: "until", Type: "string", Description: "format: RFC3339 or YYYY-MM-DD", Source: "query"},
		{Name: "until-id", OASName: "until_id", Type: "string", Description: "until id", Source: "query"},
	},
}

var UpdateWatchlistByIDOp = Op{
	Name: "UpdateWatchlistByID", Summary: "Update watchlist by id",
	Long:    "Update the name and/or content of watchlist",
	Example: `  alpaca watchlist update --watchlist-id <id> --name "Updated" --symbols AAPL,MSFT`,
	Flags: []FlagDef{
		{Name: "name", OASName: "name", Type: "string", Description: "watchlist name", Source: "body"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "list of asset symbols to include in the watchlist", Source: "body"},
		{Name: "watchlist-id", OASName: "watchlist_id", Type: "string", Description: "watchlist id", Required: true, Source: "path"},
	},
}

var UpdateWatchlistByNameOp = Op{
	Name: "UpdateWatchlistByName", Summary: "Update watchlist by name",
	Long:    "Update the name and/or content of watchlist",
	Example: `  alpaca watchlist update-by-name --name "Tech Stocks" --new-name "Technology" --symbols AAPL,MSFT`,
	Flags: []FlagDef{
		{Name: "name", OASName: "name", Type: "string", Description: "name of the watchlist", Required: true, Source: "query"},
		{Name: "symbols", OASName: "symbols", Type: "string", Description: "list of asset symbols to include in the watchlist", Source: "body"},
	},
}

var accountConfigurationsResponseFields = []ResponseField{
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
}

var assetsResponseFields = []ResponseField{
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
}

var cryptoLatestBarsRespResponseFields = []ResponseField{
	{Name: "bars", Type: "map[string]object", Description: "bars"},
}

var cryptoLatestOrderbooksRespResponseFields = []ResponseField{
	{Name: "orderbooks", Type: "map[string]object", Description: "orderbooks"},
}

var cryptoLatestQuotesRespResponseFields = []ResponseField{
	{Name: "quotes", Type: "map[string]object", Description: "quotes"},
}

var cryptoLatestTradesRespResponseFields = []ResponseField{
	{Name: "trades", Type: "map[string]object", Description: "trades"},
}

var cryptoTransferResponseFields = []ResponseField{
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
}

var cryptoWalletResponseFields = []ResponseField{
	{Name: "address", Type: "string", Description: "address"},
	{Name: "chain", Type: "string", Description: "chain"},
	{Name: "created_at", Type: "string", Description: "timestamp (RFC3339) of account creation"},
}

var optionSnapshotsRespResponseFields = []ResponseField{
	{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
	{Name: "snapshots", Type: "map[string]object", Description: "snapshots"},
}

var orderResponseFields = []ResponseField{
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
}

var positionResponseFields = []ResponseField{
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
}

var watchlistResponseFields = []ResponseField{
	{Name: "account_id", Type: "string", Description: "account ID"},
	{Name: "assets", Type: "[]object", Description: "the content of this watchlist, in the order as registered by the client"},
	{Name: "created_at", Type: "string", Description: "created at"},
	{Name: "id", Type: "string", Description: "watchlist id"},
	{Name: "name", Type: "string", Description: "user-defined watchlist name (up to 64 characters)"},
	{Name: "updated_at", Type: "string", Description: "updated at"},
}

var whitelistedAddressResponseFields = []ResponseField{
	{Name: "address", Type: "string", Description: "whitelisted address"},
	{Name: "asset", Type: "string", Description: "symbol of underlying asset for the whitelisted address"},
	{Name: "chain", Type: "string", Description: "underlying network this address represents"},
	{Name: "created_at", Type: "string", Description: "timestamp (RFC3339) of account creation"},
	{Name: "id", Type: "string", Description: "unique ID for whitelisted address"},
	{Name: "status", Type: "enum", Description: "status of whitelisted address which is either ACTIVE or PENDING", EnumValues: []string{"APPROVED", "PENDING"}},
}

var (
	responseSchemas     map[string][]ResponseField
	responseSchemasOnce sync.Once
)

// ResponseSchema returns response fields for an operation (lazy-loaded).
func ResponseSchema(opName string) ([]ResponseField, bool) {
	responseSchemasOnce.Do(func() {
		responseSchemas = map[string][]ResponseField{
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
			"CryptoLatestBars":       cryptoLatestBarsRespResponseFields,
			"CryptoLatestOrderbooks": cryptoLatestOrderbooksRespResponseFields,
			"CryptoLatestQuotes":     cryptoLatestQuotesRespResponseFields,
			"CryptoLatestTrades":     cryptoLatestTradesRespResponseFields,
			"CryptoPerpLatestBars":   cryptoLatestBarsRespResponseFields,
			"CryptoPerpLatestFuturesPricing": {
				{Name: "pricing", Type: "map[string]object", Description: "pricing"},
			},
			"CryptoPerpLatestOrderbooks": cryptoLatestOrderbooksRespResponseFields,
			"CryptoPerpLatestQuotes":     cryptoLatestQuotesRespResponseFields,
			"CryptoPerpLatestTrades":     cryptoLatestTradesRespResponseFields,
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
			"OptionChain": optionSnapshotsRespResponseFields,
			"OptionLatestQuotes": {
				{Name: "quotes", Type: "map[string]object", Description: "quotes"},
			},
			"OptionLatestTrades": {
				{Name: "trades", Type: "map[string]object", Description: "trades"},
			},
			"OptionSnapshots": optionSnapshotsRespResponseFields,
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
			"AddAssetToWatchlist":                watchlistResponseFields,
			"AddAssetToWatchlistByName":          watchlistResponseFields,
			"CreateCryptoPerpTransferForAccount": cryptoTransferResponseFields,
			"CreateCryptoTransferForAccount":     cryptoTransferResponseFields,
			"CreateWhitelistedAddress":           whitelistedAddressResponseFields,
			"CreateWhitelistedPerpAddress":       whitelistedAddressResponseFields,
			"DeleteAllOpenPositions": {
				{Name: "body", Type: "object", Description: "orders API allows a user to monitor, place and cancel their orders with Alpaca.\n\nEach order has a unique identifier p..."},
				{Name: "status", Type: "integer", Description: "HTTP status code for the attempt to close this position"},
				{Name: "symbol", Type: "string", Description: "symbol name of the asset"},
			},
			"DeleteAllOrders": {
				{Name: "id", Type: "string", Description: "orderId"},
				{Name: "status", Type: "integer", Description: "http response code"},
			},
			"DeleteOpenPosition": orderResponseFields,
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
			"GetV2Assets":                assetsResponseFields,
			"GetV2AssetsSymbolOrAssetID": assetsResponseFields,
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
			"GetAccountConfig": accountConfigurationsResponseFields,
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
			"GetAllOpenPositions":          positionResponseFields,
			"GetAllOrders":                 orderResponseFields,
			"GetCryptoFundingTransfer":     cryptoTransferResponseFields,
			"GetCryptoPerpFundingTransfer": cryptoTransferResponseFields,
			"GetOpenPosition":              positionResponseFields,
			"GetOrderByClientOrderID":      orderResponseFields,
			"GetOrderByOrderID":            orderResponseFields,
			"GetTokenizationRequests": {
				{Name: "account", Type: "string", Description: "alpaca account ID associated with this tokenization request"},
				{Name: "created_at", Type: "string", Description: "created at"},
				{Name: "fees", Type: "string", Description: "fees charged for this tokenization request"},
				{Name: "issuer", Type: "enum", Description: "tokenized asset's issuer", EnumValues: []string{"st0x", "xstocks"}},
				{Name: "issuer_account", Type: "string", Description: "issuer's account ID associated with this tokenization request"},
				{Name: "issuer_request_id", Type: "string", Description: "unique identifier of the tokenization request set by the issuer"},
				{Name: "network", Type: "enum", Description: "token's blockchain network", EnumValues: []string{"arbitrum", "base", "binance", "ethereum", "mantle", "solana", "ton", "tron"}},
				{Name: "qty", Type: "string", Description: "quantity to convert for this tokenization request. It can be fractional"},
				{Name: "status", Type: "enum", Description: "status of the tokenization request", EnumValues: []string{"completed", "pending", "rejected"}},
				{Name: "token_symbol", Type: "string", Description: "tokenized asset symbol"},
				{Name: "tokenization_request_id", Type: "string", Description: "unique identifier of the tokenization request set by Alpaca"},
				{Name: "tx_hash", Type: "string", Description: "transaction hash of the completed request on the blockchain"},
				{Name: "type", Type: "enum", Description: "tokenization request type", EnumValues: []string{"mint", "redeem"}},
				{Name: "underlying_symbol", Type: "string", Description: "underlying asset symbol"},
				{Name: "updated_at", Type: "string", Description: "updated at"},
				{Name: "wallet_address", Type: "string", Description: "wallet address associated with this tokenization request"},
			},
			"GetWatchlistByID":   watchlistResponseFields,
			"GetWatchlistByName": watchlistResponseFields,
			"GetWatchlists": {
				{Name: "account_id", Type: "string", Description: "account ID"},
				{Name: "created_at", Type: "string", Description: "created at"},
				{Name: "id", Type: "string", Description: "watchlist id"},
				{Name: "name", Type: "string", Description: "user-defined watchlist name (up to 64 characters)"},
				{Name: "updated_at", Type: "string", Description: "updated at"},
			},
			"ListCryptoFundingTransfers":     cryptoTransferResponseFields,
			"ListCryptoFundingWallets":       cryptoWalletResponseFields,
			"ListCryptoPerpFundingTransfers": cryptoTransferResponseFields,
			"ListCryptoPerpFundingWallets":   cryptoWalletResponseFields,
			"ListWhitelistedAddress":         whitelistedAddressResponseFields,
			"ListWhitelistedPerpAddress":     whitelistedAddressResponseFields,
			"OptionBars": {
				{Name: "bars", Type: "map[string][]object", Description: "bars"},
				{Name: "currency", Type: "string", Description: "currency"},
				{Name: "next_page_token", Type: "string", Description: "pagination token for the next page"},
			},
			"PatchAccountConfig":  accountConfigurationsResponseFields,
			"PatchOrderByOrderID": orderResponseFields,
			"PostOrder":           orderResponseFields,
			"PostTokenizationMint": {
				{Name: "created_at", Type: "string", Description: "created at"},
				{Name: "issuer", Type: "enum", Description: "tokenized asset's issuer", EnumValues: []string{"st0x", "xstocks"}},
				{Name: "network", Type: "enum", Description: "token's blockchain network", EnumValues: []string{"arbitrum", "base", "binance", "ethereum", "mantle", "solana", "ton", "tron"}},
				{Name: "qty", Type: "string", Description: "quantity to convert for this tokenization request. It can be fractional"},
				{Name: "status", Type: "enum", Description: "status of the tokenization request", EnumValues: []string{"completed", "pending", "rejected"}},
				{Name: "token_symbol", Type: "string", Description: "tokenized asset symbol"},
				{Name: "tokenization_request_id", Type: "string", Description: "unique identifier of the tokenization request set by Alpaca"},
				{Name: "underlying_symbol", Type: "string", Description: "underlying asset symbol"},
			},
			"PostWatchlist":            watchlistResponseFields,
			"RemoveAssetFromWatchlist": watchlistResponseFields,
			"UpdateWatchlistByID":      watchlistResponseFields,
			"UpdateWatchlistByName":    watchlistResponseFields,
		}
	})
	fields, ok := responseSchemas[opName]
	return fields, ok
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
	GetTokenizationRequestsOp,
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
	PostTokenizationMintOp,
	PostWatchlistOp,
	RemoveAssetFromWatchlistOp,
	SetCryptoPerpAccountLeverageOp,
	SubscribeToActivitiesSSEOp,
	UpdateWatchlistByIDOp,
	UpdateWatchlistByNameOp,
}

var opByName map[string]Op

func init() {
	opByName = make(map[string]Op, len(AllOps))
	for _, op := range AllOps {
		opByName[op.Name] = op
	}
}

// OpByName returns the Op with the given name, if any.
func OpByName(name string) (Op, bool) {
	op, ok := opByName[name]
	return op, ok
}
