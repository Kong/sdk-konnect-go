# AIGatewayMCPServerConversionListener


## Supported Types

### AIGatewayMCPServerBaseACLPropertiesConsumer

```go
aiGatewayMCPServerConversionListener := components.CreateAIGatewayMCPServerConversionListenerConsumer(components.AIGatewayMCPServerBaseACLPropertiesConsumer{/* values here */})
```

### AIGatewayMCPServerBaseACLPropertiesOauth

```go
aiGatewayMCPServerConversionListener := components.CreateAIGatewayMCPServerConversionListenerOauthAccessToken(components.AIGatewayMCPServerBaseACLPropertiesOauth{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayMCPServerConversionListener.Type {
	case components.AIGatewayMCPServerConversionListenerUnionTypeConsumer:
		// aiGatewayMCPServerConversionListener.AIGatewayMCPServerBaseACLPropertiesConsumer is populated
	case components.AIGatewayMCPServerConversionListenerUnionTypeOauthAccessToken:
		// aiGatewayMCPServerConversionListener.AIGatewayMCPServerBaseACLPropertiesOauth is populated
}
```
