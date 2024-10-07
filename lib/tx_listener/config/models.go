package config

// TxLoaderConfig represents the configuration of the application.
type TxLoaderConfig struct {
	ContractAddress string   `envconfig:"CONTRACT_ADDRESS"`
	RPCURL          string   `envconfig:"RPC_URL"`
	ChainIDs        []string `envconfig:"CHAIN_IDS"`
}
