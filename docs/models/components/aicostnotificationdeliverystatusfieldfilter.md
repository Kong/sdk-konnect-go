# AICostNotificationDeliveryStatusFieldFilter

Filters a field by its value with `eq`, `neq`, or `oeq`. Accepts `pending`, `delivered`, and `failed`. A bare value means `eq`.


## Supported Types

### AICostNotificationDeliveryStatus

```go
aiCostNotificationDeliveryStatusFieldFilter := components.CreateAICostNotificationDeliveryStatusFieldFilterAICostNotificationDeliveryStatus(components.AICostNotificationDeliveryStatus{/* values here */})
```

### AICostNotificationDeliveryStatusFieldEqualsFilter

```go
aiCostNotificationDeliveryStatusFieldFilter := components.CreateAICostNotificationDeliveryStatusFieldFilterAICostNotificationDeliveryStatusFieldEqualsFilter(components.AICostNotificationDeliveryStatusFieldEqualsFilter{/* values here */})
```

### AICostNotificationDeliveryStatusFieldNEQFilter

```go
aiCostNotificationDeliveryStatusFieldFilter := components.CreateAICostNotificationDeliveryStatusFieldFilterAICostNotificationDeliveryStatusFieldNEQFilter(components.AICostNotificationDeliveryStatusFieldNEQFilter{/* values here */})
```

### StringFieldOEQFilter

```go
aiCostNotificationDeliveryStatusFieldFilter := components.CreateAICostNotificationDeliveryStatusFieldFilterStringFieldOEQFilter(components.StringFieldOEQFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiCostNotificationDeliveryStatusFieldFilter.Type {
	case components.AICostNotificationDeliveryStatusFieldFilterTypeAICostNotificationDeliveryStatus:
		// aiCostNotificationDeliveryStatusFieldFilter.AICostNotificationDeliveryStatus is populated
	case components.AICostNotificationDeliveryStatusFieldFilterTypeAICostNotificationDeliveryStatusFieldEqualsFilter:
		// aiCostNotificationDeliveryStatusFieldFilter.AICostNotificationDeliveryStatusFieldEqualsFilter is populated
	case components.AICostNotificationDeliveryStatusFieldFilterTypeAICostNotificationDeliveryStatusFieldNEQFilter:
		// aiCostNotificationDeliveryStatusFieldFilter.AICostNotificationDeliveryStatusFieldNEQFilter is populated
	case components.AICostNotificationDeliveryStatusFieldFilterTypeStringFieldOEQFilter:
		// aiCostNotificationDeliveryStatusFieldFilter.StringFieldOEQFilter is populated
}
```
