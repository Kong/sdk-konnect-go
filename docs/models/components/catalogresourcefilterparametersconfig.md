# CatalogResourceFilterParametersConfig


## Supported Types

### CatalogResourceConfigFieldFilter

```go
catalogResourceFilterParametersConfig := components.CreateCatalogResourceFilterParametersConfigCatalogResourceConfigFieldFilter(components.CatalogResourceConfigFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogResourceFilterParametersConfig.Type {
	case components.CatalogResourceFilterParametersConfigTypeCatalogResourceConfigFieldFilter:
		// catalogResourceFilterParametersConfig.CatalogResourceConfigFieldFilter is populated
}
```
