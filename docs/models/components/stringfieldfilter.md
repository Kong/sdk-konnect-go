# StringFieldFilter

Filters on the given string field value by either exact or fuzzy match.


## Supported Types

### StringFieldEqualsFilter

```go
stringFieldFilter := components.CreateStringFieldFilterStringFieldEqualsFilter(components.StringFieldEqualsFilter{/* values here */})
```

### StringFieldContainsFilter

```go
stringFieldFilter := components.CreateStringFieldFilterStringFieldContainsFilter(components.StringFieldContainsFilter{/* values here */})
```

### StringFieldOContainsFilter

```go
stringFieldFilter := components.CreateStringFieldFilterStringFieldOContainsFilter(components.StringFieldOContainsFilter{/* values here */})
```

### StringFieldOEQFilter

```go
stringFieldFilter := components.CreateStringFieldFilterStringFieldOEQFilter(components.StringFieldOEQFilter{/* values here */})
```

### StringFieldNEQFilter

```go
stringFieldFilter := components.CreateStringFieldFilterStringFieldNEQFilter(components.StringFieldNEQFilter{/* values here */})
```

