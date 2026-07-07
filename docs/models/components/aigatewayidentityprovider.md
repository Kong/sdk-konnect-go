# AIGatewayIdentityProvider


## Supported Types

### AIGatewayIdentityProviderKeyAuthAIGatewayIdentityProviderKeyAuthConfig

```go
aiGatewayIdentityProvider := components.CreateAIGatewayIdentityProviderKeyAuth(components.AIGatewayIdentityProviderKeyAuthAIGatewayIdentityProviderKeyAuthConfig{/* values here */})
```

### AIGatewayIdentityProviderOpenIDConnectAIGatewayIdentityProviderOpenIDConnectConfig

```go
aiGatewayIdentityProvider := components.CreateAIGatewayIdentityProviderOpenidConnect(components.AIGatewayIdentityProviderOpenIDConnectAIGatewayIdentityProviderOpenIDConnectConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayIdentityProvider.Type {
	case components.AIGatewayIdentityProviderTypeKeyAuth:
		// aiGatewayIdentityProvider.AIGatewayIdentityProviderKeyAuthAIGatewayIdentityProviderKeyAuthConfig is populated
	case components.AIGatewayIdentityProviderTypeOpenidConnect:
		// aiGatewayIdentityProvider.AIGatewayIdentityProviderOpenIDConnectAIGatewayIdentityProviderOpenIDConnectConfig is populated
}
```
