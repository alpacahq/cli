// Code generated from api/specs/trading-api.json; DO NOT EDIT.

package api

type AccountStatus string

type ActivityType string

type AssetAttribute string

type AssetClass string

type CorporateActionCaType string

type CryptoChain string

type CryptoTransferStatus string

type Exchange string

type ExchangeForPosition string

type LocateStatus string

type OptionContractStyle string

type OptionContractType string

type OrderClass string

type OrderSide string

type OrderStatus string

type OrderType string

type PositionIntent string

type TimeInForce string

type TokenizationIssuer string

type TokenizationNetwork string

type TokenizationRequestStatus string

type TokenizationRequestType string

type TransferDirection string

type Market string

type Phase string

type Account struct {
	AccountBlocked           bool          `json:"account_blocked,omitempty"`
	AccountNumber            string        `json:"account_number,omitempty"`
	AccruedFees              string        `json:"accrued_fees,omitempty"`
	BalanceAsof              string        `json:"balance_asof,omitempty"`
	BuyingPower              string        `json:"buying_power,omitempty"`
	Cash                     string        `json:"cash,omitempty"`
	CreatedAt                string        `json:"created_at,omitempty"`
	CryptoStatus             AccountStatus `json:"crypto_status,omitempty"`
	Currency                 string        `json:"currency,omitempty"`
	Equity                   string        `json:"equity,omitempty"`
	ID                       string        `json:"id"`
	InitialMargin            string        `json:"initial_margin,omitempty"`
	IntradayAdjustments      string        `json:"intraday_adjustments,omitempty"`
	LastEquity               string        `json:"last_equity,omitempty"`
	LastMaintenanceMargin    string        `json:"last_maintenance_margin,omitempty"`
	LongMarketValue          string        `json:"long_market_value,omitempty"`
	MaintenanceMargin        string        `json:"maintenance_margin,omitempty"`
	Multiplier               string        `json:"multiplier,omitempty"`
	NonMarginableBuyingPower string        `json:"non_marginable_buying_power,omitempty"`
	OptionsApprovedLevel     int           `json:"options_approved_level,omitempty"`
	OptionsBuyingPower       string        `json:"options_buying_power,omitempty"`
	OptionsTradingLevel      int           `json:"options_trading_level,omitempty"`
	PendingRegTafFees        string        `json:"pending_reg_taf_fees,omitempty"`
	PendingTransferIn        string        `json:"pending_transfer_in,omitempty"`
	PendingTransferOut       string        `json:"pending_transfer_out,omitempty"`
	PortfolioValue           string        `json:"portfolio_value,omitempty"`
	RegtBuyingPower          string        `json:"regt_buying_power,omitempty"`
	ShortMarketValue         string        `json:"short_market_value,omitempty"`
	ShortingEnabled          bool          `json:"shorting_enabled,omitempty"`
	SMA                      string        `json:"sma,omitempty"`
	Status                   AccountStatus `json:"status"`
	TradeSuspendedByUser     bool          `json:"trade_suspended_by_user,omitempty"`
	TradingBlocked           bool          `json:"trading_blocked,omitempty"`
	TransfersBlocked         bool          `json:"transfers_blocked,omitempty"`
}

type AccountConfigurations struct {
	DisableOvernightTrading bool   `json:"disable_overnight_trading,omitempty"`
	FractionalTrading       bool   `json:"fractional_trading,omitempty"`
	MaxMarginMultiplier     string `json:"max_margin_multiplier,omitempty"`
	MaxOptionsTradingLevel  int    `json:"max_options_trading_level,omitempty"`
	NoShorting              bool   `json:"no_shorting,omitempty"`
	PtpNoExceptionEntry     bool   `json:"ptp_no_exception_entry,omitempty"`
	SuspendTrade            bool   `json:"suspend_trade,omitempty"`
	TradeConfirmEmail       string `json:"trade_confirm_email,omitempty"`
}

type ActivityEventV2CommonFields struct {
	ActivitySubtype string `json:"activity_subtype,omitempty"`
	ActivityType    string `json:"activity_type"`
	At              string `json:"at"`
	Currency        string `json:"currency"`
	EventID         string `json:"event_id"`
	ExecutedAt      string `json:"executed_at"`
	NetAmount       string `json:"net_amount,omitempty"`
	PreviousID      string `json:"previous_id,omitempty"`
	Price           string `json:"price,omitempty"`
	Qty             string `json:"qty,omitempty"`
	RefID           string `json:"ref_id"`
	SettleDate      string `json:"settle_date"`
	Status          string `json:"status"`
	SwapFeeBps      string `json:"swap_fee_bps,omitempty"`
	SwapRate        string `json:"swap_rate,omitempty"`
}

