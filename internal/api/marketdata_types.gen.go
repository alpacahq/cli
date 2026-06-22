// Code generated from api/specs/market-data-api.json; DO NOT EDIT.

package api

type CryptoHistoricalLoc string

type CryptoLatestLoc string

type CryptoPerpLoc string

type MarketType string

type OptionFeed string

type Region string

type Sort string

type StockHistoricalFeed string

type StockLatestFeed string

type StockTape string

type CashDividend struct {
	Currency       string  `json:"currency,omitempty"`
	Cusip          string  `json:"cusip"`
	DueBillOffDate string  `json:"due_bill_off_date,omitempty"`
	DueBillOnDate  string  `json:"due_bill_on_date,omitempty"`
	ExDate         string  `json:"ex_date"`
	Foreign        bool    `json:"foreign"`
	ID             string  `json:"id"`
	Isin           string  `json:"isin,omitempty"`
	PayableDate    string  `json:"payable_date,omitempty"`
	ProcessDate    string  `json:"process_date"`
	Rate           float64 `json:"rate"`
	RecordDate     string  `json:"record_date,omitempty"`
	Special        bool    `json:"special"`
	SubType        string  `json:"sub_type,omitempty"`
	Symbol         string  `json:"symbol"`
}

type CashMerger struct {
	AcquireeCusip  string  `json:"acquiree_cusip"`
	AcquireeIsin   string  `json:"acquiree_isin,omitempty"`
	AcquireeSymbol string  `json:"acquiree_symbol"`
	AcquirerCusip  string  `json:"acquirer_cusip,omitempty"`
	AcquirerIsin   string  `json:"acquirer_isin,omitempty"`
	AcquirerSymbol string  `json:"acquirer_symbol,omitempty"`
	Currency       string  `json:"currency,omitempty"`
	EffectiveDate  string  `json:"effective_date"`
	ID             string  `json:"id"`
	PayableDate    string  `json:"payable_date,omitempty"`
	ProcessDate    string  `json:"process_date"`
	Rate           float64 `json:"rate"`
}

type CorporateActions struct {
	CashDividends       []CashDividend       `json:"cash_dividends,omitempty"`
	CashMergers         []CashMerger         `json:"cash_mergers,omitempty"`
	ForwardSplits       []ForwardSplit       `json:"forward_splits,omitempty"`
	NameChanges         []NameChange         `json:"name_changes,omitempty"`
	PartialCalls        []PartialCall        `json:"partial_calls,omitempty"`
	Redemptions         []Redemption         `json:"redemptions,omitempty"`
	Reorganizations     []Reorganization     `json:"reorganizations,omitempty"`
	ReverseSplits       []ReverseSplit       `json:"reverse_splits,omitempty"`
	RightsDistributions []RightsDistribution `json:"rights_distributions,omitempty"`
	SpinOffs            []SpinOff            `json:"spin_offs,omitempty"`
	StockAndCashMergers []StockAndCashMerger `json:"stock_and_cash_mergers,omitempty"`
	StockDividends      []StockDividend      `json:"stock_dividends,omitempty"`
	StockMergers        []StockMerger        `json:"stock_mergers,omitempty"`
	UnitSplits          []UnitSplit          `json:"unit_splits,omitempty"`
	WorthlessRemovals   []WorthlessRemoval   `json:"worthless_removals,omitempty"`
}

type CorporateActionsResp struct {
	CorporateActions CorporateActions `json:"corporate_actions"`
	NextPageToken    string           `json:"next_page_token"`
}

type CryptoBar struct {
	C  float64 `json:"c"`
	H  float64 `json:"h"`
	L  float64 `json:"l"`
	N  int     `json:"n"`
	O  float64 `json:"o"`
	T  string  `json:"t"`
	V  float64 `json:"v"`
	Vw float64 `json:"vw"`
}

type CryptoBarsResp struct {
	Bars          map[string][]CryptoBar `json:"bars"`
	NextPageToken string                 `json:"next_page_token"`
}

type CryptoLatestBarsResp struct {
	Bars map[string]CryptoBar `json:"bars"`
}

type CryptoLatestOrderbooksResp struct {
	Orderbooks map[string]CryptoOrderbook `json:"orderbooks"`
}

type CryptoLatestQuotesResp struct {
	Quotes map[string]CryptoQuote `json:"quotes"`
}

