package services

import (
	"github.com/ethereum/go-ethereum/event"
	"github.com/neor-it/transactions-listener/pkg/contracts"
)

// SubscriptionsService is the interface for subscribing to events from the AssetForwarder contract.
type SubscriptionsService interface {
	// SubscribeToFundsDeposited subscribes to FundsDeposited events from the given AssetForwarder contract.
	SubscribeToFundsDeposited() (event.Subscription, chan *contracts.AssetForwarderFundsDeposited, error)
}

// subscriptionsService is a struct that implements the SubscriptionsService interface for subscribing to events.
type subscriptionsService struct {
	assetForwarder *contracts.AssetForwarderFilterer
}

// NewSubscriptionService creates a new subscription service with the given AssetForwarder filterer.
func NewSubscriptionService(assetForwarder *contracts.AssetForwarderFilterer) SubscriptionsService {
	return &subscriptionsService{
		assetForwarder: assetForwarder,
	}
}

// SubscribeToFundsDeposited subscribes to FundsDeposited events from the given AssetForwarder contract.
func (s *subscriptionsService) SubscribeToFundsDeposited() (event.Subscription, chan *contracts.AssetForwarderFundsDeposited, error) {
	eventsChan := make(chan *contracts.AssetForwarderFundsDeposited)

	// Subscribe to FundsDeposited events from the AssetForwarder contract.
	fundsDepositedSubscription, err := s.assetForwarder.WatchFundsDeposited(nil, eventsChan)
	if err != nil {
		return nil, nil, err
	}

	// Return the subscription and events channel.
	return fundsDepositedSubscription, eventsChan, nil
}