type ActivityV2DetailTRD struct {
	AssetID       string `json:"asset_id"`
	ClientOrderID string `json:"client_order_id,omitempty"`
	Commission    string `json:"commission,omitempty"`
	CumQty        string `json:"cum_qty"`
	Cusip         string `json:"cusip,omitempty"`
	ExecutionType string `json:"execution_type"`
	LeavesQty     string `json:"leaves_qty"`
	OrderID       string `json:"order_id"`
	OrderStatus   string `json:"order_status"`
	Side          string `json:"side"`
	Symbol        string `json:"symbol"`
}

type AddAssetToWatchlistRequest struct {
	Symbol string `json:"symbol,omitempty"`
}

type AdvancedInstructions struct {
	Algorithm     string `json:"algorithm,omitempty"`
	Destination   string `json:"destination,omitempty"`
	DisplayQty    string `json:"display_qty,omitempty"`
	EndTime       string `json:"end_time,omitempty"`
	MaxPercentage string `json:"max_percentage,omitempty"`
	StartTime     string `json:"start_time,omitempty"`
}

type Assets struct {
	Attributes                   []AssetAttribute `json:"attributes,omitempty"`
	BorrowStatus                 string           `json:"borrow_status,omitempty"`
	Class                        AssetClass       `json:"class"`
	Cusip                        *string          `json:"cusip,omitempty"`
	EasyToBorrow                 bool             `json:"easy_to_borrow,omitempty"`
	Exchange                     Exchange         `json:"exchange"`
	Fractionable                 bool             `json:"fractionable"`
	ID                           string           `json:"id"`
	MaintenanceMarginRequirement float64          `json:"maintenance_margin_requirement,omitempty"`
	MarginRequirementLong        string           `json:"margin_requirement_long,omitempty"`
	MarginRequirementShort       string           `json:"margin_requirement_short,omitempty"`
	Marginable                   bool             `json:"marginable"`
	MinOrderSize                 string           `json:"min_order_size,omitempty"`
	MinTradeIncrement            string           `json:"min_trade_increment,omitempty"`
	Name                         string           `json:"name"`
	PriceIncrement               string           `json:"price_increment,omitempty"`
	Shortable                    bool             `json:"shortable"`
	Status                       string           `json:"status"`
	Symbol                       string           `json:"symbol"`
	Tradable                     bool             `json:"tradable"`
}

type CanceledOrderResponse struct {
	ID     string `json:"id,omitempty"`
	Status int    `json:"status,omitempty"`
}

type CommonAcatActivityV2 struct {
	ExternalID string `json:"external_id"`
	HoldDate   string `json:"hold_date,omitempty"`
	RequestID  string `json:"request_id"`
}

type CommonCDIVActivityV2 struct {
	Cusip          string `json:"cusip"`
	DueBillOffDate string `json:"due_bill_off_date,omitempty"`
	DueBillOnDate  string `json:"due_bill_on_date,omitempty"`
	ExDate         string `json:"ex_date,omitempty"`
	Foreign        bool   `json:"foreign"`
	Isin           string `json:"isin,omitempty"`
	PayableDate    string `json:"payable_date,omitempty"`
	Rate           string `json:"rate"`
	RecordDate     string `json:"record_date,omitempty"`
	Special        bool   `json:"special"`
	Symbol         string `json:"symbol"`
}

type CommonFixedIncomeInterestActivityV2 struct {
	AccruedInterestRate string `json:"accrued_interest_rate,omitempty"`
	CashPayout          string `json:"cash_payout"`
	Cusip               string `json:"cusip"`
	EntitledQty         string `json:"entitled_qty"`
	Isin                string `json:"isin,omitempty"`
	PaymentDate         string `json:"payment_date,omitempty"`
	Price               string `json:"price,omitempty"`
	Rate                string `json:"rate,omitempty"`
	RecordDate          string `json:"record_date,omitempty"`
}

