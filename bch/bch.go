package bch

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
)

type Transaction struct {
	Sender  string  `json:"sender"`
	Reciver string  `json:"reciver"`
	Amount  float64 `json:"amount"`
}

func (trans *Transaction) GetBytes() []byte {
	// fmt.Println(trans.Sender, len(trans.Sender))
	// transLen := len(trans.Sender) + len(trans.Reciver) + 8
	var buff []byte
	buff = append([]byte(trans.Sender), []byte(trans.Reciver)...)
	buff = append(buff, binary.AppendUvarint(buff, math.Float64bits(trans.Amount))...)
	return buff
}
func (trans *Transaction) GetTXID() [sha256.Size]byte {
	return sha256.Sum256(trans.GetBytes())
}

type Block struct {
	Index        uint64            `json:"index"`
	Timestamp    uint64            `json:"timestamp"`
	Transactions []Transaction     `json:"transactions"`
	Proof        uint64            `json:"proof"`
	Hash         [sha256.Size]byte `json:"hash"`
	PreviousHash [sha256.Size]byte `json:"previous_hash"`
}

func (blk *Block) PrintInfo(prefix string) {
	fmt.Printf("%v{\n%vIndex: %d\n%vTimestamp: %d\n%vTrasnactions:\n", prefix, prefix, blk.Index, prefix, blk.Timestamp, prefix)
	for _, trans := range blk.Transactions {
		trans.PrintInfo(prefix + "\t")
	}
	fmt.Printf("%vProof: %d\n%vHash: %x\n%vPreviuosHash: %x\n", prefix, blk.Proof, prefix, blk.Hash, prefix, blk.PreviousHash)
	fmt.Printf("%v}\n", prefix)
}
func (trans *Transaction) PrintInfo(prefix string) {
	fmt.Printf("%v{\n", prefix)
	fmt.Printf("%vSender: %v\n%vReciver: %v\n%vAmount: %v\n", prefix, trans.Sender, prefix, trans.Reciver, prefix, trans.Amount)
	fmt.Printf("%v}\n", prefix)
}
func (blk *Block) GetBytes() []byte {
	var buff []byte
	buff = binary.AppendUvarint(buff, blk.Index)
	buff = binary.AppendUvarint(buff, blk.Timestamp)
	for _, trans := range blk.Transactions {
		buff = append(buff, trans.GetBytes()...)
	}
	buff = binary.AppendUvarint(buff, blk.Proof)
	buff = append(buff, blk.PreviousHash[:32]...)
	return buff
}

// Calculates hash of block based on all fileds except for Hash field
func (blk *Block) GetHash() [sha256.Size]byte {
	return sha256.Sum256(blk.GetBytes())
}

// Returns true if hash has two 0x00 bytes at start
func ValidHash(hash [sha256.Size]byte) bool {
	return bytes.Equal(hash[:2], []byte{0x00, 0x00})
}

// Returns genesis block with valid mined hash
func GetGenesisBlock(previousHash [sha256.Size]byte, timestamp uint64) Block {
	block := Block{
		Index:        0,
		Timestamp:    timestamp,
		Proof:        0,
		PreviousHash: previousHash,
	}
	var hashCandidate [sha256.Size]byte = block.GetHash()
	for !ValidHash(hashCandidate) {
		block.Proof += 1
		hashCandidate = block.GetHash()
	}
	block.Hash = hashCandidate
	return block
}

// Creates a block with given transactions and timestamp and mines it to get correct hash. Returns mined correct block
func MineBlock(previousBlock *Block, transactions []Transaction, timestamp uint64) Block {
	block := Block{
		Index:        previousBlock.Index + 1,
		Timestamp:    timestamp,
		Transactions: transactions,
		Proof:        0,
		PreviousHash: previousBlock.Hash,
	}
	var hashCandidate [sha256.Size]byte = block.GetHash()
	for !ValidHash(hashCandidate) {
		block.Proof += 1
		hashCandidate = block.GetHash()
	}
	block.Hash = hashCandidate
	return block
}

// Generates hash of block and returnes true is it equals Hash fild of this block
func ValidBlock(previousBlock Block, currentBlock Block) bool {
	return currentBlock.Index == previousBlock.Index+1 &&
		currentBlock.PreviousHash == previousBlock.Hash &&
		ValidHash(previousBlock.Hash) &&
		ValidHash(currentBlock.Hash) &&
		currentBlock.Hash == currentBlock.GetHash()
}

// Імплементація серіалізації блока у JSON щоб отримати хеш у вигляді шіснадцяткового числа
func (givenBlock *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		&struct {
			Index        uint64        `json:"index"`
			Timestamp    uint64        `json:"timestamp"`
			Transactions []Transaction `json:"transactions"`
			Proof        uint64        `json:"proof"`
			Hash         string        `json:"hash"`
			PreviousHash string        `json:"previous_hash"`
		}{
			Index:        givenBlock.Index,
			Timestamp:    givenBlock.Timestamp,
			Transactions: givenBlock.Transactions,
			Proof:        givenBlock.Proof,
			Hash:         fmt.Sprintf("%X", givenBlock.Hash),
			PreviousHash: fmt.Sprintf("%X", givenBlock.PreviousHash),
		},
	)
}
