# DateTimeFieldFilter

Filters on the given datetime (RFC-3339) field value.


## Supported Types

### DateTimeFieldEqualsFilter

```go
dateTimeFieldFilter := components.CreateDateTimeFieldFilterDateTimeFieldEqualsFilter(components.DateTimeFieldEqualsFilter{/* values here */})
```

### DateTimeFieldLTFilter

```go
dateTimeFieldFilter := components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(components.DateTimeFieldLTFilter{/* values here */})
```

### DateTimeFieldLTEFilter

```go
dateTimeFieldFilter := components.CreateDateTimeFieldFilterDateTimeFieldLTEFilter(components.DateTimeFieldLTEFilter{/* values here */})
```

### DateTimeFieldGTFilter

```go
dateTimeFieldFilter := components.CreateDateTimeFieldFilterDateTimeFieldGTFilter(components.DateTimeFieldGTFilter{/* values here */})
```

### DateTimeFieldGTEFilter

```go
dateTimeFieldFilter := components.CreateDateTimeFieldFilterDateTimeFieldGTEFilter(components.DateTimeFieldGTEFilter{/* values here */})
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
