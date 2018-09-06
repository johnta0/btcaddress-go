package btcaddress_test

import (
	"testing"
	"../btcaddress-go"
)

func TestGenerateKeyPair(t *testing.T) {
	// Generate a key pair
	kp, err := btcaddress.GenerateKeyPair()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("PrivateKey: ", kp.D)
	t.Log("Public key: ", kp.PublicKey)
}

func TestToScalar(t *testing.T) {
	// kp, _ := btcaddress.GenerateKeyPair()
} 
