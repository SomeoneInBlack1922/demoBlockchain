package main

import (
	"blockchain_m/bch"
	"crypto/sha256"

	// "bytes"
	// "math"
	// "encoding/binary"
	"fmt"
	"time"
)

func main() {
	//Створення генезис блока
	fmt.Println("Generating genezis block...")
	mySacondnameInHash := [sha256.Size]byte{'D', 'a', 't', 's', 'k', 'o'}
	genBlock := bch.GetGenesisBlock(mySacondnameInHash, uint64(time.Now().Unix()))
	fmt.Printf("Got genesis block:\n")
	genBlock.PrintInfo("\t")
	//Перевірка генезис блока функцією ValidBlock()
	if genBlock.GetHash() == genBlock.Hash {
		fmt.Printf("Genesis block has been validated!\n")
	} else {
		fmt.Printf("Validation of genesis block has failed\n")
	}
	//Транзакції для прикладу з яких я створю блок
	trans1, trans2 :=
		bch.Transaction{Sender: "Mikolai", Reciver: "Andriy", Amount: 12},
		bch.Transaction{Sender: "Sergiy", Reciver: "Anton", Amount: 3}
	fmt.Printf("Starting to mine block for transactions:\nTrans1:\n")
	trans1.PrintInfo("\t")
	fmt.Printf("Trans2:\n")
	trans2.PrintInfo("\t")
	//Майнінг блока з транзакціями з приклада
	block := bch.MineBlock(&genBlock, []bch.Transaction{trans1, trans2}, uint64(time.Now().Unix()))
	fmt.Printf("Got block:\n")
	block.PrintInfo("\t")
	if bch.ValidBlock(genBlock, block) {
		fmt.Println("Block has been validated!")
	} else {
		fmt.Println("Block is invalid")
	}
}
