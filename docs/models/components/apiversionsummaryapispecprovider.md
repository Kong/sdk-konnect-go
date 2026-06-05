# APIVersionSummaryAPISpecProvider

Provenance of the spec for the current version. Present when a spec exists.


## Supported Types

### RawAPISpecProvider

```go
apiVersionSummaryAPISpecProvider := components.CreateAPIVersionSummaryAPISpecProviderRawAPISpecProvider(components.RawAPISpecProvider{/* values here */})
```

### URLAPISpecProvider

```go
apiVersionSummaryAPISpecProvider := components.CreateAPIVersionSummaryAPISpecProviderURLAPISpecProvider(components.URLAPISpecProvider{/* values here */})
```

### IntegrationAPISpecProviderPayload

```go
apiVersionSummaryAPISpecProvider := components.CreateAPIVersionSummaryAPISpecProviderIntegrationAPISpecProviderPayload(components.IntegrationAPISpecProviderPayload{/* values here */})
```

### ResourceBoundIntegrationAPISpecProviderPayload

```go
apiVersionSummaryAPISpecProvider := components.CreateAPIVersionSummaryAPISpecProviderResourceBoundIntegrationAPISpecProviderPayload(components.ResourceBoundIntegrationAPISpecProviderPayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch apiVersionSummaryAPISpecProvider.Type {
	case components.APIVersionSummaryAPISpecProviderTypeRawAPISpecProvider:
		// apiVersionSummaryAPISpecProvider.RawAPISpecProvider is populated
	case components.APIVersionSummaryAPISpecProviderTypeURLAPISpecProvider:
		// apiVersionSummaryAPISpecProvider.URLAPISpecProvider is populated
	case components.APIVersionSummaryAPISpecProviderTypeIntegrationAPISpecProviderPayload:
		// apiVersionSummaryAPISpecProvider.IntegrationAPISpecProviderPayload is populated
	case components.APIVersionSummaryAPISpecProviderTypeResourceBoundIntegrationAPISpecProviderPayload:
		// apiVersionSummaryAPISpecProvider.ResourceBoundIntegrationAPISpecProviderPayload is populated
}
```
