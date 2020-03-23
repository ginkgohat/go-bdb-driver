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
	conditionConst  = 131072
	conditionConsts = "131072"
	conditionType   = "ed25519-sha-256"
)
