# AddOnConfigKindFieldEqualsFilter

Filter add-on config kind by exact match.


## Supported Types

### AddOnConfigKind

```go
addOnConfigKindFieldEqualsFilter := components.CreateAddOnConfigKindFieldEqualsFilterAddOnConfigKind(components.AddOnConfigKind{/* values here */})
```

### AddOnConfigKindFieldEqualsComparison

```go
addOnConfigKindFieldEqualsFilter := components.CreateAddOnConfigKindFieldEqualsFilterAddOnConfigKindFieldEqualsComparison(components.AddOnConfigKindFieldEqualsComparison{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch addOnConfigKindFieldEqualsFilter.Type {
	case components.AddOnConfigKindFieldEqualsFilterTypeAddOnConfigKind:
		// addOnConfigKindFieldEqualsFilter.AddOnConfigKind is populated
	case components.AddOnConfigKindFieldEqualsFilterTypeAddOnConfigKindFieldEqualsComparison:
		// addOnConfigKindFieldEqualsFilter.AddOnConfigKindFieldEqualsComparison is populated
}
```
