# PortalCreateIdentityProviderConfig


## Supported Types

### CreateOIDCIdentityProviderConfig

```go
portalCreateIdentityProviderConfig := components.CreatePortalCreateIdentityProviderConfigCreateOIDCIdentityProviderConfig(components.CreateOIDCIdentityProviderConfig{/* values here */})
```

### PortalSAMLIdentityProviderConfigInput

```go
portalCreateIdentityProviderConfig := components.CreatePortalCreateIdentityProviderConfigPortalSAMLIdentityProviderConfigInput(components.PortalSAMLIdentityProviderConfigInput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch portalCreateIdentityProviderConfig.Type {
	case components.PortalCreateIdentityProviderConfigTypeCreateOIDCIdentityProviderConfig:
		// portalCreateIdentityProviderConfig.CreateOIDCIdentityProviderConfig is populated
	case components.PortalCreateIdentityProviderConfigTypePortalSAMLIdentityProviderConfigInput:
		// portalCreateIdentityProviderConfig.PortalSAMLIdentityProviderConfigInput is populated
}
```
