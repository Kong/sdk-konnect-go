# BillingEntitlementAccessCheckResultType

The type of the entitlement.

If not provided, the feature has no entitlement defined (has access is always
false in this case)

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingEntitlementAccessCheckResultTypeMetered

// Open enum: custom values can be created with a direct type cast
custom := components.BillingEntitlementAccessCheckResultType("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `BillingEntitlementAccessCheckResultTypeMetered` | metered                                          |
| `BillingEntitlementAccessCheckResultTypeStatic`  | static                                           |
| `BillingEntitlementAccessCheckResultTypeBoolean` | boolean                                          |