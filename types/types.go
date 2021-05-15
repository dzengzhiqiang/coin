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

type SpotBalances struct {
	Balances []SpotBalance `json:"balances" db:"balances"`
}

type SpotTradeRequest struct {
	Symbol   string  `json:"symbol" db:"symbol"`
	Quantity float64 `json:"quantity" db:"quantity"`
}

type SpotTradeResult struct {
}

type SpotTradeResponse struct {
	State  TradeState
	Result SpotTradeResult
}

func NewSymbol(strBase, strQuote string) string {
	return fmt.Sprintf("%s/%s", strBase, strQuote)
}
