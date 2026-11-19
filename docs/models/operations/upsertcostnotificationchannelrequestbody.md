# UpsertCostNotificationChannelRequestBody


## Supported Types

### AICostEmailNotificationChannelInput

```go
upsertCostNotificationChannelRequestBody := operations.CreateUpsertCostNotificationChannelRequestBodyEmail(components.AICostEmailNotificationChannelInput{/* values here */})
```

### AICostSlackNotificationChannelReplaceInput

```go
upsertCostNotificationChannelRequestBody := operations.CreateUpsertCostNotificationChannelRequestBodySlack(components.AICostSlackNotificationChannelReplaceInput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch upsertCostNotificationChannelRequestBody.Type {
	case operations.UpsertCostNotificationChannelRequestBodyTypeEmail:
		// upsertCostNotificationChannelRequestBody.AICostEmailNotificationChannelInput is populated
	case operations.UpsertCostNotificationChannelRequestBodyTypeSlack:
		// upsertCostNotificationChannelRequestBody.AICostSlackNotificationChannelReplaceInput is populated
}
```
