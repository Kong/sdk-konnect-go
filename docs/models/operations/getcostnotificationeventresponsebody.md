# GetCostNotificationEventResponseBody

The notification event.


## Supported Types

### AICostNotificationThresholdEvent

```go
getCostNotificationEventResponseBody := operations.CreateGetCostNotificationEventResponseBodyThreshold(components.AICostNotificationThresholdEvent{/* values here */})
```

### AICostNotificationTestEvent

```go
getCostNotificationEventResponseBody := operations.CreateGetCostNotificationEventResponseBodyTest(components.AICostNotificationTestEvent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getCostNotificationEventResponseBody.Type {
	case operations.GetCostNotificationEventResponseBodyTypeThreshold:
		// getCostNotificationEventResponseBody.AICostNotificationThresholdEvent is populated
	case operations.GetCostNotificationEventResponseBodyTypeTest:
		// getCostNotificationEventResponseBody.AICostNotificationTestEvent is populated
}
```
