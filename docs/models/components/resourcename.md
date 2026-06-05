# ResourceName


## Supported Types

### StringFieldFilter

```go
resourceName := components.CreateResourceNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch resourceName.Type {
	case components.ResourceNameTypeStringFieldFilter:
		// resourceName.StringFieldFilter is populated
}
```
