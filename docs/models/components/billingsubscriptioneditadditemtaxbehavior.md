# BillingSubscriptionEditAddItemTaxBehavior

Tax behavior.

This enum is used to specify whether tax is included in the price or excluded
from the price. If not specified, the billing profile is used to determine the
tax behavior. If not specified in the billing profile, the provider's default
behavior is used.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionEditAddItemTaxBehaviorInclusive

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionEditAddItemTaxBehavior("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `BillingSubscriptionEditAddItemTaxBehaviorInclusive` | inclusive                                            |
| `BillingSubscriptionEditAddItemTaxBehaviorExclusive` | exclusive                                            |