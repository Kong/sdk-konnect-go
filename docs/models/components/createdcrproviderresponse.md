# CreateDcrProviderResponse

A response containing the newly created DCR provider object.


## Supported Types

### DCRProviderAuth0

```go
createDcrProviderResponse := components.CreateCreateDcrProviderResponseAuth0(components.DCRProviderAuth0{/* values here */})
```

### DCRProviderAzureAD

```go
createDcrProviderResponse := components.CreateCreateDcrProviderResponseAzureAd(components.DCRProviderAzureAD{/* values here */})
```

### DCRProviderCurity

```go
createDcrProviderResponse := components.CreateCreateDcrProviderResponseCurity(components.DCRProviderCurity{/* values here */})
```

### DCRProviderOKTA

```go
createDcrProviderResponse := components.CreateCreateDcrProviderResponseOkta(components.DCRProviderOKTA{/* values here */})
```

### DCRProviderHTTP

```go
createDcrProviderResponse := components.CreateCreateDcrProviderResponseHTTP(components.DCRProviderHTTP{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createDcrProviderResponse.Type {
	case components.CreateDcrProviderResponseTypeAuth0:
		// createDcrProviderResponse.DCRProviderAuth0 is populated
	case components.CreateDcrProviderResponseTypeAzureAd:
		// createDcrProviderResponse.DCRProviderAzureAD is populated
	case components.CreateDcrProviderResponseTypeCurity:
		// createDcrProviderResponse.DCRProviderCurity is populated
	case components.CreateDcrProviderResponseTypeOkta:
		// createDcrProviderResponse.DCRProviderOKTA is populated
	case components.CreateDcrProviderResponseTypeHTTP:
		// createDcrProviderResponse.DCRProviderHTTP is populated
}
```
