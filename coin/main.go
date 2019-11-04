package main

import (
	"core"
	"fmt"
	"strconv"
	"time"
)

func main() {
	blkChn := core.NewBlockchain()
	fmt.Println(blkChn)
	blkChn.AddBlock("send 1 btc to Ivan")
	blkChn.AddBlock("send 2 btc to IWa")
	for _, blk := range blkChn.Blocks {
		fmt.Printf("date[%s],preHash[%x],curHash[%x],data[%s]\n", time.Unix(blk.Timestamp, 0).Format("2006-01-02 15:04:05"), blk.PreBlkHsh, blk.Hash, blk.Data)
	}

	fmt.Println(strconv.ParseInt("0000008b0f41ec78bab747864db66bcb9fb89920ee75f43fdaaeb5544f7f76ca", 16, 16))

}
