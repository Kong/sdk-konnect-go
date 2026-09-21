# AIGatewayDatastoreUsageEntryType

The kind of entity holding the reference. A policy holds one in its `datastores` array. A model holds one when its balancer is `semantic`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayDatastoreUsageEntryTypePolicy

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayDatastoreUsageEntryType("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `AIGatewayDatastoreUsageEntryTypePolicy` | policy                                   |
| `AIGatewayDatastoreUsageEntryTypeModel`  | model                                    |