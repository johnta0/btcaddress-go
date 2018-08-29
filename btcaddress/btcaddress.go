package btcaddress

import(
	"crypto/rand"
	"crypto/ecdsa"
	// "crypto/sha256"
	// "golang.org/x/crypto/ripemd160"

	"github.com/btcsuite/btcd/btcec"
)

func GenerateKeyPair() (*ecdsa.PrivateKey, ) {
	key, _ := ecdsa.GenerateKey(btcec.S256(), rand.Reader)
	return key
}

// func GenerateAddress(keyPair ecdsa.PrivateKey) string {
// 	hash1 := sha256.Sum256([]byte(keyPair.PublicKey)
// }
