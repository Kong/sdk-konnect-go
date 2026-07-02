# Dimensions

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.DimensionsTime

// Open enum: custom values can be created with a direct type cast
custom := components.Dimensions("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `DimensionsTime`        | time                    |
| `DimensionsService`     | service                 |
| `DimensionsSeverity`    | severity                |
| `DimensionsState`       | state                   |
| `DimensionsType`        | type                    |
| `DimensionsSource`      | source                  |
| `DimensionsEnvironment` | environment             |
| `DimensionsRegion`      | region                  |