# AuthServerScopes

## Overview

Scopes define the extent of access that an access token grants to a Client. The management API will give you the ability to create, configure and manage multiple Scopes per Auth Server, and restrict their usage by Client.

### Available Operations

* [ListAuthServerScopes](#listauthserverscopes) - List all auth server scopes
* [CreateAuthServerScope](#createauthserverscope) - Create a new auth server scope
* [GetAuthServerScope](#getauthserverscope) - Get an auth server scope
* [UpdateAuthServerScope](#updateauthserverscope) - Update an auth server scope
* [DeleteAuthServerScope](#deleteauthserverscope) - Delete an auth server scope

## ListAuthServerScopes

List all the scopes for a given auth server.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAuthServerScopes" method="get" path="/v1/auth-servers/{authServerId}/scopes" -->
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

    res, err := s.AuthServerScopes.ListAuthServerScopes(ctx, operations.ListAuthServerScopesRequest{
        AuthServerID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ScopeList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListAuthServerScopesRequest](../../models/operations/listauthserverscopesrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListAuthServerScopesResponse](../../models/operations/listauthserverscopesresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## CreateAuthServerScope

Create a new scope for a given auth server. The `name` attribute must be unique within the auth server.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAuthServerScope" method="post" path="/v1/auth-servers/{authServerId}/scopes" -->
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

    res, err := s.AuthServerScopes.CreateAuthServerScope(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.CreateScope{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Scope != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |                                                                  |
| `authServerID`                                                   | `string`                                                         | :heavy_check_mark:                                               | The auth server ID                                               | d32d905a-ed33-46a3-a093-d8f536af9a8a                             |
| `createScope`                                                    | [components.CreateScope](../../models/components/createscope.md) | :heavy_check_mark:                                               | Scope to be created                                              |                                                                  |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |                                                                  |

### Response

**[*operations.CreateAuthServerScopeResponse](../../models/operations/createauthserverscoperesponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.ConflictError   | 409                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## GetAuthServerScope

Get a scope for given auth server.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAuthServerScope" method="get" path="/v1/auth-servers/{authServerId}/scopes/{scopeId}" -->
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

    res, err := s.AuthServerScopes.GetAuthServerScope(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", "c5e12516-182c-4928-ae04-05374b3b1cca")
    if err != nil {
        log.Fatal(err)
    }
    if res.Scope != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `authServerID`                                           | `string`                                                 | :heavy_check_mark:                                       | The auth server ID                                       | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `scopeID`                                                | `string`                                                 | :heavy_check_mark:                                       | The scope ID                                             | c5e12516-182c-4928-ae04-05374b3b1cca                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAuthServerScopeResponse](../../models/operations/getauthserverscoperesponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## UpdateAuthServerScope

Update a scope for a given auth server.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAuthServerScope" method="patch" path="/v1/auth-servers/{authServerId}/scopes/{scopeId}" -->
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

    res, err := s.AuthServerScopes.UpdateAuthServerScope(ctx, operations.UpdateAuthServerScopeRequest{
        AuthServerID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        ScopeID: "c5e12516-182c-4928-ae04-05374b3b1cca",
        UpdateScope: components.UpdateScope{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Scope != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateAuthServerScopeRequest](../../models/operations/updateauthserverscoperequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateAuthServerScopeResponse](../../models/operations/updateauthserverscoperesponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.ConflictError   | 409                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## DeleteAuthServerScope

Delete a scope for a given auth server. If the scope is included in any clients, it will be removed from those clients.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteAuthServerScope" method="delete" path="/v1/auth-servers/{authServerId}/scopes/{scopeId}" -->
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

    res, err := s.AuthServerScopes.DeleteAuthServerScope(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", "c5e12516-182c-4928-ae04-05374b3b1cca")
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
| `authServerID`                                           | `string`                                                 | :heavy_check_mark:                                       | The auth server ID                                       | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `scopeID`                                                | `string`                                                 | :heavy_check_mark:                                       | The scope ID                                             | c5e12516-182c-4928-ae04-05374b3b1cca                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAuthServerScopeResponse](../../models/operations/deleteauthserverscoperesponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |