# UpdateInvoiceLine

A top-level line item on an invoice.

Each line represents a single charge, typically associated with a rate card from
a subscription. Detailed (child) lines are nested under `detailed_lines` when
present.


## Supported Types

### UpdateInvoiceStandardLine

```go
updateInvoiceLine := components.CreateUpdateInvoiceLineStandardLine(components.UpdateInvoiceStandardLine{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateInvoiceLine.Type {
	case components.UpdateInvoiceLineTypeStandardLine:
		// updateInvoiceLine.UpdateInvoiceStandardLine is populated
}
```
