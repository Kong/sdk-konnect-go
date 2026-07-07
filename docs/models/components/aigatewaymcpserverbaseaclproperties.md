# AIGatewayMCPServerBaseACLProperties


## Supported Types

### AIGatewayMCPServerBaseACLPropertiesConsumer

```go
aiGatewayMCPServerBaseACLProperties := components.CreateAIGatewayMCPServerBaseACLPropertiesConsumer(components.AIGatewayMCPServerBaseACLPropertiesConsumer{/* values here */})
```

### AIGatewayMCPServerBaseACLPropertiesOauth

```go
aiGatewayMCPServerBaseACLProperties := components.CreateAIGatewayMCPServerBaseACLPropertiesOauthAccessToken(components.AIGatewayMCPServerBaseACLPropertiesOauth{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayMCPServerBaseACLProperties.Type {
	case components.AIGatewayMCPServerBaseACLPropertiesTypeConsumer:
		// aiGatewayMCPServerBaseACLProperties.AIGatewayMCPServerBaseACLPropertiesConsumer is populated
	case components.AIGatewayMCPServerBaseACLPropertiesTypeOauthAccessToken:
		// aiGatewayMCPServerBaseACLProperties.AIGatewayMCPServerBaseACLPropertiesOauth is populated
}
```
