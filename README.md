# Blockchain Assignment 1

A simple blockchain implementation in Go to demonstrate foundational blockchain concepts such as block structure, hash linking, immutability, tamper detection, and chain verification logic (without Proof-of-Work).

## Student Information
- **Roll Number**: 221603

## Detailed Breakdown

### `assignment01bca/blockchain.go` (The Blockchain Core)
This file defines the heart of the custom blockchain and contains the backend logic.
*   **Struct Definitions**: Creates a blueprint for an individual `block` (containing `Index`, `Timestamp`, `Transaction`, `Nonce`, `PreviousHash`, and `Hash`) and the `Blockchain` representing the full chain array.
*   **`CalculateHash` / `CreateHash`**: Core security utility that concatenates every single piece of information about a block (Transaction + Nonce + PreviousHash + Index + Timestamp), digests it using Go's `crypto/sha256` library, and outputs a constant-length hexadecimal encoded string.
*   **`NewBlockchain` and `NewBlock`**: Blueprint initializers. `NewBlockchain` automatically sparks a brand new chain starting with the unlinked "Genesis Block". (My Genesis Block explicitly handles `Roll Num 221603` logic).
*   **`AddBlock`**: Automatically calculates a shiny new hash and correctly links the new block's `PreviousHash` parameter to the previously added block's calculated `Hash`.
*   **`ChangeBlock` & `VerifyChain`**: Core security tests. `ChangeBlock` is an intentional weakness designed for this assignment to overwrite transaction strings _without_ recalculating the respective block hash. `VerifyChain` loops through every single block sequentially, recalculating its current elements, and loudly returning a failure if elements/hashes do not perfectly match.

### `main.go` (The Execution Flow)
The master start-up script where the required logic tests from the syllabus run sequentially.
1.  Bootstraps the personalized blockchain (`assignment01bca.NewBlockchain()`).
2.  Mints 3 specific blocks to track transfers (One uniquely injecting the last 3 digits of the roll number `603` natively).
3.  Lists the unbroken chain visually onto standard output and cleanly `VerifyChain()`s that the hashes map 1:1.
4.  Attacks the 1st block dynamically updating its payload to a fake malicious value.
5.  Repeats the validation procedure to successfully error out proving the cryptographic links function as intended.

## Frequently Asked Demo Questions

**Q1: If you change a transaction in a block, why does its hash change? What exactly goes into calculating the hash?**
**Answer**: Our blockchain implements cryptographic hashing logic, primarily SHA-256. A hash function takes an input (no matter the length) and creates a fixed-size deterministic string footprint. The exact input required for my blockchain is mapping `Transaction string` + `integer Nonce` + `Previous Hash` + `block Index` + `Timestamp` formatted strictly into one massive string. Changing just one letter in the `Transaction` drastically shifts the SHA-256 output. 

**Q2: If `Block 1` is tampered with, how does `VerifyChain()` know? Would it also affect `Block 2`?**
**Answer**: `VerifyChain()` works on an aggressive loop check over the whole struct array. When evaluating `Block 1`, it immediately recalculates what `Block 1`'s hash _should_ be based on its current elements via `CreateHash()`. It checks this new hash against `Block 1`'s recorded `Hash` parameter and fails loudly when they don't match.
Additionally, because `Block 2`'s `PreviousHash` depends entirely on `Block 1`'s true hash, a hacker wouldn't just need to secretly update `Block 1`'s hash to match its new payload, they would also subsequently break the chain link into `Block 2`, meaning they'd have to regenerate `Block 2`'s hash... and `Block 3`'s hash... entirely resetting the chain.

**Q3: How did you calculate the nonce for your Genesis block?**
**Answer**: The Genesis block nonce parameter natively leverages the student's Roll Number. Given `221603`, iterating heavily yields the mathematical additive: `2 + 2 + 1 + 6 + 0 + 3 = 14`. Thus, the Genesis integer mapping uses `14`.

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