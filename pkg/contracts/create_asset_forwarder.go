package contracts

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

// CreateAssetForwarder creates an AssetForwarder filterer with the given client and contract address of the AssetForwarder contract.
func CreateAssetForwarder(client *ethclient.Client, assetForwarderAddress common.Address) (*AssetForwarderFilterer, error) {
	assetForwarderFilter, err := NewAssetForwarderFilterer(assetForwarderAddress, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create AssetForwarder filterer")
	}

	return assetForwarderFilter, nil
}
