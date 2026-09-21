# TryItUIAudience

The audience for the Try It UI feature.

`all` means that the Try It UI will be available to all users, including unauthenticated users.

`authenticated` means that the Try It UI will only be available to authenticated users.

`registered` means that the Try It UI will only be available to users who have registered for the API.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.TryItUIAudienceAll

// Open enum: custom values can be created with a direct type cast
custom := components.TryItUIAudience("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `TryItUIAudienceAll`           | all                            |
| `TryItUIAudienceAuthenticated` | authenticated                  |
| `TryItUIAudienceRegistered`    | registered                     |