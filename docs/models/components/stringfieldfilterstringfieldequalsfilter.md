# StringFieldFilterStringFieldEqualsFilter

Filters on the given string field value by exact match.


## Supported Types

### 

```go
stringFieldFilterStringFieldEqualsFilter := components.CreateStringFieldFilterStringFieldEqualsFilterStr(string{/* values here */})
```

### StringFieldEqualsFilterStringFieldEqualsComparison

```go
stringFieldFilterStringFieldEqualsFilter := components.CreateStringFieldFilterStringFieldEqualsFilterStringFieldEqualsFilterStringFieldEqualsComparison(components.StringFieldEqualsFilterStringFieldEqualsComparison{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch stringFieldFilterStringFieldEqualsFilter.Type {
	case components.StringFieldFilterStringFieldEqualsFilterTypeStr:
		// stringFieldFilterStringFieldEqualsFilter.Str is populated
	case components.StringFieldFilterStringFieldEqualsFilterTypeStringFieldEqualsFilterStringFieldEqualsComparison:
		// stringFieldFilterStringFieldEqualsFilter.StringFieldEqualsFilterStringFieldEqualsComparison is populated
}
```
