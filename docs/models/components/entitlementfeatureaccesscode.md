# EntitlementFeatureAccessCode

Machine-readable error code.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.EntitlementFeatureAccessCodeUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.EntitlementFeatureAccessCode("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `EntitlementFeatureAccessCodeUnknown`            | unknown                                          |
| `EntitlementFeatureAccessCodeUsageLimitReached`  | usage_limit_reached                              |
| `EntitlementFeatureAccessCodeFeatureUnavailable` | feature_unavailable                              |
| `EntitlementFeatureAccessCodeFeatureNotFound`    | feature_not_found                                |
| `EntitlementFeatureAccessCodeNoCreditAvailable`  | no_credit_available                              |