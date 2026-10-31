# DeploymentType

How this AI Gateway's control plane is deployed.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.DeploymentTypeHybrid

// Open enum: custom values can be created with a direct type cast
custom := components.DeploymentType("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `DeploymentTypeHybrid`     | hybrid                     |
| `DeploymentTypeManaged`    | managed                    |
| `DeploymentTypeServerless` | serverless                 |