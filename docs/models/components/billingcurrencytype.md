# BillingCurrencyType

Currency type for custom currencies. It should be a unique code but not
conflicting with any existing standard currency codes.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingCurrencyTypeFiat

// Open enum: custom values can be created with a direct type cast
custom := components.BillingCurrencyType("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `BillingCurrencyTypeFiat`   | fiat                        |
| `BillingCurrencyTypeCustom` | custom                      |