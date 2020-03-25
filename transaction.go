package bdb

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"strings"

	cryptoconditions "github.com/Mashatan/go-cryptoconditions"
	"github.com/mr-tron/base58/base58"
	"github.com/pkg/errors"

	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/sha3"
)

func NewTransaction(
	operation string,
	asset Asset,
	metadata Metadata,
	inputs []Input,
	outputs []Output,
) (*Transaction, error) {

	if !(operation == create || operation == transfer) {
		return &Transaction{}, errors.New("Not a valid operation - expecting 'CREATE' or 'TRANSFER'")
	}

	return &Transaction{
		Asset:     asset,
		ID:        nil,
		Inputs:    inputs,
		Metadata:  metadata,
		Operation: operation,
		Outputs:   outputs,
		Version:   "2.0",
	}, nil
}

func NewCreateTransaction(
	asset Asset,
	metadata Metadata,
	outputs []Output,
	issuers []ed25519.PublicKey,
) (*Transaction, error) {
	// Create list of unfulfilled unfulfilledInputs
	unfulfilledInputs := createInputs([]ed25519.PublicKey{issuers[0]})
	return NewTransaction("CREATE", asset, metadata, unfulfilledInputs, outputs)
}

func createInputs(publicKeys []ed25519.PublicKey) []Input {
	var inputs []Input
	for _, pubKey := range publicKeys {
		input := Input{
			// New input is unfulfilled - fulfillment string is ctnull
			Fulfillment:  nil,
			Fulfills:     nil,
			OwnersBefore: []string{base58.Encode(pubKey)},
		}
		inputs = append(inputs, input)
	}
	return inputs
}

func createInputsFromUnspentTransactions(unspentTransactions []Transaction) ([]Input, error) {
	var inputs []Input

	var unspentOutputs []Output
	for _, ut := range unspentTransactions {
		unspentOutputs = append(unspentOutputs, ut.Outputs...)
	}

	for _, uo := range unspentOutputs {
		input := Input{
			OwnersBefore: uo.PublicKeys,
		}
		inputs = append(inputs, input)

	}
	return inputs, nil
}

// TODO clarify starting point transfer txn: Outputlocations or unspent transactions?
func NewTransferTransaction(
	unspentTransactions []Transaction,
	outputs []Output,
	metadata Metadata,
) (*Transaction, error) {
	inputs, err := createInputsFromUnspentTransactions(unspentTransactions)
	if err != nil {
		return nil, err
	}

	var asset Asset
	// FIXME make sure all unspent txns point to same asset
	if unspentTransactions[0].Operation == "CREATE" {
		asset.ID = unspentTransactions[0].ID
	} else {
		asset = unspentTransactions[0].Asset
	}

	return NewTransaction("TRANSFER", asset, metadata, inputs, outputs)
}

func NewOutput(condition cryptoconditions.Conditions, amount string) (Output, error) {
	if amount == "" {
		amount = "1"
	}

	if condition.Type() == cryptoconditions.CTThresholdSha256 {
		return Output{}, errors.New("No support for treshold-sha-256 yet")
	}

	return Output{
		Condition: Condition{
			Uri: condition.URI(),
			Details: ConditionDetail{
				PublicKey: base58.Encode(condition.Fingerprint()),
				Type:      strings.ToLower(condition.Type().String()),
			},
		},
		Amount:     amount,
		PublicKeys: []string{base58.Encode(condition.Fingerprint())},
	}, nil
}

/*
	The ID of a transaction is the SHA3-256 hash of the transaction.
*/
func (t *Transaction) createID() (string, error) {

	// Strip ID of txn
	tn := &Transaction{
		ID:        nil,
		Version:   t.Version,
		Inputs:    t.Inputs,
		Outputs:   t.Outputs,
		Operation: t.Operation,
		Asset:     t.Asset,
		Metadata:  t.Metadata,
	}
	// Serialize transaction - encoding/json follows RFC7159 and BDB marshalling
	dbytes := tn.JSON()

	// Return hash of serialized txn object
	h := sha3.Sum256(dbytes)
	return hex.EncodeToString(h[:]), nil
}

func (t *Transaction) String() string {
	return string(t.JSON())
}

func (t *Transaction) JSON() []byte {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(t)

	return bf.Bytes()
}
