package core

import (
	"fmt"
	"testing"
)

func TestGenerateCounterAndHash(t *testing.T) {
	str := "I like donuts"
	counter, hash, err := GenerateCounterAndHash([]byte(str), 1)
	if err != nil {
		fmt.Printf("counter:%d\nhash:%x\n", counter, hash)
	}

}
