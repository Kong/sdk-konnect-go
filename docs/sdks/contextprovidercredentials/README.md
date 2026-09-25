# ContextProviderCredentials

## Overview

### Available Operations

* [ListContextProviderCredentials](#listcontextprovidercredentials) - List Context Provider Credentials
* [CreateContextProviderCredential](#createcontextprovidercredential) - Create a Context Provider Credential
* [GetContextProviderCredential](#getcontextprovidercredential) - Get a Context Provider Credential
* [PatchContextProviderCredential](#patchcontextprovidercredential) - Update a Context Provider Credential
* [DeleteContextProviderCredential](#deletecontextprovidercredential) - Delete a Context Provider Credential

## ListContextProviderCredentials

Returns a list of context provider credentials for the organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-context-provider-credentials" method="get" path="/v1/context-provider-credentials" -->
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

    res, err := s.ContextProviderCredentials.ListContextProviderCredentials(ctx, operations.ListContextProviderCredentialsRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListContextProviderCredentialsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.ListContextProviderCredentialsRequest](../../models/operations/listcontextprovidercredentialsrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.ListContextProviderCredentialsResponse](../../models/operations/listcontextprovidercredentialsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateContextProviderCredential

Create a context provider credential for the organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-context-provider-credential" method="post" path="/v1/context-provider-credentials" -->
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

    res, err := s.ContextProviderCredentials.CreateContextProviderCredential(ctx, components.CreateContextProviderCredentialRequest{
        Name: "bob-github-pat",
        DisplayName: "Bob's Github Pat",
        Description: "Expires on 2027-01-30",
        Provider: components.CreateContextProviderCredentialRequestProviderGithub,
        Config: components.CreateContextProviderCredentialConfigPayloadBearer(
            components.BearerContextProviderCredentialConfigPayload{
                AuthType: components.BearerContextProviderCredentialConfigPayloadAuthTypeBearer,
                Secrets: components.BearerCredentialSecrets{
                    Token: "<value>",
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ContextProviderCredential != nil {
        switch res.ContextProviderCredential.Config.Type {
            case components.ContextProviderCredentialConfigTypeBearer:
                // res.ContextProviderCredential.Config.BearerContextProviderCredentialConfig is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [components.CreateContextProviderCredentialRequest](../../models/components/createcontextprovidercredentialrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.CreateContextProviderCredentialResponse](../../models/operations/createcontextprovidercredentialresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetContextProviderCredential

Retrieve a single context provider credential by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-context-provider-credential" method="get" path="/v1/context-provider-credentials/{contextProviderCredentialId}" -->
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

    res, err := s.ContextProviderCredentials.GetContextProviderCredential(ctx, "ff789428-969e-428c-a5ca-7bf92a4f21a2")
    if err != nil {
        log.Fatal(err)
    }
    if res.ContextProviderCredential != nil {
        switch res.ContextProviderCredential.Config.Type {
            case components.ContextProviderCredentialConfigTypeBearer:
                // res.ContextProviderCredential.Config.BearerContextProviderCredentialConfig is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `contextProviderCredentialID`                            | `string`                                                 | :heavy_check_mark:                                       | The ID of the context provider credential.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetContextProviderCredentialResponse](../../models/operations/getcontextprovidercredentialresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchContextProviderCredential

Partially update a context provider credential. `provider` cannot be changed after creation. `config`, if provided, must be specified in its entirety — it replaces the existing configuration rather than merging into it.


### Example Usage

<!-- UsageSnippet language="go" operationID="patch-context-provider-credential" method="patch" path="/v1/context-provider-credentials/{contextProviderCredentialId}" -->
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

    res, err := s.ContextProviderCredentials.PatchContextProviderCredential(ctx, "fde85815-b6af-4185-98bb-9e3c506bfd18", components.PatchContextProviderCredentialRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ContextProviderCredential != nil {
        switch res.ContextProviderCredential.Config.Type {
            case components.ContextProviderCredentialConfigTypeBearer:
                // res.ContextProviderCredential.Config.BearerContextProviderCredentialConfig is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `contextProviderCredentialID`                                                                                        | `string`                                                                                                             | :heavy_check_mark:                                                                                                   | The ID of the context provider credential.                                                                           |
| `patchContextProviderCredentialRequest`                                                                              | [components.PatchContextProviderCredentialRequest](../../models/components/patchcontextprovidercredentialrequest.md) | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.PatchContextProviderCredentialResponse](../../models/operations/patchcontextprovidercredentialresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteContextProviderCredential

Delete a context provider credential. Skills that reference this credential are not deleted and keep serving their last promoted version, but subsequent syncs for them will fail.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-context-provider-credential" method="delete" path="/v1/context-provider-credentials/{contextProviderCredentialId}" -->
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

    res, err := s.ContextProviderCredentials.DeleteContextProviderCredential(ctx, "4d64cd1e-c917-4a3c-bb95-68ddaa57ca20")
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
| `contextProviderCredentialID`                            | `string`                                                 | :heavy_check_mark:                                       | The ID of the context provider credential.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteContextProviderCredentialResponse](../../models/operations/deletecontextprovidercredentialresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |