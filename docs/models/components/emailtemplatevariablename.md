# EmailTemplateVariableName

Enum for available portal email template variables.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.EmailTemplateVariableNamePortalDisplayName

// Open enum: custom values can be created with a direct type cast
custom := components.EmailTemplateVariableName("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `EmailTemplateVariableNamePortalDisplayName`   | portal_display_name                            |
| `EmailTemplateVariableNamePortalDomain`        | portal_domain                                  |
| `EmailTemplateVariableNameDeveloperEmail`      | developer_email                                |
| `EmailTemplateVariableNameDeveloperFullname`   | developer_fullname                             |
| `EmailTemplateVariableNameDevPortalReplyTo`    | dev_portal_reply_to                            |
| `EmailTemplateVariableNameDeveloperStatus`     | developer_status                               |
| `EmailTemplateVariableNameAPIDocumentationURL` | api_documentation_url                          |
| `EmailTemplateVariableNameAPIName`             | api_name                                       |
| `EmailTemplateVariableNameAPIVersion`          | api_version                                    |
| `EmailTemplateVariableNameApplicationName`     | application_name                               |
| `EmailTemplateVariableNameCredentialName`      | credential_name                                |
| `EmailTemplateVariableNameCredentialExpiresAt` | credential_expires_at                          |