type CryptoLatestTradesResp struct {
	Trades map[string]CryptoTrade `json:"trades"`
}

type CryptoOrderbook struct {
	A []CryptoOrderbookEntry `json:"a"`
	B []CryptoOrderbookEntry `json:"b"`
	T string                 `json:"t"`
}

type CryptoOrderbookEntry struct {
	P float64 `json:"p"`
	S float64 `json:"s"`
}

type CryptoPerpFuturesPricing struct {
	Fr float64 `json:"fr"`
	Ft string  `json:"ft"`
	IP float64 `json:"ip"`
	Mp float64 `json:"mp"`
	Oi float64 `json:"oi"`
	T  string  `json:"t"`
}

type CryptoPerpLatestFuturesPricingResp struct {
	Pricing map[string]CryptoPerpFuturesPricing `json:"pricing"`
}

type CryptoQuote struct {
	Ap float64 `json:"ap"`
	As float64 `json:"as"`
	Bp float64 `json:"bp"`
	Bs float64 `json:"bs"`
	T  string  `json:"t"`
}

type CryptoQuotesResp struct {
	NextPageToken string                   `json:"next_page_token"`
	Quotes        map[string][]CryptoQuote `json:"quotes"`
}

type CryptoSnapshot struct {
	DailyBar     CryptoBar   `json:"dailyBar,omitempty"`
	LatestQuote  CryptoQuote `json:"latestQuote,omitempty"`
	LatestTrade  CryptoTrade `json:"latestTrade,omitempty"`
	MinuteBar    CryptoBar   `json:"minuteBar,omitempty"`
	PrevDailyBar CryptoBar   `json:"prevDailyBar,omitempty"`
}

type CryptoSnapshotsResp struct {
	Snapshots map[string]CryptoSnapshot `json:"snapshots"`
}

type CryptoTrade struct {
	I   int     `json:"i"`
	P   float64 `json:"p"`
	S   float64 `json:"s"`
	T   string  `json:"t"`
	Tks string  `json:"tks"`
}

type CryptoTradesResp struct {
	NextPageToken string                   `json:"next_page_token"`
	Trades        map[string][]CryptoTrade `json:"trades"`
}

type FixedIncomeLatestPricesResp struct {
	Prices map[string]FixedIncomePrice `json:"prices"`
}

type FixedIncomeLatestQuotesResp struct {
	Quotes map[string]FixedIncomeQuote `json:"quotes"`
}

type FixedIncomePrice struct {
	P   float64 `json:"p"`
	T   string  `json:"t"`
	Ytm float64 `json:"ytm,omitempty"`
	Ytw float64 `json:"ytw,omitempty"`
}

type FixedIncomeQuote struct {
	Ams  int     `json:"ams"`
	Ap   float64 `json:"ap"`
	As   int     `json:"as"`
	Aytm float64 `json:"aytm"`
	Aytw float64 `json:"aytw"`
	Bms  int     `json:"bms"`
	Bp   float64 `json:"bp"`
	Bs   int     `json:"bs"`
	Bytm float64 `json:"bytm"`
	Bytw float64 `json:"bytw"`
	T    string  `json:"t"`
}

type ForexLatestRatesResp struct {
	Rates map[string]ForexRate `json:"rates"`
}

type ForexRate struct {
	Ap float64 `json:"ap"`
	Bp float64 `json:"bp"`
	Mp float64 `json:"mp"`
	T  string  `json:"t"`
}

type ForexRatesResp struct {
	NextPageToken string                 `json:"next_page_token"`
	Rates         map[string][]ForexRate `json:"rates"`
}

type ForwardSplit struct {
	Currency              string  `json:"currency,omitempty"`
	Cusip                 string  `json:"cusip"`
	DueBillRedemptionDate string  `json:"due_bill_redemption_date,omitempty"`
	ExDate                string  `json:"ex_date"`
	ID                    string  `json:"id"`
	Isin                  string  `json:"isin,omitempty"`
	NewRate               float64 `json:"new_rate"`
	OldRate               float64 `json:"old_rate"`
	PayableDate           string  `json:"payable_date,omitempty"`
	ProcessDate           string  `json:"process_date"`
	RecordDate            string  `json:"record_date,omitempty"`
	Symbol                string  `json:"symbol"`
}

type IndexLatestValuesResp struct {
	Values map[string]IndexValue `json:"values"`
}

