// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IAssetForwarderDepositData is an auto generated low-level Go binding around an user-defined struct.
type IAssetForwarderDepositData struct {
	PartnerId        *big.Int
	Amount           *big.Int
	DestAmount       *big.Int
	SrcToken         common.Address
	RefundRecipient  common.Address
	DestChainIdBytes [32]byte
}

// IAssetForwarderDestDetails is an auto generated low-level Go binding around an user-defined struct.
type IAssetForwarderDestDetails struct {
	DomainId uint32
	Fee      *big.Int
	IsSet    bool
}

// IAssetForwarderRelayData is an auto generated low-level Go binding around an user-defined struct.
type IAssetForwarderRelayData struct {
	Amount     *big.Int
	SrcChainId [32]byte
	DepositId  *big.Int
	DestToken  common.Address
	Recipient  common.Address
}

// IAssetForwarderRelayDataMessage is an auto generated low-level Go binding around an user-defined struct.
type IAssetForwarderRelayDataMessage struct {
	Amount     *big.Int
	SrcChainId [32]byte
	DepositId  *big.Int
	DestToken  common.Address
	Recipient  common.Address
	Message    []byte
}

// AssetForwarderMetaData contains all meta data concerning the AssetForwarder contract.
var AssetForwarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wrappedNativeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gatewayContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdcAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenMessenger\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_routerMiddlewareBase\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_minGasThreshhold\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AmountTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGateway\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRefundData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageAlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageExcecutionFailedWithLowGas\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"}],\"name\":\"CommunityPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eventNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"initiatewithdrawal\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"DepositInfoUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"}],\"name\":\"FundsDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"FundsDepositedWithMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"FundsPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"execFlag\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"execData\",\"type\":\"bytes\"}],\"name\":\"FundsPaidWithMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdcNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"iUSDCDeposited\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TRANSFER_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_GAS_THRESHHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESOURCE_SETTER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"communityPause\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"destDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"domainId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executeRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatewayContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssetForwarder.DepositData\",\"name\":\"depositData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"}],\"name\":\"iDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"initiatewithdrawal\",\"type\":\"bool\"}],\"name\":\"iDepositInfoUpdate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssetForwarder.DepositData\",\"name\":\"depositData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"iDepositMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"iDepositUSDC\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"requestSender\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"packet\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"iReceive\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcChainId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structIAssetForwarder.RelayData\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"name\":\"iRelay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcChainId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structIAssetForwarder.RelayDataMessage\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"name\":\"iRelayMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCommunityPauseEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseStakeAmountMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseStakeAmountMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"rescue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"routerMiddlewareBase\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_destChainIdBytes\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domainId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"internalType\":\"structIAssetForwarder.DestDetails[]\",\"name\":\"_destDetails\",\"type\":\"tuple[]\"}],\"name\":\"setDestDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleCommunityPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenMessenger\",\"outputs\":[{\"internalType\":\"contractITokenMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gatewayContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_routerMiddlewareBase\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"minPauseStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPauseStakeAmount\",\"type\":\"uint256\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenMessenger\",\"type\":\"address\"}],\"name\":\"updateTokenMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStakeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedNativeToken\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AssetForwarderABI is the input ABI used to generate the binding from.
// Deprecated: Use AssetForwarderMetaData.ABI instead.
var AssetForwarderABI = AssetForwarderMetaData.ABI

// AssetForwarder is an auto generated Go binding around an Ethereum contract.
type AssetForwarder struct {
	AssetForwarderCaller     // Read-only binding to the contract
	AssetForwarderTransactor // Write-only binding to the contract
	AssetForwarderFilterer   // Log filterer for contract events
}

// AssetForwarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetForwarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetForwarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetForwarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetForwarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetForwarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetForwarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetForwarderSession struct {
	Contract     *AssetForwarder   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetForwarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetForwarderCallerSession struct {
	Contract *AssetForwarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AssetForwarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetForwarderTransactorSession struct {
	Contract     *AssetForwarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AssetForwarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetForwarderRaw struct {
	Contract *AssetForwarder // Generic contract binding to access the raw methods on
}

// AssetForwarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetForwarderCallerRaw struct {
	Contract *AssetForwarderCaller // Generic read-only contract binding to access the raw methods on
}

// AssetForwarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetForwarderTransactorRaw struct {
	Contract *AssetForwarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetForwarder creates a new instance of AssetForwarder, bound to a specific deployed contract.
func NewAssetForwarder(address common.Address, backend bind.ContractBackend) (*AssetForwarder, error) {
	contract, err := bindAssetForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetForwarder{AssetForwarderCaller: AssetForwarderCaller{contract: contract}, AssetForwarderTransactor: AssetForwarderTransactor{contract: contract}, AssetForwarderFilterer: AssetForwarderFilterer{contract: contract}}, nil
}

// NewAssetForwarderCaller creates a new read-only instance of AssetForwarder, bound to a specific deployed contract.
func NewAssetForwarderCaller(address common.Address, caller bind.ContractCaller) (*AssetForwarderCaller, error) {
	contract, err := bindAssetForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetForwarderCaller{contract: contract}, nil
}

// NewAssetForwarderTransactor creates a new write-only instance of AssetForwarder, bound to a specific deployed contract.
func NewAssetForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetForwarderTransactor, error) {
	contract, err := bindAssetForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetForwarderTransactor{contract: contract}, nil
}

// NewAssetForwarderFilterer creates a new log filterer instance of AssetForwarder, bound to a specific deployed contract.
func NewAssetForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetForwarderFilterer, error) {
	contract, err := bindAssetForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetForwarderFilterer{contract: contract}, nil
}

// bindAssetForwarder binds a generic wrapper to an already deployed contract.
func bindAssetForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AssetForwarderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetForwarder *AssetForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetForwarder.Contract.AssetForwarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetForwarder *AssetForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetForwarder.Contract.AssetForwarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetForwarder *AssetForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetForwarder.Contract.AssetForwarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetForwarder *AssetForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetForwarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetForwarder *AssetForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetForwarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetForwarder *AssetForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetForwarder.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AssetForwarder *AssetForwarderCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AssetForwarder *AssetForwarderSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AssetForwarder.Contract.DEFAULTADMINROLE(&_AssetForwarder.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AssetForwarder *AssetForwarderCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AssetForwarder.Contract.DEFAULTADMINROLE(&_AssetForwarder.CallOpts)
}

// MAXTRANSFERSIZE is a free data retrieval call binding the contract method 0xddd224f1.
//
// Solidity: function MAX_TRANSFER_SIZE() view returns(uint256)
func (_AssetForwarder *AssetForwarderCaller) MAXTRANSFERSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "MAX_TRANSFER_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTRANSFERSIZE is a free data retrieval call binding the contract method 0xddd224f1.
//
// Solidity: function MAX_TRANSFER_SIZE() view returns(uint256)
func (_AssetForwarder *AssetForwarderSession) MAXTRANSFERSIZE() (*big.Int, error) {
	return _AssetForwarder.Contract.MAXTRANSFERSIZE(&_AssetForwarder.CallOpts)
}

// MAXTRANSFERSIZE is a free data retrieval call binding the contract method 0xddd224f1.
//
// Solidity: function MAX_TRANSFER_SIZE() view returns(uint256)
func (_AssetForwarder *AssetForwarderCallerSession) MAXTRANSFERSIZE() (*big.Int, error) {
	return _AssetForwarder.Contract.MAXTRANSFERSIZE(&_AssetForwarder.CallOpts)
}

// MINGASTHRESHHOLD is a free data retrieval call binding the contract method 0x4b7b9476.
//
// Solidity: function MIN_GAS_THRESHHOLD() view returns(uint256)
func (_AssetForwarder *AssetForwarderCaller) MINGASTHRESHHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "MIN_GAS_THRESHHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINGASTHRESHHOLD is a free data retrieval call binding the contract method 0x4b7b9476.
//
// Solidity: function MIN_GAS_THRESHHOLD() view returns(uint256)
func (_AssetForwarder *AssetForwarderSession) MINGASTHRESHHOLD() (*big.Int, error) {
	return _AssetForwarder.Contract.MINGASTHRESHHOLD(&_AssetForwarder.CallOpts)
}

