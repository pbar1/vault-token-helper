package cmd

import (
	"fmt"
	"log"

	"github.com/pbar1/vault-token-helper/pkg/tokenhelper"
	"github.com/spf13/cobra"
)

// hookCmd represents the hook command
var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Configures the Vault CLI to use this token helper",
	Long:  `Configures the Vault CLI to use this token helper`,
	Run: func(cmd *cobra.Command, args []string) {
		err := tokenhelper.Hook()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Vault is now configured to use this token helper")
	},
}

func init() {
	rootCmd.AddCommand(hookCmd)
}
