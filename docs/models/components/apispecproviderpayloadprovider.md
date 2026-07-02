# APISpecProviderPayloadProvider


## Supported Types

### URLAPISpecProvider

```go
apiSpecProviderPayloadProvider := components.CreateAPISpecProviderPayloadProviderURLAPISpecProvider(components.URLAPISpecProvider{/* values here */})
```

### IntegrationAPISpecProviderPayload

```go
apiSpecProviderPayloadProvider := components.CreateAPISpecProviderPayloadProviderIntegrationAPISpecProviderPayload(components.IntegrationAPISpecProviderPayload{/* values here */})
```

### ResourceBoundIntegrationAPISpecProviderPayload

```go
apiSpecProviderPayloadProvider := components.CreateAPISpecProviderPayloadProviderResourceBoundIntegrationAPISpecProviderPayload(components.ResourceBoundIntegrationAPISpecProviderPayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch apiSpecProviderPayloadProvider.Type {
	case components.APISpecProviderPayloadProviderTypeURLAPISpecProvider:
		// apiSpecProviderPayloadProvider.URLAPISpecProvider is populated
	case components.APISpecProviderPayloadProviderTypeIntegrationAPISpecProviderPayload:
		// apiSpecProviderPayloadProvider.IntegrationAPISpecProviderPayload is populated
	case components.APISpecProviderPayloadProviderTypeResourceBoundIntegrationAPISpecProviderPayload:
		// apiSpecProviderPayloadProvider.ResourceBoundIntegrationAPISpecProviderPayload is populated
}
```
