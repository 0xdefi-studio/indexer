package models

type BankrollTx struct {
	ID        string
	Sender    string
	Receiver  string
	Owner     string
	Assets    string
	Shares    string
	TokenType string
	TxType    string

	CurrentSenderBalance string
	CurrentBalance       string

	BlockNum        uint64
	TransactionHash string
	Timestamp       uint64
}
