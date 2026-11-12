# ResponseType

OAuth 2.0 response type

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ResponseTypeNone

// Open enum: custom values can be created with a direct type cast
custom := components.ResponseType("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `ResponseTypeNone`    | none                  |
| `ResponseTypeToken`   | token                 |
| `ResponseTypeCode`    | code                  |
| `ResponseTypeIDToken` | id_token              |