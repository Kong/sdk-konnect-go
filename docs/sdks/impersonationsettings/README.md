# ImpersonationSettings

## Overview

### Available Operations

* [GetImpersonationSettings](#getimpersonationsettings) - Get Impersonation Settings
* [UpdateImpersonationSettings](#updateimpersonationsettings) - Update Impersonation Settings

## GetImpersonationSettings

Returns Impersonation Settings, which determines if user impersonation is allowed for an organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-impersonation-settings" method="get" path="/v3/organizations/impersonation" example="Get Impersonation Settings" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ImpersonationSettings.GetImpersonationSettings(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetImpersonationSettingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetImpersonationSettingsResponse](../../models/operations/getimpersonationsettingsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateImpersonationSettings

Updates Impersonation Settings.

### Example Usage: Authentication Settings cannot be all Disabled

<!-- UsageSnippet language="go" operationID="update-impersonation-settings" method="patch" path="/v3/organizations/impersonation" example="Authentication Settings cannot be all Disabled" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ImpersonationSettings.UpdateImpersonationSettings(ctx, &components.UpdateImpersonationSettingsRequest{
        Enabled: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateImpersonationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: Cannot be Blank

<!-- UsageSnippet language="go" operationID="update-impersonation-settings" method="patch" path="/v3/organizations/impersonation" example="Cannot be Blank" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ImpersonationSettings.UpdateImpersonationSettings(ctx, &components.UpdateImpersonationSettingsRequest{
        Enabled: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateImpersonationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: IdP configuration is required

<!-- UsageSnippet language="go" operationID="update-impersonation-settings" method="patch" path="/v3/organizations/impersonation" example="IdP configuration is required" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ImpersonationSettings.UpdateImpersonationSettings(ctx, &components.UpdateImpersonationSettingsRequest{
        Enabled: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateImpersonationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: Invalid ID format

<!-- UsageSnippet language="go" operationID="update-impersonation-settings" method="patch" path="/v3/organizations/impersonation" example="Invalid ID format" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ImpersonationSettings.UpdateImpersonationSettings(ctx, &components.UpdateImpersonationSettingsRequest{
        Enabled: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateImpersonationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: Must be a valid UUID v4

<!-- UsageSnippet language="go" operationID="update-impersonation-settings" method="patch" path="/v3/organizations/impersonation" example="Must be a valid UUID v4" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ImpersonationSettings.UpdateImpersonationSettings(ctx, &components.UpdateImpersonationSettingsRequest{
        Enabled: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateImpersonationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: Not Found

<!-- UsageSnippet language="go" operationID="update-impersonation-settings" method="patch" path="/v3/organizations/impersonation" example="Not Found" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ImpersonationSettings.UpdateImpersonationSettings(ctx, &components.UpdateImpersonationSettingsRequest{
        Enabled: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateImpersonationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: OIDC needs an IdP configuration

<!-- UsageSnippet language="go" operationID="update-impersonation-settings" method="patch" path="/v3/organizations/impersonation" example="OIDC needs an IdP configuration" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ImpersonationSettings.UpdateImpersonationSettings(ctx, &components.UpdateImpersonationSettingsRequest{
        Enabled: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateImpersonationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: Password Complexity

<!-- UsageSnippet language="go" operationID="update-impersonation-settings" method="patch" path="/v3/organizations/impersonation" example="Password Complexity" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ImpersonationSettings.UpdateImpersonationSettings(ctx, &components.UpdateImpersonationSettingsRequest{
        Enabled: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateImpersonationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: Request Format is Invalid

<!-- UsageSnippet language="go" operationID="update-impersonation-settings" method="patch" path="/v3/organizations/impersonation" example="Request Format is Invalid" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ImpersonationSettings.UpdateImpersonationSettings(ctx, &components.UpdateImpersonationSettingsRequest{
        Enabled: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateImpersonationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: System teams cannot be modified

<!-- UsageSnippet language="go" operationID="update-impersonation-settings" method="patch" path="/v3/organizations/impersonation" example="System teams cannot be modified" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ImpersonationSettings.UpdateImpersonationSettings(ctx, &components.UpdateImpersonationSettingsRequest{
        Enabled: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateImpersonationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="update-impersonation-settings" method="patch" path="/v3/organizations/impersonation" example="Unauthorized" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ImpersonationSettings.UpdateImpersonationSettings(ctx, &components.UpdateImpersonationSettingsRequest{
        Enabled: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateImpersonationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: Unsupported filter operation

<!-- UsageSnippet language="go" operationID="update-impersonation-settings" method="patch" path="/v3/organizations/impersonation" example="Unsupported filter operation" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ImpersonationSettings.UpdateImpersonationSettings(ctx, &components.UpdateImpersonationSettingsRequest{
        Enabled: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateImpersonationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: Update Impersonation Settings

<!-- UsageSnippet language="go" operationID="update-impersonation-settings" method="patch" path="/v3/organizations/impersonation" example="Update Impersonation Settings" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ImpersonationSettings.UpdateImpersonationSettings(ctx, &components.UpdateImpersonationSettingsRequest{
        Enabled: sdkkonnectgo.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateImpersonationSettingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [components.UpdateImpersonationSettingsRequest](../../models/components/updateimpersonationsettingsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.UpdateImpersonationSettingsResponse](../../models/operations/updateimpersonationsettingsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |