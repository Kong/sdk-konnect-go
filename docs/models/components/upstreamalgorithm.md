# UpstreamAlgorithm

Which load balancing algorithm to use.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.UpstreamAlgorithmConsistentHashing

// Open enum: custom values can be created with a direct type cast
custom := components.UpstreamAlgorithm("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `UpstreamAlgorithmConsistentHashing` | consistent-hashing                   |
| `UpstreamAlgorithmLatency`           | latency                              |
| `UpstreamAlgorithmLeastConnections`  | least-connections                    |
| `UpstreamAlgorithmRoundRobin`        | round-robin                          |
| `UpstreamAlgorithmStickySessions`    | sticky-sessions                      |