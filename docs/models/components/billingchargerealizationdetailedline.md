# BillingChargeRealizationDetailedLine

A detailed (child) line type of a charge realization run.

This is distinct from an invoice's own detailed lines: it represents the
rated/priced breakdown produced by the realization run itself, before that
breakdown is (or is not yet) reflected on an invoice line. Credit-then-invoice
runs include credit allocations in these lines, while credits-only runs keep the
gross rated detail.


## Supported Types

### BillingChargeRealizationDetailedLineFlatFee

```go
billingChargeRealizationDetailedLine := components.CreateBillingChargeRealizationDetailedLineFlatFee(components.BillingChargeRealizationDetailedLineFlatFee{/* values here */})
```

### BillingChargeRealizationDetailedLineUsageBased

```go
billingChargeRealizationDetailedLine := components.CreateBillingChargeRealizationDetailedLineUsageBased(components.BillingChargeRealizationDetailedLineUsageBased{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingChargeRealizationDetailedLine.Type {
	case components.BillingChargeRealizationDetailedLineTypeFlatFee:
		// billingChargeRealizationDetailedLine.BillingChargeRealizationDetailedLineFlatFee is populated
	case components.BillingChargeRealizationDetailedLineTypeUsageBased:
		// billingChargeRealizationDetailedLine.BillingChargeRealizationDetailedLineUsageBased is populated
}
```
