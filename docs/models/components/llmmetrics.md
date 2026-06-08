# LLMMetrics

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.LLMMetricsTotalTokens

// Open enum: custom values can be created with a direct type cast
custom := components.LLMMetrics("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `LLMMetricsTotalTokens`                      | total_tokens                                 |
| `LLMMetricsPromptTokens`                     | prompt_tokens                                |
| `LLMMetricsCompletionTokens`                 | completion_tokens                            |
| `LLMMetricsAiRequestCount`                   | ai_request_count                             |
| `LLMMetricsCost`                             | cost                                         |
| `LLMMetricsErrorRate`                        | error_rate                                   |
| `LLMMetricsLlmCacheEmbeddingsLatencyAverage` | llm_cache_embeddings_latency_average         |
| `LLMMetricsLlmCacheFetchLatencyAverage`      | llm_cache_fetch_latency_average              |
| `LLMMetricsLlmLatencyAverage`                | llm_latency_average                          |
| `LLMMetricsLlmEmbeddingsTokens`              | llm_embeddings_tokens                        |
| `LLMMetricsLlmEmbeddingsCost`                | llm_embeddings_cost                          |