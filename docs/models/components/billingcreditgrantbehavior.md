# BillingCreditGrantBehavior

Tax behavior applied to the invoice line item.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingCreditGrantBehaviorInclusive

// Open enum: custom values can be created with a direct type cast
custom := components.BillingCreditGrantBehavior("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `BillingCreditGrantBehaviorInclusive` | inclusive                             |
| `BillingCreditGrantBehaviorExclusive` | exclusive                             |