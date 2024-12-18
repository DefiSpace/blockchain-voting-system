package blockchain

import (
	"fmt"
)

// Vote represents a vote cast by a voter for a specific candidate.
type Vote struct {
	VoterName     string
	CandidateName string
	Timestamp     string
}

// NewVote creates a new vote instance.
func NewVote(voterName, candidateName, timestamp string) *Vote {
	return &Vote{
		VoterName:     voterName,
		CandidateName: candidateName,
		Timestamp:     timestamp,
	}
}

// CastVote adds the vote to the blockchain.
func (bc *Blockchain) CastVote(voterName, candidateName string) {
	// Create a new vote and add it to the blockchain.
	vote := NewVote(voterName, candidateName, fmt.Sprintf("%v", time.Now()))
	bc.AddVote(vote)
}

// AddVote adds a vote to the blockchain as a new block.
func (bc *Blockchain) AddVote(vote *Vote) {
	lastBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := Block{
		Timestamp:     vote.Timestamp,
		VoterName:     vote.VoterName,
		CandidateName: vote.CandidateName,
		Hash:          fmt.Sprintf("%d", len(bc.Chain)+1), // Simple hash as an example
		PrevHash:      lastBlock.Hash,
	}
	bc.Chain = append(bc.Chain, newBlock)
}
