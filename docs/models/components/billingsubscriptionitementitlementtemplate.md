# BillingSubscriptionItemEntitlementTemplate

The entitlement template granted to subscribers of a plan or addon containing
this rate card. Requires `feature` to be set.


## Supported Types

### BillingRateCardMeteredEntitlement

```go
billingSubscriptionItemEntitlementTemplate := components.CreateBillingSubscriptionItemEntitlementTemplateMetered(components.BillingRateCardMeteredEntitlement{/* values here */})
```

### BillingRateCardStaticEntitlement

```go
billingSubscriptionItemEntitlementTemplate := components.CreateBillingSubscriptionItemEntitlementTemplateStatic(components.BillingRateCardStaticEntitlement{/* values here */})
```

### BillingRateCardBooleanEntitlement

```go
billingSubscriptionItemEntitlementTemplate := components.CreateBillingSubscriptionItemEntitlementTemplateBoolean(components.BillingRateCardBooleanEntitlement{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingSubscriptionItemEntitlementTemplate.Type {
	case components.BillingSubscriptionItemEntitlementTemplateTypeMetered:
		// billingSubscriptionItemEntitlementTemplate.BillingRateCardMeteredEntitlement is populated
	case components.BillingSubscriptionItemEntitlementTemplateTypeStatic:
		// billingSubscriptionItemEntitlementTemplate.BillingRateCardStaticEntitlement is populated
	case components.BillingSubscriptionItemEntitlementTemplateTypeBoolean:
		// billingSubscriptionItemEntitlementTemplate.BillingRateCardBooleanEntitlement is populated
}
```
