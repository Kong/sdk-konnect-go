# Configs

JSON-B object containing the configuration for the OIDC strategy under the key 'openid-connect' or the configuration for the Key Auth strategy under the key 'key-auth'


## Supported Types

### UpdateAppAuthStrategyRequestOpenIDConnect

```go
configs := components.CreateConfigsUpdateAppAuthStrategyRequestOpenIDConnect(components.UpdateAppAuthStrategyRequestOpenIDConnect{/* values here */})
```

### UpdateAppAuthStrategyRequestKeyAuth

```go
configs := components.CreateConfigsUpdateAppAuthStrategyRequestKeyAuth(components.UpdateAppAuthStrategyRequestKeyAuth{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch configs.Type {
	case components.ConfigsTypeUpdateAppAuthStrategyRequestOpenIDConnect:
		// configs.UpdateAppAuthStrategyRequestOpenIDConnect is populated
	case components.ConfigsTypeUpdateAppAuthStrategyRequestKeyAuth:
		// configs.UpdateAppAuthStrategyRequestKeyAuth is populated
}
```
