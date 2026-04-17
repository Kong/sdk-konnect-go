# IntegrationInstanceAuthConfig

A response containing integration instance auth config.


## Supported Types

### OauthAuthConfig

```go
integrationInstanceAuthConfig := components.CreateIntegrationInstanceAuthConfigOauthAuthConfig(components.OauthAuthConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationInstanceAuthConfig.Type {
	case components.IntegrationInstanceAuthConfigTypeOauthAuthConfig:
		// integrationInstanceAuthConfig.OauthAuthConfig is populated
}
```
