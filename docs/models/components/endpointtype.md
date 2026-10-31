# EndpointType

The AWS Bedrock inference endpoint that requests to this target are sent to.
`runtime` sends requests to the Bedrock Runtime endpoint
(`bedrock-runtime.{region}.amazonaws.com`). `mantle` sends requests to the
Bedrock Mantle endpoint (`bedrock-mantle.{region}.api.aws`), which accepts
OpenAI and Anthropic request formats natively.

A target with `endpoint_type` set to `mantle` is rejected if its model
includes a format with `type` set to `bedrock`, because the Mantle endpoint
does not support the Bedrock Converse or InvokeModel APIs.

**Requires a minimum runtime version of `2.3`**.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.EndpointTypeRuntime

// Open enum: custom values can be created with a direct type cast
custom := components.EndpointType("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `EndpointTypeRuntime` | runtime               |
| `EndpointTypeMantle`  | mantle                |