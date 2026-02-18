# CreateDcrProviderResponse

A response containing the newly created DCR provider object.


## Supported Types

### DCRProviderAuth0

```go
createDcrProviderResponse := components.CreateCreateDcrProviderResponseDcrProviderAuth0(components.DCRProviderAuth0{/* values here */})
```

### DCRProviderAzureAD

```go
createDcrProviderResponse := components.CreateCreateDcrProviderResponseDcrProviderAzureAd(components.DCRProviderAzureAD{/* values here */})
```

### DCRProviderCurity

```go
createDcrProviderResponse := components.CreateCreateDcrProviderResponseDcrProviderCurity(components.DCRProviderCurity{/* values here */})
```

### DCRProviderOKTA

```go
createDcrProviderResponse := components.CreateCreateDcrProviderResponseDcrProviderOkta(components.DCRProviderOKTA{/* values here */})
```

### DCRProviderHTTP

```go
createDcrProviderResponse := components.CreateCreateDcrProviderResponseDcrProviderHTTP(components.DCRProviderHTTP{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createDcrProviderResponse.Type {
	case components.CreateDcrProviderResponseTypeDcrProviderAuth0:
		// createDcrProviderResponse.DCRProviderAuth0 is populated
	case components.CreateDcrProviderResponseTypeDcrProviderAzureAd:
		// createDcrProviderResponse.DCRProviderAzureAD is populated
	case components.CreateDcrProviderResponseTypeDcrProviderCurity:
		// createDcrProviderResponse.DCRProviderCurity is populated
	case components.CreateDcrProviderResponseTypeDcrProviderOkta:
		// createDcrProviderResponse.DCRProviderOKTA is populated
	case components.CreateDcrProviderResponseTypeDcrProviderHTTP:
		// createDcrProviderResponse.DCRProviderHTTP is populated
}
```
