# AddOnOwner

Owner for the add-on.


## Supported Types

### ControlPlane

```go
addOnOwner := components.CreateAddOnOwnerControlPlane(components.ControlPlane{/* values here */})
```

### ControlPlaneGroup

```go
addOnOwner := components.CreateAddOnOwnerControlPlaneGroup(components.ControlPlaneGroup{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch addOnOwner.Type {
	case components.AddOnOwnerTypeControlPlane:
		// addOnOwner.ControlPlane is populated
	case components.AddOnOwnerTypeControlPlaneGroup:
		// addOnOwner.ControlPlaneGroup is populated
}
```
