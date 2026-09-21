# Targets

The data plane nodes the log level change applies to. Use `all` to target every node connected to the control plane, or `node_ids` to target specific nodes.


## Supported Types

### DataPlaneNodeLogLevelTargetAll

```go
targets := components.CreateTargetsDataPlaneNodeLogLevelTargetAll(components.DataPlaneNodeLogLevelTargetAll{/* values here */})
```

### DataPlaneNodeLogLevelTargetNodeIds

```go
targets := components.CreateTargetsDataPlaneNodeLogLevelTargetNodeIds(components.DataPlaneNodeLogLevelTargetNodeIds{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch targets.Type {
	case components.TargetsTypeDataPlaneNodeLogLevelTargetAll:
		// targets.DataPlaneNodeLogLevelTargetAll is populated
	case components.TargetsTypeDataPlaneNodeLogLevelTargetNodeIds:
		// targets.DataPlaneNodeLogLevelTargetNodeIds is populated
}
```
