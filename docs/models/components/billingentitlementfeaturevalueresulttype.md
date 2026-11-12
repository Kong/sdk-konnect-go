# BillingEntitlementFeatureValueResultType

The type of the entitlement.

If not provided, the feature has no entitlement defined (has access is always
false in this case)

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingEntitlementFeatureValueResultTypeMetered

// Open enum: custom values can be created with a direct type cast
custom := components.BillingEntitlementFeatureValueResultType("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `BillingEntitlementFeatureValueResultTypeMetered` | metered                                           |
| `BillingEntitlementFeatureValueResultTypeStatic`  | static                                            |
| `BillingEntitlementFeatureValueResultTypeBoolean` | boolean                                           |