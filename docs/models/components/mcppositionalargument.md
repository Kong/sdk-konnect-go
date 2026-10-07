# MCPPositionalArgument

A positional input is a value inserted verbatim into the command line.


## Supported Types

### MCPPositionalArgument1

```go
mcpPositionalArgument := components.CreateMCPPositionalArgumentMCPPositionalArgument1(components.MCPPositionalArgument1{/* values here */})
```

### MCPPositionalArgument2

```go
mcpPositionalArgument := components.CreateMCPPositionalArgumentMCPPositionalArgument2(components.MCPPositionalArgument2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch mcpPositionalArgument.Type {
	case components.MCPPositionalArgumentTypeMCPPositionalArgument1:
		// mcpPositionalArgument.MCPPositionalArgument1 is populated
	case components.MCPPositionalArgumentTypeMCPPositionalArgument2:
		// mcpPositionalArgument.MCPPositionalArgument2 is populated
}
```
