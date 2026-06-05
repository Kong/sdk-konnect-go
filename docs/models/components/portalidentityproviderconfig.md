# PortalIdentityProviderConfig


## Supported Types

### OIDCIdentityProviderConfigOutput

```go
portalIdentityProviderConfig := components.CreatePortalIdentityProviderConfigOIDCIdentityProviderConfigOutput(components.OIDCIdentityProviderConfigOutput{/* values here */})
```

### PortalSAMLIdentityProviderConfig

```go
portalIdentityProviderConfig := components.CreatePortalIdentityProviderConfigPortalSAMLIdentityProviderConfig(components.PortalSAMLIdentityProviderConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch portalIdentityProviderConfig.Type {
	case components.PortalIdentityProviderConfigTypeOIDCIdentityProviderConfigOutput:
		// portalIdentityProviderConfig.OIDCIdentityProviderConfigOutput is populated
	case components.PortalIdentityProviderConfigTypePortalSAMLIdentityProviderConfig:
		// portalIdentityProviderConfig.PortalSAMLIdentityProviderConfig is populated
}
```
