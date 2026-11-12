# ProrationMode

The proration mode of the rate card.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ProrationModeNoProration

// Open enum: custom values can be created with a direct type cast
custom := components.ProrationMode("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `ProrationModeNoProration`   | no_proration                 |
| `ProrationModeProratePrices` | prorate_prices               |