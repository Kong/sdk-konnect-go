# MCPServerSignal

A single signal envelope containing exactly one typed capability signal.


## Supported Types

### MCPServerSignalV1

```go
mcpServerSignal := components.CreateMCPServerSignalMcp(components.MCPServerSignalV1{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch mcpServerSignal.Type {
	case components.MCPServerSignalTypeMcp:
		// mcpServerSignal.MCPServerSignalV1 is populated
}
```
