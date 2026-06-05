# EntityType

The type of entity that is being published. Can be either an API or an API Package.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.EntityTypeAPI

// Open enum: custom values can be created with a direct type cast
custom := components.EntityType("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `EntityTypeAPI`        | api                    |
| `EntityTypeAPIPackage` | api_package            |