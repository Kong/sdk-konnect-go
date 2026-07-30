# UpdateAIGatewayModelRequest

**Pre-release Feature**
This feature is currently in beta and is subject to change.

Configuration for an AI Gateway model.


## Supported Types

### AIGatewayModelAPI

```go
updateAIGatewayModelRequest := components.CreateUpdateAIGatewayModelRequestAPI(components.AIGatewayModelAPI{/* values here */})
```

### AIGatewayModelModel

```go
updateAIGatewayModelRequest := components.CreateUpdateAIGatewayModelRequestModel(components.AIGatewayModelModel{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateAIGatewayModelRequest.Type {
	case components.UpdateAIGatewayModelRequestTypeAPI:
		// updateAIGatewayModelRequest.AIGatewayModelAPI is populated
	case components.UpdateAIGatewayModelRequestTypeModel:
		// updateAIGatewayModelRequest.AIGatewayModelModel is populated
}
```
