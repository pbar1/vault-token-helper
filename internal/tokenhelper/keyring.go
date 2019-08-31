package tokenhelper

import "github.com/99designs/keyring"

type KeyringHelper struct {
	kr keyring.Keyring
	vaultAddr string
}

// NewKeyringHelper creates a KeyringHelper, an implementation of the Vault token
// helper interface using 99designs Keyring
func NewKeyringHelper(vaultAddr string) (*KeyringHelper, error) {
	backends := keyring.AvailableBackends()

	// TODO: figure out sane defaults for more than just macOS Keychain
	ring, err := keyring.Open(keyring.Config{
		AllowedBackends:                backends,
		ServiceName:                    "vault-token-helper",
		KeychainName:                   "login",
		KeychainTrustApplication:       true,
		KeychainSynchronizable:         false,
		KeychainAccessibleWhenUnlocked: false,
		KeychainPasswordFunc:           nil,
		FilePasswordFunc:               nil,
		FileDir:                        "",
		KWalletAppID:                   "",
		KWalletFolder:                  "",
		LibSecretCollectionName:        "",
		PassDir:                        "",
		PassCmd:                        "",
		PassPrefix:                     "",
		WinCredPrefix:                  "",
	})
	if err != nil {
		return nil, err
	}

	return &KeyringHelper{
		kr:        ring,
		vaultAddr: vaultAddr,
	}, nil
}

// Get gets a token from the keyring
func (h KeyringHelper) Get() ([]byte, error) {
	item, err := h.kr.Get(h.vaultAddr)
	if err == keyring.ErrKeyNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return item.Data, nil
}

// Store stores a token in the keyring
func (h KeyringHelper) Store(token []byte) error {
	err := h.kr.Set(keyring.Item{
		Key:                         h.vaultAddr,
		Data:                        token,
		Label:                       "vault-token-helper",
		Description:                 "Hashicorp Vault token",
	})
	return err
}

// Erase erases a token from the keyring
func (h KeyringHelper) Erase() error {
	err := h.kr.Remove(h.vaultAddr)
	return err
}
