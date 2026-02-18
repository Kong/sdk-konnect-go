# UpdateIdentityProviderConfig


## Supported Types

### ConfigureOIDCIdentityProviderConfig

```go
updateIdentityProviderConfig := components.CreateUpdateIdentityProviderConfigConfigureOIDCIdentityProviderConfig(components.ConfigureOIDCIdentityProviderConfig{/* values here */})
```

### SAMLIdentityProviderConfigInput

```go
updateIdentityProviderConfig := components.CreateUpdateIdentityProviderConfigSAMLIdentityProviderConfigInput(components.SAMLIdentityProviderConfigInput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateIdentityProviderConfig.Type {
	case components.UpdateIdentityProviderConfigTypeConfigureOIDCIdentityProviderConfig:
		// updateIdentityProviderConfig.ConfigureOIDCIdentityProviderConfig is populated
	case components.UpdateIdentityProviderConfigTypeSAMLIdentityProviderConfigInput:
		// updateIdentityProviderConfig.SAMLIdentityProviderConfigInput is populated
}
```
