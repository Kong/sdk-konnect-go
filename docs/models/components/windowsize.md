# WindowSize

The window size in seconds, one of a second, a minute, an hour, or a day.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.WindowSizeOne

// Open enum: custom values can be created with a direct type cast
custom := components.WindowSize(999)
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `WindowSizeOne`                          | 1                                        |
| `WindowSizeSixty`                        | 60                                       |
| `WindowSizeThreeThousandSixHundred`      | 3600                                     |
| `WindowSizeEightySixThousandFourHundred` | 86400                                    |