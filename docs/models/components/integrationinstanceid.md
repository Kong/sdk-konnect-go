# IntegrationInstanceID


## Supported Types

### UUIDFieldFilter

```go
integrationInstanceID := components.CreateIntegrationInstanceIDUUIDFieldFilter(components.UUIDFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationInstanceID.Type {
	case components.IntegrationInstanceIDTypeUUIDFieldFilter:
		// integrationInstanceID.UUIDFieldFilter is populated
}
```
