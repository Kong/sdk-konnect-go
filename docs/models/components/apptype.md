# AppType

The app type that the tax code is associated with.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AppTypeSandbox

// Open enum: custom values can be created with a direct type cast
custom := components.AppType("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `AppTypeSandbox`           | sandbox                    |
| `AppTypeStripe`            | stripe                     |
| `AppTypeExternalInvoicing` | external_invoicing         |