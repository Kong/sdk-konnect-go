# ListCustomerChargesParamsFilterULIDFieldFilter

Filter charges by the ID of their associated feature.


## Supported Types

### 

```go
listCustomerChargesParamsFilterULIDFieldFilter := components.CreateListCustomerChargesParamsFilterULIDFieldFilterStr(string{/* values here */})
```

### ListCustomerChargesParamsFilterULIDFieldFilter2

```go
listCustomerChargesParamsFilterULIDFieldFilter := components.CreateListCustomerChargesParamsFilterULIDFieldFilterListCustomerChargesParamsFilterULIDFieldFilter2(components.ListCustomerChargesParamsFilterULIDFieldFilter2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listCustomerChargesParamsFilterULIDFieldFilter.Type {
	case components.ListCustomerChargesParamsFilterULIDFieldFilterTypeStr:
		// listCustomerChargesParamsFilterULIDFieldFilter.Str is populated
	case components.ListCustomerChargesParamsFilterULIDFieldFilterTypeListCustomerChargesParamsFilterULIDFieldFilter2:
		// listCustomerChargesParamsFilterULIDFieldFilter.ListCustomerChargesParamsFilterULIDFieldFilter2 is populated
}
```
