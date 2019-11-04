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
}

func (b *Block) SetHash() {
	timeStamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PreBlkHsh, b.Data, timeStamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, preBlkHsh []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), preBlkHsh, []byte{}}
	block.SetHash()
	return block
}
