package types

import (
	"fmt"
	"github.com/civet148/sqlca"
)

type TradeState int

const (
	TradeState_Failed  TradeState = -1
	TradeState_OK      TradeState = 0
	TradeState_Pending TradeState = 1
)

func (t TradeState) String() string {
	switch t {
	case TradeState_Failed:
		return "TradeState_Failed"
	case TradeState_OK:
		return "TradeState_OK"
	case TradeState_Pending:
		return "TradeState_Pending"
	}
	return fmt.Sprintf("TradeState_Unknown<%d>", t)
}

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
}

type SpotTradeResp struct {
	State  TradeState
	Result SpotTradeResult
}

func NewSymbol(strBase, strQuote string) string {
	return fmt.Sprintf("%s/%s", strBase, strQuote)
}
