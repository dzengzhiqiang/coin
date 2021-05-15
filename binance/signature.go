package binance

import (
	"coin/config"
	"github.com/civet148/gotools/cryptos/hmacsha"
	"net/url"
)

type Signature struct {
	cfg *config.Config
}

func NewSignature(cfg *config.Config) *Signature {

	return &Signature{
		cfg: cfg,
	}
}

//query string + HMAC + SHA256
func (s *Signature) Sign(values url.Values) (strQuery, strSign string) {
	strQuery = values.Encode()
	strSign = hmacsha.HmacSha256Hex(strQuery, s.cfg.AppSecret)
	return
}
