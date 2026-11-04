# PortalAgentAudienceNullable

The audience allowed to interact with the Portal Agent. Currently only `all` (available to every developer, including unauthenticated ones) is supported. `authenticated` (restricted to signed-in developers only) is planned for a future release.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PortalAgentAudienceNullableAll
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `PortalAgentAudienceNullableAll` | all                              |