package main

import (
	"bytes"
	"fmt"
	"log"
	"sort"
	"strings"
)

type cmdDef struct {
	parent      string
	use         string
	self        bool
	examples    string
	long        string
	jsonOnly    bool
	defaults    map[string]string
	normalize   []string
	bodyAliases map[string]string // body field kebab name → CLI alias (resolves flag collisions)
	rawMethod   string            // overrides client method; signals (string, url.Values) call signature
}

type parentDef struct {
	use     string
	short   string
	parent  string
	aliases []string
}

var cmdParents = map[string]parentDef{
	"account":       {use: "account", short: "Manage your trading account"},
	"accountConfig": {use: "config", short: "Manage account configuration", parent: "account"},
	"activity":      {use: "activity", short: "Account activities (fills, dividends, transfers, etc.)", parent: "account"},
	"asset":         {use: "asset", short: "Browse assets"},
	"position":      {use: "position", short: "Manage positions"},
	"corporateAction": {
		use: "corporate-action", short: "Corporate actions announcements",
		aliases: []string{"ca"},
	},
	"option":          {use: "option", short: "Options trading"},
	"order":           {use: "order", short: "Manage orders"},
	"wallet":          {use: "wallet", short: "Crypto funding wallets and transfers"},
	"walletTransfer":  {use: "transfer", short: "Manage crypto transfers", parent: "wallet"},
	"walletWhitelist": {use: "whitelist", short: "Manage whitelisted crypto addresses", parent: "wallet"},
	"clock":           {use: "clock", short: "Market clock"},
	"calendar":        {use: "calendar", short: "Market calendar"},
	"data":            {use: "data", short: "Access market data"},
	"dataOption":      {use: "option", short: "Options market data", parent: "data"},
	"dataForex":       {use: "forex", short: "Foreign exchange rate data", parent: "data"},
	"dataLatest":      {use: "latest", short: "Get latest market data", parent: "data"},
	"dataMeta":        {use: "meta", short: "Stock exchange and condition reference data", parent: "data"},
	"screener":        {use: "screener", short: "Stock and crypto screener and market movers", parent: "data"},
	"watchlist":       {use: "watchlist", short: "Manage watchlists"},
	"cryptoPerp":      {use: "crypto-perp", short: "Crypto perpetuals (futures)"},
	"cryptoPerpWallet": {
		use: "wallet", short: "Crypto perpetuals funding wallets and transfers",
		parent: "cryptoPerp",
	},
	"cryptoPerpTransfer": {
		use: "transfer", short: "Manage crypto perpetuals transfers",
		parent: "cryptoPerpWallet",
	},
	"cryptoPerpWhitelist": {
		use: "whitelist", short: "Manage whitelisted perpetuals crypto addresses",
		parent: "cryptoPerpWallet",
	},
	"cryptoPerpData": {
		use: "data", short: "Crypto perpetuals market data",
		parent: "cryptoPerp",
	},
	"dataCrypto": {use: "crypto", short: "Crypto market data", parent: "data"},
}

