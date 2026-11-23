# SchemaVersion

The schema version of each export's data files within the dataset, which fixes their columns and values.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SchemaVersionOne

// Open enum: custom values can be created with a direct type cast
custom := components.SchemaVersion(999)
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `SchemaVersionOne` | 1                  |
| `SchemaVersionTwo` | 2                  |