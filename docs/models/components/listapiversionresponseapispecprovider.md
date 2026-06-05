# ListAPIVersionResponseAPISpecProvider

Provenance of the spec for the current version. Present when a spec exists.


## Supported Types

### RawAPISpecProvider

```go
listAPIVersionResponseAPISpecProvider := components.CreateListAPIVersionResponseAPISpecProviderRawAPISpecProvider(components.RawAPISpecProvider{/* values here */})
```

### URLAPISpecProvider

```go
listAPIVersionResponseAPISpecProvider := components.CreateListAPIVersionResponseAPISpecProviderURLAPISpecProvider(components.URLAPISpecProvider{/* values here */})
```

### IntegrationAPISpecProviderPayload

```go
listAPIVersionResponseAPISpecProvider := components.CreateListAPIVersionResponseAPISpecProviderIntegrationAPISpecProviderPayload(components.IntegrationAPISpecProviderPayload{/* values here */})
```

### ResourceBoundIntegrationAPISpecProviderPayload

```go
listAPIVersionResponseAPISpecProvider := components.CreateListAPIVersionResponseAPISpecProviderResourceBoundIntegrationAPISpecProviderPayload(components.ResourceBoundIntegrationAPISpecProviderPayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listAPIVersionResponseAPISpecProvider.Type {
	case components.ListAPIVersionResponseAPISpecProviderTypeRawAPISpecProvider:
		// listAPIVersionResponseAPISpecProvider.RawAPISpecProvider is populated
	case components.ListAPIVersionResponseAPISpecProviderTypeURLAPISpecProvider:
		// listAPIVersionResponseAPISpecProvider.URLAPISpecProvider is populated
	case components.ListAPIVersionResponseAPISpecProviderTypeIntegrationAPISpecProviderPayload:
		// listAPIVersionResponseAPISpecProvider.IntegrationAPISpecProviderPayload is populated
	case components.ListAPIVersionResponseAPISpecProviderTypeResourceBoundIntegrationAPISpecProviderPayload:
		// listAPIVersionResponseAPISpecProvider.ResourceBoundIntegrationAPISpecProviderPayload is populated
}
```
