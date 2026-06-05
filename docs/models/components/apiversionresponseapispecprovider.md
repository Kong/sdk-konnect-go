# APIVersionResponseAPISpecProvider

The provider used to supply the spec content. In write requests, provide
this instead of or alongside content. If omitted on writes, the server
infers type raw from spec.content.



## Supported Types

### RawAPISpecProvider

```go
apiVersionResponseAPISpecProvider := components.CreateAPIVersionResponseAPISpecProviderRawAPISpecProvider(components.RawAPISpecProvider{/* values here */})
```

### URLAPISpecProvider

```go
apiVersionResponseAPISpecProvider := components.CreateAPIVersionResponseAPISpecProviderURLAPISpecProvider(components.URLAPISpecProvider{/* values here */})
```

### IntegrationAPISpecProviderPayload

```go
apiVersionResponseAPISpecProvider := components.CreateAPIVersionResponseAPISpecProviderIntegrationAPISpecProviderPayload(components.IntegrationAPISpecProviderPayload{/* values here */})
```

### ResourceBoundIntegrationAPISpecProviderPayload

```go
apiVersionResponseAPISpecProvider := components.CreateAPIVersionResponseAPISpecProviderResourceBoundIntegrationAPISpecProviderPayload(components.ResourceBoundIntegrationAPISpecProviderPayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch apiVersionResponseAPISpecProvider.Type {
	case components.APIVersionResponseAPISpecProviderTypeRawAPISpecProvider:
		// apiVersionResponseAPISpecProvider.RawAPISpecProvider is populated
	case components.APIVersionResponseAPISpecProviderTypeURLAPISpecProvider:
		// apiVersionResponseAPISpecProvider.URLAPISpecProvider is populated
	case components.APIVersionResponseAPISpecProviderTypeIntegrationAPISpecProviderPayload:
		// apiVersionResponseAPISpecProvider.IntegrationAPISpecProviderPayload is populated
	case components.APIVersionResponseAPISpecProviderTypeResourceBoundIntegrationAPISpecProviderPayload:
		// apiVersionResponseAPISpecProvider.ResourceBoundIntegrationAPISpecProviderPayload is populated
}
```
