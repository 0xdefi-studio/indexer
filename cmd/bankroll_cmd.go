package cmd

import (
	"github.com/0xdefi-studio/indexer/bankroll"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(bankrollCmd)
}

var bankrollCmd = &cobra.Command{
	Use:   bankroll.BANKROLL,
	Short: "Run 0xdefi-studio indexer for bankroll",
	RunE: func(cmd *cobra.Command, args []string) error {
		walker, startBlock, step := IndexBase(
			bankroll.BANKROLL,
			"CreateBankrollSchema",
			"bankroll/bankrollAbi.json",
			"InitBankrollABI",
			RPC_URL,
		)
		var addresses = []common.Address{bankroll.BankrollWmntContractHex, bankroll.BankrollMinuContractHex}
		walker.Run(startBlock, step, addresses, bankroll.BANKROLL, DISCORD_WEBHOOK_URL, TELEGRAM_BOT_KEY, TELEGRAM_CHAT_ID)
		return nil
	},
}
