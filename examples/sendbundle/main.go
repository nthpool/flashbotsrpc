package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/metachris/flashbotsrpc"
)

var privateKey, _ = crypto.GenerateKey() // creating a new private key for testing. you probably want to use an existing key.

func main() {
	rpc := flashbotsrpc.New("https://relay.flashbots.net")
	rpc.Debug = true

	sendBundleArgs := flashbotsrpc.FlashbotsSendBundleRequest{
		Txs:         []string{"YOUR_HASH"},
		BlockNumber: fmt.Sprintf("0x%x", 13281018),
	}

	result, err := rpc.FlashbotsSendBundle(privateKey, sendBundleArgs)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	// Print result
	fmt.Printf("%+v\n", result)
}
