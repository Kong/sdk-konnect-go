# PrivateDNSStateFieldFilter


## Supported Types

### PrivateDNSStateFieldEqualsFilter

```go
privateDNSStateFieldFilter := components.CreatePrivateDNSStateFieldFilterPrivateDNSStateFieldEqualsFilter(components.PrivateDNSStateFieldEqualsFilter{/* values here */})
```

### PrivateDNSStateFieldNotEqualsFilter

```go
privateDNSStateFieldFilter := components.CreatePrivateDNSStateFieldFilterPrivateDNSStateFieldNotEqualsFilter(components.PrivateDNSStateFieldNotEqualsFilter{/* values here */})
```

### PrivateDNSStateFieldOrEqualityFilter

```go
privateDNSStateFieldFilter := components.CreatePrivateDNSStateFieldFilterPrivateDNSStateFieldOrEqualityFilter(components.PrivateDNSStateFieldOrEqualityFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch privateDNSStateFieldFilter.Type {
	case components.PrivateDNSStateFieldFilterTypePrivateDNSStateFieldEqualsFilter:
		// privateDNSStateFieldFilter.PrivateDNSStateFieldEqualsFilter is populated
	case components.PrivateDNSStateFieldFilterTypePrivateDNSStateFieldNotEqualsFilter:
		// privateDNSStateFieldFilter.PrivateDNSStateFieldNotEqualsFilter is populated
	case components.PrivateDNSStateFieldFilterTypePrivateDNSStateFieldOrEqualityFilter:
		// privateDNSStateFieldFilter.PrivateDNSStateFieldOrEqualityFilter is populated
}
```
