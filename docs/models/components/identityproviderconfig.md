# IdentityProviderConfig


## Supported Types

### OIDCIdentityProviderConfigOutput

```go
identityProviderConfig := components.CreateIdentityProviderConfigOIDCIdentityProviderConfigOutput(components.OIDCIdentityProviderConfigOutput{/* values here */})
```

### SAMLIdentityProviderConfig

```go
identityProviderConfig := components.CreateIdentityProviderConfigSAMLIdentityProviderConfig(components.SAMLIdentityProviderConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch identityProviderConfig.Type {
	case components.IdentityProviderConfigTypeOIDCIdentityProviderConfigOutput:
		// identityProviderConfig.OIDCIdentityProviderConfigOutput is populated
	case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
		// identityProviderConfig.SAMLIdentityProviderConfig is populated
}
```
