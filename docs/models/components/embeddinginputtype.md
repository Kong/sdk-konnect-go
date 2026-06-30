# EmbeddingInputType

The purpose of the input text to calculate embedding vectors.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.EmbeddingInputTypeClassification

// Open enum: custom values can be created with a direct type cast
custom := components.EmbeddingInputType("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `EmbeddingInputTypeClassification` | classification                     |
| `EmbeddingInputTypeClustering`     | clustering                         |
| `EmbeddingInputTypeImage`          | image                              |
| `EmbeddingInputTypeSearchDocument` | search_document                    |
| `EmbeddingInputTypeSearchQuery`    | search_query                       |