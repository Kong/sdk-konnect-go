# TransitGatewayStateFieldFilter


## Supported Types

### TransitGatewayStateFieldEqualsFilter

```go
transitGatewayStateFieldFilter := components.CreateTransitGatewayStateFieldFilterTransitGatewayStateFieldEqualsFilter(components.TransitGatewayStateFieldEqualsFilter{/* values here */})
```

### TransitGatewayStateFieldNotEqualsFilter

```go
transitGatewayStateFieldFilter := components.CreateTransitGatewayStateFieldFilterTransitGatewayStateFieldNotEqualsFilter(components.TransitGatewayStateFieldNotEqualsFilter{/* values here */})
```

### TransitGatewayStateFieldOrEqualityFilter

```go
transitGatewayStateFieldFilter := components.CreateTransitGatewayStateFieldFilterTransitGatewayStateFieldOrEqualityFilter(components.TransitGatewayStateFieldOrEqualityFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch transitGatewayStateFieldFilter.Type {
	case components.TransitGatewayStateFieldFilterTypeTransitGatewayStateFieldEqualsFilter:
		// transitGatewayStateFieldFilter.TransitGatewayStateFieldEqualsFilter is populated
	case components.TransitGatewayStateFieldFilterTypeTransitGatewayStateFieldNotEqualsFilter:
		// transitGatewayStateFieldFilter.TransitGatewayStateFieldNotEqualsFilter is populated
	case components.TransitGatewayStateFieldFilterTypeTransitGatewayStateFieldOrEqualityFilter:
		// transitGatewayStateFieldFilter.TransitGatewayStateFieldOrEqualityFilter is populated
}
```
