# Vault Token Helper

**Requires `VAULT_ADDR` environment variable be set!**

- Encrypts tokens at rest in macOS Keychain (via [99designs/keyring][1])
- Stores a token for each Vault address

### Setup

To configure a token helper, edit (or create) the file `~/.vault` and add a line
similar to:

```
token_helper = "/path/to/vault-token-helper"
```

You will need to use the fully qualified path to the `vault-token-helper`
binary. The binary should be executable. For more information, refer to the
[Hashicorp Vault docs on Token Helpers][2].


[1]: https://github.com/99designs/keyring
[2]: https://www.vaultproject.io/docs/commands/token-helper.html
