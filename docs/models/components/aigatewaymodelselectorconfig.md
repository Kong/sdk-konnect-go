# AIGatewayModelSelectorConfig

Configuration for overriding routing to this model using a selector.
When not set, a default model selector will be created using the model's name and format.



## Supported Types

### AIGatewayModelSelectorConfigBody

```go
aiGatewayModelSelectorConfig := components.CreateAIGatewayModelSelectorConfigAIGatewayModelSelectorConfigBody(components.AIGatewayModelSelectorConfigBody{/* values here */})
```

### AIGatewayModelSelectorConfigHeaders

```go
aiGatewayModelSelectorConfig := components.CreateAIGatewayModelSelectorConfigAIGatewayModelSelectorConfigHeaders(components.AIGatewayModelSelectorConfigHeaders{/* values here */})
```

### AIGatewayModelSelectorConfigPath

```go
aiGatewayModelSelectorConfig := components.CreateAIGatewayModelSelectorConfigAIGatewayModelSelectorConfigPath(components.AIGatewayModelSelectorConfigPath{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelSelectorConfig.Type {
	case components.AIGatewayModelSelectorConfigTypeAIGatewayModelSelectorConfigBody:
		// aiGatewayModelSelectorConfig.AIGatewayModelSelectorConfigBody is populated
	case components.AIGatewayModelSelectorConfigTypeAIGatewayModelSelectorConfigHeaders:
		// aiGatewayModelSelectorConfig.AIGatewayModelSelectorConfigHeaders is populated
	case components.AIGatewayModelSelectorConfigTypeAIGatewayModelSelectorConfigPath:
		// aiGatewayModelSelectorConfig.AIGatewayModelSelectorConfigPath is populated
}
```
