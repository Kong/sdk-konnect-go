# SystemAccounts
(*SystemAccounts*)

## Overview

### Available Operations

* [GetSystemAccounts](#getsystemaccounts) - List System Accounts
* [PostSystemAccounts](#postsystemaccounts) - Create System Account
* [GetSystemAccountsID](#getsystemaccountsid) - Fetch System Account
* [PatchSystemAccountsID](#patchsystemaccountsid) - Update System Account
* [DeleteSystemAccountsID](#deletesystemaccountsid) - Delete System Account

## GetSystemAccounts

Returns an array of system accounts (SA) in the organization. Returns 400 if any filter parameters are invalid.

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.SystemAccounts.GetSystemAccounts(ctx, operations.GetSystemAccountsRequest{
        PageSize: sdkkonnectgo.Int64(10),
        PageNumber: sdkkonnectgo.Int64(1),
        Filter: &operations.Filter{
            KonnectManaged: sdkkonnectgo.Bool(true),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SystemAccountCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetSystemAccountsRequest](../../models/operations/getsystemaccountsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetSystemAccountsResponse](../../models/operations/getsystemaccountsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PostSystemAccounts

Creates a system account. Returns a 409 if a system account with the same name already exists.

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.SystemAccounts.PostSystemAccounts(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SystemAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.CreateSystemAccount](../../models/components/createsystemaccount.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.PostSystemAccountsResponse](../../models/operations/postsystemaccountsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetSystemAccountsID

Returns the system account (SA) for the SA ID specified as a path parameter.

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.SystemAccounts.GetSystemAccountsID(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SystemAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | ID of the system account.                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSystemAccountsIDResponse](../../models/operations/getsystemaccountsidresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchSystemAccountsID

Updates the specified system account. Returns a 409 if the updated name is the same as another system account in the organization.

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.SystemAccounts.PatchSystemAccountsID(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SystemAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |
| `accountID`                                                                       | *string*                                                                          | :heavy_check_mark:                                                                | ID of the system account.                                                         |
| `updateSystemAccount`                                                             | [*components.UpdateSystemAccount](../../models/components/updatesystemaccount.md) | :heavy_minus_sign:                                                                | The request schema for the update system account request.                         |
| `opts`                                                                            | [][operations.Option](../../models/operations/option.md)                          | :heavy_minus_sign:                                                                | The options for this request.                                                     |

### Response

**[*operations.PatchSystemAccountsIDResponse](../../models/operations/patchsystemaccountsidresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteSystemAccountsID

Deletes the specified system account. Returns 404 if the requested account was not found.

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.SystemAccounts.DeleteSystemAccountsID(ctx, "<value>")
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteSystemAccountsIDResponse](../../models/operations/deletesystemaccountsidresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |