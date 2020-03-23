package bdb

type Node struct {
	API struct {
		Endpoint `json:"v1"`
	} `json:"api"`
	Docs     string `json:"docs"`
	Software string `json:"software"`
	Version  string `json:"version"`
}

type Endpoint struct {
	Assets       string `json:"assets"`
	Blocks       string `json:"blocks"`
	Docs         string `json:"docs"`
	Metadata     string `json:"metadata"`
	Outputs      string `json:"outputs"`
	Streams      string `json:"streams"`
	Transactions string `json:"transactions"`
	Validators   string `json:"validators"`
}

type Transaction struct {
	Asset Asset `json:"asset"`
	// ID has to convert to null value in JSON
	ID        *string  `json:"id"`
	Inputs    []Input  `json:"inputs"`
	Metadata  Metadata `json:"metadata"`
	Operation string   `json:"operation"`
	Outputs   []Output `json:"outputs"`
	Version   string   `json:"version"`
}

type Input struct {
	// Fulfillment can have both uri string or object with pubKey and other info
	// ID has to convert to null value in JSON
	Fulfillment  *string         `json:"fulfillment"`
	Fulfills     *OutputLocation `json:"fulfills"`
	OwnersBefore []string        `json:"owners_before"`
}

type Output struct {
	Amount     string    `json:"amount"`
	Condition  Condition `json:"condition"`
	PublicKeys []string  `json:"public_keys"`
}

type Asset struct {
	Data map[string]interface{} `json:"data,omitempty"`
	ID   *string                `json:"id,omitempty"`
}

type Metadata map[string]interface{}

type AssetResponse struct {
	Data          map[string]interface{} `json:"data"`
	TransactionID string                 `json:"id"`
}

type MetadataResponse struct {
	Metadata      map[string]interface{} `json:"metadata"`
	TransactionID string                 `json:"id"`
}

type OutputLocation struct {
	TransactionID string `json:"transaction_id,omitempty"`
	// Test if this should be json.Number
	OutputIndex int64 `json:"output_index,omitempty"`
}

type Condition struct {
	Details ConditionDetail `json:"details"`
	Uri     string          `json:"uri"`
}

type ConditionDetail struct {
	PublicKey string `json:"public_key"`
	Type      string `json:"type"`
}

type Validator struct {
	PublicKey   map[string]interface{} `json:"public_key"`
	VotingPower int                    `json:"voting_power"`
}

type Block struct {
	Height       int           `json:"height"`
	Transactions []Transaction `json:"transactions"`
}
