# DcrProviderResponse

A response containing a single DCR provider object. Sensitive fields will be removed from the response.


## Supported Types

### DCRProviderAuth0DCRProviderAuth0

```go
dcrProviderResponse := components.CreateDcrProviderResponseDcrProviderAuth0(components.DCRProviderAuth0DCRProviderAuth0{/* values here */})
```

### DCRProviderAzureADDCRProviderAzureAD

```go
dcrProviderResponse := components.CreateDcrProviderResponseDcrProviderAzureAd(components.DCRProviderAzureADDCRProviderAzureAD{/* values here */})
```

### DCRProviderCurityDCRProviderCurity

```go
dcrProviderResponse := components.CreateDcrProviderResponseDcrProviderCurity(components.DCRProviderCurityDCRProviderCurity{/* values here */})
```

### DCRProviderOKTADCRProviderOKTA

```go
dcrProviderResponse := components.CreateDcrProviderResponseDcrProviderOkta(components.DCRProviderOKTADCRProviderOKTA{/* values here */})
```

### DCRProviderHTTPDCRProviderHTTP

```go
dcrProviderResponse := components.CreateDcrProviderResponseDcrProviderHTTP(components.DCRProviderHTTPDCRProviderHTTP{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch dcrProviderResponse.Type {
	case components.DcrProviderResponseTypeDcrProviderAuth0:
		// dcrProviderResponse.DCRProviderAuth0DCRProviderAuth0 is populated
	case components.DcrProviderResponseTypeDcrProviderAzureAd:
		// dcrProviderResponse.DCRProviderAzureADDCRProviderAzureAD is populated
	case components.DcrProviderResponseTypeDcrProviderCurity:
		// dcrProviderResponse.DCRProviderCurityDCRProviderCurity is populated
	case components.DcrProviderResponseTypeDcrProviderOkta:
		// dcrProviderResponse.DCRProviderOKTADCRProviderOKTA is populated
	case components.DcrProviderResponseTypeDcrProviderHTTP:
		// dcrProviderResponse.DCRProviderHTTPDCRProviderHTTP is populated
}
```
