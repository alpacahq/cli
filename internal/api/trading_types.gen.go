// Code generated from api/specs/trading-api.json; DO NOT EDIT.

package api

type AccountStatus string

const (
	AccountStatusACCOUNTUPDATED   AccountStatus = "ACCOUNT_UPDATED"
	AccountStatusACTIVE           AccountStatus = "ACTIVE"
	AccountStatusAPPROVALPENDING  AccountStatus = "APPROVAL_PENDING"
	AccountStatusONBOARDING       AccountStatus = "ONBOARDING"
	AccountStatusREJECTED         AccountStatus = "REJECTED"
	AccountStatusSUBMISSIONFAILED AccountStatus = "SUBMISSION_FAILED"
	AccountStatusSUBMITTED        AccountStatus = "SUBMITTED"
)

var AccountStatusValues = []string{"ACCOUNT_UPDATED", "ACTIVE", "APPROVAL_PENDING", "ONBOARDING", "REJECTED", "SUBMISSION_FAILED", "SUBMITTED"}

type ActivityType string

const (
	ActivityTypeACATC   ActivityType = "ACATC"
	ActivityTypeACATS   ActivityType = "ACATS"
	ActivityTypeCFEE    ActivityType = "CFEE"
	ActivityTypeCSD     ActivityType = "CSD"
	ActivityTypeCSW     ActivityType = "CSW"
	ActivityTypeDIV     ActivityType = "DIV"
	ActivityTypeDIVCGL  ActivityType = "DIVCGL"
	ActivityTypeDIVCGS  ActivityType = "DIVCGS"
	ActivityTypeDIVFEE  ActivityType = "DIVFEE"
	ActivityTypeDIVFT   ActivityType = "DIVFT"
	ActivityTypeDIVNRA  ActivityType = "DIVNRA"
	ActivityTypeDIVROC  ActivityType = "DIVROC"
	ActivityTypeDIVTW   ActivityType = "DIVTW"
	ActivityTypeDIVTXEX ActivityType = "DIVTXEX"
	ActivityTypeFEE     ActivityType = "FEE"
	ActivityTypeFILL    ActivityType = "FILL"
	ActivityTypeFOPT    ActivityType = "FOPT"
	ActivityTypeINT     ActivityType = "INT"
	ActivityTypeINTNRA  ActivityType = "INTNRA"
	ActivityTypeINTTW   ActivityType = "INTTW"
	ActivityTypeJNL     ActivityType = "JNL"
	ActivityTypeJNLC    ActivityType = "JNLC"
	ActivityTypeJNLS    ActivityType = "JNLS"
	ActivityTypeMA      ActivityType = "MA"
	ActivityTypeMISC    ActivityType = "MISC"
	ActivityTypeNC      ActivityType = "NC"
	ActivityTypeOPASN   ActivityType = "OPASN"
	ActivityTypeOPCA    ActivityType = "OPCA"
	ActivityTypeOPCSH   ActivityType = "OPCSH"
	ActivityTypeOPEXC   ActivityType = "OPEXC"
	ActivityTypeOPEXP   ActivityType = "OPEXP"
	ActivityTypeOPTRD   ActivityType = "OPTRD"
	ActivityTypePTC     ActivityType = "PTC"
	ActivityTypePTR     ActivityType = "PTR"
	ActivityTypeREORG   ActivityType = "REORG"
	ActivityTypeSPIN    ActivityType = "SPIN"
	ActivityTypeSPLIT   ActivityType = "SPLIT"
	ActivityTypeTRANS   ActivityType = "TRANS"
)

var ActivityTypeValues = []string{"ACATC", "ACATS", "CFEE", "CSD", "CSW", "DIV", "DIVCGL", "DIVCGS", "DIVFEE", "DIVFT", "DIVNRA", "DIVROC", "DIVTW", "DIVTXEX", "FEE", "FILL", "FOPT", "INT", "INTNRA", "INTTW", "JNL", "JNLC", "JNLS", "MA", "MISC", "NC", "OPASN", "OPCA", "OPCSH", "OPEXC", "OPEXP", "OPTRD", "PTC", "PTR", "REORG", "SPIN", "SPLIT", "TRANS"}

type AssetClass string

const (
	AssetClassCrypto   AssetClass = "crypto"
	AssetClassUsEquity AssetClass = "us_equity"
	AssetClassUsOption AssetClass = "us_option"
)

var AssetClassValues = []string{"crypto", "us_equity", "us_option"}

type CryptoTransferStatus string

const (
	CryptoTransferStatusCOMPLETE   CryptoTransferStatus = "COMPLETE"
	CryptoTransferStatusFAILED     CryptoTransferStatus = "FAILED"
	CryptoTransferStatusPROCESSING CryptoTransferStatus = "PROCESSING"
)

var CryptoTransferStatusValues = []string{"COMPLETE", "FAILED", "PROCESSING"}

type Exchange string

const (
	ExchangeAMEX     Exchange = "AMEX"
	ExchangeARCA     Exchange = "ARCA"
	ExchangeBATS     Exchange = "BATS"
	ExchangeNASDAQ   Exchange = "NASDAQ"
	ExchangeNYSE     Exchange = "NYSE"
	ExchangeNYSEARCA Exchange = "NYSEARCA"
	ExchangeOTC      Exchange = "OTC"
)

var ExchangeValues = []string{"AMEX", "ARCA", "BATS", "NASDAQ", "NYSE", "NYSEARCA", "OTC"}

type ExchangeForPosition string

const (
	ExchangeForPositionAMEX     ExchangeForPosition = "AMEX"
	ExchangeForPositionARCA     ExchangeForPosition = "ARCA"
	ExchangeForPositionBATS     ExchangeForPosition = "BATS"
	ExchangeForPositionNASDAQ   ExchangeForPosition = "NASDAQ"
	ExchangeForPositionNYSE     ExchangeForPosition = "NYSE"
	ExchangeForPositionNYSEARCA ExchangeForPosition = "NYSEARCA"
	ExchangeForPositionOTC      ExchangeForPosition = "OTC"
)

var ExchangeForPositionValues = []string{"AMEX", "ARCA", "BATS", "NASDAQ", "NYSE", "NYSEARCA", "OTC"}

type OrderClass string

const (
	OrderClassBracket OrderClass = "bracket"
	OrderClassMleg    OrderClass = "mleg"
	OrderClassOco     OrderClass = "oco"
	OrderClassOto     OrderClass = "oto"
	OrderClassSimple  OrderClass = "simple"
)

var OrderClassValues = []string{"bracket", "mleg", "oco", "oto", "simple"}

type OrderSide string

const (
	OrderSideBuy  OrderSide = "buy"
	OrderSideSell OrderSide = "sell"
)

var OrderSideValues = []string{"buy", "sell"}

type OrderStatus string

const (
	OrderStatusAccepted           OrderStatus = "accepted"
	OrderStatusAcceptedForBidding OrderStatus = "accepted_for_bidding"
	OrderStatusCalculated         OrderStatus = "calculated"
	OrderStatusCanceled           OrderStatus = "canceled"
	OrderStatusDoneForDay         OrderStatus = "done_for_day"
	OrderStatusExpired            OrderStatus = "expired"
	OrderStatusFilled             OrderStatus = "filled"
	OrderStatusNew                OrderStatus = "new"
	OrderStatusPartiallyFilled    OrderStatus = "partially_filled"
	OrderStatusPendingCancel      OrderStatus = "pending_cancel"
	OrderStatusPendingNew         OrderStatus = "pending_new"
	OrderStatusPendingReplace     OrderStatus = "pending_replace"
	OrderStatusRejected           OrderStatus = "rejected"
	OrderStatusReplaced           OrderStatus = "replaced"
	OrderStatusStopped            OrderStatus = "stopped"
	OrderStatusSuspended          OrderStatus = "suspended"
)

var OrderStatusValues = []string{"accepted", "accepted_for_bidding", "calculated", "canceled", "done_for_day", "expired", "filled", "new", "partially_filled", "pending_cancel", "pending_new", "pending_replace", "rejected", "replaced", "stopped", "suspended"}

