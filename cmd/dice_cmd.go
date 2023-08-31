package cmd

import (
	"github.com/0xdefi-studio/indexer/dice"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(diceCmd)
	rootCmd.AddCommand(diceOutcomeCmd)
}

var diceCmd = &cobra.Command{
	Use:   dice.DICE,
	Short: "Run 0xdefi-studio indexer for dice",
	RunE: func(cmd *cobra.Command, args []string) error {
		walker, startBlock, step := IndexBase(
			dice.DICE,
			"CreateDiceSchema",
			"dice/diceAbi.json",
			"InitDiceABI",
			RPC_URL,
		)
		walker.SetDicePrivateKey(PRIVATE_KEY)
		//give me addresses = fomo3d.ZhartaETHAddressesHexList + []common.Address{fomo3d.ZhartaUSDCAddressHex, fomo3d.ZhartaETHSquiggledDaoAddressHex}
		var addresses = []common.Address{dice.DiceContractHex, dice.DiceContractV2Hex, dice.DiceMinuContractHex}
		//addresses = append(addresses, fomo3d.ZhartaETHAddressesHexList...)
		walker.Run(startBlock, step, addresses, dice.DICE, DISCORD_WEBHOOK_URL, TELEGRAM_BOT_KEY, TELEGRAM_CHAT_ID)
		return nil
	},
}

var diceOutcomeCmd = &cobra.Command{
	Use:   dice.DICE_OUTCOME,
	Short: "Run 0xdefi-studio indexer for dice outcome",
	RunE: func(cmd *cobra.Command, args []string) error {
		walker, startBlock, step := IndexBase(
			dice.DICE_OUTCOME,
			"CreateDiceOutcomeSchema",
			"dice/diceAbi.json",
			"InitDiceABI",
			RPC_URL,
		)
		walker.SetDicePrivateKey(PRIVATE_KEY)
		//give me addresses = fomo3d.ZhartaETHAddressesHexList + []common.Address{fomo3d.ZhartaUSDCAddressHex, fomo3d.ZhartaETHSquiggledDaoAddressHex}
		var addresses = []common.Address{dice.DiceContractHex, dice.DiceContractV2Hex, dice.DiceMinuContractHex}
		//addresses = append(addresses, fomo3d.ZhartaETHAddressesHexList...)
		walker.Run(startBlock, step, addresses, dice.DICE_OUTCOME, DISCORD_WEBHOOK_URL, TELEGRAM_BOT_KEY, TELEGRAM_CHAT_ID)
		return nil
	},
}
