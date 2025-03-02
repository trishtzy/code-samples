package config

import (
	"log"
	"strconv"

	"github.com/caarlos0/env/v11"
)

type config struct {
	EthereumPrivateKey string `env:"ETHEREUM_PRIVATE_KEY"`
	ParadexPrivateKey  string `env:"PARADEX_PRIVATE_KEY"`
	ChainID            string `env:"PARADEX_CHAIN_ID" envDefault:"1"`
	ParadexVersion     string `env:"PARADEX_VERSION" envDefault:"1"`
}

var App config

func LoadEnv() {
	err := env.Parse(&App)
	if err != nil {
		log.Fatalf("Failed to parse environment variables: %v", err)
	}
}

func (cfg config) GetChainID() int64 {
	chainID, _ := strconv.ParseInt(cfg.ChainID, 10, 64)
	return chainID
}
