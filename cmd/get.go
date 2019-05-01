package cmd

import (
	"log"

	"github.com/pbar1/vault-token-helper/pkg/tokenhelper"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves a token from the keyring",
	Long:  `Retrieves a token from the keyring`,
	Run: func(cmd *cobra.Command, args []string) {
		err := tokenhelper.Get()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