var cmdRegistry = map[string]cmdDef{
	// --- account ---
	"GetAccount": {
		parent:   "account",
		use:      "get",
		examples: "  alpaca account get",
	},
	"GetAccountConfig": {
		parent:   "accountConfig",
		use:      "get",
		examples: "  alpaca account config get",
	},
	"PatchAccountConfig": {
		parent: "accountConfig",
		use:    "set",
		examples: `  alpaca account config set --no-shorting true
  alpaca account config set --dtbp-check entry`,
	},
	"GetAccountActivities": {
		parent: "activity",
		use:    "list",
		examples: `  alpaca account activity list
  alpaca account activity list --activity-types FILL,TRANS --page-size 20
  alpaca account activity list --direction desc`,
	},
	"GetAccountActivitiesByActivityType": {
		parent: "activity",
		use:    "list-by-type",
		examples: `  alpaca account activity list-by-type --activity-type FILL --page-size 20
  alpaca account activity list-by-type --activity-type DIV --after 2025-01-01`,
	},
	"GetAccountPortfolioHistory": {
		parent:   "account",
		use:      "portfolio",
		jsonOnly: true,
		long:     "Returns portfolio equity and P&L history. Output is always JSON due to complex time-series structure.",
		examples: `  alpaca account portfolio
  alpaca account portfolio --period 1M --timeframe 1D`,
	},

	// --- asset ---
	"GetV2Assets": {
		parent: "asset",
		use:    "list",
		examples: `  alpaca asset list
  alpaca asset list --asset-class us_equity --status active
  alpaca asset list --exchange NYSE`,
	},
	"GetV2AssetsSymbolOrAssetID": {
		parent: "asset",
		use:    "get",
		examples: `  alpaca asset get --symbol-or-asset-id AAPL
  alpaca asset get --symbol-or-asset-id BTC/USD`,
	},
	"UsTreasuries": {
		parent: "asset",
		use:    "treasury",
		examples: `  alpaca asset treasury
  alpaca asset treasury --bond-status active`,
	},
	"UsCorporates": {
		parent: "asset",
		use:    "bond",
		examples: `  alpaca asset bond
  alpaca asset bond --bond-status active`,
	},

	// --- corporate action ---
	"GetV2CorporateActionsAnnouncements": {
		parent: "corporateAction",
		use:    "list",
		examples: `  alpaca corporate-action list --ca-types reverse_split --since 2025-01-01 --until 2025-12-31
  alpaca corporate-action list --ca-types cash_dividend --symbol AAPL --since 2025-01-01 --until 2025-06-30`,
	},
	"GetV2CorporateActionsAnnouncementsID": {
		parent:   "corporateAction",
		use:      "get",
		examples: "  alpaca corporate-action get --id <announcement-id>",
	},

	// --- order (trivial ones) ---
	"GetAllOrders": {
		parent:   "order",
		use:      "list",
		defaults: map[string]string{"status": "open"},
		examples: `  alpaca order list
  alpaca order list --status closed --limit 20
  alpaca order list --symbols AAPL,MSFT --after 2025-01-01`,
	},
	"GetOrderByOrderID": {
		parent:   "order",
		use:      "get",
		examples: "  alpaca order get --order-id 61e69015-8549-4baf-b96f-9c4f3e8d0c35",
	},
	"GetOrderByClientOrderID": {
		parent:   "order",
		use:      "get-by-client-id",
		examples: "  alpaca order get-by-client-id --client-order-id my-order-123",
	},
	"DeleteOrderByOrderID": {
		parent:   "order",
		use:      "cancel",
		examples: "  alpaca order cancel --order-id <id>",
	},
	"PatchOrderByOrderID": {
		parent:   "order",
		use:      "replace",
		examples: "  alpaca order replace --order-id <id> --qty 20 --limit-price 190.00",
	},
	"DeleteAllOrders": {
		parent:   "order",
		use:      "cancel-all",
		examples: "  alpaca order cancel-all",
	},

	// --- clock / calendar (self: true = parent IS the command) ---
	"LegacyClock": {
		parent:   "clock",
		self:     true,
		examples: "  alpaca clock",
	},
	"Clock": {
		parent:   "clock",
		use:      "markets",
		jsonOnly: true,
		examples: "  alpaca clock markets --markets XNYS,XNAS",
	},
	"LegacyCalendar": {
		parent: "calendar",
		self:   true,
		examples: `  alpaca calendar
  alpaca calendar --start 2025-01-01 --end 2025-12-31`,
	},
	"Calendar": {
		parent:   "calendar",
		use:      "market",
		examples: "  alpaca calendar market --market XNYS --start 2025-01-01",
	},

	// --- position ---
	"GetAllOpenPositions": {
		parent: "position",
		use:    "list",
		examples: `  alpaca position list
  alpaca position list --csv`,
	},
	"GetOpenPosition": {
		parent:    "position",
		use:       "get",
		normalize: []string{"symbol-or-asset-id"},
		examples: `  alpaca position get --symbol-or-asset-id AAPL
  alpaca position get --symbol-or-asset-id BTC/USD`,
	},
	"DeleteOpenPosition": {
		parent:    "position",
		use:       "close",
		normalize: []string{"symbol-or-asset-id"},
		examples: `  alpaca position close --symbol-or-asset-id AAPL
  alpaca position close --symbol-or-asset-id AAPL --qty 5
  alpaca position close --symbol-or-asset-id AAPL --percentage 50`,
	},
	"DeleteAllOpenPositions": {
		parent:   "position",
		use:      "close-all",
		examples: "  alpaca position close-all",
	},

	// --- option ---
	"GetOptionsContracts": {
		parent: "option",
		use:    "contracts",
		long:   "List option contracts for an underlying symbol. For market data (greeks, pricing), use `data option chain`.",
		examples: `  alpaca option contracts --underlying-symbols AAPL
  alpaca option contracts --underlying-symbols AAPL --expiration-date 2025-06-20 --type call
  alpaca option contracts --underlying-symbols SPY --strike-price-gte 400 --strike-price-lte 450`,
	},
	"GetOptionContractSymbolOrID": {
		parent:   "option",
		use:      "get",
		examples: "  alpaca option get --symbol-or-id AAPL250620C00200000",
	},
	"OptionExercise": {
		parent:   "option",
		use:      "exercise",
		examples: "  alpaca option exercise --symbol-or-contract-id AAPL250620C00200000",
	},
	"OptionDoNotExercise": {
		parent:   "option",
		use:      "do-not-exercise",
		examples: "  alpaca option do-not-exercise --symbol-or-contract-id AAPL250620C00200000",
	},

	// --- wallet ---
	"ListCryptoFundingWallets": {
		parent:   "wallet",
		use:      "list",
		examples: "  alpaca wallet list",
	},
	"ListCryptoFundingTransfers": {
		parent:   "walletTransfer",
		use:      "list",
		examples: "  alpaca wallet transfer list",
	},
	"GetCryptoFundingTransfer": {
		parent:   "walletTransfer",
		use:      "get",
		examples: "  alpaca wallet transfer get --transfer-id <id>",
	},
	"CreateCryptoTransferForAccount": {
		parent:   "walletTransfer",
		use:      "create",
		examples: "  alpaca wallet transfer create --amount 0.5 --address 0xabc... --asset BTC",
	},
	"GetCryptoTransferEstimate": {
		parent: "walletTransfer",
		use:    "estimate",
		examples: `  alpaca wallet transfer estimate --asset BTC --amount 0.5 \
    --from-address 0xabc... --to-address 0xdef...`,
	},
	"ListWhitelistedAddress": {
		parent:   "walletWhitelist",
		use:      "list",
		examples: "  alpaca wallet whitelist list",
	},
	"CreateWhitelistedAddress": {
		parent:   "walletWhitelist",
		use:      "add",
		examples: "  alpaca wallet whitelist add --address 0xabc... --asset ETH",
	},
	"DeleteWhitelistedAddress": {
		parent:   "walletWhitelist",
		use:      "delete",
		examples: "  alpaca wallet whitelist delete --whitelisted-address-id <id>",
	},

	// --- crypto perp ---
	"GetCryptoPerpAccountVitals": {
		parent:   "cryptoPerp",
		use:      "vitals",
		jsonOnly: true,
		examples: "  alpaca crypto-perp vitals",
	},
	"GetCryptoPerpAccountLeverage": {
		parent:   "cryptoPerp",
		use:      "leverage",
		jsonOnly: true,
		examples: "  alpaca crypto-perp leverage",
	},
	"SetCryptoPerpAccountLeverage": {
		parent:   "cryptoPerp",
		use:      "set-leverage",
		jsonOnly: true,
		examples: "  alpaca crypto-perp set-leverage --asset BTC --leverage 5",
	},
	"ListCryptoPerpFundingWallets": {
		parent:   "cryptoPerpWallet",
		use:      "list",
		examples: "  alpaca crypto-perp wallet list",
	},
	"GetCryptoPerpTransferEstimate": {
		parent:   "cryptoPerpTransfer",
		use:      "estimate",
		jsonOnly: true,
		examples: "  alpaca crypto-perp wallet transfer estimate --asset BTC --amount 0.5",
	},
	"ListCryptoPerpFundingTransfers": {
		parent:   "cryptoPerpTransfer",
		use:      "list",
		examples: "  alpaca crypto-perp wallet transfer list",
	},
	"CreateCryptoPerpTransferForAccount": {
		parent:   "cryptoPerpTransfer",
		use:      "create",
		examples: "  alpaca crypto-perp wallet transfer create --amount 0.5 --address 0xabc... --asset BTC",
	},
	"GetCryptoPerpFundingTransfer": {
		parent:   "cryptoPerpTransfer",
		use:      "get",
		examples: "  alpaca crypto-perp wallet transfer get --transfer-id <id>",
	},
	"ListWhitelistedPerpAddress": {
		parent:   "cryptoPerpWhitelist",
		use:      "list",
		examples: "  alpaca crypto-perp wallet whitelist list",
	},
	"CreateWhitelistedPerpAddress": {
		parent:   "cryptoPerpWhitelist",
		use:      "add",
		examples: "  alpaca crypto-perp wallet whitelist add --address 0xabc... --asset ETH",
	},
	"DeleteWhitelistedPerpAddress": {
		parent:   "cryptoPerpWhitelist",
		use:      "delete",
		examples: "  alpaca crypto-perp wallet whitelist delete --whitelisted-address-id <id>",
	},

	// --- watchlist ---
	"GetWatchlists": {
		parent:   "watchlist",
		use:      "list",
		examples: "  alpaca watchlist list",
	},
	"GetWatchlistByID": {
		parent:   "watchlist",
		use:      "get",
		examples: "  alpaca watchlist get --watchlist-id <id>",
	},
	"PostWatchlist": {
		parent:   "watchlist",
		use:      "create",
		examples: `  alpaca watchlist create --name "Tech Stocks" --symbols AAPL,MSFT,GOOG`,
	},
	"UpdateWatchlistByID": {
		parent:   "watchlist",
		use:      "update",
		examples: `  alpaca watchlist update --watchlist-id <id> --name "Updated" --symbols AAPL,MSFT`,
	},
	"DeleteWatchlistByID": {
		parent:   "watchlist",
		use:      "delete",
		examples: "  alpaca watchlist delete --watchlist-id <id>",
	},
	"AddAssetToWatchlist": {
		parent:   "watchlist",
		use:      "add",
		examples: "  alpaca watchlist add --watchlist-id <id> --symbol AAPL",
	},
	"RemoveAssetFromWatchlist": {
		parent:   "watchlist",
		use:      "remove",
		examples: "  alpaca watchlist remove --watchlist-id <id> --symbol AAPL",
	},
	"GetWatchlistByName": {
		parent:   "watchlist",
		use:      "get-by-name",
		examples: `  alpaca watchlist get-by-name --name "Tech Stocks"`,
	},
	"UpdateWatchlistByName": {
		parent:      "watchlist",
		use:         "update-by-name",
		bodyAliases: map[string]string{"name": "new-name"},
		examples:    `  alpaca watchlist update-by-name --name "Tech Stocks" --new-name "Technology" --symbols AAPL,MSFT`,
	},
	"DeleteWatchlistByName": {
		parent:   "watchlist",
		use:      "delete-by-name",
		examples: `  alpaca watchlist delete-by-name --name "Tech Stocks"`,
	},
	"AddAssetToWatchlistByName": {
		parent:   "watchlist",
		use:      "add-by-name",
		examples: `  alpaca watchlist add-by-name --name "Tech Stocks" --symbol NVDA`,
	},

	// --- data: single-symbol (unified routing) ---
	"StockBarSingle": {
		parent:    "data",
		use:       "bars",
		rawMethod: "Bars",
		examples: `  alpaca data bars --symbol AAPL --start 2025-01-01 --timeframe 1Day
  alpaca data bars --symbol BTC/USD --start 2025-01-01 --timeframe 1Hour
  alpaca data bars --symbol AAPL --start 2025-01-01 --end 2025-06-01 --limit 100`,
	},
	"StockQuoteSingle": {
		parent:    "data",
		use:       "quotes",
		rawMethod: "Quotes",
		examples: `  alpaca data quotes --symbol AAPL --start 2025-01-01
  alpaca data quotes --symbol AAPL --start 2025-01-01 --end 2025-01-31 --limit 50`,
	},
	"StockTradeSingle": {
		parent:    "data",
		use:       "trades",
		rawMethod: "Trades",
		examples: `  alpaca data trades --symbol AAPL --start 2025-01-01
  alpaca data trades --symbol AAPL --start 2025-01-01 --limit 100`,
	},
	"StockSnapshotSingle": {
		parent:    "data",
		use:       "snapshot",
		rawMethod: "Snapshot",
		jsonOnly:  true,
		long:      "Returns the latest snapshot for a symbol. Output is always JSON due to complex nested structure.",
		examples: `  alpaca data snapshot --symbol AAPL
  alpaca data snapshot --symbol BTC/USD --feed sip`,
	},
	"StockLatestTradeSingle": {
		parent:    "dataLatest",
		use:       "trade",
		rawMethod: "LatestTrade",
		examples: `  alpaca data latest trade --symbol AAPL
  alpaca data latest trade --symbol AAPL --feed sip`,
	},
	"StockLatestQuoteSingle": {
		parent:    "dataLatest",
		use:       "quote",
		rawMethod: "LatestQuote",
		examples:  "  alpaca data latest quote --symbol AAPL",
	},
	"StockLatestBarSingle": {
		parent:    "dataLatest",
		use:       "bar",
		rawMethod: "LatestBar",
		examples: `  alpaca data latest bar --symbol AAPL
  alpaca data latest bar --symbol BTC/USD`,
	},

	// --- data: news ---
	"News": {
		parent: "data",
		use:    "news",
		examples: `  alpaca data news
  alpaca data news --symbols AAPL,MSFT --limit 10`,
	},

	// --- data: option ---
	"OptionBars": {
		parent:   "dataOption",
		use:      "bars",
		jsonOnly: true,
		examples: `  alpaca data option bars --symbols AAPL250620C00200000 --start 2025-01-01
  alpaca data option bars --symbols AAPL250620C00200000,AAPL250620P00200000 --timeframe 1Day`,
	},
	"OptionTrades": {
		parent:   "dataOption",
		use:      "trades",
		jsonOnly: true,
		examples: "  alpaca data option trades --symbols AAPL250620C00200000 --start 2025-01-01",
	},
	"OptionSnapshots": {
		parent:   "dataOption",
		use:      "snapshot",
		jsonOnly: true,
		examples: `  alpaca data option snapshot --symbols AAPL250620C00200000
  alpaca data option snapshot --symbols AAPL250620C00200000,AAPL250620P00200000`,
	},
	"OptionChain": {
		parent:   "dataOption",
		use:      "chain",
		jsonOnly: true,
		examples: `  alpaca data option chain --underlying-symbol AAPL
  alpaca data option chain --underlying-symbol SPY --expiration-date 2025-06-20 --type call`,
	},
	"OptionLatestQuotes": {
		parent:   "dataOption",
		use:      "latest-quotes",
		jsonOnly: true,
		examples: "  alpaca data option latest-quotes --symbols AAPL250620C00200000",
	},
	"OptionLatestTrades": {
		parent:   "dataOption",
		use:      "latest-trades",
		jsonOnly: true,
		examples: "  alpaca data option latest-trades --symbols AAPL250620C00200000",
	},
	"OptionMetaExchanges": {
		parent:   "dataOption",
		use:      "exchanges",
		jsonOnly: true,
		examples: "  alpaca data option exchanges",
	},
	"OptionMetaConditions": {
		parent:   "dataOption",
		use:      "conditions",
		jsonOnly: true,
		examples: "  alpaca data option conditions --ticktype trade",
	},

	// --- data: forex ---
	"Rates": {
		parent:   "dataForex",
		use:      "rates",
		jsonOnly: true,
		examples: `  alpaca data forex rates --currency-pairs EUR/USD,GBP/USD --start 2025-01-01
  alpaca data forex rates --currency-pairs USD/JPY --timeframe 1Hour`,
	},
	"LatestRates": {
		parent:   "dataForex",
		use:      "latest",
		jsonOnly: true,
		examples: "  alpaca data forex latest --currency-pairs EUR/USD,GBP/USD",
	},

	// --- data: crypto orderbook ---
	"CryptoLatestOrderbooks": {
		parent:   "data",
		use:      "crypto-orderbook",
		jsonOnly: true,
		defaults: map[string]string{"loc": "us"},
		examples: "  alpaca data crypto-orderbook --symbols BTC/USD,ETH/USD",
	},

	// --- data: auctions ---
	"StockAuctions": {
		parent:   "data",
		use:      "auctions",
		jsonOnly: true,
		examples: `  alpaca data auctions --symbols AAPL --start 2025-01-01
  alpaca data auctions --symbols AAPL,MSFT --limit 10`,
	},
	"StockAuctionSingle": {
		parent:   "data",
		use:      "auction",
		jsonOnly: true,
		examples: "  alpaca data auction --symbol AAPL --start 2025-01-01",
	},

	// --- data: corporate actions (market data) ---
	"CorporateActions": {
		parent:   "data",
		use:      "corporate-actions",
		jsonOnly: true,
		examples: "  alpaca data corporate-actions --symbols AAPL --types forward_split --start 2025-01-01",
	},

	// --- data: fixed income ---
	"FixedIncomeLatestPrices": {
		parent:   "data",
		use:      "fixed-income",
		jsonOnly: true,
		examples: "  alpaca data fixed-income --isins 912797KR1,912797LB5",
	},

	// --- data: logo ---
	"Logos": {
		parent:   "data",
		use:      "logo",
		jsonOnly: true,
		examples: "  alpaca data logo --symbol AAPL",
	},

	// --- data: meta ---
	"StockMetaExchanges": {
		parent:   "dataMeta",
		use:      "exchanges",
		jsonOnly: true,
		examples: "  alpaca data meta exchanges",
	},
	"StockMetaConditions": {
		parent:   "dataMeta",
		use:      "conditions",
		jsonOnly: true,
		examples: "  alpaca data meta conditions --ticktype trade",
	},

	// --- data: screener ---
	"MostActives": {
		parent: "screener",
		use:    "most-actives",
		examples: `  alpaca data screener most-actives
  alpaca data screener most-actives --by trades --top 10`,
	},
	"Movers": {
		parent:   "screener",
		use:      "movers",
		defaults: map[string]string{"market_type": "stocks"},
		examples: `  alpaca data screener movers
  alpaca data screener movers --market-type crypto --top 5`,
	},

	// --- data: multi-symbol stock ---
	"StockBars": {
		parent:   "data",
		use:      "multi-bars",
		jsonOnly: true,
		examples: `  alpaca data multi-bars --symbols AAPL,MSFT --start 2025-01-01 --timeframe 1Day`,
	},
	"StockQuotes": {
		parent:   "data",
		use:      "multi-quotes",
		jsonOnly: true,
		examples: `  alpaca data multi-quotes --symbols AAPL,MSFT --start 2025-01-01`,
	},
	"StockTrades": {
		parent:   "data",
		use:      "multi-trades",
		jsonOnly: true,
		examples: `  alpaca data multi-trades --symbols AAPL,MSFT --start 2025-01-01`,
	},
	"StockSnapshots": {
		parent:   "data",
		use:      "multi-snapshots",
		jsonOnly: true,
		examples: `  alpaca data multi-snapshots --symbols AAPL,MSFT`,
	},
	"StockLatestBars": {
		parent:   "dataLatest",
		use:      "bars",
		jsonOnly: true,
		examples: `  alpaca data latest bars --symbols AAPL,MSFT`,
	},
	"StockLatestQuotes": {
		parent:   "dataLatest",
		use:      "quotes",
		jsonOnly: true,
		examples: `  alpaca data latest quotes --symbols AAPL,MSFT`,
	},
	"StockLatestTrades": {
		parent:   "dataLatest",
		use:      "trades",
		jsonOnly: true,
		examples: `  alpaca data latest trades --symbols AAPL,MSFT`,
	},

	// --- data: crypto ---
	"CryptoBars": {
		parent:   "dataCrypto",
		use:      "bars",
		jsonOnly: true,
		defaults: map[string]string{"loc": "us"},
		examples: `  alpaca data crypto bars --symbols BTC/USD --start 2025-01-01 --timeframe 1Day`,
	},
	"CryptoQuotes": {
		parent:   "dataCrypto",
		use:      "quotes",
		jsonOnly: true,
		defaults: map[string]string{"loc": "us"},
		examples: `  alpaca data crypto quotes --symbols BTC/USD --start 2025-01-01`,
	},
	"CryptoTrades": {
		parent:   "dataCrypto",
		use:      "trades",
		jsonOnly: true,
		defaults: map[string]string{"loc": "us"},
		examples: `  alpaca data crypto trades --symbols BTC/USD --start 2025-01-01`,
	},
	"CryptoSnapshots": {
		parent:   "dataCrypto",
		use:      "snapshots",
		jsonOnly: true,
		defaults: map[string]string{"loc": "us"},
		examples: `  alpaca data crypto snapshots --symbols BTC/USD`,
	},
	"CryptoLatestBars": {
		parent:   "dataCrypto",
		use:      "latest-bars",
		jsonOnly: true,
		defaults: map[string]string{"loc": "us"},
		examples: `  alpaca data crypto latest-bars --symbols BTC/USD`,
	},
	"CryptoLatestQuotes": {
		parent:   "dataCrypto",
		use:      "latest-quotes",
		jsonOnly: true,
		defaults: map[string]string{"loc": "us"},
		examples: `  alpaca data crypto latest-quotes --symbols BTC/USD`,
	},
	"CryptoLatestTrades": {
		parent:   "dataCrypto",
		use:      "latest-trades",
		jsonOnly: true,
		defaults: map[string]string{"loc": "us"},
		examples: `  alpaca data crypto latest-trades --symbols BTC/USD`,
	},

	// --- data: crypto perp data ---
	"CryptoPerpLatestBars": {
		parent:   "cryptoPerpData",
		use:      "latest-bars",
		jsonOnly: true,
		defaults: map[string]string{"loc": "us"},
		examples: `  alpaca crypto-perp data latest-bars --symbols BTC/USD`,
	},
	"CryptoPerpLatestFuturesPricing": {
		parent:   "cryptoPerpData",
		use:      "latest-futures-pricing",
		jsonOnly: true,
		defaults: map[string]string{"loc": "us"},
		examples: `  alpaca crypto-perp data latest-futures-pricing --symbols BTC/USD`,
	},
	"CryptoPerpLatestOrderbooks": {
		parent:   "cryptoPerpData",
		use:      "latest-orderbooks",
		jsonOnly: true,
		defaults: map[string]string{"loc": "us"},
		examples: `  alpaca crypto-perp data latest-orderbooks --symbols BTC/USD`,
	},
	"CryptoPerpLatestQuotes": {
		parent:   "cryptoPerpData",
		use:      "latest-quotes",
		jsonOnly: true,
		defaults: map[string]string{"loc": "us"},
		examples: `  alpaca crypto-perp data latest-quotes --symbols BTC/USD`,
	},
	"CryptoPerpLatestTrades": {
		parent:   "cryptoPerpData",
		use:      "latest-trades",
		jsonOnly: true,
		defaults: map[string]string{"loc": "us"},
		examples: `  alpaca crypto-perp data latest-trades --symbols BTC/USD`,
	},
}

var cmdSkip = map[string]string{
	"PostOrder": "hand-written: bracket orders, enums, dry-run, JSON parsing",
}

func checkExhaustive(allEndpoints []*endpointInfo, allSchemas []*schemaInfo) {
	epByOp := map[string]*endpointInfo{}
	for _, ep := range allEndpoints {
		epByOp[ep.goName] = ep
		_, inRegistry := cmdRegistry[ep.goName]
		_, inSkip := cmdSkip[ep.goName]
		if !inRegistry && !inSkip {
			log.Fatalf("unregistered operation %q (goName=%q) — add to cmdRegistry or cmdSkip in cmd/generate/commands.go", ep.operationID, ep.goName)
		}
	}

	for opID, def := range cmdRegistry {
		if def.examples == "" {
			log.Fatalf("cmdRegistry[%q] has empty examples — every generated command must have examples", opID)
		}
	}

	// Detect body field collisions with query/path params.
	// A collision means a body field is silently dropped by FlagDef dedup.
	// bodyAliases must resolve every collision explicitly.
	for opID, def := range cmdRegistry {
		ep := epByOp[opID]
		if ep == nil || ep.bodyRef == "" {
			continue
		}
		var bodySchema *schemaInfo
		for _, s := range allSchemas {
			if s.goName == ep.bodyRef && s.kind == "struct" {
				bodySchema = s
				break
			}
		}
		if bodySchema == nil {
			continue
		}
		nonBodyNames := map[string]bool{}
		for _, p := range ep.pathParams {
			nonBodyNames[strings.ReplaceAll(p.name, "_", "-")] = true
		}
		for _, p := range ep.queryParams {
			nonBodyNames[strings.ReplaceAll(p.name, "_", "-")] = true
		}
		for fieldName := range bodySchema.props {
			flagName := strings.ReplaceAll(fieldName, "_", "-")
			if !nonBodyNames[flagName] {
				continue
			}
			if _, aliased := def.bodyAliases[flagName]; aliased {
				continue
			}
			log.Fatalf("cmdRegistry[%q]: body field %q collides with a query/path param — add bodyAliases entry to resolve", opID, flagName)
		}
	}
}

func genCommands(allEndpoints []*endpointInfo, allSchemas []*schemaInfo) string {
	epByOp := map[string]*endpointInfo{}
	for _, ep := range allEndpoints {
		epByOp[ep.goName] = ep
	}

	// Emit command body first, then prepend header with conditional imports.
	var body bytes.Buffer

	// Emit parent command vars, sorted by key
	var parentKeys []string
	for k := range cmdParents {
		parentKeys = append(parentKeys, k)
	}
	sort.Strings(parentKeys)

	for _, key := range parentKeys {
		pdef := cmdParents[key]
		varName := key + "Cmd"
		fmt.Fprintf(&body, "var %s = &cobra.Command{\n", varName)
		fmt.Fprintf(&body, "\tUse:   %q,\n", pdef.use)
		fmt.Fprintf(&body, "\tShort: %q,\n", pdef.short)
		if len(pdef.aliases) > 0 {
			fmt.Fprintf(&body, "\tAliases: []string{")
			for i, a := range pdef.aliases {
				if i > 0 {
					fmt.Fprintf(&body, ", ")
				}
				fmt.Fprintf(&body, "%q", a)
			}
			fmt.Fprintf(&body, "},\n")
		}
		fmt.Fprintf(&body, "}\n\n")
	}

	// Emit command vars, sorted by operation ID
	var opIDs []string
	for opID := range cmdRegistry {
		opIDs = append(opIDs, opID)
	}
	sort.Strings(opIDs)

	for _, opID := range opIDs {
		def := cmdRegistry[opID]
		ep := epByOp[opID]
		if ep == nil {
			log.Fatalf("cmdRegistry references unknown operation %q", opID)
		}
		emitCommand(&body, opID, def, ep, allSchemas)
	}

	// Emit init() for subcommand wiring within groups
	emitInit(&body, opIDs, allEndpoints)

	// Assemble final output with header
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "// Code generated by cmd/generate; DO NOT EDIT.\n\n")
	fmt.Fprintf(&buf, "package cmd\n\n")
	fmt.Fprintf(&buf, "import (\n")
	bodyStr := body.String()
	if strings.Contains(bodyStr, "json.Unmarshal") {
		fmt.Fprintf(&buf, "\t\"encoding/json\"\n")
	}
	fmt.Fprintf(&buf, "\t\"fmt\"\n")
	if strings.Contains(bodyStr, "strings.") {
		fmt.Fprintf(&buf, "\t\"strings\"\n")
	}
	fmt.Fprintf(&buf, "\n")
	fmt.Fprintf(&buf, "\t\"github.com/alpacahq/cli/internal/api\"\n")
	fmt.Fprintf(&buf, "\t\"github.com/alpacahq/cli/internal/cmdutil\"\n")
	fmt.Fprintf(&buf, "\t\"github.com/spf13/cobra\"\n")
	fmt.Fprintf(&buf, ")\n\n")
	buf.Write(body.Bytes())

	return buf.String()
}

