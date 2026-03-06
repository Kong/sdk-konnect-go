# ConsumeValueValidationAction

Defines a behavior when record value is not valid.
* mark - marks a record with kong/server header and client ID value
  to help to identify the clients violating schema.
* skip - skips delivering a record.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ConsumeValueValidationActionMark

// Open enum: custom values can be created with a direct type cast
custom := components.ConsumeValueValidationAction("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `ConsumeValueValidationActionMark` | mark                               |
| `ConsumeValueValidationActionSkip` | skip                               |