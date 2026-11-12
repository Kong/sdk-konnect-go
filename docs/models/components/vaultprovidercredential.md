# VaultProviderCredential

A provider's shared secret, or one principal's authorization of the provider, indicated by credential_type. Credential values are never returned.


## Supported Types

### VaultProviderCredentialShared

```go
vaultProviderCredential := components.CreateVaultProviderCredentialShared(components.VaultProviderCredentialShared{/* values here */})
```

### VaultProviderCredentialUser

```go
vaultProviderCredential := components.CreateVaultProviderCredentialUser(components.VaultProviderCredentialUser{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch vaultProviderCredential.Type {
	case components.VaultProviderCredentialTypeShared:
		// vaultProviderCredential.VaultProviderCredentialShared is populated
	case components.VaultProviderCredentialTypeUser:
		// vaultProviderCredential.VaultProviderCredentialUser is populated
}
```
