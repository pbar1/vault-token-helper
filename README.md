# Vault Token Helper

**Requires `VAULT_ADDR` environment variable be set**

Place the binary somewhere on your `PATH`, then run the following to configure Vault to use the token helper:

```
vault-token-helper hook
```

Currently only supports macOS Keychain as backend
