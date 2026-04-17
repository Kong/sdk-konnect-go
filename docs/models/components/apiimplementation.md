# APIImplementation

An entity that implements an API


## Supported Types

### ServiceReference

```go
apiImplementation := components.CreateAPIImplementationServiceReference(components.ServiceReference{/* values here */})
```

### ControlPlaneReference

```go
apiImplementation := components.CreateAPIImplementationControlPlaneReference(components.ControlPlaneReference{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch apiImplementation.Type {
	case components.APIImplementationTypeServiceReference:
		// apiImplementation.ServiceReference is populated
	case components.APIImplementationTypeControlPlaneReference:
		// apiImplementation.ControlPlaneReference is populated
}
```
