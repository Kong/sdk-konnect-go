# SystemAccountsAccessTokens

## Overview

### Available Operations

* [GetSystemAccountIDAccessTokens](#getsystemaccountidaccesstokens) - List System Account Access Tokens
* [PostSystemAccountsIDAccessTokens](#postsystemaccountsidaccesstokens) - Create System Account Access Token
* [GetSystemAccountsIDAccessTokensID](#getsystemaccountsidaccesstokensid) - Get a System Account Access Token
* [PatchSystemAccountsIDAccessTokensID](#patchsystemaccountsidaccesstokensid) - Update System Account Access Token
* [DeleteSystemAccountsIDAccessTokensID](#deletesystemaccountsidaccesstokensid) - Delete System Account Access Token

## GetSystemAccountIDAccessTokens

Returns the access tokens for the specified system account. Returns 400 if any filter parameters are invalid.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-system-account-id-access-tokens" method="get" path="/v3/system-accounts/{accountId}/access-tokens" -->
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

    res, err := s.SystemAccountsAccessTokens.GetSystemAccountIDAccessTokens(ctx, operations.GetSystemAccountIDAccessTokensRequest{
        AccountID: "<id>",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SystemAccountAccessTokenCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetSystemAccountIDAccessTokensRequest](../../models/operations/getsystemaccountidaccesstokensrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.GetSystemAccountIDAccessTokensResponse](../../models/operations/getsystemaccountidaccesstokensresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PostSystemAccountsIDAccessTokens

Creates an access token for the specified system account (SA). The access token can be used for authenticating API and CLI requests. The token will only be displayed once on creation. Returns a 409 if the system account already has a token with the same name.

### Example Usage

<!-- UsageSnippet language="go" operationID="post-system-accounts-id-access-tokens" method="post" path="/v3/system-accounts/{accountId}/access-tokens" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.SystemAccountsAccessTokens.PostSystemAccountsIDAccessTokens(ctx, "<id>", &components.CreateSystemAccountAccessToken{
        Name: "Sample Access Token",
        ExpiresAt: types.MustTimeFromString("2019-08-24T14:15:22Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SystemAccountAccessTokenCreated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |
| `accountID`                                                                                             | *string*                                                                                                | :heavy_check_mark:                                                                                      | ID of the system account.                                                                               |
| `createSystemAccountAccessToken`                                                                        | [*components.CreateSystemAccountAccessToken](../../models/components/createsystemaccountaccesstoken.md) | :heavy_minus_sign:                                                                                      | The request body to create a system account access token.                                               |
| `opts`                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                | :heavy_minus_sign:                                                                                      | The options for this request.                                                                           |

### Response

**[*operations.PostSystemAccountsIDAccessTokensResponse](../../models/operations/postsystemaccountsidaccesstokensresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetSystemAccountsIDAccessTokensID

Returns the system account (SA) access token for the SA Access Token ID specified as a path parameter.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-system-accounts-id-access-tokens-id" method="get" path="/v3/system-accounts/{accountId}/access-tokens/{tokenId}" -->
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

    res, err := s.SystemAccountsAccessTokens.GetSystemAccountsIDAccessTokensID(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SystemAccountAccessToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | ID of the system account.                                |
| `tokenID`                                                | *string*                                                 | :heavy_check_mark:                                       | ID of the system account access token.                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSystemAccountsIDAccessTokensIDResponse](../../models/operations/getsystemaccountsidaccesstokensidresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchSystemAccountsIDAccessTokensID

Updates the specified access token. Returns a 409 if the updated name is the same as another token belonging to the specified system user.

### Example Usage

<!-- UsageSnippet language="go" operationID="patch-system-accounts-id-access-tokens-id" method="patch" path="/v3/system-accounts/{accountId}/access-tokens/{tokenId}" -->
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

    res, err := s.SystemAccountsAccessTokens.PatchSystemAccountsIDAccessTokensID(ctx, operations.PatchSystemAccountsIDAccessTokensIDRequest{
        AccountID: "<id>",
        TokenID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SystemAccountAccessToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.PatchSystemAccountsIDAccessTokensIDRequest](../../models/operations/patchsystemaccountsidaccesstokensidrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.PatchSystemAccountsIDAccessTokensIDResponse](../../models/operations/patchsystemaccountsidaccesstokensidresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteSystemAccountsIDAccessTokensID

Deletes the specified token. Returns 404 if the token was not found.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-system-accounts-id-access-tokens-id" method="delete" path="/v3/system-accounts/{accountId}/access-tokens/{tokenId}" -->
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

    res, err := s.SystemAccountsAccessTokens.DeleteSystemAccountsIDAccessTokensID(ctx, "<id>", "<id>")
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
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | ID of the system account.                                |
| `tokenID`                                                | *string*                                                 | :heavy_check_mark:                                       | ID of the system account access token.                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteSystemAccountsIDAccessTokensIDResponse](../../models/operations/deletesystemaccountsidaccesstokensidresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |