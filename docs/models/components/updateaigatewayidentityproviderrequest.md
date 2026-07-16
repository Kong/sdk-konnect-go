# UpdateAIGatewayIdentityProviderRequest

**Pre-release Feature**
This feature is currently in beta and is subject to change.


## Supported Types

### AIGatewayIdentityProviderKeyAuth

```go
updateAIGatewayIdentityProviderRequest := components.CreateUpdateAIGatewayIdentityProviderRequestKeyAuth(components.AIGatewayIdentityProviderKeyAuth{/* values here */})
```

### AIGatewayIdentityProviderOpenIDConnect

```go
updateAIGatewayIdentityProviderRequest := components.CreateUpdateAIGatewayIdentityProviderRequestOpenidConnect(components.AIGatewayIdentityProviderOpenIDConnect{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateAIGatewayIdentityProviderRequest.Type {
	case components.UpdateAIGatewayIdentityProviderRequestTypeKeyAuth:
		// updateAIGatewayIdentityProviderRequest.AIGatewayIdentityProviderKeyAuth is populated
	case components.UpdateAIGatewayIdentityProviderRequestTypeOpenidConnect:
		// updateAIGatewayIdentityProviderRequest.AIGatewayIdentityProviderOpenIDConnect is populated
}
```
