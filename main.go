package main

import (
	"asset-query/internal/helpers"
	"context"
	"fmt"

	"github.com/unicornultrafoundation/go-u2u/common"
)

func main() {

	ctx := context.Background()

	client, err := helpers.InitChainClient("https://rpc-nebulas-testnet.uniultra.xyz")
	if err != nil {
		fmt.Println(err)
	}

	contractAddress := common.HexToAddress("0xa595cd76B7705D5795C10406B9e1D97BF3779861")

	testHelper, err := helpers.SIMPLESTORAGEGetValue(ctx, client, &contractAddress)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("testHelper", testHelper)

}
