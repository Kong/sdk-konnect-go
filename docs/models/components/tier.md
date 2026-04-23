# Tier

Capacity tier that determines both cache size and performance characteristics:
- micro: ~0.5 GiB capacity
- small: ~1 GiB capacity
- medium: ~3 GiB capacity
- large: ~6 GiB capacity
- xlarge: ~12 GiB capacity
- 2xlarge: ~25 GiB capacity
- 4xlarge: ~52 GiB capacity
- 8xlarge: ~100 GiB capacity
- 12xlarge: ~150 GiB capacity
- 16xlarge: ~200 GiB capacity
- 24xlarge: ~300 GiB capacity


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.TierMicro

// Open enum: custom values can be created with a direct type cast
custom := components.Tier("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `TierMicro`            | micro                  |
| `TierSmall`            | small                  |
| `TierMedium`           | medium                 |
| `TierLarge`            | large                  |
| `TierXlarge`           | xlarge                 |
| `TierTwoxlarge`        | 2xlarge                |
| `TierFourxlarge`       | 4xlarge                |
| `TierEightxlarge`      | 8xlarge                |
| `TierTwelvexlarge`     | 12xlarge               |
| `TierSixteenxlarge`    | 16xlarge               |
| `TierTwentyFourxlarge` | 24xlarge               |