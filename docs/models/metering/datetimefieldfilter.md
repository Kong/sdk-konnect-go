# DateTimeFieldFilter

Filter events by event time.


## Supported Types

### DateTimeFieldEqualsFilter

```go
dateTimeFieldFilter := components.CreateDateTimeFieldFilterDateTimeFieldEqualsFilter(metering.DateTimeFieldEqualsFilter{/* values here */})
```

### DateTimeFieldLTFilter

```go
dateTimeFieldFilter := components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(metering.DateTimeFieldLTFilter{/* values here */})
```

### DateTimeFieldLTEFilter

```go
dateTimeFieldFilter := components.CreateDateTimeFieldFilterDateTimeFieldLTEFilter(metering.DateTimeFieldLTEFilter{/* values here */})
```

### DateTimeFieldGTFilter

```go
dateTimeFieldFilter := components.CreateDateTimeFieldFilterDateTimeFieldGTFilter(metering.DateTimeFieldGTFilter{/* values here */})
```

### DateTimeFieldGTEFilter

```go
dateTimeFieldFilter := components.CreateDateTimeFieldFilterDateTimeFieldGTEFilter(metering.DateTimeFieldGTEFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch dateTimeFieldFilter.Type {
	case components.DateTimeFieldFilterTypeDateTimeFieldEqualsFilter:
		// dateTimeFieldFilter.DateTimeFieldEqualsFilter is populated
	case components.DateTimeFieldFilterTypeDateTimeFieldLTFilter:
		// dateTimeFieldFilter.DateTimeFieldLTFilter is populated
	case components.DateTimeFieldFilterTypeDateTimeFieldLTEFilter:
		// dateTimeFieldFilter.DateTimeFieldLTEFilter is populated
	case components.DateTimeFieldFilterTypeDateTimeFieldGTFilter:
		// dateTimeFieldFilter.DateTimeFieldGTFilter is populated
	case components.DateTimeFieldFilterTypeDateTimeFieldGTEFilter:
		// dateTimeFieldFilter.DateTimeFieldGTEFilter is populated
}
```
