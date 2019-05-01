package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vault-token-helper",
	Short: "Token helper for Hashicorp Vault",
	Long:  `Token helper for Hashicorp Vault`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.BindEnv("vault-addr", "VAULT_ADDR")
	if viper.GetString("vault-addr") == "" {
		log.Fatal("VAULT_ADDR environment variable must be set")
	}
}
