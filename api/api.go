package api

import (
	"coin/proto"
	"coin/types"
)

type CoinApi interface {
	//run service loop
	Run() (err error)

	//close service
	Close()

	//query spot account balance
	SpotBalances() (*proto.SpotAccount, types.BizCode)

	//query spot coin price
	SpotPrice(symbol string) (*proto.CoinPrice, types.BizCode)

	//buy spot coin
	SpotBuy(req *proto.SpotTradeReq) (*proto.SpotTradeResp, types.BizCode)

	//sell spot coin
	SpotSell(req *proto.SpotTradeReq) (*proto.SpotTradeResp, types.BizCode)
}
