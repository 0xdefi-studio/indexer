package models

type DicePlayTx struct {
	ID                      string
	RequestId               uint64
	PlayerAddress           string
	IsSuccess               int8
	CallbackTransactionHash string

	BlockNum        uint64
	TransactionHash string
	Timestamp       uint64
}

type DiceOutcome struct {
	ID             string
	PlayerAddress  string
	Wager          string
	Payout         string
	TokenAddress   string
	CurrentBalance string
	Contract       string
	DiceOutcomes   string

	BlockNum        uint64
	TransactionHash string
	Timestamp       uint64
}
