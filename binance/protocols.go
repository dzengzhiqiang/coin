package binance

import "coin/types"

const (
	paramRecvWindow = "recvWindow"
	paramTimestamp  = "timestamp"
	paramSignature  = "signature"
)

type spotAccount struct {
	MakerCommission  int32               `json:"makerCommission"`
	TakerCommission  int32               `json:"takerCommission"`
	BuyerCommission  int32               `json:"buyerCommission"`
	SellerCommission int32               `json:"sellerCommission"`
	CanTrade         bool                `json:"canTrade"`
	CanWithdraw      bool                `json:"canWithdraw"`
	CanDeposit       bool                `json:"canDeposit"`
	UpdateTime       int64               `json:"updateTime"`
	AccountType      string              `json:"accountType"`
	Balances         []types.SpotBalance `json:"balances"`
	Permissions      []string            `json:"permissions"`
}
