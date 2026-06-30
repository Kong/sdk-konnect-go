# MCPResourceSourceDetail

The source of the MCP resource, indicating how it originated.


## Supported Types

### MCPResourceSourceAPICatalog

```go
mcpResourceSourceDetail := components.CreateMCPResourceSourceDetailAPICatalog(components.MCPResourceSourceAPICatalog{/* values here */})
```

### MCPResourceSourceRaw

```go
mcpResourceSourceDetail := components.CreateMCPResourceSourceDetailRaw(components.MCPResourceSourceRaw{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch mcpResourceSourceDetail.Type {
	case components.MCPResourceSourceDetailTypeAPICatalog:
		// mcpResourceSourceDetail.MCPResourceSourceAPICatalog is populated
	case components.MCPResourceSourceDetailTypeRaw:
		// mcpResourceSourceDetail.MCPResourceSourceRaw is populated
}
```
