# Deliveries


## Supported Types

### AICostNotificationDeliveryPending

```go
deliveries := operations.CreateDeliveriesPending(components.AICostNotificationDeliveryPending{/* values here */})
```

### AICostNotificationDeliveryDelivered

```go
deliveries := operations.CreateDeliveriesDelivered(components.AICostNotificationDeliveryDelivered{/* values here */})
```

### AICostNotificationDeliveryFailed

```go
deliveries := operations.CreateDeliveriesFailed(components.AICostNotificationDeliveryFailed{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch deliveries.Type {
	case operations.DeliveriesTypePending:
		// deliveries.AICostNotificationDeliveryPending is populated
	case operations.DeliveriesTypeDelivered:
		// deliveries.AICostNotificationDeliveryDelivered is populated
	case operations.DeliveriesTypeFailed:
		// deliveries.AICostNotificationDeliveryFailed is populated
}
```
