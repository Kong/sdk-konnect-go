# EmailTemplateName

Short name email template name.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.EmailTemplateNameConfirmEmailAddress

// Open enum: custom values can be created with a direct type cast
custom := components.EmailTemplateName("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `EmailTemplateNameConfirmEmailAddress`     | confirm-email-address                      |
| `EmailTemplateNameAppRegistrationApproved` | app-registration-approved                  |
| `EmailTemplateNameAppRegistrationRejected` | app-registration-rejected                  |
| `EmailTemplateNameAppRegistrationRevoked`  | app-registration-revoked                   |
| `EmailTemplateNameResetPassword`           | reset-password                             |
| `EmailTemplateNameAccountAccessApproved`   | account-access-approved                    |
| `EmailTemplateNameAccountAccessRejected`   | account-access-rejected                    |
| `EmailTemplateNameAccountAccessRevoked`    | account-access-revoked                     |
| `EmailTemplateNameAPIKeyExpiringSoon`      | api-key-expiring-soon                      |
| `EmailTemplateNameAPIKeyExpired`           | api-key-expired                            |
| `EmailTemplateNameDeveloperInvitation`     | developer-invitation                       |