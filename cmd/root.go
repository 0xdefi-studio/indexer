package cmd

import (
	"fmt"
	"github.com/0xdefi-studio/indexer/chain"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var walker *chain.Manager

var (
	// Used for flags.
	cfgFile              string
	DISCORD_WEBHOOK_URL  string
	FOMO3D_START_BLOCK   int64
	RPC_URL              string
	DBUser               string
	DBPassword           string
	DBName               string
	DBPort               int
	DBHost               string
	Step                 int64
	IsDebug              bool
	TELEGRAM_BOT_KEY     string
	TELEGRAM_CHAT_ID     int64
	DICE_START_BLOCK     int64
	PRIVATE_KEY          string
	BANKROLL_START_BLOCK int64

	rootCmd = &cobra.Command{
		Use:   "0xdefi-studio",
		Short: "0xdefi-studio command line tool",
		Long:  `0xdefi-studio command line tool.`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		DISCORD_WEBHOOK_URL = viper.GetString("DISCORD_WEBHOOK_URL")
		FOMO3D_START_BLOCK = viper.GetInt64("FOMO3D_START_BLOCK")
		RPC_URL = viper.GetString("RPC_URL")
		DBName = viper.GetString("dbName")
		DBUser = viper.GetString("dbUser")
		DBPassword = viper.GetString("dbPassword")
		DBPort = viper.GetInt("dbPort")
		DBHost = viper.GetString("dbHost")
		Step = viper.GetInt64("Step")
		IsDebug = viper.GetBool("isDebug")
		TELEGRAM_BOT_KEY = viper.GetString("TELEGRAM_BOT_KEY")
		TELEGRAM_CHAT_ID = viper.GetInt64("TELEGRAM_CHAT_ID")
		DICE_START_BLOCK = viper.GetInt64("DICE_START_BLOCK")
		PRIVATE_KEY = viper.GetString("PRIVATE_KEY")
		BANKROLL_START_BLOCK = viper.GetInt64("BANKROLL_START_BLOCK")
	}
}
