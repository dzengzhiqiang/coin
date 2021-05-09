package main

import (
	"coin/api"
	"fmt"
	"github.com/civet148/log"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"time"
)

const (
	VERSION      = "v0.1.0"
	PROGRAM_NAME = "coin"
	UPDATE_DATE  = "2021-05-08"
)

const (
	CMD_NAME_BINANCE = "binance"
	CMD_NAME_HUOBI   = "huobi"
)

const (
	SUB_CMD_NAME_RUN     = "run"
	SUB_CMD_NAME_BALANCE = "balance"
	SUB_CMD_NAME_TRADE   = "trade"
	SUB_CMD_NAME_PRICE   = "price"
)

const (
	CMD_FLAG_NAME_APPID     = "app-id"
	CMD_FLAG_NAME_APPKEY    = "app-key"
	CMD_FLAG_NAME_APPSECRET = "app-secret"
	CMD_FLAG_NAME_CONFIG    = "config"
)

const (
	COIN_ENV_NAME_APP_ID     = "COIN_ENV_APP_ID"
	COIN_ENV_NAME_APP_KEY    = "COIN_ENV_APP_KEY"
	COIN_ENV_NAME_APP_SECRET = "COIN_ENV_APP_SECRET"
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
					os.Exit(0)
				}
			}
		}
	}()
}

func main() {
	gracefulExit()

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
}

var binanceCmd = &cli.Command{
	Name:  CMD_NAME_BINANCE,
	Usage: "run binance coin trader",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    CMD_FLAG_NAME_APPID,
			Usage:   "app id",
			EnvVars: []string{COIN_ENV_NAME_APP_ID},
		},
		&cli.StringFlag{
			Name:    CMD_FLAG_NAME_APPKEY,
			Usage:   "app key",
			EnvVars: []string{COIN_ENV_NAME_APP_KEY},
		},
		&cli.StringFlag{
			Name:    CMD_FLAG_NAME_APPSECRET,
			Usage:   "app secret",
			EnvVars: []string{COIN_ENV_NAME_APP_SECRET},
		},
		&cli.StringFlag{
			Name:        CMD_FLAG_NAME_CONFIG,
			Usage:       "config file path",
			DefaultText: "coin.toml",
		},
	},
	Subcommands: []*cli.Command{
		runSubCmd,
		balanceSubCmd,
		priceSubCmd,
		tradeSubCmd,
	},
	Action: func(cctx *cli.Context) error {

		return nil
	},
}

var huobiCmd = &cli.Command{
	Name:  CMD_NAME_HUOBI,
	Usage: "run huobi coin trader",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    CMD_FLAG_NAME_APPID,
			Usage:   "app id",
			EnvVars: []string{COIN_ENV_NAME_APP_ID},
		},
		&cli.StringFlag{
			Name:    CMD_FLAG_NAME_APPKEY,
			Usage:   "app key",
			EnvVars: []string{COIN_ENV_NAME_APP_KEY},
		},
		&cli.StringFlag{
			Name:    CMD_FLAG_NAME_APPSECRET,
			Usage:   "app secret",
			EnvVars: []string{COIN_ENV_NAME_APP_SECRET},
		},
		&cli.StringFlag{
			Name:        CMD_FLAG_NAME_CONFIG,
			Usage:       "config file path",
			DefaultText: "coin.toml",
		},
	},
	Subcommands: []*cli.Command{
		runSubCmd,
		balanceSubCmd,
		priceSubCmd,
		tradeSubCmd,
	},
	Action: func(cctx *cli.Context) error {

		return nil
	},
}

var runSubCmd = &cli.Command{
	Name:  SUB_CMD_NAME_RUN,
	Usage: "run as a service",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {

		return nil
	},
}

var balanceSubCmd = &cli.Command{
	Name:  SUB_CMD_NAME_BALANCE,
	Usage: "query balances",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {

		return nil
	},
}

var priceSubCmd = &cli.Command{
	Name:  SUB_CMD_NAME_PRICE,
	Usage: "query coin price",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {

		return nil
	},
}

var tradeSubCmd = &cli.Command{
	Name:  SUB_CMD_NAME_TRADE,
	Usage: "coin trade",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {

		return nil
	},
}