type IndexValue struct {
	T string  `json:"t"`
	V float64 `json:"v"`
}

type IndexValuesResp struct {
	NextPageToken string                  `json:"next_page_token"`
	Values        map[string][]IndexValue `json:"values"`
}

type MostActive struct {
	Symbol     string `json:"symbol"`
	TradeCount int    `json:"trade_count"`
	Volume     int    `json:"volume"`
}

type MostActivesResp struct {
	LastUpdated string       `json:"last_updated"`
	MostActives []MostActive `json:"most_actives"`
}

type Mover struct {
	Change        float64 `json:"change"`
	PercentChange float64 `json:"percent_change"`
	Price         float64 `json:"price"`
	Symbol        string  `json:"symbol"`
}

type MoversResp struct {
	Gainers     []Mover    `json:"gainers"`
	LastUpdated string     `json:"last_updated"`
	Losers      []Mover    `json:"losers"`
	MarketType  MarketType `json:"market_type"`
}

type NameChange struct {
	Currency    string `json:"currency,omitempty"`
	ID          string `json:"id"`
	NewCusip    string `json:"new_cusip"`
	NewIsin     string `json:"new_isin,omitempty"`
	NewSymbol   string `json:"new_symbol"`
	OldCusip    string `json:"old_cusip"`
	OldIsin     string `json:"old_isin,omitempty"`
	OldSymbol   string `json:"old_symbol"`
	ProcessDate string `json:"process_date"`
}

type News struct {
	Author    string      `json:"author"`
	Content   string      `json:"content"`
	CreatedAt string      `json:"created_at"`
	Headline  string      `json:"headline"`
	ID        int         `json:"id"`
	Images    []NewsImage `json:"images"`
	Source    string      `json:"source"`
	Summary   string      `json:"summary"`
	Symbols   []string    `json:"symbols"`
	UpdatedAt string      `json:"updated_at"`
	URL       *string     `json:"url,omitempty"`
}

type NewsImage struct {
	Size string `json:"size"`
	URL  string `json:"url"`
}

type NewsResp struct {
	News          []News `json:"news"`
	NextPageToken string `json:"next_page_token"`
}

type OptionBar struct {
	C  float64 `json:"c"`
	H  float64 `json:"h"`
	L  float64 `json:"l"`
	N  int     `json:"n"`
	O  float64 `json:"o"`
	T  string  `json:"t"`
	V  int     `json:"v"`
	Vw float64 `json:"vw"`
}

type OptionBarsResp struct {
	Bars          map[string][]OptionBar `json:"bars"`
	Currency      string                 `json:"currency,omitempty"`
	NextPageToken string                 `json:"next_page_token"`
}

type OptionGreeks struct {
	Delta float64 `json:"delta"`
	Gamma float64 `json:"gamma"`
	Rho   float64 `json:"rho"`
	Theta float64 `json:"theta"`
	Vega  float64 `json:"vega"`
}

type OptionLatestQuotesResp struct {
	Quotes map[string]OptionQuote `json:"quotes"`
}

type OptionLatestTradesResp struct {
	Trades map[string]OptionTrade `json:"trades"`
}

type OptionQuote struct {
	Ap float64 `json:"ap"`
	As int     `json:"as"`
	Ax string  `json:"ax"`
	Bp float64 `json:"bp"`
	Bs int     `json:"bs"`
	Bx string  `json:"bx"`
	C  string  `json:"c"`
	T  string  `json:"t"`
}

type OptionSnapshot struct {
	DailyBar          OptionBar    `json:"dailyBar,omitempty"`
	Greeks            OptionGreeks `json:"greeks,omitempty"`
	ImpliedVolatility float64      `json:"impliedVolatility,omitempty"`
	LatestQuote       OptionQuote  `json:"latestQuote,omitempty"`
	LatestTrade       OptionTrade  `json:"latestTrade,omitempty"`
	MinuteBar         OptionBar    `json:"minuteBar,omitempty"`
	PrevDailyBar      OptionBar    `json:"prevDailyBar,omitempty"`
}

type OptionSnapshotsResp struct {
	NextPageToken string                    `json:"next_page_token"`
	Snapshots     map[string]OptionSnapshot `json:"snapshots"`
}

