# UpdateCatalogServiceAPISpecCreateAPISpecProvider

Represents the provider, or source, of an API specs.
Note that `provider` cannot be patched for API specs that are bound to a resource.
Trying to do so will result in a 400 response.



## Supported Types

### URLAPISpecProvider

```go
updateCatalogServiceAPISpecCreateAPISpecProvider := components.CreateUpdateCatalogServiceAPISpecCreateAPISpecProviderURLAPISpecProvider(components.URLAPISpecProvider{/* values here */})
```

### RawAPISpecProviderPayload

```go
updateCatalogServiceAPISpecCreateAPISpecProvider := components.CreateUpdateCatalogServiceAPISpecCreateAPISpecProviderRawAPISpecProviderPayload(components.RawAPISpecProviderPayload{/* values here */})
```

### IntegrationAPISpecProviderPayload

```go
updateCatalogServiceAPISpecCreateAPISpecProvider := components.CreateUpdateCatalogServiceAPISpecCreateAPISpecProviderIntegrationAPISpecProviderPayload(components.IntegrationAPISpecProviderPayload{/* values here */})
```

### ResourceBoundIntegrationAPISpecProviderPayload

```go
updateCatalogServiceAPISpecCreateAPISpecProvider := components.CreateUpdateCatalogServiceAPISpecCreateAPISpecProviderResourceBoundIntegrationAPISpecProviderPayload(components.ResourceBoundIntegrationAPISpecProviderPayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateCatalogServiceAPISpecCreateAPISpecProvider.Type {
	case components.UpdateCatalogServiceAPISpecCreateAPISpecProviderTypeURLAPISpecProvider:
		// updateCatalogServiceAPISpecCreateAPISpecProvider.URLAPISpecProvider is populated
	case components.UpdateCatalogServiceAPISpecCreateAPISpecProviderTypeRawAPISpecProviderPayload:
		// updateCatalogServiceAPISpecCreateAPISpecProvider.RawAPISpecProviderPayload is populated
	case components.UpdateCatalogServiceAPISpecCreateAPISpecProviderTypeIntegrationAPISpecProviderPayload:
		// updateCatalogServiceAPISpecCreateAPISpecProvider.IntegrationAPISpecProviderPayload is populated
	case components.UpdateCatalogServiceAPISpecCreateAPISpecProviderTypeResourceBoundIntegrationAPISpecProviderPayload:
		// updateCatalogServiceAPISpecCreateAPISpecProvider.ResourceBoundIntegrationAPISpecProviderPayload is populated
}
```
