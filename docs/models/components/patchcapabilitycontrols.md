# PatchCapabilityControls

The capability controls to merge into the existing configuration for a mapped context source. Omitted deny keys preserve their current value, an empty array clears a capability, and a non-empty array replaces.



## Supported Types

### PatchAPICapabilityControls

```go
patchCapabilityControls := components.CreatePatchCapabilityControlsAPI(components.PatchAPICapabilityControls{/* values here */})
```

### PatchMCPCapabilityControls

```go
patchCapabilityControls := components.CreatePatchCapabilityControlsMcpServer(components.PatchMCPCapabilityControls{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch patchCapabilityControls.Type {
	case components.PatchCapabilityControlsTypeAPI:
		// patchCapabilityControls.PatchAPICapabilityControls is populated
	case components.PatchCapabilityControlsTypeMcpServer:
		// patchCapabilityControls.PatchMCPCapabilityControls is populated
}
```
