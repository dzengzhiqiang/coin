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
func (m *HuoBi) SpotBalances() *types.SpotBalances {
	log.Debugf("query spot balance")
	return nil
}

//query spot coin price
func (m *HuoBi) SpotPrice(cs *types.CoinSymbol) *types.CoinPrice {
	log.Debugf("query spot price")
	return nil
}

//trade spot coin
func (m *HuoBi) SpotTrade(trade types.SpotTradeRequest) *types.SpotTradeResponse {
	log.Debugf("spot trade")
	return nil
}
