# EntitlementTemplate

The entitlement template granted to subscribers of a plan or addon containing
this rate card. Requires `feature` to be set.


## Supported Types

### BillingRateCardMeteredEntitlement

```go
entitlementTemplate := components.CreateEntitlementTemplateMetered(components.BillingRateCardMeteredEntitlement{/* values here */})
```

### BillingRateCardStaticEntitlement

```go
entitlementTemplate := components.CreateEntitlementTemplateStatic(components.BillingRateCardStaticEntitlement{/* values here */})
```

### BillingRateCardBooleanEntitlement

```go
entitlementTemplate := components.CreateEntitlementTemplateBoolean(components.BillingRateCardBooleanEntitlement{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch entitlementTemplate.Type {
	case components.EntitlementTemplateTypeMetered:
		// entitlementTemplate.BillingRateCardMeteredEntitlement is populated
	case components.EntitlementTemplateTypeStatic:
		// entitlementTemplate.BillingRateCardStaticEntitlement is populated
	case components.EntitlementTemplateTypeBoolean:
		// entitlementTemplate.BillingRateCardBooleanEntitlement is populated
}
```
