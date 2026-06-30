# CreateIdentityProviderConfig


## Supported Types

### OIDCIdentityProviderConfig

```go
createIdentityProviderConfig := components.CreateCreateIdentityProviderConfigOIDCIdentityProviderConfig(components.OIDCIdentityProviderConfig{/* values here */})
```

### SAMLIdentityProviderConfigInput

```go
createIdentityProviderConfig := components.CreateCreateIdentityProviderConfigSAMLIdentityProviderConfigInput(components.SAMLIdentityProviderConfigInput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createIdentityProviderConfig.Type {
	case components.CreateIdentityProviderConfigTypeOIDCIdentityProviderConfig:
		// createIdentityProviderConfig.OIDCIdentityProviderConfig is populated
	case components.CreateIdentityProviderConfigTypeSAMLIdentityProviderConfigInput:
		// createIdentityProviderConfig.SAMLIdentityProviderConfigInput is populated
}
```
