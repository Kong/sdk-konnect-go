# ListGrantsParamsFilterULIDFieldFilter

Filter grants by the ID of the customer that owns the entitlement.


## Supported Types

### 

```go
listGrantsParamsFilterULIDFieldFilter := components.CreateListGrantsParamsFilterULIDFieldFilterStr(string{/* values here */})
```

### ListGrantsParamsFilterULIDFieldFilter2

```go
listGrantsParamsFilterULIDFieldFilter := components.CreateListGrantsParamsFilterULIDFieldFilterListGrantsParamsFilterULIDFieldFilter2(components.ListGrantsParamsFilterULIDFieldFilter2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listGrantsParamsFilterULIDFieldFilter.Type {
	case components.ListGrantsParamsFilterULIDFieldFilterTypeStr:
		// listGrantsParamsFilterULIDFieldFilter.Str is populated
	case components.ListGrantsParamsFilterULIDFieldFilterTypeListGrantsParamsFilterULIDFieldFilter2:
		// listGrantsParamsFilterULIDFieldFilter.ListGrantsParamsFilterULIDFieldFilter2 is populated
}
```
