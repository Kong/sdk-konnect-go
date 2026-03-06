# APIAccess

Type of API access data-plane groups will support for a configuration.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.APIAccessPrivate

// Open enum: custom values can be created with a direct type cast
custom := components.APIAccess("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `APIAccessPrivate`           | private                      |
| `APIAccessPublic`            | public                       |
| `APIAccessPrivatePlusPublic` | private+public               |