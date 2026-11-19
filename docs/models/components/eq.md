# Eq

The value to match.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.EqPrincipal

// Open enum: custom values can be created with a direct type cast
custom := components.Eq("custom_value")
```


## Values

| Name          | Value         |
| ------------- | ------------- |
| `EqPrincipal` | principal     |
| `EqConsumer`  | consumer      |
| `EqOidc`      | oidc          |