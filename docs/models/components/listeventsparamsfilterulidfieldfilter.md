# ListEventsParamsFilterULIDFieldFilter

Filter events by the associated customer ID.


## Supported Types

### 

```go
listEventsParamsFilterULIDFieldFilter := components.CreateListEventsParamsFilterULIDFieldFilterStr(string{/* values here */})
```

### ListEventsParamsFilterULIDFieldFilter2

```go
listEventsParamsFilterULIDFieldFilter := components.CreateListEventsParamsFilterULIDFieldFilterListEventsParamsFilterULIDFieldFilter2(components.ListEventsParamsFilterULIDFieldFilter2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listEventsParamsFilterULIDFieldFilter.Type {
	case components.ListEventsParamsFilterULIDFieldFilterTypeStr:
		// listEventsParamsFilterULIDFieldFilter.Str is populated
	case components.ListEventsParamsFilterULIDFieldFilterTypeListEventsParamsFilterULIDFieldFilter2:
		// listEventsParamsFilterULIDFieldFilter.ListEventsParamsFilterULIDFieldFilter2 is populated
}
```
