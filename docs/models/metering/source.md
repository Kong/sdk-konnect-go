# Source

Where this price came from.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/metering"
)

value := metering.SourceManual

// Open enum: custom values can be created with a direct type cast
custom := metering.Source("custom_value")
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `SourceManual` | manual         |
| `SourceSystem` | system         |