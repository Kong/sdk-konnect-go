# IdentityProviderConfig


## Supported Types

### OIDCIdentityProviderConfig

```go
identityProviderConfig := components.CreateIdentityProviderConfigOIDCIdentityProviderConfig(components.OIDCIdentityProviderConfig{/* values here */})
```

### SAMLIdentityProviderConfig

```go
identityProviderConfig := components.CreateIdentityProviderConfigSAMLIdentityProviderConfig(components.SAMLIdentityProviderConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch identityProviderConfig.Type {
	case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
		// identityProviderConfig.OIDCIdentityProviderConfig is populated
	case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
		// identityProviderConfig.SAMLIdentityProviderConfig is populated
}
```
