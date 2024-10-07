package services

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/neor-it/transactions-listener/lib/tx_listener/config"
	"github.com/neor-it/transactions-listener/pkg/contracts"
	"github.com/neor-it/transactions-listener/pkg/utils"
	"log"
	"time"
)

const (
	// waitingTime is the duration to wait before checking the transaction status again.
	waitingTime = 5 * time.Second
)

// EventProcessorService is an interface for processing events service in the application.
type EventProcessorService interface {
	// ProcessEvents processes the events from the given channel using the given EventProcessor.
	ProcessEvents(ctx context.Context, fundsDepositedChan chan *contracts.AssetForwarderFundsDeposited, processor EventProcessorService)
}

// eventProcessorService is a struct that implements the EventProcessor interface for processing events.
type eventProcessorService struct {
	client   *ethclient.Client
	chainIDs []string
}

// NewEventProcessorService creates a new EventProcessorService instance.
func NewEventProcessorService(client *ethclient.Client, cfg *config.TxLoaderConfig) EventProcessorService {
	return &eventProcessorService{
		client:   client,
		chainIDs: cfg.ChainIDs,
	}
}

// ProcessEvents processes the events from the given channel using the given EventProcessor.
func (p *eventProcessorService) ProcessEvents(ctx context.Context, fundsDepositedChan chan *contracts.AssetForwarderFundsDeposited, processor EventProcessorService) {
	for {
		select {
		case event := <-fundsDepositedChan:
			p.processFundsDeposited(ctx, event)
		}
	}
}

// processFundsDeposited processes a FundsDeposited event from the AssetForwarder contract.
func (p *eventProcessorService) processFundsDeposited(ctx context.Context, event *contracts.AssetForwarderFundsDeposited) {
	// Convert the destination chain ID to a string.
	destinationChainID := utils.Bytes32ToString(event.DestChainIdBytes)

	// Check if the destination chain ID is in the list of chain IDs.
	for _, chainId := range p.chainIDs {
		if chainId == destinationChainID {
			// Start a new goroutine to handle the transaction receipt checking asynchronously.
			go p.processTransactionReceipt(ctx, event, destinationChainID)

			break
		}
	}
}

// processTransactionReceipt processes the transaction receipt and prints the event details if the transaction was successful.
func (p *eventProcessorService) processTransactionReceipt(ctx context.Context, event *contracts.AssetForwarderFundsDeposited, destinationChainID string) {
	// Loop indefinitely until the transaction is either successful or failed.
	for {
		// Get the transaction using the transaction hash.
		tx, isPending, err := p.client.TransactionByHash(ctx, event.Raw.TxHash)
		// If there is an error in getting the transaction receipt, log the error and return from the goroutine.
		if err != nil {
			log.Printf("Failed to get transaction receipt: %v", err)

			return
		}

		if isPending {
			// If the transaction is still pending, wait for a while before checking again.
			time.Sleep(waitingTime)

			continue
		}

		if tx == nil {
			log.Printf("Transaction: %s is nil", event.Raw.TxHash)

			return
		}

		p.printEventDetails(destinationChainID, event)

		break
	}
}

// printEventDetails prints the details of a FundsDeposited event.
func (p *eventProcessorService) printEventDetails(chainID string, event *contracts.AssetForwarderFundsDeposited) {
	log.Println("Transfer to chain detected!")
	log.Println("TxHash", event.Raw.TxHash.Hex())
	log.Println("Chain ID:", chainID)
	log.Println("Amount:", event.Amount)
	log.Println("DepositId:", event.DepositId)
	log.Println("Depositor:", event.Depositor.Hex())
	log.Println("DestAmount:", event.DestAmount.Int64())
	log.Println("DestToken:", common.BytesToAddress(event.DestToken).Hex())
	log.Println("Recipient:", common.BytesToAddress(event.Recipient).Hex())
}
