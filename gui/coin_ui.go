package gui

import (
	"coin/config"
	"github.com/civet148/log"
	"github.com/jroimartin/gocui"
)

type CoinUI struct {
	cfg *config.Config
	gui *gocui.Gui
}

func NewCoinUI(cfg *config.Config) *CoinUI  {
	gui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Errorf(err.Error())
		panic(err.Error())
	}
	return &CoinUI{
		cfg: cfg,
		gui: gui,
	}
}

func (m *CoinUI) Close() {
	m.gui.Close()
}