type CommonMAActivityV2 struct {
	AcquireeCusip  string `json:"acquiree_cusip"`
	AcquireeIsin   string `json:"acquiree_isin,omitempty"`
	AcquireeRate   string `json:"acquiree_rate,omitempty"`
	AcquireeSymbol string `json:"acquiree_symbol"`
	AcquirerCusip  string `json:"acquirer_cusip,omitempty"`
	AcquirerIsin   string `json:"acquirer_isin,omitempty"`
	AcquirerRate   string `json:"acquirer_rate,omitempty"`
	AcquirerSymbol string `json:"acquirer_symbol,omitempty"`
	EffectiveDate  string `json:"effective_date"`
	PayableDate    string `json:"payable_date"`
}

type CommonNCActivityV2 struct {
	NewCusip  string `json:"new_cusip"`
	NewSymbol string `json:"new_symbol"`
	OldCusip  string `json:"old_cusip"`
	OldSymbol string `json:"old_symbol"`
}

type CommonNTAActivityV2 struct {
	GroupID    string `json:"group_id,omitempty"`
	SystemDate string `json:"system_date"`
}

type CommonREOActivityV2 struct {
	CashPayout    string `json:"cash_payout,omitempty"`
	CashRate      string `json:"cash_rate,omitempty"`
	Cusip         string `json:"cusip,omitempty"`
	EffectiveDate string `json:"effective_date,omitempty"`
	Isin          string `json:"isin,omitempty"`
	NewRate       string `json:"new_rate,omitempty"`
	PayableDate   string `json:"payable_date"`
	Qty           string `json:"qty,omitempty"`
	RemovedQty    string `json:"removed_qty,omitempty"`
	SourceCusip   string `json:"source_cusip"`
	SourceIsin    string `json:"source_isin,omitempty"`
	SourceQty     string `json:"source_qty,omitempty"`
	SourceRate    string `json:"source_rate,omitempty"`
	SourceSymbol  string `json:"source_symbol"`
	Symbol        string `json:"symbol,omitempty"`
}

type CommonSDIVActivityV2 struct {
	Cusip       string `json:"cusip"`
	ExDate      string `json:"ex_date,omitempty"`
	Isin        string `json:"isin,omitempty"`
	PayableDate string `json:"payable_date,omitempty"`
	Rate        string `json:"rate"`
	RecordDate  string `json:"record_date,omitempty"`
	Symbol      string `json:"symbol"`
}

type CommonSpinoffActivityV2 struct {
	DueBillRedemptionDate string `json:"due_bill_redemption_date,omitempty"`
	ExDate                string `json:"ex_date,omitempty"`
	NewCusip              string `json:"new_cusip"`
	NewIsin               string `json:"new_isin,omitempty"`
	NewPrice              string `json:"new_price"`
	NewRate               string `json:"new_rate"`
	NewSymbol             string `json:"new_symbol"`
	PayableDate           string `json:"payable_date,omitempty"`
	RecordDate            string `json:"record_date,omitempty"`
	SourceCusip           string `json:"source_cusip"`
	SourceIsin            string `json:"source_isin,omitempty"`
	SourcePrice           string `json:"source_price"`
	SourceRate            string `json:"source_rate"`
	SourceSymbol          string `json:"source_symbol"`
}

type CommonSplitActivityV2 struct {
	NewCusip    string `json:"new_cusip"`
	NewIsin     string `json:"new_isin,omitempty"`
	NewRate     string `json:"new_rate"`
	OldCusip    string `json:"old_cusip"`
	OldIsin     string `json:"old_isin,omitempty"`
	OldRate     string `json:"old_rate"`
	PayableDate string `json:"payable_date,omitempty"`
}

type CommonVOFSubtypeActivityV2 struct {
	NewCusip     string `json:"new_cusip,omitempty"`
	NewSymbol    string `json:"new_symbol,omitempty"`
	SourceCusip  string `json:"source_cusip"`
	SourceSymbol string `json:"source_symbol"`
}

