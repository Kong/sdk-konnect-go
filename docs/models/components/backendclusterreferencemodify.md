# BackendClusterReferenceModify

The backend cluster associated with the virtual cluster.

Either `id` or `name` must be provided. Following changes to the backend cluster name won't affect the
reference, as the system will create the entities relationship by `id`.



## Supported Types

### BackendClusterReferenceByID

```go
backendClusterReferenceModify := components.CreateBackendClusterReferenceModifyBackendClusterReferenceByID(components.BackendClusterReferenceByID{/* values here */})
```

### BackendClusterReferenceByName

```go
backendClusterReferenceModify := components.CreateBackendClusterReferenceModifyBackendClusterReferenceByName(components.BackendClusterReferenceByName{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch backendClusterReferenceModify.Type {
	case components.BackendClusterReferenceModifyTypeBackendClusterReferenceByID:
		// backendClusterReferenceModify.BackendClusterReferenceByID is populated
	case components.BackendClusterReferenceModifyTypeBackendClusterReferenceByName:
		// backendClusterReferenceModify.BackendClusterReferenceByName is populated
}
```
