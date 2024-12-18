package blockchain

import (
	"fmt"
	"time"
)

// Block represents each block in the blockchain
type Block struct {
	Timestamp     string
	VoterName     string
	CandidateName string
	Hash          string
	PrevHash      string
}

// Blockchain represents a chain of blocks
type Blockchain struct {
	Chain []Block
}

// NewBlockchain creates a new blockchain with the genesis block
func NewBlockchain() *Blockchain {
	blockchain := &Blockchain{}
	blockchain.AddGenesisBlock()
	return blockchain
}

// AddGenesisBlock adds the first block to the blockchain
func (bc *Blockchain) AddGenesisBlock() {
	genesisBlock := Block{
		Timestamp:     time.Now().String(),
		VoterName:     "System",
		CandidateName: "Genesis",
		Hash:          "genesis_hash",
		PrevHash:      "",
	}
	bc.Chain = append(bc.Chain, genesisBlock)
}

// AddVote adds a new vote block to the blockchain
func (bc *Blockchain) AddVote(voterName, candidateName string) {
	lastBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := Block{
		Timestamp:     time.Now().String(),
		VoterName:     voterName,
		CandidateName: candidateName,
		Hash:          fmt.Sprintf("%d", len(bc.Chain)+1), // Simple hash as an example
		PrevHash:      lastBlock.Hash,
	}
	bc.Chain = append(bc.Chain, newBlock)
}

// GetBlockchain returns the entire blockchain
func (bc *Blockchain) GetBlockchain() []Block {
	return bc.Chain
}
