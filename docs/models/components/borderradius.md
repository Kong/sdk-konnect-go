# BorderRadius

The corner roundness to apply to this portal's theme.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BorderRadiusSoft

// Open enum: custom values can be created with a direct type cast
custom := components.BorderRadius("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `BorderRadiusSoft`  | soft                |
| `BorderRadiusSharp` | sharp               |
| `BorderRadiusRound` | round               |