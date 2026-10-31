# AuthServer

## Overview

Auth Servers expose an OAuth 2.0 and OpenID Connect server interface for generating access tokens. The management API will give you the ability to create, configure and manage multiple Auth Servers per Konnect organization. Auth Servers are a regional Konnect entity.

### Available Operations

* [ListAuthServers](#listauthservers) - List all auth servers
* [CreateAuthServer](#createauthserver) - Create a new auth server
* [GetAuthServer](#getauthserver) - Get an auth server
* [UpdateAuthServer](#updateauthserver) - Update an auth server
* [DeleteAuthServer](#deleteauthserver) - Delete an auth server

## ListAuthServers

List all available auth servers.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAuthServers" method="get" path="/v1/auth-servers" -->
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

    res, err := s.AuthServer.ListAuthServers(ctx, operations.ListAuthServersRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthServerList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListAuthServersRequest](../../models/operations/listauthserversrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListAuthServersResponse](../../models/operations/listauthserversresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateAuthServer

Create a new auth server. Each auth server has a unique, randomly generated, public issuer URL.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAuthServer" method="post" path="/v1/auth-servers" -->
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

    res, err := s.AuthServer.CreateAuthServer(ctx, components.CreateAuthServer{
        Name: "<value>",
        Audience: "<value>",
        Labels: map[string]string{
            "env": "test",
        },
        TrustedOrigins: []string{
            "https://example.com",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthServer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.CreateAuthServer](../../models/components/createauthserver.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.CreateAuthServerResponse](../../models/operations/createauthserverresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.ConflictError   | 409                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## GetAuthServer

Get an auth server.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAuthServer" method="get" path="/v1/auth-servers/{authServerId}" -->
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

    res, err := s.AuthServer.GetAuthServer(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthServer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `authServerID`                                           | `string`                                                 | :heavy_check_mark:                                       | The auth server ID                                       | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAuthServerResponse](../../models/operations/getauthserverresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## UpdateAuthServer

Update an auth server.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAuthServer" method="patch" path="/v1/auth-servers/{authServerId}" -->
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

    res, err := s.AuthServer.UpdateAuthServer(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.UpdateAuthServer{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
        TrustedOrigins: []string{
            "https://example.com",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthServer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |                                                                            |
| `authServerID`                                                             | `string`                                                                   | :heavy_check_mark:                                                         | The auth server ID                                                         | d32d905a-ed33-46a3-a093-d8f536af9a8a                                       |
| `updateAuthServer`                                                         | [components.UpdateAuthServer](../../models/components/updateauthserver.md) | :heavy_check_mark:                                                         | Auth server to be updated                                                  |                                                                            |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |                                                                            |

### Response

**[*operations.UpdateAuthServerResponse](../../models/operations/updateauthserverresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.ConflictError   | 409                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## DeleteAuthServer

Delete an auth server. If force=true, delete the auth server and all its associated resources. If force=false, only allow deletion if no resources are associated with the auth server. This action is irreversible.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteAuthServer" method="delete" path="/v1/auth-servers/{authServerId}" -->
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

    res, err := s.AuthServer.DeleteAuthServer(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", operations.DeleteAuthServerForceFalse.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    | Example                                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                          | :heavy_check_mark:                                                                                                                                                             | The context to use for the request.                                                                                                                                            |                                                                                                                                                                                |
| `authServerID`                                                                                                                                                                 | `string`                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | The auth server ID                                                                                                                                                             | d32d905a-ed33-46a3-a093-d8f536af9a8a                                                                                                                                           |
| `forceDestroy`                                                                                                                                                                 | [*operations.DeleteAuthServerForce](../../models/operations/deleteauthserverforce.md)                                                                                          | :heavy_minus_sign:                                                                                                                                                             | If true, delete the specified auth server and all its associated resources. If false, only allow deletion if no clients, scopes or claims are associated with the auth server. |                                                                                                                                                                                |
| `opts`                                                                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                                                                       | :heavy_minus_sign:                                                                                                                                                             | The options for this request.                                                                                                                                                  |                                                                                                                                                                                |

### Response

**[*operations.DeleteAuthServerResponse](../../models/operations/deleteauthserverresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |