package cmd

import (
	"fmt"

	"github.com/alpacahq/cli/internal/output"
)

func orderColumns() []output.Column {
	return []output.Column{
		{Header: "ID", Field: "id"},
		{Header: "SYMBOL", Field: "symbol"},
		{Header: "SIDE", Field: "side"},
		{Header: "QTY", Field: "qty"},
		{Header: "TYPE", Field: "type"},
		{Header: "STATUS", Field: "status"},
		{Header: "LIMIT PRICE", Field: "limit_price"},
		{Header: "STOP PRICE", Field: "stop_price"},
		{Header: "FILLED QTY", Field: "filled_qty"},
		{Header: "FILLED AVG", Field: "filled_avg_price"},
		{Header: "TIME IN FORCE", Field: "time_in_force"},
		{Header: "SUBMITTED", Field: "submitted_at"},
	}
}

func positionColumns() []output.Column {
	return []output.Column{
		{Header: "SYMBOL", Field: "symbol"},
		{Header: "QTY", Field: "qty"},
		{Header: "SIDE", Field: "side"},
		{Header: "AVG ENTRY", Field: "avg_entry_price"},
		{Header: "CURRENT", Field: "current_price"},
		{Header: "MKT VALUE", Field: "market_value"},
		{Header: "P/L", Field: "unrealized_pl", Format: func(v any) string {
			return output.DollarPL(fmt.Sprintf("%v", v))
		}},
		{Header: "P/L %", Field: "unrealized_plpc", Format: func(v any) string {
			return output.PercentPL(fmt.Sprintf("%v", v))
		}},
	}
}

func accountColumns() []output.Column {
	return []output.Column{
		{Header: "ACCOUNT #", Field: "account_number"},
		{Header: "STATUS", Field: "status"},
		{Header: "EQUITY", Field: "equity"},
		{Header: "CASH", Field: "cash"},
		{Header: "BUYING POWER", Field: "buying_power"},
		{Header: "PORTFOLIO VALUE", Field: "portfolio_value"},
		{Header: "CURRENCY", Field: "currency"},
		{Header: "PDT", Field: "pattern_day_trader"},
		{Header: "TRADING BLOCKED", Field: "trading_blocked"},
	}
}

func accountConfigColumns() []output.Column {
	return []output.Column{
		{Header: "DTBP CHECK", Field: "dtbp_check"},
		{Header: "FRACTIONAL TRADING", Field: "fractional_trading"},
		{Header: "MAX MARGIN MULTIPLIER", Field: "max_margin_multiplier"},
		{Header: "NO SHORTING", Field: "no_shorting"},
		{Header: "PDT CHECK", Field: "pdt_check"},
		{Header: "SUSPEND TRADE", Field: "suspend_trade"},
		{Header: "TRADE CONFIRM EMAIL", Field: "trade_confirm_email"},
	}
}

func assetListColumns() []output.Column {
	return []output.Column{
		{Header: "SYMBOL", Field: "symbol"},
		{Header: "NAME", Field: "name"},
		{Header: "CLASS", Field: "class"},
		{Header: "EXCHANGE", Field: "exchange"},
		{Header: "STATUS", Field: "status"},
		{Header: "TRADABLE", Field: "tradable"},
		{Header: "SHORTABLE", Field: "shortable"},
		{Header: "FRACTIONABLE", Field: "fractionable"},
	}
}

func assetDetailColumns() []output.Column {
	return []output.Column{
		{Header: "SYMBOL", Field: "symbol"},
		{Header: "NAME", Field: "name"},
		{Header: "CLASS", Field: "class"},
		{Header: "EXCHANGE", Field: "exchange"},
		{Header: "STATUS", Field: "status"},
		{Header: "TRADABLE", Field: "tradable"},
		{Header: "SHORTABLE", Field: "shortable"},
		{Header: "FRACTIONABLE", Field: "fractionable"},
		{Header: "MARGINABLE", Field: "marginable"},
		{Header: "EASY TO BORROW", Field: "easy_to_borrow"},
	}
}

func watchlistColumns() []output.Column {
	return []output.Column{
		{Header: "ID", Field: "id"},
		{Header: "NAME", Field: "name"},
		{Header: "CREATED", Field: "created_at"},
		{Header: "UPDATED", Field: "updated_at"},
	}
}

func optionChainColumns() []output.Column {
	return []output.Column{
		{Header: "SYMBOL", Field: "symbol"},
		{Header: "TYPE", Field: "type"},
		{Header: "STRIKE", Field: "strike_price"},
		{Header: "EXPIRY", Field: "expiration_date"},
		{Header: "STATUS", Field: "status"},
		{Header: "UNDERLYING", Field: "underlying_symbol"},
	}
}

func calendarColumns() []output.Column {
	return []output.Column{
		{Header: "DATE", Field: "date"},
		{Header: "OPEN", Field: "open"},
		{Header: "CLOSE", Field: "close"},
		{Header: "SESSION OPEN", Field: "session_open"},
		{Header: "SESSION CLOSE", Field: "session_close"},
	}
}

func newsColumns() []output.Column {
	return []output.Column{
		{Header: "DATE", Field: "created_at"},
		{Header: "HEADLINE", Field: "headline"},
		{Header: "SOURCE", Field: "source"},
		{Header: "SYMBOLS", Field: "symbols"},
		{Header: "URL", Field: "url"},
	}
}

