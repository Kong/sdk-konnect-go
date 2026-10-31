# ListCreditGrantsParamsFilterCreditGrantLifecycleStatus

Filter credit grants by status.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ListCreditGrantsParamsFilterCreditGrantLifecycleStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.ListCreditGrantsParamsFilterCreditGrantLifecycleStatus("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `ListCreditGrantsParamsFilterCreditGrantLifecycleStatusPending` | pending                                                         |
| `ListCreditGrantsParamsFilterCreditGrantLifecycleStatusActive`  | active                                                          |
| `ListCreditGrantsParamsFilterCreditGrantLifecycleStatusExpired` | expired                                                         |
| `ListCreditGrantsParamsFilterCreditGrantLifecycleStatusVoided`  | voided                                                          |