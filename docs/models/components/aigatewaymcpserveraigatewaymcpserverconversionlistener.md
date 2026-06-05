# AIGatewayMCPServerAIGatewayMCPServerConversionListener


## Supported Types

### AIGatewayMCPServerConversionListenerAIGatewayMCPServerBaseACLPropertiesConsumer

```go
aiGatewayMCPServerAIGatewayMCPServerConversionListener := components.CreateAIGatewayMCPServerAIGatewayMCPServerConversionListenerConsumer(components.AIGatewayMCPServerConversionListenerAIGatewayMCPServerBaseACLPropertiesConsumer{/* values here */})
```

### AIGatewayMCPServerConversionListenerAIGatewayMCPServerBaseACLPropertiesOauth

```go
aiGatewayMCPServerAIGatewayMCPServerConversionListener := components.CreateAIGatewayMCPServerAIGatewayMCPServerConversionListenerOauthAccessToken(components.AIGatewayMCPServerConversionListenerAIGatewayMCPServerBaseACLPropertiesOauth{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayMCPServerAIGatewayMCPServerConversionListener.Type {
	case components.AIGatewayMCPServerAIGatewayMCPServerConversionListenerTypeConsumer:
		// aiGatewayMCPServerAIGatewayMCPServerConversionListener.AIGatewayMCPServerConversionListenerAIGatewayMCPServerBaseACLPropertiesConsumer is populated
	case components.AIGatewayMCPServerAIGatewayMCPServerConversionListenerTypeOauthAccessToken:
		// aiGatewayMCPServerAIGatewayMCPServerConversionListener.AIGatewayMCPServerConversionListenerAIGatewayMCPServerBaseACLPropertiesOauth is populated
}
```
