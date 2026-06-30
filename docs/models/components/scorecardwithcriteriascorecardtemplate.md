# ScorecardWithCriteriaScorecardTemplate

The name of the scorecard template used to create the scorecard.
Otherwise, `null`.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ScorecardWithCriteriaScorecardTemplateKongBestPractices

// Open enum: custom values can be created with a direct type cast
custom := components.ScorecardWithCriteriaScorecardTemplate("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ScorecardWithCriteriaScorecardTemplateKongBestPractices`     | kong_best_practices                                           |
| `ScorecardWithCriteriaScorecardTemplateServiceDocumentation`  | service_documentation                                         |
| `ScorecardWithCriteriaScorecardTemplateServiceMaturity`       | service_maturity                                              |
| `ScorecardWithCriteriaScorecardTemplateSecurityAndCompliance` | security_and_compliance                                       |