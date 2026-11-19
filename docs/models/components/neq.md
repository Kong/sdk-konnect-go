# Neq

The value to exclude.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.NeqPrincipal

// Open enum: custom values can be created with a direct type cast
custom := components.Neq("custom_value")
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `NeqPrincipal` | principal      |
| `NeqConsumer`  | consumer       |
| `NeqOidc`      | oidc           |