# AIGatewayMCPServer

**Pre-release Feature**
This feature is currently in beta and is subject to change.


## Supported Types

### AIGatewayMCPServerConversionOnlyResponse

```go
aiGatewayMCPServer := components.CreateAIGatewayMCPServerConversionOnly(components.AIGatewayMCPServerConversionOnlyResponse{/* values here */})
```

### AIGatewayMCPServerConversionListenerResponse

```go
aiGatewayMCPServer := components.CreateAIGatewayMCPServerConversionListener(components.AIGatewayMCPServerConversionListenerResponse{/* values here */})
```

### AIGatewayMCPServerListenerResponse

```go
aiGatewayMCPServer := components.CreateAIGatewayMCPServerListener(components.AIGatewayMCPServerListenerResponse{/* values here */})
```

### AIGatewayMCPServerPassthroughListenerResponse

```go
aiGatewayMCPServer := components.CreateAIGatewayMCPServerPassthroughListener(components.AIGatewayMCPServerPassthroughListenerResponse{/* values here */})
```

### AIGatewayMCPServerUpstreamServerResponse

```go
aiGatewayMCPServer := components.CreateAIGatewayMCPServerUpstreamServer(components.AIGatewayMCPServerUpstreamServerResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayMCPServer.Type {
	case components.AIGatewayMCPServerTypeConversionOnly:
		// aiGatewayMCPServer.AIGatewayMCPServerConversionOnlyResponse is populated
	case components.AIGatewayMCPServerTypeConversionListener:
		// aiGatewayMCPServer.AIGatewayMCPServerConversionListenerResponse is populated
	case components.AIGatewayMCPServerTypeListener:
		// aiGatewayMCPServer.AIGatewayMCPServerListenerResponse is populated
	case components.AIGatewayMCPServerTypePassthroughListener:
		// aiGatewayMCPServer.AIGatewayMCPServerPassthroughListenerResponse is populated
	case components.AIGatewayMCPServerTypeUpstreamServer:
		// aiGatewayMCPServer.AIGatewayMCPServerUpstreamServerResponse is populated
}
```
