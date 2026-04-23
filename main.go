package main

import (
	"fmt"

	"assignment01bca/assignment01bca"
)

func main() {
	// Create blockchain with personalized genesis block
	blockchain := assignment01bca.NewBlockchain()

	// Create at least 3 more blocks
	// Includes last 3 digits of roll number (603) in one transaction as required
	blockchain.AddBlock("Fatima pays 603 rupees to Ahmed", 21)
	blockchain.AddBlock("Bob sends documents to Sara", 35)
	blockchain.AddBlock("University fee submitted by student 221603", 42)

	// Print blockchain
	fmt.Println("\nInitial blockchain:")
	blockchain.ListBlocks()

	// Verify blockchain before tampering
	fmt.Println("\nVerification before tampering:")
	blockchain.VerifyChain()

	// Tamper with Block 1
	fmt.Println("\nTampering with Block 1...")
	blockchain.ChangeBlock(1, "Fatima pays 9999 rupees to Hacker")
	fmt.Println("\nBlockchain after tampering:")
	blockchain.ListBlocks()

	// Verify blockchain after tampering
	fmt.Println("\nVerification after tampering:")
	blockchain.VerifyChain()
}
