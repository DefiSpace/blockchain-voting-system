package blockchain

// Voter represents a person casting a vote
type Voter struct {
	Name string
}

// Vote represents a single vote for a candidate
type Vote struct {
	CandidateName string
}
