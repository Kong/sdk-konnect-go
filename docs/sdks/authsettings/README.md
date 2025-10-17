# AuthSettings
(*AuthSettings*)

## Overview

### Available Operations

* [GetAuthenticationSettings](#getauthenticationsettings) - Get Auth Settings
* [UpdateAuthenticationSettings](#updateauthenticationsettings) - Update Auth Settings
* [GetIdpConfiguration](#getidpconfiguration) - Fetch IdP Configuration
* [UpdateIdpConfiguration](#updateidpconfiguration) - Update IdP Configuration
* [GetTeamGroupMappings](#getteamgroupmappings) - Fetch Team Group Mappings
* [PatchTeamGroupMappings](#patchteamgroupmappings) - Patch Mappings by Team ID
* [UpdateIdpTeamMappings](#updateidpteammappings) - Update Team Mappings
* [GetIdpTeamMappings](#getidpteammappings) - Fetch Team Mapping
* [GetIdentityProviders](#getidentityproviders) - Retrieve Identity Providers
* [CreateIdentityProvider](#createidentityprovider) - Create Identity Provider
* [GetIdentityProvider](#getidentityprovider) - Get Identity Provider
* [UpdateIdentityProvider](#updateidentityprovider) - Update Identity Provider
* [DeleteIdentityProvider](#deleteidentityprovider) - Delete Identity Provider

## GetAuthenticationSettings

Returns authentication configuration, which determines how users can log in and how they are assigned to teams.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-authentication-settings" method="get" path="/v3/authentication-settings" -->
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

    res, err := s.AuthSettings.GetAuthenticationSettings(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthenticationSettings != nil {
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

**[*operations.GetAuthenticationSettingsResponse](../../models/operations/getauthenticationsettingsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateAuthenticationSettings

Updates authentication configuration.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-authentication-settings" method="patch" path="/v3/authentication-settings" -->
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

    res, err := s.AuthSettings.UpdateAuthenticationSettings(ctx, &components.UpdateAuthenticationSettings{
        BasicAuthEnabled: sdkkonnectgo.Pointer(true),
        OidcAuthEnabled: sdkkonnectgo.Pointer(false),
        SamlAuthEnabled: sdkkonnectgo.Pointer(false),
        IdpMappingEnabled: sdkkonnectgo.Pointer(true),
        KonnectMappingEnabled: sdkkonnectgo.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthenticationSettings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [components.UpdateAuthenticationSettings](../../models/components/updateauthenticationsettings.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateAuthenticationSettingsResponse](../../models/operations/updateauthenticationsettingsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetIdpConfiguration

Fetch the IdP configuration.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-idp-configuration" method="get" path="/v3/identity-provider" -->
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

    res, err := s.AuthSettings.GetIdpConfiguration(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Idp != nil {
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

**[*operations.GetIdpConfigurationResponse](../../models/operations/getidpconfigurationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateIdpConfiguration

Update the IdP configuration.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-idp-configuration" method="patch" path="/v3/identity-provider" -->
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

    res, err := s.AuthSettings.UpdateIdpConfiguration(ctx, &components.UpdateIDPConfiguration{
        Issuer: sdkkonnectgo.Pointer("https://myidp.com/oauth2"),
        LoginPath: sdkkonnectgo.Pointer("myapp"),
        ClientID: sdkkonnectgo.Pointer("YOUR_CLIENT_ID"),
        ClientSecret: sdkkonnectgo.Pointer("YOUR_CLIENT_SECRET"),
        ClaimMappings: &components.UpdateIDPConfigurationClaimMappings{
            Groups: sdkkonnectgo.Pointer("custom-group-claim"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Idp != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.UpdateIDPConfiguration](../../models/components/updateidpconfiguration.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdateIdpConfigurationResponse](../../models/operations/updateidpconfigurationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetTeamGroupMappings

Retrieves the mappings between Konnect Teams and Identity Provider Groups.
Returns a 400 error if an Identity Provider has not yet been configured.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-team-group-mappings" method="get" path="/v3/identity-provider/team-group-mappings" -->
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

    res, err := s.AuthSettings.GetTeamGroupMappings(ctx, sdkkonnectgo.Pointer[int64](10), sdkkonnectgo.Pointer[int64](1))
    if err != nil {
        log.Fatal(err)
    }
    if res.TeamGroupMappingCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             | Example                                                                                                 |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |                                                                                                         |
| `pageSize`                                                                                              | **int64*                                                                                                | :heavy_minus_sign:                                                                                      | The maximum number of items to include per page. The last page of a collection may include fewer items. | 10                                                                                                      |
| `pageNumber`                                                                                            | **int64*                                                                                                | :heavy_minus_sign:                                                                                      | Determines which page of the entities to retrieve.                                                      | 1                                                                                                       |
| `opts`                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                | :heavy_minus_sign:                                                                                      | The options for this request.                                                                           |                                                                                                         |

### Response

**[*operations.GetTeamGroupMappingsResponse](../../models/operations/getteamgroupmappingsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchTeamGroupMappings

Allows partial updates to the mappings between Konnect Teams and Identity Provider Groups.
The request body must be keyed on team ID. For a given team ID, the given group list is a
complete replacement. To remove all mappings for a given team, provide an empty group list.

Returns a 400 error if an Identity Provider has not yet been configured, or if a team ID in
the request body is not found or is not a UUID.

### Example Usage

<!-- UsageSnippet language="go" operationID="patch-team-group-mappings" method="patch" path="/v3/identity-provider/team-group-mappings" -->
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

    res, err := s.AuthSettings.PatchTeamGroupMappings(ctx, &components.PatchTeamGroupMappings{
        Data: []components.Data{
            components.Data{
                TeamID: sdkkonnectgo.Pointer("af91db4c-6e51-403e-a2bf-33d27ae50c0a"),
                Groups: []string{
                    "Service Developers",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TeamGroupMappingCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.PatchTeamGroupMappings](../../models/components/patchteamgroupmappings.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.PatchTeamGroupMappingsResponse](../../models/operations/patchteamgroupmappingsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateIdpTeamMappings

Updates the IdP group to Konnect team mapping.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-idp-team-mappings" method="put" path="/v3/identity-provider/team-mappings" -->
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

    res, err := s.AuthSettings.UpdateIdpTeamMappings(ctx, &components.UpdateTeamMappings{
        Mappings: []components.Mappings{
            components.Mappings{
                Group: sdkkonnectgo.Pointer("Service Developers"),
                TeamIds: []string{
                    "af91db4c-6e51-403e-a2bf-33d27ae50c0a",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TeamMappingCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.UpdateTeamMappings](../../models/components/updateteammappings.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.UpdateIdpTeamMappingsResponse](../../models/operations/updateidpteammappingsresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.BadRequestError    | 400                          | application/problem+json     |
| sdkerrors.ForbiddenError     | 401                          | application/problem+json     |
| sdkerrors.NotFoundError      | 404                          | application/problem+json     |
| sdkerrors.PreconditionFailed | 412                          | application/problem+json     |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## GetIdpTeamMappings

Fetch the IdP group to Konnect team mapping.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-idp-team-mappings" method="get" path="/v3/identity-provider/team-mappings" -->
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

    res, err := s.AuthSettings.GetIdpTeamMappings(ctx, sdkkonnectgo.Pointer[int64](10), sdkkonnectgo.Pointer[int64](1))
    if err != nil {
        log.Fatal(err)
    }
    if res.TeamMappingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             | Example                                                                                                 |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |                                                                                                         |
| `pageSize`                                                                                              | **int64*                                                                                                | :heavy_minus_sign:                                                                                      | The maximum number of items to include per page. The last page of a collection may include fewer items. | 10                                                                                                      |
| `pageNumber`                                                                                            | **int64*                                                                                                | :heavy_minus_sign:                                                                                      | Determines which page of the entities to retrieve.                                                      | 1                                                                                                       |
| `opts`                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                | :heavy_minus_sign:                                                                                      | The options for this request.                                                                           |                                                                                                         |

### Response

**[*operations.GetIdpTeamMappingsResponse](../../models/operations/getidpteammappingsresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.UnauthorizedError  | 401                          | application/problem+json     |
| sdkerrors.ForbiddenError     | 403                          | application/problem+json     |
| sdkerrors.NotFoundError      | 404                          | application/problem+json     |
| sdkerrors.PreconditionFailed | 412                          | application/problem+json     |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## GetIdentityProviders

Retrieves the identity providers available within the organization. This operation provides information about 
various identity providers for SAML or OIDC authentication integrations.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-identity-providers" method="get" path="/v3/identity-providers" -->
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

    res, err := s.AuthSettings.GetIdentityProviders(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.IdentityProviders != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                   | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ctx`                                                                       | [context.Context](https://pkg.go.dev/context#Context)                       | :heavy_check_mark:                                                          | The context to use for the request.                                         |
| `filter`                                                                    | [*operations.QueryParamFilter](../../models/operations/queryparamfilter.md) | :heavy_minus_sign:                                                          | Filter identity providers returned in the response.                         |
| `opts`                                                                      | [][operations.Option](../../models/operations/option.md)                    | :heavy_minus_sign:                                                          | The options for this request.                                               |

### Response

**[*operations.GetIdentityProvidersResponse](../../models/operations/getidentityprovidersresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateIdentityProvider

Creates a new identity provider. This operation allows the creation of a new identity provider for 
authentication purposes.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-identity-provider" method="post" path="/v3/identity-providers" -->
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

    res, err := s.AuthSettings.CreateIdentityProvider(ctx, components.CreateIdentityProvider{
        Type: components.IdentityProviderTypeOidc.ToPointer(),
        LoginPath: sdkkonnectgo.Pointer("myapp"),
        Enabled: sdkkonnectgo.Pointer(true),
        Config: sdkkonnectgo.Pointer(components.CreateCreateIdentityProviderConfigSAMLIdentityProviderConfigInput(
            components.SAMLIdentityProviderConfigInput{
                IdpMetadataURL: sdkkonnectgo.Pointer("https://mocksaml.com/api/saml/metadata"),
                IdpMetadataXML: sdkkonnectgo.Pointer("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
                "<EntityDescriptor xmlns=\"urn:oasis:names:tc:SAML:2.0:metadata\">\n" +
                "  <!-- SAML metadata content here -->\n" +
                "</EntityDescriptor>\n" +
                ""),
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.CreateIdentityProvider](../../models/components/createidentityprovider.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateIdentityProviderResponse](../../models/operations/createidentityproviderresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetIdentityProvider

Retrieves the configuration of a single identity provider. This operation returns information about a 
specific identity provider's settings and authentication integration details.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-identity-provider" method="get" path="/v3/identity-providers/{id}" -->
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

    res, err := s.AuthSettings.GetIdentityProvider(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | ID of the identity provider.                             | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetIdentityProviderResponse](../../models/operations/getidentityproviderresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateIdentityProvider

Updates the configuration of an existing identity provider. This operation allows modifications to be made 
to an existing identity provider's configuration.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-identity-provider" method="patch" path="/v3/identity-providers/{id}" -->
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

    res, err := s.AuthSettings.UpdateIdentityProvider(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.UpdateIdentityProvider{
        Enabled: sdkkonnectgo.Pointer(true),
        LoginPath: sdkkonnectgo.Pointer("myapp"),
        Config: sdkkonnectgo.Pointer(components.CreateUpdateIdentityProviderConfigConfigureOIDCIdentityProviderConfig(
            components.ConfigureOIDCIdentityProviderConfig{
                IssuerURL: "https://konghq.okta.com/oauth2/default",
                ClientID: "YOUR_CLIENT_ID",
                ClientSecret: sdkkonnectgo.Pointer("YOUR_CLIENT_SECRET"),
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

| Parameter                                                                                                                                                   | Type                                                                                                                                                        | Required                                                                                                                                                    | Description                                                                                                                                                 | Example                                                                                                                                                     |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                                                       | :heavy_check_mark:                                                                                                                                          | The context to use for the request.                                                                                                                         |                                                                                                                                                             |
| `id`                                                                                                                                                        | *string*                                                                                                                                                    | :heavy_check_mark:                                                                                                                                          | ID of the identity provider.                                                                                                                                | d32d905a-ed33-46a3-a093-d8f536af9a8a                                                                                                                        |
| `updateIdentityProvider`                                                                                                                                    | [components.UpdateIdentityProvider](../../models/components/updateidentityprovider.md)                                                                      | :heavy_check_mark:                                                                                                                                          | An object representing the configuration for updating an identity provider. This configuration may pertain  to either an OIDC or a SAML identity provider.<br/> |                                                                                                                                                             |
| `opts`                                                                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                                                                    | :heavy_minus_sign:                                                                                                                                          | The options for this request.                                                                                                                               |                                                                                                                                                             |

### Response

**[*operations.UpdateIdentityProviderResponse](../../models/operations/updateidentityproviderresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteIdentityProvider

Deletes an existing identity provider configuration. This operation removes a specific identity provider 
from the organization.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-identity-provider" method="delete" path="/v3/identity-providers/{id}" -->
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

    res, err := s.AuthSettings.DeleteIdentityProvider(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | ID of the identity provider.                             | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteIdentityProviderResponse](../../models/operations/deleteidentityproviderresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |