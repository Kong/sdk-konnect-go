# NetworkStateFieldEqualsFilter

Filter a network state by exact match.


## Supported Types

### NetworkState

```go
networkStateFieldEqualsFilter := components.CreateNetworkStateFieldEqualsFilterNetworkState(components.NetworkState{/* values here */})
```

### NetworkStateFieldEqualsComparison

```go
networkStateFieldEqualsFilter := components.CreateNetworkStateFieldEqualsFilterNetworkStateFieldEqualsComparison(components.NetworkStateFieldEqualsComparison{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch networkStateFieldEqualsFilter.Type {
	case components.NetworkStateFieldEqualsFilterTypeNetworkState:
		// networkStateFieldEqualsFilter.NetworkState is populated
	case components.NetworkStateFieldEqualsFilterTypeNetworkStateFieldEqualsComparison:
		// networkStateFieldEqualsFilter.NetworkStateFieldEqualsComparison is populated
}
```
