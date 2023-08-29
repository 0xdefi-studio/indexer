package chain

import (
	"github.com/0xdefi-studio/indexer/bankroll"
	"github.com/0xdefi-studio/indexer/dice"
	"github.com/0xdefi-studio/indexer/fomo3d"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (m *Manager) ParseFomo3dOnEndTx(log types.Log) (*fomo3d.Fomo3dOnEndTx, error) {
	var onEndTx fomo3d.Fomo3dOnEndTx
	err := m.Fomo3dV1ABI.UnpackIntoInterface(&onEndTx, "onEndTx", log.Data)
	if err != nil {
		return nil, err
	}
	return &onEndTx, nil
}

func (m *Manager) ParseDicePlayTx(log types.Log) (*dice.DiceDicePlayEvent, error) {
	var playTx dice.DiceDicePlayEvent
	err := m.DiceABI.UnpackIntoInterface(&playTx, "Dice_Play_Event", log.Data)
	if err != nil {
		return nil, err
	}
	playTx.PlayerAddress = common.HexToAddress(log.Topics[1].Hex())
	return &playTx, nil
}

func (m *Manager) ParseBankrollDepositTx(log types.Log) (*bankroll.BankrollDeposit, error) {
	var deposit bankroll.BankrollDeposit
	err := m.BankrollABI.UnpackIntoInterface(&deposit, "Deposit", log.Data)
	if err != nil {
		return nil, err
	}
	deposit.Caller = common.HexToAddress(log.Topics[1].Hex())
	deposit.Owner = common.HexToAddress(log.Topics[2].Hex())
	return &deposit, nil
}

func (m *Manager) ParseBankrollWithdrawTx(log types.Log) (*bankroll.BankrollWithdraw, error) {
	var withdraw bankroll.BankrollWithdraw
	err := m.BankrollABI.UnpackIntoInterface(&withdraw, "Withdraw", log.Data)
	if err != nil {
		return nil, err
	}
	withdraw.Caller = common.HexToAddress(log.Topics[1].Hex())
	withdraw.Receiver = common.HexToAddress(log.Topics[2].Hex())
	withdraw.Owner = common.HexToAddress(log.Topics[3].Hex())
	return &withdraw, nil
}
