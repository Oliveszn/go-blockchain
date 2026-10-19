package chain

import "testing"

func TestNewChainIsValid(t *testing.T) {
	c := NewBlockChain()
	if len(c.blocks) != 1 {
		t.Fatalf("expected 1 block, got %d", len(c.blocks))
	}
	if ok, err := c.IsValidChain(); !ok {
		t.Fatalf("new chain should be valid: %v", err)
	}
}

func TestChainWithBlocksIsValid(t *testing.T) {
	c := NewBlockChain()
	c.AddBlock("a")
	c.AddBlock("b")
	c.AddBlock("c")
	if len(c.blocks) != 4 {
		t.Fatalf("expected 4 blocks, got %d", len(c.blocks))
	}
	if ok, err := c.IsValidChain(); !ok {
		t.Fatalf("chain should be valid: %v", err)
	}
}

func TestTamperedDataIsDetected(t *testing.T) {
	c := NewBlockChain()
	c.AddBlock("a")
	c.AddBlock("b")
	c.blocks[1].Data = "evil"
	if ok, _ := c.IsValidChain(); ok {
		t.Fatal("tampered chain should be invalid")
	}
}

func TestTamperAndRehashIsStillDetected(t *testing.T) {
	c := NewBlockChain()
	c.AddBlock("a")
	c.AddBlock("b")
	c.AddBlock("c")

	c.blocks[2].Data = "evil"
	c.blocks[2].Hash = c.blocks[2].CalculateHash() // make it look self-consistent

	ok, err := c.IsValidChain()
	if ok {
		t.Fatal("rehashed tampering should still be invalid")
	}
	t.Log(err) // see which block gets blamed
}

func TestBrokenLinkIsDetected(t *testing.T) {
	c := NewBlockChain()
	c.AddBlock("a")
	c.AddBlock("b")
	c.blocks[2].PrevBlockHash = "deadbeef"
	c.blocks[2].Hash = c.blocks[2].CalculateHash()
	if ok, _ := c.IsValidChain(); ok {
		t.Fatal("broken link should be invalid")
	}
}
