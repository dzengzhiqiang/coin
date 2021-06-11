package main

import (
	"coin/api"
	"coin/binance"
	"coin/config"
	"coin/huobi"
	"coin/types"
	"fmt"
	"github.com/civet148/log"
	"github.com/civet148/sqlca"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"time"
)

const (
	VERSION      = "v0.1.0"
	PROGRAM_NAME = "coin"
	UPDATE_DATE  = "2021-05-10"
)

const (
	CMD_NAME_BINANCE = "binance"
	CMD_NAME_HUOBI   = "huobi"
)

const (
	SUB_CMD_NAME_RUN     = "run"
	SUB_CMD_NAME_ACCOUNT = "account"
	SUB_CMD_NAME_BUY     = "buy"
	SUB_CMD_NAME_SELL    = "sell"
	SUB_CMD_NAME_PRICE   = "price"
)

const (
	CMD_FLAG_NAME_APPID     = "app-id"
	CMD_FLAG_NAME_APPKEY    = "app-key"
	CMD_FLAG_NAME_APPSECRET = "app-secret"
	CMD_FLAG_NAME_CONFIG    = "config"
)

const (
	COIN_ENV_APP_ID     = "COIN_ENV_APP_ID"
	COIN_ENV_APP_KEY    = "COIN_ENV_APP_KEY"
	COIN_ENV_APP_SECRET = "COIN_ENV_APP_SECRET"
)

var coin api.CoinApi

func gracefulExit() {
	//capture signal of Ctrl+C and gracefully exit
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	defer close(sigChannel)
	go func() {
		for {
			select {
			case <-sigChannel:
				{
					if coin != nil {
						coin.Close()
					}
					time.Sleep(500 * time.Millisecond)
					log.Infof("program exit...")
					os.Exit(0)
				}
			}
		}
	}()
}

func init() {
	log.SetLevel(0)
	log.CloseConsole(true)
	log.Open("coin.log")
}

func main() {
	local := []*cli.Command{
		binanceCmd,
		huobiCmd,
	}
	app := &cli.App{
		Name:     PROGRAM_NAME,
		Usage:    "coin manager",
		Version:  fmt.Sprintf("v%s %s", VERSION, UPDATE_DATE),
		Flags:    []cli.Flag{},
		Commands: local,
	}
	if err := app.Run(os.Args); err != nil {
		log.Errorf("exit in error %s", err)
		os.Exit(1)
		return
	}
	time.Sleep(5 * time.Second)
}

var binanceCmd = &cli.Command{
	Name:  CMD_NAME_BINANCE,
	Usage: "run binance coin trader",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    CMD_FLAG_NAME_APPID,
			Usage:   "app id",
			EnvVars: []string{COIN_ENV_APP_ID},
		},
		&cli.StringFlag{
			Name:    CMD_FLAG_NAME_APPKEY,
			Usage:   "app key",
			EnvVars: []string{COIN_ENV_APP_KEY},
		},
		&cli.StringFlag{
			Name:    CMD_FLAG_NAME_APPSECRET,
			Usage:   "app secret",
			EnvVars: []string{COIN_ENV_APP_SECRET},
		},
		&cli.StringFlag{
			Name:  CMD_FLAG_NAME_CONFIG,
			Usage: "config file path",
		},
	},
	Before: func(ctx *cli.Context) error {
		cfg := config.NewConfig(ctx.String(CMD_FLAG_NAME_CONFIG))
		if cfg == nil {
			return fmt.Errorf("load config file failed")
		}
		if ctx.IsSet(CMD_FLAG_NAME_APPID) {
			cfg.AppId = ctx.String(CMD_FLAG_NAME_APPID)
		}
		if ctx.IsSet(CMD_FLAG_NAME_APPKEY) {
			cfg.AppKey = ctx.String(CMD_FLAG_NAME_APPKEY)
		}
		if ctx.IsSet(CMD_FLAG_NAME_APPSECRET) {
			cfg.AppSecret = ctx.String(CMD_FLAG_NAME_APPSECRET)
		}
		coin = binance.NewCoinManager(cfg)
		return nil
	},
	Subcommands: []*cli.Command{
		runSubCmd,
		accountSubCmd,
		priceSubCmd,
		buySubCmd,
		sellSubCmd,
	},
}

