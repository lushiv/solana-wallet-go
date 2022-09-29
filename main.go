package main

import (
	"fmt"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"

	"github.com/portto/solana-go-sdk/types"
)

func generate_mnemonic() (result string) {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic_key, _ := bip39.NewMnemonic(entropy)
	return mnemonic_key

}

func creating_a_new_solana_wallet() {
	wallet := types.NewAccount()
	fmt.Println("Wallet Address:", wallet.PublicKey.ToBase58())
	fmt.Println("Private Key:", wallet.PrivateKey)

}
func creating_a_new_solana_hd_wallet(mnemonic string) {
	//solana_derivation_path := "m/44'/501'/0'/0'"
	seed := bip39.NewSeed(mnemonic, "secret_key")
	private_key, _ := bip32.NewMasterKey(seed)
	public_key := private_key.PublicKey()
	fmt.Println("The private key is : ", private_key)
	fmt.Println("The public key is : ", public_key)

}

func main() {
	var mnemonic string = ""
	mnemonic = generate_mnemonic()
	println("mnemonic Address is : ", mnemonic)
	creating_a_new_solana_wallet()
	creating_a_new_solana_hd_wallet(mnemonic)
}
