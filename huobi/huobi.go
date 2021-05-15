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
func (m *HuoBi) SpotBalances() (balances *types.SpotBalances, code types.BizCode) {
	log.Debugf("query spot balance")
	return
}

//query spot coin price
func (m *HuoBi) SpotPrice(symbol string) (price *types.CoinPrice, code types.BizCode) {
	log.Debugf("query spot price")

	return
}

//trade spot coin
func (m *HuoBi) SpotTrade(trade types.SpotTradeRequest) (response *types.SpotTradeResponse, code types.BizCode) {
	log.Debugf("spot trade")
	return
}
