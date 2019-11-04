package core

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	preBlk := bc.LastBlock()
	newBlk := NewBlock(data, preBlk.Hash)
	bc.Blocks = append(bc.Blocks, newBlk)
}

func (bc *Blockchain) LastBlock() *Block {
	len := len(bc.Blocks)
	block := bc.Blocks[len-1]
	return block
}

//creates a Blockchain with a genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
