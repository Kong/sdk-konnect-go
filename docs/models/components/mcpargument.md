# MCPArgument


## Supported Types

### MCPPositionalArgument

```go
mcpArgument := components.CreateMCPArgumentMCPPositionalArgument(components.MCPPositionalArgument{/* values here */})
```

### MCPNamedArgument

```go
mcpArgument := components.CreateMCPArgumentMCPNamedArgument(components.MCPNamedArgument{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch mcpArgument.Type {
	case components.MCPArgumentTypeMCPPositionalArgument:
		// mcpArgument.MCPPositionalArgument is populated
	case components.MCPArgumentTypeMCPNamedArgument:
		// mcpArgument.MCPNamedArgument is populated
}
```
