package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(telegramCmd)
}

var telegramCmd = &cobra.Command{
	Use:   "telegram",
	Short: "Run telegram server for TG-notification",
	RunE: func(cmd *cobra.Command, args []string) error {
		walker, _, _ := IndexBase(
			"telegram",
			"CreateTelegramSchema",
			"fomo3d/fomo3d.json",
			"InitFomo3dV1ABI",
			RPC_URL,
		)
		walker.TelegramServerRun(TELEGRAM_BOT_KEY)
		return nil
	},
}
