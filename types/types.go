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

type CoinSymbol struct {
	Symbol string `json:"symbol" db:"symbol"`
	Base   string `json:"base" db:"base"`
	Quote  string `json:"quote" db:"quote"`
}

type CoinPrice struct {
	CoinSymbol *CoinSymbol   `json:"coin_symbol" db:"coin_symbol"`
	Price      sqlca.Decimal `json:"price" db:"price"`
	PriceTime  int64         `json:"price_time" db:"price_time"`
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
	CoinSymbol *CoinSymbol
	Quantity   float64
}

type SpotTradeResult struct {
}

type SpotTradeResponse struct {
	State  TradeState
	Result SpotTradeResult
}

func NewCoinSymbol(strBase, strQuote string) *CoinSymbol {
	return &CoinSymbol{
		Symbol: fmt.Sprintf("%s/%s", strBase, strQuote),
		Base:   strBase,
		Quote:  strQuote,
	}
}
