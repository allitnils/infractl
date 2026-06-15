package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var outputFormat string
var subscriptionID string

var rootCmd = &cobra.Command{
	Use:   "infractl",
	Short: "Azure infrastructure auditing CLI",
	Long:  `infractl audits Azure infrastructure for compliance, drift, cost, and policy violations.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: $HOME/.infractl.yaml)")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "output format: table, json")
	rootCmd.PersistentFlags().StringVarP(&subscriptionID, "subscription", "s", "", "Azure subscription ID")
	_ = viper.BindPFlag("subscription", rootCmd.PersistentFlags().Lookup("subscription"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".infractl")
	}
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
}
