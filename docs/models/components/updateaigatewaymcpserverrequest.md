# UpdateAIGatewayMCPServerRequest

**Pre-release Feature**
This feature is currently in beta and is subject to change.


## Supported Types

### AIGatewayMCPServerConversionOnly

```go
updateAIGatewayMCPServerRequest := components.CreateUpdateAIGatewayMCPServerRequestConversionOnly(components.AIGatewayMCPServerConversionOnly{/* values here */})
```

### AIGatewayMCPServerConversionListener

```go
updateAIGatewayMCPServerRequest := components.CreateUpdateAIGatewayMCPServerRequestConversionListener(components.AIGatewayMCPServerConversionListener{/* values here */})
```

### AIGatewayMCPServerListener

```go
updateAIGatewayMCPServerRequest := components.CreateUpdateAIGatewayMCPServerRequestListener(components.AIGatewayMCPServerListener{/* values here */})
```

### AIGatewayMCPServerPassthroughListener

```go
updateAIGatewayMCPServerRequest := components.CreateUpdateAIGatewayMCPServerRequestPassthroughListener(components.AIGatewayMCPServerPassthroughListener{/* values here */})
```

### AIGatewayMCPServerUpstreamServer

```go
updateAIGatewayMCPServerRequest := components.CreateUpdateAIGatewayMCPServerRequestUpstreamServer(components.AIGatewayMCPServerUpstreamServer{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateAIGatewayMCPServerRequest.Type {
	case components.UpdateAIGatewayMCPServerRequestTypeConversionOnly:
		// updateAIGatewayMCPServerRequest.AIGatewayMCPServerConversionOnly is populated
	case components.UpdateAIGatewayMCPServerRequestTypeConversionListener:
		// updateAIGatewayMCPServerRequest.AIGatewayMCPServerConversionListener is populated
	case components.UpdateAIGatewayMCPServerRequestTypeListener:
		// updateAIGatewayMCPServerRequest.AIGatewayMCPServerListener is populated
	case components.UpdateAIGatewayMCPServerRequestTypePassthroughListener:
		// updateAIGatewayMCPServerRequest.AIGatewayMCPServerPassthroughListener is populated
	case components.UpdateAIGatewayMCPServerRequestTypeUpstreamServer:
		// updateAIGatewayMCPServerRequest.AIGatewayMCPServerUpstreamServer is populated
}
```