type CorporateAnnouncement struct {
	CaSubType               string                `json:"ca_sub_type,omitempty"`
	CaType                  CorporateActionCaType `json:"ca_type,omitempty"`
	Cash                    string                `json:"cash,omitempty"`
	CorporateActionID       string                `json:"corporate_action_id,omitempty"`
	DeclarationDate         string                `json:"declaration_date,omitempty"`
	EffectiveDate           string                `json:"effective_date,omitempty"`
	ExDate                  string                `json:"ex_date,omitempty"`
	ID                      string                `json:"id,omitempty"`
	InitiatingOriginalCusip string                `json:"initiating_original_cusip,omitempty"`
	InitiatingSymbol        string                `json:"initiating_symbol,omitempty"`
	NewRate                 string                `json:"new_rate,omitempty"`
	OldRate                 string                `json:"old_rate,omitempty"`
	PayableDate             string                `json:"payable_date,omitempty"`
	RecordDate              string                `json:"record_date,omitempty"`
	TargetOriginalCusip     string                `json:"target_original_cusip,omitempty"`
	TargetSymbol            string                `json:"target_symbol,omitempty"`
}

type CreateCryptoTransferRequest struct {
	Address string      `json:"address"`
	Amount  string      `json:"amount"`
	Asset   string      `json:"asset"`
	Chain   CryptoChain `json:"chain,omitempty"`
}

type CreateLocateRequest struct {
	AllOrNone  bool    `json:"all_or_none,omitempty"`
	LimitPrice *string `json:"limit_price,omitempty"`
	Qty        int     `json:"qty"`
	Symbol     string  `json:"symbol"`
}

type CreateOrderRequest struct {
	AdvancedInstructions AdvancedInstructions `json:"advanced_instructions,omitempty"`
	ClientOrderID        string               `json:"client_order_id,omitempty"`
	ExtendedHours        bool                 `json:"extended_hours,omitempty"`
	Legs                 []MLegOrderLeg       `json:"legs,omitempty"`
	LimitPrice           string               `json:"limit_price,omitempty"`
	Notional             string               `json:"notional,omitempty"`
	OrderClass           OrderClass           `json:"order_class,omitempty"`
	PositionIntent       PositionIntent       `json:"position_intent,omitempty"`
	Qty                  string               `json:"qty,omitempty"`
	Side                 OrderSide            `json:"side,omitempty"`
	StopLoss             map[string]any       `json:"stop_loss,omitempty"`
	StopPrice            string               `json:"stop_price,omitempty"`
	Symbol               string               `json:"symbol,omitempty"`
	TakeProfit           map[string]any       `json:"take_profit,omitempty"`
	TimeInForce          TimeInForce          `json:"time_in_force"`
	TrailPercent         string               `json:"trail_percent,omitempty"`
	TrailPrice           string               `json:"trail_price,omitempty"`
	Type                 OrderType            `json:"type"`
}

type CreateWatchlistRequest struct {
	Name    string   `json:"name"`
	Symbols []string `json:"symbols,omitempty"`
}

type CreateWhitelistedAddressRequest struct {
	Address string      `json:"address,omitempty"`
	Asset   string      `json:"asset,omitempty"`
	Chain   CryptoChain `json:"chain,omitempty"`
}

type CryptoTransfer struct {
	Amount      string               `json:"amount,omitempty"`
	Asset       string               `json:"asset,omitempty"`
	Chain       string               `json:"chain,omitempty"`
	CreatedAt   string               `json:"created_at,omitempty"`
	Direction   TransferDirection    `json:"direction,omitempty"`
	Fees        string               `json:"fees,omitempty"`
	FromAddress string               `json:"from_address,omitempty"`
	ID          string               `json:"id,omitempty"`
	NetworkFee  string               `json:"network_fee,omitempty"`
	Status      CryptoTransferStatus `json:"status,omitempty"`
	ToAddress   string               `json:"to_address,omitempty"`
	TxHash      string               `json:"tx_hash,omitempty"`
	UsdValue    string               `json:"usd_value,omitempty"`
}

type CryptoWallet struct {
	Address   string `json:"address,omitempty"`
	Chain     string `json:"chain,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type Error struct {
	Code    float64 `json:"code"`
	Message string  `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

type ListLocateQuotesResponse struct {
	Errors []LocateQuoteError `json:"errors,omitempty"`
	Quotes []LocateQuote      `json:"quotes"`
}

type ListLocatesResponse struct {
	Locates       []Locate `json:"locates"`
	NextPageToken *string  `json:"next_page_token"`
}

type Locate struct {
	AllOrNone       bool         `json:"all_or_none"`
	CreatedAt       string       `json:"created_at"`
	ExpiresAt       *string      `json:"expires_at,omitempty"`
	ID              string       `json:"id"`
	LimitPrice      *string      `json:"limit_price,omitempty"`
	LocatedPrice    *string      `json:"located_price,omitempty"`
	LocatedQty      *int         `json:"located_qty,omitempty"`
	RejectionReason *string      `json:"rejection_reason,omitempty"`
	RequestedQty    int          `json:"requested_qty"`
	Status          LocateStatus `json:"status"`
	Symbol          string       `json:"symbol"`
	TotalFee        *string      `json:"total_fee,omitempty"`
}

type LocateError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
}