// MINGASTHRESHHOLD is a free data retrieval call binding the contract method 0x4b7b9476.
//
// Solidity: function MIN_GAS_THRESHHOLD() view returns(uint256)
func (_AssetForwarder *AssetForwarderCallerSession) MINGASTHRESHHOLD() (*big.Int, error) {
	return _AssetForwarder.Contract.MINGASTHRESHHOLD(&_AssetForwarder.CallOpts)
}

// PAUSER is a free data retrieval call binding the contract method 0xd9dc8694.
//
// Solidity: function PAUSER() view returns(bytes32)
func (_AssetForwarder *AssetForwarderCaller) PAUSER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "PAUSER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSER is a free data retrieval call binding the contract method 0xd9dc8694.
//
// Solidity: function PAUSER() view returns(bytes32)
func (_AssetForwarder *AssetForwarderSession) PAUSER() ([32]byte, error) {
	return _AssetForwarder.Contract.PAUSER(&_AssetForwarder.CallOpts)
}

// PAUSER is a free data retrieval call binding the contract method 0xd9dc8694.
//
// Solidity: function PAUSER() view returns(bytes32)
func (_AssetForwarder *AssetForwarderCallerSession) PAUSER() ([32]byte, error) {
	return _AssetForwarder.Contract.PAUSER(&_AssetForwarder.CallOpts)
}

// RESOURCESETTER is a free data retrieval call binding the contract method 0xdb618e2a.
//
// Solidity: function RESOURCE_SETTER() view returns(bytes32)
func (_AssetForwarder *AssetForwarderCaller) RESOURCESETTER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "RESOURCE_SETTER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESOURCESETTER is a free data retrieval call binding the contract method 0xdb618e2a.
//
// Solidity: function RESOURCE_SETTER() view returns(bytes32)
func (_AssetForwarder *AssetForwarderSession) RESOURCESETTER() ([32]byte, error) {
	return _AssetForwarder.Contract.RESOURCESETTER(&_AssetForwarder.CallOpts)
}

// RESOURCESETTER is a free data retrieval call binding the contract method 0xdb618e2a.
//
// Solidity: function RESOURCE_SETTER() view returns(bytes32)
func (_AssetForwarder *AssetForwarderCallerSession) RESOURCESETTER() ([32]byte, error) {
	return _AssetForwarder.Contract.RESOURCESETTER(&_AssetForwarder.CallOpts)
}

// DepositNonce is a free data retrieval call binding the contract method 0xde35f5cb.
//
// Solidity: function depositNonce() view returns(uint256)
func (_AssetForwarder *AssetForwarderCaller) DepositNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "depositNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositNonce is a free data retrieval call binding the contract method 0xde35f5cb.
//
// Solidity: function depositNonce() view returns(uint256)
func (_AssetForwarder *AssetForwarderSession) DepositNonce() (*big.Int, error) {
	return _AssetForwarder.Contract.DepositNonce(&_AssetForwarder.CallOpts)
}

// DepositNonce is a free data retrieval call binding the contract method 0xde35f5cb.
//
// Solidity: function depositNonce() view returns(uint256)
func (_AssetForwarder *AssetForwarderCallerSession) DepositNonce() (*big.Int, error) {
	return _AssetForwarder.Contract.DepositNonce(&_AssetForwarder.CallOpts)
}

