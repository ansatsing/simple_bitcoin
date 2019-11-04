package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)

const PRE_ZERO_NUM = 4 //针对16进制
const TARGET = 1 << (256 - PRE_ZERO_NUM/4)

//type ProofOfWork struct {
//
//}

func GenerateCounterAndHash(data []byte) (int, []byte, error) {
	var counter int64 = 1
	for {
		buffer := new(bytes.Buffer)
		binary.Write(buffer, binary.BigEndian, counter)
		_data := bytes.Join([][]byte{data, buffer.Bytes()}, []byte{})
		hash := sha256.Sum256(_data)
	}
}