func barColumns() []output.Column {
	return []output.Column{
		{Header: "TIMESTAMP", Field: "t"},
		{Header: "OPEN", Field: "o"},
		{Header: "HIGH", Field: "h"},
		{Header: "LOW", Field: "l"},
		{Header: "CLOSE", Field: "c"},
		{Header: "VOLUME", Field: "v"},
		{Header: "VWAP", Field: "vw"},
	}
}

func quoteColumns() []output.Column {
	return []output.Column{
		{Header: "TIMESTAMP", Field: "t"},
		{Header: "BID", Field: "bp"},
		{Header: "BID SIZE", Field: "bs"},
		{Header: "ASK", Field: "ap"},
		{Header: "ASK SIZE", Field: "as"},
	}
}

func tradeColumns() []output.Column {
	return []output.Column{
		{Header: "TIMESTAMP", Field: "t"},
		{Header: "PRICE", Field: "p"},
		{Header: "SIZE", Field: "s"},
		{Header: "EXCHANGE", Field: "x"},
	}
}

func tradeActivityColumns() []output.Column {
	return []output.Column{
		{Header: "ID", Field: "id"},
		{Header: "TYPE", Field: "activity_type"},
		{Header: "SYMBOL", Field: "symbol"},
		{Header: "SIDE", Field: "side"},
		{Header: "QTY", Field: "qty"},
		{Header: "PRICE", Field: "price"},
		{Header: "CUM QTY", Field: "cum_qty"},
		{Header: "ORDER ID", Field: "order_id"},
		{Header: "TIME", Field: "transaction_time"},
	}
}

func nonTradeActivityColumns() []output.Column {
	return []output.Column{
		{Header: "ID", Field: "id"},
		{Header: "TYPE", Field: "activity_type"},
		{Header: "SYMBOL", Field: "symbol"},
		{Header: "NET AMOUNT", Field: "net_amount"},
		{Header: "QTY", Field: "qty"},
		{Header: "PER SHARE", Field: "per_share_amount"},
		{Header: "STATUS", Field: "status"},
		{Header: "DATE", Field: "date"},
	}
}

func screenerMostActivesColumns() []output.Column {
	return []output.Column{
		{Header: "SYMBOL", Field: "symbol"},
		{Header: "VOLUME", Field: "volume"},
		{Header: "TRADE COUNT", Field: "trade_count"},
	}
}

func screenerMoverColumns() []output.Column {
	return []output.Column{
		{Header: "SYMBOL", Field: "symbol"},
		{Header: "CHANGE %", Field: "percent_change"},
		{Header: "CHANGE", Field: "change"},
		{Header: "PRICE", Field: "price"},
	}
}

func forexRateColumns() []output.Column {
	return []output.Column{
		{Header: "TIMESTAMP", Field: "t"},
		{Header: "BID", Field: "bp"},
		{Header: "ASK", Field: "ap"},
		{Header: "MID", Field: "mp"},
	}
}

func treasuryColumns() []output.Column {
	return []output.Column{
		{Header: "CUSIP", Field: "cusip"},
		{Header: "DESCRIPTION", Field: "description_short"},
		{Header: "SUBTYPE", Field: "subtype"},
		{Header: "STATUS", Field: "bond_status"},
		{Header: "COUPON", Field: "coupon"},
		{Header: "MATURITY", Field: "maturity_date"},
		{Header: "CLOSE PRICE", Field: "close_price"},
		{Header: "CLOSE YTM", Field: "close_yield_to_maturity"},
	}
}

func bondColumns() []output.Column {
	return []output.Column{
		{Header: "CUSIP", Field: "cusip"},
		{Header: "TICKER", Field: "ticker"},
		{Header: "DESCRIPTION", Field: "description_short"},
		{Header: "STATUS", Field: "bond_status"},
		{Header: "COUPON", Field: "coupon"},
		{Header: "MATURITY", Field: "maturity_date"},
		{Header: "CLOSE PRICE", Field: "close_price"},
	}
}

func walletColumns() []output.Column {
	return []output.Column{
		{Header: "ADDRESS", Field: "address"},
		{Header: "CHAIN", Field: "chain"},
		{Header: "CREATED", Field: "created_at"},
	}
}

func transferColumns() []output.Column {
	return []output.Column{
		{Header: "ID", Field: "id"},
		{Header: "ASSET", Field: "asset"},
		{Header: "AMOUNT", Field: "amount"},
		{Header: "DIRECTION", Field: "direction"},
		{Header: "STATUS", Field: "status"},
		{Header: "CHAIN", Field: "chain"},
		{Header: "CREATED", Field: "created_at"},
	}
}

func whitelistColumns() []output.Column {
	return []output.Column{
		{Header: "ID", Field: "id"},
		{Header: "ADDRESS", Field: "address"},
		{Header: "ASSET", Field: "asset"},
		{Header: "CHAIN", Field: "chain"},
		{Header: "STATUS", Field: "status"},
		{Header: "CREATED", Field: "created_at"},
	}
}

func corporateActionColumns() []output.Column {
	return []output.Column{
		{Header: "ID", Field: "id"},
		{Header: "TYPE", Field: "ca_type"},
		{Header: "SUB TYPE", Field: "ca_sub_type"},
		{Header: "SYMBOL", Field: "symbol"},
		{Header: "EX DATE", Field: "ex_date"},
		{Header: "RECORD DATE", Field: "record_date"},
		{Header: "PAYABLE DATE", Field: "payable_date"},
	}
}
