# DcrProviderResponse

A response containing a single DCR provider object. Sensitive fields will be removed from the response.


## Supported Types

### DCRProviderAuth0DCRProviderAuth0

```go
dcrProviderResponse := components.CreateDcrProviderResponseAuth0(components.DCRProviderAuth0DCRProviderAuth0{/* values here */})
```

### DCRProviderAzureADDCRProviderAzureAD

```go
dcrProviderResponse := components.CreateDcrProviderResponseAzureAd(components.DCRProviderAzureADDCRProviderAzureAD{/* values here */})
```

### DCRProviderCurityDCRProviderCurity

```go
dcrProviderResponse := components.CreateDcrProviderResponseCurity(components.DCRProviderCurityDCRProviderCurity{/* values here */})
```

### DCRProviderOKTADCRProviderOKTA

```go
dcrProviderResponse := components.CreateDcrProviderResponseOkta(components.DCRProviderOKTADCRProviderOKTA{/* values here */})
```

### DCRProviderKongIdentityDCRProviderKongIdentity

```go
dcrProviderResponse := components.CreateDcrProviderResponseKongIdentity(components.DCRProviderKongIdentityDCRProviderKongIdentity{/* values here */})
```

### DCRProviderHTTPDCRProviderHTTP

```go
dcrProviderResponse := components.CreateDcrProviderResponseHTTP(components.DCRProviderHTTPDCRProviderHTTP{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch dcrProviderResponse.Type {
	case components.DcrProviderResponseTypeAuth0:
		// dcrProviderResponse.DCRProviderAuth0DCRProviderAuth0 is populated
	case components.DcrProviderResponseTypeAzureAd:
		// dcrProviderResponse.DCRProviderAzureADDCRProviderAzureAD is populated
	case components.DcrProviderResponseTypeCurity:
		// dcrProviderResponse.DCRProviderCurityDCRProviderCurity is populated
	case components.DcrProviderResponseTypeOkta:
		// dcrProviderResponse.DCRProviderOKTADCRProviderOKTA is populated
	case components.DcrProviderResponseTypeKongIdentity:
		// dcrProviderResponse.DCRProviderKongIdentityDCRProviderKongIdentity is populated
	case components.DcrProviderResponseTypeHTTP:
		// dcrProviderResponse.DCRProviderHTTPDCRProviderHTTP is populated
}
```
