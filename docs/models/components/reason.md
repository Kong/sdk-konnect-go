# Reason

The reason this discount was applied.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ReasonMaximumSpend

// Open enum: custom values can be created with a direct type cast
custom := components.Reason("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `ReasonMaximumSpend`       | maximum_spend              |
| `ReasonRatecardPercentage` | ratecard_percentage        |
| `ReasonRatecardUsage`      | ratecard_usage             |