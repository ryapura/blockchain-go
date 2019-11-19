package main

import "fmt"

func main() {
	blockChain := NewBlockChain()
	blockChain.AddBlock("BUY 50 BTC - 452000.00 USD from @ryapura to @rauly")
	blockChain.AddBlock("SELL 18000.00 USD - 2 BTC from @raul.yapura to @ryapura")
	blockChain.AddBlock("BUY 30 BTC - 285000.00 USD from @trader1 to @trader2")
	blockChain.AddBlock("SELL 19000.00 USD - 2.5 BTC from @trader2 to @trader3")
	//recoremos la blackchain
	for _, block := range blockChain.blocks {
		fmt.Printf("PREV HASH : %x\n", block.PreviousBlockHash)
		fmt.Printf("DATA BLOCK: %s\n", block.Data)
		fmt.Printf("NEW  HASH : %x\n", block.Hash)
		fmt.Println("-------------------------------------------------------------------------\n\n")

	}
}
