# ControlPlaneGroups
(*ControlPlaneGroups*)

## Overview

### Available Operations

* [GetControlPlanesIDGroupMemberStatus](#getcontrolplanesidgroupmemberstatus) - Control Plane Group Member Status
* [GetControlPlanesIDGroupMemberships](#getcontrolplanesidgroupmemberships) - List Control Plane Group Memberships
* [PutControlPlanesIDGroupMemberships](#putcontrolplanesidgroupmemberships) - Upsert Control Plane Group Members
* [PostControlPlanesIDGroupMembershipsAdd](#postcontrolplanesidgroupmembershipsadd) - Add Control Plane Group Members
* [PostControlPlanesIDGroupMembershipsRemove](#postcontrolplanesidgroupmembershipsremove) - Remove Control Plane Group Members
* [GetControlPlanesIDGroupStatus](#getcontrolplanesidgroupstatus) - Get Control Plane Group Status

## GetControlPlanesIDGroupMemberStatus

Determines the group membership status of a control plane.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ControlPlaneGroups.GetControlPlanesIDGroupMemberStatus(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.GroupMemberStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | ID of a control plane                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetControlPlanesIDGroupMemberStatusResponse](../../models/operations/getcontrolplanesidgroupmemberstatusresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.BadRequestError    | 400                          | application/problem+json     |
| sdkerrors.UnauthorizedError  | 401                          | application/problem+json     |
| sdkerrors.ForbiddenError     | 403                          | application/problem+json     |
| sdkerrors.NotFoundError      | 404                          | application/problem+json     |
| sdkerrors.ServiceUnavailable | 503                          | application/problem+json     |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## GetControlPlanesIDGroupMemberships

Returns an array of control planes that are a member of this control plane group.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ControlPlaneGroups.GetControlPlanesIDGroupMemberships(ctx, operations.GetControlPlanesIDGroupMembershipsRequest{
        ID: "<id>",
        PageSize: sdkkonnectgo.Int64(10),
        PageAfter: sdkkonnectgo.String("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListGroupMemberships != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetControlPlanesIDGroupMembershipsRequest](../../models/operations/getcontrolplanesidgroupmembershipsrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.GetControlPlanesIDGroupMembershipsResponse](../../models/operations/getcontrolplanesidgroupmembershipsresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.BadRequestError    | 400                          | application/problem+json     |
| sdkerrors.UnauthorizedError  | 401                          | application/problem+json     |
| sdkerrors.ForbiddenError     | 403                          | application/problem+json     |
| sdkerrors.NotFoundError      | 404                          | application/problem+json     |
| sdkerrors.ServiceUnavailable | 503                          | application/problem+json     |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## PutControlPlanesIDGroupMemberships

Adds one or more control planes as a member of a control plane group.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ControlPlaneGroups.PutControlPlanesIDGroupMemberships(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `id`                                                                                           | *string*                                                                                       | :heavy_check_mark:                                                                             | ID of a control plane group                                                                    |
| `groupMembership`                                                                              | [*components.GroupMembership](../../models/components/groupmembership.md)                      | :heavy_minus_sign:                                                                             | Request body for upserting a list of child control planes to a control plane group membership. |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PutControlPlanesIDGroupMembershipsResponse](../../models/operations/putcontrolplanesidgroupmembershipsresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.BadRequestError    | 400                          | application/problem+json     |
| sdkerrors.UnauthorizedError  | 401                          | application/problem+json     |
| sdkerrors.ForbiddenError     | 403                          | application/problem+json     |
| sdkerrors.NotFoundError      | 404                          | application/problem+json     |
| sdkerrors.ServiceUnavailable | 503                          | application/problem+json     |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## PostControlPlanesIDGroupMembershipsAdd

Adds one or more control planes as a member of a control plane group.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ControlPlaneGroups.PostControlPlanesIDGroupMembershipsAdd(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |
| `id`                                                                                        | *string*                                                                                    | :heavy_check_mark:                                                                          | ID of a control plane group                                                                 |
| `groupMembership`                                                                           | [*components.GroupMembership](../../models/components/groupmembership.md)                   | :heavy_minus_sign:                                                                          | Request body for adding a list of child control planes to a control plane group membership. |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |

### Response

**[*operations.PostControlPlanesIDGroupMembershipsAddResponse](../../models/operations/postcontrolplanesidgroupmembershipsaddresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.BadRequestError    | 400                          | application/problem+json     |
| sdkerrors.UnauthorizedError  | 401                          | application/problem+json     |
| sdkerrors.ForbiddenError     | 403                          | application/problem+json     |
| sdkerrors.NotFoundError      | 404                          | application/problem+json     |
| sdkerrors.ServiceUnavailable | 503                          | application/problem+json     |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## PostControlPlanesIDGroupMembershipsRemove

Removes one or more control planes from the members of a control plane group.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ControlPlaneGroups.PostControlPlanesIDGroupMembershipsRemove(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                       | Type                                                                                            | Required                                                                                        | Description                                                                                     |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `ctx`                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                           | :heavy_check_mark:                                                                              | The context to use for the request.                                                             |
| `id`                                                                                            | *string*                                                                                        | :heavy_check_mark:                                                                              | ID of a control plane group                                                                     |
| `groupMembership`                                                                               | [*components.GroupMembership](../../models/components/groupmembership.md)                       | :heavy_minus_sign:                                                                              | Request body for removing a list of child control planes from a control plane group membership. |
| `opts`                                                                                          | [][operations.Option](../../models/operations/option.md)                                        | :heavy_minus_sign:                                                                              | The options for this request.                                                                   |

### Response

**[*operations.PostControlPlanesIDGroupMembershipsRemoveResponse](../../models/operations/postcontrolplanesidgroupmembershipsremoveresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.BadRequestError    | 400                          | application/problem+json     |
| sdkerrors.UnauthorizedError  | 401                          | application/problem+json     |
| sdkerrors.ForbiddenError     | 403                          | application/problem+json     |
| sdkerrors.NotFoundError      | 404                          | application/problem+json     |
| sdkerrors.ServiceUnavailable | 503                          | application/problem+json     |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## GetControlPlanesIDGroupStatus

Returns the status of a control plane group, including existing conflicts.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ControlPlaneGroups.GetControlPlanesIDGroupStatus(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetGroupStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | ID of a control plane group                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetControlPlanesIDGroupStatusResponse](../../models/operations/getcontrolplanesidgroupstatusresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.BadRequestError    | 400                          | application/problem+json     |
| sdkerrors.UnauthorizedError  | 401                          | application/problem+json     |
| sdkerrors.ForbiddenError     | 403                          | application/problem+json     |
| sdkerrors.NotFoundError      | 404                          | application/problem+json     |
| sdkerrors.ServiceUnavailable | 503                          | application/problem+json     |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |