# Direction

`asc` for ascending or `desc` for descending order. Defaults to `asc`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/operations"
)

value := operations.DirectionAsc

// Open enum: custom values can be created with a direct type cast
custom := operations.Direction("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `DirectionAsc`  | asc             |
| `DirectionDesc` | desc            |