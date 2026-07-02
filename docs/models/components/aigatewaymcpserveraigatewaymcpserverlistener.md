# AIGatewayMCPServerAIGatewayMCPServerListener


## Supported Types

### AIGatewayMCPServerListenerAIGatewayMCPServerBaseACLPropertiesConsumer

```go
aiGatewayMCPServerAIGatewayMCPServerListener := components.CreateAIGatewayMCPServerAIGatewayMCPServerListenerConsumer(components.AIGatewayMCPServerListenerAIGatewayMCPServerBaseACLPropertiesConsumer{/* values here */})
```

### AIGatewayMCPServerListenerAIGatewayMCPServerBaseACLPropertiesOauth

```go
aiGatewayMCPServerAIGatewayMCPServerListener := components.CreateAIGatewayMCPServerAIGatewayMCPServerListenerOauthAccessToken(components.AIGatewayMCPServerListenerAIGatewayMCPServerBaseACLPropertiesOauth{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayMCPServerAIGatewayMCPServerListener.Type {
	case components.AIGatewayMCPServerAIGatewayMCPServerListenerTypeConsumer:
		// aiGatewayMCPServerAIGatewayMCPServerListener.AIGatewayMCPServerListenerAIGatewayMCPServerBaseACLPropertiesConsumer is populated
	case components.AIGatewayMCPServerAIGatewayMCPServerListenerTypeOauthAccessToken:
		// aiGatewayMCPServerAIGatewayMCPServerListener.AIGatewayMCPServerListenerAIGatewayMCPServerBaseACLPropertiesOauth is populated
}
```
