# CreatedAt

Filters on the given datetime (RFC-3339) field value.


## Supported Types

### 

```go
createdAt := components.CreateCreatedAtDateTime(time.Time{/* values here */})
```

### CreatedAtDateTimeFieldEqualsFilter

```go
createdAt := components.CreateCreatedAtCreatedAtDateTimeFieldEqualsFilter(components.CreatedAtDateTimeFieldEqualsFilter{/* values here */})
```

### CreatedAtDateTimeFieldLTFilter

```go
createdAt := components.CreateCreatedAtCreatedAtDateTimeFieldLTFilter(components.CreatedAtDateTimeFieldLTFilter{/* values here */})
```

### CreatedAtDateTimeFieldLTEFilter

```go
createdAt := components.CreateCreatedAtCreatedAtDateTimeFieldLTEFilter(components.CreatedAtDateTimeFieldLTEFilter{/* values here */})
```

### CreatedAtDateTimeFieldGTFilter

```go
createdAt := components.CreateCreatedAtCreatedAtDateTimeFieldGTFilter(components.CreatedAtDateTimeFieldGTFilter{/* values here */})
```

### CreatedAtDateTimeFieldGTEFilter

```go
createdAt := components.CreateCreatedAtCreatedAtDateTimeFieldGTEFilter(components.CreatedAtDateTimeFieldGTEFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createdAt.Type {
	case components.CreatedAtTypeDateTime:
		// createdAt.DateTime is populated
	case components.CreatedAtTypeCreatedAtDateTimeFieldEqualsFilter:
		// createdAt.CreatedAtDateTimeFieldEqualsFilter is populated
	case components.CreatedAtTypeCreatedAtDateTimeFieldLTFilter:
		// createdAt.CreatedAtDateTimeFieldLTFilter is populated
	case components.CreatedAtTypeCreatedAtDateTimeFieldLTEFilter:
		// createdAt.CreatedAtDateTimeFieldLTEFilter is populated
	case components.CreatedAtTypeCreatedAtDateTimeFieldGTFilter:
		// createdAt.CreatedAtDateTimeFieldGTFilter is populated
	case components.CreatedAtTypeCreatedAtDateTimeFieldGTEFilter:
		// createdAt.CreatedAtDateTimeFieldGTEFilter is populated
}
```
