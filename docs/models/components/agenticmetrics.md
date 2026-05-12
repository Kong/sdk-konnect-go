# AgenticMetrics

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AgenticMetricsA2aLatencyAverage

// Open enum: custom values can be created with a direct type cast
custom := components.AgenticMetrics("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `AgenticMetricsA2aLatencyAverage`      | a2a_latency_average                    |
| `AgenticMetricsA2aResponseSizeSum`     | a2a_response_size_sum                  |
| `AgenticMetricsErrorRate`              | error_rate                             |
| `AgenticMetricsKongLatencyAverage`     | kong_latency_average                   |
| `AgenticMetricsKongLatencyP50`         | kong_latency_p50                       |
| `AgenticMetricsKongLatencyP95`         | kong_latency_p95                       |
| `AgenticMetricsKongLatencyP99`         | kong_latency_p99                       |
| `AgenticMetricsMcpResponseSizeSum`     | mcp_response_size_sum                  |
| `AgenticMetricsRequestCount`           | request_count                          |
| `AgenticMetricsRequestPerMinute`       | request_per_minute                     |
| `AgenticMetricsRequestSizeAverage`     | request_size_average                   |
| `AgenticMetricsRequestSizeP50`         | request_size_p50                       |
| `AgenticMetricsRequestSizeP95`         | request_size_p95                       |
| `AgenticMetricsRequestSizeP99`         | request_size_p99                       |
| `AgenticMetricsRequestSizeSum`         | request_size_sum                       |
| `AgenticMetricsResponseLatencyAverage` | response_latency_average               |
| `AgenticMetricsResponseLatencyP50`     | response_latency_p50                   |
| `AgenticMetricsResponseLatencyP95`     | response_latency_p95                   |
| `AgenticMetricsResponseLatencyP99`     | response_latency_p99                   |
| `AgenticMetricsResponseSizeAverage`    | response_size_average                  |
| `AgenticMetricsResponseSizeP50`        | response_size_p50                      |
| `AgenticMetricsResponseSizeP95`        | response_size_p95                      |
| `AgenticMetricsResponseSizeP99`        | response_size_p99                      |
| `AgenticMetricsResponseSizeSum`        | response_size_sum                      |
| `AgenticMetricsUpstreamLatencyAverage` | upstream_latency_average               |
| `AgenticMetricsUpstreamLatencyP50`     | upstream_latency_p50                   |
| `AgenticMetricsUpstreamLatencyP95`     | upstream_latency_p95                   |
| `AgenticMetricsUpstreamLatencyP99`     | upstream_latency_p99                   |