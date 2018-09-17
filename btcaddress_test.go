package btcaddress_test

import (
	"math/big"
	"testing"
	"crypto/ecdsa"
	"../btcaddress-go"
	"github.com/btcsuite/btcd/btcec"
)

const (
    DECIMAL int = 10
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
	pubkeyX, _ := new(big.Int).SetString("15372347899851644783916699817703347896837050808584440115680587892706676598630", DECIMAL)
	pubkeyY, _ := new(big.Int).SetString("15747420107305469165456188672163633096043500936022991061831670415114251650195", DECIMAL)
	pubkey := ecdsa.PublicKey{btcec.S256() , pubkeyX, pubkeyY}

	scalarPubkey := btcaddress.ToScalar(pubkey)
	if (scalarPubkey != "0221fc70c9ce6dfe24cab6c8035adedac2f04531be6113a5b6fbd2ac60e7006766") {
		t.Error(
			"expected: ", "0221fc70c9ce6dfe24cab6c8035adedac2f04531be6113a5b6fbd2ac60e7006766",
			"recieved: ", scalarPubkey,
		)
	}
} 
