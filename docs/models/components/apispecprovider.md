# APISpecProvider

Represents the provider, or source, of API spec contents.


## Supported Types

### URLAPISpecProvider

```go
apiSpecProvider := components.CreateAPISpecProviderURLAPISpecProvider(components.URLAPISpecProvider{/* values here */})
```

### RawAPISpecProvider

```go
apiSpecProvider := components.CreateAPISpecProviderRawAPISpecProvider(components.RawAPISpecProvider{/* values here */})
```

### IntegrationAPISpecProvider

```go
apiSpecProvider := components.CreateAPISpecProviderIntegrationAPISpecProvider(components.IntegrationAPISpecProvider{/* values here */})
```

### ResourceBoundIntegrationAPISpecProvider

```go
apiSpecProvider := components.CreateAPISpecProviderResourceBoundIntegrationAPISpecProvider(components.ResourceBoundIntegrationAPISpecProvider{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch apiSpecProvider.Type {
	case components.APISpecProviderTypeURLAPISpecProvider:
		// apiSpecProvider.URLAPISpecProvider is populated
	case components.APISpecProviderTypeRawAPISpecProvider:
		// apiSpecProvider.RawAPISpecProvider is populated
	case components.APISpecProviderTypeIntegrationAPISpecProvider:
		// apiSpecProvider.IntegrationAPISpecProvider is populated
	case components.APISpecProviderTypeResourceBoundIntegrationAPISpecProvider:
		// apiSpecProvider.ResourceBoundIntegrationAPISpecProvider is populated
}
```
