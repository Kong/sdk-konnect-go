# BillingChargesExpand

Expands for customer charges.

Values:

- `real_time_usage`: The charge's real-time usage; it sets the `usage` and the
`totals.realtime` fields, and fills the `outstanding` realization's `usage`
with the not-yet-booked remainder of the live read.
- `customer`: The complete customer entity of the charge.
- `feature`: The complete feature entity of the charge.
- `subscription`: The complete subscription of the charge, when present.
- `realization.invoice`: The invoice header of each realization (the invoice
entity without its `lines` and `customer` snapshot), in place of the ID
reference.
- `realization.totals`: The `totals` of each realization run, including credit
allocations.
- `realization.detailed_lines`: The `detailed_lines` of each realization run.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingChargesExpandRealTimeUsage

// Open enum: custom values can be created with a direct type cast
custom := components.BillingChargesExpand("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `BillingChargesExpandRealTimeUsage`            | real_time_usage                                |
| `BillingChargesExpandCustomer`                 | customer                                       |
| `BillingChargesExpandFeature`                  | feature                                        |
| `BillingChargesExpandSubscription`             | subscription                                   |
| `BillingChargesExpandRealizationInvoice`       | realization.invoice                            |
| `BillingChargesExpandRealizationTotals`        | realization.totals                             |
| `BillingChargesExpandRealizationDetailedLines` | realization.detailed_lines                     |