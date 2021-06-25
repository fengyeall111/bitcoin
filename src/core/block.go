package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// block header
type Block struct {
	Timestamp     int64  // 区块创建的时间
	Data          []byte // 区块包含的数据
	PrevBlockHash []byte // 上一个区块的哈希值,组成链的必要参数
	Hash          []byte // 区块自身的哈希值，用与校验区块额数据有效性
}

// 初始化创建一个新区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
	}
	block.SetHash()
	return block
}

// 设置区块的hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// 创建第一个块，创世纪块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
