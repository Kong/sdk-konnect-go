# TeamMembership
(*TeamMembership*)

## Overview

### Available Operations

* [ListTeamUsers](#listteamusers) - List Team Users
* [AddUserToTeam](#addusertoteam) - Add User
* [RemoveUserFromTeam](#removeuserfromteam) - Remove User
* [ListUserTeams](#listuserteams) - List User Teams

## ListTeamUsers

Returns a paginated list of users that belong to the team specified in the path parameter.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-team-users" method="get" path="/v3/teams/{teamId}/users" -->
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

    res, err := s.TeamMembership.ListTeamUsers(ctx, operations.ListTeamUsersRequest{
        TeamID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &operations.ListTeamUsersQueryParamFilter{
            Active: sdkkonnectgo.Pointer(true),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListTeamUsersRequest](../../models/operations/listteamusersrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ListTeamUsersResponse](../../models/operations/listteamusersresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## AddUserToTeam

Adds a user to a team.

### Example Usage

<!-- UsageSnippet language="go" operationID="add-user-to-team" method="post" path="/v3/teams/{teamId}/users" -->
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

    res, err := s.TeamMembership.AddUserToTeam(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", &components.AddUserToTeam{
        UserID: "df120cb4-f60b-47bc-a2f8-6a28e6a3c63b",
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

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           | Example                                                               |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ctx`                                                                 | [context.Context](https://pkg.go.dev/context#Context)                 | :heavy_check_mark:                                                    | The context to use for the request.                                   |                                                                       |
| `teamID`                                                              | *string*                                                              | :heavy_check_mark:                                                    | ID of the team.                                                       | d32d905a-ed33-46a3-a093-d8f536af9a8a                                  |
| `addUserToTeam`                                                       | [*components.AddUserToTeam](../../models/components/addusertoteam.md) | :heavy_minus_sign:                                                    | The request schema for adding a user to a team.                       |                                                                       |
| `opts`                                                                | [][operations.Option](../../models/operations/option.md)              | :heavy_minus_sign:                                                    | The options for this request.                                         |                                                                       |

### Response

**[*operations.AddUserToTeamResponse](../../models/operations/addusertoteamresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.ConflictError   | 409                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## RemoveUserFromTeam

Removes a user from a team.
If the user was removed, returns a 204 empty response. Returns 404 if the user or team were not found.

### Example Usage

<!-- UsageSnippet language="go" operationID="remove-user-from-team" method="delete" path="/v3/teams/{teamId}/users/{userId}" -->
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

    res, err := s.TeamMembership.RemoveUserFromTeam(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", "d32d905a-ed33-46a3-a093-d8f536af9a8a")
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
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | User ID                                                  | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `teamID`                                                 | *string*                                                 | :heavy_check_mark:                                       | Team ID.                                                 | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RemoveUserFromTeamResponse](../../models/operations/removeuserfromteamresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ListUserTeams

Returns a paginated list of a teams that the user belongs to.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-user-teams" method="get" path="/v3/users/{userId}/teams" -->
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

    res, err := s.TeamMembership.ListUserTeams(ctx, operations.ListUserTeamsRequest{
        UserID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TeamCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListUserTeamsRequest](../../models/operations/listuserteamsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ListUserTeamsResponse](../../models/operations/listuserteamsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |