# DateTimeFieldFilter

Filters on the given datetime (RFC-3339) field value.


## Supported Types

### 

```go
dateTimeFieldFilter := components.CreateDateTimeFieldFilterDateTime(time.Time{/* values here */})
```

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

