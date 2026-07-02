# CreateCatalogIntegration


## Supported Types

### CreatePrivateCatalogIntegration

```go
createCatalogIntegration := components.CreateCreateCatalogIntegrationCreatePrivateCatalogIntegration(components.CreatePrivateCatalogIntegration{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createCatalogIntegration.Type {
	case components.CreateCatalogIntegrationTypeCreatePrivateCatalogIntegration:
		// createCatalogIntegration.CreatePrivateCatalogIntegration is populated
}
```
