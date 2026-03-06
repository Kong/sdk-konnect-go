# CatalogServiceScorecardScorecardTemplate

The name of the scorecard template used to create the scorecard.
Otherwise, `null`.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CatalogServiceScorecardScorecardTemplateKongBestPractices

// Open enum: custom values can be created with a direct type cast
custom := components.CatalogServiceScorecardScorecardTemplate("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `CatalogServiceScorecardScorecardTemplateKongBestPractices`     | kong_best_practices                                             |
| `CatalogServiceScorecardScorecardTemplateServiceDocumentation`  | service_documentation                                           |
| `CatalogServiceScorecardScorecardTemplateServiceMaturity`       | service_maturity                                                |
| `CatalogServiceScorecardScorecardTemplateSecurityAndCompliance` | security_and_compliance                                         |