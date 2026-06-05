# ListCreditTransactionsParamsFilterType

Filter credit transactions by type.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ListCreditTransactionsParamsFilterTypeFunded

// Open enum: custom values can be created with a direct type cast
custom := components.ListCreditTransactionsParamsFilterType("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `ListCreditTransactionsParamsFilterTypeFunded`   | funded                                           |
| `ListCreditTransactionsParamsFilterTypeConsumed` | consumed                                         |
| `ListCreditTransactionsParamsFilterTypeExpired`  | expired                                          |