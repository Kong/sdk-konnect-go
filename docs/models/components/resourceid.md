# ResourceID


## Supported Types

### StringFieldFilterExact

```go
resourceID := components.CreateResourceIDStringFieldFilterExact(components.StringFieldFilterExact{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch resourceID.Type {
	case components.ResourceIDTypeStringFieldFilterExact:
		// resourceID.StringFieldFilterExact is populated
}
```
