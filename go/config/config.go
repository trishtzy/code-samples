package config

import (
	"log"
	"math/big"
	"strconv"

	"github.com/caarlos0/env/v11"
	"github.com/ethereum/go-ethereum/common/math"
)

type config struct {
	EthereumPrivateKey string `env:"ETHEREUM_PRIVATE_KEY"`
	ChainID            string `env:"PARADEX_CHAIN_ID" envDefault:"11155111"`
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

func (cfg config) GetChainIDBigInt() *math.HexOrDecimal256 {
	chainID := big.NewInt(cfg.GetChainID())
	return (*math.HexOrDecimal256)(chainID)
}
