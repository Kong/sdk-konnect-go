# CreateScorecardScorecardTemplate

The name of the scorecard template used to create the scorecard.
Otherwise, `null`.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreateScorecardScorecardTemplateKongBestPractices

// Open enum: custom values can be created with a direct type cast
custom := components.CreateScorecardScorecardTemplate("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `CreateScorecardScorecardTemplateKongBestPractices`     | kong_best_practices                                     |
| `CreateScorecardScorecardTemplateServiceDocumentation`  | service_documentation                                   |
| `CreateScorecardScorecardTemplateServiceMaturity`       | service_maturity                                        |
| `CreateScorecardScorecardTemplateSecurityAndCompliance` | security_and_compliance                                 |