type OrderType string

const (
	OrderTypeLimit        OrderType = "limit"
	OrderTypeMarket       OrderType = "market"
	OrderTypeStop         OrderType = "stop"
	OrderTypeStopLimit    OrderType = "stop_limit"
	OrderTypeTrailingStop OrderType = "trailing_stop"
)

var OrderTypeValues = []string{"limit", "market", "stop", "stop_limit", "trailing_stop"}

type PositionIntent string

const (
	PositionIntentBuyToClose  PositionIntent = "buy_to_close"
	PositionIntentBuyToOpen   PositionIntent = "buy_to_open"
	PositionIntentSellToClose PositionIntent = "sell_to_close"
	PositionIntentSellToOpen  PositionIntent = "sell_to_open"
)

var PositionIntentValues = []string{"buy_to_close", "buy_to_open", "sell_to_close", "sell_to_open"}

type TimeInForce string

const (
	TimeInForceCls TimeInForce = "cls"
	TimeInForceDay TimeInForce = "day"
	TimeInForceFok TimeInForce = "fok"
	TimeInForceGtc TimeInForce = "gtc"
	TimeInForceIoc TimeInForce = "ioc"
	TimeInForceOpg TimeInForce = "opg"
)

var TimeInForceValues = []string{"cls", "day", "fok", "gtc", "ioc", "opg"}

type TransferDirection string

const (
	TransferDirectionINCOMING TransferDirection = "INCOMING"
	TransferDirectionOUTGOING TransferDirection = "OUTGOING"
)

var TransferDirectionValues = []string{"INCOMING", "OUTGOING"}

type BondStatus string

const (
	BondStatusMatured     BondStatus = "matured"
	BondStatusOutstanding BondStatus = "outstanding"
	BondStatusPreIssuance BondStatus = "pre_issuance"
)

var BondStatusValues = []string{"matured", "outstanding", "pre_issuance"}

type CallType string

const (
	CallTypeMakeWhole  CallType = "make_whole"
	CallTypeOrdinary   CallType = "ordinary"
	CallTypeRegulatory CallType = "regulatory"
	CallTypeSpecial    CallType = "special"
)

var CallTypeValues = []string{"make_whole", "ordinary", "regulatory", "special"}

type CouponFrequency string

const (
	CouponFrequencyAnnual     CouponFrequency = "annual"
	CouponFrequencyMonthly    CouponFrequency = "monthly"
	CouponFrequencyQuarterly  CouponFrequency = "quarterly"
	CouponFrequencySemiAnnual CouponFrequency = "semi_annual"
	CouponFrequencyZero       CouponFrequency = "zero"
)

var CouponFrequencyValues = []string{"annual", "monthly", "quarterly", "semi_annual", "zero"}

type CouponType string

const (
	CouponTypeFixed    CouponType = "fixed"
	CouponTypeFloating CouponType = "floating"
	CouponTypeZero     CouponType = "zero"
)

var CouponTypeValues = []string{"fixed", "floating", "zero"}

type DayCount string

const (
	DayCountV30x360  DayCount = "30/360"
	DayCountV30x365  DayCount = "30/365"
	DayCountV30Ex360 DayCount = "30E/360"
	DayCountAx360    DayCount = "A/360"
	DayCountAx364    DayCount = "A/364"
	DayCountAx365    DayCount = "A/365"
	DayCountAxA      DayCount = "A/A"
	DayCountBx252    DayCount = "B/252"
)

var DayCountValues = []string{"30/360", "30/365", "30E/360", "A/360", "A/364", "A/365", "A/A", "B/252"}

type Market string

