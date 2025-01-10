package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"

	"github.com/fairmath/shuttle/cmd/config"
	"github.com/fairmath/shuttle/internal/client/tendermint/api"
	"github.com/fairmath/shuttle/internal/server"
)

func runServer(cfg config.Config) error {
	log := config.LoadLogger(cfg.LogLevel)

	cl, err := api.NewCosmosAPI(cfg.HTTPTendermintURL, cfg.RPCTendermintURL, cfg.WSTendermintURL)
	if err != nil {
		return fmt.Errorf("connect to tendermint api: %w", err)
	}

	srv, err := server.NewServer(cfg, cl, log)
	if err != nil {
		return fmt.Errorf("new server: %w", err)
	}

	if err := srv.Run(); err != nil {
		log.Error("run server", zap.Error(err))
	}

	return nil
}

func main() {
	cfg := config.NewConfig()

	app := &cli.App{
		Name:  "github.com/fairmath/shuttle",
		Usage: "Cosmos - Ethereum rpc proxy",
		Description: `Shuttle is a proxy server for translation ethereum json rpc requests to any cosmos-based chain.\n
Developed to support integration between blockscout and cosmos-based blockchains.`,
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "Start cosmos-ethereum chains proxy",
				Action: func(_ *cli.Context) error {
					return runServer(*cfg)
				},
				Flags: cfg.BuildFlags(),
			},
		},
		EnableBashCompletion: true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
