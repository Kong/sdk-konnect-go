# CatalogResourceFilterParametersIntegrationInstanceID


## Supported Types

### UUIDFieldFilter

```go
catalogResourceFilterParametersIntegrationInstanceID := components.CreateCatalogResourceFilterParametersIntegrationInstanceIDUUIDFieldFilter(components.UUIDFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogResourceFilterParametersIntegrationInstanceID.Type {
	case components.CatalogResourceFilterParametersIntegrationInstanceIDTypeUUIDFieldFilter:
		// catalogResourceFilterParametersIntegrationInstanceID.UUIDFieldFilter is populated
}
```