func emitCommand(buf *bytes.Buffer, opID string, def cmdDef, ep *endpointInfo, schemas []*schemaInfo) {
	opVar := opID + "Op"
	parentVar := def.parent + "Cmd"
	clientVar := "tradingClient"
	if ep.specSource == "marketdata" {
		clientVar = "dataClient"
	}

	// Build configure closures
	var configures []string

	if def.jsonOnly {
		configures = append(configures, "jsonOnly")
	}

	if len(def.defaults) > 0 {
		var pairs []string
		var defKeys []string
		for k := range def.defaults {
			defKeys = append(defKeys, k)
		}
		sort.Strings(defKeys)
		for _, k := range defKeys {
			pairs = append(pairs, fmt.Sprintf("%q: %q", k, def.defaults[k]))
		}
		configures = append(configures, fmt.Sprintf("flagOpts(&cmdutil.FlagOpts{Defaults: map[string]string{%s}})", strings.Join(pairs, ", ")))
	}

	// Register aliased body flags via configure closure
	if len(def.bodyAliases) > 0 {
		var bodySchema *schemaInfo
		for _, s := range schemas {
			if s.goName == ep.bodyRef && s.kind == "struct" {
				bodySchema = s
				break
			}
		}
		var aliasKeys []string
		for k := range def.bodyAliases {
			aliasKeys = append(aliasKeys, k)
		}
		sort.Strings(aliasKeys)
		var aliasLines []string
		for _, oasKebab := range aliasKeys {
			alias := def.bodyAliases[oasKebab]
			desc := ""
			if bodySchema != nil {
				oasName := strings.ReplaceAll(oasKebab, "-", "_")
				if ps, ok := bodySchema.props[oasName]; ok {
					desc, _ = ps["description"].(string)
				}
			}
			aliasLines = append(aliasLines, fmt.Sprintf("\t\tc.Flags().String(%q, \"\", %q)", alias, desc))
		}
		configures = append(configures, fmt.Sprintf("func(c *cobra.Command) {\n%s\n\t}", strings.Join(aliasLines, "\n")))
	}

	// Build configure func for long + examples
	hasConfigFunc := def.long != "" || def.examples != ""
	if hasConfigFunc {
		var lines []string
		if def.long != "" {
			lines = append(lines, fmt.Sprintf("\t\tc.Long = %q", def.long))
		}
		lines = append(lines, fmt.Sprintf("\t\tc.Example = %s", backtickQuote(def.examples)))
		configures = append(configures, fmt.Sprintf("func(c *cobra.Command) {\n%s\n\t}", strings.Join(lines, "\n")))
	}

	// Build the fetch callback
	fetchBody := buildFetchBody(opID, def, ep, clientVar, schemas)

	if def.self {
		// attachCmd on existing parent
		fmt.Fprintf(buf, "func init() {\n")
		fmt.Fprintf(buf, "\tattachCmd(%s, api.%s, func(cmd *cobra.Command, args []string) (any, error) {\n", parentVar, opVar)
		fmt.Fprintf(buf, "\t\t%s\n", fetchBody)
		fmt.Fprintf(buf, "\t}")
		for _, c := range configures {
			fmt.Fprintf(buf, ", %s", c)
		}
		fmt.Fprintf(buf, ")\n")
		fmt.Fprintf(buf, "}\n\n")
	} else {
		varName := cmdVarName(opID, def)
		fmt.Fprintf(buf, "var %s = fetchCmd(%q, api.%s, func(cmd *cobra.Command, args []string) (any, error) {\n", varName, def.use, opVar)
		fmt.Fprintf(buf, "\t%s\n", fetchBody)
		fmt.Fprintf(buf, "}")
		for _, c := range configures {
			fmt.Fprintf(buf, ", %s", c)
		}
		fmt.Fprintf(buf, ")\n\n")
	}
}

