# RouteExpressionPathHandling

Controls how the Service path, Route path and requested path are combined when sending a request to the upstream. See above for a detailed description of each behavior.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.RouteExpressionPathHandlingV0

// Open enum: custom values can be created with a direct type cast
custom := components.RouteExpressionPathHandling("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `RouteExpressionPathHandlingV0` | v0                              |
| `RouteExpressionPathHandlingV1` | v1                              |