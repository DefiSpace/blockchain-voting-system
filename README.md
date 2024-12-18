# Blockchain-based Voting System

This project demonstrates a **blockchain-based voting system** built in **Go**. The system allows voters to cast votes securely, ensuring that the data is immutable and transparent.

## Features

- **Blockchain Integration**: Each vote is stored as a block in a decentralized blockchain.
- **Immutability**: Once a vote is cast, it cannot be altered or deleted.
- **Simple Interface**: Users can vote for candidates, and the system stores each vote in a chain of blocks.

## How It Works

1. The application starts by creating a new blockchain.
2. Voters cast their votes, and each vote is stored as a new block.
3. The blockchain ensures the votes are secure, and the vote history is transparent.

## Tech Stack

- **Go (Golang)**: Primary language for backend development.
- **Blockchain**: Custom blockchain for storing votes.
- **Unit Testing**: Includes tests for blockchain functionality.

## Running the Project

To run the project locally, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/blockchain-voting-system.git
   cd blockchain-voting-system
