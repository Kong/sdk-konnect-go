# RouteWithoutParents


## Supported Types

### RouteJSON

```go
routeWithoutParents := components.CreateRouteWithoutParentsRouteJSON(components.RouteJSON{/* values here */})
```

### RouteExpression

```go
routeWithoutParents := components.CreateRouteWithoutParentsRouteExpression(components.RouteExpression{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch routeWithoutParents.Type {
	case components.RouteWithoutParentsTypeRouteJSON:
		// routeWithoutParents.RouteJSON is populated
	case components.RouteWithoutParentsTypeRouteExpression:
		// routeWithoutParents.RouteExpression is populated
}
```
