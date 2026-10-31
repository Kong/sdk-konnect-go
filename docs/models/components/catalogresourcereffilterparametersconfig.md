# CatalogResourceRefFilterParametersConfig


## Supported Types

### CatalogResourceConfigFieldFilter

```go
catalogResourceRefFilterParametersConfig := components.CreateCatalogResourceRefFilterParametersConfigCatalogResourceConfigFieldFilter(components.CatalogResourceConfigFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogResourceRefFilterParametersConfig.Type {
	case components.CatalogResourceRefFilterParametersConfigTypeCatalogResourceConfigFieldFilter:
		// catalogResourceRefFilterParametersConfig.CatalogResourceConfigFieldFilter is populated
}
```