// DestDetails is a free data retrieval call binding the contract method 0x981a8a02.
//
// Solidity: function destDetails(bytes32 ) view returns(uint32 domainId, uint256 fee, bool isSet)
func (_AssetForwarder *AssetForwarderCaller) DestDetails(opts *bind.CallOpts, arg0 [32]byte) (struct {
	DomainId uint32
	Fee      *big.Int
	IsSet    bool
}, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "destDetails", arg0)

	outstruct := new(struct {
		DomainId uint32
		Fee      *big.Int
		IsSet    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DomainId = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Fee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsSet = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// DestDetails is a free data retrieval call binding the contract method 0x981a8a02.
//
// Solidity: function destDetails(bytes32 ) view returns(uint32 domainId, uint256 fee, bool isSet)
func (_AssetForwarder *AssetForwarderSession) DestDetails(arg0 [32]byte) (struct {
	DomainId uint32
	Fee      *big.Int
	IsSet    bool
}, error) {
	return _AssetForwarder.Contract.DestDetails(&_AssetForwarder.CallOpts, arg0)
}

// DestDetails is a free data retrieval call binding the contract method 0x981a8a02.
//
// Solidity: function destDetails(bytes32 ) view returns(uint32 domainId, uint256 fee, bool isSet)
func (_AssetForwarder *AssetForwarderCallerSession) DestDetails(arg0 [32]byte) (struct {
	DomainId uint32
	Fee      *big.Int
	IsSet    bool
}, error) {
	return _AssetForwarder.Contract.DestDetails(&_AssetForwarder.CallOpts, arg0)
}

// ExecuteRecord is a free data retrieval call binding the contract method 0xfd5ad37c.
//
// Solidity: function executeRecord(bytes32 ) view returns(bool)
func (_AssetForwarder *AssetForwarderCaller) ExecuteRecord(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "executeRecord", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExecuteRecord is a free data retrieval call binding the contract method 0xfd5ad37c.
//
// Solidity: function executeRecord(bytes32 ) view returns(bool)
func (_AssetForwarder *AssetForwarderSession) ExecuteRecord(arg0 [32]byte) (bool, error) {
	return _AssetForwarder.Contract.ExecuteRecord(&_AssetForwarder.CallOpts, arg0)
}

// ExecuteRecord is a free data retrieval call binding the contract method 0xfd5ad37c.
//
// Solidity: function executeRecord(bytes32 ) view returns(bool)
func (_AssetForwarder *AssetForwarderCallerSession) ExecuteRecord(arg0 [32]byte) (bool, error) {
	return _AssetForwarder.Contract.ExecuteRecord(&_AssetForwarder.CallOpts, arg0)
}

// GatewayContract is a free data retrieval call binding the contract method 0xeb0cde1d.
//
// Solidity: function gatewayContract() view returns(address)
func (_AssetForwarder *AssetForwarderCaller) GatewayContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "gatewayContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GatewayContract is a free data retrieval call binding the contract method 0xeb0cde1d.
//
// Solidity: function gatewayContract() view returns(address)
func (_AssetForwarder *AssetForwarderSession) GatewayContract() (common.Address, error) {
	return _AssetForwarder.Contract.GatewayContract(&_AssetForwarder.CallOpts)
}

// GatewayContract is a free data retrieval call binding the contract method 0xeb0cde1d.
//
// Solidity: function gatewayContract() view returns(address)
func (_AssetForwarder *AssetForwarderCallerSession) GatewayContract() (common.Address, error) {
	return _AssetForwarder.Contract.GatewayContract(&_AssetForwarder.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AssetForwarder *AssetForwarderCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AssetForwarder *AssetForwarderSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AssetForwarder.Contract.GetRoleAdmin(&_AssetForwarder.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AssetForwarder *AssetForwarderCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AssetForwarder.Contract.GetRoleAdmin(&_AssetForwarder.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AssetForwarder *AssetForwarderCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AssetForwarder *AssetForwarderSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AssetForwarder.Contract.HasRole(&_AssetForwarder.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AssetForwarder *AssetForwarderCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AssetForwarder.Contract.HasRole(&_AssetForwarder.CallOpts, role, account)
}

// IsCommunityPauseEnabled is a free data retrieval call binding the contract method 0x4463182f.
//
// Solidity: function isCommunityPauseEnabled() view returns(bool)
func (_AssetForwarder *AssetForwarderCaller) IsCommunityPauseEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "isCommunityPauseEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCommunityPauseEnabled is a free data retrieval call binding the contract method 0x4463182f.
//
// Solidity: function isCommunityPauseEnabled() view returns(bool)
func (_AssetForwarder *AssetForwarderSession) IsCommunityPauseEnabled() (bool, error) {
	return _AssetForwarder.Contract.IsCommunityPauseEnabled(&_AssetForwarder.CallOpts)
}

// IsCommunityPauseEnabled is a free data retrieval call binding the contract method 0x4463182f.
//
// Solidity: function isCommunityPauseEnabled() view returns(bool)
func (_AssetForwarder *AssetForwarderCallerSession) IsCommunityPauseEnabled() (bool, error) {
	return _AssetForwarder.Contract.IsCommunityPauseEnabled(&_AssetForwarder.CallOpts)
}

// PauseStakeAmountMax is a free data retrieval call binding the contract method 0xf215f148.
//
// Solidity: function pauseStakeAmountMax() view returns(uint256)
func (_AssetForwarder *AssetForwarderCaller) PauseStakeAmountMax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "pauseStakeAmountMax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PauseStakeAmountMax is a free data retrieval call binding the contract method 0xf215f148.
//
// Solidity: function pauseStakeAmountMax() view returns(uint256)
func (_AssetForwarder *AssetForwarderSession) PauseStakeAmountMax() (*big.Int, error) {
	return _AssetForwarder.Contract.PauseStakeAmountMax(&_AssetForwarder.CallOpts)
}

// PauseStakeAmountMax is a free data retrieval call binding the contract method 0xf215f148.
//
// Solidity: function pauseStakeAmountMax() view returns(uint256)
func (_AssetForwarder *AssetForwarderCallerSession) PauseStakeAmountMax() (*big.Int, error) {
	return _AssetForwarder.Contract.PauseStakeAmountMax(&_AssetForwarder.CallOpts)
}

// PauseStakeAmountMin is a free data retrieval call binding the contract method 0xb19941a9.
//
// Solidity: function pauseStakeAmountMin() view returns(uint256)
func (_AssetForwarder *AssetForwarderCaller) PauseStakeAmountMin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "pauseStakeAmountMin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PauseStakeAmountMin is a free data retrieval call binding the contract method 0xb19941a9.
//
// Solidity: function pauseStakeAmountMin() view returns(uint256)
func (_AssetForwarder *AssetForwarderSession) PauseStakeAmountMin() (*big.Int, error) {
	return _AssetForwarder.Contract.PauseStakeAmountMin(&_AssetForwarder.CallOpts)
}

// PauseStakeAmountMin is a free data retrieval call binding the contract method 0xb19941a9.
//
// Solidity: function pauseStakeAmountMin() view returns(uint256)
func (_AssetForwarder *AssetForwarderCallerSession) PauseStakeAmountMin() (*big.Int, error) {
	return _AssetForwarder.Contract.PauseStakeAmountMin(&_AssetForwarder.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AssetForwarder *AssetForwarderCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AssetForwarder *AssetForwarderSession) Paused() (bool, error) {
	return _AssetForwarder.Contract.Paused(&_AssetForwarder.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AssetForwarder *AssetForwarderCallerSession) Paused() (bool, error) {
	return _AssetForwarder.Contract.Paused(&_AssetForwarder.CallOpts)
}

// RouterMiddlewareBase is a free data retrieval call binding the contract method 0xc44e947e.
//
// Solidity: function routerMiddlewareBase() view returns(bytes32)
func (_AssetForwarder *AssetForwarderCaller) RouterMiddlewareBase(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "routerMiddlewareBase")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RouterMiddlewareBase is a free data retrieval call binding the contract method 0xc44e947e.
//
// Solidity: function routerMiddlewareBase() view returns(bytes32)
func (_AssetForwarder *AssetForwarderSession) RouterMiddlewareBase() ([32]byte, error) {
	return _AssetForwarder.Contract.RouterMiddlewareBase(&_AssetForwarder.CallOpts)
}

// RouterMiddlewareBase is a free data retrieval call binding the contract method 0xc44e947e.
//
// Solidity: function routerMiddlewareBase() view returns(bytes32)
func (_AssetForwarder *AssetForwarderCallerSession) RouterMiddlewareBase() ([32]byte, error) {
	return _AssetForwarder.Contract.RouterMiddlewareBase(&_AssetForwarder.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AssetForwarder *AssetForwarderCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AssetForwarder *AssetForwarderSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AssetForwarder.Contract.SupportsInterface(&_AssetForwarder.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AssetForwarder *AssetForwarderCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AssetForwarder.Contract.SupportsInterface(&_AssetForwarder.CallOpts, interfaceId)
}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_AssetForwarder *AssetForwarderCaller) TokenMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "tokenMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_AssetForwarder *AssetForwarderSession) TokenMessenger() (common.Address, error) {
	return _AssetForwarder.Contract.TokenMessenger(&_AssetForwarder.CallOpts)
}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_AssetForwarder *AssetForwarderCallerSession) TokenMessenger() (common.Address, error) {
	return _AssetForwarder.Contract.TokenMessenger(&_AssetForwarder.CallOpts)
}

// TotalStakedAmount is a free data retrieval call binding the contract method 0x567e98f9.
//
// Solidity: function totalStakedAmount() view returns(uint256)
func (_AssetForwarder *AssetForwarderCaller) TotalStakedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "totalStakedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakedAmount is a free data retrieval call binding the contract method 0x567e98f9.
//
// Solidity: function totalStakedAmount() view returns(uint256)
func (_AssetForwarder *AssetForwarderSession) TotalStakedAmount() (*big.Int, error) {
	return _AssetForwarder.Contract.TotalStakedAmount(&_AssetForwarder.CallOpts)
}

// TotalStakedAmount is a free data retrieval call binding the contract method 0x567e98f9.
//
// Solidity: function totalStakedAmount() view returns(uint256)
func (_AssetForwarder *AssetForwarderCallerSession) TotalStakedAmount() (*big.Int, error) {
	return _AssetForwarder.Contract.TotalStakedAmount(&_AssetForwarder.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_AssetForwarder *AssetForwarderCaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_AssetForwarder *AssetForwarderSession) Usdc() (common.Address, error) {
	return _AssetForwarder.Contract.Usdc(&_AssetForwarder.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_AssetForwarder *AssetForwarderCallerSession) Usdc() (common.Address, error) {
	return _AssetForwarder.Contract.Usdc(&_AssetForwarder.CallOpts)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_AssetForwarder *AssetForwarderCaller) WrappedNativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "wrappedNativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_AssetForwarder *AssetForwarderSession) WrappedNativeToken() (common.Address, error) {
	return _AssetForwarder.Contract.WrappedNativeToken(&_AssetForwarder.CallOpts)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_AssetForwarder *AssetForwarderCallerSession) WrappedNativeToken() (common.Address, error) {
	return _AssetForwarder.Contract.WrappedNativeToken(&_AssetForwarder.CallOpts)
}

// CommunityPause is a paid mutator transaction binding the contract method 0x6696821b.
//
// Solidity: function communityPause() payable returns()
func (_AssetForwarder *AssetForwarderTransactor) CommunityPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "communityPause")
}

// CommunityPause is a paid mutator transaction binding the contract method 0x6696821b.
//
// Solidity: function communityPause() payable returns()
func (_AssetForwarder *AssetForwarderSession) CommunityPause() (*types.Transaction, error) {
	return _AssetForwarder.Contract.CommunityPause(&_AssetForwarder.TransactOpts)
}

// CommunityPause is a paid mutator transaction binding the contract method 0x6696821b.
//
// Solidity: function communityPause() payable returns()
func (_AssetForwarder *AssetForwarderTransactorSession) CommunityPause() (*types.Transaction, error) {
	return _AssetForwarder.Contract.CommunityPause(&_AssetForwarder.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AssetForwarder *AssetForwarderTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AssetForwarder *AssetForwarderSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetForwarder.Contract.GrantRole(&_AssetForwarder.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AssetForwarder *AssetForwarderTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetForwarder.Contract.GrantRole(&_AssetForwarder.TransactOpts, role, account)
}

// IDeposit is a paid mutator transaction binding the contract method 0xf452ed4d.
//
// Solidity: function iDeposit((uint256,uint256,uint256,address,address,bytes32) depositData, bytes destToken, bytes recipient) payable returns()
func (_AssetForwarder *AssetForwarderTransactor) IDeposit(opts *bind.TransactOpts, depositData IAssetForwarderDepositData, destToken []byte, recipient []byte) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "iDeposit", depositData, destToken, recipient)
}

// IDeposit is a paid mutator transaction binding the contract method 0xf452ed4d.
//
// Solidity: function iDeposit((uint256,uint256,uint256,address,address,bytes32) depositData, bytes destToken, bytes recipient) payable returns()
func (_AssetForwarder *AssetForwarderSession) IDeposit(depositData IAssetForwarderDepositData, destToken []byte, recipient []byte) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IDeposit(&_AssetForwarder.TransactOpts, depositData, destToken, recipient)
}

// IDeposit is a paid mutator transaction binding the contract method 0xf452ed4d.
//
// Solidity: function iDeposit((uint256,uint256,uint256,address,address,bytes32) depositData, bytes destToken, bytes recipient) payable returns()
func (_AssetForwarder *AssetForwarderTransactorSession) IDeposit(depositData IAssetForwarderDepositData, destToken []byte, recipient []byte) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IDeposit(&_AssetForwarder.TransactOpts, depositData, destToken, recipient)
}

// IDepositInfoUpdate is a paid mutator transaction binding the contract method 0xad7c17ee.
//
// Solidity: function iDepositInfoUpdate(address srcToken, uint256 feeAmount, uint256 depositId, bool initiatewithdrawal) payable returns()
func (_AssetForwarder *AssetForwarderTransactor) IDepositInfoUpdate(opts *bind.TransactOpts, srcToken common.Address, feeAmount *big.Int, depositId *big.Int, initiatewithdrawal bool) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "iDepositInfoUpdate", srcToken, feeAmount, depositId, initiatewithdrawal)
}

// IDepositInfoUpdate is a paid mutator transaction binding the contract method 0xad7c17ee.
//
// Solidity: function iDepositInfoUpdate(address srcToken, uint256 feeAmount, uint256 depositId, bool initiatewithdrawal) payable returns()
func (_AssetForwarder *AssetForwarderSession) IDepositInfoUpdate(srcToken common.Address, feeAmount *big.Int, depositId *big.Int, initiatewithdrawal bool) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IDepositInfoUpdate(&_AssetForwarder.TransactOpts, srcToken, feeAmount, depositId, initiatewithdrawal)
}

// IDepositInfoUpdate is a paid mutator transaction binding the contract method 0xad7c17ee.
//
// Solidity: function iDepositInfoUpdate(address srcToken, uint256 feeAmount, uint256 depositId, bool initiatewithdrawal) payable returns()
func (_AssetForwarder *AssetForwarderTransactorSession) IDepositInfoUpdate(srcToken common.Address, feeAmount *big.Int, depositId *big.Int, initiatewithdrawal bool) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IDepositInfoUpdate(&_AssetForwarder.TransactOpts, srcToken, feeAmount, depositId, initiatewithdrawal)
}

// IDepositMessage is a paid mutator transaction binding the contract method 0x0421caf0.
//
// Solidity: function iDepositMessage((uint256,uint256,uint256,address,address,bytes32) depositData, bytes destToken, bytes recipient, bytes message) payable returns()
func (_AssetForwarder *AssetForwarderTransactor) IDepositMessage(opts *bind.TransactOpts, depositData IAssetForwarderDepositData, destToken []byte, recipient []byte, message []byte) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "iDepositMessage", depositData, destToken, recipient, message)
}

