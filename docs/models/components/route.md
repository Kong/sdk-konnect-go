# Route


## Supported Types

### RouteJSON

```go
route := components.CreateRouteRouteJSON(components.RouteJSON{/* values here */})
```

### RouteExpression

```go
route := components.CreateRouteRouteExpression(components.RouteExpression{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch route.Type {
	case components.RouteTypeRouteJSON:
		// route.RouteJSON is populated
	case components.RouteTypeRouteExpression:
		// route.RouteExpression is populated
}
```
