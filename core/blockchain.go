package core

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	preBlk := bc.LastBlock()
	newBlk := NewBlock(data, preBlk.Hash)
	bc.blocks = append(bc.blocks, newBlk)
}

func (bc *Blockchain) LastBlock() *Block {
	len := len(bc.blocks)
	block := bc.blocks[len-1]
	return block
}
