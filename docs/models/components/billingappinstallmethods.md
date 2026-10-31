# BillingAppInstallMethods

Supported installation methods for an app.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingAppInstallMethodsWithOauth2

// Open enum: custom values can be created with a direct type cast
custom := components.BillingAppInstallMethods("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `BillingAppInstallMethodsWithOauth2`            | with_oauth2                                     |
| `BillingAppInstallMethodsWithAPIKey`            | with_api_key                                    |
| `BillingAppInstallMethodsNoCredentialsRequired` | no_credentials_required                         |