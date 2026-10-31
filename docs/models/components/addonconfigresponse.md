# AddOnConfigResponse

Configuration object for different types of add-ons.


## Supported Types

### ManagedCacheAddOnConfigResponse

```go
addOnConfigResponse := components.CreateAddOnConfigResponseManagedCacheAddOnConfigResponse(components.ManagedCacheAddOnConfigResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch addOnConfigResponse.Type {
	case components.AddOnConfigResponseTypeManagedCacheAddOnConfigResponse:
		// addOnConfigResponse.ManagedCacheAddOnConfigResponse is populated
}
```
