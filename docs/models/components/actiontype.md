# ActionType

Represents which action was performed on a resource:
  - `map`: a resource was mapped to a service
  - `unmap`: a resource was unmapped from a service
  - `archive`: a resource was archived
  - `restore`: a resource was restored


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ActionTypeMap

// Open enum: custom values can be created with a direct type cast
custom := components.ActionType("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `ActionTypeMap`     | map                 |
| `ActionTypeUnmap`   | unmap               |
| `ActionTypeArchive` | archive             |
| `ActionTypeRestore` | restore             |