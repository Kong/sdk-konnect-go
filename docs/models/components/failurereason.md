# FailureReason

Reason why delivery did not succeed: render_failed means the notification could not be prepared; http_error means the endpoint returned an unsuccessful HTTP response; retries_exhausted means no more attempts will be made; destination_inactive means the destination could not receive notifications; client_error means the request was rejected; expired means the delivery reached its expiry time.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.FailureReasonRenderFailed

// Open enum: custom values can be created with a direct type cast
custom := components.FailureReason("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `FailureReasonRenderFailed`        | render_failed                      |
| `FailureReasonHTTPError`           | http_error                         |
| `FailureReasonRetriesExhausted`    | retries_exhausted                  |
| `FailureReasonDestinationInactive` | destination_inactive               |
| `FailureReasonClientError`         | client_error                       |
| `FailureReasonExpired`             | expired                            |