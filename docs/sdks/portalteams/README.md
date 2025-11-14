# PortalTeams
(*PortalTeams*)

## Overview

APIs related to configuration of Konnect Developer Portal developer teams.

### Available Operations

* [ListPortalTeams](#listportalteams) - List Teams
* [CreatePortalTeam](#createportalteam) - Create Team
* [GetPortalTeam](#getportalteam) - Get Team
* [UpdatePortalTeam](#updateportalteam) - Update Team
* [DeletePortalTeam](#deleteportalteam) - Delete Team

## ListPortalTeams

Lists the developer teams in a portal. Each team can contain any developer and developers can be part of multiple teams.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-portal-teams" method="get" path="/v3/portals/{portalId}/teams" -->
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

    res, err := s.PortalTeams.ListPortalTeams(ctx, operations.ListPortalTeamsRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPortalTeamsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListPortalTeamsRequest](../../models/operations/listportalteamsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListPortalTeamsResponse](../../models/operations/listportalteamsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreatePortalTeam

Creates a developer team in a portal. Developers can be added to teams to provide RBAC access to API products. Teams can be assigned roles that grant permissions to perform an action on a resource.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-portal-team" method="post" path="/v3/portals/{portalId}/teams" -->
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

    res, err := s.PortalTeams.CreatePortalTeam(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalCreateTeamRequest{
        Name: "IDM - Developers",
        Description: sdkkonnectgo.Pointer("The Identity Management (IDM) team."),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalTeamResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               | Example                                                                                   |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |                                                                                           |
| `portalID`                                                                                | *string*                                                                                  | :heavy_check_mark:                                                                        | ID of the portal.                                                                         | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                      |
| `portalCreateTeamRequest`                                                                 | [*components.PortalCreateTeamRequest](../../models/components/portalcreateteamrequest.md) | :heavy_minus_sign:                                                                        | Create a team in a portal.                                                                | {<br/>"name": "IDM - Developers",<br/>"description": "The Identity Management (IDM) team."<br/>} |
| `opts`                                                                                    | [][operations.Option](../../models/operations/option.md)                                  | :heavy_minus_sign:                                                                        | The options for this request.                                                             |                                                                                           |

### Response

**[*operations.CreatePortalTeamResponse](../../models/operations/createportalteamresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPortalTeam

Get an individual team.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-portal-team" method="get" path="/v3/portals/{portalId}/teams/{teamId}" -->
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

    res, err := s.PortalTeams.GetPortalTeam(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", "f32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalTeamResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `teamID`                                                 | *string*                                                 | :heavy_check_mark:                                       | ID of the team.                                          | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `portalID`                                               | *string*                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPortalTeamResponse](../../models/operations/getportalteamresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePortalTeam

Updates an individual developer team for a portal.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-portal-team" method="patch" path="/v3/portals/{portalId}/teams/{teamId}" -->
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

    res, err := s.PortalTeams.UpdatePortalTeam(ctx, operations.UpdatePortalTeamRequest{
        TeamID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        PortalUpdateTeamRequest: &components.PortalUpdateTeamRequest{
            Name: sdkkonnectgo.Pointer("IDM - Developers"),
            Description: sdkkonnectgo.Pointer("The Identity Management (IDM) API team."),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalTeamResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdatePortalTeamRequest](../../models/operations/updateportalteamrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.UpdatePortalTeamResponse](../../models/operations/updateportalteamresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeletePortalTeam

Deletes a developer team from a portal. Deleting a team also deletes its assigned roles. Members of the team are not deleted, but they will lose any access provided through the team.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-portal-team" method="delete" path="/v3/portals/{portalId}/teams/{teamId}" -->
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

    res, err := s.PortalTeams.DeletePortalTeam(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", "f32d905a-ed33-46a3-a093-d8f536af9a8a")
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
| `teamID`                                                 | *string*                                                 | :heavy_check_mark:                                       | ID of the team.                                          | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `portalID`                                               | *string*                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeletePortalTeamResponse](../../models/operations/deleteportalteamresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |