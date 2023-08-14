package cmd

import (
	"encoding/json"
	"github.com/0xdefi-studio/indexer/utils"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init() {
	rootCmd.AddCommand(abimerge)
}

var abimerge = &cobra.Command{
	Use:   "abimerge",
	Short: "Merge abi from different abi files",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			log.Println("Wrong argument")
			os.Exit(1)
		}
		log.Printf("Start scan abi files from %s\n", args[0])
		abis := utils.ReadABI(args[0])
		merged := utils.MergeABI(abis)
		m, _ := json.Marshal(merged)
		log.Println(string(m))
		return nil
	},
}
