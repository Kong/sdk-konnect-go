# PortalAuthSettings
(*PortalAuthSettings*)

## Overview

APIs related to configuration of Konnect Developer Portal auth settings.

### Available Operations

* [GetPortalAuthenticationSettings](#getportalauthenticationsettings) - Get Auth Settings
* [UpdatePortalAuthenticationSettings](#updateportalauthenticationsettings) - Update Auth Settings
* [ListPortalTeamGroupMappings](#listportalteamgroupmappings) - List Team Group Mappings
* [UpdatePortalTeamGroupMappings](#updateportalteamgroupmappings) - Update Team Group Mappings
* [GetPortalIdentityProviders](#getportalidentityproviders) - Retrieve Identity Providers
* [CreatePortalIdentityProvider](#createportalidentityprovider) - Create Identity Provider
* [GetPortalIdentityProvider](#getportalidentityprovider) - Get Identity Provider
* [UpdatePortalIdentityProvider](#updateportalidentityprovider) - Update Identity Provider
* [DeletePortalIdentityProvider](#deleteportalidentityprovider) - Delete Identity Provider

## GetPortalAuthenticationSettings

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns the developer authentication configuration for a portal, which determines how developers can log in and how they are assigned to teams.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
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

Updates the developer authentication configuration for a portal. Developers can be allowed to login using basic auth (email & password) or use Single-Sign-On (SSO) through an OIDC Identity Provider (IdP). Developers can be automatically assigned to teams by mapping claims from thier IdP account.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuthSettings.UpdatePortalAuthenticationSettings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalAuthenticationSettingsUpdateRequest{
        BasicAuthEnabled: sdkkonnectgo.Bool(true),
        OidcAuthEnabled: sdkkonnectgo.Bool(true),
        SamlAuthEnabled: sdkkonnectgo.Bool(true),
        OidcTeamMappingEnabled: sdkkonnectgo.Bool(true),
        KonnectMappingEnabled: sdkkonnectgo.Bool(false),
        IdpMappingEnabled: sdkkonnectgo.Bool(true),
        OidcIssuer: sdkkonnectgo.String("https://identity.example.com/v2"),
        OidcClientID: sdkkonnectgo.String("x7id0o42lklas0blidl2"),
        OidcScopes: []string{
            "email",
            "openid",
            "profile",
        },
        OidcClaimMappings: &components.PortalClaimMappings{
            Groups: sdkkonnectgo.String("custom-group-claim"),
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

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuthSettings.ListPortalTeamGroupMappings(ctx, operations.ListPortalTeamGroupMappingsRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        PageSize: sdkkonnectgo.Int64(10),
        PageNumber: sdkkonnectgo.Int64(1),
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

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuthSettings.UpdatePortalTeamGroupMappings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalTeamGroupMappingsUpdateRequest{
        Data: []components.PortalTeamGroupMappingsUpdateRequestData{
            components.PortalTeamGroupMappingsUpdateRequestData{
                TeamID: sdkkonnectgo.String("af91db4c-6e51-403e-a2bf-33d27ae50c0a"),
                Groups: []string{
                    "Service Developer",
                },
            },
            components.PortalTeamGroupMappingsUpdateRequestData{
                TeamID: sdkkonnectgo.String("bc11db4c-6e51-403e-a2bf-33d27ae50c0a"),
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

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
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

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuthSettings.CreatePortalIdentityProvider(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", components.CreateIdentityProvider{
        Type: components.IdentityProviderTypeOidc.ToPointer(),
        LoginPath: sdkkonnectgo.String("the-oidc-konnect-org"),
        Enabled: sdkkonnectgo.Bool(true),
        Config: sdkkonnectgo.Pointer(components.CreateCreateIdentityProviderConfigConfigureOIDCIdentityProviderConfig(
            components.ConfigureOIDCIdentityProviderConfig{
                IssuerURL: "https://konghq.okta.com/oauth2/default",
                ClientID: "0oaqhb43ckTZ02j1F357",
                ClientSecret: sdkkonnectgo.String("BbqwI8xP9E4evOK"),
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
        // handle response
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

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuthSettings.GetPortalIdentityProvider(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", "d32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProvider != nil {
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

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuthSettings.UpdatePortalIdentityProvider(ctx, operations.UpdatePortalIdentityProviderRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        UpdateIdentityProvider: components.UpdateIdentityProvider{
            Enabled: sdkkonnectgo.Bool(true),
            LoginPath: sdkkonnectgo.String("the-oidc-konnect-org"),
            Config: sdkkonnectgo.Pointer(components.CreateUpdateIdentityProviderConfigConfigureOIDCIdentityProviderConfig(
                components.ConfigureOIDCIdentityProviderConfig{
                    IssuerURL: "https://konghq.okta.com/oauth2/default",
                    ClientID: "0oaqhb43ckTZ02j1F357",
                    ClientSecret: sdkkonnectgo.String("BbqwI8xP9E4evOK"),
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
        // handle response
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

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
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