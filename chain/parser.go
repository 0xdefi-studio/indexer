package chain

import (
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
