# InvalidParameters


## Supported Types

### InvalidParameterStandard

```go
invalidParameters := components.CreateInvalidParametersInvalidParameterStandard(components.InvalidParameterStandard{/* values here */})
```

### InvalidParameterMinimumLength

```go
invalidParameters := components.CreateInvalidParametersInvalidParameterMinimumLength(components.InvalidParameterMinimumLength{/* values here */})
```

### InvalidParameterMaximumLength

```go
invalidParameters := components.CreateInvalidParametersInvalidParameterMaximumLength(components.InvalidParameterMaximumLength{/* values here */})
```

### InvalidParameterChoiceItem

```go
invalidParameters := components.CreateInvalidParametersInvalidParameterChoiceItem(components.InvalidParameterChoiceItem{/* values here */})
```

### InvalidParameterDependentItem

```go
invalidParameters := components.CreateInvalidParametersInvalidParameterDependentItem(components.InvalidParameterDependentItem{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch invalidParameters.Type {
	case components.InvalidParametersTypeInvalidParameterStandard:
		// invalidParameters.InvalidParameterStandard is populated
	case components.InvalidParametersTypeInvalidParameterMinimumLength:
		// invalidParameters.InvalidParameterMinimumLength is populated
	case components.InvalidParametersTypeInvalidParameterMaximumLength:
		// invalidParameters.InvalidParameterMaximumLength is populated
	case components.InvalidParametersTypeInvalidParameterChoiceItem:
		// invalidParameters.InvalidParameterChoiceItem is populated
	case components.InvalidParametersTypeInvalidParameterDependentItem:
		// invalidParameters.InvalidParameterDependentItem is populated
}
```
