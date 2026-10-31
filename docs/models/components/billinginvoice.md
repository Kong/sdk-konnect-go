# BillingInvoice

An invoice issued to a customer.

The `type` field determines the concrete variant:

- `standard`: a standard invoice for charges owed.


## Supported Types

### BillingInvoiceStandard

```go
billingInvoice := components.CreateBillingInvoiceStandard(components.BillingInvoiceStandard{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingInvoice.Type {
	case components.BillingInvoiceTypeStandard:
		// billingInvoice.BillingInvoiceStandard is populated
}
```
