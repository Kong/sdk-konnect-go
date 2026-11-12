# AIGatewayModelProviderConfigAuthBasicLocation

Specify whether the param name and value options go in a query string, or the POST form/JSON body.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayModelProviderConfigAuthBasicLocationBody

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayModelProviderConfigAuthBasicLocation("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `AIGatewayModelProviderConfigAuthBasicLocationBody`  | body                                                 |
| `AIGatewayModelProviderConfigAuthBasicLocationQuery` | query                                                |