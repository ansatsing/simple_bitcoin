package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"math/big"
)

const PRE_ZERO_NUM = 4 //针对16进制
//var target   = big.NewInt(1 << (256 - PRE_ZERO_NUM*4))
//var target   = 1 << (256 - PRE_ZERO_NUM*4)

//type ProofOfWork struct {
//
//}.
func GenerateCounterAndHash(data []byte, num int) (int, []byte, error) {
	if data == nil {
		return 0, nil, errors.New("data can not be empty")
	}
	if num < 1 {
		return 0, nil, errors.New("num must greater than zero")
	}
	counter := 1
	target := big.NewInt(1)
	target.Lsh(target, uint(256-num*4))
	fmt.Printf("%64x\n", target)
	var hashInt big.Int
	for {
		buffer := new(bytes.Buffer)
		binary.Write(buffer, binary.BigEndian, counter)
		cnt := buffer.Bytes()[:]
		_data := bytes.Join([][]byte{data, cnt}, []byte{})
		//_data := bytes.Join([][]byte{data, IntToHex(int64(counter))}, []byte{})
		hash := sha256.Sum256(_data)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(target) == -1 {
			return counter, hash[:], nil
		}
		buffer.
			counter++
	}
}
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
