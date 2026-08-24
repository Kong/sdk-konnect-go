# CreateAIGatewayRequestDeploymentType

How this AI Gateway's control plane is deployed. Set at creation time and cannot be changed afterward.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreateAIGatewayRequestDeploymentTypeHybrid

// Open enum: custom values can be created with a direct type cast
custom := components.CreateAIGatewayRequestDeploymentType("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `CreateAIGatewayRequestDeploymentTypeHybrid`     | hybrid                                           |
| `CreateAIGatewayRequestDeploymentTypeManaged`    | managed                                          |
| `CreateAIGatewayRequestDeploymentTypeServerless` | serverless                                       |