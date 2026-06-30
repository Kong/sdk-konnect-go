# DeviceAuthorizationGrantTokenBadRequestDeviceAuthorizationGrantTokenBadRequestError

A single error code from the following enum.

  authorization_pending
     The authorization request is still pending as the end user hasn't yet completed the user-interaction steps.

  slow_down
     A variant of "authorization_pending", the authorization request is still pending and polling should continue, but the interval MUST
     be increased by 5 seconds for this and all subsequent requests.

  access_denied
     The authorization request was denied.

  expired_token
     The "device_code" has expired, and the device authorization session has concluded.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/sdkerrors"
)

value := sdkerrors.DeviceAuthorizationGrantTokenBadRequestDeviceAuthorizationGrantTokenBadRequestErrorAuthorizationPending

// Open enum: custom values can be created with a direct type cast
custom := sdkerrors.DeviceAuthorizationGrantTokenBadRequestDeviceAuthorizationGrantTokenBadRequestError("custom_value")
```


## Values

| Name                                                                                                      | Value                                                                                                     |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `DeviceAuthorizationGrantTokenBadRequestDeviceAuthorizationGrantTokenBadRequestErrorAuthorizationPending` | authorization_pending                                                                                     |
| `DeviceAuthorizationGrantTokenBadRequestDeviceAuthorizationGrantTokenBadRequestErrorSlowDown`             | slow_down                                                                                                 |
| `DeviceAuthorizationGrantTokenBadRequestDeviceAuthorizationGrantTokenBadRequestErrorAccessDenied`         | access_denied                                                                                             |
| `DeviceAuthorizationGrantTokenBadRequestDeviceAuthorizationGrantTokenBadRequestErrorExpiredToken`         | expired_token                                                                                             |