// IDepositMessage is a paid mutator transaction binding the contract method 0x0421caf0.
//
// Solidity: function iDepositMessage((uint256,uint256,uint256,address,address,bytes32) depositData, bytes destToken, bytes recipient, bytes message) payable returns()
func (_AssetForwarder *AssetForwarderSession) IDepositMessage(depositData IAssetForwarderDepositData, destToken []byte, recipient []byte, message []byte) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IDepositMessage(&_AssetForwarder.TransactOpts, depositData, destToken, recipient, message)
}

// IDepositMessage is a paid mutator transaction binding the contract method 0x0421caf0.
//
// Solidity: function iDepositMessage((uint256,uint256,uint256,address,address,bytes32) depositData, bytes destToken, bytes recipient, bytes message) payable returns()
func (_AssetForwarder *AssetForwarderTransactorSession) IDepositMessage(depositData IAssetForwarderDepositData, destToken []byte, recipient []byte, message []byte) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IDepositMessage(&_AssetForwarder.TransactOpts, depositData, destToken, recipient, message)
}

// IDepositUSDC is a paid mutator transaction binding the contract method 0x3e28c7d2.
//
// Solidity: function iDepositUSDC(uint256 partnerId, bytes32 destChainIdBytes, bytes32 recipient, uint256 amount) payable returns()
func (_AssetForwarder *AssetForwarderTransactor) IDepositUSDC(opts *bind.TransactOpts, partnerId *big.Int, destChainIdBytes [32]byte, recipient [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "iDepositUSDC", partnerId, destChainIdBytes, recipient, amount)
}

// IDepositUSDC is a paid mutator transaction binding the contract method 0x3e28c7d2.
//
// Solidity: function iDepositUSDC(uint256 partnerId, bytes32 destChainIdBytes, bytes32 recipient, uint256 amount) payable returns()
func (_AssetForwarder *AssetForwarderSession) IDepositUSDC(partnerId *big.Int, destChainIdBytes [32]byte, recipient [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IDepositUSDC(&_AssetForwarder.TransactOpts, partnerId, destChainIdBytes, recipient, amount)
}

// IDepositUSDC is a paid mutator transaction binding the contract method 0x3e28c7d2.
//
// Solidity: function iDepositUSDC(uint256 partnerId, bytes32 destChainIdBytes, bytes32 recipient, uint256 amount) payable returns()
func (_AssetForwarder *AssetForwarderTransactorSession) IDepositUSDC(partnerId *big.Int, destChainIdBytes [32]byte, recipient [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IDepositUSDC(&_AssetForwarder.TransactOpts, partnerId, destChainIdBytes, recipient, amount)
}

// IReceive is a paid mutator transaction binding the contract method 0x1aa6485a.
//
// Solidity: function iReceive(string requestSender, bytes packet, string ) returns(bytes)
func (_AssetForwarder *AssetForwarderTransactor) IReceive(opts *bind.TransactOpts, requestSender string, packet []byte, arg2 string) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "iReceive", requestSender, packet, arg2)
}

// IReceive is a paid mutator transaction binding the contract method 0x1aa6485a.
//
// Solidity: function iReceive(string requestSender, bytes packet, string ) returns(bytes)
func (_AssetForwarder *AssetForwarderSession) IReceive(requestSender string, packet []byte, arg2 string) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IReceive(&_AssetForwarder.TransactOpts, requestSender, packet, arg2)
}

// IReceive is a paid mutator transaction binding the contract method 0x1aa6485a.
//
// Solidity: function iReceive(string requestSender, bytes packet, string ) returns(bytes)
func (_AssetForwarder *AssetForwarderTransactorSession) IReceive(requestSender string, packet []byte, arg2 string) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IReceive(&_AssetForwarder.TransactOpts, requestSender, packet, arg2)
}

// IRelay is a paid mutator transaction binding the contract method 0x64778c1f.
//
// Solidity: function iRelay((uint256,bytes32,uint256,address,address) relayData) payable returns()
func (_AssetForwarder *AssetForwarderTransactor) IRelay(opts *bind.TransactOpts, relayData IAssetForwarderRelayData) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "iRelay", relayData)
}

// IRelay is a paid mutator transaction binding the contract method 0x64778c1f.
//
// Solidity: function iRelay((uint256,bytes32,uint256,address,address) relayData) payable returns()
func (_AssetForwarder *AssetForwarderSession) IRelay(relayData IAssetForwarderRelayData) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IRelay(&_AssetForwarder.TransactOpts, relayData)
}

// IRelay is a paid mutator transaction binding the contract method 0x64778c1f.
//
// Solidity: function iRelay((uint256,bytes32,uint256,address,address) relayData) payable returns()
func (_AssetForwarder *AssetForwarderTransactorSession) IRelay(relayData IAssetForwarderRelayData) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IRelay(&_AssetForwarder.TransactOpts, relayData)
}

// IRelayMessage is a paid mutator transaction binding the contract method 0x6fb003da.
//
// Solidity: function iRelayMessage((uint256,bytes32,uint256,address,address,bytes) relayData) payable returns()
func (_AssetForwarder *AssetForwarderTransactor) IRelayMessage(opts *bind.TransactOpts, relayData IAssetForwarderRelayDataMessage) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "iRelayMessage", relayData)
}

// IRelayMessage is a paid mutator transaction binding the contract method 0x6fb003da.
//
// Solidity: function iRelayMessage((uint256,bytes32,uint256,address,address,bytes) relayData) payable returns()
func (_AssetForwarder *AssetForwarderSession) IRelayMessage(relayData IAssetForwarderRelayDataMessage) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IRelayMessage(&_AssetForwarder.TransactOpts, relayData)
}

