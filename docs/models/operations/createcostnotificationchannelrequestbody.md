# CreateCostNotificationChannelRequestBody


## Supported Types

### AICostEmailNotificationChannelInput

```go
createCostNotificationChannelRequestBody := operations.CreateCreateCostNotificationChannelRequestBodyEmail(components.AICostEmailNotificationChannelInput{/* values here */})
```

### AICostSlackNotificationChannelInput

```go
createCostNotificationChannelRequestBody := operations.CreateCreateCostNotificationChannelRequestBodySlack(components.AICostSlackNotificationChannelInput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createCostNotificationChannelRequestBody.Type {
	case operations.CreateCostNotificationChannelRequestBodyTypeEmail:
		// createCostNotificationChannelRequestBody.AICostEmailNotificationChannelInput is populated
	case operations.CreateCostNotificationChannelRequestBodyTypeSlack:
		// createCostNotificationChannelRequestBody.AICostSlackNotificationChannelInput is populated
}
```
