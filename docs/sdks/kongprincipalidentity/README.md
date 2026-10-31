# KongPrincipalIdentity

## Overview

Identities represent external authentication associations for a principal. Each identity links a principal
to an external system such as an OIDC provider, an auth server client, a control plane consumer, or a custom identifier.


### Available Operations

* [ListIdentities](#listidentities) - List all identities for a principal
* [CreateIdentity](#createidentity) - Create an identity for a principal
* [GetIdentity](#getidentity) - Get an identity
* [ReplaceIdentity](#replaceidentity) - Replace an identity
* [UpdateIdentity](#updateidentity) - Update an identity
* [DeleteIdentity](#deleteidentity) - Delete an identity

## ListIdentities

Retrieves a paginated list of all identities associated with the specified principal.


### Example Usage

<!-- UsageSnippet language="go" operationID="listIdentities" method="get" path="/v2/directories/{directoryId}/principals/{principalId}/identities" -->
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

    res, err := s.KongPrincipalIdentity.ListIdentities(ctx, operations.ListIdentitiesRequest{
        DirectoryID: "bcefadc3-9616-47fe-80fb-e7318f5022d4",
        PrincipalID: "2fb1411b-45ab-48d2-9fae-cd58c76da0fe",
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
    if res.ListKongPrincipalIdentities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListIdentitiesRequest](../../models/operations/listidentitiesrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListIdentitiesResponse](../../models/operations/listidentitiesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateIdentity

Creates a new identity association for the specified principal.


### Example Usage

<!-- UsageSnippet language="go" operationID="createIdentity" method="post" path="/v2/directories/{directoryId}/principals/{principalId}/identities" -->
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

    res, err := s.KongPrincipalIdentity.CreateIdentity(ctx, operations.CreateIdentityRequest{
        DirectoryID: "ff508ee3-f3c6-44d7-b9ef-43513decec8b",
        PrincipalID: "f4db939a-b2be-46c7-99ed-261e9e77b76f",
        KongPrincipalIdentityBody: components.CreateKongPrincipalIdentityBodyCustom(
            components.KongPrincipalIdentityCustom{
                Type: components.KongPrincipalIdentityCustomTypeCustom,
                Key: "<key>",
                Value: "<value>",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongPrincipalIdentity != nil {
        switch res.KongPrincipalIdentity.Type {
            case components.KongPrincipalIdentityTypeOidc:
                // res.KongPrincipalIdentity.KongPrincipalIdentityOIDCResponse is populated
            case components.KongPrincipalIdentityTypeAuthServerClient:
                // res.KongPrincipalIdentity.KongPrincipalIdentityAuthServerClientResponse is populated
            case components.KongPrincipalIdentityTypeControlPlaneConsumer:
                // res.KongPrincipalIdentity.KongPrincipalIdentityCPConsumerResponse is populated
            case components.KongPrincipalIdentityTypeCustom:
                // res.KongPrincipalIdentity.KongPrincipalIdentityCustomResponse is populated
        }

    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateIdentityRequest](../../models/operations/createidentityrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateIdentityResponse](../../models/operations/createidentityresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetIdentity

Retrieves detailed information about a specific principal identity by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getIdentity" method="get" path="/v2/directories/{directoryId}/principals/{principalId}/identities/{identityId}" -->
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

    res, err := s.KongPrincipalIdentity.GetIdentity(ctx, operations.GetIdentityRequest{
        DirectoryID: "89febfa0-d5e8-4428-a63b-e8b6af2053ef",
        PrincipalID: "a1f801a0-a6d6-43ad-8778-8e1284ce5c07",
        IdentityID: "3fd14482-c6a7-48d6-9936-f508af0552f2",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongPrincipalIdentity != nil {
        switch res.KongPrincipalIdentity.Type {
            case components.KongPrincipalIdentityTypeOidc:
                // res.KongPrincipalIdentity.KongPrincipalIdentityOIDCResponse is populated
            case components.KongPrincipalIdentityTypeAuthServerClient:
                // res.KongPrincipalIdentity.KongPrincipalIdentityAuthServerClientResponse is populated
            case components.KongPrincipalIdentityTypeControlPlaneConsumer:
                // res.KongPrincipalIdentity.KongPrincipalIdentityCPConsumerResponse is populated
            case components.KongPrincipalIdentityTypeCustom:
                // res.KongPrincipalIdentity.KongPrincipalIdentityCustomResponse is populated
        }

    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetIdentityRequest](../../models/operations/getidentityrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.GetIdentityResponse](../../models/operations/getidentityresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ReplaceIdentity

Completely replaces an existing principal identity with the provided data.
All fields must be specified; any omitted fields will be reset to their defaults.


### Example Usage

<!-- UsageSnippet language="go" operationID="replaceIdentity" method="put" path="/v2/directories/{directoryId}/principals/{principalId}/identities/{identityId}" -->
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

    res, err := s.KongPrincipalIdentity.ReplaceIdentity(ctx, operations.ReplaceIdentityRequest{
        DirectoryID: "afded8a2-e169-4daf-8c86-6c4c6563fe95",
        PrincipalID: "da9a2255-e89f-4390-be81-7a84b5639158",
        IdentityID: "ea06476f-4b1b-41e5-89e4-fb05512fa5f8",
        KongPrincipalIdentityBody: components.CreateKongPrincipalIdentityBodyAuthServerClient(
            components.KongPrincipalIdentityAuthServerClient{
                Type: components.KongPrincipalIdentityAuthServerClientTypeAuthServerClient,
                AuthServerID: "52749349-21f9-414a-a46e-098326545de5",
                ClientID: "<id>",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongPrincipalIdentity != nil {
        switch res.KongPrincipalIdentity.Type {
            case components.KongPrincipalIdentityTypeOidc:
                // res.KongPrincipalIdentity.KongPrincipalIdentityOIDCResponse is populated
            case components.KongPrincipalIdentityTypeAuthServerClient:
                // res.KongPrincipalIdentity.KongPrincipalIdentityAuthServerClientResponse is populated
            case components.KongPrincipalIdentityTypeControlPlaneConsumer:
                // res.KongPrincipalIdentity.KongPrincipalIdentityCPConsumerResponse is populated
            case components.KongPrincipalIdentityTypeCustom:
                // res.KongPrincipalIdentity.KongPrincipalIdentityCustomResponse is populated
        }

    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ReplaceIdentityRequest](../../models/operations/replaceidentityrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ReplaceIdentityResponse](../../models/operations/replaceidentityresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateIdentity

Partially updates an existing principal identity by modifying only the provided fields.
Unspecified fields will retain their current values.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateIdentity" method="patch" path="/v2/directories/{directoryId}/principals/{principalId}/identities/{identityId}" -->
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

    res, err := s.KongPrincipalIdentity.UpdateIdentity(ctx, operations.UpdateIdentityRequest{
        DirectoryID: "5c59b3d4-1c9d-4f1b-9073-0e315d8e40a1",
        PrincipalID: "bdf1e9b0-2173-4b66-8239-edd63f4f8d11",
        IdentityID: "e47aa317-a004-4b5c-975e-12703a8424ba",
        UpdatePrincipalIdentityBody: components.UpdatePrincipalIdentityBody{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongPrincipalIdentity != nil {
        switch res.KongPrincipalIdentity.Type {
            case components.KongPrincipalIdentityTypeOidc:
                // res.KongPrincipalIdentity.KongPrincipalIdentityOIDCResponse is populated
            case components.KongPrincipalIdentityTypeAuthServerClient:
                // res.KongPrincipalIdentity.KongPrincipalIdentityAuthServerClientResponse is populated
            case components.KongPrincipalIdentityTypeControlPlaneConsumer:
                // res.KongPrincipalIdentity.KongPrincipalIdentityCPConsumerResponse is populated
            case components.KongPrincipalIdentityTypeCustom:
                // res.KongPrincipalIdentity.KongPrincipalIdentityCustomResponse is populated
        }

    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateIdentityRequest](../../models/operations/updateidentityrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateIdentityResponse](../../models/operations/updateidentityresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteIdentity

Permanently deletes an identity association from the principal.
This operation cannot be undone.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteIdentity" method="delete" path="/v2/directories/{directoryId}/principals/{principalId}/identities/{identityId}" -->
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

    res, err := s.KongPrincipalIdentity.DeleteIdentity(ctx, operations.DeleteIdentityRequest{
        DirectoryID: "404c1da4-2073-49d2-a013-85521fce7e41",
        PrincipalID: "04a131cd-0c74-474d-979e-763c83a6d7df",
        IdentityID: "7b29ee4b-d354-46f7-8b96-c875b067193e",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteIdentityRequest](../../models/operations/deleteidentityrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.DeleteIdentityResponse](../../models/operations/deleteidentityresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |