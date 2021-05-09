package binance

import (
	"coin/api"
	"coin/config"
	"coin/types"
	"github.com/civet148/log"
)

type Binance struct {
	cfg *config.Config
}

func NewCoinManager(cfg *config.Config) api.CoinApi {
	return &Binance{
		cfg: cfg,
	}
}

//run service loop
func (m *Binance) Run() (err error) {
	log.Debugf("running...")
	return
}

//close service
func (m *Binance) Close() {
	log.Debugf("closing...")
}

//query spot account balance
func (m *Binance) SpotBalances() *types.SpotBalances {
	log.Debugf("query spot balance")
	return nil
}

//query spot coin price
func (m *Binance) SpotPrice(cs *types.CoinSymbol) *types.CoinPrice {
	log.Debugf("query spot price")
	return nil
}

//trade spot coin
func (m *Binance) SpotTrade(trade types.SpotTradeRequest) *types.SpotTradeResponse {
	log.Debugf("spot trade")
	return nil
}
