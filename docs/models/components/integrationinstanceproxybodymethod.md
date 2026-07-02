# IntegrationInstanceProxyBodyMethod

The HTTP request method.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.IntegrationInstanceProxyBodyMethodDelete

// Open enum: custom values can be created with a direct type cast
custom := components.IntegrationInstanceProxyBodyMethod("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `IntegrationInstanceProxyBodyMethodDelete` | DELETE                                     |
| `IntegrationInstanceProxyBodyMethodGet`    | GET                                        |
| `IntegrationInstanceProxyBodyMethodPatch`  | PATCH                                      |
| `IntegrationInstanceProxyBodyMethodPost`   | POST                                       |
| `IntegrationInstanceProxyBodyMethodPut`    | PUT                                        |