type OptionTrade struct {
	C string  `json:"c"`
	P float64 `json:"p"`
	S int     `json:"s"`
	T string  `json:"t"`
	X string  `json:"x"`
}

type OptionTradesResp struct {
	Currency      string                   `json:"currency,omitempty"`
	NextPageToken string                   `json:"next_page_token"`
	Trades        map[string][]OptionTrade `json:"trades"`
}

type PartialCall struct {
	Currency               string  `json:"currency,omitempty"`
	Cusip                  string  `json:"cusip,omitempty"`
	DividendRate           float64 `json:"dividend_rate,omitempty"`
	ID                     string  `json:"id"`
	Isin                   string  `json:"isin,omitempty"`
	LotteryDate            string  `json:"lottery_date,omitempty"`
	LotteryType            string  `json:"lottery_type,omitempty"`
	PayableDate            string  `json:"payable_date,omitempty"`
	Price                  float64 `json:"price,omitempty"`
	ProcessDate            string  `json:"process_date"`
	RecordDate             string  `json:"record_date,omitempty"`
	ResultsPublicationDate string  `json:"results_publication_date,omitempty"`
	Symbol                 string  `json:"symbol"`
}

type Redemption struct {
	Currency    string  `json:"currency,omitempty"`
	Cusip       string  `json:"cusip"`
	ID          string  `json:"id"`
	Isin        string  `json:"isin,omitempty"`
	PayableDate string  `json:"payable_date,omitempty"`
	ProcessDate string  `json:"process_date"`
	Rate        float64 `json:"rate"`
	Symbol      string  `json:"symbol"`
}

type Reorganization struct {
	CashRate       float64                       `json:"cash_rate,omitempty"`
	Currency       string                        `json:"currency,omitempty"`
	Cusip          string                        `json:"cusip"`
	EffectiveDate  string                        `json:"effective_date"`
	ID             string                        `json:"id"`
	Isin           string                        `json:"isin,omitempty"`
	PayableDate    string                        `json:"payable_date,omitempty"`
	ProcessDate    string                        `json:"process_date"`
	StockMovements []ReorganizationStockMovement `json:"stock_movements,omitempty"`
	Symbol         string                        `json:"symbol"`
}

type ReorganizationStockMovement struct {
	Cusip      string  `json:"cusip"`
	Isin       string  `json:"isin,omitempty"`
	NewRate    float64 `json:"new_rate"`
	SourceRate float64 `json:"source_rate"`
	Symbol     string  `json:"symbol"`
}

type ReverseSplit struct {
	Currency    string  `json:"currency,omitempty"`
	ExDate      string  `json:"ex_date"`
	ID          string  `json:"id"`
	NewCusip    string  `json:"new_cusip"`
	NewIsin     string  `json:"new_isin,omitempty"`
	NewRate     float64 `json:"new_rate"`
	OldCusip    string  `json:"old_cusip"`
	OldIsin     string  `json:"old_isin,omitempty"`
	OldRate     float64 `json:"old_rate"`
	PayableDate string  `json:"payable_date,omitempty"`
	ProcessDate string  `json:"process_date"`
	RecordDate  string  `json:"record_date,omitempty"`
	Symbol      string  `json:"symbol"`
}

type RightsDistribution struct {
	Currency       string  `json:"currency,omitempty"`
	ExDate         string  `json:"ex_date"`
	ExpirationDate string  `json:"expiration_date,omitempty"`
	ID             string  `json:"id"`
	NewCusip       string  `json:"new_cusip"`
	NewIsin        string  `json:"new_isin,omitempty"`
	NewSymbol      string  `json:"new_symbol"`
	PayableDate    string  `json:"payable_date"`
	ProcessDate    string  `json:"process_date"`
	Rate           float64 `json:"rate"`
	RecordDate     string  `json:"record_date,omitempty"`
	SourceCusip    string  `json:"source_cusip"`
	SourceIsin     string  `json:"source_isin,omitempty"`
	SourceSymbol   string  `json:"source_symbol"`
}

