package types

import (
	"fmt"
	"strings"
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

type Symbol string

//one input:  DOGE/USDT
//tow inputs: DOGE USDT
func NewSymbol(args ...string) Symbol {
	var s Symbol
	if len(args) == 1 {
		s = Symbol(args[0])
	} else if len(args) >= 2 {
		strSymbol := fmt.Sprintf("%s/%s", args[0], args[1])
		s = Symbol(strSymbol)
	}
	return s
}

func (s Symbol) String() string {
	return string(s)
}

func (s Symbol) GoString() string {
	return string(s)
}

func (s Symbol) Base() (strBase string) {
	bq := strings.Split(string(s), "/")
	if len(bq) == 2 {
		strBase = bq[0]
	}
	return
}

func (s Symbol) Quote() (strQuote string) {
	bq := strings.Split(string(s), "/")
	if len(bq) == 2 {
		strQuote = bq[1]
	}
	return
}

func (s Symbol) FromString(strSymbol string) Symbol {
	return Symbol(strSymbol)
}
