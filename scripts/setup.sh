#!/bin/bash

# Setup Go environment (if Go is not installed)
if ! [ -x "$(command -v go)" ]; then
  echo "Error: Go is not installed. Please install Go first." >&2
  exit 1
fi

# Install dependencies (if needed)
echo "Installing dependencies..."
go mod tidy

# Run the application
echo "Starting the Blockchain Voting System..."
go run cmd/main.go
