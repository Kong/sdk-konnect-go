# BillingChargeRealizationInvoice

The reference of the invoice related to the realization.


## Supported Types

### ChargeRealizationInvoice

```go
billingChargeRealizationInvoice := components.CreateBillingChargeRealizationInvoiceChargeRealizationInvoice(components.ChargeRealizationInvoice{/* values here */})
```

### ChargeRealizationInvoiceReference

```go
billingChargeRealizationInvoice := components.CreateBillingChargeRealizationInvoiceChargeRealizationInvoiceReference(components.ChargeRealizationInvoiceReference{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingChargeRealizationInvoice.Type {
	case components.BillingChargeRealizationInvoiceTypeChargeRealizationInvoice:
		// billingChargeRealizationInvoice.ChargeRealizationInvoice is populated
	case components.BillingChargeRealizationInvoiceTypeChargeRealizationInvoiceReference:
		// billingChargeRealizationInvoice.ChargeRealizationInvoiceReference is populated
}
```
