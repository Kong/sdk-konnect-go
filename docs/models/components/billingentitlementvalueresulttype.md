# BillingEntitlementValueResultType

The type of the entitlement.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingEntitlementValueResultTypeMetered

// Open enum: custom values can be created with a direct type cast
custom := components.BillingEntitlementValueResultType("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `BillingEntitlementValueResultTypeMetered` | metered                                    |
| `BillingEntitlementValueResultTypeStatic`  | static                                     |
| `BillingEntitlementValueResultTypeBoolean` | boolean                                    |