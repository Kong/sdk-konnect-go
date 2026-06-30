# CreateDcrProviderRequest

Request body for creating a DCR provider. The provider_type cannot be updated after creation.


## Supported Types

### CreateDcrProviderRequestAuth0

```go
createDcrProviderRequest := components.CreateCreateDcrProviderRequestAuth0(components.CreateDcrProviderRequestAuth0{/* values here */})
```

### CreateDcrProviderRequestAzureAd

```go
createDcrProviderRequest := components.CreateCreateDcrProviderRequestAzureAd(components.CreateDcrProviderRequestAzureAd{/* values here */})
```

### CreateDcrProviderRequestCurity

```go
createDcrProviderRequest := components.CreateCreateDcrProviderRequestCurity(components.CreateDcrProviderRequestCurity{/* values here */})
```

### CreateDcrProviderRequestOkta

```go
createDcrProviderRequest := components.CreateCreateDcrProviderRequestOkta(components.CreateDcrProviderRequestOkta{/* values here */})
```

### CreateDcrProviderRequestHTTP

```go
createDcrProviderRequest := components.CreateCreateDcrProviderRequestHTTP(components.CreateDcrProviderRequestHTTP{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createDcrProviderRequest.Type {
	case components.CreateDcrProviderRequestTypeAuth0:
		// createDcrProviderRequest.CreateDcrProviderRequestAuth0 is populated
	case components.CreateDcrProviderRequestTypeAzureAd:
		// createDcrProviderRequest.CreateDcrProviderRequestAzureAd is populated
	case components.CreateDcrProviderRequestTypeCurity:
		// createDcrProviderRequest.CreateDcrProviderRequestCurity is populated
	case components.CreateDcrProviderRequestTypeOkta:
		// createDcrProviderRequest.CreateDcrProviderRequestOkta is populated
	case components.CreateDcrProviderRequestTypeHTTP:
		// createDcrProviderRequest.CreateDcrProviderRequestHTTP is populated
}
```
