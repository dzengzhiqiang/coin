package proto

import (
	"coin/types"
	"github.com/civet148/sqlca"
)

type CoinPrice struct {
	Symbol string        `json:"symbol" db:"symbol"`
	Price  sqlca.Decimal `json:"price" db:"price"`
}

type SpotBalance struct {
	Asset  string        `json:"asset" db:"asset"`
	Free   sqlca.Decimal `json:"free" db:"free"`
	Locked sqlca.Decimal `json:"locked" db:"locked"`
}

type SpotAccount struct {
	MakerCommission  int32         `json:"maker_commission"`
	TakerCommission  int32         `json:"taker_commission"`
	BuyerCommission  int32         `json:"buyer_commission"`
	SellerCommission int32         `json:"seller_commission"`
	CanTrade         bool          `json:"can_trade"`
	CanWithdraw      bool          `json:"can_withdraw"`
	CanDeposit       bool          `json:"can_deposit"`
	UpdateTime       int64         `json:"update_time"`
	AccountType      string        `json:"account_type"`
	Balances         []SpotBalance `json:"balances"`
	Permissions      []string      `json:"permissions"`
}

type SpotTradeReq struct {
	Symbol   string        `json:"symbol"`
	Quantity sqlca.Decimal `json:"quantity"`
}

type SpotTradeResult struct {
	Symbol              string        `json:"symbol"`                     // 交易对
	OrderId             string        `json:"order_id"`                   // 系统的订单ID
	ClientOrderId       string        `json:"client_order_id"`            // 客户自己设置的ID
	TransactTime        int64         `json:"transact_time"`              // 交易的时间戳
	Price               sqlca.Decimal `json:"price"`                      // 订单价格
	OrigQty             sqlca.Decimal `json:"orig_quantity"`              // 用户设置的原始订单数量
	ExecutedQty         sqlca.Decimal `json:"executed_quantity"`          // 交易的订单数量
	CummulativeQuoteQty sqlca.Decimal `json:"cummulative_quote_quantity"` // 累计交易的金额
	Type                string        `json:"type"`                       // 订单类型， 比如市价单，现价单等
	Side                string        `json:"side"`                       // 订单方向，买还是卖(SELL,BUY)
}

type SpotTradeResp struct {
	State  types.TradeState
	Result SpotTradeResult
}
