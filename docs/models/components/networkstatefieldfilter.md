# NetworkStateFieldFilter


## Supported Types

### NetworkStateFieldEqualsFilter

```go
networkStateFieldFilter := components.CreateNetworkStateFieldFilterNetworkStateFieldEqualsFilter(components.NetworkStateFieldEqualsFilter{/* values here */})
```

### NetworkStateFieldNotEqualsFilter

```go
networkStateFieldFilter := components.CreateNetworkStateFieldFilterNetworkStateFieldNotEqualsFilter(components.NetworkStateFieldNotEqualsFilter{/* values here */})
```

### NetworkStateFieldOrEqualityFilter

```go
networkStateFieldFilter := components.CreateNetworkStateFieldFilterNetworkStateFieldOrEqualityFilter(components.NetworkStateFieldOrEqualityFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch networkStateFieldFilter.Type {
	case components.NetworkStateFieldFilterTypeNetworkStateFieldEqualsFilter:
		// networkStateFieldFilter.NetworkStateFieldEqualsFilter is populated
	case components.NetworkStateFieldFilterTypeNetworkStateFieldNotEqualsFilter:
		// networkStateFieldFilter.NetworkStateFieldNotEqualsFilter is populated
	case components.NetworkStateFieldFilterTypeNetworkStateFieldOrEqualityFilter:
		// networkStateFieldFilter.NetworkStateFieldOrEqualityFilter is populated
}
```
