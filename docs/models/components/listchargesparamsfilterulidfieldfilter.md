# ListChargesParamsFilterULIDFieldFilter

Filter charges by the ID of their associated feature.


## Supported Types

### 

```go
listChargesParamsFilterULIDFieldFilter := components.CreateListChargesParamsFilterULIDFieldFilterStr(string{/* values here */})
```

### ListChargesParamsFilterULIDFieldFilter2

```go
listChargesParamsFilterULIDFieldFilter := components.CreateListChargesParamsFilterULIDFieldFilterListChargesParamsFilterULIDFieldFilter2(components.ListChargesParamsFilterULIDFieldFilter2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listChargesParamsFilterULIDFieldFilter.Type {
	case components.ListChargesParamsFilterULIDFieldFilterTypeStr:
		// listChargesParamsFilterULIDFieldFilter.Str is populated
	case components.ListChargesParamsFilterULIDFieldFilterTypeListChargesParamsFilterULIDFieldFilter2:
		// listChargesParamsFilterULIDFieldFilter.ListChargesParamsFilterULIDFieldFilter2 is populated
}
```
