# solana-wallet-go
## Step 1: Set up your development environment
initialize your go.mod file* 
* `go mod init sol-wallet`

## Step 2: Install the Solana SDK for Go
Install the solana-go-sdk package in your project
* `go get -v github.com/portto/solana-go-sdk@v1.14.0`
* `go mod download filippo.io/edwards25519`
* `go mod download github.com/mr-tron/base58`
* `go get github.com/tyler-smith/go-bip39`
* `go get github.com/tyler-smith/go-bip32`