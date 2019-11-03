package core

import "time"

type Block struct {
	Timestamp	int64
	Data		[]byte
	PreBlkHsh	[]byte
	Hash		[]byte
}

func (b *Block) SetHash() {

}

func NewBlock(data string, preBlkHsh []byte) *Block {
	block:=&Block{time.Now().Unix(),[]byte(data),preBlkHsh,[]byte{}}
	block.SetHash()
	return block
}