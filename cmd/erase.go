package cmd

import (
	"log"

	"github.com/pbar1/vault-token-helper/pkg/tokenhelper"
	"github.com/spf13/cobra"
)

// eraseCmd represents the erase command
var eraseCmd = &cobra.Command{
	Use:   "erase",
	Short: "Deletes a token from the keyring",
	Long:  `Deletes a token from the keyring`,
	Run: func(cmd *cobra.Command, args []string) {
		err := tokenhelper.Erase()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(eraseCmd)
}
