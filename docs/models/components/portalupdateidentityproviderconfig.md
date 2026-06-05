# PortalUpdateIdentityProviderConfig


## Supported Types

### OIDCIdentityProviderConfig

```go
portalUpdateIdentityProviderConfig := components.CreatePortalUpdateIdentityProviderConfigOIDCIdentityProviderConfig(components.OIDCIdentityProviderConfig{/* values here */})
```

### PortalSAMLIdentityProviderConfigInput

```go
portalUpdateIdentityProviderConfig := components.CreatePortalUpdateIdentityProviderConfigPortalSAMLIdentityProviderConfigInput(components.PortalSAMLIdentityProviderConfigInput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch portalUpdateIdentityProviderConfig.Type {
	case components.PortalUpdateIdentityProviderConfigTypeOIDCIdentityProviderConfig:
		// portalUpdateIdentityProviderConfig.OIDCIdentityProviderConfig is populated
	case components.PortalUpdateIdentityProviderConfigTypePortalSAMLIdentityProviderConfigInput:
		// portalUpdateIdentityProviderConfig.PortalSAMLIdentityProviderConfigInput is populated
}
```
