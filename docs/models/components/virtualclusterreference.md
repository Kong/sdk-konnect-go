# VirtualClusterReference

A reference to a virtual cluster.


## Supported Types

### VirtualClusterReferenceByID

```go
virtualClusterReference := components.CreateVirtualClusterReferenceVirtualClusterReferenceByID(components.VirtualClusterReferenceByID{/* values here */})
```

### VirtualClusterReferenceByName

```go
virtualClusterReference := components.CreateVirtualClusterReferenceVirtualClusterReferenceByName(components.VirtualClusterReferenceByName{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch virtualClusterReference.Type {
	case components.VirtualClusterReferenceTypeVirtualClusterReferenceByID:
		// virtualClusterReference.VirtualClusterReferenceByID is populated
	case components.VirtualClusterReferenceTypeVirtualClusterReferenceByName:
		// virtualClusterReference.VirtualClusterReferenceByName is populated
}
```
