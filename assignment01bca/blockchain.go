package assignment01bca

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "strconv"
    "time"
)

// block represents a single block in the blockchain.
type block struct {
    Index        int
    Timestamp    string
    Transaction  string
    Nonce        int
    PreviousHash string
    Hash         string
}

// Blockchain stores all blocks.
type Blockchain struct {
    Blocks []*block
}

// CalculateHash calculates SHA-256 over:
// transaction + nonce + previousHash + index + timestamp
func CalculateHash(stringToHash string) string {
    hash := sha256.Sum256([]byte(stringToHash))
    return hex.EncodeToString(hash[:])
}

// CreateHash generates the hash for a block.
func (b *block) CreateHash() string {
    data := b.Transaction +
        strconv.Itoa(b.Nonce) +
        b.PreviousHash +
        strconv.Itoa(b.Index) +
        b.Timestamp
    return CalculateHash(data)
}

// NewBlockchain creates a blockchain with a personalized genesis block.
// For roll number 221603:
// Genesis transaction = "Genesis Block - 221603"
// Genesis nonce = 14 (sum of digits)
func NewBlockchain() *Blockchain {
    bc := &Blockchain{}
    genesis := NewBlock("Genesis Block - 221603", 14, "0")
    genesis.Index = 0
    genesis.Hash = genesis.CreateHash()
    bc.Blocks = append(bc.Blocks, genesis)
    return bc
}

// NewBlock creates a new block with the required signature from the assignment.
func NewBlock(transaction string, nonce int, previousHash string) *block {
    b := &block{
        Transaction:  transaction,
        Nonce:        nonce,
        PreviousHash: previousHash,
        Timestamp:    time.Now().Format(time.RFC3339Nano),
    }
    return b
}

// AddBlock adds a new block to the blockchain.
func (bc *Blockchain) AddBlock(transaction string, nonce int) {
    lastBlock := bc.Blocks[len(bc.Blocks)-1]

    newBlock := NewBlock(transaction, nonce, lastBlock.Hash)
    newBlock.Index = len(bc.Blocks)
    newBlock.Hash = newBlock.CreateHash()

    bc.Blocks = append(bc.Blocks, newBlock)
}

// ListBlocks prints all blocks in a nice readable format.
func (bc *Blockchain) ListBlocks() {
    fmt.Println("================ BLOCKCHAIN ================")
    for _, b := range bc.Blocks {
        fmt.Printf("Block #%d\n", b.Index)
        fmt.Printf("Timestamp     : %s\n", b.Timestamp)
        fmt.Printf("Transaction   : %s\n", b.Transaction)
        fmt.Printf("Nonce         : %d\n", b.Nonce)
        fmt.Printf("Previous Hash : %s\n", b.PreviousHash)
        fmt.Printf("Current Hash  : %s\n", b.Hash)
        fmt.Println("--------------------------------------------")
    }
}

// ChangeBlock changes the transaction of a given block reference.
// Intentionally does not recalculate hash so tampering can be detected.
func (bc *Blockchain) ChangeBlock(index int, newTransaction string) {
    if index < 0 || index >= len(bc.Blocks) {
        fmt.Println("Invalid block index. No changes made.")
        return
    }

    bc.Blocks[index].Transaction = newTransaction
    fmt.Printf("Block #%d transaction has been changed to: %s\n", index, newTransaction)
}

// VerifyChain verifies hashes and previous-hash links.
func (bc *Blockchain) VerifyChain() bool {
    fmt.Println("=============== VERIFY CHAIN ===============")

    for i := 0; i < len(bc.Blocks); i++ {
        current := bc.Blocks[i]
        recalculatedHash := current.CreateHash()

        if current.Hash != recalculatedHash {
            fmt.Printf("Chain verification failed at Block #%d\n", current.Index)
            fmt.Println("Reason: Current block hash does not match recalculated hash.")
            fmt.Printf("Stored Hash      : %s\n", current.Hash)
            fmt.Printf("Recalculated Hash: %s\n", recalculatedHash)
            return false
        }

        if i > 0 {
            previous := bc.Blocks[i-1]
            if current.PreviousHash != previous.Hash {
                fmt.Printf("Chain verification failed at Block #%d\n", current.Index)
                fmt.Println("Reason: Previous hash link is broken.")
                fmt.Printf("Expected Previous Hash: %s\n", previous.Hash)
                fmt.Printf("Actual Previous Hash  : %s\n", current.PreviousHash)
                return false
            }
        }
    }

    fmt.Println("Blockchain is valid. No tampering detected.")
    return true
}
