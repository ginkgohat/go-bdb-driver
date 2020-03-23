package bdb

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"strings"

	"github.com/go-interledger/cryptoconditions"
	"github.com/mr-tron/base58/base58"
	"github.com/pkg/errors"
)

type KeyPair struct {
	PrivateKey ed25519.PrivateKey `json:"privateKey"`
	PublicKey  ed25519.PublicKey  `json:"publicKey"`
}

// NewKeyPair
func NewKeyPair() (*KeyPair, error) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, errors.Wrap(err, "Could not generate new ED25519 KeyPair")
	}

	return &KeyPair{
		PublicKey:  pubKey,
		PrivateKey: privKey,
	}, nil
}

// Seed
func (k *KeyPair) Seed() []byte {
	return k.PrivateKey.Seed()
}

// GenPublicKey
func GenPublicKey(priv ed25519.PrivateKey) ed25519.PublicKey {
	publicKey := make([]byte, ed25519.PublicKeySize)
	copy(publicKey, priv[32:])
	return publicKey
}

// GenKeyPairFromSeed
func GenKeyPairFromSeed(seed []byte) *KeyPair {
	privKey := ed25519.NewKeyFromSeed(seed)
	keys := &KeyPair{
		PrivateKey: privKey,
		PublicKey:  GenPublicKey(privKey),
	}
	return keys
}

func NewEd25519Condition(pubKey ed25519.PublicKey) *cryptoconditions.Condition {
	return cryptoconditions.NewSimpleCondition(cryptoconditions.CTEd25519Sha256, pubKey, conditionConst)
}

/*
	The privateKeys slice expects keys in the same order as the accompanying public key
	in the transaction.Inputs

*/
func (t *Transaction) Sign(keyPairs []*KeyPair) error {

	// Set transaction ID to ctnull value
	t.ID = nil

	signedTx := *t

	// Compute signatures of inputs
	for idx, input := range signedTx.Inputs {
		var serializedTxn strings.Builder
		s, err := t.String()
		if err != nil {
			return err
		}
		serializedTxn.WriteString(s)

		keyPair := keyPairs[idx]

		// If fulfills is not empty add to make unique serialization Txn
		if input.Fulfills != nil {
			serializedTxn.WriteString(input.Fulfills.TransactionID)
			serializedTxn.WriteString(string(input.Fulfills.OutputIndex))
		}

		bytesToSign := []byte(serializedTxn.String())

		// rand reader is ignored within Sign method; crypto.Hash(0) is sanity check to
		// make sure bytes_to_sign is not hashed already - ed25519.PrivateKey cannot sign hashed msg
		signature, err := keyPair.PrivateKey.Sign(rand.Reader, bytesToSign[:], crypto.Hash(0))

		// https://tools.ietf.org/html/draft-thomas-crypto-conditions-03#section-8.5
		ed25519Fulfillment, err := cryptoconditions.NewEd25519Sha256(keyPair.PublicKey, signature)
		if err != nil {
			return errors.Wrap(err, "Could not create fulfillment")
		}

		// TODO - Not sure whether this should be ed25519Fulfillment.Encode()
		// TODO - or ed25519Fulfillment.Condition().Encode()
		ff, err := ed25519Fulfillment.Encode()
		if err != nil {
			return err
		}
		ffSt := base58.Encode(ff)
		signedTx.Inputs[idx].Fulfillment = &ffSt
	}
	//Create ID of transaction (hash of body)
	id, err := signedTx.createID()
	if err != nil {
		return errors.Wrap(err, "Could not create ID")
	}
	t.Inputs = signedTx.Inputs
	t.ID = &id

	return nil
}
