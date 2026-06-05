# HashOn

What to use as hashing input. Using `none` results in a weighted-round-robin scheme with no hashing.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.HashOnConsumer

// Open enum: custom values can be created with a direct type cast
custom := components.HashOn("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `HashOnConsumer`   | consumer           |
| `HashOnCookie`     | cookie             |
| `HashOnHeader`     | header             |
| `HashOnIP`         | ip                 |
| `HashOnNone`       | none               |
| `HashOnPath`       | path               |
| `HashOnQueryArg`   | query_arg          |
| `HashOnURICapture` | uri_capture        |