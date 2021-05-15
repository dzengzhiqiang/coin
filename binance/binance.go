package binance

import (
	"coin/api"
	"coin/config"
	"coin/types"
	"coin/utils"
	"encoding/json"
	"fmt"
	"github.com/civet148/httpc"
	"github.com/civet148/log"
	"net/url"
)

/*
	接口的基本信息
	GET 方法的接口, 参数必须在 query string中发送。
	POST, PUT, 和 DELETE 方法的接口,参数可以在内容形式为application/x-www-form-urlencoded的 query string 中发送，也可以在 request body 中发送。 如果你喜欢，也可以混合这两种方式发送参数。
	对参数的顺序不做要求。
	但如果同一个参数名在query string和request body中都有，query string中的会被优先采用。
*/

const (
	binance_api_host = "https://api.binance.com"
)
const (
	api_v3_account = "/api/v3/account"
)

const (
	httpCodeWebAccessFirewall = 403 //违反WAF限制(Web应用程序防火墙)
	httpCodeBanned            = 418 //IP封禁
	httpCodeRateLimit         = 429 //访问频率限制

)

const (
	X_MBX_APIKEY = "X-MBX-APIKEY"
)

type Binance struct {
	cfg    *config.Config
	sign   *Signature
	client *httpc.Client
}

func NewCoinManager(cfg *config.Config) api.CoinApi {
	return &Binance{
		cfg:    cfg,
		sign:   NewSignature(cfg),
		client: httpc.NewHttpClient(15),
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
func (m *Binance) SpotBalances() (balances *types.SpotBalances, code types.BizCode) {
	log.Debugf("query spot balance")
	m.client.Header().Set(X_MBX_APIKEY, m.cfg.AppKey)

	strQuery, strSign := m.sign.Sign(url.Values{
		paramTimestamp: []string{utils.UnixMilliSecond().String()},
	})
	strUrl := m.makeUrl(binance_api_host, api_v3_account, strQuery, strSign)
	log.Infof("GET %s", strUrl)
	r, err := m.client.Get(strUrl, nil)
	if err != nil {
		log.Errorf("GET [%s] error [%s]", api_v3_account, err.Error())
		return nil, types.BizCode_ServerError
	}
	log.Debugf("body [%s]", r.Body)
	var sa spotAccount
	if err = json.Unmarshal(r.Body, &sa); err != nil {
		log.Errorf("unmarshal account error [%s]", err)
		return nil, types.BizCode_UnmashalError
	}
	return
}

//query spot coin price
func (m *Binance) SpotPrice(cs *types.CoinSymbol) (price *types.CoinPrice, code types.BizCode) {
	log.Debugf("query spot price")

	return
}

//trade spot coin
func (m *Binance) SpotTrade(trade types.SpotTradeRequest) (response *types.SpotTradeResponse, code types.BizCode) {
	log.Debugf("spot trade")
	return
}

func (m *Binance) makeUrl(strHost, strApi, strQuery, strSign string) string {
	return fmt.Sprintf("%s%s?%s&%s=%s", strHost, strApi, strQuery, paramSignature, strSign)
}
