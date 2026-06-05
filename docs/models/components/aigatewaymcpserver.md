# AIGatewayMCPServer


## Supported Types

### AIGatewayMCPServerAIGatewayMCPServerConversionOnly

```go
aiGatewayMCPServer := components.CreateAIGatewayMCPServerConversionOnly(components.AIGatewayMCPServerAIGatewayMCPServerConversionOnly{/* values here */})
```

### AIGatewayMCPServerAIGatewayMCPServerConversionListener

```go
aiGatewayMCPServer := components.CreateAIGatewayMCPServerConversionListener(components.AIGatewayMCPServerAIGatewayMCPServerConversionListener{/* values here */})
```

### AIGatewayMCPServerAIGatewayMCPServerListener

```go
aiGatewayMCPServer := components.CreateAIGatewayMCPServerListener(components.AIGatewayMCPServerAIGatewayMCPServerListener{/* values here */})
```

### AIGatewayMCPServerAIGatewayMCPServerPassthroughListener

```go
aiGatewayMCPServer := components.CreateAIGatewayMCPServerPassthroughListener(components.AIGatewayMCPServerAIGatewayMCPServerPassthroughListener{/* values here */})
```

### AIGatewayMCPServerAIGatewayMCPServerUpstreamServer

```go
aiGatewayMCPServer := components.CreateAIGatewayMCPServerUpstreamServer(components.AIGatewayMCPServerAIGatewayMCPServerUpstreamServer{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayMCPServer.Type {
	case components.AIGatewayMCPServerTypeConversionOnly:
		// aiGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerConversionOnly is populated
	case components.AIGatewayMCPServerTypeConversionListener:
		// aiGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerConversionListener is populated
	case components.AIGatewayMCPServerTypeListener:
		// aiGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerListener is populated
	case components.AIGatewayMCPServerTypePassthroughListener:
		// aiGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerPassthroughListener is populated
	case components.AIGatewayMCPServerTypeUpstreamServer:
		// aiGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerUpstreamServer is populated
}
```
