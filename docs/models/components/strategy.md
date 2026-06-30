# Strategy

which vector database driver to use

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.StrategyPgvector

// Open enum: custom values can be created with a direct type cast
custom := components.Strategy("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `StrategyPgvector` | pgvector           |
| `StrategyRedis`    | redis              |