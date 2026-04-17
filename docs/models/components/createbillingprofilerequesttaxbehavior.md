# CreateBillingProfileRequestTaxBehavior

Tax behavior.

If not specified the billing profile is used to determine the tax behavior.
If not specified in the billing profile, the provider's default behavior is used.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreateBillingProfileRequestTaxBehaviorInclusive

// Open enum: custom values can be created with a direct type cast
custom := components.CreateBillingProfileRequestTaxBehavior("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `CreateBillingProfileRequestTaxBehaviorInclusive` | inclusive                                         |
| `CreateBillingProfileRequestTaxBehaviorExclusive` | exclusive                                         |