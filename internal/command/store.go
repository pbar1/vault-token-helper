package command

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/pbar1/vault-token-helper/internal/tokenhelper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Stores a token in the keyring",
	Long:  `Stores a token for the current Vault in the keyring. Expects a token to be written to STDIN.`,
	Run: func(cmd *cobra.Command, args []string) {
		vaultAddr := viper.GetString("vault-addr")

		keyringHelper, err := tokenhelper.NewKeyringHelper(vaultAddr)
		if err != nil {
			log.Fatalf("unable to create token helper: %v\n", err)
		}

		reader := bufio.NewReader(os.Stdin)
		token, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			log.Fatalf("unable to read token from stdin: %v\n", err)
		}

		err = keyringHelper.Store(token)
		if err != nil {
			log.Fatalf("unable to store token: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)
}
