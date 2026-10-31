# CapabilityControls

The capabilities of a mapped MCP resource that the MCP server does not expose.



## Supported Types

### APICapabilityControls

```go
capabilityControls := components.CreateCapabilityControlsAPI(components.APICapabilityControls{/* values here */})
```

### MCPCapabilityControls

```go
capabilityControls := components.CreateCapabilityControlsMcpServer(components.MCPCapabilityControls{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch capabilityControls.Type {
	case components.CapabilityControlsTypeAPI:
		// capabilityControls.APICapabilityControls is populated
	case components.CapabilityControlsTypeMcpServer:
		// capabilityControls.MCPCapabilityControls is populated
}
```
