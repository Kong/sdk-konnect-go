# MCPResourceInfo


## Supported Types

### APIResource

```go
mcpResourceInfo := components.CreateMCPResourceInfoAPI(components.APIResource{/* values here */})
```

### McpServerResource

```go
mcpResourceInfo := components.CreateMCPResourceInfoMcpServer(components.McpServerResource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch mcpResourceInfo.Type {
	case components.MCPResourceInfoTypeAPI:
		// mcpResourceInfo.APIResource is populated
	case components.MCPResourceInfoTypeMcpServer:
		// mcpResourceInfo.McpServerResource is populated
}
```
