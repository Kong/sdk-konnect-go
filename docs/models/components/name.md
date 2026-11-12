# Name

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.NameAll

// Open enum: custom values can be created with a direct type cast
custom := components.Name("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `NameAll`             | all                   |
| `NameAlter`           | alter                 |
| `NameAlterConfigs`    | alter_configs         |
| `NameCreate`          | create                |
| `NameDelete`          | delete                |
| `NameDescribe`        | describe              |
| `NameDescribeConfigs` | describe_configs      |
| `NameIdempotentWrite` | idempotent_write      |
| `NameRead`            | read                  |
| `NameWrite`           | write                 |