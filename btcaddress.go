package btcaddress

import(
	"fmt"
	"encoding/hex"
	"crypto/rand"
	"crypto/ecdsa"
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"

	"github.com/m0t0k1ch1/base58"
	"github.com/btcsuite/btcd/btcec"
)

// GenerateKeyPair returns (private key public key) key pair
func GenerateKeyPair() (*ecdsa.PrivateKey, error ) {
	key, err := ecdsa.GenerateKey(btcec.S256(), rand.Reader)
	if err != nil { return nil, err } 
	return key, nil
}

// ToScalar returns compressed pubkey generated from EC point public key
func ToScalar(pubkey ecdsa.PublicKey) string {
	pubkeyX := fmt.Sprintf("%x", pubkey.X)
	// pubkeyY := fmt.Sprintf("%x", keyPair.PublicKey.Y)
	// if pubkeyY is even => prefix is 02
	// if pubkeyY is odd => prefix is 03
	return "02" + pubkeyX
}

// GeneratePkh returns pubkey-hash
func GeneratePkh(pubkey string) string {
	hash1 := sha256.Sum256([]byte(pubkey))
	hash1String := hex.EncodeToString(hash1[:])

	rip := ripemd160.New()
	hash2 := rip.Sum([]byte(hash1String))
	return hex.EncodeToString(hash2[:])
}

// Base58Encode returns base58-encoded string
func Base58Encode(pkh string) string {
	pkhBytes, _ := hex.DecodeString(pkh)
	b58 := base58.NewBitcoinBase58()
	addr, _ := b58.EncodeToString(pkhBytes)
	return addr
}
