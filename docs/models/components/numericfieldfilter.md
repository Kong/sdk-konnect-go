# NumericFieldFilter

Filter by a numeric value.


## Supported Types

### 

```go
numericFieldFilter := components.CreateNumericFieldFilterNumber(float64{/* values here */})
```

### NumericFieldEqualsFilter

```go
numericFieldFilter := components.CreateNumericFieldFilterNumericFieldEqualsFilter(components.NumericFieldEqualsFilter{/* values here */})
```

### NumericFieldLTFilter

```go
numericFieldFilter := components.CreateNumericFieldFilterNumericFieldLTFilter(components.NumericFieldLTFilter{/* values here */})
```

### NumericFieldLTEFilter

```go
numericFieldFilter := components.CreateNumericFieldFilterNumericFieldLTEFilter(components.NumericFieldLTEFilter{/* values here */})
```

### NumericFieldGTFilter

```go
numericFieldFilter := components.CreateNumericFieldFilterNumericFieldGTFilter(components.NumericFieldGTFilter{/* values here */})
```

### NumericFieldGTEFilter

```go
numericFieldFilter := components.CreateNumericFieldFilterNumericFieldGTEFilter(components.NumericFieldGTEFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch numericFieldFilter.Type {
	case components.NumericFieldFilterTypeNumber:
		// numericFieldFilter.Number is populated
	case components.NumericFieldFilterTypeNumericFieldEqualsFilter:
		// numericFieldFilter.NumericFieldEqualsFilter is populated
	case components.NumericFieldFilterTypeNumericFieldLTFilter:
		// numericFieldFilter.NumericFieldLTFilter is populated
	case components.NumericFieldFilterTypeNumericFieldLTEFilter:
		// numericFieldFilter.NumericFieldLTEFilter is populated
	case components.NumericFieldFilterTypeNumericFieldGTFilter:
		// numericFieldFilter.NumericFieldGTFilter is populated
	case components.NumericFieldFilterTypeNumericFieldGTEFilter:
		// numericFieldFilter.NumericFieldGTEFilter is populated
}
```
