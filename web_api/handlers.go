package main

import (
	"blockchain_m/bch"
	"strconv"

	"time"

	"github.com/gin-gonic/gin"
)

func getLastBlockHandler(context *gin.Context) {
	serialisedBlock, _ := blockchain[len(blockchain)-1].MarshalJSON()
	context.Data(200, "Content-Type: application/json", serialisedBlock)
}
func getBlockByIdHandler(context *gin.Context) {
	blockIdString := context.Query("blockId")
	if blockIdString == "" {
		context.JSON(400, "required parameter 'blockId' was not provided")
		return
	}
	blockId, conversionError := strconv.Atoi(blockIdString)

	if conversionError != nil {
		context.JSON(400, "block id could not be read")
		return
	}

	if blockId < len(blockchain) {
		serialisedBlock, _ := blockchain[blockId].MarshalJSON()
		context.Data(200, "Content-Type: application/json", serialisedBlock)

	} else {
		context.JSON(400, "no block with such index")
	}
}

func addTransactionHandler(context *gin.Context) {
	var inputTransaction bch.Transaction
	parsingError := context.BindJSON(&inputTransaction)
	if parsingError != nil {
		context.JSON(400, "failed to parse transaction")
		return
	}
	mempool = append(mempool, inputTransaction)
	context.JSON(200, inputTransaction)
}

func getMempoolSizeHandler(context *gin.Context) {
	context.JSON(200, len(mempool))
}
func getMempoolHandler(context *gin.Context) {
	context.JSON(200, mempool)
}

func mineBlockHandler(context *gin.Context) {
	if len(mempool) > 0 {
		minedBlock := bch.MineBlock(
			&blockchain[len(blockchain)-1],
			mempool,
			uint64(time.Now().Unix()),
		)
		blockchain = append(blockchain, minedBlock)
		mempool = make([]bch.Transaction, 0)
		serialisedBlock, _ := minedBlock.MarshalJSON()
		context.Data(200, "Content-Type: application/json", serialisedBlock)
	} else {
		context.JSON(400, "there are no transactions to add to a new block")
	}
}

// type ErrorResponce struct {
// 	Error string `json:"error"`
// }
