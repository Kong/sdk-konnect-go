# ResourceType

This rule applies to access only for type of resource

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ResourceTypeTopic

// Open enum: custom values can be created with a direct type cast
custom := components.ResourceType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `ResourceTypeTopic`           | topic                         |
| `ResourceTypeGroup`           | group                         |
| `ResourceTypeTransactionalID` | transactional_id              |
| `ResourceTypeCluster`         | cluster                       |