package binance

import (
	"coin/proto"
	"github.com/civet148/sqlca"
)

const (
	paramRecvWindow       = "recvWindow"
	paramTimestamp        = "timestamp"
	paramSignature        = "signature"
	paramSymbol           = "symbol"
	paramSide             = "side"
	paramType             = "type"
	paramTimeInForce      = "timeInForce"
	paramQuantity         = "quantity"
	paramQuoteOrderQty    = "quoteOrderQty"
	paramPrice            = "price"
	paramNewClientOrderId = "newClientOrderId"
	paramStopPrice        = "stopPrice"
	paramIcebergQty       = "icebergQty"
	paramNewOrderRespType = "newOrderRespType"
)

type spotAccountResp struct {
	MakerCommission  int32               `json:"makerCommission"`
	TakerCommission  int32               `json:"takerCommission"`
	BuyerCommission  int32               `json:"buyerCommission"`
	SellerCommission int32               `json:"sellerCommission"`
	CanTrade         bool                `json:"canTrade"`
	CanWithdraw      bool                `json:"canWithdraw"`
	CanDeposit       bool                `json:"canDeposit"`
	UpdateTime       int64               `json:"updateTime"`
	AccountType      string              `json:"accountType"`
	Balances         []proto.SpotBalance `json:"balances"`
	Permissions      []string            `json:"permissions"`
}

//https://2pd.github.io/zh/spot/restful_endpoints/spot_endpoints.html#%E4%B8%8B%E5%8D%95-trade
type spotOrderReq struct {
	Symbol           string        `json:"symbol"`           //STRING	YES
	Side             string        `json:"side"`             //ENUM		YES
	Type             string        `json:"type"`             //ENUM		YES
	TimeInForce      string        `json:"timeInForce"`      //ENUM		NO
	Quantity         sqlca.Decimal `json:"quantity"`         //DECIMAL	NO
	QuoteOrderQty    sqlca.Decimal `json:"quoteOrderQty"`    //DECIMAL	NO
	Price            sqlca.Decimal `json:"price"`            //DECIMAL	NO
	NewClientOrderId string        `json:"newClientOrderId"` //STRING	NO	客户自定义的唯一订单ID。 如果未发送，则自动生成
	StopPrice        sqlca.Decimal `json:"stopPrice"`        //DECIMAL	NO	仅 STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, 和TAKE_PROFIT_LIMIT 需要此参数。
	IcebergQty       sqlca.Decimal `json:"icebergQty"`       //DECIMAL	NO	仅使用 LIMIT, STOP_LOSS_LIMIT, 和 TAKE_PROFIT_LIMIT 创建新的 iceberg 订单时需要此参数
	NewOrderRespType string        `json:"newOrderRespType"` //ENUM	NO	设置响应JSON。 ACK，RESULT或FULL； "MARKET"和" LIMIT"订单类型默认为"FULL"，所有其他订单默认为"ACK"。
	RecvWindow       string        `json:"recvWindow"`       //LONG	NO	赋值不能大于 60000
	Timestamp        int           `json:"timestamp"`        //LONG	YES
}

type spotOrderResp struct {
	Symbol              string        `json:"symbol"`              // 交易对
	OrderId             int64         `json:"orderId"`             // 系统的订单ID
	OrderListId         int64         `json:"orderListId"`         // OCO订单的ID，不然就是-1
	ClientOrderId       string        `json:"clientOrderId"`       // 客户自己设置的ID
	TransactTime        int64         `json:"transactTime"`        // 交易的时间戳
	Price               sqlca.Decimal `json:"price"`               // 订单价格
	OrigQty             sqlca.Decimal `json:"origQty"`             // 用户设置的原始订单数量
	ExecutedQty         sqlca.Decimal `json:"executedQty"`         // 交易的订单数量
	CummulativeQuoteQty sqlca.Decimal `json:"cummulativeQuoteQty"` // 累计交易的金额
	Status              string        `json:"status"`              // 订单状态(FILLED)
	TimeInForce         string        `json:"timeInForce"`         // 订单的时效方式(GTC)
	Type                string        `json:"type"`                // 订单类型， 比如市价单，现价单等
	Side                string        `json:"side"`                // 订单方向，买还是卖(SELL,BUY)
	Fills               []spotFill    `json:"fills"`               // 订单中交易的信息
}

type spotFill struct {
	Price           sqlca.Decimal `json:"price"`           // 交易的价格
	Qty             sqlca.Decimal `json:"qty"`             // 交易的数量
	Commission      sqlca.Decimal `json:"commission"`      // 手续费金额
	CommissionAsset string        `json:"commissionAsset"` // 手续费的币种
}
