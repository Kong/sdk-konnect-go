# Measures

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/operations"
)

value := operations.MeasuresEvents

// Open enum: custom values can be created with a direct type cast
custom := operations.Measures("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `MeasuresEvents`                     | events                               |
| `MeasuresUnpricedEvents`             | unpriced_events                      |
| `MeasuresReportedCost`               | reported_cost                        |
| `MeasuresReportedCacheCost`          | reported_cache_cost                  |
| `MeasuresInputTokens`                | input_tokens                         |
| `MeasuresInputTextTokens`            | input_text_tokens                    |
| `MeasuresInputAudioTokens`           | input_audio_tokens                   |
| `MeasuresInputImageTokens`           | input_image_tokens                   |
| `MeasuresInputCacheReadTokens`       | input_cache_read_tokens              |
| `MeasuresInputCacheReadTextTokens`   | input_cache_read_text_tokens         |
| `MeasuresInputCacheReadAudioTokens`  | input_cache_read_audio_tokens        |
| `MeasuresInputCacheReadImageTokens`  | input_cache_read_image_tokens        |
| `MeasuresInputCacheReadVideoTokens`  | input_cache_read_video_tokens        |
| `MeasuresInputCacheWriteTokens`      | input_cache_write_tokens             |
| `MeasuresInputCacheWriteTextTokens`  | input_cache_write_text_tokens        |
| `MeasuresInputCacheWriteAudioTokens` | input_cache_write_audio_tokens       |
| `MeasuresInputCacheWriteImageTokens` | input_cache_write_image_tokens       |
| `MeasuresInputCacheWriteVideoTokens` | input_cache_write_video_tokens       |
| `MeasuresOutputTokens`               | output_tokens                        |
| `MeasuresOutputTextTokens`           | output_text_tokens                   |
| `MeasuresOutputAudioTokens`          | output_audio_tokens                  |
| `MeasuresOutputReasoningTokens`      | output_reasoning_tokens              |
| `MeasuresAcceptedPredictionTokens`   | accepted_prediction_tokens           |
| `MeasuresRejectedPredictionTokens`   | rejected_prediction_tokens           |
| `MeasuresRealtimeAudioMs`            | realtime_audio_ms                    |
| `MeasuresReportedTotalTokens`        | reported_total_tokens                |
| `MeasuresTotalInputTokens`           | total_input_tokens                   |
| `MeasuresSubjects`                   | subjects                             |