# PortalCreateIdentityProviderConfig


## Supported Types

### OIDCIdentityProviderConfig

```go
portalCreateIdentityProviderConfig := components.CreatePortalCreateIdentityProviderConfigOIDCIdentityProviderConfig(components.OIDCIdentityProviderConfig{/* values here */})
```

### PortalSAMLIdentityProviderConfigInput

```go
portalCreateIdentityProviderConfig := components.CreatePortalCreateIdentityProviderConfigPortalSAMLIdentityProviderConfigInput(components.PortalSAMLIdentityProviderConfigInput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch portalCreateIdentityProviderConfig.Type {
	case components.PortalCreateIdentityProviderConfigTypeOIDCIdentityProviderConfig:
		// portalCreateIdentityProviderConfig.OIDCIdentityProviderConfig is populated
	case components.PortalCreateIdentityProviderConfigTypePortalSAMLIdentityProviderConfigInput:
		// portalCreateIdentityProviderConfig.PortalSAMLIdentityProviderConfigInput is populated
}
```
