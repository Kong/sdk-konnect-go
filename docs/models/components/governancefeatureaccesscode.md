# GovernanceFeatureAccessCode

Machine-readable error code.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.GovernanceFeatureAccessCodeUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.GovernanceFeatureAccessCode("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `GovernanceFeatureAccessCodeUnknown`            | unknown                                         |
| `GovernanceFeatureAccessCodeUsageLimitReached`  | usage_limit_reached                             |
| `GovernanceFeatureAccessCodeFeatureUnavailable` | feature_unavailable                             |
| `GovernanceFeatureAccessCodeFeatureNotFound`    | feature_not_found                               |
| `GovernanceFeatureAccessCodeNoCreditAvailable`  | no_credit_available                             |