# BillingEntitlementAccessResultType

The type of the entitlement.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingEntitlementAccessResultTypeMetered

// Open enum: custom values can be created with a direct type cast
custom := components.BillingEntitlementAccessResultType("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `BillingEntitlementAccessResultTypeMetered` | metered                                     |
| `BillingEntitlementAccessResultTypeStatic`  | static                                      |
| `BillingEntitlementAccessResultTypeBoolean` | boolean                                     |