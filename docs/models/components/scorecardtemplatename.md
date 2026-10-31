# ScorecardTemplateName

Well-known name of the template. Uniquely identifies the template.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ScorecardTemplateNameKongBestPractices

// Open enum: custom values can be created with a direct type cast
custom := components.ScorecardTemplateName("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `ScorecardTemplateNameKongBestPractices`     | kong_best_practices                          |
| `ScorecardTemplateNameServiceDocumentation`  | service_documentation                        |
| `ScorecardTemplateNameServiceMaturity`       | service_maturity                             |
| `ScorecardTemplateNameSecurityAndCompliance` | security_and_compliance                      |