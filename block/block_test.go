package block

import "testing"

func TestCalculateHashIsDeterministic(t *testing.T) {
	b := &Block{Index: 1, Timestamp: 1700000000, Data: "hello", PrevBlockHash: "abc"}
	if b.CalculateHash() != b.CalculateHash() {
		t.Fatal("same block produced different hashes")
	}
}

func TestCalculateHashChangesWithContent(t *testing.T) {
	b := &Block{Index: 1, Timestamp: 1700000000, Data: "hello", PrevBlockHash: "abc"}
	before := b.CalculateHash()
	b.Data = "hellp" // tiny change
	if b.CalculateHash() == before {
		t.Fatal("changing the data did not change the hash")
	}
}

func TestNewBlockLinksToPrevious(t *testing.T) {
	g := NewGenesisBlock()
	b := NewBlock(g, "next")
	if b.PrevBlockHash != g.Hash {
		t.Fatal("block is not linked to its predecessor")
	}
	if b.Hash != b.CalculateHash() {
		t.Fatal("stored hash does not match recomputed hash")
	}
}
