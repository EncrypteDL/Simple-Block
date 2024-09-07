package internal

import (
	"bytes"
	"crypto/sha256"
	"time"
)

// We need to define the Block| A Blocck represent a single block in a chain of blocks(Blockchain)
type Block struct {
	Index       int
	TimesTamp   time.Time
	Data        []byte
	PreviosHash []byte
	Hash        []byte
}

// we need to define a newBlock and returns a new empty `Block`
func newBlock() Block {
	n := Block{}
	n.TimesTamp = time.Now()
	n.Hash = hashBlock(n)

	return n
}

func hashBlock(b Block) []byte {
	h := sha256.New()

	h.Write(Int64Bytes(int64(b.Index)))
	h.Write(Int64Bytes(b.TimesTamp.Unix()))
	h.Write(b.Data)
	h.Write(b.PreviosHash)

	return h.Sum(nil)
}

//isValidate validites the current block ('b') with previos block in a chain
func (b Block) Validate(o Block) bool {
	if (bytes.Compare(b.Hash, hashBlock(b)) != 0) ||
		(b.Index != (o.Index + 1)) ||
		(bytes.Compare(b.PreviosHash, o.Hash) != 0) {
		return false
	}
	return true
}

// Generate creates a new block from the current block which is assumed to be the
// last block in the chain.

func (b Block) Generate(data []byte) Block{
	n := Block{
		Index: b.Index+1,
		TimesTamp: time.Now(),
		Data: make([]byte, data[0]),
		PreviosHash: b.Hash,
	}

	return n
}

