# TransitGatewayStateFieldEqualsFilter

Filter transit-gateway state by exact match.


## Supported Types

### TransitGatewayState

```go
transitGatewayStateFieldEqualsFilter := components.CreateTransitGatewayStateFieldEqualsFilterTransitGatewayState(components.TransitGatewayState{/* values here */})
```

### TransitGatewayStateFieldEqualsComparison

```go
transitGatewayStateFieldEqualsFilter := components.CreateTransitGatewayStateFieldEqualsFilterTransitGatewayStateFieldEqualsComparison(components.TransitGatewayStateFieldEqualsComparison{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch transitGatewayStateFieldEqualsFilter.Type {
	case components.TransitGatewayStateFieldEqualsFilterTypeTransitGatewayState:
		// transitGatewayStateFieldEqualsFilter.TransitGatewayState is populated
	case components.TransitGatewayStateFieldEqualsFilterTypeTransitGatewayStateFieldEqualsComparison:
		// transitGatewayStateFieldEqualsFilter.TransitGatewayStateFieldEqualsComparison is populated
}
```
