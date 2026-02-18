# ResourceIntegrationDisplayName


## Supported Types

### StringFieldFilter

```go
resourceIntegrationDisplayName := components.CreateResourceIntegrationDisplayNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch resourceIntegrationDisplayName.Type {
	case components.ResourceIntegrationDisplayNameTypeStringFieldFilter:
		// resourceIntegrationDisplayName.StringFieldFilter is populated
}
```
