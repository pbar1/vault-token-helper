package command

import (
	"log"

	"github.com/pbar1/vault-token-helper/internal/tokenhelper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

)

var eraseCmd = &cobra.Command{
	Use:   "erase",
	Short: "Erases a token from the keyring",
	Long:  `Erases the token for the current Vault from the keyring.`,
	Run: func(cmd *cobra.Command, args []string) {
		vaultAddr := viper.GetString("vault-addr")

		keyringHelper, err := tokenhelper.NewKeyringHelper(vaultAddr)
		if err != nil {
			log.Fatalf("unable to create token helper: %v\n", err)
		}

		err = keyringHelper.Erase()
		if err != nil {
			log.Fatalf("unable to erase token: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(eraseCmd)
}
