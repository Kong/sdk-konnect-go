# BillingChargeFlatFeeCustomer

The customer owning the charge.

By default, only the `id` of the customer is returned. For more details use the
`customer` expand.


## Supported Types

### CustomerBillingCustomer

```go
billingChargeFlatFeeCustomer := components.CreateBillingChargeFlatFeeCustomerCustomerBillingCustomer(components.CustomerBillingCustomer{/* values here */})
```

### CustomerReference

```go
billingChargeFlatFeeCustomer := components.CreateBillingChargeFlatFeeCustomerCustomerReference(components.CustomerReference{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingChargeFlatFeeCustomer.Type {
	case components.BillingChargeFlatFeeCustomerTypeCustomerBillingCustomer:
		// billingChargeFlatFeeCustomer.CustomerBillingCustomer is populated
	case components.BillingChargeFlatFeeCustomerTypeCustomerReference:
		// billingChargeFlatFeeCustomer.CustomerReference is populated
}
```
