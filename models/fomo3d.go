package models

type Fomo3dOnEndTx struct {
	ID            string
	PlayerAddress string
	EthIn         string
	KeysBought    string
	WinnerAddr    string
	AmountWon     string
	NewPot        string
	P3DAmount     string
	GenAmount     string
	PotAmount     string
	AirDropPot    string

	BlockNum        uint64
	TransactionHash string
	Timestamp       uint64
}
