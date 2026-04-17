# Op

The operation to perform.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.OpAdd

// Open enum: custom values can be created with a direct type cast
custom := components.Op("custom_value")
```


## Values

| Name        | Value       |
| ----------- | ----------- |
| `OpAdd`     | add         |
| `OpRemove`  | remove      |
| `OpReplace` | replace     |