# AIGatewayMCPServerListener


## Supported Types

### AIGatewayMCPServerBaseACLPropertiesConsumerAIGatewayMCPServerBaseACLPropertiesConsumer

```go
aiGatewayMCPServerListener := components.CreateAIGatewayMCPServerListenerConsumer(components.AIGatewayMCPServerBaseACLPropertiesConsumerAIGatewayMCPServerBaseACLPropertiesConsumer{/* values here */})
```

### AIGatewayMCPServerBaseACLPropertiesOauthAIGatewayMCPServerBaseACLPropertiesOauth

```go
aiGatewayMCPServerListener := components.CreateAIGatewayMCPServerListenerOauthAccessToken(components.AIGatewayMCPServerBaseACLPropertiesOauthAIGatewayMCPServerBaseACLPropertiesOauth{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayMCPServerListener.Type {
	case components.AIGatewayMCPServerListenerUnionTypeConsumer:
		// aiGatewayMCPServerListener.AIGatewayMCPServerBaseACLPropertiesConsumerAIGatewayMCPServerBaseACLPropertiesConsumer is populated
	case components.AIGatewayMCPServerListenerUnionTypeOauthAccessToken:
		// aiGatewayMCPServerListener.AIGatewayMCPServerBaseACLPropertiesOauthAIGatewayMCPServerBaseACLPropertiesOauth is populated
}
```
