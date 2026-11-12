# LLMCostPriceSource

Where this price came from.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.LLMCostPriceSourceManual

// Open enum: custom values can be created with a direct type cast
custom := components.LLMCostPriceSource("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `LLMCostPriceSourceManual` | manual                     |
| `LLMCostPriceSourceSystem` | system                     |