// IRelayMessage is a paid mutator transaction binding the contract method 0x6fb003da.
//
// Solidity: function iRelayMessage((uint256,bytes32,uint256,address,address,bytes) relayData) payable returns()
func (_AssetForwarder *AssetForwarderTransactorSession) IRelayMessage(relayData IAssetForwarderRelayDataMessage) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IRelayMessage(&_AssetForwarder.TransactOpts, relayData)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_AssetForwarder *AssetForwarderTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_AssetForwarder *AssetForwarderSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _AssetForwarder.Contract.Multicall(&_AssetForwarder.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_AssetForwarder *AssetForwarderTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _AssetForwarder.Contract.Multicall(&_AssetForwarder.TransactOpts, data)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AssetForwarder *AssetForwarderTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AssetForwarder *AssetForwarderSession) Pause() (*types.Transaction, error) {
	return _AssetForwarder.Contract.Pause(&_AssetForwarder.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AssetForwarder *AssetForwarderTransactorSession) Pause() (*types.Transaction, error) {
	return _AssetForwarder.Contract.Pause(&_AssetForwarder.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AssetForwarder *AssetForwarderTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AssetForwarder *AssetForwarderSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetForwarder.Contract.RenounceRole(&_AssetForwarder.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AssetForwarder *AssetForwarderTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetForwarder.Contract.RenounceRole(&_AssetForwarder.TransactOpts, role, account)
}

// Rescue is a paid mutator transaction binding the contract method 0x839006f2.
//
// Solidity: function rescue(address token) returns()
func (_AssetForwarder *AssetForwarderTransactor) Rescue(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "rescue", token)
}

// Rescue is a paid mutator transaction binding the contract method 0x839006f2.
//
// Solidity: function rescue(address token) returns()
func (_AssetForwarder *AssetForwarderSession) Rescue(token common.Address) (*types.Transaction, error) {
	return _AssetForwarder.Contract.Rescue(&_AssetForwarder.TransactOpts, token)
}

// Rescue is a paid mutator transaction binding the contract method 0x839006f2.
//
// Solidity: function rescue(address token) returns()
func (_AssetForwarder *AssetForwarderTransactorSession) Rescue(token common.Address) (*types.Transaction, error) {
	return _AssetForwarder.Contract.Rescue(&_AssetForwarder.TransactOpts, token)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AssetForwarder *AssetForwarderTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AssetForwarder *AssetForwarderSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetForwarder.Contract.RevokeRole(&_AssetForwarder.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AssetForwarder *AssetForwarderTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetForwarder.Contract.RevokeRole(&_AssetForwarder.TransactOpts, role, account)
}

// SetDestDetails is a paid mutator transaction binding the contract method 0x0171958a.
//
// Solidity: function setDestDetails(bytes32[] _destChainIdBytes, (uint32,uint256,bool)[] _destDetails) returns()
func (_AssetForwarder *AssetForwarderTransactor) SetDestDetails(opts *bind.TransactOpts, _destChainIdBytes [][32]byte, _destDetails []IAssetForwarderDestDetails) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "setDestDetails", _destChainIdBytes, _destDetails)
}

// SetDestDetails is a paid mutator transaction binding the contract method 0x0171958a.
//
// Solidity: function setDestDetails(bytes32[] _destChainIdBytes, (uint32,uint256,bool)[] _destDetails) returns()
func (_AssetForwarder *AssetForwarderSession) SetDestDetails(_destChainIdBytes [][32]byte, _destDetails []IAssetForwarderDestDetails) (*types.Transaction, error) {
	return _AssetForwarder.Contract.SetDestDetails(&_AssetForwarder.TransactOpts, _destChainIdBytes, _destDetails)
}

// SetDestDetails is a paid mutator transaction binding the contract method 0x0171958a.
//
// Solidity: function setDestDetails(bytes32[] _destChainIdBytes, (uint32,uint256,bool)[] _destDetails) returns()
func (_AssetForwarder *AssetForwarderTransactorSession) SetDestDetails(_destChainIdBytes [][32]byte, _destDetails []IAssetForwarderDestDetails) (*types.Transaction, error) {
	return _AssetForwarder.Contract.SetDestDetails(&_AssetForwarder.TransactOpts, _destChainIdBytes, _destDetails)
}

// ToggleCommunityPause is a paid mutator transaction binding the contract method 0xf627df94.
//
// Solidity: function toggleCommunityPause() returns()
func (_AssetForwarder *AssetForwarderTransactor) ToggleCommunityPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "toggleCommunityPause")
}

// ToggleCommunityPause is a paid mutator transaction binding the contract method 0xf627df94.
//
// Solidity: function toggleCommunityPause() returns()
func (_AssetForwarder *AssetForwarderSession) ToggleCommunityPause() (*types.Transaction, error) {
	return _AssetForwarder.Contract.ToggleCommunityPause(&_AssetForwarder.TransactOpts)
}

// ToggleCommunityPause is a paid mutator transaction binding the contract method 0xf627df94.
//
// Solidity: function toggleCommunityPause() returns()
func (_AssetForwarder *AssetForwarderTransactorSession) ToggleCommunityPause() (*types.Transaction, error) {
	return _AssetForwarder.Contract.ToggleCommunityPause(&_AssetForwarder.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AssetForwarder *AssetForwarderTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AssetForwarder *AssetForwarderSession) Unpause() (*types.Transaction, error) {
	return _AssetForwarder.Contract.Unpause(&_AssetForwarder.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AssetForwarder *AssetForwarderTransactorSession) Unpause() (*types.Transaction, error) {
	return _AssetForwarder.Contract.Unpause(&_AssetForwarder.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0x5ac62700.
//
// Solidity: function update(uint256 index, address _gatewayContract, bytes _routerMiddlewareBase, uint256 minPauseStakeAmount, uint256 maxPauseStakeAmount) returns()
func (_AssetForwarder *AssetForwarderTransactor) Update(opts *bind.TransactOpts, index *big.Int, _gatewayContract common.Address, _routerMiddlewareBase []byte, minPauseStakeAmount *big.Int, maxPauseStakeAmount *big.Int) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "update", index, _gatewayContract, _routerMiddlewareBase, minPauseStakeAmount, maxPauseStakeAmount)
}

// Update is a paid mutator transaction binding the contract method 0x5ac62700.
//
// Solidity: function update(uint256 index, address _gatewayContract, bytes _routerMiddlewareBase, uint256 minPauseStakeAmount, uint256 maxPauseStakeAmount) returns()
func (_AssetForwarder *AssetForwarderSession) Update(index *big.Int, _gatewayContract common.Address, _routerMiddlewareBase []byte, minPauseStakeAmount *big.Int, maxPauseStakeAmount *big.Int) (*types.Transaction, error) {
	return _AssetForwarder.Contract.Update(&_AssetForwarder.TransactOpts, index, _gatewayContract, _routerMiddlewareBase, minPauseStakeAmount, maxPauseStakeAmount)
}

// Update is a paid mutator transaction binding the contract method 0x5ac62700.
//
// Solidity: function update(uint256 index, address _gatewayContract, bytes _routerMiddlewareBase, uint256 minPauseStakeAmount, uint256 maxPauseStakeAmount) returns()
func (_AssetForwarder *AssetForwarderTransactorSession) Update(index *big.Int, _gatewayContract common.Address, _routerMiddlewareBase []byte, minPauseStakeAmount *big.Int, maxPauseStakeAmount *big.Int) (*types.Transaction, error) {
	return _AssetForwarder.Contract.Update(&_AssetForwarder.TransactOpts, index, _gatewayContract, _routerMiddlewareBase, minPauseStakeAmount, maxPauseStakeAmount)
}

// UpdateTokenMessenger is a paid mutator transaction binding the contract method 0x78e0f9bd.
//
// Solidity: function updateTokenMessenger(address _tokenMessenger) returns()
func (_AssetForwarder *AssetForwarderTransactor) UpdateTokenMessenger(opts *bind.TransactOpts, _tokenMessenger common.Address) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "updateTokenMessenger", _tokenMessenger)
}

// UpdateTokenMessenger is a paid mutator transaction binding the contract method 0x78e0f9bd.
//
// Solidity: function updateTokenMessenger(address _tokenMessenger) returns()
func (_AssetForwarder *AssetForwarderSession) UpdateTokenMessenger(_tokenMessenger common.Address) (*types.Transaction, error) {
	return _AssetForwarder.Contract.UpdateTokenMessenger(&_AssetForwarder.TransactOpts, _tokenMessenger)
}

