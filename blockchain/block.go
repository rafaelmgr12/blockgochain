package blockchain

import (
	"bytes"
	"crypto/md5"
)

type Block struct {
	Hash     string
	Data     string
	PrevHash string
	Nonce    int
}

func (b *Block) ComputeHash() {
	concatenatedData := bytes.Join([][]byte{[]byte(b.Data), []byte(b.PrevHash)}, []byte{})

	computedHash := md5.Sum(concatenatedData)

	b.Hash = string(computedHash[:])
}

func CreateBlock(data string, preHash string) *Block {
	block := &Block{Data: data, PrevHash: preHash}
	block.ComputeHash()
	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", "")
}
