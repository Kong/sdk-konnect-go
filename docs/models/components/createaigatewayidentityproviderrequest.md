# CreateAIGatewayIdentityProviderRequest


## Supported Types

### AIGatewayIdentityProviderKeyAuth

```go
createAIGatewayIdentityProviderRequest := components.CreateCreateAIGatewayIdentityProviderRequestKeyAuth(components.AIGatewayIdentityProviderKeyAuth{/* values here */})
```

### AIGatewayIdentityProviderOpenIDConnect

```go
createAIGatewayIdentityProviderRequest := components.CreateCreateAIGatewayIdentityProviderRequestOpenidConnect(components.AIGatewayIdentityProviderOpenIDConnect{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createAIGatewayIdentityProviderRequest.Type {
	case components.CreateAIGatewayIdentityProviderRequestTypeKeyAuth:
		// createAIGatewayIdentityProviderRequest.AIGatewayIdentityProviderKeyAuth is populated
	case components.CreateAIGatewayIdentityProviderRequestTypeOpenidConnect:
		// createAIGatewayIdentityProviderRequest.AIGatewayIdentityProviderOpenIDConnect is populated
}
```
