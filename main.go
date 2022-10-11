package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	connection, err := ethclient.Dial("https://mainnet.infura.io/v3/03bf216d226d47adb501d642f9dc0c12")

	if err != nil {
		log.Fatal("Failed to make a connection to the Ethereum Node", err)
	}

	context := context.Background()
	txn, pending, _ := connection.TransactionByHash(context, common.HexToHash("0xbca59819540b55e32a9a422e44559df023c2951b97b50815fa029b67ffe2a4a7"))
	if pending {
		fmt.Println("The transaction is pending", txn)
	} else {
		fmt.Println("Transaction is not pending", txn)
	}
}