type LocateQuote struct {
	AvailableQty int    `json:"available_qty"`
	Price        string `json:"price,omitempty"`
	QuotedAt     string `json:"quoted_at"`
	Symbol       string `json:"symbol"`
}

type LocateQuoteError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Symbol  string `json:"symbol"`
}

type MLegOrderLeg struct {
	PositionIntent PositionIntent `json:"position_intent,omitempty"`
	RatioQty       string         `json:"ratio_qty"`
	Side           OrderSide      `json:"side,omitempty"`
	Symbol         string         `json:"symbol"`
}

type NonTradeActivities struct {
	ActivitySubType string       `json:"activity_sub_type,omitempty"`
	ActivityType    ActivityType `json:"activity_type,omitempty"`
	CreatedAt       string       `json:"created_at,omitempty"`
	Currency        string       `json:"currency,omitempty"`
	Cusip           string       `json:"cusip,omitempty"`
	Date            string       `json:"date,omitempty"`
	GroupID         string       `json:"group_id,omitempty"`
	ID              string       `json:"id,omitempty"`
	NetAmount       string       `json:"net_amount,omitempty"`
	PerShareAmount  string       `json:"per_share_amount,omitempty"`
	Qty             string       `json:"qty,omitempty"`
	Status          string       `json:"status,omitempty"`
	Symbol          string       `json:"symbol,omitempty"`
}

type OptionContract struct {
	ClosePrice        string              `json:"close_price,omitempty"`
	ClosePriceDate    string              `json:"close_price_date,omitempty"`
	Deliverables      []OptionDeliverable `json:"deliverables,omitempty"`
	ExpirationDate    string              `json:"expiration_date"`
	ID                string              `json:"id"`
	Multiplier        string              `json:"multiplier"`
	Name              string              `json:"name"`
	OpenInterest      string              `json:"open_interest,omitempty"`
	OpenInterestDate  string              `json:"open_interest_date,omitempty"`
	RootSymbol        string              `json:"root_symbol,omitempty"`
	Size              string              `json:"size"`
	Status            string              `json:"status"`
	StrikePrice       string              `json:"strike_price"`
	Style             OptionContractStyle `json:"style"`
	Symbol            string              `json:"symbol"`
	Tradable          bool                `json:"tradable"`
	Type              OptionContractType  `json:"type"`
	UnderlyingAssetID string              `json:"underlying_asset_id"`
	UnderlyingSymbol  string              `json:"underlying_symbol"`
}

type OptionContractsResponse struct {
	NextPageToken   string           `json:"next_page_token,omitempty"`
	OptionContracts []OptionContract `json:"option_contracts"`
}

type OptionDeliverable struct {
	AllocationPercentage string `json:"allocation_percentage"`
	Amount               string `json:"amount"`
	AssetID              string `json:"asset_id,omitempty"`
	DelayedSettlement    bool   `json:"delayed_settlement"`
	SettlementMethod     string `json:"settlement_method"`
	SettlementType       string `json:"settlement_type"`
	Symbol               string `json:"symbol"`
	Type                 string `json:"type"`
}

