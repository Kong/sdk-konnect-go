# TokenVaultCredentials

## Overview

Token vault credentials are the credentials that a principal has obtained by authorizing a token-vault provider. This resource is read-only: providers are authorized through the enrollment
flow, not through this API, so only listing, retrieving, and de-authorizing (delete) are supported.


### Available Operations

* [ListVaultCredentials](#listvaultcredentials) - List vault credentials for a principal
* [GetVaultCredential](#getvaultcredential) - Get a vault credential
* [DeleteVaultCredential](#deletevaultcredential) - De-authorize a provider from a principal

## ListVaultCredentials

Returns a paginated list of token-vault credentials that the principal has obtained by
authorizing providers.


### Example Usage

<!-- UsageSnippet language="go" operationID="listVaultCredentials" method="get" path="/v2/directories/{directoryId}/principals/{principalId}/vault-credentials" -->
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

    res, err := s.TokenVaultCredentials.ListVaultCredentials(ctx, operations.ListVaultCredentialsRequest{
        DirectoryID: "752738ff-a5f9-44fb-93cc-f2a922fdc9b5",
        PrincipalID: "8900c3f8-4e61-466f-a2c8-3986bb3be10c",
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
    if res.ListVaultCredentials != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListVaultCredentialsRequest](../../models/operations/listvaultcredentialsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListVaultCredentialsResponse](../../models/operations/listvaultcredentialsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetVaultCredential

Retrieves detailed information about the credentials the principal has obtained from a single provider.

### Example Usage

<!-- UsageSnippet language="go" operationID="getVaultCredential" method="get" path="/v2/directories/{directoryId}/principals/{principalId}/vault-credentials/{vaultCredentialId}" -->
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

    res, err := s.TokenVaultCredentials.GetVaultCredential(ctx, operations.GetVaultCredentialRequest{
        DirectoryID: "ac337d71-dc84-4a03-8e30-be3c4b85d5a4",
        PrincipalID: "b43a2ed0-130e-4500-8bca-a72a4c1bc91a",
        VaultCredentialID: "0392be87-8ed9-4911-9fee-99e68af87f1c",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VaultCredential != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetVaultCredentialRequest](../../models/operations/getvaultcredentialrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetVaultCredentialResponse](../../models/operations/getvaultcredentialresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteVaultCredential

Removes the principal's link to the provider, deleting its stored credential. The principal must re-enroll to link the provider again.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteVaultCredential" method="delete" path="/v2/directories/{directoryId}/principals/{principalId}/vault-credentials/{vaultCredentialId}" -->
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

    res, err := s.TokenVaultCredentials.DeleteVaultCredential(ctx, operations.DeleteVaultCredentialRequest{
        DirectoryID: "70c4c01d-53e9-40f0-8eb6-cefa27831fa0",
        PrincipalID: "9460f6ec-4d54-482d-ac10-1898bb7d11e3",
        VaultCredentialID: "fa487968-82cc-4e70-9601-56f617bb227a",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.DeleteVaultCredentialRequest](../../models/operations/deletevaultcredentialrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.DeleteVaultCredentialResponse](../../models/operations/deletevaultcredentialresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |