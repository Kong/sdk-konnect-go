# WebhookStatus

Current status of a webhook. `active` indicates the webhook is sending or ready to send requests.
`inactive` indicates the webhook has been turned off due to failed attempts.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.WebhookStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.WebhookStatus("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `WebhookStatusActive`   | active                  |
| `WebhookStatusInactive` | inactive                |