func buildFetchBody(opID string, def cmdDef, ep *endpointInfo, clientVar string, schemas []*schemaInfo) string {
	methodName := ep.goName
	isVoid := ep.responseEmpty || ep.responseRef == ""
	isPatch := ep.method == "PATCH" || ep.method == "PUT"
	isPost := ep.method == "POST"
	hasBodyRef := ep.bodyRef != ""
	hasBodyInline := len(ep.bodyInline) > 0
	hasQueryParams := len(ep.queryParams) > 0

	// rawMethod: override client method with (string, url.Values) signature
	if def.rawMethod != "" {
		symbolParam := ""
		for _, pp := range ep.pathParams {
			if strings.Contains(pp.name, "symbol") {
				symbolParam = pp.name
				break
			}
		}
		if symbolParam == "" && len(ep.pathParams) > 0 {
			symbolParam = ep.pathParams[0].name
		}
		symbolFlag := strings.ReplaceAll(symbolParam, "_", "-")
		paramsExpr := lcFirst(ep.goName) + "ParamsFromFlags(cmd).Values()"
		return fmt.Sprintf("return %s.%s(cmdutil.Str(cmd, %q), %s)", clientVar, def.rawMethod, symbolFlag, paramsExpr)
	}

	// For PATCH/PUT with body and bodyAliases → inline body with aliases
	if isPatch && hasBodyRef && len(def.bodyAliases) > 0 {
		return buildPatchBodyWithAliases(ep, def, clientVar, schemas)
	}

	// For PATCH with body
	if isPatch && hasBodyRef {
		bodyFuncName := lcFirst(ep.bodyRef) + "BodyFromFlags"
		var args []string
		for _, pp := range ep.pathParams {
			args = append(args, pathParamExpr(pp, def))
		}
		args = append(args, "body")

		lines := fmt.Sprintf("body, changed := %s(cmd)\n", bodyFuncName)
		for _, jf := range structRefBodyFields(ep.bodyRef, schemas) {
			lines += fmt.Sprintf("\tif cmdutil.Changed(cmd, %q) {\n", jf.flagName)
			lines += fmt.Sprintf("\t\tif err := json.Unmarshal([]byte(cmdutil.Str(cmd, %q)), &body.%s); err != nil {\n", jf.flagName, jf.goField)
			lines += fmt.Sprintf("\t\t\treturn nil, fmt.Errorf(\"--%s: %%w\", err)\n", jf.flagName)
			lines += "\t\t}\n"
			lines += "\t\tchanged = true\n"
			lines += "\t}\n"
		}
		lines += "\tif !changed {\n"
		lines += "\t\treturn nil, fmt.Errorf(\"specify at least one flag to change (see '%s --help')\", cmd.CommandPath())\n"
		lines += "\t}\n"
		lines += fmt.Sprintf("\treturn %s.%s(%s)", clientVar, methodName, strings.Join(args, ", "))
		return lines
	}

	// For POST with body ref (flat scalar)
	if isPost && hasBodyRef {
		bodyExpr := buildPostBody(ep, def, schemas)
		if bodyExpr != "" {
			var args []string
			for _, pp := range ep.pathParams {
				args = append(args, pathParamExpr(pp, def))
			}
			if hasQueryParams {
				args = append(args, lcFirst(ep.goName)+"ParamsFromFlags(cmd)")
			}
			args = append(args, bodyExpr)
			call := fmt.Sprintf("%s.%s(%s)", clientVar, methodName, strings.Join(args, ", "))
			if isVoid {
				return "return voidResponse(" + call + ")"
			}
			return "return " + call
		}
		// Try multi-statement block for mixed types (e.g. string + []string fields)
		if block := buildPostBodyBlock(ep, def, clientVar, schemas); block != "" {
			return block
		}
	}

	// For POST with inline body schema
	if isPost && hasBodyInline {
		bodyExpr := buildInlinePostBody(ep)
		if bodyExpr != "" {
			var args []string
			for _, pp := range ep.pathParams {
				args = append(args, pathParamExpr(pp, def))
			}
			if hasQueryParams {
				args = append(args, lcFirst(ep.goName)+"ParamsFromFlags(cmd)")
			}
			args = append(args, bodyExpr)
			call := fmt.Sprintf("%s.%s(%s)", clientVar, methodName, strings.Join(args, ", "))
			if isVoid {
				return "return voidResponse(" + call + ")"
			}
			return "return " + call
		}
	}

	// Standard: build argument list
	var args []string
	for _, pp := range ep.pathParams {
		args = append(args, pathParamExpr(pp, def))
	}
	if hasQueryParams {
		args = append(args, lcFirst(ep.goName)+"ParamsFromFlags(cmd)")
	}

	call := fmt.Sprintf("%s.%s(%s)", clientVar, methodName, strings.Join(args, ", "))
	if isVoid {
		return "return voidResponse(" + call + ")"
	}
	return "return " + call
}

func buildInlinePostBody(ep *endpointInfo) string {
	props, ok := ep.bodyInline["properties"].(map[string]any)
	if !ok || len(props) == 0 {
		return ""
	}

	var propNames []string
	for name := range props {
		propNames = append(propNames, name)
	}
	sort.Strings(propNames)

	// Only handle flat scalar (all string fields)
	for _, name := range propNames {
		ps, ok := props[name].(map[string]any)
		if !ok {
			return ""
		}
		t, _ := ps["type"].(string)
		if t != "string" {
			return ""
		}
	}

	typeName := ep.goName + "Request"
	var lines []string
	lines = append(lines, fmt.Sprintf("&api.%s{", typeName))
	for _, name := range propNames {
		flagName := strings.ReplaceAll(name, "_", "-")
		goField := toGoFieldName(name)
		lines = append(lines, fmt.Sprintf("\t\t%s: cmdutil.Str(cmd, %q),", goField, flagName))
	}
	lines = append(lines, "\t}")
	return strings.Join(lines, "\n")
}

func buildPostBody(ep *endpointInfo, def cmdDef, schemas []*schemaInfo) string {
	var bodySchema *schemaInfo
	for _, s := range schemas {
		if s.goName == ep.bodyRef && s.kind == "struct" {
			bodySchema = s
			break
		}
	}
	if bodySchema == nil {
		return ""
	}

	// Check all fields are strings (flat scalar)
	var propNames []string
	for name := range bodySchema.props {
		propNames = append(propNames, name)
	}
	sort.Strings(propNames)

	for _, name := range propNames {
		ps := bodySchema.props[name]
		if _, ok := ps["$ref"].(string); ok {
			return "" // has enum/ref — not flat scalar
		}
		t, _ := ps["type"].(string)
		if t != "string" {
			return "" // not flat scalar
		}
	}

	if len(propNames) == 0 {
		return ""
	}

	var lines []string
	lines = append(lines, fmt.Sprintf("&api.%s{", ep.bodyRef))
	for _, name := range propNames {
		flagName := strings.ReplaceAll(name, "_", "-")
		goField := toGoFieldName(name)
		lines = append(lines, fmt.Sprintf("\t\t%s: cmdutil.Str(cmd, %q),", goField, flagName))
	}
	lines = append(lines, "\t}")
	return strings.Join(lines, "\n")
}

func buildPatchBodyWithAliases(ep *endpointInfo, def cmdDef, clientVar string, schemas []*schemaInfo) string {
	var bodySchema *schemaInfo
	for _, s := range schemas {
		if s.goName == ep.bodyRef && s.kind == "struct" {
			bodySchema = s
			break
		}
	}
	if bodySchema == nil {
		return ""
	}

	var propNames []string
	for name := range bodySchema.props {
		propNames = append(propNames, name)
	}
	sort.Strings(propNames)

	var b strings.Builder
	fmt.Fprintf(&b, "body := &api.%s{}\n", ep.bodyRef)
	b.WriteString("\tvar changed bool\n")

	for _, name := range propNames {
		ps := bodySchema.props[name]
		flagName := strings.ReplaceAll(name, "_", "-")
		if alias, ok := def.bodyAliases[flagName]; ok {
			flagName = alias
		}
		goField := toGoFieldName(name)

		if _, ok := ps["$ref"]; ok {
			continue
		}
		switch t, _ := ps["type"].(string); t {
		case "string":
			fmt.Fprintf(&b, "\tif cmdutil.Changed(cmd, %q) {\n", flagName)
			fmt.Fprintf(&b, "\t\tbody.%s = cmdutil.Str(cmd, %q)\n", goField, flagName)
			b.WriteString("\t\tchanged = true\n")
			b.WriteString("\t}\n")
		case "array":
			if items, ok := ps["items"].(map[string]any); ok {
				if itemType, _ := items["type"].(string); itemType == "string" {
					fmt.Fprintf(&b, "\tif cmdutil.Changed(cmd, %q) {\n", flagName)
					fmt.Fprintf(&b, "\t\tif s := cmdutil.Str(cmd, %q); s != \"\" {\n", flagName)
					fmt.Fprintf(&b, "\t\t\tbody.%s = strings.Split(s, \",\")\n", goField)
					b.WriteString("\t\t}\n")
					b.WriteString("\t\tchanged = true\n")
					b.WriteString("\t}\n")
				}
			}
		}
	}

	b.WriteString("\tif !changed {\n")
	b.WriteString("\t\treturn nil, fmt.Errorf(\"specify at least one flag to change (see '%s --help')\", cmd.CommandPath())\n")
	b.WriteString("\t}\n")

	var args []string
	for _, pp := range ep.pathParams {
		args = append(args, pathParamExpr(pp, def))
	}
	if len(ep.queryParams) > 0 {
		args = append(args, lcFirst(ep.goName)+"ParamsFromFlags(cmd)")
	}
	args = append(args, "body")

	fmt.Fprintf(&b, "\treturn %s.%s(%s)", clientVar, ep.goName, strings.Join(args, ", "))
	return b.String()
}

