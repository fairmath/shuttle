package config

import "github.com/urfave/cli/v2"

type Config struct {
	HTTPTendermintURL string
	WSTendermintURL   string
	ListenAddr        string
	LogLevel          string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) BuildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "ws-tendermint-url",
			Usage:       "set up a Tendermint node url",
			Destination: &c.WSTendermintURL,
			Required:    true,
			Value:       "http://localhost/websocket:1317",
			EnvVars:     []string{"TENDERMINT_URL"},
		},
		&cli.StringFlag{
			Name:        "http-tendermint-url",
			Usage:       "set up a Tendermint node url",
			Destination: &c.HTTPTendermintURL,
			Required:    true,
			Value:       "http://localhost:1317",
			EnvVars:     []string{"TENDERMINT_URL"},
		},
		&cli.StringFlag{
			Name:        "http-listen-addr",
			Usage:       "set up listening address",
			Destination: &c.ListenAddr,
			Required:    true,
			Value:       "0.0.0.0:9182",
			EnvVars:     []string{"HTTP_LISTEN_ADDR"},
		},
		&cli.StringFlag{
			Name:        "log-level",
			Usage:       "set up log level",
			Destination: &c.LogLevel,
			Value:       "info",
			EnvVars:     []string{"LOG_LEVEL"},
		},
	}
}
