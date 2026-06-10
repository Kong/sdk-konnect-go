# ListEventsParamsFilterDateTimeFieldFilter

Filter events by the time the event was ingested.


## Supported Types

### DateTimeFieldFilterDateTimeFieldEqualsFilter

```go
listEventsParamsFilterDateTimeFieldFilter := components.CreateListEventsParamsFilterDateTimeFieldFilterDateTimeFieldFilterDateTimeFieldEqualsFilter(metering.DateTimeFieldFilterDateTimeFieldEqualsFilter{/* values here */})
```

### DateTimeFieldFilterDateTimeFieldLTFilter

```go
listEventsParamsFilterDateTimeFieldFilter := components.CreateListEventsParamsFilterDateTimeFieldFilterDateTimeFieldFilterDateTimeFieldLTFilter(metering.DateTimeFieldFilterDateTimeFieldLTFilter{/* values here */})
```

### DateTimeFieldFilterDateTimeFieldLTEFilter

```go
listEventsParamsFilterDateTimeFieldFilter := components.CreateListEventsParamsFilterDateTimeFieldFilterDateTimeFieldFilterDateTimeFieldLTEFilter(metering.DateTimeFieldFilterDateTimeFieldLTEFilter{/* values here */})
```

### DateTimeFieldFilterDateTimeFieldGTFilter

```go
listEventsParamsFilterDateTimeFieldFilter := components.CreateListEventsParamsFilterDateTimeFieldFilterDateTimeFieldFilterDateTimeFieldGTFilter(metering.DateTimeFieldFilterDateTimeFieldGTFilter{/* values here */})
```

### DateTimeFieldFilterDateTimeFieldGTEFilter

```go
listEventsParamsFilterDateTimeFieldFilter := components.CreateListEventsParamsFilterDateTimeFieldFilterDateTimeFieldFilterDateTimeFieldGTEFilter(metering.DateTimeFieldFilterDateTimeFieldGTEFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listEventsParamsFilterDateTimeFieldFilter.Type {
	case components.ListEventsParamsFilterDateTimeFieldFilterTypeDateTimeFieldFilterDateTimeFieldEqualsFilter:
		// listEventsParamsFilterDateTimeFieldFilter.DateTimeFieldFilterDateTimeFieldEqualsFilter is populated
	case components.ListEventsParamsFilterDateTimeFieldFilterTypeDateTimeFieldFilterDateTimeFieldLTFilter:
		// listEventsParamsFilterDateTimeFieldFilter.DateTimeFieldFilterDateTimeFieldLTFilter is populated
	case components.ListEventsParamsFilterDateTimeFieldFilterTypeDateTimeFieldFilterDateTimeFieldLTEFilter:
		// listEventsParamsFilterDateTimeFieldFilter.DateTimeFieldFilterDateTimeFieldLTEFilter is populated
	case components.ListEventsParamsFilterDateTimeFieldFilterTypeDateTimeFieldFilterDateTimeFieldGTFilter:
		// listEventsParamsFilterDateTimeFieldFilter.DateTimeFieldFilterDateTimeFieldGTFilter is populated
	case components.ListEventsParamsFilterDateTimeFieldFilterTypeDateTimeFieldFilterDateTimeFieldGTEFilter:
		// listEventsParamsFilterDateTimeFieldFilter.DateTimeFieldFilterDateTimeFieldGTEFilter is populated
}
```
