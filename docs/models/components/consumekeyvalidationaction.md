# ~~ConsumeKeyValidationAction~~

Deprecated. Use `failure_mode`.

Defines a behavior when record key is not valid.
* mark - marks a record with kong/server header and client ID value
  to help to identify the clients violating schema.
* skip - skips delivering a record.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ConsumeKeyValidationActionMark

// Open enum: custom values can be created with a direct type cast
custom := components.ConsumeKeyValidationAction("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `ConsumeKeyValidationActionMark` | mark                             |
| `ConsumeKeyValidationActionSkip` | skip                             |