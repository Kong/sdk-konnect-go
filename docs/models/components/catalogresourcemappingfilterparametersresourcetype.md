# CatalogResourceMappingFilterParametersResourceType


## Supported Types

### StringFieldFilter

```go
catalogResourceMappingFilterParametersResourceType := components.CreateCatalogResourceMappingFilterParametersResourceTypeStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogResourceMappingFilterParametersResourceType.Type {
	case components.CatalogResourceMappingFilterParametersResourceTypeTypeStringFieldFilter:
		// catalogResourceMappingFilterParametersResourceType.StringFieldFilter is populated
}
```
