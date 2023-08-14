package cmd

import (
	"github.com/0xdefi-studio/indexer/fomo3d"
	"github.com/spf13/cobra"
)

var platform string
var borrower string

func init() {
	queryCmd.Flags().StringVarP(&platform, "platform", "p", "", "platform")
	queryCmd.Flags().StringVarP(&borrower, "borrower", "b", "", "borrower")
	rootCmd.AddCommand(queryCmd)
}

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Run 0xdefi-studio indexer",
	RunE: func(cmd *cobra.Command, args []string) error {
		switch platform {
		case "fomo3d":
			walker, _, _ = IndexBase(
				fomo3d.FOMO3DV1,
				"CreateFomo3dV1Schema",
				"fomo3d/fomo3d.json",
				"InitFomo3dV1ABI",
				RPC_URL,
			)

		}

		return nil
	},
}
