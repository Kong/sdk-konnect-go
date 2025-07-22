# AttributesFieldFilter

Filters on the given string field value by either exact or fuzzy match.


## Supported Types

### StringFieldEqualsFilter

```go
attributesFieldFilter := components.CreateAttributesFieldFilterStringFieldEqualsFilter(components.StringFieldEqualsFilter{/* values here */})
```

### StringFieldContainsFilter

```go
attributesFieldFilter := components.CreateAttributesFieldFilterStringFieldContainsFilter(components.StringFieldContainsFilter{/* values here */})
```

### StringFieldOContainsFilter

```go
attributesFieldFilter := components.CreateAttributesFieldFilterStringFieldOContainsFilter(components.StringFieldOContainsFilter{/* values here */})
```

### StringFieldOEQFilter

```go
attributesFieldFilter := components.CreateAttributesFieldFilterStringFieldOEQFilter(components.StringFieldOEQFilter{/* values here */})
```

### StringFieldNEQFilter

```go
attributesFieldFilter := components.CreateAttributesFieldFilterStringFieldNEQFilter(components.StringFieldNEQFilter{/* values here */})
```

