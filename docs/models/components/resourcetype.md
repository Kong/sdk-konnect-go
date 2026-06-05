# ResourceType


## Supported Types

### StringFieldFilter

```go
resourceType := components.CreateResourceTypeStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch resourceType.Type {
	case components.ResourceTypeTypeStringFieldFilter:
		// resourceType.StringFieldFilter is populated
}
```
