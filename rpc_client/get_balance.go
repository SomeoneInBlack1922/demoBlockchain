package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func get_balance() {
	if len(os.Args) < 3 {
		fmt.Println("Недостатньо аргументів для операції get_balance. Треба 1 аргумент - публічна адреса акаунта")
		return
	}
	publicKeyBytes, err := hex.DecodeString(os.Args[2])
	error_stop(err, "Читання адреси акаунта")
	userAdress := common.BytesToAddress(publicKeyBytes)
	gotClient, err := ethclient.Dial(clientAdressString)
	error_stop(err, "зв'язку з блокчейном")
	lastBlockNumber, err := gotClient.BlockNumber(emptyContext)
	error_stop(err, "видобування номера останнього блока")
	bigIntBlockNumber := new(big.Int).SetUint64(lastBlockNumber)
	userBlance, err := gotClient.BalanceAt(emptyContext, userAdress, bigIntBlockNumber)
	error_stop(err, "читання баланса користувача")
	fmt.Printf("Баланс за цією адрсою: %d", userBlance)
}
