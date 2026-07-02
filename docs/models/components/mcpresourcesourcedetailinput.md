# MCPResourceSourceDetailInput

The source of the MCP resource used in create and update requests.


## Supported Types

### MCPResourceSourceAPICatalog

```go
mcpResourceSourceDetailInput := components.CreateMCPResourceSourceDetailInputAPICatalog(components.MCPResourceSourceAPICatalog{/* values here */})
```

### MCPResourceSourceRawInput

```go
mcpResourceSourceDetailInput := components.CreateMCPResourceSourceDetailInputRaw(components.MCPResourceSourceRawInput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch mcpResourceSourceDetailInput.Type {
	case components.MCPResourceSourceDetailInputTypeAPICatalog:
		// mcpResourceSourceDetailInput.MCPResourceSourceAPICatalog is populated
	case components.MCPResourceSourceDetailInputTypeRaw:
		// mcpResourceSourceDetailInput.MCPResourceSourceRawInput is populated
}
```
