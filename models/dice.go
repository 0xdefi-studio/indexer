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
