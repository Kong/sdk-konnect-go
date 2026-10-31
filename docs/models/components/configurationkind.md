# ConfigurationKind

**Pre-release Feature**
This feature is currently in beta and is subject to change.

Kind of the Cloud Gateway deployment. If serverless.v1 is specified, the following fields
should be omitted (will be ignored if provided): autoscale, cloud_gateway_network_id, version.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ConfigurationKindDedicatedV0

// Open enum: custom values can be created with a direct type cast
custom := components.ConfigurationKind("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ConfigurationKindDedicatedV0`  | dedicated.v0                    |
| `ConfigurationKindServerlessV1` | serverless.v1                   |