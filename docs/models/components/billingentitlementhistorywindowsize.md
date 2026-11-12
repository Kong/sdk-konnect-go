# BillingEntitlementHistoryWindowSize

The meter query granularities the usage history can be grouped into. Sub-hour
windows are too expensive to compute and monthly windows are not supported.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingEntitlementHistoryWindowSizePt1H

// Open enum: custom values can be created with a direct type cast
custom := components.BillingEntitlementHistoryWindowSize("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `BillingEntitlementHistoryWindowSizePt1H` | PT1H                                      |
| `BillingEntitlementHistoryWindowSizeP1D`  | P1D                                       |