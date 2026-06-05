# AdvancedMetrics

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AdvancedMetricsErrorRate

// Open enum: custom values can be created with a direct type cast
custom := components.AdvancedMetrics("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `AdvancedMetricsErrorRate`                  | error_rate                                  |
| `AdvancedMetricsKongInternalLatencyAverage` | kong_internal_latency_average               |
| `AdvancedMetricsKongInternalLatencyP50`     | kong_internal_latency_p50                   |
| `AdvancedMetricsKongInternalLatencyP95`     | kong_internal_latency_p95                   |
| `AdvancedMetricsKongInternalLatencyP99`     | kong_internal_latency_p99                   |
| `AdvancedMetricsKongLatencyAverage`         | kong_latency_average                        |
| `AdvancedMetricsKongLatencyP50`             | kong_latency_p50                            |
| `AdvancedMetricsKongLatencyP95`             | kong_latency_p95                            |
| `AdvancedMetricsKongLatencyP99`             | kong_latency_p99                            |
| `AdvancedMetricsRequestCount`               | request_count                               |
| `AdvancedMetricsRequestPerMinute`           | request_per_minute                          |
| `AdvancedMetricsRequestSizeAverage`         | request_size_average                        |
| `AdvancedMetricsRequestSizeP50`             | request_size_p50                            |
| `AdvancedMetricsRequestSizeP95`             | request_size_p95                            |
| `AdvancedMetricsRequestSizeP99`             | request_size_p99                            |
| `AdvancedMetricsRequestSizeSum`             | request_size_sum                            |
| `AdvancedMetricsResponseLatencyAverage`     | response_latency_average                    |
| `AdvancedMetricsResponseLatencyP50`         | response_latency_p50                        |
| `AdvancedMetricsResponseLatencyP95`         | response_latency_p95                        |
| `AdvancedMetricsResponseLatencyP99`         | response_latency_p99                        |
| `AdvancedMetricsResponseSizeAverage`        | response_size_average                       |
| `AdvancedMetricsResponseSizeP50`            | response_size_p50                           |
| `AdvancedMetricsResponseSizeP95`            | response_size_p95                           |
| `AdvancedMetricsResponseSizeP99`            | response_size_p99                           |
| `AdvancedMetricsResponseSizeSum`            | response_size_sum                           |
| `AdvancedMetricsUpstreamLatencyAverage`     | upstream_latency_average                    |
| `AdvancedMetricsUpstreamLatencyP50`         | upstream_latency_p50                        |
| `AdvancedMetricsUpstreamLatencyP95`         | upstream_latency_p95                        |
| `AdvancedMetricsUpstreamLatencyP99`         | upstream_latency_p99                        |