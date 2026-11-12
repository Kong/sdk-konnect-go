# PartialVectordbStrategy

which vector database driver to use

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PartialVectordbStrategyPgvector

// Open enum: custom values can be created with a direct type cast
custom := components.PartialVectordbStrategy("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PartialVectordbStrategyPgvector` | pgvector                          |
| `PartialVectordbStrategyRedis`    | redis                             |