# Config


## Supported Types

### CatalogResourceConfigFieldFilter

```go
config := components.CreateConfigCatalogResourceConfigFieldFilter(components.CatalogResourceConfigFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch config.Type {
	case components.ConfigTypeCatalogResourceConfigFieldFilter:
		// config.CatalogResourceConfigFieldFilter is populated
}
```
