# Me

## Overview

### Available Operations

* [GetUsersMe](#getusersme) - Get My User Account
* [DeleteUsersMe](#deleteusersme) - Delete My User Account
* [PatchUsersMe](#patchusersme) - Update My User Account
* [GetUsersMePermissions](#getusersmepermissions) - Get My Permissions
* [RetrieveUsersMePermissions](#retrieveusersmepermissions) - Retrieve My Permissions
* [GetOrganizationsMe](#getorganizationsme) - Get My Organization
* [UpdateOrganizationsMe](#updateorganizationsme) - Update My Organization

## GetUsersMe

Returns the user account for the user identified in the token of the request.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-users-me" method="get" path="/v3/users/me" -->
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

    res, err := s.Me.GetUsersMe(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetUsersMeResponse](../../models/operations/getusersmeresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteUsersMe

Deletes the user account for the user identified in the token of the request.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-users-me" method="delete" path="/v3/users/me" -->
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
            KonnectAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.Me.DeleteUsersMe(ctx)
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteUsersMeResponse](../../models/operations/deleteusersmeresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchUsersMe

Updates the user account for the user identified in the token of the request.

### Example Usage

<!-- UsageSnippet language="go" operationID="patch-users-me" method="patch" path="/v3/users/me" -->
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
            KonnectAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.Me.PatchUsersMe(ctx, &components.UpdateUser{
        FullName: sdkkonnectgo.Pointer("James C. Woods"),
        PreferredName: sdkkonnectgo.Pointer("Jimmy"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [components.UpdateUser](../../models/components/updateuser.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*operations.PatchUsersMeResponse](../../models/operations/patchusersmeresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetUsersMePermissions

Returns the permissions for the current user

### Example Usage

<!-- UsageSnippet language="go" operationID="get-users-me-permissions" method="get" path="/v3/users/me/permissions" -->
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
            KonnectAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.Me.GetUsersMePermissions(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UserPermissions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                             | Type                                                                                                                  | Required                                                                                                              | Description                                                                                                           |
| --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                 | :heavy_check_mark:                                                                                                    | The context to use for the request.                                                                                   |
| `filter`                                                                                                              | [*operations.GetUsersMePermissionsQueryParamFilter](../../models/operations/getusersmepermissionsqueryparamfilter.md) | :heavy_minus_sign:                                                                                                    | Filter permissions returned in the response.                                                                          |
| `opts`                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                              | :heavy_minus_sign:                                                                                                    | The options for this request.                                                                                         |

### Response

**[*operations.GetUsersMePermissionsResponse](../../models/operations/getusersmepermissionsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## RetrieveUsersMePermissions

Returns the permissions for the current user, grouped under caller-supplied keys. Each key declares its own filters, and every permission matching any of that key's filters is returned under the key. The same permission may be returned under more than one key. A key with no filters matches every permission.

### Example Usage

<!-- UsageSnippet language="go" operationID="retrieve-users-me-permissions" method="post" path="/v3/users/me/retrieve-permissions" -->
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

    res, err := s.Me.RetrieveUsersMePermissions(ctx, components.RetrieveMyPermissionsBatchRequest{
        Keys: []components.PermissionKey{
            components.PermissionKey{
                Key: "control_planes",
                Filters: []components.PermissionFilter{
                    components.PermissionFilter{
                        Resource: sdkkonnectgo.Pointer("runtimegroups/*"),
                        Service: sdkkonnectgo.Pointer("reg"),
                        Region: sdkkonnectgo.Pointer("us"),
                        TopLevel: sdkkonnectgo.Pointer(true),
                        Actions: []string{
                            "read",
                        },
                    },
                },
            },
        },
    }, &components.RetrieveMyPermissionsCursorPageQuery{
        Size: sdkkonnectgo.Pointer[int64](10),
        After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RetrieveMyPermissionsBatchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                           | Type                                                                                                                | Required                                                                                                            | Description                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                               | :heavy_check_mark:                                                                                                  | The context to use for the request.                                                                                 |
| `retrieveMyPermissionsBatchRequest`                                                                                 | [components.RetrieveMyPermissionsBatchRequest](../../models/components/retrievemypermissionsbatchrequest.md)        | :heavy_check_mark:                                                                                                  | The keys, and the filters for each key, to group the caller's permissions under.                                    |
| `page`                                                                                                              | [*components.RetrieveMyPermissionsCursorPageQuery](../../models/components/retrievemypermissionscursorpagequery.md) | :heavy_minus_sign:                                                                                                  | Determines which page of the collection to retrieve.                                                                |
| `opts`                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                            | :heavy_minus_sign:                                                                                                  | The options for this request.                                                                                       |

### Response

**[*operations.RetrieveUsersMePermissionsResponse](../../models/operations/retrieveusersmepermissionsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetOrganizationsMe

Returns the organization of the user identified in the token of the request.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-organizations-me" method="get" path="/v3/organizations/me" -->
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

    res, err := s.Me.GetOrganizationsMe(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.MeOrganization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetOrganizationsMeResponse](../../models/operations/getorganizationsmeresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateOrganizationsMe

Updates the current user's organization. When updating the owner, the new owner
must be an organization admin.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-organizations-me" method="patch" path="/v3/organizations/me" -->
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

    res, err := s.Me.UpdateOrganizationsMe(ctx, &components.UpdateOrganization{
        OwnerID: sdkkonnectgo.Pointer("d9c56d92-0339-4dfb-be0c-2253162edba9"),
        Name: sdkkonnectgo.Pointer("Kong Inc"),
        MfaEnabled: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MeOrganization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.UpdateOrganization](../../models/components/updateorganization.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.UpdateOrganizationsMeResponse](../../models/operations/updateorganizationsmeresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |