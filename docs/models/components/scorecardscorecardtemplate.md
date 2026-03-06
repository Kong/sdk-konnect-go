# ScorecardScorecardTemplate

The name of the scorecard template used to create the scorecard.
Otherwise, `null`.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ScorecardScorecardTemplateKongBestPractices

// Open enum: custom values can be created with a direct type cast
custom := components.ScorecardScorecardTemplate("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `ScorecardScorecardTemplateKongBestPractices`     | kong_best_practices                               |
| `ScorecardScorecardTemplateServiceDocumentation`  | service_documentation                             |
| `ScorecardScorecardTemplateServiceMaturity`       | service_maturity                                  |
| `ScorecardScorecardTemplateSecurityAndCompliance` | security_and_compliance                           |