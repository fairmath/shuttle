package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fairmath/shuttle/cmd/config"
	"github.com/fairmath/shuttle/internal/client/tendermint/api"
	"github.com/fairmath/shuttle/internal/server"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func runServer(cfg config.Config) error {
	log := config.LoadLogger(cfg.LogLevel)

	cl := api.NewCosmosApi(cfg.TendermintURL, "/")

	srv, err := server.NewServer(cfg.ListenAddr, cl, log)
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
