package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ChainConfig struct {
	RPCURL     string `json:"rpcUrl"`
	ChainID    uint64 `json:"chainId"`
	BridgeAddr string `json:"bridgeAddr"`
	StartBlock uint64 `json:"startBlock,omitempty"`
}

type Config struct {
	SourceChain         ChainConfig `json:"sourceChain"`
	DestinationChain    ChainConfig `json:"destinationChain"`
	PollInterval        int         `json:"pollInterval"`
	ValidatorPrivateKey string      `json:"validatorPrivateKey"`
}

func LoadConfig() (*Config, error) {
	// First try to load from environment
	configPath := os.Getenv("RELAYER_CONFIG")
	if configPath == "" {
		configPath = "config.json"
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config: %v", err)
	}

	// Validate config
	if config.SourceChain.RPCURL == "" {
		return nil, fmt.Errorf("source chain RPC URL is required")
	}
	if config.DestinationChain.RPCURL == "" {
		return nil, fmt.Errorf("destination chain RPC URL is required")
	}
	if config.ValidatorPrivateKey == "" {
		return nil, fmt.Errorf("validator private key is required")
	}

	return &config, nil
}
