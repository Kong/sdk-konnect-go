# AIGatewayIdentityProvider

**Pre-release Feature**
This feature is currently in beta and is subject to change.


## Supported Types

### AIGatewayIdentityProviderKeyAuthResponse

```go
aiGatewayIdentityProvider := components.CreateAIGatewayIdentityProviderKeyAuth(components.AIGatewayIdentityProviderKeyAuthResponse{/* values here */})
```

### AIGatewayIdentityProviderOpenIDConnectResponse

```go
aiGatewayIdentityProvider := components.CreateAIGatewayIdentityProviderOpenidConnect(components.AIGatewayIdentityProviderOpenIDConnectResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayIdentityProvider.Type {
	case components.AIGatewayIdentityProviderTypeKeyAuth:
		// aiGatewayIdentityProvider.AIGatewayIdentityProviderKeyAuthResponse is populated
	case components.AIGatewayIdentityProviderTypeOpenidConnect:
		// aiGatewayIdentityProvider.AIGatewayIdentityProviderOpenIDConnectResponse is populated
}
```
