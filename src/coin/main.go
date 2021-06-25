package main

import (
	"fmt"
	"private_blockchain/core"
)

func main() {
	// 初始化区块链，创建第一个区块
	bc := core.NewBlockChain()
	// 加入一个区块，发送一个比特币伊文
	bc.AddBlock("send 1 BTC to Ivan")
	bc.AddBlock("send 2 more BTC to Ivan")
	for _, block := range bc.Blocks {
		fmt.Printf("prev hash %x\n", block.PrevBlockHash)
		fmt.Printf("Data %x\n", block.Data)
		fmt.Printf("Hash %x\n", block.Hash)
		fmt.Println()
	}
}
