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
	SpotBalances() (*types.SpotBalances, types.BizCode)

	//query spot coin price
	SpotPrice(symbol string) (*types.CoinPrice, types.BizCode)

	//trade spot coin
	SpotTrade(trade types.SpotTradeRequest) (*types.SpotTradeResponse, types.BizCode)
}
