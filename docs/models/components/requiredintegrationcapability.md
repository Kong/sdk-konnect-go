# RequiredIntegrationCapability

Required capability from connected integrations to use this criteria template.
Applicable for templates that are not provided by a given integration (functional
across integrations).


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.RequiredIntegrationCapabilityOnCallProvider

// Open enum: custom values can be created with a direct type cast
custom := components.RequiredIntegrationCapability("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `RequiredIntegrationCapabilityOnCallProvider`      | on_call_provider                                   |
| `RequiredIntegrationCapabilityIncidentProvider`    | incident_provider                                  |
| `RequiredIntegrationCapabilityPullRequestProvider` | pull_request_provider                              |
| `RequiredIntegrationCapabilityWorkflowProvider`    | workflow_provider                                  |