# ListInvoicesParamsFilterULIDFieldFilter

Filter by customer ID.


## Supported Types

### 

```go
listInvoicesParamsFilterULIDFieldFilter := components.CreateListInvoicesParamsFilterULIDFieldFilterStr(string{/* values here */})
```

### ListInvoicesParamsFilterULIDFieldFilter2

```go
listInvoicesParamsFilterULIDFieldFilter := components.CreateListInvoicesParamsFilterULIDFieldFilterListInvoicesParamsFilterULIDFieldFilter2(components.ListInvoicesParamsFilterULIDFieldFilter2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listInvoicesParamsFilterULIDFieldFilter.Type {
	case components.ListInvoicesParamsFilterULIDFieldFilterTypeStr:
		// listInvoicesParamsFilterULIDFieldFilter.Str is populated
	case components.ListInvoicesParamsFilterULIDFieldFilterTypeListInvoicesParamsFilterULIDFieldFilter2:
		// listInvoicesParamsFilterULIDFieldFilter.ListInvoicesParamsFilterULIDFieldFilter2 is populated
}
```
