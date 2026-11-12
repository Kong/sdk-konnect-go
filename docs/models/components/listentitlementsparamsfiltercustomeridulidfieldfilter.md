# ListEntitlementsParamsFilterCustomerIDULIDFieldFilter

Filter entitlements by customer ID.


## Supported Types

### 

```go
listEntitlementsParamsFilterCustomerIDULIDFieldFilter := components.CreateListEntitlementsParamsFilterCustomerIDULIDFieldFilterStr(string{/* values here */})
```

### ListEntitlementsParamsFilterULIDFieldFilterCustomerID2

```go
listEntitlementsParamsFilterCustomerIDULIDFieldFilter := components.CreateListEntitlementsParamsFilterCustomerIDULIDFieldFilterListEntitlementsParamsFilterULIDFieldFilterCustomerID2(components.ListEntitlementsParamsFilterULIDFieldFilterCustomerID2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listEntitlementsParamsFilterCustomerIDULIDFieldFilter.Type {
	case components.ListEntitlementsParamsFilterCustomerIDULIDFieldFilterTypeStr:
		// listEntitlementsParamsFilterCustomerIDULIDFieldFilter.Str is populated
	case components.ListEntitlementsParamsFilterCustomerIDULIDFieldFilterTypeListEntitlementsParamsFilterULIDFieldFilterCustomerID2:
		// listEntitlementsParamsFilterCustomerIDULIDFieldFilter.ListEntitlementsParamsFilterULIDFieldFilterCustomerID2 is populated
}
```
