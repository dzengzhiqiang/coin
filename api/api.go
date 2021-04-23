package api

type CoinPrice struct {
	From   string
	To     string
	Remark string
}

type CoinApi interface {
	CoinGetPrices() ([]CoinPrice, error) //market trends of coin price
	CoinBuy()
	ConSell()
}
