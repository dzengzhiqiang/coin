package api

import (
	"coin/types"
)

type CoinApi interface {
	//run service loop
	Run() (err error)

	//close service
	Close()

	//query spot account balance
	SpotBalances() *types.SpotBalances

	//query spot coin price
	SpotPrice(cs *types.CoinSymbol) *types.CoinPrice

	//trade spot coin
	SpotTrade(trade types.SpotTradeRequest) *types.SpotTradeResponse
}
