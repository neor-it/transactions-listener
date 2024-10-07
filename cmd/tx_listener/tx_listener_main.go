package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/neor-it/transactions-listener/lib/tx_listener/config"
	"github.com/neor-it/transactions-listener/lib/tx_listener/svc/app"
	pkgcommon "github.com/neor-it/transactions-listener/pkg/common"
	"log"
)

func main() {
	// Recover from panics.
	defer pkgcommon.CatchPanic()

	// Create a new context.
	ctx := context.Background()

	// Load the configuration.
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load the configuration:", err)
	}

	// Connect to the Ethereum client.
	client, err := ethclient.Dial(cfg.RPCURL)
	if err != nil {
		log.Fatal("Error connecting to the Ethereum client:", err)
	}

	log.Println("Transactions listener service is started successfully")

	// Create a new service instance with the client and configuration.
	service := app.NewService(client, cfg)

	// Start the listener service.
	err = service.StartListener(ctx)
	if err != nil {
		log.Fatal("Failed to start the transactions listener service:", err)
	}
}
