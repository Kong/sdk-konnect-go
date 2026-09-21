# BillingChargeUsageBasedCustomer

The customer owning the charge.

By default, only the `id` of the customer is returned. For more details use the
`customer` expand.


## Supported Types

### BillingChargeUsageBasedCustomerBillingCustomer

```go
billingChargeUsageBasedCustomer := components.CreateBillingChargeUsageBasedCustomerBillingChargeUsageBasedCustomerBillingCustomer(components.BillingChargeUsageBasedCustomerBillingCustomer{/* values here */})
```

### CustomerCustomerReference

```go
billingChargeUsageBasedCustomer := components.CreateBillingChargeUsageBasedCustomerCustomerCustomerReference(components.CustomerCustomerReference{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingChargeUsageBasedCustomer.Type {
	case components.BillingChargeUsageBasedCustomerTypeBillingChargeUsageBasedCustomerBillingCustomer:
		// billingChargeUsageBasedCustomer.BillingChargeUsageBasedCustomerBillingCustomer is populated
	case components.BillingChargeUsageBasedCustomerTypeCustomerCustomerReference:
		// billingChargeUsageBasedCustomer.CustomerCustomerReference is populated
}
```
