package main

import (
	"context"
	"fmt"
	"os"
)

var emptyContext context.Context // Порожній контекст бо його вимагають методи ефіра
// var privateKey ecdsa.PrivateKey  // Публічний ключ що буде використаний для кожної операції
// var publicKey common.Address     //Приватний ключ що буде використаний для кожної операції
func error_stop(err error, text string) {
	if err != nil {
		panic(fmt.Sprintf("Сталася помилка: %v під час %v", err, text))
	}
}

var clientAdressString string

func main() {
	emptyContext = context.Background() // Порожній контекст бо його вимагають методи ефіра
	clientAdressString = "http://localhost:8545"

	//Крок 1 - читання операції і прийняття рішення що робити далі
	if len(os.Args) < 2 {
		fmt.Println("Ви не ввели операцію - вона має бути першим аргументом")
		return
	}
	switch operation := os.Args[1]; operation {
	case "help":
		help()
		return
	case "store_account":
		store_account()
		return
	case "get_balance":
		get_balance()
	default:
		fmt.Printf("Операції %v нема. Якби мені не було лінь я би тут вивів список доступних операції. Але мені лінь\n", os.Args[1])
	}

	// keystoreFileRead, err := os.ReadFile(os.Args[1])
	// error_stop(err, "Читання keystore файла")

	// account, err := keystore.DecryptKey(keystoreFileRead, "")
	// error_stop(err, "декодування файла keystore")

	// fmt.Printf("Приватний ключ %x\nПублічний ключ %x\nОперація %v\n", *account.PrivateKey, account.Address, os.Args[2])
	// fmt.Printf("Arguments: %v\n", os.Args[1:])

	// gotClient, err := ethclient.Dial("http://localhost:8545")
	// if err != nil {
	// 	fmt.Println("Failed to connect to local chain")
	// 	return
	// }
	// blockNumber, err := gotClient.BlockNumber(context)
	// if err != nil {
	// 	fmt.Printf("Failed to get block number: %v\n", err)
	// 	return
	// }
	// bigIntBlockNumber := new(big.Int).SetUint64(blockNumber)
	// // fmt.Printf("%v\n", blockNumber)
	// metasBalance, err := gotClient.BalanceAt(context, common.HexToAddress("0x0D334543139D501F1aC04435f236f179c9366ab6"), bigIntBlockNumber)
	// if err != nil {
	// 	fmt.Printf("Failed to get balance: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Balance of Metas: %v\n", metasBalance)
}
