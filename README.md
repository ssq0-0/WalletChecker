# Wallet checker

A script to check the number of tokens on the wallet from different projects. 

---

## Features

- **Linea LXP:** Module for checking the number of LXP on a wallet in the Linea network

*(More features coming soon!)*

---

## Installation

### Requirements

- **Go** (Version 1.22 or newer)
- Git (for cloning the repository)
- Optional: `make` for simplified build and run commands.

### Steps

1. Clone the repository:
```bash
git clone https://github.com/ssq0-0/WalletChecker.git
cd base
go mod download
go build -o walletchecker ./core/main.go   
```
2. Run the application:

```bash
./walletchecker
```

3. Or use make(if installed):
```bash
make run
```
4. Or use docker: 
```bash
docker build -t walletchecker:latest .

docker run -it \                                                                                    
  -v "$(pwd)/account:/app/account" \
  walletchecker
```
---

### Wallets (`wallets.csv`)

This section defines the wallets used by the software. Each wallet is described by the following fields:

- **`address`**: The addres of your wallet.
---

### Modules (`modules`)

When you run the program, a list of modules will be displayed in the console. The user's task is to select a module, otherwise the program will be terminated

### For additional assistance or troubleshooting, refer to the official documentation or reach out via [support channel](https://t.me/cheifssq).