const (
	MarketBMO     Market = "BMO"
	MarketBNYM    Market = "BNYM"
	MarketBOATS   Market = "BOATS"
	MarketCEUX    Market = "CEUX"
	MarketCHIX    Market = "CHIX"
	MarketHKEX    Market = "HKEX"
	MarketIEX     Market = "IEX"
	MarketIEXG    Market = "IEXG"
	MarketISE     Market = "ISE"
	MarketLSE     Market = "LSE"
	MarketMTA     Market = "MTA"
	MarketMTAA    Market = "MTAA"
	MarketNASDAQ  Market = "NASDAQ"
	MarketNYSE    Market = "NYSE"
	MarketOCEA    Market = "OCEA"
	MarketOPRA    Market = "OPRA"
	MarketOTC     Market = "OTC"
	MarketOTCM    Market = "OTCM"
	MarketSIFMA   Market = "SIFMA"
	MarketTADAWUL Market = "TADAWUL"
	MarketXAMS    Market = "XAMS"
	MarketXBRU    Market = "XBRU"
	MarketXDUB    Market = "XDUB"
	MarketXETR    Market = "XETR"
	MarketXETRA   Market = "XETRA"
	MarketXHKG    Market = "XHKG"
	MarketXLIS    Market = "XLIS"
	MarketXLON    Market = "XLON"
	MarketXNAS    Market = "XNAS"
	MarketXNYS    Market = "XNYS"
	MarketXPAR    Market = "XPAR"
	MarketXSAU    Market = "XSAU"
)

var MarketValues = []string{"BMO", "BNYM", "BOATS", "CEUX", "CHIX", "HKEX", "IEX", "IEXG", "ISE", "LSE", "MTA", "MTAA", "NASDAQ", "NYSE", "OCEA", "OPRA", "OTC", "OTCM", "SIFMA", "TADAWUL", "XAMS", "XBRU", "XDUB", "XETR", "XETRA", "XHKG", "XLIS", "XLON", "XNAS", "XNYS", "XPAR", "XSAU"}

type Phase string

const (
	PhaseClosed Phase = "closed"
	PhaseCore   Phase = "core"
	PhaseLunch  Phase = "lunch"
	PhasePost   Phase = "post"
	PhasePre    Phase = "pre"
)

var PhaseValues = []string{"closed", "core", "lunch", "post", "pre"}

type SpOutlook string

const (
	SpOutlookDeveloping    SpOutlook = "developing"
	SpOutlookNegative      SpOutlook = "negative"
	SpOutlookNotMeaningful SpOutlook = "not_meaningful"
	SpOutlookNotRated      SpOutlook = "not_rated"
	SpOutlookPositive      SpOutlook = "positive"
	SpOutlookStable        SpOutlook = "stable"
)

var SpOutlookValues = []string{"developing", "negative", "not_meaningful", "not_rated", "positive", "stable"}

type TreasurySubtype string

const (
	TreasurySubtypeBill     TreasurySubtype = "bill"
	TreasurySubtypeBond     TreasurySubtype = "bond"
	TreasurySubtypeFloating TreasurySubtype = "floating"
	TreasurySubtypeNote     TreasurySubtype = "note"
	TreasurySubtypeStrips   TreasurySubtype = "strips"
	TreasurySubtypeTips     TreasurySubtype = "tips"
)

var TreasurySubtypeValues = []string{"bill", "bond", "floating", "note", "strips", "tips"}

