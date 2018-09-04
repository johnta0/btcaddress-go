package main

import (
	"fmt"
	"./btcaddress"
)

func main() {
	keyPair := *btcaddress.GenerateKeyPair()
	pubkey := btcaddress.ToScalar(keyPair)

	fmt.Println("Private key: ", keyPair.D)
	fmt.Println("Public key: ", keyPair.PublicKey)
	fmt.Println("Public key (scalar): ", pubkey)

	address := btcaddress.GenerateP2PKHAddress(pubkey)
	fmt.Println("address: ", address)
}