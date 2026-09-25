# ListChargesParamsFilterCustomerIDULIDFieldFilter

Filter charges by the ID of their customer.


## Supported Types

### 

```go
listChargesParamsFilterCustomerIDULIDFieldFilter := components.CreateListChargesParamsFilterCustomerIDULIDFieldFilterStr(string{/* values here */})
```

### ListChargesParamsFilterULIDFieldFilterCustomerID2

```go
listChargesParamsFilterCustomerIDULIDFieldFilter := components.CreateListChargesParamsFilterCustomerIDULIDFieldFilterListChargesParamsFilterULIDFieldFilterCustomerID2(components.ListChargesParamsFilterULIDFieldFilterCustomerID2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listChargesParamsFilterCustomerIDULIDFieldFilter.Type {
	case components.ListChargesParamsFilterCustomerIDULIDFieldFilterTypeStr:
		// listChargesParamsFilterCustomerIDULIDFieldFilter.Str is populated
	case components.ListChargesParamsFilterCustomerIDULIDFieldFilterTypeListChargesParamsFilterULIDFieldFilterCustomerID2:
		// listChargesParamsFilterCustomerIDULIDFieldFilter.ListChargesParamsFilterULIDFieldFilterCustomerID2 is populated
}
```
