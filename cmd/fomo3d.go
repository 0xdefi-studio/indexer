package cmd

import (
	"github.com/0xdefi-studio/indexer/fomo3d"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(fomo3dV1Cmd)
}

var fomo3dV1Cmd = &cobra.Command{
	Use:   "fomo3dV1",
	Short: "Run 0xdefi-studio indexer for fomo3dV1",
	RunE: func(cmd *cobra.Command, args []string) error {
		walker, startBlock, step := IndexBase(
			fomo3d.FOMO3DV1,
			"CreateFomo3dV1Schema",
			"fomo3d/fomo3d.json",
			"InitFomo3dV1ABI",
			RPC_URL,
		)
		//give me addresses = fomo3d.ZhartaETHAddressesHexList + []common.Address{fomo3d.ZhartaUSDCAddressHex, fomo3d.ZhartaETHSquiggledDaoAddressHex}
		var addresses = []common.Address{fomo3d.FOMO3DContractV1Hex}
		//addresses = append(addresses, fomo3d.ZhartaETHAddressesHexList...)
		walker.Run(startBlock, step, addresses, fomo3d.FOMO3DV1, DISCORD_WEBHOOK_URL, TELEGRAM_BOT_KEY, TELEGRAM_CHAT_ID)
		return nil
	},
}
