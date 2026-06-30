# CatalogResourceRefFilterParametersIntegrationInstanceID


## Supported Types

### UUIDFieldFilter

```go
catalogResourceRefFilterParametersIntegrationInstanceID := components.CreateCatalogResourceRefFilterParametersIntegrationInstanceIDUUIDFieldFilter(components.UUIDFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogResourceRefFilterParametersIntegrationInstanceID.Type {
	case components.CatalogResourceRefFilterParametersIntegrationInstanceIDTypeUUIDFieldFilter:
		// catalogResourceRefFilterParametersIntegrationInstanceID.UUIDFieldFilter is populated
}
```
