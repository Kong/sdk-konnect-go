# ResourceConfig


## Supported Types

### CatalogResourceConfigFieldFilter

```go
resourceConfig := components.CreateResourceConfigCatalogResourceConfigFieldFilter(components.CatalogResourceConfigFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch resourceConfig.Type {
	case components.ResourceConfigTypeCatalogResourceConfigFieldFilter:
		// resourceConfig.CatalogResourceConfigFieldFilter is populated
}
```
