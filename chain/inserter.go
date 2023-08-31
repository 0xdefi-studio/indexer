package chain

import (
	"fmt"
	"github.com/0xdefi-studio/indexer/bankroll"
	"github.com/0xdefi-studio/indexer/dice"
	"github.com/0xdefi-studio/indexer/fomo3d"
	models2 "github.com/0xdefi-studio/indexer/models"
	"math/big"
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

func getBankrollTokenType(contract string) string {
	if contract == bankroll.BANKROLL_WMNT_ADDRESS {
		return bankroll.BANKROLL_WMNT
	}
	return bankroll.BANKROLL_MINU
}

func (m *Manager) GetBankrollPrevBalance(tokenType string) (*big.Int, error) {
	prevBalance := big.NewInt(0)
	var prevBankroll models2.BankrollTx
	// get prevBankroll by tokenType from db
	err := m.DB.Model(&prevBankroll).Where("token_type = ?", tokenType).Order("block_num desc").First()
	if err == nil {
		var ok bool
		prevBalance, ok = prevBalance.SetString(prevBankroll.CurrentBalance, 10)
		if ok == false {
			fmt.Println("InsertBankrollDepositTx error: SetString!")
			return nil, err
		}
	}
	return prevBalance, nil
}

func (m *Manager) GetBankrollSenderPrevBalance(tokenType string, userAddress string) (*big.Int, error) {
	prevBalance := big.NewInt(0)
	var prevBankroll models2.BankrollTx
	// get prevBankroll by tokenType from db
	err := m.DB.Model(&prevBankroll).Where("token_type = ? and sender = ?", tokenType, userAddress).Order("block_num desc").First()
	if err == nil {
		var ok bool
		prevBalance, ok = prevBalance.SetString(prevBankroll.CurrentSenderBalance, 10)
		if ok == false {
			fmt.Println("GetBankrollSenderPrevBalance error: SetString!")
			return nil, err
		}
	}
	return prevBalance, nil
}

func (m *Manager) GetDiceOutcomePrevBalance(tokenAddress string) (*big.Int, error) {
	prevBalance := big.NewInt(0)
	var prevDiceOutcome models2.DiceOutcome
	// get prevBankroll by tokenType from db
	err := m.DB.Model(&prevDiceOutcome).Where("contract = ?", tokenAddress).Order("block_num desc").First()
	if err == nil {
		var ok bool
		prevBalance, ok = prevBalance.SetString(prevDiceOutcome.CurrentBalance, 10)
		if ok == false {
			fmt.Println("GetDiceOutcomePrevBalance error: SetString!")
			return nil, err
		}
	}
	return prevBalance, nil
}

func (m *Manager) InsertBankrollDepositTx(depositTx *bankroll.BankrollDeposit, contract string, blockNum uint64, txHash string, timestamp uint64) (*models2.BankrollTx, error) {
	// 1.0 get sender
	sender := strings.ToLower(depositTx.Caller.String())
	// 2.0 get tokenType like wmnt or minu
	tokenType := getBankrollTokenType(contract)
	// 3.0 get prevBalance to calculate currentBalance
	prevBalance, err1 := m.GetBankrollPrevBalance(tokenType)
	prevSenderBalance, err2 := m.GetBankrollSenderPrevBalance(tokenType, sender)
	if err1 != nil {
		fmt.Println("InsertBankrollDepositTx error: GetBankrollPrevBalance!")
		return nil, err1
	}
	if err2 != nil {
		fmt.Println("InsertBankrollDepositTx error: InsertBankrollDepositTx!")
		return nil, err1
	}
	// 3.2 calculate currentBalance by adding because it's deposit
	prevBalance.Add(prevBalance, depositTx.Assets)
	prevSenderBalance.Add(prevSenderBalance, depositTx.Assets)
	l := &models2.BankrollTx{
		//CompressedData:      strings.ToLower(loan.WalletIndexed.Hex()),
		ID:        fmt.Sprintf("%v_%v_%v", sender, txHash[:8], contract),
		Sender:    sender,
		Owner:     depositTx.Owner.String(),
		Assets:    depositTx.Assets.String(),
		Shares:    depositTx.Shares.String(),
		TokenType: tokenType,
		TxType:    "deposit",

		CurrentBalance:       prevBalance.String(),
		CurrentSenderBalance: prevSenderBalance.String(),

		BlockNum:        blockNum,
		TransactionHash: txHash,
		Timestamp:       timestamp,
	}
	_, err := m.DB.Model(l).Insert()
	return l, err
}

func (m *Manager) InsertBankrollWithdrawTx(withdrawTx *bankroll.BankrollWithdraw, contract string, blockNum uint64, txHash string, timestamp uint64) (*models2.BankrollTx, error) {
	// 1.0 get sender
	sender := strings.ToLower(withdrawTx.Caller.String())
	// 2.0 get tokenType like wmnt or minu
	tokenType := getBankrollTokenType(contract)
	// 3.0 get prevBalance to calculate currentBalance
	prevBalance, err1 := m.GetBankrollPrevBalance(tokenType)
	prevSenderBalance, err2 := m.GetBankrollSenderPrevBalance(tokenType, sender)
	if err1 != nil {
		fmt.Println("InsertBankrollDepositTx error: GetBankrollPrevBalance!")
		return nil, err1
	}
	if err2 != nil {
		fmt.Println("InsertBankrollDepositTx error: InsertBankrollDepositTx!")
		return nil, err1
	}
	// 3.2 calculate currentBalance by adding because it's deposit
	prevBalance.Sub(prevBalance, withdrawTx.Assets)
	prevSenderBalance.Sub(prevSenderBalance, withdrawTx.Assets)
	l := &models2.BankrollTx{
		//CompressedData:      strings.ToLower(loan.WalletIndexed.Hex()),
		ID:        fmt.Sprintf("%v_%v_%v", sender, txHash[:8], tokenType),
		Sender:    sender,
		Receiver:  withdrawTx.Receiver.String(),
		Owner:     withdrawTx.Owner.String(),
		Assets:    withdrawTx.Assets.String(),
		Shares:    withdrawTx.Shares.String(),
		TokenType: tokenType,
		TxType:    "withdraw",

		CurrentBalance:       prevBalance.String(),
		CurrentSenderBalance: prevSenderBalance.String(),

		BlockNum:        blockNum,
		TransactionHash: txHash,
		Timestamp:       timestamp,
	}
	_, err := m.DB.Model(l).Insert()
	return l, err
}

func (m *Manager) InsertDiceOutcomeTx(outcome *dice.DiceDiceOutcomeEvent, contract string, blockNum uint64, txHash string, timestamp uint64) (*models2.DiceOutcome, error) {
	// 1.0 get sender
	sender := strings.ToLower(outcome.PlayerAddress.String())
	tokenAddress := strings.ToLower(outcome.TokenAddress.String())
	// 2.0 get prevBalance to calculate currentBalance
	prevBalance, err := m.GetDiceOutcomePrevBalance(contract)
	if err != nil {
		fmt.Println("InsertDiceOutcomeTx error: prevBalance fetching!")
		return nil, err
	}
	// 3.2 calculate currentBalance by adding because it's deposit
	currBalanceChange := big.NewInt(0)
	currBalanceChange.Sub(outcome.Wager, outcome.Payout)
	prevBalance.Add(prevBalance, currBalanceChange)
	l := &models2.DiceOutcome{
		//CompressedData:      strings.ToLower(loan.WalletIndexed.Hex()),
		ID:            fmt.Sprintf("%v_%v_%v", sender, txHash[:8], contract),
		PlayerAddress: sender,
		Wager:         outcome.Wager.String(),
		Payout:        outcome.Payout.String(),
		TokenAddress:  tokenAddress,
		Contract:      contract,

		CurrentBalance: prevBalance.String(),

		BlockNum:        blockNum,
		TransactionHash: txHash,
		Timestamp:       timestamp,
	}
	_, err = m.DB.Model(l).Insert()
	return l, err
}
