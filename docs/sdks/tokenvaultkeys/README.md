# TokenVaultKeys

## Overview

Vault keys are the public encryption keys used to protect token vault data for
a token-vault enabled directory.


### Available Operations

* [CreateVaultKey](#createvaultkey) - Create a vault key
* [ListVaultKeys](#listvaultkeys) - List vault keys
* [GetVaultKey](#getvaultkey) - Get a vault key
* [UpdateVaultKey](#updatevaultkey) - Update a vault key
* [DeleteVaultKey](#deletevaultkey) - Delete a vault key

## CreateVaultKey

Adds a PEM-encoded RSA or EC public key used to encrypt token vault data for a
vault-enabled directory. Keys must be at least 2048 bits (RSA) or 256 bits (EC).
A directory can hold at most 5 vault keys.


### Example Usage

<!-- UsageSnippet language="go" operationID="createVaultKey" method="post" path="/v2/directories/{directoryId}/vault/keys" -->
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

    res, err := s.TokenVaultKeys.CreateVaultKey(ctx, "ff61fb67-4add-4afb-9ed0-b1d22f6a37d4", components.CreateVaultKeyBody{
        Name: "<value>",
        PublicKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...\n-----END PUBLIC KEY-----",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VaultKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `directoryID`                                                                  | `string`                                                                       | :heavy_check_mark:                                                             | ID of the directory                                                            |
| `createVaultKeyBody`                                                           | [components.CreateVaultKeyBody](../../models/components/createvaultkeybody.md) | :heavy_check_mark:                                                             | Vault key to be created                                                        |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CreateVaultKeyResponse](../../models/operations/createvaultkeyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListVaultKeys

Returns a paginated list of the public encryption keys registered for a
vault-enabled directory.


### Example Usage

<!-- UsageSnippet language="go" operationID="listVaultKeys" method="get" path="/v2/directories/{directoryId}/vault/keys" -->
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

    res, err := s.TokenVaultKeys.ListVaultKeys(ctx, operations.ListVaultKeysRequest{
        DirectoryID: "ad15ce1d-7b83-44d3-b0aa-816466d981e5",
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
    if res.ListVaultKeys != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListVaultKeysRequest](../../models/operations/listvaultkeysrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ListVaultKeysResponse](../../models/operations/listvaultkeysresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetVaultKey

Retrieves a single vault key by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getVaultKey" method="get" path="/v2/directories/{directoryId}/vault/keys/{keyId}" -->
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

    res, err := s.TokenVaultKeys.GetVaultKey(ctx, "ceb6afaf-edd9-4c02-bfd7-7d8092981146", "487f1918-5b62-4606-8c20-c22eaa880667")
    if err != nil {
        log.Fatal(err)
    }
    if res.VaultKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `directoryID`                                            | `string`                                                 | :heavy_check_mark:                                       | ID of the directory                                      |
| `keyID`                                                  | `string`                                                 | :heavy_check_mark:                                       | ID of the vault key                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetVaultKeyResponse](../../models/operations/getvaultkeyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateVaultKey

Partially updates an existing vault key. Only the key's name can be changed;
the public key material is immutable. A PUT operation is not offered because
key material cannot be replaced.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateVaultKey" method="patch" path="/v2/directories/{directoryId}/vault/keys/{keyId}" -->
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

    res, err := s.TokenVaultKeys.UpdateVaultKey(ctx, operations.UpdateVaultKeyRequest{
        DirectoryID: "cdab025a-837c-4e5f-b90b-121917dd8ce5",
        KeyID: "773e04c3-0b36-4d3c-95c5-936d0822ef4b",
        PatchVaultKeyBody: components.PatchVaultKeyBody{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VaultKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateVaultKeyRequest](../../models/operations/updatevaultkeyrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateVaultKeyResponse](../../models/operations/updatevaultkeyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteVaultKey

Deletes a vault key from a directory. Deleting a key does not affect
credentials that were already encrypted with it; the key simply stops
receiving newly encrypted credentials.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteVaultKey" method="delete" path="/v2/directories/{directoryId}/vault/keys/{keyId}" -->
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

    res, err := s.TokenVaultKeys.DeleteVaultKey(ctx, "6aab3107-a5ca-4704-8e05-3b7babfe8984", "a423959d-b618-46ca-8f88-0e3dc8cc953d")
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
| `keyID`                                                  | `string`                                                 | :heavy_check_mark:                                       | ID of the vault key                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteVaultKeyResponse](../../models/operations/deletevaultkeyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |