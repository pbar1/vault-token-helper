package command

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "vault-token-helper",
	Short: "Token helper for Hashicorp Vault",
	Long:  `Token helper for Hashicorp Vault`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version, gitCommit, buildDate string) {
	viper.Set("version", version)
	viper.Set("git-commit", gitCommit)
	viper.Set("build-date", buildDate)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("unable to execute root command: %v\n", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	err := viper.BindEnv("vault-addr", "VAULT_ADDR")
	if err != nil {
		log.Fatalf("unable to bind environment variable VAULT_ADDR: %v\n", err)
	}

	if viper.GetString("vault-addr") == "" {
		log.Fatalln("VAULT_ADDR environment variable must be set")
	}
}
