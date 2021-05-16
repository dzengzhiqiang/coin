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
	SpotBalances() (*types.SpotAccount, types.BizCode)

	//query spot coin price
	SpotPrice(symbol string) (*types.CoinPrice, types.BizCode)

	//buy spot coin
	SpotBuy(req *types.SpotTradeReq) (*types.SpotTradeResp, types.BizCode)

	//sell spot coin
	SpotSell(req *types.SpotTradeReq) (*types.SpotTradeResp, types.BizCode)
}
