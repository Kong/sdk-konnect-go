# BillingSubscriptionEditAddItemEntitlementTemplate

The entitlement template granted to subscribers of a plan or addon containing
this rate card. Requires `feature` to be set.


## Supported Types

### BillingRateCardMeteredEntitlement

```go
billingSubscriptionEditAddItemEntitlementTemplate := components.CreateBillingSubscriptionEditAddItemEntitlementTemplateMetered(components.BillingRateCardMeteredEntitlement{/* values here */})
```

### BillingRateCardStaticEntitlement

```go
billingSubscriptionEditAddItemEntitlementTemplate := components.CreateBillingSubscriptionEditAddItemEntitlementTemplateStatic(components.BillingRateCardStaticEntitlement{/* values here */})
```

### BillingRateCardBooleanEntitlement

```go
billingSubscriptionEditAddItemEntitlementTemplate := components.CreateBillingSubscriptionEditAddItemEntitlementTemplateBoolean(components.BillingRateCardBooleanEntitlement{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingSubscriptionEditAddItemEntitlementTemplate.Type {
	case components.BillingSubscriptionEditAddItemEntitlementTemplateTypeMetered:
		// billingSubscriptionEditAddItemEntitlementTemplate.BillingRateCardMeteredEntitlement is populated
	case components.BillingSubscriptionEditAddItemEntitlementTemplateTypeStatic:
		// billingSubscriptionEditAddItemEntitlementTemplate.BillingRateCardStaticEntitlement is populated
	case components.BillingSubscriptionEditAddItemEntitlementTemplateTypeBoolean:
		// billingSubscriptionEditAddItemEntitlementTemplate.BillingRateCardBooleanEntitlement is populated
}
```
