# Destination


## Supported Types

### HTTPDestination

```go
destination := components.CreateDestinationHTTP(components.HTTPDestination{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch destination.Type {
	case components.DestinationTypeHTTP:
		// destination.HTTPDestination is populated
}
```
