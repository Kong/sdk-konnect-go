# TeamMembership

## Overview

### Available Operations

* [ListTeamUsers](#listteamusers) - List Team Users
* [AddUserToTeam](#addusertoteam) - Add User
* [RemoveUserFromTeam](#removeuserfromteam) - Remove User
* [ListUserTeams](#listuserteams) - List User Teams

## ListTeamUsers

Returns a paginated list of users that belong to the team specified in the path parameter.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-team-users" method="get" path="/v3/teams/{teamId}/users" example="User Collection Response" -->
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

### Example Usage: Authentication Settings cannot be all Disabled

<!-- UsageSnippet language="go" operationID="add-user-to-team" method="post" path="/v3/teams/{teamId}/users" example="Authentication Settings cannot be all Disabled" -->
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
### Example Usage: Cannot be Blank

<!-- UsageSnippet language="go" operationID="add-user-to-team" method="post" path="/v3/teams/{teamId}/users" example="Cannot be Blank" -->
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
### Example Usage: Example request body

<!-- UsageSnippet language="go" operationID="add-user-to-team" method="post" path="/v3/teams/{teamId}/users" example="Example request body" -->
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
### Example Usage: IdP configuration is required

<!-- UsageSnippet language="go" operationID="add-user-to-team" method="post" path="/v3/teams/{teamId}/users" example="IdP configuration is required" -->
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
### Example Usage: Invalid ID format

<!-- UsageSnippet language="go" operationID="add-user-to-team" method="post" path="/v3/teams/{teamId}/users" example="Invalid ID format" -->
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
### Example Usage: Must be a valid UUID v4

<!-- UsageSnippet language="go" operationID="add-user-to-team" method="post" path="/v3/teams/{teamId}/users" example="Must be a valid UUID v4" -->
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
### Example Usage: Not Found

<!-- UsageSnippet language="go" operationID="add-user-to-team" method="post" path="/v3/teams/{teamId}/users" example="Not Found" -->
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
### Example Usage: OIDC needs an IdP configuration

<!-- UsageSnippet language="go" operationID="add-user-to-team" method="post" path="/v3/teams/{teamId}/users" example="OIDC needs an IdP configuration" -->
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
### Example Usage: Password Complexity

<!-- UsageSnippet language="go" operationID="add-user-to-team" method="post" path="/v3/teams/{teamId}/users" example="Password Complexity" -->
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
### Example Usage: Request Format is Invalid

<!-- UsageSnippet language="go" operationID="add-user-to-team" method="post" path="/v3/teams/{teamId}/users" example="Request Format is Invalid" -->
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
### Example Usage: System teams cannot be modified

<!-- UsageSnippet language="go" operationID="add-user-to-team" method="post" path="/v3/teams/{teamId}/users" example="System teams cannot be modified" -->
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
### Example Usage: Unsupported filter operation

<!-- UsageSnippet language="go" operationID="add-user-to-team" method="post" path="/v3/teams/{teamId}/users" example="Unsupported filter operation" -->
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
### Example Usage: User is Already Active

<!-- UsageSnippet language="go" operationID="add-user-to-team" method="post" path="/v3/teams/{teamId}/users" example="User is Already Active" -->
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

<!-- UsageSnippet language="go" operationID="list-user-teams" method="get" path="/v3/users/{userId}/teams" example="Team Collection" -->
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