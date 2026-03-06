# PublishedStatus

Whether the resource is visible on a given portal. Defaults to unpublished.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PublishedStatusPublished

// Open enum: custom values can be created with a direct type cast
custom := components.PublishedStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `PublishedStatusPublished`   | published                    |
| `PublishedStatusUnpublished` | unpublished                  |