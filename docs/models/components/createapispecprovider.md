# CreateAPISpecProvider

Represents the provider, or source, of an API spec


## Supported Types

### URLAPISpecProvider

```go
createAPISpecProvider := components.CreateCreateAPISpecProviderURLAPISpecProvider(components.URLAPISpecProvider{/* values here */})
```

### RawAPISpecProviderPayload

```go
createAPISpecProvider := components.CreateCreateAPISpecProviderRawAPISpecProviderPayload(components.RawAPISpecProviderPayload{/* values here */})
```

### IntegrationAPISpecProviderPayload

```go
createAPISpecProvider := components.CreateCreateAPISpecProviderIntegrationAPISpecProviderPayload(components.IntegrationAPISpecProviderPayload{/* values here */})
```

### ResourceBoundIntegrationAPISpecProviderPayload

```go
createAPISpecProvider := components.CreateCreateAPISpecProviderResourceBoundIntegrationAPISpecProviderPayload(components.ResourceBoundIntegrationAPISpecProviderPayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createAPISpecProvider.Type {
	case components.CreateAPISpecProviderTypeURLAPISpecProvider:
		// createAPISpecProvider.URLAPISpecProvider is populated
	case components.CreateAPISpecProviderTypeRawAPISpecProviderPayload:
		// createAPISpecProvider.RawAPISpecProviderPayload is populated
	case components.CreateAPISpecProviderTypeIntegrationAPISpecProviderPayload:
		// createAPISpecProvider.IntegrationAPISpecProviderPayload is populated
	case components.CreateAPISpecProviderTypeResourceBoundIntegrationAPISpecProviderPayload:
		// createAPISpecProvider.ResourceBoundIntegrationAPISpecProviderPayload is populated
}
```
