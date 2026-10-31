# MCPRemoteTransport

Transport protocol configuration for remote context


## Supported Types

### MCPStreamableHTTPTransportMCPStreamableHTTPTransport

```go
mcpRemoteTransport := components.CreateMCPRemoteTransportMCPStreamableHTTPTransportMCPStreamableHTTPTransport(components.MCPStreamableHTTPTransportMCPStreamableHTTPTransport{/* values here */})
```

### MCPSseTransportMCPSseTransport

```go
mcpRemoteTransport := components.CreateMCPRemoteTransportMCPSseTransportMCPSseTransport(components.MCPSseTransportMCPSseTransport{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch mcpRemoteTransport.Type {
	case components.MCPRemoteTransportTypeMCPStreamableHTTPTransportMCPStreamableHTTPTransport:
		// mcpRemoteTransport.MCPStreamableHTTPTransportMCPStreamableHTTPTransport is populated
	case components.MCPRemoteTransportTypeMCPSseTransportMCPSseTransport:
		// mcpRemoteTransport.MCPSseTransportMCPSseTransport is populated
}
```
