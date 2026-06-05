# CreateAPIVersionRequestProvider

Represent spec provider information used for fetching the API spec. For raw, provide the raw content in the `content` property instead of using this provider.



## Supported Types

### URLAPISpecProvider

```go
createAPIVersionRequestProvider := components.CreateCreateAPIVersionRequestProviderURLAPISpecProvider(components.URLAPISpecProvider{/* values here */})
```

### IntegrationAPISpecProviderPayload

```go
createAPIVersionRequestProvider := components.CreateCreateAPIVersionRequestProviderIntegrationAPISpecProviderPayload(components.IntegrationAPISpecProviderPayload{/* values here */})
```

### ResourceBoundIntegrationAPISpecProviderPayload

```go
createAPIVersionRequestProvider := components.CreateCreateAPIVersionRequestProviderResourceBoundIntegrationAPISpecProviderPayload(components.ResourceBoundIntegrationAPISpecProviderPayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createAPIVersionRequestProvider.Type {
	case components.CreateAPIVersionRequestProviderTypeURLAPISpecProvider:
		// createAPIVersionRequestProvider.URLAPISpecProvider is populated
	case components.CreateAPIVersionRequestProviderTypeIntegrationAPISpecProviderPayload:
		// createAPIVersionRequestProvider.IntegrationAPISpecProviderPayload is populated
	case components.CreateAPIVersionRequestProviderTypeResourceBoundIntegrationAPISpecProviderPayload:
		// createAPIVersionRequestProvider.ResourceBoundIntegrationAPISpecProviderPayload is populated
}
```
