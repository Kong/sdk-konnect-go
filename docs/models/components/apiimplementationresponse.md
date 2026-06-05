# APIImplementationResponse

An entity that implements an API


## Supported Types

### APIImplementationResponseServiceReference

```go
apiImplementationResponse := components.CreateAPIImplementationResponseAPIImplementationResponseServiceReference(components.APIImplementationResponseServiceReference{/* values here */})
```

### APIImplementationResponseControlPlaneReference

```go
apiImplementationResponse := components.CreateAPIImplementationResponseAPIImplementationResponseControlPlaneReference(components.APIImplementationResponseControlPlaneReference{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch apiImplementationResponse.Type {
	case components.APIImplementationResponseTypeAPIImplementationResponseServiceReference:
		// apiImplementationResponse.APIImplementationResponseServiceReference is populated
	case components.APIImplementationResponseTypeAPIImplementationResponseControlPlaneReference:
		// apiImplementationResponse.APIImplementationResponseControlPlaneReference is populated
}
```
