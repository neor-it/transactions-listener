package app

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/neor-it/transactions-listener/lib/tx_listener/config"
	"github.com/neor-it/transactions-listener/lib/tx_listener/svc"
	"github.com/neor-it/transactions-listener/lib/tx_listener/svc/services"
	"github.com/neor-it/transactions-listener/pkg/contracts"
	"github.com/pkg/errors"
)

// service is a struct that implements the Service instance.
type service struct {
	client *ethclient.Client
	cfg    *config.TxLoaderConfig
}

// NewService creates a new instance of the Service interface.
func NewService(client *ethclient.Client, cfg *config.TxLoaderConfig) svc.Service {
	return &service{
		client: client,
		cfg:    cfg,
	}
}

// StartListener starts the transactions listener service with the given client and contract address.
func (s *service) StartListener(ctx context.Context) error {
	// Convert the contract address to a common.Address format.
	contractAddress := common.HexToAddress(s.cfg.ContractAddress)

	// Create an AssetForwarder filterer instance.
	assetForwarder, err := contracts.CreateAssetForwarder(s.client, contractAddress)
	if err != nil {
		return errors.Wrap(err, "failed to create AssetForwarder filterer instance")
	}

	// Create a new subscription service instance with the given AssetForwarder filterer.
	subscriptionService := services.NewSubscriptionService(assetForwarder)
	if subscriptionService == nil {
		return errors.New("failed to create subscription instance")
	}

	// Subscribe to FundsDeposited events.
	fundsDepositedSubscription, fundsDepositedChan, err := subscriptionService.SubscribeToFundsDeposited()
	if err != nil {
		return errors.Wrap(err, "failed to subscribe to FundsDeposited events")
	}
	defer fundsDepositedSubscription.Unsubscribe()

	// Create a new event processor instance with the given configuration.
	eventProcessor := services.NewEventProcessorService(s.client, s.cfg)

	// Process the events from the channel using the event processor.
	eventProcessor.ProcessEvents(ctx, fundsDepositedChan, eventProcessor)

	return nil
}