// UpdateTokenMessenger is a paid mutator transaction binding the contract method 0x78e0f9bd.
//
// Solidity: function updateTokenMessenger(address _tokenMessenger) returns()
func (_AssetForwarder *AssetForwarderTransactorSession) UpdateTokenMessenger(_tokenMessenger common.Address) (*types.Transaction, error) {
	return _AssetForwarder.Contract.UpdateTokenMessenger(&_AssetForwarder.TransactOpts, _tokenMessenger)
}

// WithdrawStakeAmount is a paid mutator transaction binding the contract method 0x8a27fecb.
//
// Solidity: function withdrawStakeAmount() returns()
func (_AssetForwarder *AssetForwarderTransactor) WithdrawStakeAmount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "withdrawStakeAmount")
}

// WithdrawStakeAmount is a paid mutator transaction binding the contract method 0x8a27fecb.
//
// Solidity: function withdrawStakeAmount() returns()
func (_AssetForwarder *AssetForwarderSession) WithdrawStakeAmount() (*types.Transaction, error) {
	return _AssetForwarder.Contract.WithdrawStakeAmount(&_AssetForwarder.TransactOpts)
}

// WithdrawStakeAmount is a paid mutator transaction binding the contract method 0x8a27fecb.
//
// Solidity: function withdrawStakeAmount() returns()
func (_AssetForwarder *AssetForwarderTransactorSession) WithdrawStakeAmount() (*types.Transaction, error) {
	return _AssetForwarder.Contract.WithdrawStakeAmount(&_AssetForwarder.TransactOpts)
}

