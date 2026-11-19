# AICostNotificationChannel

A notification channel is a destination for budget alerts: a list of email addresses or a Slack incoming webhook. A budget's notification rule sends its alerts to the channels it lists, and one channel can serve many rules.

`type` tells the two kinds apart. An email channel carries its settings in `email`, and a Slack channel carries them in `slack`.


## Supported Types

### AICostEmailNotificationChannel

```go
aiCostNotificationChannel := components.CreateAICostNotificationChannelEmail(components.AICostEmailNotificationChannel{/* values here */})
```

### AICostSlackNotificationChannel

```go
aiCostNotificationChannel := components.CreateAICostNotificationChannelSlack(components.AICostSlackNotificationChannel{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiCostNotificationChannel.Type {
	case components.AICostNotificationChannelTypeEmail:
		// aiCostNotificationChannel.AICostEmailNotificationChannel is populated
	case components.AICostNotificationChannelTypeSlack:
		// aiCostNotificationChannel.AICostSlackNotificationChannel is populated
}
```
