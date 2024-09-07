package internal

import "errors"

var (
	// ErrInvalidBlock is an error returned if an invalid block is added to
	// a block chain where the new block's has and index are invalid
	ErrInvalidBlock = errors.New("error: invalid block")
)

// Chain is a slice of blocks to form the "block chain". Each block is
// connected to the previous block cryptographically by each block's hash
// being in corporated into the next block in the chain. The larger the chain
// grows the higher the integrity of the chain and the more difficult it is to
// temper with or modify previous blocks in the chain.
type Chain []Block

// NewChain creates a new "block chain" `Chain` with an initial block already
// created called the "genesis block"
func NewChain() *Chain {
	return &Chain{newBlock()}
}

// Add adds a new block (`block`) to the chain verifying its validty
// If the block is invalid an error is returned otherwise the block is appended
// to the block chain.
func (c *Chain) Add(block Block) error {
	prevBlock := (*c)[len(*c)-1]
	if !block.Validate(prevBlock) {
		return ErrInvalidBlock
	}
	*c = append(*c, block)
	return nil
}

// Write creates a new block with the given data (`data`) and appends it to the
// block chain. This implements the `io.Writer` interface so you can treat the
// block chain as a valid Writer.
func (c *Chain) Write(data []byte) (int, error) {
	prevBlock := (*c)[len(*c)-1]
	block := prevBlock.Generate(data)
	if err := c.Add(block); err != nil {
		return 0, ErrInvalidBlock
	}
	return len(data), nil
}

// Verify verifies the cryptographic hashes of every block inthe chain
// ensuring all blocks are valid and their integrity in tact
func (c *Chain) Verify() bool {
	prevBlock := (*c)[0]
	for _, block := range (*c)[1:] {
		if !block.Validate(prevBlock) {
			return false
		}
		prevBlock = block
	}
	return true
}
