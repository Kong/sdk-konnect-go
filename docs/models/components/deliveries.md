# Deliveries


## Supported Types

### AICostNotificationDeliveryPending

```go
deliveries := components.CreateDeliveriesPending(components.AICostNotificationDeliveryPending{/* values here */})
```

### AICostNotificationDeliveryDelivered

```go
deliveries := components.CreateDeliveriesDelivered(components.AICostNotificationDeliveryDelivered{/* values here */})
```

### AICostNotificationDeliveryFailed

```go
deliveries := components.CreateDeliveriesFailed(components.AICostNotificationDeliveryFailed{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch deliveries.Type {
	case components.DeliveriesTypePending:
		// deliveries.AICostNotificationDeliveryPending is populated
	case components.DeliveriesTypeDelivered:
		// deliveries.AICostNotificationDeliveryDelivered is populated
	case components.DeliveriesTypeFailed:
		// deliveries.AICostNotificationDeliveryFailed is populated
}
```
