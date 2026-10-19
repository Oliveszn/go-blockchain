package block

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

type Block struct {
	Index         int    `json:"index"`
	Timestamp     int64  `json:"timestamp"`
	Data          string `json:"data"`
	PrevBlockHash string `json:"prev_block_hash"`
	Hash          string `json:"hash"`
	Nonce         int64  `json:"nonce"`
}

func (b *Block) CalculateHash() string {
	c := *b     // copy, so we never mutate the original
	c.Hash = "" // the hash can't include itself

	data, _ := json.Marshal(c)
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func NewBlock(prev *Block, data string) *Block {
	b := &Block{
		Index:         prev.Index + 1,
		Timestamp:     time.Now().Unix(),
		Data:          data,
		PrevBlockHash: prev.Hash,
	}
	b.Hash = b.CalculateHash()
	return b
}

func GenesisBlock() *Block {
	b := &Block{
		Index:     0,
		Timestamp: time.Now().Unix(),
		Data:      "Genesis Block",
	}
	b.Hash = b.CalculateHash()
	return b
}
