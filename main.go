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
	CMD_FLAG_NAME_APPID     = "app-id"
	CMD_FLAG_NAME_APPKEY    = "app-key"
	CMD_FLAG_NAME_APPSECRET = "app-secret"
)

var manager api.ManagerApi

func main() {
	//capture signal of Ctrl+C and gracefully exit
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	defer close(sigChannel)
	go func() {
		for {
			select {
			case <-sigChannel:
				{
					if manager != nil {
						manager.Close()
					}
					time.Sleep(500 * time.Millisecond)
					os.Exit(0)
				}
			}
		}
	}()

	local := []*cli.Command{
		binanceCmd,
		huobiCmd,
	}
	app := &cli.App{
		Name:     PROGRAM_NAME,
		Usage:    "stos manager",
		Version:  fmt.Sprintf("v%s %s", VERSION, UPDATE_DATE),
		Flags:    []cli.Flag{},
		Commands: local,
		Action:   nil,
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

	},
	Action: func(cctx *cli.Context) error {

		return nil
	},
}

var huobiCmd = &cli.Command{
	Name:  CMD_NAME_HUOBI,
	Usage: "run huobi coin trader",
	Flags: []cli.Flag{

	},
	Action: func(cctx *cli.Context) error {

		return nil
	},
}
