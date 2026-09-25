# BillingEntitlementMeasureUsageFromPreset

Preset for the time from which usage is measured.

- `current_period_start`: the start of the current usage period.
- `now`: the entitlement creation time.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingEntitlementMeasureUsageFromPresetCurrentPeriodStart

// Open enum: custom values can be created with a direct type cast
custom := components.BillingEntitlementMeasureUsageFromPreset("custom_value")
```


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `BillingEntitlementMeasureUsageFromPresetCurrentPeriodStart` | current_period_start                                         |
| `BillingEntitlementMeasureUsageFromPresetNow`                | now                                                          |