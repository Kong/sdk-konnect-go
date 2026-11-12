# BillingChargeFlatFeeCustomer

The customer owning the charge.

By default, only the `id` of the customer is returned. For more details use the
`customer` expand.


## Supported Types

### BillingChargeFlatFeeCustomerBillingCustomer

```go
billingChargeFlatFeeCustomer := components.CreateBillingChargeFlatFeeCustomerBillingChargeFlatFeeCustomerBillingCustomer(components.BillingChargeFlatFeeCustomerBillingCustomer{/* values here */})
```

### CustomerReference

```go
billingChargeFlatFeeCustomer := components.CreateBillingChargeFlatFeeCustomerCustomerReference(components.CustomerReference{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingChargeFlatFeeCustomer.Type {
	case components.BillingChargeFlatFeeCustomerTypeBillingChargeFlatFeeCustomerBillingCustomer:
		// billingChargeFlatFeeCustomer.BillingChargeFlatFeeCustomerBillingCustomer is populated
	case components.BillingChargeFlatFeeCustomerTypeCustomerReference:
		// billingChargeFlatFeeCustomer.CustomerReference is populated
}
```
