# ProviderTemplateName

The static name of the vault provider template to base this provider on (e.g. "github"). For static_secret, provide only name and credential_type ("shared"); the credential is provisioned separately via the vault provider credentials API, and client_id/client_secret are rejected.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ProviderTemplateNameGithub

// Open enum: custom values can be created with a direct type cast
custom := components.ProviderTemplateName("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `ProviderTemplateNameGithub`        | github                              |
| `ProviderTemplateNameSlack`         | slack                               |
| `ProviderTemplateNameStaticSecret`  | static_secret                       |
| `ProviderTemplateNameGoogle`        | google                              |
| `ProviderTemplateNameAtlassianRovo` | atlassian-rovo                      |
| `ProviderTemplateNameDatadog`       | datadog                             |
| `ProviderTemplateNameCustomOauth`   | custom-oauth                        |
| `ProviderTemplateNameMiro`          | miro                                |