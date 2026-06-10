# BillingPlanStatus

The status of the plan. Computed based on the effective start and end dates:

- `draft`: `effective_from` is not set.
- `scheduled`: `now < effective_from`.
- `active`: `effective_from <= now` and (`effective_to` is not set or
`now < effective_to`).
- `archived`: `effective_to <= now`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/metering"
)

value := metering.BillingPlanStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := metering.BillingPlanStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `BillingPlanStatusDraft`     | draft                        |
| `BillingPlanStatusActive`    | active                       |
| `BillingPlanStatusArchived`  | archived                     |
| `BillingPlanStatusScheduled` | scheduled                    |