# Alpaca CLI Command Reference

Quick reference for common tasks. All examples use `--json --quiet` for machine-readable output.

## Trading

### Submit orders

```bash
# Market order
alpaca order submit AAPL --side buy --qty 10 --type market --json --quiet

# Limit order
alpaca order submit AAPL --side buy --qty 10 --type limit --limit-price 185.00 --json --quiet

# Stop-limit order
alpaca order submit AAPL --side sell --qty 5 --type stop_limit \
  --stop-price 180.00 --limit-price 179.50 --json --quiet

# Trailing stop (percent)
alpaca order submit AAPL --side sell --qty 5 --type trailing_stop \
  --trail-percent 2.5 --json --quiet

# Preview without submitting
alpaca order submit AAPL --side buy --qty 10 --type market --dry-run
```

### Manage orders

```bash
alpaca order list --json --quiet                       # all orders
alpaca order list --status open --json --quiet          # open orders only
alpaca order get <order-id> --json --quiet              # order details
alpaca order cancel <order-id> --json --quiet           # cancel one
alpaca order cancel-all --json --quiet                  # cancel all open
alpaca order replace <order-id> --qty 20 --json --quiet # modify order
```

### Positions

```bash
alpaca position list --json --quiet                 # all open positions
alpaca position get AAPL --json --quiet              # single position
alpaca position close AAPL --json --quiet            # close position
alpaca position close-all --json --quiet             # close everything
```

### Options

```bash
alpaca option chain AAPL --json --quiet                    # options chain
alpaca option get <contract-id> --json --quiet             # contract details
alpaca option exercise <contract-id> --json --quiet        # exercise
alpaca option do-not-exercise <contract-id> --json --quiet # mark do-not-exercise
```

## Market Data

### Price data

```bash
# Historical bars
alpaca data bars AAPL --start 2025-01-01 --timeframe 1Day --json --quiet

# Latest trade/quote/bar
alpaca data latest trade AAPL --json --quiet
alpaca data latest quote AAPL --json --quiet
alpaca data latest bar AAPL --json --quiet

# Full snapshot
alpaca data snapshot AAPL --json --quiet
```

### Options data

```bash
alpaca data option chain AAPL --json --quiet          # greeks and pricing
alpaca data option bars --json --quiet                # option historical bars
alpaca data option trades --json --quiet              # option historical trades
alpaca data option snapshot --json --quiet            # option snapshots
alpaca data option latest-quotes --json --quiet
alpaca data option latest-trades --json --quiet
alpaca data option exchanges --json --quiet           # option exchanges
alpaca data option conditions --json --quiet          # option trade conditions
```

### Additional data

```bash
alpaca data forex rates --json --quiet               # historical forex rates
alpaca data forex latest --json --quiet              # latest forex rates
alpaca data crypto-orderbook --json --quiet          # crypto orderbooks
alpaca data auctions --json --quiet                  # stock auction data
alpaca data corporate-actions --json --quiet         # corporate actions (market data)
alpaca data fixed-income --json --quiet              # fixed income prices
alpaca data logo AAPL --json --quiet                 # company logo URL
alpaca data meta exchanges --json --quiet            # exchange code reference
alpaca data meta conditions trade --json --quiet     # trade condition codes
```

### Screeners and news

```bash
alpaca screener most-actives --json --quiet
alpaca screener movers --json --quiet
alpaca news --json --quiet
```

## Account

```bash
alpaca account get --json --quiet                    # equity, buying power
alpaca account config get --json --quiet             # account settings
alpaca activity list --json --quiet                   # fills, dividends, etc.
alpaca portfolio history --json --quiet               # portfolio value over time
```

## Market status

```bash
alpaca clock --json --quiet                          # is market open?
alpaca calendar --json --quiet                       # trading calendar
```

## Assets and watchlists

```bash
alpaca asset list --json --quiet                     # browse equities/crypto
alpaca asset get AAPL --json --quiet                 # asset details
alpaca asset treasury --json --quiet                 # US Treasury bonds
alpaca asset bond --json --quiet                     # US Corporate bonds

# Watchlists (by ID)
alpaca watchlist list --json --quiet
alpaca watchlist create "My List" --symbols AAPL,MSFT --json --quiet
alpaca watchlist get <id> --json --quiet
alpaca watchlist add <id> GOOG --json --quiet
alpaca watchlist remove <id> MSFT --json --quiet
alpaca watchlist delete <id>

# Watchlists (by name)
alpaca watchlist get-by-name "My List" --json --quiet
alpaca watchlist add-by-name "My List" NVDA --json --quiet
alpaca watchlist update-by-name "My List" --symbols AAPL,GOOG --json --quiet
alpaca watchlist delete-by-name "My List"

# Wallet / crypto transfers
alpaca wallet transfer estimate --amount 0.1 --asset BTC --address 0x... --json --quiet
```

## Raw API access

For endpoints not covered by named commands:

```bash
alpaca api get /v2/account --json --quiet
alpaca api post /v2/orders --data '{"symbol":"AAPL","qty":"1","side":"buy","type":"market","time_in_force":"day"}' --json --quiet
alpaca api patch /v2/account/configurations --data '{"dtbp_check":"both"}' --json --quiet
alpaca api delete /v2/orders/<id> --json --quiet
```

Pipe payloads from stdin:

```bash
cat order.json | alpaca api post /v2/orders --json --quiet
```

## Agent workflow example

```bash
export ALPACA_API_KEY=PK...
export ALPACA_SECRET_KEY=...

# Check market status
clock=$(alpaca clock --json --quiet)
is_open=$(echo "$clock" | jq -r '.is_open')

# Trade only if market is open
if [ "$is_open" = "true" ]; then
  alpaca order submit AAPL --side buy --qty 10 --type market --json --quiet
fi

# Handle errors
if ! result=$(alpaca order get <id> --json --quiet 2>err.json); then
  echo "Failed: $(jq -r .error err.json)"
fi
```
