# PortalIdentityProviderConfig

The configuration data for the authentication integration.


## Supported Types

### PortalOIDCIdentityProviderConfig

```go
portalIdentityProviderConfig := components.CreatePortalIdentityProviderConfigPortalOIDCIdentityProviderConfig(components.PortalOIDCIdentityProviderConfig{/* values here */})
```

### PortalSAMLIdentityProviderConfig

```go
portalIdentityProviderConfig := components.CreatePortalIdentityProviderConfigPortalSAMLIdentityProviderConfig(components.PortalSAMLIdentityProviderConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch portalIdentityProviderConfig.Type {
	case components.PortalIdentityProviderConfigTypePortalOIDCIdentityProviderConfig:
		// portalIdentityProviderConfig.PortalOIDCIdentityProviderConfig is populated
	case components.PortalIdentityProviderConfigTypePortalSAMLIdentityProviderConfig:
		// portalIdentityProviderConfig.PortalSAMLIdentityProviderConfig is populated
}
```
