# CreateIdentityProviderConfig


## Supported Types

### ConfigureOIDCIdentityProviderConfig

```go
createIdentityProviderConfig := components.CreateCreateIdentityProviderConfigConfigureOIDCIdentityProviderConfig(components.ConfigureOIDCIdentityProviderConfig{/* values here */})
```

### SAMLIdentityProviderConfigInput

```go
createIdentityProviderConfig := components.CreateCreateIdentityProviderConfigSAMLIdentityProviderConfigInput(components.SAMLIdentityProviderConfigInput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createIdentityProviderConfig.Type {
	case components.CreateIdentityProviderConfigTypeConfigureOIDCIdentityProviderConfig:
		// createIdentityProviderConfig.ConfigureOIDCIdentityProviderConfig is populated
	case components.CreateIdentityProviderConfigTypeSAMLIdentityProviderConfigInput:
		// createIdentityProviderConfig.SAMLIdentityProviderConfigInput is populated
}
```
