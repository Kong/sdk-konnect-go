# APIVersionRequestProvider

Represent spec provider information used for fetching the API spec. For raw, provide the raw content in the `content` property instead of using this provider.



## Supported Types

### URLAPISpecProvider

```go
apiVersionRequestProvider := components.CreateAPIVersionRequestProviderURLAPISpecProvider(components.URLAPISpecProvider{/* values here */})
```

### IntegrationAPISpecProviderPayload

```go
apiVersionRequestProvider := components.CreateAPIVersionRequestProviderIntegrationAPISpecProviderPayload(components.IntegrationAPISpecProviderPayload{/* values here */})
```

### ResourceBoundIntegrationAPISpecProviderPayload

```go
apiVersionRequestProvider := components.CreateAPIVersionRequestProviderResourceBoundIntegrationAPISpecProviderPayload(components.ResourceBoundIntegrationAPISpecProviderPayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch apiVersionRequestProvider.Type {
	case components.APIVersionRequestProviderTypeURLAPISpecProvider:
		// apiVersionRequestProvider.URLAPISpecProvider is populated
	case components.APIVersionRequestProviderTypeIntegrationAPISpecProviderPayload:
		// apiVersionRequestProvider.IntegrationAPISpecProviderPayload is populated
	case components.APIVersionRequestProviderTypeResourceBoundIntegrationAPISpecProviderPayload:
		// apiVersionRequestProvider.ResourceBoundIntegrationAPISpecProviderPayload is populated
}
```
