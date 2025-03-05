package helpers

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	u2u "github.com/unicornultrafoundation/go-u2u"
	"github.com/unicornultrafoundation/go-u2u/accounts/abi"
	"github.com/unicornultrafoundation/go-u2u/common"
	"github.com/unicornultrafoundation/go-u2u/ethclient"
	"github.com/unicornultrafoundation/go-u2u/rpc"
)

const SIMPLESTORAGE = `[
  {
    "inputs": [],
    "name": "getValue",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "sender",
    "outputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
]`

func InitChainClient(url string) (*ethclient.Client, error) {
	return ethclient.Dial(url)
}

func InitNewRPCClient(url string) (*rpc.Client, error) {
	return rpc.Dial(url)
}

func SIMPLESTORAGEGetValue(ctx context.Context, client *ethclient.Client,
	contractAddress *common.Address) (*big.Int, error) {
	// Create ABI instance
	contractABI, err := abi.JSON(strings.NewReader(SIMPLESTORAGE))
	if err != nil {
		fmt.Println("err", err)

		return nil, err
	}

	// Prepare the function call data
	data, err := contractABI.Pack("getValue")
	if err != nil {

		return nil, err
	}

	// Call the contract
	msg := u2u.CallMsg{
		To:   contractAddress,
		Data: data,
	}

	// Execute the call
	result, err := client.CallContract(ctx, msg, nil)
	if err != nil {

		return nil, err
	}

	// Unpack the result
	var unpackVar *big.Int
	err = contractABI.UnpackIntoInterface(&unpackVar, "getValue", result)
	if err != nil {
		fmt.Println("err", err)

		return nil, err
	}
	return unpackVar, nil
}
