# CreateAIGatewayModelRequest

Configuration for an AI Gateway model.


## Supported Types

### AIGatewayModelAPI

```go
createAIGatewayModelRequest := components.CreateCreateAIGatewayModelRequestAPI(components.AIGatewayModelAPI{/* values here */})
```

### AIGatewayModelModel

```go
createAIGatewayModelRequest := components.CreateCreateAIGatewayModelRequestModel(components.AIGatewayModelModel{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createAIGatewayModelRequest.Type {
	case components.CreateAIGatewayModelRequestTypeAPI:
		// createAIGatewayModelRequest.AIGatewayModelAPI is populated
	case components.CreateAIGatewayModelRequestTypeModel:
		// createAIGatewayModelRequest.AIGatewayModelModel is populated
}
```
