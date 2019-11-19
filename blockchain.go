package main

// Blockchain ... < Blockchain struct  >
type Blockchain struct {
	blocks []*Block
}

// AddBlock ... <  add block -simple example >
func (blockChain *Blockchain) AddBlock(data string) {
	previousBlock := blockChain.blocks[len(blockChain.blocks)-1]
	newBlock := NewBlock(data, previousBlock.Hash)
	blockChain.blocks = append(blockChain.blocks, newBlock)
}

// NewOriginBlock ... <  origin block - simple example >
func NewOriginBlock() *Block {
	return NewBlock("Origin Block", []byte{})
}

//NewBlockChain ... <create new blockChain from origin block>
func NewBlockChain() *Blockchain {
	return &Blockchain{blocks: []*Block{NewOriginBlock()}}
}
