# AddOnStateFieldEqualsFilter

Filter add-on state by exact match.


## Supported Types

### AddOnState

```go
addOnStateFieldEqualsFilter := components.CreateAddOnStateFieldEqualsFilterAddOnState(components.AddOnState{/* values here */})
```

### AddOnStateFieldEqualsComparison

```go
addOnStateFieldEqualsFilter := components.CreateAddOnStateFieldEqualsFilterAddOnStateFieldEqualsComparison(components.AddOnStateFieldEqualsComparison{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch addOnStateFieldEqualsFilter.Type {
	case components.AddOnStateFieldEqualsFilterTypeAddOnState:
		// addOnStateFieldEqualsFilter.AddOnState is populated
	case components.AddOnStateFieldEqualsFilterTypeAddOnStateFieldEqualsComparison:
		// addOnStateFieldEqualsFilter.AddOnStateFieldEqualsComparison is populated
}
```