type Order struct {
	AssetClass     AssetClass     `json:"asset_class,omitempty"`
	AssetID        string         `json:"asset_id,omitempty"`
	CanceledAt     *string        `json:"canceled_at,omitempty"`
	ClientOrderID  string         `json:"client_order_id,omitempty"`
	CreatedAt      string         `json:"created_at,omitempty"`
	ExpiredAt      *string        `json:"expired_at,omitempty"`
	ExpiresAt      string         `json:"expires_at,omitempty"`
	ExtendedHours  bool           `json:"extended_hours,omitempty"`
	FailedAt       *string        `json:"failed_at,omitempty"`
	FilledAt       *string        `json:"filled_at,omitempty"`
	FilledAvgPrice *string        `json:"filled_avg_price,omitempty"`
	FilledQty      string         `json:"filled_qty,omitempty"`
	Hwm            *string        `json:"hwm,omitempty"`
	ID             string         `json:"id,omitempty"`
	Legs           []OrderLeg     `json:"legs,omitempty"`
	LimitPrice     *string        `json:"limit_price,omitempty"`
	Notional       *string        `json:"notional"`
	OrderClass     OrderClass     `json:"order_class,omitempty"`
	OrderType      string         `json:"order_type,omitempty"`
	PositionIntent PositionIntent `json:"position_intent,omitempty"`
	Qty            *string        `json:"qty,omitempty"`
	RatioQty       *string        `json:"ratio_qty,omitempty"`
	ReplacedAt     *string        `json:"replaced_at,omitempty"`
	ReplacedBy     *string        `json:"replaced_by,omitempty"`
	Replaces       *string        `json:"replaces,omitempty"`
	Side           OrderSide      `json:"side,omitempty"`
	Status         OrderStatus    `json:"status,omitempty"`
	StopPrice      *string        `json:"stop_price,omitempty"`
	SubmittedAt    *string        `json:"submitted_at,omitempty"`
	Symbol         string         `json:"symbol,omitempty"`
	TimeInForce    TimeInForce    `json:"time_in_force"`
	TrailPercent   *string        `json:"trail_percent,omitempty"`
	TrailPrice     *string        `json:"trail_price,omitempty"`
	Type           OrderType      `json:"type"`
	UpdatedAt      *string        `json:"updated_at,omitempty"`
}

type OrderLeg struct {
	AssetClass     AssetClass     `json:"asset_class,omitempty"`
	AssetID        string         `json:"asset_id,omitempty"`
	CanceledAt     *string        `json:"canceled_at,omitempty"`
	ClientOrderID  string         `json:"client_order_id,omitempty"`
	CreatedAt      string         `json:"created_at,omitempty"`
	ExpiredAt      *string        `json:"expired_at,omitempty"`
	ExtendedHours  bool           `json:"extended_hours,omitempty"`
	FailedAt       *string        `json:"failed_at,omitempty"`
	FilledAt       *string        `json:"filled_at,omitempty"`
	FilledAvgPrice *string        `json:"filled_avg_price,omitempty"`
	FilledQty      string         `json:"filled_qty,omitempty"`
	Hwm            *string        `json:"hwm,omitempty"`
	ID             string         `json:"id,omitempty"`
	Legs           any            `json:"legs,omitempty"`
	LimitPrice     *string        `json:"limit_price,omitempty"`
	Notional       *string        `json:"notional"`
	OrderClass     OrderClass     `json:"order_class,omitempty"`
	OrderType      string         `json:"order_type,omitempty"`
	PositionIntent PositionIntent `json:"position_intent,omitempty"`
	Qty            *string        `json:"qty"`
	ReplacedAt     *string        `json:"replaced_at,omitempty"`
	ReplacedBy     *string        `json:"replaced_by,omitempty"`
	Replaces       *string        `json:"replaces,omitempty"`
	Side           OrderSide      `json:"side"`
	Status         OrderStatus    `json:"status,omitempty"`
	StopPrice      *string        `json:"stop_price,omitempty"`
	SubmittedAt    *string        `json:"submitted_at,omitempty"`
	Symbol         string         `json:"symbol"`
	TimeInForce    TimeInForce    `json:"time_in_force"`
	TrailPercent   *string        `json:"trail_percent,omitempty"`
	TrailPrice     *string        `json:"trail_price,omitempty"`
	Type           OrderType      `json:"type"`
	UpdatedAt      *string        `json:"updated_at,omitempty"`
}

type PatchOrderRequest struct {
	AdvancedInstructions AdvancedInstructions `json:"advanced_instructions,omitempty"`
	ClientOrderID        string               `json:"client_order_id,omitempty"`
	LimitPrice           string               `json:"limit_price,omitempty"`
	Notional             string               `json:"notional,omitempty"`
	Qty                  string               `json:"qty,omitempty"`
	StopPrice            string               `json:"stop_price,omitempty"`
	TimeInForce          TimeInForce          `json:"time_in_force,omitempty"`
	Trail                string               `json:"trail,omitempty"`
}

