package bdb

const (
	// url path
	apiPath         = "api/v1"
	transactionPath = "transactions"
	assetPath       = "assets"
	blockPath       = "blocks"
	metadataPath    = "metadata"
	outputPath      = "outputs"
	validatorPath   = "validators"

	// operation
	create   = "CREATE"
	transfer = "TRANSFER"

	// spent | unspent
	spent   = "spent"
	unspent = "unspent"

	// condition
	conditionCost      = 131072
	conditionCostStr   = "131072"
	conditionURLPrefix = "ni:///sha-256"
	conditionfpt       = "ed25519-sha-256"

	// post mode
	async  = "async"
	sync   = "sync"
	commit = "commit"
)
