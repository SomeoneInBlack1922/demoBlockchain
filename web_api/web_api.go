package main

import (
	// "blockchain_m/bch"
	// "fmt"
	// "time"

	"blockchain_m/bch"

	"github.com/gin-gonic/gin"
)

// Не бачу особливого сенсу створювати структуру для цього
var blockchain []bch.Block
var mempool []bch.Transaction

func main() {
	//Створення колекцій для мемппула та блокчейна
	blockchain = append(blockchain, bch.GetGenesisBlock([32]byte{}, 0))
	mempool = make([]bch.Transaction, 0)
	// blockchain = make([]bch.Block, 16)
	// blockchain
	// mempool := make([]bch.Transaction, 16)

	// gin.SetMode(gin.ReleaseMode)

	// router := gin.New()
	router := gin.Default()
	router.GET("/getLastBlock", getLastBlockHandler)
	router.GET("/getBlockById", getBlockByIdHandler)
	router.POST("/addTransaction", addTransactionHandler)
	router.GET("/getMempoolSize", getMempoolSizeHandler)
	router.GET("/getMempool", getMempoolHandler)
	router.POST("/mineBlock", mineBlockHandler)
	router.Run() // listen and serve on 0.0.0.0:8080

}
