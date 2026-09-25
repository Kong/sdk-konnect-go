# BillingEntitlement

An entitlement grants a customer access to a feature.


## Supported Types

### BillingEntitlementMetered

```go
billingEntitlement := components.CreateBillingEntitlementMetered(components.BillingEntitlementMetered{/* values here */})
```

### BillingEntitlementStatic

```go
billingEntitlement := components.CreateBillingEntitlementStatic(components.BillingEntitlementStatic{/* values here */})
```

### BillingEntitlementBoolean

```go
billingEntitlement := components.CreateBillingEntitlementBoolean(components.BillingEntitlementBoolean{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingEntitlement.Type {
	case components.BillingEntitlementTypeMetered:
		// billingEntitlement.BillingEntitlementMetered is populated
	case components.BillingEntitlementTypeStatic:
		// billingEntitlement.BillingEntitlementStatic is populated
	case components.BillingEntitlementTypeBoolean:
		// billingEntitlement.BillingEntitlementBoolean is populated
}
```
