# ResourceConfigurationQualifier

Enumeration of configuration qualifiers available for organization-wide configuration.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ResourceConfigurationQualifierDataPlaneGroupIdleTimeoutMinutes

// Open enum: custom values can be created with a direct type cast
custom := components.ResourceConfigurationQualifier("custom_value")
```


## Values

| Name                                                             | Value                                                            |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ResourceConfigurationQualifierDataPlaneGroupIdleTimeoutMinutes` | data-plane-group-idle-timeout-minutes                            |
| `ResourceConfigurationQualifierAutoPilotBaseRpsMaxValue`         | auto-pilot-base-rps-max-value                                    |