// AssetForwarderCommunityPausedIterator is returned from FilterCommunityPaused and is used to iterate over the raw logs and unpacked data for CommunityPaused events raised by the AssetForwarder contract.
type AssetForwarderCommunityPausedIterator struct {
	Event *AssetForwarderCommunityPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetForwarderCommunityPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderCommunityPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetForwarderCommunityPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetForwarderCommunityPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderCommunityPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderCommunityPaused represents a CommunityPaused event raised by the AssetForwarder contract.
type AssetForwarderCommunityPaused struct {
	Pauser       common.Address
	StakedAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCommunityPaused is a free log retrieval operation binding the contract event 0x9593b43c20e09177a4170170ac564123ad8138e040e21eec96d1ae9db9ee5d6d.
//
// Solidity: event CommunityPaused(address indexed pauser, uint256 stakedAmount)
func (_AssetForwarder *AssetForwarderFilterer) FilterCommunityPaused(opts *bind.FilterOpts, pauser []common.Address) (*AssetForwarderCommunityPausedIterator, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "CommunityPaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return &AssetForwarderCommunityPausedIterator{contract: _AssetForwarder.contract, event: "CommunityPaused", logs: logs, sub: sub}, nil
}

// WatchCommunityPaused is a free log subscription operation binding the contract event 0x9593b43c20e09177a4170170ac564123ad8138e040e21eec96d1ae9db9ee5d6d.
//
// Solidity: event CommunityPaused(address indexed pauser, uint256 stakedAmount)
func (_AssetForwarder *AssetForwarderFilterer) WatchCommunityPaused(opts *bind.WatchOpts, sink chan<- *AssetForwarderCommunityPaused, pauser []common.Address) (event.Subscription, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "CommunityPaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderCommunityPaused)
				if err := _AssetForwarder.contract.UnpackLog(event, "CommunityPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommunityPaused is a log parse operation binding the contract event 0x9593b43c20e09177a4170170ac564123ad8138e040e21eec96d1ae9db9ee5d6d.
//
// Solidity: event CommunityPaused(address indexed pauser, uint256 stakedAmount)
func (_AssetForwarder *AssetForwarderFilterer) ParseCommunityPaused(log types.Log) (*AssetForwarderCommunityPaused, error) {
	event := new(AssetForwarderCommunityPaused)
	if err := _AssetForwarder.contract.UnpackLog(event, "CommunityPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderDepositInfoUpdateIterator is returned from FilterDepositInfoUpdate and is used to iterate over the raw logs and unpacked data for DepositInfoUpdate events raised by the AssetForwarder contract.
type AssetForwarderDepositInfoUpdateIterator struct {
	Event *AssetForwarderDepositInfoUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetForwarderDepositInfoUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderDepositInfoUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetForwarderDepositInfoUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetForwarderDepositInfoUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderDepositInfoUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderDepositInfoUpdate represents a DepositInfoUpdate event raised by the AssetForwarder contract.
type AssetForwarderDepositInfoUpdate struct {
	SrcToken           common.Address
	FeeAmount          *big.Int
	DepositId          *big.Int
	EventNonce         *big.Int
	Initiatewithdrawal bool
	Depositor          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDepositInfoUpdate is a free log retrieval operation binding the contract event 0x86896302632bf6dc8a3ac0ae7ddf17d5a5d5c1ca1aad37b4b920a587c51135b1.
//
// Solidity: event DepositInfoUpdate(address srcToken, uint256 feeAmount, uint256 depositId, uint256 eventNonce, bool initiatewithdrawal, address depositor)
func (_AssetForwarder *AssetForwarderFilterer) FilterDepositInfoUpdate(opts *bind.FilterOpts) (*AssetForwarderDepositInfoUpdateIterator, error) {

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "DepositInfoUpdate")
	if err != nil {
		return nil, err
	}
	return &AssetForwarderDepositInfoUpdateIterator{contract: _AssetForwarder.contract, event: "DepositInfoUpdate", logs: logs, sub: sub}, nil
}

// WatchDepositInfoUpdate is a free log subscription operation binding the contract event 0x86896302632bf6dc8a3ac0ae7ddf17d5a5d5c1ca1aad37b4b920a587c51135b1.
//
// Solidity: event DepositInfoUpdate(address srcToken, uint256 feeAmount, uint256 depositId, uint256 eventNonce, bool initiatewithdrawal, address depositor)
func (_AssetForwarder *AssetForwarderFilterer) WatchDepositInfoUpdate(opts *bind.WatchOpts, sink chan<- *AssetForwarderDepositInfoUpdate) (event.Subscription, error) {

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "DepositInfoUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderDepositInfoUpdate)
				if err := _AssetForwarder.contract.UnpackLog(event, "DepositInfoUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositInfoUpdate is a log parse operation binding the contract event 0x86896302632bf6dc8a3ac0ae7ddf17d5a5d5c1ca1aad37b4b920a587c51135b1.
//
// Solidity: event DepositInfoUpdate(address srcToken, uint256 feeAmount, uint256 depositId, uint256 eventNonce, bool initiatewithdrawal, address depositor)
func (_AssetForwarder *AssetForwarderFilterer) ParseDepositInfoUpdate(log types.Log) (*AssetForwarderDepositInfoUpdate, error) {
	event := new(AssetForwarderDepositInfoUpdate)
	if err := _AssetForwarder.contract.UnpackLog(event, "DepositInfoUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderFundsDepositedIterator is returned from FilterFundsDeposited and is used to iterate over the raw logs and unpacked data for FundsDeposited events raised by the AssetForwarder contract.
type AssetForwarderFundsDepositedIterator struct {
	Event *AssetForwarderFundsDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetForwarderFundsDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderFundsDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetForwarderFundsDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetForwarderFundsDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderFundsDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderFundsDeposited represents a FundsDeposited event raised by the AssetForwarder contract.
type AssetForwarderFundsDeposited struct {
	PartnerId        *big.Int
	Amount           *big.Int
	DestChainIdBytes [32]byte
	DestAmount       *big.Int
	DepositId        *big.Int
	SrcToken         common.Address
	Depositor        common.Address
	Recipient        []byte
	DestToken        []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFundsDeposited is a free log retrieval operation binding the contract event 0x6f223106c8e3df857d691613d18d1478cc7c629a1fdf16c7b461d36729fcc7ad.
//
// Solidity: event FundsDeposited(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 destAmount, uint256 depositId, address srcToken, address depositor, bytes recipient, bytes destToken)
func (_AssetForwarder *AssetForwarderFilterer) FilterFundsDeposited(opts *bind.FilterOpts) (*AssetForwarderFundsDepositedIterator, error) {

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "FundsDeposited")
	if err != nil {
		return nil, err
	}
	return &AssetForwarderFundsDepositedIterator{contract: _AssetForwarder.contract, event: "FundsDeposited", logs: logs, sub: sub}, nil
}

// WatchFundsDeposited is a free log subscription operation binding the contract event 0x6f223106c8e3df857d691613d18d1478cc7c629a1fdf16c7b461d36729fcc7ad.
//
// Solidity: event FundsDeposited(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 destAmount, uint256 depositId, address srcToken, address depositor, bytes recipient, bytes destToken)
func (_AssetForwarder *AssetForwarderFilterer) WatchFundsDeposited(opts *bind.WatchOpts, sink chan<- *AssetForwarderFundsDeposited) (event.Subscription, error) {

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "FundsDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderFundsDeposited)
				if err := _AssetForwarder.contract.UnpackLog(event, "FundsDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsDeposited is a log parse operation binding the contract event 0x6f223106c8e3df857d691613d18d1478cc7c629a1fdf16c7b461d36729fcc7ad.
//
// Solidity: event FundsDeposited(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 destAmount, uint256 depositId, address srcToken, address depositor, bytes recipient, bytes destToken)
func (_AssetForwarder *AssetForwarderFilterer) ParseFundsDeposited(log types.Log) (*AssetForwarderFundsDeposited, error) {
	event := new(AssetForwarderFundsDeposited)
	if err := _AssetForwarder.contract.UnpackLog(event, "FundsDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderFundsDepositedWithMessageIterator is returned from FilterFundsDepositedWithMessage and is used to iterate over the raw logs and unpacked data for FundsDepositedWithMessage events raised by the AssetForwarder contract.
type AssetForwarderFundsDepositedWithMessageIterator struct {
	Event *AssetForwarderFundsDepositedWithMessage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetForwarderFundsDepositedWithMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderFundsDepositedWithMessage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetForwarderFundsDepositedWithMessage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetForwarderFundsDepositedWithMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderFundsDepositedWithMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderFundsDepositedWithMessage represents a FundsDepositedWithMessage event raised by the AssetForwarder contract.
type AssetForwarderFundsDepositedWithMessage struct {
	PartnerId        *big.Int
	Amount           *big.Int
	DestChainIdBytes [32]byte
	DestAmount       *big.Int
	DepositId        *big.Int
	SrcToken         common.Address
	Recipient        []byte
	Depositor        common.Address
	DestToken        []byte
	Message          []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFundsDepositedWithMessage is a free log retrieval operation binding the contract event 0x3dbc28a2fa93575c89d951d683c45ddb951a2ecf6bc9b9704a61589fa0fcb70f.
//
// Solidity: event FundsDepositedWithMessage(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 destAmount, uint256 depositId, address srcToken, bytes recipient, address depositor, bytes destToken, bytes message)
func (_AssetForwarder *AssetForwarderFilterer) FilterFundsDepositedWithMessage(opts *bind.FilterOpts) (*AssetForwarderFundsDepositedWithMessageIterator, error) {

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "FundsDepositedWithMessage")
	if err != nil {
		return nil, err
	}
	return &AssetForwarderFundsDepositedWithMessageIterator{contract: _AssetForwarder.contract, event: "FundsDepositedWithMessage", logs: logs, sub: sub}, nil
}

// WatchFundsDepositedWithMessage is a free log subscription operation binding the contract event 0x3dbc28a2fa93575c89d951d683c45ddb951a2ecf6bc9b9704a61589fa0fcb70f.
//
// Solidity: event FundsDepositedWithMessage(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 destAmount, uint256 depositId, address srcToken, bytes recipient, address depositor, bytes destToken, bytes message)
func (_AssetForwarder *AssetForwarderFilterer) WatchFundsDepositedWithMessage(opts *bind.WatchOpts, sink chan<- *AssetForwarderFundsDepositedWithMessage) (event.Subscription, error) {

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "FundsDepositedWithMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderFundsDepositedWithMessage)
				if err := _AssetForwarder.contract.UnpackLog(event, "FundsDepositedWithMessage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsDepositedWithMessage is a log parse operation binding the contract event 0x3dbc28a2fa93575c89d951d683c45ddb951a2ecf6bc9b9704a61589fa0fcb70f.
//
// Solidity: event FundsDepositedWithMessage(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 destAmount, uint256 depositId, address srcToken, bytes recipient, address depositor, bytes destToken, bytes message)
func (_AssetForwarder *AssetForwarderFilterer) ParseFundsDepositedWithMessage(log types.Log) (*AssetForwarderFundsDepositedWithMessage, error) {
	event := new(AssetForwarderFundsDepositedWithMessage)
	if err := _AssetForwarder.contract.UnpackLog(event, "FundsDepositedWithMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderFundsPaidIterator is returned from FilterFundsPaid and is used to iterate over the raw logs and unpacked data for FundsPaid events raised by the AssetForwarder contract.
type AssetForwarderFundsPaidIterator struct {
	Event *AssetForwarderFundsPaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetForwarderFundsPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderFundsPaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetForwarderFundsPaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetForwarderFundsPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderFundsPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderFundsPaid represents a FundsPaid event raised by the AssetForwarder contract.
type AssetForwarderFundsPaid struct {
	MessageHash [32]byte
	Forwarder   common.Address
	Nonce       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFundsPaid is a free log retrieval operation binding the contract event 0x0f3ca0b27903ec13ef88a7ea8be837cc19b0d7f71a735f2083215739a8004464.
//
// Solidity: event FundsPaid(bytes32 messageHash, address forwarder, uint256 nonce)
func (_AssetForwarder *AssetForwarderFilterer) FilterFundsPaid(opts *bind.FilterOpts) (*AssetForwarderFundsPaidIterator, error) {

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "FundsPaid")
	if err != nil {
		return nil, err
	}
	return &AssetForwarderFundsPaidIterator{contract: _AssetForwarder.contract, event: "FundsPaid", logs: logs, sub: sub}, nil
}

// WatchFundsPaid is a free log subscription operation binding the contract event 0x0f3ca0b27903ec13ef88a7ea8be837cc19b0d7f71a735f2083215739a8004464.
//
// Solidity: event FundsPaid(bytes32 messageHash, address forwarder, uint256 nonce)
func (_AssetForwarder *AssetForwarderFilterer) WatchFundsPaid(opts *bind.WatchOpts, sink chan<- *AssetForwarderFundsPaid) (event.Subscription, error) {

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "FundsPaid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderFundsPaid)
				if err := _AssetForwarder.contract.UnpackLog(event, "FundsPaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsPaid is a log parse operation binding the contract event 0x0f3ca0b27903ec13ef88a7ea8be837cc19b0d7f71a735f2083215739a8004464.
//
// Solidity: event FundsPaid(bytes32 messageHash, address forwarder, uint256 nonce)
func (_AssetForwarder *AssetForwarderFilterer) ParseFundsPaid(log types.Log) (*AssetForwarderFundsPaid, error) {
	event := new(AssetForwarderFundsPaid)
	if err := _AssetForwarder.contract.UnpackLog(event, "FundsPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderFundsPaidWithMessageIterator is returned from FilterFundsPaidWithMessage and is used to iterate over the raw logs and unpacked data for FundsPaidWithMessage events raised by the AssetForwarder contract.
type AssetForwarderFundsPaidWithMessageIterator struct {
	Event *AssetForwarderFundsPaidWithMessage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetForwarderFundsPaidWithMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderFundsPaidWithMessage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetForwarderFundsPaidWithMessage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetForwarderFundsPaidWithMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderFundsPaidWithMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderFundsPaidWithMessage represents a FundsPaidWithMessage event raised by the AssetForwarder contract.
type AssetForwarderFundsPaidWithMessage struct {
	MessageHash [32]byte
	Forwarder   common.Address
	Nonce       *big.Int
	ExecFlag    bool
	ExecData    []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFundsPaidWithMessage is a free log retrieval operation binding the contract event 0x21937deaa62558dad619c8d730a7d1d7ef41731fc194c32973511e1455cb37ad.
//
// Solidity: event FundsPaidWithMessage(bytes32 messageHash, address forwarder, uint256 nonce, bool execFlag, bytes execData)
func (_AssetForwarder *AssetForwarderFilterer) FilterFundsPaidWithMessage(opts *bind.FilterOpts) (*AssetForwarderFundsPaidWithMessageIterator, error) {

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "FundsPaidWithMessage")
	if err != nil {
		return nil, err
	}
	return &AssetForwarderFundsPaidWithMessageIterator{contract: _AssetForwarder.contract, event: "FundsPaidWithMessage", logs: logs, sub: sub}, nil
}

// WatchFundsPaidWithMessage is a free log subscription operation binding the contract event 0x21937deaa62558dad619c8d730a7d1d7ef41731fc194c32973511e1455cb37ad.
//
// Solidity: event FundsPaidWithMessage(bytes32 messageHash, address forwarder, uint256 nonce, bool execFlag, bytes execData)
func (_AssetForwarder *AssetForwarderFilterer) WatchFundsPaidWithMessage(opts *bind.WatchOpts, sink chan<- *AssetForwarderFundsPaidWithMessage) (event.Subscription, error) {

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "FundsPaidWithMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderFundsPaidWithMessage)
				if err := _AssetForwarder.contract.UnpackLog(event, "FundsPaidWithMessage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsPaidWithMessage is a log parse operation binding the contract event 0x21937deaa62558dad619c8d730a7d1d7ef41731fc194c32973511e1455cb37ad.
//
// Solidity: event FundsPaidWithMessage(bytes32 messageHash, address forwarder, uint256 nonce, bool execFlag, bytes execData)
func (_AssetForwarder *AssetForwarderFilterer) ParseFundsPaidWithMessage(log types.Log) (*AssetForwarderFundsPaidWithMessage, error) {
	event := new(AssetForwarderFundsPaidWithMessage)
	if err := _AssetForwarder.contract.UnpackLog(event, "FundsPaidWithMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AssetForwarder contract.
type AssetForwarderPausedIterator struct {
	Event *AssetForwarderPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetForwarderPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetForwarderPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetForwarderPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderPaused represents a Paused event raised by the AssetForwarder contract.
type AssetForwarderPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AssetForwarder *AssetForwarderFilterer) FilterPaused(opts *bind.FilterOpts) (*AssetForwarderPausedIterator, error) {

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AssetForwarderPausedIterator{contract: _AssetForwarder.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AssetForwarder *AssetForwarderFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AssetForwarderPaused) (event.Subscription, error) {

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderPaused)
				if err := _AssetForwarder.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AssetForwarder *AssetForwarderFilterer) ParsePaused(log types.Log) (*AssetForwarderPaused, error) {
	event := new(AssetForwarderPaused)
	if err := _AssetForwarder.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AssetForwarder contract.
type AssetForwarderRoleAdminChangedIterator struct {
	Event *AssetForwarderRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetForwarderRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetForwarderRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetForwarderRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderRoleAdminChanged represents a RoleAdminChanged event raised by the AssetForwarder contract.
type AssetForwarderRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AssetForwarder *AssetForwarderFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AssetForwarderRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AssetForwarderRoleAdminChangedIterator{contract: _AssetForwarder.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AssetForwarder *AssetForwarderFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AssetForwarderRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderRoleAdminChanged)
				if err := _AssetForwarder.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AssetForwarder *AssetForwarderFilterer) ParseRoleAdminChanged(log types.Log) (*AssetForwarderRoleAdminChanged, error) {
	event := new(AssetForwarderRoleAdminChanged)
	if err := _AssetForwarder.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AssetForwarder contract.
type AssetForwarderRoleGrantedIterator struct {
	Event *AssetForwarderRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetForwarderRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetForwarderRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetForwarderRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderRoleGranted represents a RoleGranted event raised by the AssetForwarder contract.
type AssetForwarderRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetForwarder *AssetForwarderFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AssetForwarderRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AssetForwarderRoleGrantedIterator{contract: _AssetForwarder.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetForwarder *AssetForwarderFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AssetForwarderRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderRoleGranted)
				if err := _AssetForwarder.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetForwarder *AssetForwarderFilterer) ParseRoleGranted(log types.Log) (*AssetForwarderRoleGranted, error) {
	event := new(AssetForwarderRoleGranted)
	if err := _AssetForwarder.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AssetForwarder contract.
type AssetForwarderRoleRevokedIterator struct {
	Event *AssetForwarderRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetForwarderRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetForwarderRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetForwarderRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderRoleRevoked represents a RoleRevoked event raised by the AssetForwarder contract.
type AssetForwarderRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetForwarder *AssetForwarderFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AssetForwarderRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AssetForwarderRoleRevokedIterator{contract: _AssetForwarder.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetForwarder *AssetForwarderFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AssetForwarderRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderRoleRevoked)
				if err := _AssetForwarder.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetForwarder *AssetForwarderFilterer) ParseRoleRevoked(log types.Log) (*AssetForwarderRoleRevoked, error) {
	event := new(AssetForwarderRoleRevoked)
	if err := _AssetForwarder.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AssetForwarder contract.
type AssetForwarderUnpausedIterator struct {
	Event *AssetForwarderUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetForwarderUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetForwarderUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetForwarderUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderUnpaused represents a Unpaused event raised by the AssetForwarder contract.
type AssetForwarderUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AssetForwarder *AssetForwarderFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AssetForwarderUnpausedIterator, error) {

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AssetForwarderUnpausedIterator{contract: _AssetForwarder.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AssetForwarder *AssetForwarderFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AssetForwarderUnpaused) (event.Subscription, error) {

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderUnpaused)
				if err := _AssetForwarder.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AssetForwarder *AssetForwarderFilterer) ParseUnpaused(log types.Log) (*AssetForwarderUnpaused, error) {
	event := new(AssetForwarderUnpaused)
	if err := _AssetForwarder.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderIUSDCDepositedIterator is returned from FilterIUSDCDeposited and is used to iterate over the raw logs and unpacked data for IUSDCDeposited events raised by the AssetForwarder contract.
type AssetForwarderIUSDCDepositedIterator struct {
	Event *AssetForwarderIUSDCDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetForwarderIUSDCDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderIUSDCDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetForwarderIUSDCDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetForwarderIUSDCDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderIUSDCDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderIUSDCDeposited represents a IUSDCDeposited event raised by the AssetForwarder contract.
type AssetForwarderIUSDCDeposited struct {
	PartnerId        *big.Int
	Amount           *big.Int
	DestChainIdBytes [32]byte
	UsdcNonce        *big.Int
	SrcToken         common.Address
	Recipient        [32]byte
	Depositor        common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterIUSDCDeposited is a free log retrieval operation binding the contract event 0x297a8bc8b87367a63661d6429dbab51be5cefd71ce6a3050fa900a8f276d66d9.
//
// Solidity: event iUSDCDeposited(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 usdcNonce, address srcToken, bytes32 recipient, address depositor)
func (_AssetForwarder *AssetForwarderFilterer) FilterIUSDCDeposited(opts *bind.FilterOpts) (*AssetForwarderIUSDCDepositedIterator, error) {

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "iUSDCDeposited")
	if err != nil {
		return nil, err
	}
	return &AssetForwarderIUSDCDepositedIterator{contract: _AssetForwarder.contract, event: "iUSDCDeposited", logs: logs, sub: sub}, nil
}

// WatchIUSDCDeposited is a free log subscription operation binding the contract event 0x297a8bc8b87367a63661d6429dbab51be5cefd71ce6a3050fa900a8f276d66d9.
//
// Solidity: event iUSDCDeposited(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 usdcNonce, address srcToken, bytes32 recipient, address depositor)
func (_AssetForwarder *AssetForwarderFilterer) WatchIUSDCDeposited(opts *bind.WatchOpts, sink chan<- *AssetForwarderIUSDCDeposited) (event.Subscription, error) {

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "iUSDCDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderIUSDCDeposited)
				if err := _AssetForwarder.contract.UnpackLog(event, "iUSDCDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIUSDCDeposited is a log parse operation binding the contract event 0x297a8bc8b87367a63661d6429dbab51be5cefd71ce6a3050fa900a8f276d66d9.
//
// Solidity: event iUSDCDeposited(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 usdcNonce, address srcToken, bytes32 recipient, address depositor)
func (_AssetForwarder *AssetForwarderFilterer) ParseIUSDCDeposited(log types.Log) (*AssetForwarderIUSDCDeposited, error) {
	event := new(AssetForwarderIUSDCDeposited)
	if err := _AssetForwarder.contract.UnpackLog(event, "iUSDCDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
