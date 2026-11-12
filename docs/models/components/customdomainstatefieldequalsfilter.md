# CustomDomainStateFieldEqualsFilter

Filter custom domain state by exact match.


## Supported Types

### CustomDomainState

```go
customDomainStateFieldEqualsFilter := components.CreateCustomDomainStateFieldEqualsFilterCustomDomainState(components.CustomDomainState{/* values here */})
```

### CustomDomainStateFieldEqualsComparison

```go
customDomainStateFieldEqualsFilter := components.CreateCustomDomainStateFieldEqualsFilterCustomDomainStateFieldEqualsComparison(components.CustomDomainStateFieldEqualsComparison{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customDomainStateFieldEqualsFilter.Type {
	case components.CustomDomainStateFieldEqualsFilterTypeCustomDomainState:
		// customDomainStateFieldEqualsFilter.CustomDomainState is populated
	case components.CustomDomainStateFieldEqualsFilterTypeCustomDomainStateFieldEqualsComparison:
		// customDomainStateFieldEqualsFilter.CustomDomainStateFieldEqualsComparison is populated
}
```
