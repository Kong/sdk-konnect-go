# AICostNotificationTestEvent

A channel test is a test message sent through a channel's real delivery path, with test wording and example figures, so what arrives shows what an alert will look like.

`status` shows where the test stands:

- `pending`: the test message is queued or being sent.
- `delivered`: Slack's webhook answered the test message with a `2xx` status, or the test email was queued for every recipient at Konnect's email service. A later rejection, suppression, or bounce by the email service doesn't change this status.
- `failed`: the test message couldn't be delivered, and `error` says why. A test isn't retried, so its first failed attempt fails it.


## Supported Types

### AICostNotificationTestEventPending

```go
aiCostNotificationTestEvent := components.CreateAICostNotificationTestEventPending(components.AICostNotificationTestEventPending{/* values here */})
```

### AICostNotificationTestEventDelivered

```go
aiCostNotificationTestEvent := components.CreateAICostNotificationTestEventDelivered(components.AICostNotificationTestEventDelivered{/* values here */})
```

### AICostNotificationTestEventFailed

```go
aiCostNotificationTestEvent := components.CreateAICostNotificationTestEventFailed(components.AICostNotificationTestEventFailed{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiCostNotificationTestEvent.Type {
	case components.AICostNotificationTestEventTypePending:
		// aiCostNotificationTestEvent.AICostNotificationTestEventPending is populated
	case components.AICostNotificationTestEventTypeDelivered:
		// aiCostNotificationTestEvent.AICostNotificationTestEventDelivered is populated
	case components.AICostNotificationTestEventTypeFailed:
		// aiCostNotificationTestEvent.AICostNotificationTestEventFailed is populated
}
```
