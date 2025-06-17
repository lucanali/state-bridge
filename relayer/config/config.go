package config

import (
	"encoding/json"
	"os"
)

type ChainConfig struct {
	RPCURL     string `json:"rpc_url"`
	ChainID    uint64 `json:"chain_id"`
	BridgeAddr string `json:"bridge_address"`
	StartBlock uint64 `json:"start_block"`
}

type Config struct {
	SourceChain         ChainConfig `json:"source_chain"`
	DestinationChain    ChainConfig `json:"destination_chain"`
	PollInterval        int64       `json:"poll_interval"`         // in seconds
	ValidatorPrivateKey string      `json:"validator_private_key"` // Single private key for validator
}

func LoadConfig() (*Config, error) {
	// First try to load from environment
	configPath := os.Getenv("RELAYER_CONFIG")
	if configPath == "" {
		configPath = "config.json"
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
