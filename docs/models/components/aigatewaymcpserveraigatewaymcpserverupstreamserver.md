# AIGatewayMCPServerAIGatewayMCPServerUpstreamServer


## Supported Types

### AIGatewayMCPServerUpstreamServerAIGatewayMCPServerBaseACLPropertiesConsumer

```go
aiGatewayMCPServerAIGatewayMCPServerUpstreamServer := components.CreateAIGatewayMCPServerAIGatewayMCPServerUpstreamServerConsumer(components.AIGatewayMCPServerUpstreamServerAIGatewayMCPServerBaseACLPropertiesConsumer{/* values here */})
```

### AIGatewayMCPServerUpstreamServerAIGatewayMCPServerBaseACLPropertiesOauth

```go
aiGatewayMCPServerAIGatewayMCPServerUpstreamServer := components.CreateAIGatewayMCPServerAIGatewayMCPServerUpstreamServerOauthAccessToken(components.AIGatewayMCPServerUpstreamServerAIGatewayMCPServerBaseACLPropertiesOauth{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayMCPServerAIGatewayMCPServerUpstreamServer.Type {
	case components.AIGatewayMCPServerAIGatewayMCPServerUpstreamServerTypeConsumer:
		// aiGatewayMCPServerAIGatewayMCPServerUpstreamServer.AIGatewayMCPServerUpstreamServerAIGatewayMCPServerBaseACLPropertiesConsumer is populated
	case components.AIGatewayMCPServerAIGatewayMCPServerUpstreamServerTypeOauthAccessToken:
		// aiGatewayMCPServerAIGatewayMCPServerUpstreamServer.AIGatewayMCPServerUpstreamServerAIGatewayMCPServerBaseACLPropertiesOauth is populated
}
```
