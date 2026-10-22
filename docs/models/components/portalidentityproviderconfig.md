# PortalIdentityProviderConfig


## Supported Types

### PortalOIDCIdentityProviderConfigOutput

```go
portalIdentityProviderConfig := components.CreatePortalIdentityProviderConfigPortalOIDCIdentityProviderConfigOutput(components.PortalOIDCIdentityProviderConfigOutput{/* values here */})
```

### PortalSAMLIdentityProviderConfig

```go
portalIdentityProviderConfig := components.CreatePortalIdentityProviderConfigPortalSAMLIdentityProviderConfig(components.PortalSAMLIdentityProviderConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch portalIdentityProviderConfig.Type {
	case components.PortalIdentityProviderConfigTypePortalOIDCIdentityProviderConfigOutput:
		// portalIdentityProviderConfig.PortalOIDCIdentityProviderConfigOutput is populated
	case components.PortalIdentityProviderConfigTypePortalSAMLIdentityProviderConfig:
		// portalIdentityProviderConfig.PortalSAMLIdentityProviderConfig is populated
}
```
