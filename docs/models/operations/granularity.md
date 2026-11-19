# Granularity

Splits each group's totals into time buckets of this calendar interval, returned in `series`:

- `PT1H`: an hour.
- `P1D`: a day.
- `P1W`: an ISO week, starting on Monday.
- `P1M`: a calendar month.

Bucket boundaries follow the calendar in `time_zone`, so the first and last buckets can extend beyond `starts_at` and `ends_at`. The range can span at most 400 buckets. Without `granularity`, each group has only its totals for the whole range, and `series` is empty.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/operations"
)

value := operations.GranularityPt1H

// Open enum: custom values can be created with a direct type cast
custom := operations.Granularity("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `GranularityPt1H` | PT1H              |
| `GranularityP1D`  | P1D               |
| `GranularityP1W`  | P1W               |
| `GranularityP1M`  | P1M               |