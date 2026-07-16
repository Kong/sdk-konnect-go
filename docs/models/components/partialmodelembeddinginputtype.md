# PartialModelEmbeddingInputType

The purpose of the input text to calculate embedding vectors.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PartialModelEmbeddingInputTypeClassification

// Open enum: custom values can be created with a direct type cast
custom := components.PartialModelEmbeddingInputType("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `PartialModelEmbeddingInputTypeClassification` | classification                                 |
| `PartialModelEmbeddingInputTypeClustering`     | clustering                                     |
| `PartialModelEmbeddingInputTypeImage`          | image                                          |
| `PartialModelEmbeddingInputTypeSearchDocument` | search_document                                |
| `PartialModelEmbeddingInputTypeSearchQuery`    | search_query                                   |