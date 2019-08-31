package command

import (
	"fmt"
	"log"

	"github.com/pbar1/vault-token-helper/internal/tokenhelper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a token from the keyring",
	Long:  `Gets the token for the current Vault from the keyring. Prints this token to STDOUT.`,
	Run: func(cmd *cobra.Command, args []string) {
		vaultAddr := viper.GetString("vault-addr")

		keyringHelper, err := tokenhelper.NewKeyringHelper(vaultAddr)
		if err != nil {
			log.Fatalf("unable to create token helper: %v\n", err)
		}

		token, err := keyringHelper.Get()
		if err != nil {
			log.Fatalf("unable to get token: %v\n", err)
		}
		fmt.Printf("%s", token)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
