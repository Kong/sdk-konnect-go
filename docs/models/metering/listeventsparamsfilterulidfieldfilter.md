# ListEventsParamsFilterULIDFieldFilter

Filter events by the associated customer ID.


## Supported Types

### 

```go
listEventsParamsFilterULIDFieldFilter := components.CreateListEventsParamsFilterULIDFieldFilterStr(string{/* values here */})
```

### ULIDFieldFilter2

```go
listEventsParamsFilterULIDFieldFilter := components.CreateListEventsParamsFilterULIDFieldFilterULIDFieldFilter2(metering.ULIDFieldFilter2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listEventsParamsFilterULIDFieldFilter.Type {
	case components.ListEventsParamsFilterULIDFieldFilterTypeStr:
		// listEventsParamsFilterULIDFieldFilter.Str is populated
	case components.ListEventsParamsFilterULIDFieldFilterTypeULIDFieldFilter2:
		// listEventsParamsFilterULIDFieldFilter.ULIDFieldFilter2 is populated
}
```
