package bdb

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/kalaspuffar/base64url"

	"github.com/mr-tron/base58"
	"github.com/stevenroose/asn1"
)

type (
	PrivateKey ed25519.PrivateKey
	PublicKey  ed25519.PublicKey
)

// KeyPair bigchaindb key pair
type KeyPair struct {
	PrivateKey PrivateKey `json:"privatekey"`
	PublicKey  PublicKey  `json:"publickey"`
}

// NewKeyPair generates a public/private key pair using entropy from rand.
func NewKeyPair() (*KeyPair, error) {
	pub, pri, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		PrivateKey: PrivateKey(pri),
		PublicKey:  PublicKey(pub),
	}, nil
}

// NewKeyPair generates a public/private key pair using private key string
// in bigchaindb,private key string is seed encode string
func NewKeyPairFromPrivateKey(privStr string) (*KeyPair, error) {
	priv, err := NewPrivateKeyFromString(privStr)
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		PrivateKey: PrivateKey(priv),
		PublicKey:  PrivateKey(priv).Public(),
	}, nil
}

// NewPrivateKeyFromString generates a private key pair using private key string
// in bigchaindb,private key string is seed encode string
func NewPrivateKeyFromString(privStr string) (PrivateKey, error) {
	// privStr is seed string
	if privStr == "" {
		return nil, errors.New("PrivateKey not be empty")
	}

	// encode seed
	seed, err := base58.Decode(privStr)
	if err != nil {
		return nil, err
	}

	if len(seed) != ed25519.SeedSize {
		return nil, errors.New("PrivateKey is error format")
	}

	priv := ed25519.NewKeyFromSeed(seed)
	return PrivateKey(priv), nil
}

// NewPublicKeyFromString generates a public key pair using public key string
func NewPublicKeyFromString(pubStr string) (PublicKey, error) {
	if pubStr == "" {
		return nil, errors.New("PublicKey not be empty")
	}

	pub, err := base58.Decode(pubStr)
	if err != nil {
		return nil, err
	}
	return PublicKey(pub), nil
}

// Public returns the PublicKey corresponding to priv.
func (priv PrivateKey) Public() PublicKey {
	publicKey := make([]byte, ed25519.PublicKeySize)
	copy(publicKey, priv[32:])
	return PublicKey(publicKey)
}

// Seed returns the private key seed corresponding to priv. It is provided for
// interoperability with RFC 8032. RFC 8032's private keys correspond to seeds
// in this package.
func (priv PrivateKey) seed() []byte {
	seed := make([]byte, ed25519.SeedSize)
	copy(seed, priv[:32])
	return seed
}

// String return privatekey seed base58 encode
// bigchaindb privatekey string not priv encode, is seed encode
func (priv PrivateKey) String() string {
	return base58.Encode(priv.seed())
}

// Sign signs the given message with priv.
// Ed25519 performs two passes over messages to be signed and therefore cannot
// handle pre-hashed messages. Thus opts.HashFunc() must return zero to
// indicate the message hasn't been hashed. This can be achieved by passing
func (priv PrivateKey) Sign(message []byte) (signature []byte, err error) {
	return ed25519.PrivateKey(priv).Sign(rand.Reader, message, crypto.Hash(0))
}

// String return publicKey seed base58 encode
func (pub PublicKey) String() string {
	return base58.Encode(pub)
}

// contionURL the Crypto Conditions URL
func (pub PublicKey) contionURL() (string, error) {
	var ac *asn1.Context

	content := struct {
		PubKey []byte `asn1:"tag:0"`
	}{
		PubKey: pub,
	}

	data, err := ac.Encode(content)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)

	return fmt.Sprintf("%s;%s?fpt=%s&cost=%s", conditionURLPrefix, base64url.Encode(hash[:]), conditionfpt, conditionCostStr), err
}

// return keypair json string (encode base58)
func (key *KeyPair) Json() string {
	return fmt.Sprintf(`{"privatekey":"%s","publickey":"%s"}`, key.PrivateKey.String(), key.PublicKey.String())
}
