# Action

How to handle the request if the rule matches

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ActionAllow

// Open enum: custom values can be created with a direct type cast
custom := components.Action("custom_value")
```


## Values

| Name          | Value         |
| ------------- | ------------- |
| `ActionAllow` | allow         |
| `ActionDeny`  | deny          |