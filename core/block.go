package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64
	Data      []byte
	PreBlkHsh []byte
	Hash      []byte
	Counter   int64
}

func (b *Block) SetHash() {
	timeStamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PreBlkHsh, b.Data, timeStamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, preBlkHsh []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), preBlkHsh, []byte{}, 0}
	block.SetHash()
	return block
}

//creates and returns genesis block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
