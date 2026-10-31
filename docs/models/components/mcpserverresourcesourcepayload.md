# McpServerResourceSourcePayload

The source of the MCP resource used in create and update requests.


## Supported Types

### MCPResourceRemoteMCPServerPayload

```go
mcpServerResourceSourcePayload := components.CreateMcpServerResourceSourcePayloadRemoteMcpServer(components.MCPResourceRemoteMCPServerPayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch mcpServerResourceSourcePayload.Type {
	case components.McpServerResourceSourcePayloadTypeRemoteMcpServer:
		// mcpServerResourceSourcePayload.MCPResourceRemoteMCPServerPayload is populated
}
```
