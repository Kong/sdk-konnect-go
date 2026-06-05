# TokensCountStrategy

Methodology to use for token usage calculation.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.TokensCountStrategyCompletionTokens

// Open enum: custom values can be created with a direct type cast
custom := components.TokensCountStrategy("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `TokensCountStrategyCompletionTokens` | completion-tokens                     |
| `TokensCountStrategyCost`             | cost                                  |
| `TokensCountStrategyLlmAccuracy`      | llm-accuracy                          |
| `TokensCountStrategyPromptTokens`     | prompt-tokens                         |
| `TokensCountStrategyTotalTokens`      | total-tokens                          |