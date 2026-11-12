# BillingInvoiceLine

A top-level line item on an invoice.

Each line represents a single charge, typically associated with a rate card from
a subscription. Detailed (child) lines are nested under `detailed_lines` when
present.


## Supported Types

### BillingInvoiceStandardLine

```go
billingInvoiceLine := components.CreateBillingInvoiceLineStandardLine(components.BillingInvoiceStandardLine{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingInvoiceLine.Type {
	case components.BillingInvoiceLineTypeStandardLine:
		// billingInvoiceLine.BillingInvoiceStandardLine is populated
}
```
