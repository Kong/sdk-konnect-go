# PortalUpdateIdentityProviderConfig


## Supported Types

### UpdateOIDCIdentityProviderConfig

```go
portalUpdateIdentityProviderConfig := components.CreatePortalUpdateIdentityProviderConfigUpdateOIDCIdentityProviderConfig(components.UpdateOIDCIdentityProviderConfig{/* values here */})
```

### UpdateSAMLIdentityProviderConfig

```go
portalUpdateIdentityProviderConfig := components.CreatePortalUpdateIdentityProviderConfigUpdateSAMLIdentityProviderConfig(components.UpdateSAMLIdentityProviderConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch portalUpdateIdentityProviderConfig.Type {
	case components.PortalUpdateIdentityProviderConfigTypeUpdateOIDCIdentityProviderConfig:
		// portalUpdateIdentityProviderConfig.UpdateOIDCIdentityProviderConfig is populated
	case components.PortalUpdateIdentityProviderConfigTypeUpdateSAMLIdentityProviderConfig:
		// portalUpdateIdentityProviderConfig.UpdateSAMLIdentityProviderConfig is populated
}
```
