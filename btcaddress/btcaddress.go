package btcaddress

import(
	"crypto/rand"
	"crypto/ecdsa"
	"github.com/btcsuite/btcd/btcec"

)

func GenerateKeyPair() (*ecdsa.PrivateKey, ) {
	key, _ := ecdsa.GenerateKey(btcec.S256(), rand.Reader)
	return key
}

func GenerateAddress(keyPair ecdsa.PrivateKey) string {
	
}
