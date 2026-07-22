# Model

Configuration for routing requests to a specific model.



## Supported Types

### AIGatewayModelRouteConfigBody

```go
model := components.CreateModelAIGatewayModelRouteConfigBody(components.AIGatewayModelRouteConfigBody{/* values here */})
```

### AIGatewayModelRouteConfigHeaders

```go
model := components.CreateModelAIGatewayModelRouteConfigHeaders(components.AIGatewayModelRouteConfigHeaders{/* values here */})
```

### AIGatewayModelRouteConfigPathAliases

```go
model := components.CreateModelAIGatewayModelRouteConfigPathAliases(components.AIGatewayModelRouteConfigPathAliases{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch model.Type {
	case components.ModelTypeAIGatewayModelRouteConfigBody:
		// model.AIGatewayModelRouteConfigBody is populated
	case components.ModelTypeAIGatewayModelRouteConfigHeaders:
		// model.AIGatewayModelRouteConfigHeaders is populated
	case components.ModelTypeAIGatewayModelRouteConfigPathAliases:
		// model.AIGatewayModelRouteConfigPathAliases is populated
}
```