type Account struct {
	AccountBlocked           bool          `json:"account_blocked,omitempty"`
	AccountNumber            string        `json:"account_number,omitempty"`
	AccruedFees              string        `json:"accrued_fees,omitempty"`
	BalanceAsof              string        `json:"balance_asof,omitempty"`
	BuyingPower              string        `json:"buying_power,omitempty"`
	Cash                     string        `json:"cash,omitempty"`
	CreatedAt                string        `json:"created_at,omitempty"`
	Currency                 string        `json:"currency,omitempty"`
	DaytradeCount            int           `json:"daytrade_count,omitempty"`
	DaytradingBuyingPower    string        `json:"daytrading_buying_power,omitempty"`
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
	PatternDayTrader         bool          `json:"pattern_day_trader,omitempty"`
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
	DTBPCheck               string `json:"dtbp_check,omitempty"`
	FractionalTrading       bool   `json:"fractional_trading,omitempty"`
	MaxMarginMultiplier     string `json:"max_margin_multiplier,omitempty"`
	MaxOptionsTradingLevel  int    `json:"max_options_trading_level,omitempty"`
	NoShorting              bool   `json:"no_shorting,omitempty"`
	PDTCheck                string `json:"pdt_check,omitempty"`
	PtpNoExceptionEntry     bool   `json:"ptp_no_exception_entry,omitempty"`
	SuspendTrade            bool   `json:"suspend_trade,omitempty"`
	TradeConfirmEmail       string `json:"trade_confirm_email,omitempty"`
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
	Attributes                   []string   `json:"attributes,omitempty"`
	Class                        AssetClass `json:"class"`
	Cusip                        string     `json:"cusip,omitempty"`
	EasyToBorrow                 bool       `json:"easy_to_borrow"`
	Exchange                     Exchange   `json:"exchange"`
	Fractionable                 bool       `json:"fractionable"`
	ID                           string     `json:"id"`
	MaintenanceMarginRequirement float64    `json:"maintenance_margin_requirement,omitempty"`
	MarginRequirementLong        string     `json:"margin_requirement_long,omitempty"`
	MarginRequirementShort       string     `json:"margin_requirement_short,omitempty"`
	Marginable                   bool       `json:"marginable"`
	Name                         string     `json:"name"`
	Shortable                    bool       `json:"shortable"`
	Status                       string     `json:"status"`
	Symbol                       string     `json:"symbol"`
	Tradable                     bool       `json:"tradable"`
}

type Calendar struct {
	Close          string `json:"close"`
	Date           string `json:"date"`
	Open           string `json:"open"`
	SettlementDate string `json:"settlement_date"`
}

type CanceledOrderResponse struct {
	ID     string `json:"id,omitempty"`
	Status int    `json:"status,omitempty"`
}

