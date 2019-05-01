package tokenhelper

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/ini.v1"
)

// Hook configures the Vault CLI to use this token helper
func Hook() error {
	selfPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	homePath, err := homedir.Dir()
	if err != nil {
		return err
	}

	vaultCfgPath := filepath.Join(homePath, ".vault")
	if _, err := os.Stat(vaultCfgPath); os.IsNotExist(err) {
		f, err := os.Create(vaultCfgPath)
		if err != nil {
			return err
		}
		f.Close()
	}

	cfg, err := ini.Load(vaultCfgPath)
	if err != nil {
		return err
	}

	// I know this is HCL and not INI, but it works
	cfg.Section("").Key("token_helper").SetValue("\"" + selfPath + "\"")
	cfg.SaveTo(vaultCfgPath)
	if err != nil {
		return err
	}

	return nil
}
