package dal

import "coin/config"

type Dal struct {
	cfg *config.Config
}

func NewDal(cfg *config.Config) *Dal {
	return &Dal{
		cfg: cfg,
	}
}
