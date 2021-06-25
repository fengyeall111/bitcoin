package core

// 串联块形成链
type BlockChain struct {
	Blocks []*Block
}

// 添加区块到链中
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// 创建一个新的创世纪区块
func NewBlockChain() *BlockChain {
	return &BlockChain{
		Blocks: []*Block{NewGenesisBlock()},
	}
}
