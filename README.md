# BlockGoChain - Build a Blockchain from Scratch in Go

Welcome to **BlockGoChain**, a hands-on project designed to help developers build a functional blockchain system from scratch using Go. This project explores key blockchain concepts such as blocks, hash functions, proof-of-work algorithms, and transaction management.

## Project Overview

The goal of this project is to implement the fundamental components of a blockchain, starting with block structures and gradually building out features like mining, transaction handling, and wallet functionality. By completing this project, you will gain a deep understanding of how blockchain systems operate, from hashing blocks to managing digital assets.

### Features Implemented

1. **Blocks and Blockchain Structure**
    - Create the fundamental structures of a blockchain, including blocks and chains of blocks.
    - Understand and compute hashes for secure block validation.
    - Create the first "Genesis" block and continue adding blocks to the chain.
2. **Proof of Work Algorithm**
    - Implement a proof-of-work algorithm for mining new blocks.
    - Verify and validate the mining process using computational power.
    - Mine new blocks and test the algorithm.
3. **Transaction Management**
    - Define a transaction structure to handle asset transfers.
    - Integrate transactions within blocks for safe and secure operations.
4. **Wallet Functionality**
    - Build wallet features to manage digital assets.
    - Add signing and verification functionality to ensure secure transactions.
    - Implement wallet functionality and test its reliability.

### Project Tasks

### 1. Blockchain Prototype

- **Task 1**: Create Block and Blockchain Structures
- **Task 2**: Compute Hashes for Blocks
- **Task 3**: Create the Genesis Block
- **Task 4**: Add New Blocks to the Blockchain
- **Task 5**: Test Blockchain Integrity

### 2. Proof of Work

- **Task 6**: Implement Proof of Work Structure
- **Task 7**: Mine New Blocks
- **Task 8**: Validate the Proof of Work Algorithm
- **Task 9**: Test the Mining Process

### 3. Transactions

- **Task 10**: Create Transaction Structures
- **Task 11**: Add Transactions to Blocks
- **Task 12**: Test Transaction Handling

### 4. Wallets

- **Task 13**: Build Wallet Functionality
- **Task 14**: Sign and Verify Transactions
- **Task 15**: Test Wallet Features

### File Structure

- `block.go` - Contains the structure and logic for individual blocks in the blockchain.
- `blockchain.go` - Manages the blockchain structure, adding blocks, and validating them.
- `proof.go` - Implements the proof of work system, including mining functionality.
- `wallet.go` - Handles wallet-related functionality, including signing and verifying transactions.
- `main.go` - The main entry point of the application, bringing all components together.

### Requirements

To run this project, ensure you have Go installed on your machine. You can install Go [here](https://golang.org/dl/).

### Running the Project

1. Clone the repository:
    
    ```bash
    git clone https://github.com/rafaelmgr12/blockgochain
    cd blockgochain
    
    ```
    
2. Install dependencies:
    
    ```bash
    go mod tidy
    ```
    
3. Run the blockchain application:
    
    ```bash
    go run main.go
    ```
