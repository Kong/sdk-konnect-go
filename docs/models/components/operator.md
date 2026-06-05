# Operator

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.OperatorIn

// Open enum: custom values can be created with a direct type cast
custom := components.Operator("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `OperatorIn`       | in                 |
| `OperatorNotIn`    | not_in             |
| `OperatorEmpty`    | empty              |
| `OperatorNotEmpty` | not_empty          |