package cmd

import (
	"fmt"

	"github.com/alpacahq/cli/internal/output"
)

// col creates a column with an explicit header.
// Timestamp formatting is auto-detected from the field name.
func col(field, header string) output.Column {
	c := output.Column{Header: header, Field: field}
	if output.IsTimestampField(field) {
		c.Format = output.FormatTimestamp
	}
	return c
}

// colFmt creates a column with an explicit header and custom format function.
func colFmt(field, header string, format func(any) string) output.Column {
	return output.Column{Header: header, Field: field, Format: format}
}

// orderColumns subsets the ~30-field order response to the most useful columns.
func orderColumns() []output.Column {
	return []output.Column{
		col("id", "ID"),
		col("symbol", "SYMBOL"),
		col("order_class", "CLASS"),
		col("side", "SIDE"),
		col("qty", "QTY"),
		col("type", "TYPE"),
		col("status", "STATUS"),
		col("limit_price", "LIMIT PRICE"),
		col("stop_price", "STOP PRICE"),
		col("filled_qty", "FILLED QTY"),
		col("filled_avg_price", "FILLED AVG"),
		col("time_in_force", "TIME IN FORCE"),
		col("submitted_at", "SUBMITTED"),
	}
}

// canceledOrderColumns shows only id + status for cancel-all results.
func canceledOrderColumns() []output.Column {
	return []output.Column{
		col("id", "ID"),
		col("status", "STATUS"),
	}
}

// positionColumns includes P/L coloring.
func positionColumns() []output.Column {
	return []output.Column{
		col("symbol", "SYMBOL"),
		col("qty", "QTY"),
		col("side", "SIDE"),
		col("avg_entry_price", "AVG ENTRY"),
		col("current_price", "CURRENT"),
		col("market_value", "MKT VALUE"),
		colFmt("unrealized_pl", "P/L", func(v any) string {
			return output.DollarPL(fmt.Sprintf("%v", v))
		}),
		colFmt("unrealized_plpc", "P/L %", func(v any) string {
			return output.PercentPL(fmt.Sprintf("%v", v))
		}),
	}
}

// barColumns maps single-letter API fields to human-readable headers.
func barColumns() []output.Column {
	return []output.Column{
		col("t", "TIMESTAMP"),
		col("o", "OPEN"),
		col("h", "HIGH"),
		col("l", "LOW"),
		col("c", "CLOSE"),
		col("v", "VOLUME"),
		col("vw", "VWAP"),
	}
}

// quoteColumns maps abbreviated API fields to human-readable headers.
func quoteColumns() []output.Column {
	return []output.Column{
		col("t", "TIMESTAMP"),
		col("bp", "BID"),
		col("bs", "BID SIZE"),
		col("ap", "ASK"),
		col("as", "ASK SIZE"),
	}
}

// tradeColumns maps single-letter API fields to human-readable headers.
func tradeColumns() []output.Column {
	return []output.Column{
		col("t", "TIMESTAMP"),
		col("p", "PRICE"),
		col("s", "SIZE"),
		col("x", "EXCHANGE"),
	}
}

// forexRateColumns maps abbreviated API fields to human-readable headers.
func forexRateColumns() []output.Column {
	return []output.Column{
		col("t", "TIMESTAMP"),
		col("bp", "BID"),
		col("ap", "ASK"),
		col("mp", "MID"),
	}
}
