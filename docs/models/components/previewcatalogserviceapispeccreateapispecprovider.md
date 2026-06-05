# PreviewCatalogServiceAPISpecCreateAPISpecProvider

Represents the provider, or source, of an API spec


## Supported Types

### URLAPISpecProvider

```go
previewCatalogServiceAPISpecCreateAPISpecProvider := components.CreatePreviewCatalogServiceAPISpecCreateAPISpecProviderURLAPISpecProvider(components.URLAPISpecProvider{/* values here */})
```

### RawAPISpecProviderPayload

```go
previewCatalogServiceAPISpecCreateAPISpecProvider := components.CreatePreviewCatalogServiceAPISpecCreateAPISpecProviderRawAPISpecProviderPayload(components.RawAPISpecProviderPayload{/* values here */})
```

### IntegrationAPISpecProviderPayload

```go
previewCatalogServiceAPISpecCreateAPISpecProvider := components.CreatePreviewCatalogServiceAPISpecCreateAPISpecProviderIntegrationAPISpecProviderPayload(components.IntegrationAPISpecProviderPayload{/* values here */})
```

### ResourceBoundIntegrationAPISpecProviderPayload

```go
previewCatalogServiceAPISpecCreateAPISpecProvider := components.CreatePreviewCatalogServiceAPISpecCreateAPISpecProviderResourceBoundIntegrationAPISpecProviderPayload(components.ResourceBoundIntegrationAPISpecProviderPayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch previewCatalogServiceAPISpecCreateAPISpecProvider.Type {
	case components.PreviewCatalogServiceAPISpecCreateAPISpecProviderTypeURLAPISpecProvider:
		// previewCatalogServiceAPISpecCreateAPISpecProvider.URLAPISpecProvider is populated
	case components.PreviewCatalogServiceAPISpecCreateAPISpecProviderTypeRawAPISpecProviderPayload:
		// previewCatalogServiceAPISpecCreateAPISpecProvider.RawAPISpecProviderPayload is populated
	case components.PreviewCatalogServiceAPISpecCreateAPISpecProviderTypeIntegrationAPISpecProviderPayload:
		// previewCatalogServiceAPISpecCreateAPISpecProvider.IntegrationAPISpecProviderPayload is populated
	case components.PreviewCatalogServiceAPISpecCreateAPISpecProviderTypeResourceBoundIntegrationAPISpecProviderPayload:
		// previewCatalogServiceAPISpecCreateAPISpecProvider.ResourceBoundIntegrationAPISpecProviderPayload is populated
}
```
