# CustomDomainStateFieldFilter


## Supported Types

### CustomDomainStateFieldEqualsFilter

```go
customDomainStateFieldFilter := components.CreateCustomDomainStateFieldFilterCustomDomainStateFieldEqualsFilter(components.CustomDomainStateFieldEqualsFilter{/* values here */})
```

### CustomDomainStateFieldNotEqualsFilter

```go
customDomainStateFieldFilter := components.CreateCustomDomainStateFieldFilterCustomDomainStateFieldNotEqualsFilter(components.CustomDomainStateFieldNotEqualsFilter{/* values here */})
```

### CustomDomainStateFieldOrEqualityFilter

```go
customDomainStateFieldFilter := components.CreateCustomDomainStateFieldFilterCustomDomainStateFieldOrEqualityFilter(components.CustomDomainStateFieldOrEqualityFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customDomainStateFieldFilter.Type {
	case components.CustomDomainStateFieldFilterTypeCustomDomainStateFieldEqualsFilter:
		// customDomainStateFieldFilter.CustomDomainStateFieldEqualsFilter is populated
	case components.CustomDomainStateFieldFilterTypeCustomDomainStateFieldNotEqualsFilter:
		// customDomainStateFieldFilter.CustomDomainStateFieldNotEqualsFilter is populated
	case components.CustomDomainStateFieldFilterTypeCustomDomainStateFieldOrEqualityFilter:
		// customDomainStateFieldFilter.CustomDomainStateFieldOrEqualityFilter is populated
}
```
