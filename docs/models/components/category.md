# Category

Categorical tag for the functionality the template relates to.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CategorySourceCodeManagement

// Open enum: custom values can be created with a direct type cast
custom := components.Category("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `CategorySourceCodeManagement`  | Source Code Management          |
| `CategoryIncidentManagement`    | Incident Management             |
| `CategoryObservability`         | Observability                   |
| `CategoryCodeQuality`           | Code Quality                    |
| `CategoryIssueTracking`         | Issue Tracking                  |
| `CategoryCommunications`        | Communications                  |
| `CategoryFeatureManagement`     | Feature Management              |
| `CategorySecurityVulnerability` | Security / Vulnerability        |
| `CategoryDocumentation`         | Documentation                   |
| `CategoryCatalogManagement`     | Catalog Management              |
| `CategoryGatewayConfiguration`  | Gateway Configuration           |