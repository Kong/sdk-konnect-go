# SubscriptionAddonRateCardEntitlementTemplate

The entitlement template granted to subscribers of a plan or addon containing
this rate card. Requires `feature` to be set.


## Supported Types

### BillingRateCardMeteredEntitlement

```go
subscriptionAddonRateCardEntitlementTemplate := components.CreateSubscriptionAddonRateCardEntitlementTemplateMetered(components.BillingRateCardMeteredEntitlement{/* values here */})
```

### BillingRateCardStaticEntitlement

```go
subscriptionAddonRateCardEntitlementTemplate := components.CreateSubscriptionAddonRateCardEntitlementTemplateStatic(components.BillingRateCardStaticEntitlement{/* values here */})
```

### BillingRateCardBooleanEntitlement

```go
subscriptionAddonRateCardEntitlementTemplate := components.CreateSubscriptionAddonRateCardEntitlementTemplateBoolean(components.BillingRateCardBooleanEntitlement{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch subscriptionAddonRateCardEntitlementTemplate.Type {
	case components.SubscriptionAddonRateCardEntitlementTemplateTypeMetered:
		// subscriptionAddonRateCardEntitlementTemplate.BillingRateCardMeteredEntitlement is populated
	case components.SubscriptionAddonRateCardEntitlementTemplateTypeStatic:
		// subscriptionAddonRateCardEntitlementTemplate.BillingRateCardStaticEntitlement is populated
	case components.SubscriptionAddonRateCardEntitlementTemplateTypeBoolean:
		// subscriptionAddonRateCardEntitlementTemplate.BillingRateCardBooleanEntitlement is populated
}
```
