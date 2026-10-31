# LLMFiltersOperator

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.LLMFiltersOperatorIn

// Open enum: custom values can be created with a direct type cast
custom := components.LLMFiltersOperator("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `LLMFiltersOperatorIn`       | in                           |
| `LLMFiltersOperatorNotIn`    | not_in                       |
| `LLMFiltersOperatorEmpty`    | empty                        |
| `LLMFiltersOperatorNotEmpty` | not_empty                    |