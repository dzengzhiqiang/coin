package binance

import (
	"coin/api"
	"coin/config"
	"coin/dal"
	"coin/proto"
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
	api_v3_account      = "/api/v3/account"      //现货账户
	api_v3_order        = "/api/v3/order"        //下单
	api_v3_avg_price    = "/api/v3/avgPrice"     //平均价格
	api_v3_ticker_price = "/api/v3/ticker/price" //最新价格
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
	dal    *dal.Dal
}

func NewCoinManager(cfg *config.Config) api.CoinApi {
	return &Binance{
		cfg:    cfg,
		dal:    dal.NewDal(cfg),
		sign:   NewSignature(cfg),
		client: httpc.NewHttpClient(60),
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
func (m *Binance) SpotBalances() (acc *proto.SpotAccount, code types.BizCode) {
	log.Debugf("query spot balance")
	m.client.Header().Set(X_MBX_APIKEY, m.cfg.AppKey)

	strQuery, strSign := m.sign.Sign(url.Values{
		paramTimestamp: []string{utils.UnixMilliSecond().String()},
	})
	strUrl := m.makeSignUrl(binance_api_host, api_v3_account, strQuery, strSign)
	log.Infof("GET %s", strUrl)
	r, err := m.client.Get(strUrl, nil)
	if err != nil {
		log.Errorf("GET [%s] error [%s]", strUrl, err.Error())
		return nil, types.BizCode_ServerError
	}
	log.Infof("http code [%v] body [%s]", r.StatusCode, r.Body)
	var sa spotAccountResp
	if err = json.Unmarshal(r.Body, &sa); err != nil {
		log.Errorf("unmarshal account error [%s]", err)
		return nil, types.BizCode_UnmashalError
	}
	acc = &proto.SpotAccount{
		MakerCommission:  sa.MakerCommission,
		TakerCommission:  sa.TakerCommission,
		BuyerCommission:  sa.BuyerCommission,
		SellerCommission: sa.SellerCommission,
		CanTrade:         sa.CanTrade,
		CanWithdraw:      sa.CanWithdraw,
		CanDeposit:       sa.CanDeposit,
		UpdateTime:       sa.UpdateTime,
		AccountType:      sa.AccountType,
		Balances:         sa.Balances,
		Permissions:      sa.Permissions,
	}
	return
}

//query spot coin price
func (m *Binance) SpotPrice(symbol string) (price *proto.CoinPrice, code types.BizCode) {
	log.Debugf("query spot price [%s]", symbol)
	m.client.Header().Set(X_MBX_APIKEY, m.cfg.AppKey)
	strUrl := m.makeValuesUrl(binance_api_host, api_v3_ticker_price, url.Values{
		paramSymbol: []string{symbol},
	})
	log.Infof("GET %s", strUrl)
	r, err := m.client.Get(strUrl, nil)
	if err != nil {
		log.Errorf("GET [%s] error [%s]", strUrl, err.Error())
		return nil, types.BizCode_ServerError
	}
	log.Infof("http code [%v] body [%s]", r.StatusCode, r.Body)
	price = &proto.CoinPrice{}
	if err = json.Unmarshal(r.Body, price); err != nil {
		log.Errorf("unmarshal account error [%s]", err)
		return nil, types.BizCode_UnmashalError
	}
	return
}

//buy spot coin
func (m *Binance) SpotBuy(req *proto.SpotTradeReq) (*proto.SpotTradeResp, types.BizCode) {
	resp, code := m.spotTrade(req, true)
	if code != types.BizCode_OK {
		log.Errorf("request [%+v] failed code [%v]", req, code)
		return nil, types.BizCode_ServerError
	}
	return resp, types.BizCode_OK
}

//sell spot coin
func (m *Binance) SpotSell(req *proto.SpotTradeReq) (*proto.SpotTradeResp, types.BizCode) {
	resp, code := m.spotTrade(req, false)
	if code != types.BizCode_OK {
		log.Errorf("request [%+v] failed code [%v]", req, code)
		return nil, types.BizCode_ServerError
	}
	return resp, types.BizCode_OK
}

func (m *Binance) makeHostUrl(strHost, strApi string) string {
	return fmt.Sprintf("%s%s", strHost, strApi)
}

func (m *Binance) makeValuesUrl(strHost, strApi string, values url.Values) string {
	strQuery := values.Encode()
	return m.makeQueryUrl(strHost, strApi, strQuery)
}

func (m *Binance) makeQueryUrl(strHost, strApi, strQuery string) string {
	if strQuery != "" {
		return fmt.Sprintf("%s%s?%s", strHost, strApi, strQuery)
	}
	return fmt.Sprintf("%s%s", strHost, strApi)
}

func (m *Binance) makeSignUrl(strHost, strApi, strQuery, strSign string) string {
	return fmt.Sprintf("%s%s?%s&%s=%s", strHost, strApi, strQuery, paramSignature, strSign)
}

func (m *Binance) makeQuerySign(strQuery, strSign string) string {
	return fmt.Sprintf("%s&%s=%s", strQuery, paramSignature, strSign)
}

func (m *Binance) spotTrade(req *proto.SpotTradeReq, buy bool) (resp *proto.SpotTradeResp, code types.BizCode) {
	strSide := "SELL"
	if buy {
		strSide = "BUY"
	}
	m.client.Header().Set(X_MBX_APIKEY, m.cfg.AppKey)

	strQuery, strSign := m.sign.Sign(url.Values{
		paramSymbol:   []string{req.Symbol},
		paramSide:     []string{strSide},
		paramType:     []string{"MARKET"},
		paramQuantity: []string{req.Quantity.String()},
		//paramQuoteOrderQty: []string{req.Quantity.String()},
		//paramPrice: []string{""},
		//paramNewClientOrderId: []string{""},
		//paramStopPrice: []string{""},
		//paramIcebergQty: []string{""},
		//paramTimeInForce: []string{"GTC"},
		//paramNewOrderRespType: []string{"FULL"},
		//paramRecvWindow: []string{"5000"},
		paramTimestamp: []string{utils.UnixMilliSecond().String()},
	})
	strUrl := m.makeHostUrl(binance_api_host, api_v3_order)
	strPost := m.makeQuerySign(strQuery, strSign)
	log.Infof("POST %s [%s]", strUrl, strPost)
	r, err := m.client.Post(httpc.CONTENT_TYPE_NAME_X_WWW_FORM_URL_ENCODED, strUrl, strPost)
	if err != nil {
		log.Errorf("POST [%s] error [%s]", strUrl, err.Error())
		return nil, types.BizCode_ServerError
	}
	log.Infof("http code [%v] body [%s]", r.StatusCode, r.Body)
	var sor spotOrderResp
	if err = json.Unmarshal(r.Body, &sor); err != nil {
		log.Errorf("unmarshal account error [%s]", err)
		return nil, types.BizCode_UnmashalError
	}
	return &proto.SpotTradeResp{
		State: types.TradeState_OK,
		Result: proto.SpotTradeResult{
			Symbol:              sor.Symbol,
			OrderId:             utils.Int64ToString(sor.OrderId),
			ClientOrderId:       sor.ClientOrderId,
			TransactTime:        sor.TransactTime,
			Price:               sor.Price,
			OrigQty:             sor.OrigQty,
			ExecutedQty:         sor.ExecutedQty,
			CummulativeQuoteQty: sor.CummulativeQuoteQty,
			Type:                sor.Type,
			Side:                sor.Side,
		},
	}, types.BizCode_OK
}
