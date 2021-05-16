package huobi

import (
	"coin/api"
	"coin/config"
	"coin/types"
	"github.com/civet148/log"
)

type HuoBi struct {
	cfg *config.Config
}

func NewCoinManager(cfg *config.Config) api.CoinApi {
	return &HuoBi{
		cfg: cfg,
	}
}

//run service loop
func (m *HuoBi) Run() (err error) {
	log.Debugf("running...")
	return
}

//close service
func (m *HuoBi) Close() {
	log.Debugf("closing...")
}

//query spot account balance
func (m *HuoBi) SpotBalances() (acc *types.SpotAccount, code types.BizCode) {
	log.Debugf("query spot balance")
	return
}

//query spot coin price
func (m *HuoBi) SpotPrice(symbol string) (price *types.CoinPrice, code types.BizCode) {
	log.Debugf("query spot price")

	return
}

//buy spot coin
func (m *HuoBi) SpotBuy(req *types.SpotTradeReq) (*types.SpotTradeResp, types.BizCode) {
	return nil, 0
}

//sell spot coin
func (m *HuoBi) SpotSell(req *types.SpotTradeReq) (*types.SpotTradeResp, types.BizCode) {
	return nil, 0
}
