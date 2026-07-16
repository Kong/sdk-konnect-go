# CreditGrantLifecycleStatus

Current lifecycle status of the grant.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreditGrantLifecycleStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.CreditGrantLifecycleStatus("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `CreditGrantLifecycleStatusPending` | pending                             |
| `CreditGrantLifecycleStatusActive`  | active                              |
| `CreditGrantLifecycleStatusExpired` | expired                             |
| `CreditGrantLifecycleStatusVoided`  | voided                              |