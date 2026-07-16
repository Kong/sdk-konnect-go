# AIGatewayModel

**Pre-release Feature**
This feature is currently in beta and is subject to change.

Configuration for an AI Gateway model.


## Supported Types

### AIGatewayModelAIGatewayModelAPI

```go
aiGatewayModel := components.CreateAIGatewayModelAPI(components.AIGatewayModelAIGatewayModelAPI{/* values here */})
```

### AIGatewayModelAIGatewayModelModel

```go
aiGatewayModel := components.CreateAIGatewayModelModel(components.AIGatewayModelAIGatewayModelModel{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModel.Type {
	case components.AIGatewayModelTypeAPI:
		// aiGatewayModel.AIGatewayModelAIGatewayModelAPI is populated
	case components.AIGatewayModelTypeModel:
		// aiGatewayModel.AIGatewayModelAIGatewayModelModel is populated
}
```
