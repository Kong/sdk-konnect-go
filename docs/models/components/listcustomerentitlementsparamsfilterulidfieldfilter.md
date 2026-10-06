# ListCustomerEntitlementsParamsFilterULIDFieldFilter

Filter entitlements by feature ID.


## Supported Types

### 

```go
listCustomerEntitlementsParamsFilterULIDFieldFilter := components.CreateListCustomerEntitlementsParamsFilterULIDFieldFilterStr(string{/* values here */})
```

### ListCustomerEntitlementsParamsFilterULIDFieldFilter2

```go
listCustomerEntitlementsParamsFilterULIDFieldFilter := components.CreateListCustomerEntitlementsParamsFilterULIDFieldFilterListCustomerEntitlementsParamsFilterULIDFieldFilter2(components.ListCustomerEntitlementsParamsFilterULIDFieldFilter2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listCustomerEntitlementsParamsFilterULIDFieldFilter.Type {
	case components.ListCustomerEntitlementsParamsFilterULIDFieldFilterTypeStr:
		// listCustomerEntitlementsParamsFilterULIDFieldFilter.Str is populated
	case components.ListCustomerEntitlementsParamsFilterULIDFieldFilterTypeListCustomerEntitlementsParamsFilterULIDFieldFilter2:
		// listCustomerEntitlementsParamsFilterULIDFieldFilter.ListCustomerEntitlementsParamsFilterULIDFieldFilter2 is populated
}
```
