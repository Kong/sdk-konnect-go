# TokenVaultProviders

## Overview

Vault providers configure credential providers for token-vault enabled directories.


### Available Operations

* [ListCredentialProviderTemplates](#listcredentialprovidertemplates) - List credential provider templates
* [GetCredentialProviderTemplate](#getcredentialprovidertemplate) - Get a credential provider template
* [CreateVaultProvider](#createvaultprovider) - Create a vault provider
* [ListVaultProviders](#listvaultproviders) - List vault providers
* [GetVaultProvider](#getvaultprovider) - Get a vault provider
* [UpdateVaultProvider](#updatevaultprovider) - Update a vault provider
* [ReplaceVaultProvider](#replacevaultprovider) - Replace a vault provider
* [DeleteVaultProvider](#deletevaultprovider) - Delete a vault provider
* [CreateVaultProviderCredential](#createvaultprovidercredential) - Create a shared credential for a provider
* [ListVaultProviderCredentials](#listvaultprovidercredentials) - List credential metadata for a provider
* [GetVaultProviderCredential](#getvaultprovidercredential) - Get a specific credential
* [UpdateVaultProviderCredential](#updatevaultprovidercredential) - Update a shared credential's value
* [DeleteVaultProviderCredential](#deletevaultprovidercredential) - Delete a specific provider credential

## ListCredentialProviderTemplates

Returns a paginated list of vault credential provider templates available to base a provider on.


### Example Usage

<!-- UsageSnippet language="go" operationID="listCredentialProviderTemplates" method="get" path="/v2/credential-provider-templates" -->
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

    res, err := s.TokenVaultProviders.ListCredentialProviderTemplates(ctx, &components.CursorPageParameters{
        Size: sdkkonnectgo.Pointer[int64](10),
        After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCredentialProviderTemplates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ctx`                                                                               | [context.Context](https://pkg.go.dev/context#Context)                               | :heavy_check_mark:                                                                  | The context to use for the request.                                                 |
| `page`                                                                              | [*components.CursorPageParameters](../../models/components/cursorpageparameters.md) | :heavy_minus_sign:                                                                  | Determines which page of the collection to retrieve.                                |
| `opts`                                                                              | [][operations.Option](../../models/operations/option.md)                            | :heavy_minus_sign:                                                                  | The options for this request.                                                       |

### Response

**[*operations.ListCredentialProviderTemplatesResponse](../../models/operations/listcredentialprovidertemplatesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCredentialProviderTemplate

Retrieves a single vault credential provider template by its name.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCredentialProviderTemplate" method="get" path="/v2/credential-provider-templates/{credentialProviderTemplateName}" -->
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

    res, err := s.TokenVaultProviders.GetCredentialProviderTemplate(ctx, components.ProviderTemplateNameGithub)
    if err != nil {
        log.Fatal(err)
    }
    if res.CredentialProviderTemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `credentialProviderTemplateName`                                                   | [components.ProviderTemplateName](../../models/components/providertemplatename.md) | :heavy_check_mark:                                                                 | Name of the vault credential provider template                                     | github                                                                             |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetCredentialProviderTemplateResponse](../../models/operations/getcredentialprovidertemplateresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateVaultProvider

Creates a new credential provider configuration for a vault-enabled directory.


### Example Usage

<!-- UsageSnippet language="go" operationID="createVaultProvider" method="post" path="/v2/directories/{directoryId}/vault/providers" -->
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

    res, err := s.TokenVaultProviders.CreateVaultProvider(ctx, "6ca3a97b-7e20-4d27-be51-5ae2e053040e", components.CreateVaultProviderBody{
        TemplateName: components.ProviderTemplateNameGithub,
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VaultProvider != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `directoryID`                                                                            | `string`                                                                                 | :heavy_check_mark:                                                                       | ID of the directory                                                                      |
| `createVaultProviderBody`                                                                | [components.CreateVaultProviderBody](../../models/components/createvaultproviderbody.md) | :heavy_check_mark:                                                                       | Vault provider to be created                                                             |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateVaultProviderResponse](../../models/operations/createvaultproviderresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListVaultProviders

Returns a paginated list of credential providers for a vault-enabled directory.


### Example Usage

<!-- UsageSnippet language="go" operationID="listVaultProviders" method="get" path="/v2/directories/{directoryId}/vault/providers" -->
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

    res, err := s.TokenVaultProviders.ListVaultProviders(ctx, operations.ListVaultProvidersRequest{
        DirectoryID: "8d4ff0c2-b53d-46ec-9f39-96c4b783f5e1",
        Page: &components.CursorPageParameters{
            Size: sdkkonnectgo.Pointer[int64](10),
            After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
            Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListVaultProviders != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListVaultProvidersRequest](../../models/operations/listvaultprovidersrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListVaultProvidersResponse](../../models/operations/listvaultprovidersresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetVaultProvider

Retrieves a single vault provider by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getVaultProvider" method="get" path="/v2/directories/{directoryId}/vault/providers/{providerId}" -->
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

    res, err := s.TokenVaultProviders.GetVaultProvider(ctx, "d8e5e185-d6b7-48e8-8c90-d9e691c5e9ad", "9923e17e-47ee-45cb-bf51-2258171d95d1")
    if err != nil {
        log.Fatal(err)
    }
    if res.VaultProvider != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `directoryID`                                            | `string`                                                 | :heavy_check_mark:                                       | ID of the directory                                      |
| `providerID`                                             | `string`                                                 | :heavy_check_mark:                                       | ID of the vault provider                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetVaultProviderResponse](../../models/operations/getvaultproviderresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateVaultProvider

Partially updates an existing vault provider.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateVaultProvider" method="patch" path="/v2/directories/{directoryId}/vault/providers/{providerId}" -->
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

    res, err := s.TokenVaultProviders.UpdateVaultProvider(ctx, operations.UpdateVaultProviderRequest{
        DirectoryID: "7bf1060b-5938-49e1-a82b-cb8b6dff06c5",
        ProviderID: "2fce0e81-0a0e-4682-8a8f-87836f29b126",
        PatchVaultProviderBody: components.PatchVaultProviderBody{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VaultProvider != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateVaultProviderRequest](../../models/operations/updatevaultproviderrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateVaultProviderResponse](../../models/operations/updatevaultproviderresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ReplaceVaultProvider

Completely replaces an existing vault provider's configuration.
All fields must be specified; any omitted fields will be reset to their defaults.


### Example Usage

<!-- UsageSnippet language="go" operationID="replaceVaultProvider" method="put" path="/v2/directories/{directoryId}/vault/providers/{providerId}" -->
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

    res, err := s.TokenVaultProviders.ReplaceVaultProvider(ctx, operations.ReplaceVaultProviderRequest{
        DirectoryID: "30e89ae4-931b-466e-982c-e11e7eae2663",
        ProviderID: "1cb91a7d-e5ff-41ed-a7ad-67eb96866643",
        ReplaceVaultProviderBody: components.ReplaceVaultProviderBody{
            TemplateName: components.ProviderTemplateNameGithub,
            Name: "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VaultProvider != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ReplaceVaultProviderRequest](../../models/operations/replacevaultproviderrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ReplaceVaultProviderResponse](../../models/operations/replacevaultproviderresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteVaultProvider

Deletes a vault provider from a directory.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteVaultProvider" method="delete" path="/v2/directories/{directoryId}/vault/providers/{providerId}" -->
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

    res, err := s.TokenVaultProviders.DeleteVaultProvider(ctx, "97a2fd3f-365d-4fd8-bfdf-7d395fdd78aa", "f7d6b04b-4960-44d3-a850-568937934a4a")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `directoryID`                                            | `string`                                                 | :heavy_check_mark:                                       | ID of the directory                                      |
| `providerID`                                             | `string`                                                 | :heavy_check_mark:                                       | ID of the vault provider                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteVaultProviderResponse](../../models/operations/deletevaultproviderresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateVaultProviderCredential

Creates a shared credential for a provider. For a static_secret provider, stores
a literal secret value (at most one, a second create returns 409). For a shared
authorization_code provider, mints an OAuth enrollment URL for an
admin to complete once instead of accepting a secret value.


### Example Usage

<!-- UsageSnippet language="go" operationID="createVaultProviderCredential" method="post" path="/v2/directories/{directoryId}/vault/providers/{providerId}/credentials" -->
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

    res, err := s.TokenVaultProviders.CreateVaultProviderCredential(ctx, operations.CreateVaultProviderCredentialRequest{
        DirectoryID: "5effe433-f32b-4fbe-b038-9f2eb32c141f",
        ProviderID: "20819d16-3ab9-4f54-885f-50fac89bb6d9",
        SetVaultProviderCredentialBody: components.SetVaultProviderCredentialBody{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VaultProviderCredentialShared != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.CreateVaultProviderCredentialRequest](../../models/operations/createvaultprovidercredentialrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.CreateVaultProviderCredentialResponse](../../models/operations/createvaultprovidercredentialresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListVaultProviderCredentials

Returns a paginated list of credential metadata for a provider. Each
item's shape depends on its credential_type. shared items are the
provider's shared credential: the admin-configured secret. Never include
enrollment_url (fetch by id to mint one; shared authorization_code
credentials only). user items
are each principal's individual authorization, embedding principal id
and display_name. Credential values are never returned.


### Example Usage

<!-- UsageSnippet language="go" operationID="listVaultProviderCredentials" method="get" path="/v2/directories/{directoryId}/vault/providers/{providerId}/credentials" -->
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

    res, err := s.TokenVaultProviders.ListVaultProviderCredentials(ctx, operations.ListVaultProviderCredentialsRequest{
        DirectoryID: "625e54ff-f4fd-4520-ba1c-cee0afa20256",
        ProviderID: "cbe2302f-0d33-4663-a4b2-bf1cc9af0e2f",
        Page: &components.CursorPageParameters{
            Size: sdkkonnectgo.Pointer[int64](10),
            After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
            Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListVaultProviderCredentials != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListVaultProviderCredentialsRequest](../../models/operations/listvaultprovidercredentialsrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ListVaultProviderCredentialsResponse](../../models/operations/listvaultprovidercredentialsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetVaultProviderCredential

Returns a single credential by id. For a shared authorization_code
credential, always mints a fresh enrollment_url on every call,
regardless of whether the credential is pending or active, so an
admin can re-enroll at any time.


### Example Usage

<!-- UsageSnippet language="go" operationID="getVaultProviderCredential" method="get" path="/v2/directories/{directoryId}/vault/providers/{providerId}/credentials/{credentialId}" -->
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

    res, err := s.TokenVaultProviders.GetVaultProviderCredential(ctx, operations.GetVaultProviderCredentialRequest{
        DirectoryID: "a4f3c8a6-7c13-41f8-b3a8-b31845b83b26",
        ProviderID: "008bbc75-4078-46c6-95c5-611a448aa0e9",
        CredentialID: "18874437-a870-4d69-92ae-c32195aa98f9",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VaultProviderCredential != nil {
        switch res.VaultProviderCredential.Type {
            case components.VaultProviderCredentialTypeShared:
                // res.VaultProviderCredential.VaultProviderCredentialShared is populated
            case components.VaultProviderCredentialTypeUser:
                // res.VaultProviderCredential.VaultProviderCredentialUser is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetVaultProviderCredentialRequest](../../models/operations/getvaultprovidercredentialrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.GetVaultProviderCredentialResponse](../../models/operations/getvaultprovidercredentialresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateVaultProviderCredential

Replaces the value of an existing credential and returns 200 with the
updated credential body, for both static_secret providers and shared
authorization_code providers (the latter re-mints a fresh enrollment URL).
Returns 404 when no secret/credential exists yet for this provider.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateVaultProviderCredential" method="put" path="/v2/directories/{directoryId}/vault/providers/{providerId}/credentials/{credentialId}" -->
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

    res, err := s.TokenVaultProviders.UpdateVaultProviderCredential(ctx, operations.UpdateVaultProviderCredentialRequest{
        DirectoryID: "45596e86-5066-4480-a07b-8101b3dd1225",
        ProviderID: "c2173233-e90f-4a57-beec-c27b9a3fd503",
        CredentialID: "8d888862-5c26-44c2-8bfe-7592fd219f75",
        SetVaultProviderCredentialBody: components.SetVaultProviderCredentialBody{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VaultProviderCredentialShared != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.UpdateVaultProviderCredentialRequest](../../models/operations/updatevaultprovidercredentialrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.UpdateVaultProviderCredentialResponse](../../models/operations/updatevaultprovidercredentialresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteVaultProviderCredential

Delete a specific provider credential

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteVaultProviderCredential" method="delete" path="/v2/directories/{directoryId}/vault/providers/{providerId}/credentials/{credentialId}" -->
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

    res, err := s.TokenVaultProviders.DeleteVaultProviderCredential(ctx, operations.DeleteVaultProviderCredentialRequest{
        DirectoryID: "83bb20ac-765e-40a3-9658-bc46d75dd3cd",
        ProviderID: "577bf00b-c85e-45df-82a8-cf8b6a115622",
        CredentialID: "8872d9da-565d-4aaa-9384-e6189ffe1f35",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.DeleteVaultProviderCredentialRequest](../../models/operations/deletevaultprovidercredentialrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.DeleteVaultProviderCredentialResponse](../../models/operations/deletevaultprovidercredentialresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |