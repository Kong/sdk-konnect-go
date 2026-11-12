# Method

The HTTP method of an API Operation.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MethodGet

// Open enum: custom values can be created with a direct type cast
custom := components.Method("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `MethodGet`     | GET             |
| `MethodPost`    | POST            |
| `MethodPut`     | PUT             |
| `MethodDelete`  | DELETE          |
| `MethodPatch`   | PATCH           |
| `MethodOptions` | OPTIONS         |
| `MethodHead`    | HEAD            |
| `MethodConnect` | CONNECT         |
| `MethodTrace`   | TRACE           |