func buildPostBodyBlock(ep *endpointInfo, def cmdDef, clientVar string, schemas []*schemaInfo) string {
	var bodySchema *schemaInfo
	for _, s := range schemas {
		if s.goName == ep.bodyRef && s.kind == "struct" {
			bodySchema = s
			break
		}
	}
	if bodySchema == nil {
		return ""
	}

	var propNames []string
	for name := range bodySchema.props {
		propNames = append(propNames, name)
	}
	sort.Strings(propNames)

	type fieldCat struct {
		goField  string
		flagName string
		kind     string // "string", "arrayString"
	}
	var fields []fieldCat
	var hasArray bool
	for _, name := range propNames {
		ps := bodySchema.props[name]
		if _, ok := ps["$ref"]; ok {
			return ""
		}
		flagName := strings.ReplaceAll(name, "_", "-")
		goField := toGoFieldName(name)
		switch t, _ := ps["type"].(string); t {
		case "string":
			fields = append(fields, fieldCat{goField, flagName, "string"})
		case "array":
			if items, ok := ps["items"].(map[string]any); ok {
				if itemType, _ := items["type"].(string); itemType == "string" {
					fields = append(fields, fieldCat{goField, flagName, "arrayString"})
					hasArray = true
					continue
				}
			}
			return ""
		default:
			return ""
		}
	}
	if !hasArray {
		return ""
	}

	var b strings.Builder
	fmt.Fprintf(&b, "body := &api.%s{\n", ep.bodyRef)
	for _, f := range fields {
		if f.kind == "string" {
			fmt.Fprintf(&b, "\t\t%s: cmdutil.Str(cmd, %q),\n", f.goField, f.flagName)
		}
	}
	b.WriteString("\t}\n")
	for _, f := range fields {
		if f.kind == "arrayString" {
			fmt.Fprintf(&b, "\tif s := cmdutil.Str(cmd, %q); s != \"\" {\n", f.flagName)
			fmt.Fprintf(&b, "\t\tbody.%s = strings.Split(s, \",\")\n", f.goField)
			b.WriteString("\t}\n")
		}
	}

	var args []string
	for _, pp := range ep.pathParams {
		args = append(args, pathParamExpr(pp, def))
	}
	if len(ep.queryParams) > 0 {
		args = append(args, lcFirst(ep.goName)+"ParamsFromFlags(cmd)")
	}
	args = append(args, "body")

	isVoid := ep.responseEmpty || ep.responseRef == ""
	call := fmt.Sprintf("%s.%s(%s)", clientVar, ep.goName, strings.Join(args, ", "))
	if isVoid {
		fmt.Fprintf(&b, "\treturn voidResponse(%s)", call)
	} else {
		fmt.Fprintf(&b, "\treturn %s", call)
	}
	return b.String()
}

type jsonField struct {
	flagName string
	goField  string
}

func structRefBodyFields(bodyRef string, schemas []*schemaInfo) []jsonField {
	var bodySchema *schemaInfo
	for _, s := range schemas {
		if s.goName == bodyRef && s.kind == "struct" {
			bodySchema = s
			break
		}
	}
	if bodySchema == nil {
		return nil
	}

	schemaByName := map[string]*schemaInfo{}
	for _, s := range schemas {
		schemaByName[s.name] = s
	}

	var propNames []string
	for name := range bodySchema.props {
		propNames = append(propNames, name)
	}
	sort.Strings(propNames)

	var fields []jsonField
	for _, name := range propNames {
		ps := bodySchema.props[name]
		ref, ok := ps["$ref"].(string)
		if !ok {
			continue
		}
		rn := refBaseName(ref)
		if s, ok := schemaByName[rn]; ok && s.kind == "struct" {
			fields = append(fields, jsonField{
				flagName: strings.ReplaceAll(name, "_", "-"),
				goField:  toGoFieldName(name),
			})
		}
	}
	return fields
}

func pathParamExpr(pp paramInfo, def cmdDef) string {
	flagName := strings.ReplaceAll(pp.name, "_", "-")
	expr := fmt.Sprintf("cmdutil.Str(cmd, %q)", flagName)
	for _, n := range def.normalize {
		if n == flagName {
			return "normalizePathParam(" + expr + ")"
		}
	}
	return expr
}

func cmdVarName(opID string, def cmdDef) string {
	candidate := lcFirst(opID) + "Cmd"
	// Avoid collisions with parent vars (e.g. Calendar op → calendarCmd would
	// collide with the "calendar" parent command var).
	parentVarName := def.parent + "Cmd"
	if candidate == parentVarName {
		candidate = lcFirst(opID) + ucFirst(def.use) + "Cmd"
	}
	return candidate
}

func ucFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func emitInit(buf *bytes.Buffer, opIDs []string, allEndpoints []*endpointInfo) {
	// Collect child→parent relationships for non-self commands
	type childEntry struct {
		varName string
		parent  string
	}
	var children []childEntry
	for _, opID := range opIDs {
		def := cmdRegistry[opID]
		if def.self {
			continue
		}
		children = append(children, childEntry{
			varName: cmdVarName(opID, def),
			parent:  def.parent,
		})
	}

	// Also wire parent→parent relationships (sub-groups)
	type parentWire struct {
		childVar  string
		parentVar string
	}
	var parentWires []parentWire
	var parentKeys []string
	for k := range cmdParents {
		parentKeys = append(parentKeys, k)
	}
	sort.Strings(parentKeys)

	for _, key := range parentKeys {
		pdef := cmdParents[key]
		if pdef.parent != "" {
			parentWires = append(parentWires, parentWire{
				childVar:  key + "Cmd",
				parentVar: pdef.parent + "Cmd",
			})
		}
	}

	fmt.Fprintf(buf, "func init() {\n")

	// Wire parent→parent first
	for _, pw := range parentWires {
		fmt.Fprintf(buf, "\t%s.AddCommand(%s)\n", pw.parentVar, pw.childVar)
	}
	if len(parentWires) > 0 && len(children) > 0 {
		fmt.Fprintf(buf, "\n")
	}

	// Wire commands to parents
	for _, c := range children {
		parentVar := c.parent + "Cmd"
		fmt.Fprintf(buf, "\t%s.AddCommand(%s)\n", parentVar, c.varName)
	}

	fmt.Fprintf(buf, "}\n")
}

func backtickQuote(s string) string {
	if strings.Contains(s, "`") {
		// Fall back to double-quoted string with escaping
		return fmt.Sprintf("%q", s)
	}
	return "`" + s + "`"
}
