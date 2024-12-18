package main

import (
	"fmt"
	"log"
	"voting/pkg/blockchain"
)

func main() {
	// Create a new Blockchain instance
	blockchain := blockchain.NewBlockchain()

	// Add a few blocks (votes)
	blockchain.AddVote("Alice", "Candidate 1")
	blockchain.AddVote("Bob", "Candidate 2")

	// Print out the blockchain
	for _, block := range blockchain.GetBlockchain() {
		fmt.Println(block)
	}
}
