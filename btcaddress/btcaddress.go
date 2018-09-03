package btcaddress

import(
	"fmt"
	"crypto/rand"
	"crypto/ecdsa"
	// "crypto/sha256"
	// "golang.org/x/crypto/ripemd160"

	"github.com/btcsuite/btcd/btcec"
)

// GenerateKeyPair returns (private key public key) key pair
func GenerateKeyPair() (*ecdsa.PrivateKey, ) {
	key, _ := ecdsa.GenerateKey(btcec.S256(), rand.Reader)
	return key
}

// ToScalar returns compressed pubkey generated from EC point public key
func ToScalar(keyPair ecdsa.PrivateKey) string {
	var pubkey_x string = fmt.Sprintf("%x", keyPair.PublicKey.X)
	// var pubkey_y string = fmt.Sprintf("%x", keyPair.PublicKey.Y)
	// if pubkey_y is even => prefix is 02
	// if pubkey_y is odd => prefix is 03
	return "02" + pubkey_x
}

// GenerateP2PKHAddress returns p2psh address from public key
func GenerateP2PKHAddress(pubkey string) string {
	hash1 := sha256.Sum256([]byte(pubkey)
}
