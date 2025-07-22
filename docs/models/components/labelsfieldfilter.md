# LabelsFieldFilter

Filters on the given string field value by either exact or fuzzy match.


## Supported Types

### StringFieldEqualsFilter

```go
labelsFieldFilter := components.CreateLabelsFieldFilterStringFieldEqualsFilter(components.StringFieldEqualsFilter{/* values here */})
```

### StringFieldContainsFilter

```go
labelsFieldFilter := components.CreateLabelsFieldFilterStringFieldContainsFilter(components.StringFieldContainsFilter{/* values here */})
```

### StringFieldOContainsFilter

```go
labelsFieldFilter := components.CreateLabelsFieldFilterStringFieldOContainsFilter(components.StringFieldOContainsFilter{/* values here */})
```

### StringFieldOEQFilter

```go
labelsFieldFilter := components.CreateLabelsFieldFilterStringFieldOEQFilter(components.StringFieldOEQFilter{/* values here */})
```

### StringFieldNEQFilter

```go
labelsFieldFilter := components.CreateLabelsFieldFilterStringFieldNEQFilter(components.StringFieldNEQFilter{/* values here */})
```

