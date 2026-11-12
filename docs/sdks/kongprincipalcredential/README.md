# KongPrincipalCredential

## Overview

Credentials are authentication secrets associated with principals. Supported credential types are
API keys (system-generated or imported) and basic auth username/password pairs.


### Available Operations

* [CreatePrincipalKey](#createprincipalkey) - Create a principal API key
* [ListPrincipalKeys](#listprincipalkeys) - List all API keys for a principal
* [UpdatePrincipalKey](#updateprincipalkey) - Update a principal API key
* [GetIdentityKey](#getidentitykey) - Get a principal API key
* [DeletePrincipalKey](#deleteprincipalkey) - Delete an API key
* [ListPrincipalBasicAuths](#listprincipalbasicauths) - List all basic auth credentials
* [CreatePrincipalBasicAuth](#createprincipalbasicauth) - Create a basic auth credential
* [GetPrincipalBasicAuth](#getprincipalbasicauth) - Get a basic auth credential
* [UpdatePrincipalBasicAuth](#updateprincipalbasicauth) - Update a basic auth credential
* [ReplacePrincipalBasicAuth](#replaceprincipalbasicauth) - Replace a basic auth credential
* [DeletePrincipalBasicAuth](#deleteprincipalbasicauth) - Delete a basic auth credential
* [ListPasswords](#listpasswords) - List all passwords for a basic auth
* [CreatePassword](#createpassword) - Create a password for a basic auth credential
* [DeleteAllPasswords](#deleteallpasswords) - Delete all passwords for a basic auth
* [GetPassword](#getpassword) - Get a password
* [UpdatePassword](#updatepassword) - Update a password
* [ReplacePassword](#replacepassword) - Replace a password
* [DeletePassword](#deletepassword) - Delete a password

## CreatePrincipalKey

Creates a new API key for the specified principal.
Keys can be system-generated (v1) or imported with a custom secret, and may have an optional expiration date.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-principal-key" method="post" path="/v2/directories/{directoryId}/principals/{principalId}/api-keys" -->
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

    res, err := s.KongPrincipalCredential.CreatePrincipalKey(ctx, operations.CreatePrincipalKeyRequest{
        DirectoryID: "57aa00d4-42e1-46f2-926d-eb1796773a80",
        PrincipalID: "28de680e-3db1-46f0-8de7-5f8b89b8f80b",
        CreatePrincipalKeyBody: components.CreateCreatePrincipalKeyBodyV1(
            components.CreatePrincipalKeyBodyv1{
                Type: components.CreatePrincipalKeyBodyv1TypeV1,
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongPrincipalKeyWithSecret != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreatePrincipalKeyRequest](../../models/operations/createprincipalkeyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreatePrincipalKeyResponse](../../models/operations/createprincipalkeyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListPrincipalKeys

Retrieves a paginated list of all API keys associated with the specified principal.
Results can be sorted by created_at, updated_at, or name.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-principal-keys" method="get" path="/v2/directories/{directoryId}/principals/{principalId}/api-keys" -->
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

    res, err := s.KongPrincipalCredential.ListPrincipalKeys(ctx, operations.ListPrincipalKeysRequest{
        DirectoryID: "bdc010b0-3c43-4209-ab61-d9ce1f2288a8",
        PrincipalID: "e2c75760-eef8-472a-b3e2-09ce52c960e8",
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
    if res.ListKongPrincipalKeys != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListPrincipalKeysRequest](../../models/operations/listprincipalkeysrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListPrincipalKeysResponse](../../models/operations/listprincipalkeysresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePrincipalKey

Partially updates an existing API key's metadata.
Note: The key secret cannot be modified after creation.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-principal-key" method="patch" path="/v2/directories/{directoryId}/principals/{principalId}/api-keys/{apiKeyId}" -->
```go
package main

import (
	"github.com/Kong/sdk-konnect-go/optionalnullable"
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/types"
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

    res, err := s.KongPrincipalCredential.UpdatePrincipalKey(ctx, operations.UpdatePrincipalKeyRequest{
        DirectoryID: "bf6936e8-ee51-4b26-9dc9-8d20cff07624",
        PrincipalID: "7136f099-1c3b-48b2-aad2-a4532ad54772",
        APIKeyID: "4ae93307-c032-4675-bc67-60797bf74224",
        UpdatePrincipalKeyBody: components.UpdatePrincipalKeyBody{
            ExpiresAt: optionalnullable.From(types.MustNewTimeFromString("2022-11-04T20:10:06.927Z")),
            ManagedBy: map[string]*string{
                "owner": sdkkonnectgo.Pointer("terraform"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongPrincipalKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdatePrincipalKeyRequest](../../models/operations/updateprincipalkeyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdatePrincipalKeyResponse](../../models/operations/updateprincipalkeyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetIdentityKey

Retrieves detailed information about a specific API key.
Note: The key secret is not returned; it is only available when the key is first created.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-identity-key" method="get" path="/v2/directories/{directoryId}/principals/{principalId}/api-keys/{apiKeyId}" -->
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

    res, err := s.KongPrincipalCredential.GetIdentityKey(ctx, operations.GetIdentityKeyRequest{
        DirectoryID: "7c7d4c1f-cb90-4d73-bfcb-8e406c603393",
        PrincipalID: "6a668bb0-6965-499f-8d5a-404632e44cf6",
        APIKeyID: "1608ff96-7d46-4d00-8e0b-e5a25940678e",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongPrincipalKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetIdentityKeyRequest](../../models/operations/getidentitykeyrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetIdentityKeyResponse](../../models/operations/getidentitykeyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeletePrincipalKey

Permanently deletes an API key, immediately revoking its access.
This operation cannot be undone.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-principal-key" method="delete" path="/v2/directories/{directoryId}/principals/{principalId}/api-keys/{apiKeyId}" -->
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

    res, err := s.KongPrincipalCredential.DeletePrincipalKey(ctx, operations.DeletePrincipalKeyRequest{
        DirectoryID: "57944b92-3e96-4fd5-9b2e-4ed2c7086dab",
        PrincipalID: "6f4cc144-7623-42a7-8571-31b4226cefd0",
        APIKeyID: "89e8482a-9f51-4f1f-a6c1-b80fe7cf18bf",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeletePrincipalKeyRequest](../../models/operations/deleteprincipalkeyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.DeletePrincipalKeyResponse](../../models/operations/deleteprincipalkeyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListPrincipalBasicAuths

Retrieves a paginated list of all basic auth credentials associated with the specified principal.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-principal-basic-auths" method="get" path="/v2/directories/{directoryId}/principals/{principalId}/basic-auths" -->
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

    res, err := s.KongPrincipalCredential.ListPrincipalBasicAuths(ctx, operations.ListPrincipalBasicAuthsRequest{
        DirectoryID: "c1717152-f121-49cc-a14b-ae0a3be6a468",
        PrincipalID: "9092d33a-16a8-4263-8d26-3b0a5fb58c48",
        Page: &components.CursorPageParameters{
            Size: sdkkonnectgo.Pointer[int64](10),
            After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
            Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListKongBasicAuths != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListPrincipalBasicAuthsRequest](../../models/operations/listprincipalbasicauthsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListPrincipalBasicAuthsResponse](../../models/operations/listprincipalbasicauthsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreatePrincipalBasicAuth

Creates a new basic auth credential for the specified principal.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-principal-basic-auth" method="post" path="/v2/directories/{directoryId}/principals/{principalId}/basic-auths" -->
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

    res, err := s.KongPrincipalCredential.CreatePrincipalBasicAuth(ctx, operations.CreatePrincipalBasicAuthRequest{
        DirectoryID: "86088ac4-7312-47bc-881c-6a4dcfef0a9b",
        PrincipalID: "d70474a7-068d-4a56-879f-faef017c9340",
        CreateBasicAuthBody: components.CreateBasicAuthBody{
            Username: "Lenore.Toy93",
            ManagedBy: map[string]string{
                "owner": "terraform",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongBasicAuth != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreatePrincipalBasicAuthRequest](../../models/operations/createprincipalbasicauthrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.CreatePrincipalBasicAuthResponse](../../models/operations/createprincipalbasicauthresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPrincipalBasicAuth

Retrieves a specific basic auth credential by ID.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-principal-basic-auth" method="get" path="/v2/directories/{directoryId}/principals/{principalId}/basic-auths/{basicAuthId}" -->
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

    res, err := s.KongPrincipalCredential.GetPrincipalBasicAuth(ctx, operations.GetPrincipalBasicAuthRequest{
        DirectoryID: "938e509b-bb28-4bdb-b1c7-0ed2e0577aa2",
        PrincipalID: "790c4dc3-4b42-44f9-883d-1da8b0bf033c",
        BasicAuthID: "6173ed16-e311-494a-a6c0-93ff25385c18",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongBasicAuth != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetPrincipalBasicAuthRequest](../../models/operations/getprincipalbasicauthrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetPrincipalBasicAuthResponse](../../models/operations/getprincipalbasicauthresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePrincipalBasicAuth

Partially updates an existing basic auth credential. Unspecified fields retain their current values.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-principal-basic-auth" method="patch" path="/v2/directories/{directoryId}/principals/{principalId}/basic-auths/{basicAuthId}" -->
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

    res, err := s.KongPrincipalCredential.UpdatePrincipalBasicAuth(ctx, operations.UpdatePrincipalBasicAuthRequest{
        DirectoryID: "4dc56174-4bbc-4abf-9c3b-17cbf38acc6e",
        PrincipalID: "bff40436-7077-4cc0-ac5e-6908762a61c8",
        BasicAuthID: "f76fa345-b72f-4409-a107-7c39a775f46b",
        UpdateBasicAuthBody: components.UpdateBasicAuthBody{
            ManagedBy: map[string]*string{
                "owner": sdkkonnectgo.Pointer("terraform"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongBasicAuth != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdatePrincipalBasicAuthRequest](../../models/operations/updateprincipalbasicauthrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpdatePrincipalBasicAuthResponse](../../models/operations/updateprincipalbasicauthresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ReplacePrincipalBasicAuth

Fully replaces an existing basic auth credential with the provided values.


### Example Usage

<!-- UsageSnippet language="go" operationID="replace-principal-basic-auth" method="put" path="/v2/directories/{directoryId}/principals/{principalId}/basic-auths/{basicAuthId}" -->
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

    res, err := s.KongPrincipalCredential.ReplacePrincipalBasicAuth(ctx, operations.ReplacePrincipalBasicAuthRequest{
        DirectoryID: "27d584bf-48bd-4b75-acab-cde93d5bda8c",
        PrincipalID: "27207874-84fc-417c-a7f3-24f0f2b508e7",
        BasicAuthID: "d8dde232-e6cd-4ee4-802a-3c487d8382c8",
        ReplaceBasicAuthBody: components.ReplaceBasicAuthBody{
            Username: "Craig92",
            ManagedBy: map[string]string{
                "owner": "terraform",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongBasicAuth != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ReplacePrincipalBasicAuthRequest](../../models/operations/replaceprincipalbasicauthrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ReplacePrincipalBasicAuthResponse](../../models/operations/replaceprincipalbasicauthresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeletePrincipalBasicAuth

Permanently deletes a basic auth credential and all associated passwords.
This operation cannot be undone.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-principal-basic-auth" method="delete" path="/v2/directories/{directoryId}/principals/{principalId}/basic-auths/{basicAuthId}" -->
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

    res, err := s.KongPrincipalCredential.DeletePrincipalBasicAuth(ctx, operations.DeletePrincipalBasicAuthRequest{
        DirectoryID: "f74cd7e5-f4a0-4aff-a205-faf5eda271ba",
        PrincipalID: "d0a80504-ef73-4aba-98df-cfdaffa7ae02",
        BasicAuthID: "cee29617-4cad-43c8-b8a2-2ae1bacd48d0",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.DeletePrincipalBasicAuthRequest](../../models/operations/deleteprincipalbasicauthrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.DeletePrincipalBasicAuthResponse](../../models/operations/deleteprincipalbasicauthresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListPasswords

Retrieves a paginated list of all passwords associated with the specified basic auth credential.


### Example Usage

<!-- UsageSnippet language="go" operationID="listPasswords" method="get" path="/v2/directories/{directoryId}/principals/{principalId}/basic-auths/{basicAuthId}/passwords" -->
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

    res, err := s.KongPrincipalCredential.ListPasswords(ctx, operations.ListPasswordsRequest{
        DirectoryID: "c6bdffc1-ad7e-49d4-a398-38b39a4f0c4e",
        PrincipalID: "10226343-6342-434d-9e4e-4e227489454c",
        BasicAuthID: "3ca07846-903b-4ac1-bb41-91b418c4e43e",
        Page: &components.CursorPageParameters{
            Size: sdkkonnectgo.Pointer[int64](10),
            After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
            Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListKongPasswords != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListPasswordsRequest](../../models/operations/listpasswordsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ListPasswordsResponse](../../models/operations/listpasswordsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreatePassword

Creates a new password for the specified basic auth credential.
The password value is only returned on creation and cannot be retrieved again.


### Example Usage

<!-- UsageSnippet language="go" operationID="createPassword" method="post" path="/v2/directories/{directoryId}/principals/{principalId}/basic-auths/{basicAuthId}/passwords" -->
```go
package main

import (
	"github.com/Kong/sdk-konnect-go/optionalnullable"
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/types"
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

    res, err := s.KongPrincipalCredential.CreatePassword(ctx, operations.CreatePasswordRequest{
        DirectoryID: "a58eb2cb-a2e5-496b-a4c0-974761374b0e",
        PrincipalID: "bf882789-3b05-49ea-a4ef-8fe35a1edb51",
        BasicAuthID: "b923c4ad-bc12-4dc4-b128-97fc6ed52548",
        CreatePasswordBody: components.CreatePasswordBody{
            ExpiresAt: optionalnullable.From(types.MustNewTimeFromString("2022-11-04T20:10:06.927Z")),
            ManagedBy: map[string]string{
                "owner": "terraform",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongPasswordWithSecret != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreatePasswordRequest](../../models/operations/createpasswordrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreatePasswordResponse](../../models/operations/createpasswordresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteAllPasswords

Permanently deletes all passwords associated with the specified basic auth credential.
This operation cannot be undone.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteAllPasswords" method="delete" path="/v2/directories/{directoryId}/principals/{principalId}/basic-auths/{basicAuthId}/passwords" -->
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

    res, err := s.KongPrincipalCredential.DeleteAllPasswords(ctx, operations.DeleteAllPasswordsRequest{
        DirectoryID: "7b1b689e-2db7-4744-bc66-9a571a754e09",
        PrincipalID: "d31fafb3-762b-44ae-9ab0-80dd5833758c",
        BasicAuthID: "a2edd480-7a5e-480e-96c9-bfc4130404b7",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteAllPasswordsRequest](../../models/operations/deleteallpasswordsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.DeleteAllPasswordsResponse](../../models/operations/deleteallpasswordsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPassword

Retrieves a specific password by ID. The password value is not returned.


### Example Usage

<!-- UsageSnippet language="go" operationID="getPassword" method="get" path="/v2/directories/{directoryId}/principals/{principalId}/basic-auths/{basicAuthId}/passwords/{passwordId}" -->
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

    res, err := s.KongPrincipalCredential.GetPassword(ctx, operations.GetPasswordRequest{
        DirectoryID: "6cf025ce-67af-4d40-afae-2415d97dc92a",
        PrincipalID: "e643ce22-53fd-4743-8586-c15aa227e2a4",
        BasicAuthID: "f3190ff0-2a14-4f54-b06d-4d7a819fe0ae",
        PasswordID: "2d60cb94-fe9f-4fb3-a18f-45a72ce6ca3c",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongPassword != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetPasswordRequest](../../models/operations/getpasswordrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.GetPasswordResponse](../../models/operations/getpasswordresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePassword

Partially updates a password's expiration date.


### Example Usage

<!-- UsageSnippet language="go" operationID="updatePassword" method="patch" path="/v2/directories/{directoryId}/principals/{principalId}/basic-auths/{basicAuthId}/passwords/{passwordId}" -->
```go
package main

import (
	"github.com/Kong/sdk-konnect-go/optionalnullable"
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/types"
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

    res, err := s.KongPrincipalCredential.UpdatePassword(ctx, operations.UpdatePasswordRequest{
        DirectoryID: "d2eab6f5-a35d-4653-8475-84b7e4ea4fec",
        PrincipalID: "8a445420-41b3-4f69-acf9-113cfe51f3e7",
        BasicAuthID: "4b4f2381-f061-40f6-aa6b-e77793362610",
        PasswordID: "1974cebf-191e-47b0-b226-a516d285c2a1",
        UpdatePasswordBody: components.UpdatePasswordBody{
            ExpiresAt: optionalnullable.From(types.MustNewTimeFromString("2022-11-04T20:10:06.927Z")),
            ManagedBy: map[string]*string{
                "owner": sdkkonnectgo.Pointer("terraform"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongPassword != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdatePasswordRequest](../../models/operations/updatepasswordrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdatePasswordResponse](../../models/operations/updatepasswordresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ReplacePassword

Fully replaces an existing password with the provided values.


### Example Usage

<!-- UsageSnippet language="go" operationID="replacePassword" method="put" path="/v2/directories/{directoryId}/principals/{principalId}/basic-auths/{basicAuthId}/passwords/{passwordId}" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/types"
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

    res, err := s.KongPrincipalCredential.ReplacePassword(ctx, operations.ReplacePasswordRequest{
        DirectoryID: "91b54b30-7d72-4626-bbbb-bc38d0edd659",
        PrincipalID: "78c28e45-dbb7-4595-a9d5-99fb71cda35c",
        BasicAuthID: "fe3fa8ea-f8d4-4816-9355-70a59434c552",
        PasswordID: "df80a2b7-ceda-467f-a25e-a0a2eff68690",
        ReplacePasswordBody: components.ReplacePasswordBody{
            Secret: "<value>",
            ExpiresAt: types.MustNewTimeFromString("2022-11-04T20:10:06.927Z"),
            ManagedBy: map[string]string{
                "owner": "terraform",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongPassword != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ReplacePasswordRequest](../../models/operations/replacepasswordrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ReplacePasswordResponse](../../models/operations/replacepasswordresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeletePassword

Permanently deletes a password, immediately revoking its access.
This operation cannot be undone.


### Example Usage

<!-- UsageSnippet language="go" operationID="deletePassword" method="delete" path="/v2/directories/{directoryId}/principals/{principalId}/basic-auths/{basicAuthId}/passwords/{passwordId}" -->
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

    res, err := s.KongPrincipalCredential.DeletePassword(ctx, operations.DeletePasswordRequest{
        DirectoryID: "308859d5-2425-4f6b-a349-fd448ac4a360",
        PrincipalID: "e398ccb2-fde1-4ce8-810c-f759ae7e99f7",
        BasicAuthID: "30b35ab9-f87b-4a26-8a3a-07bfadf5477b",
        PasswordID: "fe99fbdf-a90c-4df0-a4a6-fc366bfc58de",
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
| `request`                                                                            | [operations.DeletePasswordRequest](../../models/operations/deletepasswordrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.DeletePasswordResponse](../../models/operations/deletepasswordresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |