# ListGrantsParamsFilterFeatureIDULIDFieldFilter

Filter grants by the ID of the entitlement's feature.


## Supported Types

### 

```go
listGrantsParamsFilterFeatureIDULIDFieldFilter := components.CreateListGrantsParamsFilterFeatureIDULIDFieldFilterStr(string{/* values here */})
```

### ListGrantsParamsFilterULIDFieldFilterFeatureID2

```go
listGrantsParamsFilterFeatureIDULIDFieldFilter := components.CreateListGrantsParamsFilterFeatureIDULIDFieldFilterListGrantsParamsFilterULIDFieldFilterFeatureID2(components.ListGrantsParamsFilterULIDFieldFilterFeatureID2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listGrantsParamsFilterFeatureIDULIDFieldFilter.Type {
	case components.ListGrantsParamsFilterFeatureIDULIDFieldFilterTypeStr:
		// listGrantsParamsFilterFeatureIDULIDFieldFilter.Str is populated
	case components.ListGrantsParamsFilterFeatureIDULIDFieldFilterTypeListGrantsParamsFilterULIDFieldFilterFeatureID2:
		// listGrantsParamsFilterFeatureIDULIDFieldFilter.ListGrantsParamsFilterULIDFieldFilterFeatureID2 is populated
}
```
