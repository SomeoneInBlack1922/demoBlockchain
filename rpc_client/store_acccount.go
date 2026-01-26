package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func store_account() {
	if len(os.Args) < 4 {
		fmt.Println("Операція store_account потребує двох параметрів - шляху папки куди буде збережено keystore файл що буде збережено та приватного ключа акаунта")
		return
	}
	keystoreFileDir := os.Args[2]
	//Отримати шлях папки у яку треба буде покласти згенерований keystore
	// iterator := len(keystoreFilePath)
	// for iterator > 0 {
	// 	if keystoreFilePath[iterator-1] == '/' {
	// 		keystoreFileDir = keystoreFilePath[:iterator]
	// 		iterator = -1
	// 		continue
	// 	}
	// 	iterator -= 1
	// }
	// if iterator == 0 {
	// 	keystoreFileDir = "./"
	// 	// keystoreFilePath = "./" + keystoreFilePath
	// }
	gotKeystore := keystore.NewKeyStore(keystoreFileDir, keystore.StandardScryptN, keystore.StandardScryptP)
	privateKeyBytes, err := hex.DecodeString(os.Args[3])
	error_stop(err, "читання шістнідцяткового рядка приватного ключа")
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	error_stop(err, "Створення типпу данних ключа із рядка байтів приватного ключа")
	gotKeystore.ImportECDSA(privateKey, "")
	fmt.Println("Збережено")
}
