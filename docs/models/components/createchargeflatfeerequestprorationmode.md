# CreateChargeFlatFeeRequestProrationMode

The proration mode of the rate card.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreateChargeFlatFeeRequestProrationModeNoProration

// Open enum: custom values can be created with a direct type cast
custom := components.CreateChargeFlatFeeRequestProrationMode("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `CreateChargeFlatFeeRequestProrationModeNoProration`   | no_proration                                           |
| `CreateChargeFlatFeeRequestProrationModeProratePrices` | prorate_prices                                         |