type SpinOff struct {
	Currency              string  `json:"currency,omitempty"`
	DueBillRedemptionDate string  `json:"due_bill_redemption_date,omitempty"`
	ExDate                string  `json:"ex_date"`
	ID                    string  `json:"id"`
	NewCusip              string  `json:"new_cusip"`
	NewIsin               string  `json:"new_isin,omitempty"`
	NewRate               float64 `json:"new_rate"`
	NewSymbol             string  `json:"new_symbol"`
	PayableDate           string  `json:"payable_date,omitempty"`
	ProcessDate           string  `json:"process_date"`
	RecordDate            string  `json:"record_date,omitempty"`
	SourceCusip           string  `json:"source_cusip"`
	SourceIsin            string  `json:"source_isin,omitempty"`
	SourceRate            float64 `json:"source_rate"`
	SourceSymbol          string  `json:"source_symbol"`
}

type StockAndCashMerger struct {
	AcquireeCusip  string  `json:"acquiree_cusip"`
	AcquireeIsin   string  `json:"acquiree_isin,omitempty"`
	AcquireeRate   float64 `json:"acquiree_rate"`
	AcquireeSymbol string  `json:"acquiree_symbol"`
	AcquirerCusip  string  `json:"acquirer_cusip"`
	AcquirerIsin   string  `json:"acquirer_isin,omitempty"`
	AcquirerRate   float64 `json:"acquirer_rate"`
	AcquirerSymbol string  `json:"acquirer_symbol"`
	CashRate       float64 `json:"cash_rate"`
	Currency       string  `json:"currency,omitempty"`
	EffectiveDate  string  `json:"effective_date"`
	ID             string  `json:"id"`
	PayableDate    string  `json:"payable_date,omitempty"`
	ProcessDate    string  `json:"process_date"`
}

type StockAuction struct {
	C string  `json:"c"`
	P float64 `json:"p"`
	S int     `json:"s,omitempty"`
	T string  `json:"t"`
	X string  `json:"x"`
}

type StockAuctionsResp struct {
	Auctions      map[string][]StockDailyAuctions `json:"auctions"`
	Currency      string                          `json:"currency,omitempty"`
	NextPageToken string                          `json:"next_page_token"`
}

type StockAuctionsRespSingle struct {
	Auctions      []StockDailyAuctions `json:"auctions"`
	Currency      string               `json:"currency,omitempty"`
	NextPageToken string               `json:"next_page_token"`
	Symbol        string               `json:"symbol"`
}

type StockBar struct {
	C  float64 `json:"c"`
	H  float64 `json:"h"`
	L  float64 `json:"l"`
	N  int     `json:"n"`
	O  float64 `json:"o"`
	T  string  `json:"t"`
	V  int     `json:"v"`
	Vw float64 `json:"vw"`
}

type StockBarsResp struct {
	Bars          map[string][]StockBar `json:"bars"`
	Currency      string                `json:"currency,omitempty"`
	NextPageToken string                `json:"next_page_token"`
}

type StockBarsRespSingle struct {
	Bars          []StockBar `json:"bars"`
	Currency      string     `json:"currency,omitempty"`
	NextPageToken string     `json:"next_page_token"`
	Symbol        string     `json:"symbol"`
}

type StockDailyAuctions struct {
	C []StockAuction `json:"c"`
	D string         `json:"d"`
	O []StockAuction `json:"o"`
}

type StockDividend struct {
	Currency    string  `json:"currency,omitempty"`
	Cusip       string  `json:"cusip"`
	ExDate      string  `json:"ex_date"`
	ID          string  `json:"id"`
	Isin        string  `json:"isin,omitempty"`
	PayableDate string  `json:"payable_date,omitempty"`
	ProcessDate string  `json:"process_date"`
	Rate        float64 `json:"rate"`
	RecordDate  string  `json:"record_date,omitempty"`
	Symbol      string  `json:"symbol"`
}

type StockLatestBarsResp struct {
	Bars     map[string]StockBar `json:"bars"`
	Currency string              `json:"currency,omitempty"`
}

type StockLatestBarsRespSingle struct {
	Bar      StockBar `json:"bar"`
	Currency string   `json:"currency,omitempty"`
	Symbol   string   `json:"symbol"`
}

type StockLatestQuotesResp struct {
	Currency string                `json:"currency,omitempty"`
	Quotes   map[string]StockQuote `json:"quotes"`
}

type StockLatestQuotesRespSingle struct {
	Currency string     `json:"currency,omitempty"`
	Quote    StockQuote `json:"quote"`
	Symbol   string     `json:"symbol"`
}

type StockLatestTradesResp struct {
	Currency string                `json:"currency,omitempty"`
	Trades   map[string]StockTrade `json:"trades"`
}

