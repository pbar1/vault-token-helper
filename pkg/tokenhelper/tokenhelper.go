package tokenhelper

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/99designs/keyring"
	"github.com/spf13/viper"
)

// Get retrieves a Vault token from the keyring and prints it to stdout
func Get() error {
	kr, err := openKeychain()
	if err != nil {
		return err
	}

	key := viper.GetString("vault-addr")

	item, err := kr.Get(key)
	if err != nil && err.Error() == "The specified item could not be found in the keyring." {
		fmt.Print()
		os.Exit(0)
	} else {
		return err
	}

	token := string(item.Data)
	fmt.Fprintf(os.Stdout, "%s", token)

	return nil
}

// Store puts a Vault token in the keyring
func Store() error {
	r := bufio.NewReader(os.Stdin)
	token, err := r.ReadString('\n')
	if err != nil && err != io.EOF {
		return err
	}

	kr, err := openKeychain()
	if err != nil {
		return err
	}

	key := viper.GetString("vault-addr")

	err = kr.Set(keyring.Item{
		Key:         key,
		Data:        []byte(token),
		Description: "Hashicorp Vault token",
	})
	if err != nil {
		return err
	}

	return nil
}

// Erase deletes a Vault token from the keyring
func Erase() error {
	kr, err := openKeychain()
	if err != nil {
		return err
	}

	key := viper.GetString("vault-addr")

	err = kr.Remove(key)
	if err != nil {
		return err
	}

	return nil
}

func openKeychain() (keyring.Keyring, error) {
	kr, err := keyring.Open(keyring.Config{
		KeychainName:             "login",
		KeychainTrustApplication: true,
	})

	if err != nil {
		return nil, err
	}

	return kr, nil
}