var huobiCmd = &cli.Command{
	Name:  CMD_NAME_HUOBI,
	Usage: "run huobi coin trader",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    CMD_FLAG_NAME_APPID,
			Usage:   "app id",
			EnvVars: []string{COIN_ENV_APP_ID},
		},
		&cli.StringFlag{
			Name:    CMD_FLAG_NAME_APPKEY,
			Usage:   "app key",
			EnvVars: []string{COIN_ENV_APP_KEY},
		},
		&cli.StringFlag{
			Name:    CMD_FLAG_NAME_APPSECRET,
			Usage:   "app secret",
			EnvVars: []string{COIN_ENV_APP_SECRET},
		},
		&cli.StringFlag{
			Name:  CMD_FLAG_NAME_CONFIG,
			Usage: "config file path",
		},
	},
	Before: func(ctx *cli.Context) error {
		cfg := config.NewConfig(ctx.String(CMD_FLAG_NAME_CONFIG))
		if cfg == nil {
			return fmt.Errorf("load config file failed")
		}
		if ctx.IsSet(CMD_FLAG_NAME_APPID) {
			cfg.AppId = ctx.String(CMD_FLAG_NAME_APPID)
		}
		if ctx.IsSet(CMD_FLAG_NAME_APPKEY) {
			cfg.AppKey = ctx.String(CMD_FLAG_NAME_APPKEY)
		}
		if ctx.IsSet(CMD_FLAG_NAME_APPSECRET) {
			cfg.AppSecret = ctx.String(CMD_FLAG_NAME_APPSECRET)
		}
		coin = huobi.NewCoinManager(cfg)
		return nil
	},
	Subcommands: []*cli.Command{
		runSubCmd,
		accountSubCmd,
		priceSubCmd,
		buySubCmd,
		sellSubCmd,
	},
}

var runSubCmd = &cli.Command{
	Name:  SUB_CMD_NAME_RUN,
	Usage: "run as a service",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		log.Debugf("run sub command action")
		return coin.Run()
	},
}

var accountSubCmd = &cli.Command{
	Name:  SUB_CMD_NAME_ACCOUNT,
	Usage: "query account balances",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		balances, _ := coin.SpotBalances()
		log.Json(balances)
		return nil
	},
}

var priceSubCmd = &cli.Command{
	Name:  SUB_CMD_NAME_PRICE,
	Usage: "query coin price",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() == 0 {
			log.Errorf("please input symbol to query price")
			return nil
		}
		price, _ := coin.SpotPrice(cctx.Args().First())
		log.Json(price)
		return nil
	},
}

var sellSubCmd = &cli.Command{
	Name:  SUB_CMD_NAME_SELL,
	Usage: "coin trade [SELL]",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		var args = cctx.Args()
		if args.Len() != 2 {
			return fmt.Errorf("args less than 2 (symbol and quantity)")
		}
		r, c := coin.SpotSell(&types.SpotTradeReq{
			Symbol:   args.Get(0),
			Quantity: sqlca.NewDecimal(args.Get(1)),
		})
		if c != types.BizCode_OK {
			return fmt.Errorf("[SELL] trade failed with code [%v]", c)
		}
		log.Json(r)
		return nil
	},
}
var buySubCmd = &cli.Command{
	Name:  SUB_CMD_NAME_BUY,
	Usage: "coin trade [BUY]",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {

		var args = cctx.Args()
		if args.Len() != 2 {
			return fmt.Errorf("args less than 2 (symbol and quantity)")
		}
		r, c := coin.SpotBuy(&types.SpotTradeReq{
			Symbol:   args.Get(0),
			Quantity: sqlca.NewDecimal(args.Get(1)),
		})
		if c != types.BizCode_OK {
			return fmt.Errorf("[BUY] trade failed with code [%v]", c)
		}
		log.Json(r)
		return nil
	},
}
