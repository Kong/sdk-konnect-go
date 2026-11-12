# ProduceKeyValidationAction

Defines a behavior when record key is not valid.
* reject - rejects a batch for topic partition. Only available for produce.
* mark - marks a record with kong/server header and client ID value


  to help to identify the clients violating schema.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ProduceKeyValidationActionReject

// Open enum: custom values can be created with a direct type cast
custom := components.ProduceKeyValidationAction("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `ProduceKeyValidationActionReject` | reject                             |
| `ProduceKeyValidationActionMark`   | mark                               |