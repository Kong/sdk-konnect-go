# ListEntitlementsParamsFilterULIDFieldFilter

Filter entitlements by feature ID.


## Supported Types

### 

```go
listEntitlementsParamsFilterULIDFieldFilter := components.CreateListEntitlementsParamsFilterULIDFieldFilterStr(string{/* values here */})
```

### ListEntitlementsParamsFilterULIDFieldFilter2

```go
listEntitlementsParamsFilterULIDFieldFilter := components.CreateListEntitlementsParamsFilterULIDFieldFilterListEntitlementsParamsFilterULIDFieldFilter2(components.ListEntitlementsParamsFilterULIDFieldFilter2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listEntitlementsParamsFilterULIDFieldFilter.Type {
	case components.ListEntitlementsParamsFilterULIDFieldFilterTypeStr:
		// listEntitlementsParamsFilterULIDFieldFilter.Str is populated
	case components.ListEntitlementsParamsFilterULIDFieldFilterTypeListEntitlementsParamsFilterULIDFieldFilter2:
		// listEntitlementsParamsFilterULIDFieldFilter.ListEntitlementsParamsFilterULIDFieldFilter2 is populated
}
```