type StockLatestTradesRespSingle struct {
	Currency string     `json:"currency,omitempty"`
	Symbol   string     `json:"symbol"`
	Trade    StockTrade `json:"trade"`
}

type StockMerger struct {
	AcquireeCusip  string  `json:"acquiree_cusip"`
	AcquireeIsin   string  `json:"acquiree_isin,omitempty"`
	AcquireeRate   float64 `json:"acquiree_rate"`
	AcquireeSymbol string  `json:"acquiree_symbol"`
	AcquirerCusip  string  `json:"acquirer_cusip"`
	AcquirerIsin   string  `json:"acquirer_isin,omitempty"`
	AcquirerRate   float64 `json:"acquirer_rate"`
	AcquirerSymbol string  `json:"acquirer_symbol"`
	Currency       string  `json:"currency,omitempty"`
	EffectiveDate  string  `json:"effective_date"`
	ID             string  `json:"id"`
	PayableDate    string  `json:"payable_date,omitempty"`
	ProcessDate    string  `json:"process_date"`
}

type StockQuote struct {
	Ap float64   `json:"ap"`
	As int       `json:"as"`
	Ax string    `json:"ax"`
	Bp float64   `json:"bp"`
	Bs int       `json:"bs"`
	Bx string    `json:"bx"`
	C  []string  `json:"c"`
	T  string    `json:"t"`
	Z  StockTape `json:"z"`
}

type StockQuotesResp struct {
	Currency      string                  `json:"currency,omitempty"`
	NextPageToken string                  `json:"next_page_token"`
	Quotes        map[string][]StockQuote `json:"quotes"`
}

type StockQuotesRespSingle struct {
	Currency      string       `json:"currency,omitempty"`
	NextPageToken string       `json:"next_page_token"`
	Quotes        []StockQuote `json:"quotes"`
	Symbol        string       `json:"symbol"`
}

type StockSnapshot struct {
	DailyBar     StockBar   `json:"dailyBar,omitempty"`
	LatestQuote  StockQuote `json:"latestQuote,omitempty"`
	LatestTrade  StockTrade `json:"latestTrade,omitempty"`
	MinuteBar    StockBar   `json:"minuteBar,omitempty"`
	PrevDailyBar StockBar   `json:"prevDailyBar,omitempty"`
}

type StockTrade struct {
	C []string  `json:"c"`
	I int       `json:"i"`
	P float64   `json:"p"`
	S int       `json:"s"`
	T string    `json:"t"`
	U string    `json:"u,omitempty"`
	X string    `json:"x"`
	Z StockTape `json:"z"`
}

type StockTradesResp struct {
	Currency      string                  `json:"currency,omitempty"`
	NextPageToken string                  `json:"next_page_token"`
	Trades        map[string][]StockTrade `json:"trades"`
}

type StockTradesRespSingle struct {
	Currency      string       `json:"currency,omitempty"`
	NextPageToken string       `json:"next_page_token"`
	Symbol        string       `json:"symbol"`
	Trades        []StockTrade `json:"trades"`
}

type UnitSplit struct {
	AlternateCusip  string  `json:"alternate_cusip"`
	AlternateIsin   string  `json:"alternate_isin,omitempty"`
	AlternateRate   float64 `json:"alternate_rate"`
	AlternateSymbol string  `json:"alternate_symbol"`
	Currency        string  `json:"currency,omitempty"`
	EffectiveDate   string  `json:"effective_date"`
	ID              string  `json:"id"`
	NewCusip        string  `json:"new_cusip"`
	NewIsin         string  `json:"new_isin,omitempty"`
	NewRate         float64 `json:"new_rate"`
	NewSymbol       string  `json:"new_symbol"`
	OldCusip        string  `json:"old_cusip"`
	OldIsin         string  `json:"old_isin,omitempty"`
	OldRate         float64 `json:"old_rate"`
	OldSymbol       string  `json:"old_symbol"`
	PayableDate     string  `json:"payable_date,omitempty"`
	ProcessDate     string  `json:"process_date"`
}

type WorthlessRemoval struct {
	Currency    string `json:"currency,omitempty"`
	Cusip       string `json:"cusip"`
	ID          string `json:"id"`
	Isin        string `json:"isin,omitempty"`
	ProcessDate string `json:"process_date"`
	Symbol      string `json:"symbol"`
}