type PortfolioHistory struct {
	BaseValue     *float64             `json:"base_value"`
	BaseValueAsof string               `json:"base_value_asof,omitempty"`
	Cashflow      map[string][]float64 `json:"cashflow,omitempty"`
	Equity        []float64            `json:"equity"`
	ProfitLoss    []float64            `json:"profit_loss"`
	ProfitLossPct []float64            `json:"profit_loss_pct"`
	Timeframe     string               `json:"timeframe"`
	Timestamp     []int                `json:"timestamp"`
}

type Position struct {
	AssetClass             AssetClass          `json:"asset_class"`
	AssetID                string              `json:"asset_id"`
	AssetMarginable        bool                `json:"asset_marginable"`
	AvgEntryPrice          string              `json:"avg_entry_price"`
	AvgEntrySwapRate       string              `json:"avg_entry_swap_rate,omitempty"`
	ChangeToday            string              `json:"change_today"`
	CostBasis              string              `json:"cost_basis"`
	CurrentPrice           string              `json:"current_price"`
	Exchange               ExchangeForPosition `json:"exchange"`
	LastdayPrice           string              `json:"lastday_price"`
	MarketValue            string              `json:"market_value"`
	PrevSwapRate           string              `json:"prev_swap_rate,omitempty"`
	Qty                    string              `json:"qty"`
	QtyAvailable           string              `json:"qty_available,omitempty"`
	Side                   string              `json:"side"`
	SwapRate               string              `json:"swap_rate,omitempty"`
	Symbol                 string              `json:"symbol"`
	UnrealizedIntradayPL   string              `json:"unrealized_intraday_pl"`
	UnrealizedIntradayPlpc string              `json:"unrealized_intraday_plpc"`
	UnrealizedPL           string              `json:"unrealized_pl"`
	UnrealizedPlpc         string              `json:"unrealized_plpc"`
	Usd                    USDPositionValues   `json:"usd,omitempty"`
}

type PositionClosedReponse struct {
	Body   Order  `json:"body,omitempty"`
	Status int    `json:"status"`
	Symbol string `json:"symbol"`
}

type TokenizationMintRequest struct {
	Issuer           TokenizationIssuer  `json:"issuer"`
	Network          TokenizationNetwork `json:"network"`
	Qty              string              `json:"qty"`
	UnderlyingSymbol string              `json:"underlying_symbol"`
	WalletAddress    string              `json:"wallet_address"`
}

type TokenizationMintResponse struct {
	CreatedAt             string                    `json:"created_at"`
	Issuer                TokenizationIssuer        `json:"issuer"`
	Network               TokenizationNetwork       `json:"network"`
	Qty                   string                    `json:"qty"`
	Status                TokenizationRequestStatus `json:"status"`
	TokenSymbol           string                    `json:"token_symbol"`
	TokenizationRequestID string                    `json:"tokenization_request_id"`
	UnderlyingSymbol      string                    `json:"underlying_symbol"`
}

type TokenizationRequest struct {
	Account                 string                    `json:"account,omitempty"`
	ClientAccountID         string                    `json:"client_account_id,omitempty"`
	ClientExternalAccountID string                    `json:"client_external_account_id,omitempty"`
	ClientRequestID         string                    `json:"client_request_id,omitempty"`
	CreatedAt               string                    `json:"created_at"`
	Fees                    *string                   `json:"fees,omitempty"`
	Issuer                  TokenizationIssuer        `json:"issuer"`
	IssuerAccount           string                    `json:"issuer_account,omitempty"`
	IssuerRequestID         *string                   `json:"issuer_request_id,omitempty"`
	Network                 TokenizationNetwork       `json:"network"`
	Qty                     string                    `json:"qty"`
	Status                  TokenizationRequestStatus `json:"status"`
	TokenSymbol             string                    `json:"token_symbol"`
	TokenizationRequestID   string                    `json:"tokenization_request_id"`
	TxHash                  *string                   `json:"tx_hash,omitempty"`
	Type                    TokenizationRequestType   `json:"type"`
	UnderlyingSymbol        string                    `json:"underlying_symbol"`
	UpdatedAt               *string                   `json:"updated_at,omitempty"`
	WalletAddress           string                    `json:"wallet_address"`
}

