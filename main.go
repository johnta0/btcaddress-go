package main

import (
	"fmt"
	"./btcaddress"
)

func main() {
	keyPair := *btcaddress.GenerateKeyPair()
	
	fmt.Println("Private key: ", keyPair.D)
	fmt.Println("Public key: ", keyPair.PublicKey)
	fmt.Println("Public key (scalar)", )
}