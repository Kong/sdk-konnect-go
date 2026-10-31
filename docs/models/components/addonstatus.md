# AddonStatus

The status of the add-on. Computed based on the effective start and end dates:

- `draft`: `effective_from` is not set.
- `active`: `effective_from <= now` and (`effective_to` is not set or
`now < effective_to`).
- `archived`: `effective_to <= now`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AddonStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.AddonStatus("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `AddonStatusDraft`    | draft                 |
| `AddonStatusActive`   | active                |
| `AddonStatusArchived` | archived              |