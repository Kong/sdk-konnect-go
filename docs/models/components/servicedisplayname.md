# ServiceDisplayName


## Supported Types

### StringFieldFilter

```go
serviceDisplayName := components.CreateServiceDisplayNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch serviceDisplayName.Type {
	case components.ServiceDisplayNameTypeStringFieldFilter:
		// serviceDisplayName.StringFieldFilter is populated
}
```
