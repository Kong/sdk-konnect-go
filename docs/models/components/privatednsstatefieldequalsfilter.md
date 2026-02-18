# PrivateDNSStateFieldEqualsFilter

Filter Private DNS state by exact match.


## Supported Types

### PrivateDNSState

```go
privateDNSStateFieldEqualsFilter := components.CreatePrivateDNSStateFieldEqualsFilterPrivateDNSState(components.PrivateDNSState{/* values here */})
```

### PrivateDNSStateFieldEqualsComparison

```go
privateDNSStateFieldEqualsFilter := components.CreatePrivateDNSStateFieldEqualsFilterPrivateDNSStateFieldEqualsComparison(components.PrivateDNSStateFieldEqualsComparison{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch privateDNSStateFieldEqualsFilter.Type {
	case components.PrivateDNSStateFieldEqualsFilterTypePrivateDNSState:
		// privateDNSStateFieldEqualsFilter.PrivateDNSState is populated
	case components.PrivateDNSStateFieldEqualsFilterTypePrivateDNSStateFieldEqualsComparison:
		// privateDNSStateFieldEqualsFilter.PrivateDNSStateFieldEqualsComparison is populated
}
```
