# AddOnStateFieldFilter


## Supported Types

### AddOnStateFieldEqualsFilter

```go
addOnStateFieldFilter := components.CreateAddOnStateFieldFilterAddOnStateFieldEqualsFilter(components.AddOnStateFieldEqualsFilter{/* values here */})
```

### AddOnStateFieldNotEqualsFilter

```go
addOnStateFieldFilter := components.CreateAddOnStateFieldFilterAddOnStateFieldNotEqualsFilter(components.AddOnStateFieldNotEqualsFilter{/* values here */})
```

### AddOnStateFieldOrEqualityFilter

```go
addOnStateFieldFilter := components.CreateAddOnStateFieldFilterAddOnStateFieldOrEqualityFilter(components.AddOnStateFieldOrEqualityFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch addOnStateFieldFilter.Type {
	case components.AddOnStateFieldFilterTypeAddOnStateFieldEqualsFilter:
		// addOnStateFieldFilter.AddOnStateFieldEqualsFilter is populated
	case components.AddOnStateFieldFilterTypeAddOnStateFieldNotEqualsFilter:
		// addOnStateFieldFilter.AddOnStateFieldNotEqualsFilter is populated
	case components.AddOnStateFieldFilterTypeAddOnStateFieldOrEqualityFilter:
		// addOnStateFieldFilter.AddOnStateFieldOrEqualityFilter is populated
}
```
