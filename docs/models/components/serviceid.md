# ServiceID


## Supported Types

### StringFieldFilterExact

```go
serviceID := components.CreateServiceIDStringFieldFilterExact(components.StringFieldFilterExact{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch serviceID.Type {
	case components.ServiceIDTypeStringFieldFilterExact:
		// serviceID.StringFieldFilterExact is populated
}
```