type TradingActivities struct {
	ActivityType    ActivityType `json:"activity_type,omitempty"`
	CumQty          string       `json:"cum_qty,omitempty"`
	ID              string       `json:"id,omitempty"`
	LeavesQty       string       `json:"leaves_qty,omitempty"`
	OrderID         string       `json:"order_id,omitempty"`
	OrderStatus     OrderStatus  `json:"order_status,omitempty"`
	Price           string       `json:"price,omitempty"`
	Qty             string       `json:"qty,omitempty"`
	Side            string       `json:"side,omitempty"`
	Symbol          string       `json:"symbol,omitempty"`
	TransactionTime string       `json:"transaction_time,omitempty"`
	Type            string       `json:"type,omitempty"`
}

type USDPositionValues struct {
	AvgEntryPrice          string `json:"avg_entry_price"`
	ChangeToday            string `json:"change_today,omitempty"`
	CostBasis              string `json:"cost_basis"`
	CurrentPrice           string `json:"current_price,omitempty"`
	LastdayPrice           string `json:"lastday_price,omitempty"`
	MarketValue            string `json:"market_value,omitempty"`
	UnrealizedIntradayPL   string `json:"unrealized_intraday_pl,omitempty"`
	UnrealizedIntradayPlpc string `json:"unrealized_intraday_plpc,omitempty"`
	UnrealizedPL           string `json:"unrealized_pl,omitempty"`
	UnrealizedPlpc         string `json:"unrealized_plpc,omitempty"`
}

type UpdateWatchlistRequest struct {
	Name    string   `json:"name,omitempty"`
	Symbols []string `json:"symbols,omitempty"`
}

type WalletFeeEstimateResponse struct {
	Fee        string `json:"fee,omitempty"`
	NetworkFee string `json:"network_fee,omitempty"`
}

type Watchlist struct {
	AccountID string   `json:"account_id"`
	Assets    []Assets `json:"assets,omitempty"`
	CreatedAt string   `json:"created_at"`
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	UpdatedAt string   `json:"updated_at"`
}

type WatchlistWithoutAsset struct {
	AccountID string `json:"account_id"`
	CreatedAt string `json:"created_at"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	UpdatedAt string `json:"updated_at"`
}

type WhitelistedAddress struct {
	Address   string `json:"address,omitempty"`
	Asset     string `json:"asset,omitempty"`
	Chain     string `json:"chain,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	ID        string `json:"id,omitempty"`
	Status    string `json:"status,omitempty"`
}

type CalendarDay struct {
	CoreEnd        string `json:"core_end"`
	CoreStart      string `json:"core_start"`
	Date           string `json:"date"`
	LunchEnd       string `json:"lunch_end,omitempty"`
	LunchStart     string `json:"lunch_start,omitempty"`
	PostEnd        string `json:"post_end,omitempty"`
	PostStart      string `json:"post_start,omitempty"`
	PreEnd         string `json:"pre_end,omitempty"`
	PreStart       string `json:"pre_start,omitempty"`
	SettlementDate string `json:"settlement_date,omitempty"`
}

type Clock struct {
	IsMarketDay     bool         `json:"is_market_day"`
	Market          PublicMarket `json:"market"`
	NextMarketClose string       `json:"next_market_close"`
	NextMarketOpen  string       `json:"next_market_open"`
	Phase           Phase        `json:"phase"`
	PhaseUntil      string       `json:"phase_until"`
	Timestamp       string       `json:"timestamp"`
}

type ClockResp struct {
	Clocks []Clock `json:"clocks"`
}

type LegacyCalendarDay struct {
	Close          string `json:"close"`
	Date           string `json:"date"`
	Open           string `json:"open"`
	SessionClose   string `json:"session_close"`
	SessionOpen    string `json:"session_open"`
	SettlementDate string `json:"settlement_date"`
}

type LegacyClock struct {
	IsOpen    bool   `json:"is_open"`
	NextClose string `json:"next_close"`
	NextOpen  string `json:"next_open"`
	Timestamp string `json:"timestamp"`
}

type PublicCalendarResp struct {
	Calendar []CalendarDay `json:"calendar"`
	Market   PublicMarket  `json:"market"`
}

type PublicMarket struct {
	Acronym  string `json:"acronym"`
	Bic      string `json:"bic,omitempty"`
	Mic      string `json:"mic,omitempty"`
	Name     string `json:"name"`
	Timezone string `json:"timezone"`
}
