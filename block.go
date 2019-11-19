package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block ...< struct block >
type Block struct {
	Timestamp         int64
	Data              []byte
	PreviousBlockHash []byte
	Hash              []byte
}

// SetHash ... < seter hash  >
func (block *Block) SetHash() {

	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	header := bytes.Join([][]byte{block.PreviousBlockHash, block.Data, timestamp}, []byte{})
	hash := sha256.Sum256(header)
	block.Hash = hash[:] //le pasamos elcontenido del hash, ya que sum256()nos devuelve un tamaño fijo
}

// NewBlock ... < create new block >
func NewBlock(data string, prevBlockHash []byte) *Block {
	newblock := &Block{
		Timestamp:         time.Now().Unix(),
		Data:              []byte(data),
		PreviousBlockHash: prevBlockHash,
		Hash:              []byte{},
	}
	newblock.SetHash()
	return newblock
}

/* for test previous a main.go
func main() {

	myBlock := NewBlock("Test block", []byte("895dsd4"))
	fmt.Println("Timestamp :", myBlock.Timestamp)
	fmt.Println("Data :", string(myBlock.Data))
	fmt.Println("Previuos Block Hash :", string(myBlock.PreviousBlockHash))
	fmt.Println("New Hash :", string(myBlock.Hash))

}*/
