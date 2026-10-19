package chain

import (
	"fmt"

	"github.com/Oliveszn/go-blockchain/block"
)

type Chain struct {
	blocks []*block.Block
}

func NewBlockChain() *Chain {
	genesisBlock := block.NewGenesisBlock()
	return &Chain{
		blocks: []*block.Block{genesisBlock},
	}
}

func (c *Chain) AddBlock(data string) {
	prevBlock := c.blocks[len(c.blocks)-1]
	newBlock := block.NewBlock(prevBlock, data)
	c.blocks = append(c.blocks, newBlock)
}

func (c *Chain) LatestBlock() *block.Block {
	return c.blocks[len(c.blocks)-1]
}

func (c *Chain) IsValidChain() (bool, error) {
	if len(c.blocks) == 0 {
		return false, fmt.Errorf("chain is empty")
	}

	genesis := c.blocks[0]
	if genesis.Index != 0 || genesis.PrevBlockHash != "" {
		return false, fmt.Errorf("first block is not a valid genesis block")
	}

	if genesis.Hash != genesis.CalculateHash() {
		return false, fmt.Errorf("invalid genesis block hash")
	}

	for i := 1; i < len(c.blocks); i++ {
		if err := validateSuccessor(c.blocks[i-1], c.blocks[i]); err != nil {
			return false, fmt.Errorf("block %d: %w", i, err)
		}
	}

	return true, nil
}

// validateSuccessor checks that curr is a valid block following prev.
func validateSuccessor(prev, curr *block.Block) error {
	if curr.Index != prev.Index+1 {
		return fmt.Errorf("index %d does not follow %d", curr.Index, prev.Index)
	}
	if curr.PrevBlockHash != prev.Hash {
		return fmt.Errorf("previous hash does not match the preceding block")
	}
	if curr.Hash != curr.CalculateHash() {
		return fmt.Errorf("stored hash does not match recomputed hash")
	}
	return nil
}
