# Blockchain Assignment 1

A simple blockchain implementation in Go to demonstrate foundational blockchain concepts such as block structure, hash linking, immutability, tamper detection, and chain verification logic (without Proof-of-Work).

## Student Information
- **Roll Number**: 221603

## Features Implemented

- **Custom Block Structure**: Each block contains an `Index`, `Timestamp`, `Transaction`, `Nonce`, `PreviousHash`, and `Hash`.
- **SHA-256 Hashing**: Block hashes are calculated securely using SHA-256 by concatenating the transaction, nonce, previous hash, index, and timestamp.
- **Personalized Genesis Block**: The genesis block is tailored with the transaction `"Genesis Block - 221603"` and a nonce of `14` (the sum of the digits of the roll number).
- **Roll Number Injection**: Following the assignment requirements, the last 3 digits of the roll number (`603`) are included in one of the block transactions.
- **Chain Verification**: The `VerifyChain()` method checks the integrity of the blockchain by recalculating hashes and ensuring the `PreviousHash` links are completely intact.
- **Tamper Detection**: The `main.go` file includes a demonstration where a block is maliciously altered using `ChangeBlock()`. The `VerifyChain()` method is then run to successfully detect the failure and broken links.

## Project Structure

```text
.
├── assignment01bca/
│   └── blockchain.go   # Core blockchain logic and structures
├── main.go             # Entry point, chain initialization, and tamper demonstration
├── go.mod              # Go module definition
└── README.md           # Project documentation
```

## How to Run

1. Make sure you have [Go](https://go.dev/dl/) installed on your system.
2. Open a terminal in the project directory (`assignment01bca`).
3. Execute the entry-point script:

```bash
go run main.go
```

### Expected Output
When you run the program, it will:
1. Print the initial, properly verified blockchain containing the Genesis block and 3 transaction blocks.
2. Run an initial chain verification (which will pass).
3. Intentionally alter the transaction in Block #1 to simulate a hack/tamper.
4. Run the chain verification again, which will rightfully detect the mismatch and print the verification failure, proving the immutability of the chain.