type Clock struct {
	IsOpen    bool   `json:"is_open,omitempty"`
	NextClose string `json:"next_close,omitempty"`
	NextOpen  string `json:"next_open,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

type CreateCryptoTransferRequest struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
	Asset   string `json:"asset"`
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
	Style             string              `json:"style"`
	Symbol            string              `json:"symbol"`
	Tradable          bool                `json:"tradable"`
	Type              string              `json:"type"`
	UnderlyingAssetID string              `json:"underlying_asset_id"`
	UnderlyingSymbol  string              `json:"underlying_symbol"`
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
	CanceledAt     string         `json:"canceled_at,omitempty"`
	ClientOrderID  string         `json:"client_order_id,omitempty"`
	CreatedAt      string         `json:"created_at,omitempty"`
	ExpiredAt      string         `json:"expired_at,omitempty"`
	ExtendedHours  bool           `json:"extended_hours,omitempty"`
	FailedAt       string         `json:"failed_at,omitempty"`
	FilledAt       string         `json:"filled_at,omitempty"`
	FilledAvgPrice string         `json:"filled_avg_price,omitempty"`
	FilledQty      string         `json:"filled_qty,omitempty"`
	Hwm            string         `json:"hwm,omitempty"`
	ID             string         `json:"id,omitempty"`
	Legs           []OrderLeg     `json:"legs,omitempty"`
	LimitPrice     string         `json:"limit_price,omitempty"`
	Notional       string         `json:"notional"`
	OrderClass     OrderClass     `json:"order_class,omitempty"`
	OrderType      string         `json:"order_type,omitempty"`
	PositionIntent PositionIntent `json:"position_intent,omitempty"`
	Qty            string         `json:"qty,omitempty"`
	ReplacedAt     string         `json:"replaced_at,omitempty"`
	ReplacedBy     string         `json:"replaced_by,omitempty"`
	Replaces       string         `json:"replaces,omitempty"`
	Side           OrderSide      `json:"side,omitempty"`
	Status         OrderStatus    `json:"status,omitempty"`
	StopPrice      string         `json:"stop_price,omitempty"`
	SubmittedAt    string         `json:"submitted_at,omitempty"`
	Symbol         string         `json:"symbol,omitempty"`
	TimeInForce    TimeInForce    `json:"time_in_force"`
	TrailPercent   string         `json:"trail_percent,omitempty"`
	TrailPrice     string         `json:"trail_price,omitempty"`
	Type           OrderType      `json:"type"`
	UpdatedAt      string         `json:"updated_at,omitempty"`
}

type OrderLeg struct {
	AssetClass     AssetClass     `json:"asset_class,omitempty"`
	AssetID        string         `json:"asset_id,omitempty"`
	CanceledAt     string         `json:"canceled_at,omitempty"`
	ClientOrderID  string         `json:"client_order_id,omitempty"`
	CreatedAt      string         `json:"created_at,omitempty"`
	ExpiredAt      string         `json:"expired_at,omitempty"`
	ExtendedHours  bool           `json:"extended_hours,omitempty"`
	FailedAt       string         `json:"failed_at,omitempty"`
	FilledAt       string         `json:"filled_at,omitempty"`
	FilledAvgPrice string         `json:"filled_avg_price,omitempty"`
	FilledQty      string         `json:"filled_qty,omitempty"`
	Hwm            string         `json:"hwm,omitempty"`
	ID             string         `json:"id,omitempty"`
	Legs           []any          `json:"legs,omitempty"`
	LimitPrice     string         `json:"limit_price,omitempty"`
	Notional       string         `json:"notional"`
	OrderClass     OrderClass     `json:"order_class,omitempty"`
	OrderType      string         `json:"order_type,omitempty"`
	PositionIntent PositionIntent `json:"position_intent,omitempty"`
	Qty            string         `json:"qty"`
	ReplacedAt     string         `json:"replaced_at,omitempty"`
	ReplacedBy     string         `json:"replaced_by,omitempty"`
	Replaces       string         `json:"replaces,omitempty"`
	Side           OrderSide      `json:"side"`
	Status         OrderStatus    `json:"status,omitempty"`
	StopPrice      string         `json:"stop_price,omitempty"`
	SubmittedAt    string         `json:"submitted_at,omitempty"`
	Symbol         string         `json:"symbol"`
	TimeInForce    TimeInForce    `json:"time_in_force"`
	TrailPercent   string         `json:"trail_percent,omitempty"`
	TrailPrice     string         `json:"trail_price,omitempty"`
	Type           OrderType      `json:"type"`
	UpdatedAt      string         `json:"updated_at,omitempty"`
}

type PatchOrderRequest struct {
	AdvancedInstructions AdvancedInstructions `json:"advanced_instructions,omitempty"`
	ClientOrderID        string               `json:"client_order_id,omitempty"`
	LimitPrice           string               `json:"limit_price,omitempty"`
	Qty                  string               `json:"qty,omitempty"`
	StopPrice            string               `json:"stop_price,omitempty"`
	TimeInForce          TimeInForce          `json:"time_in_force,omitempty"`
	Trail                string               `json:"trail,omitempty"`
}

type PortfolioHistory struct {
	BaseValue     float64        `json:"base_value"`
	BaseValueAsof string         `json:"base_value_asof,omitempty"`
	Cashflow      map[string]any `json:"cashflow,omitempty"`
	Equity        []float64      `json:"equity"`
	ProfitLoss    []float64      `json:"profit_loss"`
	ProfitLossPct []float64      `json:"profit_loss_pct"`
	Timeframe     string         `json:"timeframe"`
	Timestamp     []int          `json:"timestamp"`
}

type Position struct {
	AssetClass             AssetClass          `json:"asset_class"`
	AssetID                string              `json:"asset_id"`
	AssetMarginable        bool                `json:"asset_marginable"`
	AvgEntryPrice          string              `json:"avg_entry_price"`
	ChangeToday            string              `json:"change_today"`
	CostBasis              string              `json:"cost_basis"`
	CurrentPrice           string              `json:"current_price"`
	Exchange               ExchangeForPosition `json:"exchange"`
	LastdayPrice           string              `json:"lastday_price"`
	MarketValue            string              `json:"market_value"`
	Qty                    string              `json:"qty"`
	QtyAvailable           string              `json:"qty_available,omitempty"`
	Side                   string              `json:"side"`
	Symbol                 string              `json:"symbol"`
	UnrealizedIntradayPL   string              `json:"unrealized_intraday_pl"`
	UnrealizedIntradayPlpc string              `json:"unrealized_intraday_plpc"`
	UnrealizedPL           string              `json:"unrealized_pl"`
	UnrealizedPlpc         string              `json:"unrealized_plpc"`
}

type PositionClosedReponse struct {
	Body   Order  `json:"body,omitempty"`
	Status int    `json:"status"`
	Symbol string `json:"symbol"`
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

type UpdateWatchlistRequest struct {
	Name    string   `json:"name"`
	Symbols []string `json:"symbols,omitempty"`
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

type ClockV3 struct {
	IsMarketDay     bool         `json:"is_market_day"`
	Market          PublicMarket `json:"market"`
	NextMarketClose string       `json:"next_market_close"`
	NextMarketOpen  string       `json:"next_market_open"`
	Phase           Phase        `json:"phase"`
	PhaseUntil      string       `json:"phase_until"`
	Timestamp       string       `json:"timestamp"`
}

type ClockResp struct {
	Clocks []ClockV3 `json:"clocks"`
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

type UsCorporate struct {
	AccruedInterest                 float64         `json:"accrued_interest,omitempty"`
	BondStatus                      BondStatus      `json:"bond_status"`
	CallType                        CallType        `json:"call_type,omitempty"`
	Callable                        bool            `json:"callable"`
	ClosePrice                      float64         `json:"close_price,omitempty"`
	ClosePriceDate                  string          `json:"close_price_date,omitempty"`
	CloseYieldToMaturity            float64         `json:"close_yield_to_maturity,omitempty"`
	CloseYieldToWorst               float64         `json:"close_yield_to_worst,omitempty"`
	Convertible                     bool            `json:"convertible"`
	CountryDomicile                 string          `json:"country_domicile"`
	Coupon                          float64         `json:"coupon"`
	CouponFrequency                 CouponFrequency `json:"coupon_frequency"`
	CouponType                      CouponType      `json:"coupon_type"`
	Cusip                           string          `json:"cusip"`
	DatedDate                       string          `json:"dated_date"`
	DayCount                        DayCount        `json:"day_count"`
	Description                     string          `json:"description"`
	DescriptionShort                string          `json:"description_short"`
	FirstCouponDate                 string          `json:"first_coupon_date,omitempty"`
	Isin                            string          `json:"isin"`
	IssueDate                       string          `json:"issue_date"`
	IssueMinimumDenomination        float64         `json:"issue_minimum_denomination"`
	IssuePrice                      float64         `json:"issue_price"`
	IssueSize                       float64         `json:"issue_size"`
	Issuer                          string          `json:"issuer"`
	LastCouponDate                  string          `json:"last_coupon_date,omitempty"`
	LiquidityInstitutionalAggregate float64         `json:"liquidity_institutional_aggregate,omitempty"`
	LiquidityInstitutionalBuy       float64         `json:"liquidity_institutional_buy,omitempty"`
	LiquidityInstitutionalSell      float64         `json:"liquidity_institutional_sell,omitempty"`
	LiquidityMicroAggregate         float64         `json:"liquidity_micro_aggregate,omitempty"`
	LiquidityMicroBuy               float64         `json:"liquidity_micro_buy,omitempty"`
	LiquidityMicroSell              float64         `json:"liquidity_micro_sell,omitempty"`
	LiquidityRetailAggregate        float64         `json:"liquidity_retail_aggregate,omitempty"`
	LiquidityRetailBuy              float64         `json:"liquidity_retail_buy,omitempty"`
	LiquidityRetailSell             float64         `json:"liquidity_retail_sell,omitempty"`
	Marginable                      bool            `json:"marginable"`
	MaturityDate                    string          `json:"maturity_date,omitempty"`
	NextCallDate                    string          `json:"next_call_date,omitempty"`
	NextCallPrice                   float64         `json:"next_call_price,omitempty"`
	NextCouponDate                  string          `json:"next_coupon_date,omitempty"`
	ParValue                        float64         `json:"par_value"`
	Perpetual                       bool            `json:"perpetual"`
	Puttable                        bool            `json:"puttable"`
	RegS                            bool            `json:"reg_s"`
	ReissueDate                     string          `json:"reissue_date,omitempty"`
	ReissuePrice                    float64         `json:"reissue_price,omitempty"`
	ReissueSize                     float64         `json:"reissue_size,omitempty"`
	Sector                          string          `json:"sector"`
	Seniority                       string          `json:"seniority"`
	SpCreditwatch                   string          `json:"sp_creditwatch,omitempty"`
	SpCreditwatchDate               string          `json:"sp_creditwatch_date,omitempty"`
	SpOutlook                       SpOutlook       `json:"sp_outlook,omitempty"`
	SpOutlookDate                   string          `json:"sp_outlook_date,omitempty"`
	SpRating                        string          `json:"sp_rating,omitempty"`
	SpRatingDate                    string          `json:"sp_rating_date,omitempty"`
	Ticker                          string          `json:"ticker"`
	Tradable                        bool            `json:"tradable"`
}

type UsCorporatesResp struct {
	UsCorporates []UsCorporate `json:"us_corporates"`
}

type UsTreasuriesResp struct {
	UsTreasuries []UsTreasury `json:"us_treasuries"`
}

type UsTreasury struct {
	BondStatus           BondStatus      `json:"bond_status"`
	ClosePrice           float64         `json:"close_price,omitempty"`
	ClosePriceDate       string          `json:"close_price_date,omitempty"`
	CloseYieldToMaturity float64         `json:"close_yield_to_maturity,omitempty"`
	CloseYieldToWorst    float64         `json:"close_yield_to_worst,omitempty"`
	Coupon               float64         `json:"coupon"`
	CouponFrequency      CouponFrequency `json:"coupon_frequency"`
	CouponType           CouponType      `json:"coupon_type"`
	Cusip                string          `json:"cusip"`
	Description          string          `json:"description"`
	DescriptionShort     string          `json:"description_short"`
	FirstCouponDate      string          `json:"first_coupon_date,omitempty"`
	Isin                 string          `json:"isin"`
	IssueDate            string          `json:"issue_date"`
	LastCouponDate       string          `json:"last_coupon_date,omitempty"`
	MaturityDate         string          `json:"maturity_date"`
	NextCouponDate       string          `json:"next_coupon_date,omitempty"`
	Subtype              TreasurySubtype `json:"subtype"`
	Tradable             bool            `json:"tradable"`
}

var AccountOptionsApprovedLevelValues = []string{"0", "1", "2", "3"}

var AccountOptionsTradingLevelValues = []string{"0", "1", "2", "3"}

var AccountConfigurationsDTBPCheckValues = []string{"both", "entry", "exit"}

var AccountConfigurationsMaxOptionsTradingLevelValues = []string{"0", "1", "2", "3"}

var AdvancedInstructionsAlgorithmValues = []string{"DMA", "TWAP", "VWAP"}

var AdvancedInstructionsDestinationValues = []string{"ARCA", "NASDAQ", "NYSE"}

var AssetsStatusValues = []string{"active", "inactive"}

var NonTradeActivitiesStatusValues = []string{"canceled", "correct", "executed"}

var OptionContractStatusValues = []string{"active", "inactive"}

var OptionContractStyleValues = []string{"american", "european"}

var OptionContractTypeValues = []string{"call", "put"}

var OptionDeliverableSettlementMethodValues = []string{"BTOB", "CADF", "CAFX", "CCC"}

var OptionDeliverableSettlementTypeValues = []string{"T+0", "T+1", "T+2", "T+3", "T+4", "T+5"}

var OptionDeliverableTypeValues = []string{"cash", "equity"}

var TradingActivitiesTypeValues = []string{"fill", "partial_fill"}

var WhitelistedAddressStatusValues = []string{"APPROVED", "PENDING"}
