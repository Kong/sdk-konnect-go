# PortalAuthSettings

## Overview

APIs related to configuration of Konnect Developer Portal auth settings.

### Available Operations

* [GetPortalAuthenticationSettings](#getportalauthenticationsettings) - Get Auth Settings
* [UpdatePortalAuthenticationSettings](#updateportalauthenticationsettings) - Update Auth Settings
* [ListPortalTeamGroupMappings](#listportalteamgroupmappings) - List Team Group Mappings
* [UpdatePortalTeamGroupMappings](#updateportalteamgroupmappings) - Update Team Group Mappings
* [GetPortalIdentityProviders](#getportalidentityproviders) - List Identity Providers
* [CreatePortalIdentityProvider](#createportalidentityprovider) - Create Identity Provider
* [GetPortalIdentityProvider](#getportalidentityprovider) - Get Identity Provider
* [UpdatePortalIdentityProvider](#updateportalidentityprovider) - Update Identity Provider
* [DeletePortalIdentityProvider](#deleteportalidentityprovider) - Delete Identity Provider

## GetPortalAuthenticationSettings

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns the developer authentication configuration for a portal, which determines how developers can log in and how they are assigned to teams.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-portal-authentication-settings" method="get" path="/v3/portals/{portalId}/authentication-settings" example="Example 1" -->
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

    res, err := s.PortalAuthSettings.GetPortalAuthenticationSettings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuthenticationSettingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `portalID`                                               | *string*                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPortalAuthenticationSettingsResponse](../../models/operations/getportalauthenticationsettingsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePortalAuthenticationSettings

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Updates the developer authentication configuration for a portal. Developers can be allowed to login using basic auth (email & password) or use Single-Sign-On through an Identity Provider. Developers can be automatically assigned to teams by mapping claims from their IdP account.

### Example Usage: Disable Basic Developer Auth

<!-- UsageSnippet language="go" operationID="update-portal-authentication-settings" method="patch" path="/v3/portals/{portalId}/authentication-settings" example="Disable Basic Developer Auth" -->
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

    res, err := s.PortalAuthSettings.UpdatePortalAuthenticationSettings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalAuthenticationSettingsUpdateRequest{
        BasicAuthEnabled: sdkkonnectgo.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuthenticationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: Enable OIDC Developer Auth

<!-- UsageSnippet language="go" operationID="update-portal-authentication-settings" method="patch" path="/v3/portals/{portalId}/authentication-settings" example="Enable OIDC Developer Auth" -->
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

    res, err := s.PortalAuthSettings.UpdatePortalAuthenticationSettings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalAuthenticationSettingsUpdateRequest{
        OidcAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcTeamMappingEnabled: sdkkonnectgo.Pointer(false),
        KonnectMappingEnabled: sdkkonnectgo.Pointer(true),
        OidcIssuer: sdkkonnectgo.Pointer("https://identity.example.com/v2"),
        OidcClientID: sdkkonnectgo.Pointer("x7id0o42lklas0blidl2"),
        OidcScopes: []string{
            "email",
            "openid",
            "profile",
        },
        OidcClaimMappings: &components.PortalAuthenticationSettingsUpdateRequestPortalClaimMappings{
            Groups: sdkkonnectgo.Pointer("custom-group-claim"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuthenticationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: Enable OIDC Team Mapping

<!-- UsageSnippet language="go" operationID="update-portal-authentication-settings" method="patch" path="/v3/portals/{portalId}/authentication-settings" example="Enable OIDC Team Mapping" -->
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

    res, err := s.PortalAuthSettings.UpdatePortalAuthenticationSettings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalAuthenticationSettingsUpdateRequest{
        OidcTeamMappingEnabled: sdkkonnectgo.Pointer(true),
        KonnectMappingEnabled: sdkkonnectgo.Pointer(true),
        OidcClaimMappings: &components.PortalAuthenticationSettingsUpdateRequestPortalClaimMappings{
            Groups: sdkkonnectgo.Pointer("custom-group-claim"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuthenticationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: Example 1

<!-- UsageSnippet language="go" operationID="update-portal-authentication-settings" method="patch" path="/v3/portals/{portalId}/authentication-settings" example="Example 1" -->
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

    res, err := s.PortalAuthSettings.UpdatePortalAuthenticationSettings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalAuthenticationSettingsUpdateRequest{
        BasicAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcAuthEnabled: sdkkonnectgo.Pointer(true),
        SamlAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcTeamMappingEnabled: sdkkonnectgo.Pointer(true),
        KonnectMappingEnabled: sdkkonnectgo.Pointer(false),
        IdpMappingEnabled: sdkkonnectgo.Pointer(true),
        OidcIssuer: sdkkonnectgo.Pointer("https://identity.example.com/v2"),
        OidcClientID: sdkkonnectgo.Pointer("x7id0o42lklas0blidl2"),
        OidcScopes: []string{
            "email",
            "openid",
            "profile",
        },
        OidcClaimMappings: &components.PortalAuthenticationSettingsUpdateRequestPortalClaimMappings{
            Groups: sdkkonnectgo.Pointer("custom-group-claim"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuthenticationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="update-portal-authentication-settings" method="patch" path="/v3/portals/{portalId}/authentication-settings" example="Unauthorized" -->
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

    res, err := s.PortalAuthSettings.UpdatePortalAuthenticationSettings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalAuthenticationSettingsUpdateRequest{
        BasicAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcTeamMappingEnabled: sdkkonnectgo.Pointer(true),
        KonnectMappingEnabled: sdkkonnectgo.Pointer(false),
        OidcIssuer: sdkkonnectgo.Pointer("https://identity.example.com/v2"),
        OidcClientID: sdkkonnectgo.Pointer("x7id0o42lklas0blidl2"),
        OidcScopes: []string{
            "email",
            "openid",
            "profile",
        },
        OidcClaimMappings: &components.PortalAuthenticationSettingsUpdateRequestPortalClaimMappings{
            Groups: sdkkonnectgo.Pointer("custom-group-claim"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuthenticationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="update-portal-authentication-settings" method="patch" path="/v3/portals/{portalId}/authentication-settings" example="UnauthorizedExample" -->
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

    res, err := s.PortalAuthSettings.UpdatePortalAuthenticationSettings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalAuthenticationSettingsUpdateRequest{
        BasicAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcTeamMappingEnabled: sdkkonnectgo.Pointer(true),
        KonnectMappingEnabled: sdkkonnectgo.Pointer(false),
        OidcIssuer: sdkkonnectgo.Pointer("https://identity.example.com/v2"),
        OidcClientID: sdkkonnectgo.Pointer("x7id0o42lklas0blidl2"),
        OidcScopes: []string{
            "email",
            "openid",
            "profile",
        },
        OidcClaimMappings: &components.PortalAuthenticationSettingsUpdateRequestPortalClaimMappings{
            Groups: sdkkonnectgo.Pointer("custom-group-claim"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuthenticationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateAuthSettingsBadRequestExample1

<!-- UsageSnippet language="go" operationID="update-portal-authentication-settings" method="patch" path="/v3/portals/{portalId}/authentication-settings" example="UpdateAuthSettingsBadRequestExample1" -->
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

    res, err := s.PortalAuthSettings.UpdatePortalAuthenticationSettings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalAuthenticationSettingsUpdateRequest{
        BasicAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcTeamMappingEnabled: sdkkonnectgo.Pointer(true),
        KonnectMappingEnabled: sdkkonnectgo.Pointer(false),
        OidcIssuer: sdkkonnectgo.Pointer("https://identity.example.com/v2"),
        OidcClientID: sdkkonnectgo.Pointer("x7id0o42lklas0blidl2"),
        OidcScopes: []string{
            "email",
            "openid",
            "profile",
        },
        OidcClaimMappings: &components.PortalAuthenticationSettingsUpdateRequestPortalClaimMappings{
            Groups: sdkkonnectgo.Pointer("custom-group-claim"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuthenticationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateAuthSettingsBadRequestExample2

<!-- UsageSnippet language="go" operationID="update-portal-authentication-settings" method="patch" path="/v3/portals/{portalId}/authentication-settings" example="UpdateAuthSettingsBadRequestExample2" -->
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

    res, err := s.PortalAuthSettings.UpdatePortalAuthenticationSettings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalAuthenticationSettingsUpdateRequest{
        BasicAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcTeamMappingEnabled: sdkkonnectgo.Pointer(true),
        KonnectMappingEnabled: sdkkonnectgo.Pointer(false),
        OidcIssuer: sdkkonnectgo.Pointer("https://identity.example.com/v2"),
        OidcClientID: sdkkonnectgo.Pointer("x7id0o42lklas0blidl2"),
        OidcScopes: []string{
            "email",
            "openid",
            "profile",
        },
        OidcClaimMappings: &components.PortalAuthenticationSettingsUpdateRequestPortalClaimMappings{
            Groups: sdkkonnectgo.Pointer("custom-group-claim"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuthenticationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateAuthSettingsBadRequestExample3

<!-- UsageSnippet language="go" operationID="update-portal-authentication-settings" method="patch" path="/v3/portals/{portalId}/authentication-settings" example="UpdateAuthSettingsBadRequestExample3" -->
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

    res, err := s.PortalAuthSettings.UpdatePortalAuthenticationSettings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalAuthenticationSettingsUpdateRequest{
        BasicAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcTeamMappingEnabled: sdkkonnectgo.Pointer(true),
        KonnectMappingEnabled: sdkkonnectgo.Pointer(false),
        OidcIssuer: sdkkonnectgo.Pointer("https://identity.example.com/v2"),
        OidcClientID: sdkkonnectgo.Pointer("x7id0o42lklas0blidl2"),
        OidcScopes: []string{
            "email",
            "openid",
            "profile",
        },
        OidcClaimMappings: &components.PortalAuthenticationSettingsUpdateRequestPortalClaimMappings{
            Groups: sdkkonnectgo.Pointer("custom-group-claim"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuthenticationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateAuthSettingsBadRequestExample4

<!-- UsageSnippet language="go" operationID="update-portal-authentication-settings" method="patch" path="/v3/portals/{portalId}/authentication-settings" example="UpdateAuthSettingsBadRequestExample4" -->
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

    res, err := s.PortalAuthSettings.UpdatePortalAuthenticationSettings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalAuthenticationSettingsUpdateRequest{
        BasicAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcTeamMappingEnabled: sdkkonnectgo.Pointer(true),
        KonnectMappingEnabled: sdkkonnectgo.Pointer(false),
        OidcIssuer: sdkkonnectgo.Pointer("https://identity.example.com/v2"),
        OidcClientID: sdkkonnectgo.Pointer("x7id0o42lklas0blidl2"),
        OidcScopes: []string{
            "email",
            "openid",
            "profile",
        },
        OidcClaimMappings: &components.PortalAuthenticationSettingsUpdateRequestPortalClaimMappings{
            Groups: sdkkonnectgo.Pointer("custom-group-claim"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuthenticationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateAuthSettingsBadRequestExample5

<!-- UsageSnippet language="go" operationID="update-portal-authentication-settings" method="patch" path="/v3/portals/{portalId}/authentication-settings" example="UpdateAuthSettingsBadRequestExample5" -->
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

    res, err := s.PortalAuthSettings.UpdatePortalAuthenticationSettings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalAuthenticationSettingsUpdateRequest{
        BasicAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcTeamMappingEnabled: sdkkonnectgo.Pointer(true),
        KonnectMappingEnabled: sdkkonnectgo.Pointer(false),
        OidcIssuer: sdkkonnectgo.Pointer("https://identity.example.com/v2"),
        OidcClientID: sdkkonnectgo.Pointer("x7id0o42lklas0blidl2"),
        OidcScopes: []string{
            "email",
            "openid",
            "profile",
        },
        OidcClaimMappings: &components.PortalAuthenticationSettingsUpdateRequestPortalClaimMappings{
            Groups: sdkkonnectgo.Pointer("custom-group-claim"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuthenticationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateAuthSettingsBadRequestExample6

<!-- UsageSnippet language="go" operationID="update-portal-authentication-settings" method="patch" path="/v3/portals/{portalId}/authentication-settings" example="UpdateAuthSettingsBadRequestExample6" -->
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

    res, err := s.PortalAuthSettings.UpdatePortalAuthenticationSettings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalAuthenticationSettingsUpdateRequest{
        BasicAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcTeamMappingEnabled: sdkkonnectgo.Pointer(true),
        KonnectMappingEnabled: sdkkonnectgo.Pointer(false),
        OidcIssuer: sdkkonnectgo.Pointer("https://identity.example.com/v2"),
        OidcClientID: sdkkonnectgo.Pointer("x7id0o42lklas0blidl2"),
        OidcScopes: []string{
            "email",
            "openid",
            "profile",
        },
        OidcClaimMappings: &components.PortalAuthenticationSettingsUpdateRequestPortalClaimMappings{
            Groups: sdkkonnectgo.Pointer("custom-group-claim"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuthenticationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateAuthSettingsBadRequestExample7

<!-- UsageSnippet language="go" operationID="update-portal-authentication-settings" method="patch" path="/v3/portals/{portalId}/authentication-settings" example="UpdateAuthSettingsBadRequestExample7" -->
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

    res, err := s.PortalAuthSettings.UpdatePortalAuthenticationSettings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalAuthenticationSettingsUpdateRequest{
        BasicAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcTeamMappingEnabled: sdkkonnectgo.Pointer(true),
        KonnectMappingEnabled: sdkkonnectgo.Pointer(false),
        OidcIssuer: sdkkonnectgo.Pointer("https://identity.example.com/v2"),
        OidcClientID: sdkkonnectgo.Pointer("x7id0o42lklas0blidl2"),
        OidcScopes: []string{
            "email",
            "openid",
            "profile",
        },
        OidcClaimMappings: &components.PortalAuthenticationSettingsUpdateRequestPortalClaimMappings{
            Groups: sdkkonnectgo.Pointer("custom-group-claim"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuthenticationSettingsResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateAuthSettingsBadRequestExample8

<!-- UsageSnippet language="go" operationID="update-portal-authentication-settings" method="patch" path="/v3/portals/{portalId}/authentication-settings" example="UpdateAuthSettingsBadRequestExample8" -->
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

    res, err := s.PortalAuthSettings.UpdatePortalAuthenticationSettings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalAuthenticationSettingsUpdateRequest{
        BasicAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcTeamMappingEnabled: sdkkonnectgo.Pointer(true),
        KonnectMappingEnabled: sdkkonnectgo.Pointer(false),
        OidcIssuer: sdkkonnectgo.Pointer("https://identity.example.com/v2"),
        OidcClientID: sdkkonnectgo.Pointer("x7id0o42lklas0blidl2"),
        OidcScopes: []string{
            "email",
            "openid",
            "profile",
        },
        OidcClaimMappings: &components.PortalAuthenticationSettingsUpdateRequestPortalClaimMappings{
            Groups: sdkkonnectgo.Pointer("custom-group-claim"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuthenticationSettingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                                                                                                                                              | Example                                                                                                                                                                                                                                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                                                                                                                          |
| `portalID`                                                                                                                                                                                                                                                                                                                                                               | *string*                                                                                                                                                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                       | ID of the portal.                                                                                                                                                                                                                                                                                                                                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                                                                                                                                                                                                                                                                                                     |
| `portalAuthenticationSettingsUpdateRequest`                                                                                                                                                                                                                                                                                                                              | [*components.PortalAuthenticationSettingsUpdateRequest](../../models/components/portalauthenticationsettingsupdaterequest.md)                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                       | Update a portal's developer authentication settings.                                                                                                                                                                                                                                                                                                                     | {<br/>"basic_auth_enabled": true,<br/>"oidc_auth_enabled": true,<br/>"oidc_team_mapping_enabled": true,<br/>"konnect_mapping_enabled": false,<br/>"oidc_issuer": "https://identity.example.com/v2",<br/>"oidc_client_id": "x7id0o42lklas0blidl2",<br/>"oidc_scopes": [<br/>"email",<br/>"openid",<br/>"profile"<br/>],<br/>"oidc_claim_mappings": {<br/>"name": "name",<br/>"email": "email",<br/>"groups": "custom-group-claim"<br/>}<br/>} |
| `opts`                                                                                                                                                                                                                                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                       | The options for this request.                                                                                                                                                                                                                                                                                                                                            |                                                                                                                                                                                                                                                                                                                                                                          |

### Response

**[*operations.UpdatePortalAuthenticationSettingsResponse](../../models/operations/updateportalauthenticationsettingsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListPortalTeamGroupMappings

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Lists mappings between Konnect portal teams and Identity Provider (IdP) groups. Returns a 400 error if an IdP has not yet been configured.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-portal-team-group-mappings" method="get" path="/v3/portals/{portalId}/identity-provider/team-group-mappings" example="Example 1" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuthSettings.ListPortalTeamGroupMappings(ctx, operations.ListPortalTeamGroupMappingsRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalTeamGroupMappingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListPortalTeamGroupMappingsRequest](../../models/operations/listportalteamgroupmappingsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ListPortalTeamGroupMappingsResponse](../../models/operations/listportalteamgroupmappingsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePortalTeamGroupMappings

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Allows partial updates to the mappings between Konnect portal teams and Identity Provider (IdP) groups. The request body must be keyed on team ID. For a given team ID, the given group list is a complete replacement. To remove all mappings for a given team, provide an empty group list.
Returns a 400 error if an IdP has not yet been configured, or if a team ID in the request body is not found or is not a UUID.

### Example Usage: Example 1

<!-- UsageSnippet language="go" operationID="update-portal-team-group-mappings" method="patch" path="/v3/portals/{portalId}/identity-provider/team-group-mappings" example="Example 1" -->
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

    res, err := s.PortalAuthSettings.UpdatePortalTeamGroupMappings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalTeamGroupMappingsUpdateRequest{
        Data: []components.PortalTeamGroupMappingsUpdateRequestData{
            components.PortalTeamGroupMappingsUpdateRequestData{
                TeamID: sdkkonnectgo.Pointer("af91db4c-6e51-403e-a2bf-33d27ae50c0a"),
                Groups: []string{
                    "Service Developer",
                },
            },
            components.PortalTeamGroupMappingsUpdateRequestData{
                TeamID: sdkkonnectgo.Pointer("bc11db4c-6e51-403e-a2bf-33d27ae50c0a"),
                Groups: []string{
                    "Service Owner",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalTeamGroupMappingResponse != nil {
        // handle response
    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="update-portal-team-group-mappings" method="patch" path="/v3/portals/{portalId}/identity-provider/team-group-mappings" example="Unauthorized" -->
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

    res, err := s.PortalAuthSettings.UpdatePortalTeamGroupMappings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalTeamGroupMappingsUpdateRequest{
        Data: []components.PortalTeamGroupMappingsUpdateRequestData{
            components.PortalTeamGroupMappingsUpdateRequestData{
                TeamID: sdkkonnectgo.Pointer("af91db4c-6e51-403e-a2bf-33d27ae50c0a"),
                Groups: []string{
                    "Service Developer",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalTeamGroupMappingResponse != nil {
        // handle response
    }
}
```
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="update-portal-team-group-mappings" method="patch" path="/v3/portals/{portalId}/identity-provider/team-group-mappings" example="UnauthorizedExample" -->
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

    res, err := s.PortalAuthSettings.UpdatePortalTeamGroupMappings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalTeamGroupMappingsUpdateRequest{
        Data: []components.PortalTeamGroupMappingsUpdateRequestData{
            components.PortalTeamGroupMappingsUpdateRequestData{
                TeamID: sdkkonnectgo.Pointer("af91db4c-6e51-403e-a2bf-33d27ae50c0a"),
                Groups: []string{
                    "Service Developer",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalTeamGroupMappingResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateTeamGroupMappingsBadRequestExample1

<!-- UsageSnippet language="go" operationID="update-portal-team-group-mappings" method="patch" path="/v3/portals/{portalId}/identity-provider/team-group-mappings" example="UpdateTeamGroupMappingsBadRequestExample1" -->
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

    res, err := s.PortalAuthSettings.UpdatePortalTeamGroupMappings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalTeamGroupMappingsUpdateRequest{
        Data: []components.PortalTeamGroupMappingsUpdateRequestData{
            components.PortalTeamGroupMappingsUpdateRequestData{
                TeamID: sdkkonnectgo.Pointer("af91db4c-6e51-403e-a2bf-33d27ae50c0a"),
                Groups: []string{
                    "Service Developer",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalTeamGroupMappingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                           | Type                                                                                                                | Required                                                                                                            | Description                                                                                                         | Example                                                                                                             |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                               | :heavy_check_mark:                                                                                                  | The context to use for the request.                                                                                 |                                                                                                                     |
| `portalID`                                                                                                          | *string*                                                                                                            | :heavy_check_mark:                                                                                                  | ID of the portal.                                                                                                   | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                                                |
| `portalTeamGroupMappingsUpdateRequest`                                                                              | [*components.PortalTeamGroupMappingsUpdateRequest](../../models/components/portalteamgroupmappingsupdaterequest.md) | :heavy_minus_sign:                                                                                                  | N/A                                                                                                                 | {<br/>"data": [<br/>{<br/>"team_id": "af91db4c-6e51-403e-a2bf-33d27ae50c0a",<br/>"groups": [<br/>"Service Developer"<br/>]<br/>}<br/>]<br/>} |
| `opts`                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                            | :heavy_minus_sign:                                                                                                  | The options for this request.                                                                                       |                                                                                                                     |

### Response

**[*operations.UpdatePortalTeamGroupMappingsResponse](../../models/operations/updateportalteamgroupmappingsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPortalIdentityProviders

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Retrieves the identity providers available within the portal. This operation provides information about
various identity providers for SAML or OIDC authentication integrations.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-portal-identity-providers" method="get" path="/v3/portals/{portalId}/identity-providers" example="Identity Provider Collection" -->
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

    res, err := s.PortalAuthSettings.GetPortalIdentityProviders(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProviders != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                       | Type                                                                                                                            | Required                                                                                                                        | Description                                                                                                                     | Example                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                           | :heavy_check_mark:                                                                                                              | The context to use for the request.                                                                                             |                                                                                                                                 |
| `portalID`                                                                                                                      | *string*                                                                                                                        | :heavy_check_mark:                                                                                                              | ID of the portal.                                                                                                               | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                                                            |
| `filter`                                                                                                                        | [*operations.GetPortalIdentityProvidersQueryParamFilter](../../models/operations/getportalidentityprovidersqueryparamfilter.md) | :heavy_minus_sign:                                                                                                              | Filter identity providers returned in the response.                                                                             |                                                                                                                                 |
| `opts`                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                        | :heavy_minus_sign:                                                                                                              | The options for this request.                                                                                                   |                                                                                                                                 |

### Response

**[*operations.GetPortalIdentityProvidersResponse](../../models/operations/getportalidentityprovidersresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreatePortalIdentityProvider

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Creates a new identity provider. This operation allows the creation of a new identity provider for
authentication purposes.

### Example Usage: OIDC Identity Provider

<!-- UsageSnippet language="go" operationID="create-portal-identity-provider" method="post" path="/v3/portals/{portalId}/identity-providers" example="OIDC Identity Provider" -->
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

    res, err := s.PortalAuthSettings.CreatePortalIdentityProvider(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", components.CreateIdentityProvider{
        Type: components.IdentityProviderTypeOidc.ToPointer(),
        LoginPath: sdkkonnectgo.Pointer("myapp"),
        Enabled: sdkkonnectgo.Pointer(true),
        Config: sdkkonnectgo.Pointer(components.CreateCreateIdentityProviderConfigSAMLIdentityProviderConfigInput(
            components.SAMLIdentityProviderConfigInput{
                IdpMetadataURL: sdkkonnectgo.Pointer("https://mocksaml.com/api/saml/metadata"),
                IdpMetadataXML: sdkkonnectgo.Pointer("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<EntityDescriptor xmlns=\"urn:oasis:names:tc:SAML:2.0:metadata\">\n  <!-- SAML metadata content here -->\n</EntityDescriptor>\n"),
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```
### Example Usage: OIDC Identity Provider Request

<!-- UsageSnippet language="go" operationID="create-portal-identity-provider" method="post" path="/v3/portals/{portalId}/identity-providers" example="OIDC Identity Provider Request" -->
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

    res, err := s.PortalAuthSettings.CreatePortalIdentityProvider(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", components.CreateIdentityProvider{
        Type: components.IdentityProviderTypeOidc.ToPointer(),
        LoginPath: sdkkonnectgo.Pointer("the-oidc-konnect-org"),
        Enabled: sdkkonnectgo.Pointer(true),
        Config: sdkkonnectgo.Pointer(components.CreateCreateIdentityProviderConfigConfigureOIDCIdentityProviderConfig(
            components.ConfigureOIDCIdentityProviderConfig{
                IssuerURL: "https://konghq.okta.com/oauth2/default",
                ClientID: "0oaqhb43ckTZ02j1F357",
                ClientSecret: sdkkonnectgo.Pointer("BbqwI8xP9E4evOK"),
                Scopes: []string{
                    "openid",
                    "email",
                    "profile",
                },
                ClaimMappings: &components.OIDCIdentityProviderClaimMappings{},
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```
### Example Usage: SAML Identity Provider Request With Metadata

<!-- UsageSnippet language="go" operationID="create-portal-identity-provider" method="post" path="/v3/portals/{portalId}/identity-providers" example="SAML Identity Provider Request With Metadata" -->
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

    res, err := s.PortalAuthSettings.CreatePortalIdentityProvider(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", components.CreateIdentityProvider{
        Type: components.IdentityProviderTypeSaml.ToPointer(),
        LoginPath: sdkkonnectgo.Pointer("the-saml-konnect-org"),
        Enabled: sdkkonnectgo.Pointer(true),
        Config: sdkkonnectgo.Pointer(components.CreateCreateIdentityProviderConfigSAMLIdentityProviderConfigInput(
            components.SAMLIdentityProviderConfigInput{
                IdpMetadataXML: sdkkonnectgo.Pointer("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<EntityDescriptor xmlns=\"urn:oasis:names:tc:SAML:2.0:metadata\">\n<!-- SAML metadata content here -->\n</EntityDescriptor>\n"),
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```
### Example Usage: SAML Identity Provider Request With MetadataURL

<!-- UsageSnippet language="go" operationID="create-portal-identity-provider" method="post" path="/v3/portals/{portalId}/identity-providers" example="SAML Identity Provider Request With MetadataURL" -->
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

    res, err := s.PortalAuthSettings.CreatePortalIdentityProvider(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", components.CreateIdentityProvider{
        Type: components.IdentityProviderTypeSaml.ToPointer(),
        LoginPath: sdkkonnectgo.Pointer("the-saml-konnect-org"),
        Enabled: sdkkonnectgo.Pointer(true),
        Config: sdkkonnectgo.Pointer(components.CreateCreateIdentityProviderConfigSAMLIdentityProviderConfigInput(
            components.SAMLIdentityProviderConfigInput{
                IdpMetadataURL: sdkkonnectgo.Pointer("https://mocksaml.com/api/saml/metadata"),
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```
### Example Usage: SAML Identity Provider With Meta Data

<!-- UsageSnippet language="go" operationID="create-portal-identity-provider" method="post" path="/v3/portals/{portalId}/identity-providers" example="SAML Identity Provider With Meta Data" -->
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

    res, err := s.PortalAuthSettings.CreatePortalIdentityProvider(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", components.CreateIdentityProvider{
        Type: components.IdentityProviderTypeOidc.ToPointer(),
        LoginPath: sdkkonnectgo.Pointer("myapp"),
        Enabled: sdkkonnectgo.Pointer(true),
        Config: sdkkonnectgo.Pointer(components.CreateCreateIdentityProviderConfigSAMLIdentityProviderConfigInput(
            components.SAMLIdentityProviderConfigInput{
                IdpMetadataURL: sdkkonnectgo.Pointer("https://mocksaml.com/api/saml/metadata"),
                IdpMetadataXML: sdkkonnectgo.Pointer("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<EntityDescriptor xmlns=\"urn:oasis:names:tc:SAML:2.0:metadata\">\n  <!-- SAML metadata content here -->\n</EntityDescriptor>\n"),
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```
### Example Usage: SAML Identity Provider With Meta Data URL

<!-- UsageSnippet language="go" operationID="create-portal-identity-provider" method="post" path="/v3/portals/{portalId}/identity-providers" example="SAML Identity Provider With Meta Data URL" -->
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

    res, err := s.PortalAuthSettings.CreatePortalIdentityProvider(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", components.CreateIdentityProvider{
        Type: components.IdentityProviderTypeOidc.ToPointer(),
        LoginPath: sdkkonnectgo.Pointer("myapp"),
        Enabled: sdkkonnectgo.Pointer(true),
        Config: sdkkonnectgo.Pointer(components.CreateCreateIdentityProviderConfigSAMLIdentityProviderConfigInput(
            components.SAMLIdentityProviderConfigInput{
                IdpMetadataURL: sdkkonnectgo.Pointer("https://mocksaml.com/api/saml/metadata"),
                IdpMetadataXML: sdkkonnectgo.Pointer("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<EntityDescriptor xmlns=\"urn:oasis:names:tc:SAML:2.0:metadata\">\n  <!-- SAML metadata content here -->\n</EntityDescriptor>\n"),
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="create-portal-identity-provider" method="post" path="/v3/portals/{portalId}/identity-providers" example="Unauthorized" -->
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

    res, err := s.PortalAuthSettings.CreatePortalIdentityProvider(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", components.CreateIdentityProvider{
        Type: components.IdentityProviderTypeOidc.ToPointer(),
        LoginPath: sdkkonnectgo.Pointer("myapp"),
        Enabled: sdkkonnectgo.Pointer(true),
        Config: sdkkonnectgo.Pointer(components.CreateCreateIdentityProviderConfigSAMLIdentityProviderConfigInput(
            components.SAMLIdentityProviderConfigInput{
                IdpMetadataURL: sdkkonnectgo.Pointer("https://mocksaml.com/api/saml/metadata"),
                IdpMetadataXML: sdkkonnectgo.Pointer("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<EntityDescriptor xmlns=\"urn:oasis:names:tc:SAML:2.0:metadata\">\n  <!-- SAML metadata content here -->\n</EntityDescriptor>\n"),
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="create-portal-identity-provider" method="post" path="/v3/portals/{portalId}/identity-providers" example="UnauthorizedExample" -->
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

    res, err := s.PortalAuthSettings.CreatePortalIdentityProvider(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", components.CreateIdentityProvider{
        Type: components.IdentityProviderTypeOidc.ToPointer(),
        LoginPath: sdkkonnectgo.Pointer("myapp"),
        Enabled: sdkkonnectgo.Pointer(true),
        Config: sdkkonnectgo.Pointer(components.CreateCreateIdentityProviderConfigSAMLIdentityProviderConfigInput(
            components.SAMLIdentityProviderConfigInput{
                IdpMetadataURL: sdkkonnectgo.Pointer("https://mocksaml.com/api/saml/metadata"),
                IdpMetadataXML: sdkkonnectgo.Pointer("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<EntityDescriptor xmlns=\"urn:oasis:names:tc:SAML:2.0:metadata\">\n  <!-- SAML metadata content here -->\n</EntityDescriptor>\n"),
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    | Example                                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |                                                                                                                                                                |
| `portalID`                                                                                                                                                     | *string*                                                                                                                                                       | :heavy_check_mark:                                                                                                                                             | ID of the portal.                                                                                                                                              | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                                                                                           |
| `createIdentityProvider`                                                                                                                                       | [components.CreateIdentityProvider](../../models/components/createidentityprovider.md)                                                                         | :heavy_check_mark:                                                                                                                                             | An object representing the configuration for creating a new identity provider. This configuration may pertain  to either an OIDC or a SAML identity provider.<br/> |                                                                                                                                                                |
| `opts`                                                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                                                       | :heavy_minus_sign:                                                                                                                                             | The options for this request.                                                                                                                                  |                                                                                                                                                                |

### Response

**[*operations.CreatePortalIdentityProviderResponse](../../models/operations/createportalidentityproviderresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPortalIdentityProvider

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Retrieves the configuration of a single identity provider. This operation returns information about a
specific identity provider's settings and authentication integration details.

### Example Usage: OIDC Identity Provider

<!-- UsageSnippet language="go" operationID="get-portal-identity-provider" method="get" path="/v3/portals/{portalId}/identity-providers/{id}" example="OIDC Identity Provider" -->
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

    res, err := s.PortalAuthSettings.GetPortalIdentityProvider(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", "d32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```
### Example Usage: SAML Identity Provider With Meta Data

<!-- UsageSnippet language="go" operationID="get-portal-identity-provider" method="get" path="/v3/portals/{portalId}/identity-providers/{id}" example="SAML Identity Provider With Meta Data" -->
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

    res, err := s.PortalAuthSettings.GetPortalIdentityProvider(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", "d32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```
### Example Usage: SAML Identity Provider With Meta Data URL

<!-- UsageSnippet language="go" operationID="get-portal-identity-provider" method="get" path="/v3/portals/{portalId}/identity-providers/{id}" example="SAML Identity Provider With Meta Data URL" -->
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

    res, err := s.PortalAuthSettings.GetPortalIdentityProvider(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", "d32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `portalID`                                               | *string*                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | ID of the identity provider.                             | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPortalIdentityProviderResponse](../../models/operations/getportalidentityproviderresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePortalIdentityProvider

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Updates the configuration of an existing identity provider. This operation allows modifications to be made
to an existing identity provider's configuration.

### Example Usage: NotFoundExample

<!-- UsageSnippet language="go" operationID="update-portal-identity-provider" method="patch" path="/v3/portals/{portalId}/identity-providers/{id}" example="NotFoundExample" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuthSettings.UpdatePortalIdentityProvider(ctx, operations.UpdatePortalIdentityProviderRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        UpdateIdentityProvider: components.UpdateIdentityProvider{
            Enabled: sdkkonnectgo.Pointer(true),
            LoginPath: sdkkonnectgo.Pointer("myapp"),
            Config: sdkkonnectgo.Pointer(components.CreateUpdateIdentityProviderConfigSAMLIdentityProviderConfigInput(
                components.SAMLIdentityProviderConfigInput{
                    IdpMetadataURL: sdkkonnectgo.Pointer("https://mocksaml.com/api/saml/metadata"),
                    IdpMetadataXML: sdkkonnectgo.Pointer("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<EntityDescriptor xmlns=\"urn:oasis:names:tc:SAML:2.0:metadata\">\n  <!-- SAML metadata content here -->\n</EntityDescriptor>\n"),
                },
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```
### Example Usage: OIDC Identity Provider

<!-- UsageSnippet language="go" operationID="update-portal-identity-provider" method="patch" path="/v3/portals/{portalId}/identity-providers/{id}" example="OIDC Identity Provider" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuthSettings.UpdatePortalIdentityProvider(ctx, operations.UpdatePortalIdentityProviderRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        UpdateIdentityProvider: components.UpdateIdentityProvider{
            Enabled: sdkkonnectgo.Pointer(true),
            LoginPath: sdkkonnectgo.Pointer("myapp"),
            Config: sdkkonnectgo.Pointer(components.CreateUpdateIdentityProviderConfigSAMLIdentityProviderConfigInput(
                components.SAMLIdentityProviderConfigInput{
                    IdpMetadataURL: sdkkonnectgo.Pointer("https://mocksaml.com/api/saml/metadata"),
                    IdpMetadataXML: sdkkonnectgo.Pointer("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<EntityDescriptor xmlns=\"urn:oasis:names:tc:SAML:2.0:metadata\">\n  <!-- SAML metadata content here -->\n</EntityDescriptor>\n"),
                },
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```
### Example Usage: SAML Identity Provider With Meta Data

<!-- UsageSnippet language="go" operationID="update-portal-identity-provider" method="patch" path="/v3/portals/{portalId}/identity-providers/{id}" example="SAML Identity Provider With Meta Data" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuthSettings.UpdatePortalIdentityProvider(ctx, operations.UpdatePortalIdentityProviderRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        UpdateIdentityProvider: components.UpdateIdentityProvider{
            Enabled: sdkkonnectgo.Pointer(true),
            LoginPath: sdkkonnectgo.Pointer("myapp"),
            Config: sdkkonnectgo.Pointer(components.CreateUpdateIdentityProviderConfigSAMLIdentityProviderConfigInput(
                components.SAMLIdentityProviderConfigInput{
                    IdpMetadataURL: sdkkonnectgo.Pointer("https://mocksaml.com/api/saml/metadata"),
                    IdpMetadataXML: sdkkonnectgo.Pointer("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<EntityDescriptor xmlns=\"urn:oasis:names:tc:SAML:2.0:metadata\">\n  <!-- SAML metadata content here -->\n</EntityDescriptor>\n"),
                },
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```
### Example Usage: SAML Identity Provider With Meta Data URL

<!-- UsageSnippet language="go" operationID="update-portal-identity-provider" method="patch" path="/v3/portals/{portalId}/identity-providers/{id}" example="SAML Identity Provider With Meta Data URL" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuthSettings.UpdatePortalIdentityProvider(ctx, operations.UpdatePortalIdentityProviderRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        UpdateIdentityProvider: components.UpdateIdentityProvider{
            Enabled: sdkkonnectgo.Pointer(true),
            LoginPath: sdkkonnectgo.Pointer("myapp"),
            Config: sdkkonnectgo.Pointer(components.CreateUpdateIdentityProviderConfigSAMLIdentityProviderConfigInput(
                components.SAMLIdentityProviderConfigInput{
                    IdpMetadataURL: sdkkonnectgo.Pointer("https://mocksaml.com/api/saml/metadata"),
                    IdpMetadataXML: sdkkonnectgo.Pointer("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<EntityDescriptor xmlns=\"urn:oasis:names:tc:SAML:2.0:metadata\">\n  <!-- SAML metadata content here -->\n</EntityDescriptor>\n"),
                },
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="update-portal-identity-provider" method="patch" path="/v3/portals/{portalId}/identity-providers/{id}" example="Unauthorized" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuthSettings.UpdatePortalIdentityProvider(ctx, operations.UpdatePortalIdentityProviderRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        UpdateIdentityProvider: components.UpdateIdentityProvider{
            Enabled: sdkkonnectgo.Pointer(true),
            LoginPath: sdkkonnectgo.Pointer("myapp"),
            Config: sdkkonnectgo.Pointer(components.CreateUpdateIdentityProviderConfigSAMLIdentityProviderConfigInput(
                components.SAMLIdentityProviderConfigInput{
                    IdpMetadataURL: sdkkonnectgo.Pointer("https://mocksaml.com/api/saml/metadata"),
                    IdpMetadataXML: sdkkonnectgo.Pointer("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<EntityDescriptor xmlns=\"urn:oasis:names:tc:SAML:2.0:metadata\">\n  <!-- SAML metadata content here -->\n</EntityDescriptor>\n"),
                },
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="update-portal-identity-provider" method="patch" path="/v3/portals/{portalId}/identity-providers/{id}" example="UnauthorizedExample" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuthSettings.UpdatePortalIdentityProvider(ctx, operations.UpdatePortalIdentityProviderRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        UpdateIdentityProvider: components.UpdateIdentityProvider{
            Enabled: sdkkonnectgo.Pointer(true),
            LoginPath: sdkkonnectgo.Pointer("myapp"),
            Config: sdkkonnectgo.Pointer(components.CreateUpdateIdentityProviderConfigSAMLIdentityProviderConfigInput(
                components.SAMLIdentityProviderConfigInput{
                    IdpMetadataURL: sdkkonnectgo.Pointer("https://mocksaml.com/api/saml/metadata"),
                    IdpMetadataXML: sdkkonnectgo.Pointer("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<EntityDescriptor xmlns=\"urn:oasis:names:tc:SAML:2.0:metadata\">\n  <!-- SAML metadata content here -->\n</EntityDescriptor>\n"),
                },
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```
### Example Usage: Update Identity Provider Login Path

<!-- UsageSnippet language="go" operationID="update-portal-identity-provider" method="patch" path="/v3/portals/{portalId}/identity-providers/{id}" example="Update Identity Provider Login Path" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuthSettings.UpdatePortalIdentityProvider(ctx, operations.UpdatePortalIdentityProviderRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        UpdateIdentityProvider: components.UpdateIdentityProvider{
            Enabled: sdkkonnectgo.Pointer(true),
            LoginPath: sdkkonnectgo.Pointer("the-konnect-org"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```
### Example Usage: Update OIDC Identity Provider

<!-- UsageSnippet language="go" operationID="update-portal-identity-provider" method="patch" path="/v3/portals/{portalId}/identity-providers/{id}" example="Update OIDC Identity Provider" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuthSettings.UpdatePortalIdentityProvider(ctx, operations.UpdatePortalIdentityProviderRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        UpdateIdentityProvider: components.UpdateIdentityProvider{
            Enabled: sdkkonnectgo.Pointer(true),
            LoginPath: sdkkonnectgo.Pointer("the-oidc-konnect-org"),
            Config: sdkkonnectgo.Pointer(components.CreateUpdateIdentityProviderConfigConfigureOIDCIdentityProviderConfig(
                components.ConfigureOIDCIdentityProviderConfig{
                    IssuerURL: "https://konghq.okta.com/oauth2/default",
                    ClientID: "0oaqhb43ckTZ02j1F357",
                    ClientSecret: sdkkonnectgo.Pointer("BbqwI8xP9E4evOK"),
                    Scopes: []string{
                        "openid",
                        "email",
                        "profile",
                    },
                    ClaimMappings: &components.OIDCIdentityProviderClaimMappings{},
                },
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```
### Example Usage: Update SAML Identity Provider

<!-- UsageSnippet language="go" operationID="update-portal-identity-provider" method="patch" path="/v3/portals/{portalId}/identity-providers/{id}" example="Update SAML Identity Provider" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuthSettings.UpdatePortalIdentityProvider(ctx, operations.UpdatePortalIdentityProviderRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        UpdateIdentityProvider: components.UpdateIdentityProvider{
            Enabled: sdkkonnectgo.Pointer(true),
            LoginPath: sdkkonnectgo.Pointer("the-saml-konnect-org"),
            Config: sdkkonnectgo.Pointer(components.CreateUpdateIdentityProviderConfigSAMLIdentityProviderConfigInput(
                components.SAMLIdentityProviderConfigInput{
                    IdpMetadataURL: sdkkonnectgo.Pointer("https://mocksaml.com/api/saml/metadata"),
                },
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
        switch res.IdentityProvider.Config.Type {
            case components.IdentityProviderConfigTypeOIDCIdentityProviderConfig:
                // res.IdentityProvider.Config.OIDCIdentityProviderConfig is populated
            case components.IdentityProviderConfigTypeSAMLIdentityProviderConfig:
                // res.IdentityProvider.Config.SAMLIdentityProviderConfig is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.UpdatePortalIdentityProviderRequest](../../models/operations/updateportalidentityproviderrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.UpdatePortalIdentityProviderResponse](../../models/operations/updateportalidentityproviderresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeletePortalIdentityProvider

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Deletes an existing identity provider configuration. This operation removes a specific identity provider
from the portal.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-portal-identity-provider" method="delete" path="/v3/portals/{portalId}/identity-providers/{id}" -->
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

    res, err := s.PortalAuthSettings.DeletePortalIdentityProvider(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", "d32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `portalID`                                               | *string*                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | ID of the identity provider.                             | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeletePortalIdentityProviderResponse](../../models/operations/deleteportalidentityproviderresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |