package chain

import (
	"fmt"
	"github.com/0xdefi-studio/indexer/dice"
	"github.com/0xdefi-studio/indexer/fomo3d"
	models2 "github.com/0xdefi-studio/indexer/models"
	"strings"
)

func (m *Manager) InsertEndTxV1(endTx *fomo3d.Fomo3dOnEndTx, blockNum uint64, txHash string, timestamp uint64) (*models2.Fomo3dOnEndTx, error) {
	userAddress := strings.ToLower(endTx.PlayerAddress.Hex())
	l := &models2.Fomo3dOnEndTx{
		//CompressedData:      strings.ToLower(loan.WalletIndexed.Hex()),
		ID:              fmt.Sprintf("%v_%v", userAddress, txHash),
		PlayerAddress:   userAddress,
		EthIn:           endTx.EthIn.String(),
		KeysBought:      endTx.KeysBought.String(),
		WinnerAddr:      endTx.WinnerAddr.String(),
		AmountWon:       endTx.AmountWon.String(),
		NewPot:          endTx.NewPot.String(),
		P3DAmount:       endTx.P3DAmount.String(),
		GenAmount:       endTx.GenAmount.String(),
		PotAmount:       endTx.PotAmount.String(),
		AirDropPot:      endTx.AirDropPot.String(),
		BlockNum:        blockNum,
		TransactionHash: txHash,
		Timestamp:       timestamp,
	}
	_, err := m.DB.Model(l).Insert()
	return l, err
}

func (m *Manager) InsertDicePlayTx(playTx *dice.DiceDicePlayEvent, contract string, blockNum uint64, txHash string, timestamp uint64) (*models2.DicePlayTx, error) {
	userAddress := strings.ToLower(playTx.PlayerAddress.String())
	l := &models2.DicePlayTx{
		//CompressedData:      strings.ToLower(loan.WalletIndexed.Hex()),
		ID:              fmt.Sprintf("%v_%v", strings.ToLower(contract), playTx.RequestId),
		PlayerAddress:   userAddress,
		RequestId:       playTx.RequestId.Uint64(),
		IsSuccess:       0,
		BlockNum:        blockNum,
		TransactionHash: txHash,
		Timestamp:       timestamp,
	}
	_, err := m.DB.Model(l).Insert()
	return l, err
}
