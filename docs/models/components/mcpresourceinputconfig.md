# MCPResourceInputConfig


## Supported Types

### MCPResourceAPIReferenceConfig

```go
mcpResourceInputConfig := components.CreateMCPResourceInputConfigAPIReference(components.MCPResourceAPIReferenceConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch mcpResourceInputConfig.Type {
	case components.MCPResourceInputConfigTypeAPIReference:
		// mcpResourceInputConfig.MCPResourceAPIReferenceConfig is populated
}
```
