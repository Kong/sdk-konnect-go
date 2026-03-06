# AddOnConfigKindFieldFilter

Filter for add-on config kind field.


## Supported Types

### AddOnConfigKindFieldEqualsFilter

```go
addOnConfigKindFieldFilter := components.CreateAddOnConfigKindFieldFilterAddOnConfigKindFieldEqualsFilter(components.AddOnConfigKindFieldEqualsFilter{/* values here */})
```

### AddOnConfigKindFieldNotEqualsFilter

```go
addOnConfigKindFieldFilter := components.CreateAddOnConfigKindFieldFilterAddOnConfigKindFieldNotEqualsFilter(components.AddOnConfigKindFieldNotEqualsFilter{/* values here */})
```

### AddOnConfigKindFieldOrEqualityFilter

```go
addOnConfigKindFieldFilter := components.CreateAddOnConfigKindFieldFilterAddOnConfigKindFieldOrEqualityFilter(components.AddOnConfigKindFieldOrEqualityFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch addOnConfigKindFieldFilter.Type {
	case components.AddOnConfigKindFieldFilterTypeAddOnConfigKindFieldEqualsFilter:
		// addOnConfigKindFieldFilter.AddOnConfigKindFieldEqualsFilter is populated
	case components.AddOnConfigKindFieldFilterTypeAddOnConfigKindFieldNotEqualsFilter:
		// addOnConfigKindFieldFilter.AddOnConfigKindFieldNotEqualsFilter is populated
	case components.AddOnConfigKindFieldFilterTypeAddOnConfigKindFieldOrEqualityFilter:
		// addOnConfigKindFieldFilter.AddOnConfigKindFieldOrEqualityFilter is populated
}
```
