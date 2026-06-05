# AddOnOwner

Owner for the add-on.


## Supported Types

### ControlPlaneAddOnOwner

```go
addOnOwner := components.CreateAddOnOwnerControlPlaneAddOnOwner(components.ControlPlaneAddOnOwner{/* values here */})
```

### ControlPlaneGroupAddOnOwner

```go
addOnOwner := components.CreateAddOnOwnerControlPlaneGroupAddOnOwner(components.ControlPlaneGroupAddOnOwner{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch addOnOwner.Type {
	case components.AddOnOwnerTypeControlPlaneAddOnOwner:
		// addOnOwner.ControlPlaneAddOnOwner is populated
	case components.AddOnOwnerTypeControlPlaneGroupAddOnOwner:
		// addOnOwner.ControlPlaneGroupAddOnOwner is populated
}
```
