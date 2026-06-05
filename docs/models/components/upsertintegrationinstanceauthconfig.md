# UpsertIntegrationInstanceAuthConfig


## Supported Types

### OauthConfig

```go
upsertIntegrationInstanceAuthConfig := components.CreateUpsertIntegrationInstanceAuthConfigOauthConfig(components.OauthConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch upsertIntegrationInstanceAuthConfig.Type {
	case components.UpsertIntegrationInstanceAuthConfigTypeOauthConfig:
		// upsertIntegrationInstanceAuthConfig.OauthConfig is populated
}
```
