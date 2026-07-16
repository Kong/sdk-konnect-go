# CreateAIGatewayMCPServerRequest

**Pre-release Feature**
This feature is currently in beta and is subject to change.


## Supported Types

### AIGatewayMCPServerConversionOnly

```go
createAIGatewayMCPServerRequest := components.CreateCreateAIGatewayMCPServerRequestConversionOnly(components.AIGatewayMCPServerConversionOnly{/* values here */})
```

### AIGatewayMCPServerConversionListener

```go
createAIGatewayMCPServerRequest := components.CreateCreateAIGatewayMCPServerRequestConversionListener(components.AIGatewayMCPServerConversionListener{/* values here */})
```

### AIGatewayMCPServerListener

```go
createAIGatewayMCPServerRequest := components.CreateCreateAIGatewayMCPServerRequestListener(components.AIGatewayMCPServerListener{/* values here */})
```

### AIGatewayMCPServerPassthroughListener

```go
createAIGatewayMCPServerRequest := components.CreateCreateAIGatewayMCPServerRequestPassthroughListener(components.AIGatewayMCPServerPassthroughListener{/* values here */})
```

### AIGatewayMCPServerUpstreamServer

```go
createAIGatewayMCPServerRequest := components.CreateCreateAIGatewayMCPServerRequestUpstreamServer(components.AIGatewayMCPServerUpstreamServer{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createAIGatewayMCPServerRequest.Type {
	case components.CreateAIGatewayMCPServerRequestTypeConversionOnly:
		// createAIGatewayMCPServerRequest.AIGatewayMCPServerConversionOnly is populated
	case components.CreateAIGatewayMCPServerRequestTypeConversionListener:
		// createAIGatewayMCPServerRequest.AIGatewayMCPServerConversionListener is populated
	case components.CreateAIGatewayMCPServerRequestTypeListener:
		// createAIGatewayMCPServerRequest.AIGatewayMCPServerListener is populated
	case components.CreateAIGatewayMCPServerRequestTypePassthroughListener:
		// createAIGatewayMCPServerRequest.AIGatewayMCPServerPassthroughListener is populated
	case components.CreateAIGatewayMCPServerRequestTypeUpstreamServer:
		// createAIGatewayMCPServerRequest.AIGatewayMCPServerUpstreamServer is populated
}
```
