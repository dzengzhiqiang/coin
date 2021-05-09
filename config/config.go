package config

import (
	"github.com/BurntSushi/toml"
	"github.com/civet148/log"
)

const (
	DEFAULT_CONFIG_FILE = "coin.toml"
)

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

func NewConfig(strConfig string) *Config {
	var cfg = &Config{}
	if strConfig == "" {
		strConfig = DEFAULT_CONFIG_FILE
	}
	if _, err := toml.DecodeFile(strConfig, cfg); err != nil {
		log.Errorf("load [%s] file error [%s]", strConfig, err.Error())
		return nil
	}
	log.Json(cfg)
	return cfg
}
