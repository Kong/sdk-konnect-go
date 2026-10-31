# ResourceIntegrationInstanceID


## Supported Types

### UUIDFieldFilter

```go
resourceIntegrationInstanceID := components.CreateResourceIntegrationInstanceIDUUIDFieldFilter(components.UUIDFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch resourceIntegrationInstanceID.Type {
	case components.ResourceIntegrationInstanceIDTypeUUIDFieldFilter:
		// resourceIntegrationInstanceID.UUIDFieldFilter is populated
}
```
