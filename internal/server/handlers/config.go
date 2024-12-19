package handlers

import "github.com/urfave/cli/v2"

type Config struct {
	Bech32AddrPrefix string
	CosmosDenom      string
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) BuildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "bech-prefix",
			Usage:       "set up cosmos address bech32 prefix",
			Destination: &cfg.Bech32AddrPrefix,
			Value:       "cosmos",
			EnvVars:     []string{"BECH_PREFIX"},
		},
		&cli.StringFlag{
			Name:        "coin-denom",
			Usage:       "set up denom in cosmos chain",
			Destination: &cfg.CosmosDenom,
			Value:       "uatom",
			EnvVars:     []string{"COIN_DENOM"},
		},
	}
}
