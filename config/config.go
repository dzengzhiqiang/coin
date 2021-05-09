package config

type Config struct {
	AppId     string `toml:"app_id"`
	AppKey    string `toml:"app_key"`
	AppSecret string `toml:"app_secret"`
	Coins     []Coin `toml:"coins"`
}

type Coin struct {
	BaseAsset   string  `toml:"base_asset"`
	QuoteAsset  string  `toml:"quote_asset"`
	SellPercent float64 `toml:"sell_percent"`
	SellRate    float64 `toml:"sell_rate"`
	BuyRate     float64 `toml:"buy_rate"`
}
