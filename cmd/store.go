package cmd

import (
	"log"

	"github.com/pbar1/vault-token-helper/pkg/tokenhelper"
	"github.com/spf13/cobra"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Puts a token in the keyring",
	Long:  `Puts a token in the keyring`,
	Run: func(cmd *cobra.Command, args []string) {
		err := tokenhelper.Store()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)
}
