# MCPTransport

Transport protocol configuration for the package


## Supported Types

### MCPStdioTransport

```go
mcpTransport := components.CreateMCPTransportMCPStdioTransport(components.MCPStdioTransport{/* values here */})
```

### MCPStreamableHTTPTransport

```go
mcpTransport := components.CreateMCPTransportMCPStreamableHTTPTransport(components.MCPStreamableHTTPTransport{/* values here */})
```

### MCPSseTransport

```go
mcpTransport := components.CreateMCPTransportMCPSseTransport(components.MCPSseTransport{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch mcpTransport.Type {
	case components.MCPTransportTypeMCPStdioTransport:
		// mcpTransport.MCPStdioTransport is populated
	case components.MCPTransportTypeMCPStreamableHTTPTransport:
		// mcpTransport.MCPStreamableHTTPTransport is populated
	case components.MCPTransportTypeMCPSseTransport:
		// mcpTransport.MCPSseTransport is populated
}
```
