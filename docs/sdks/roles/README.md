# Roles

## Overview

### Available Operations

* [GetPredefinedRoles](#getpredefinedroles) - Get Predefined Roles
* [ListTeamRoles](#listteamroles) - List Team Roles
* [TeamsAssignRole](#teamsassignrole) - Assign Team Role
* [TeamsRemoveRole](#teamsremoverole) - Remove Team Role
* [ListUserRoles](#listuserroles) - List User Roles
* [UsersAssignRole](#usersassignrole) - Assign Role
* [UsersRemoveRole](#usersremoverole) - Remove Role

## GetPredefinedRoles

Retrieves the predefined, or system managed, roles.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-predefined-roles" method="get" path="/v3/roles" example="Predefined Roles" -->
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

    res, err := s.Roles.GetPredefinedRoles(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Roles != nil {
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

**[*operations.GetPredefinedRolesResponse](../../models/operations/getpredefinedrolesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListTeamRoles

Lists the roles belonging to a team. Returns 400 if any filter parameters are invalid.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-team-roles" method="get" path="/v3/teams/{teamId}/assigned-roles" example="CollectionExample" -->
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

    res, err := s.Roles.ListTeamRoles(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRoleCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           | Example                                                                                               |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                 | :heavy_check_mark:                                                                                    | The context to use for the request.                                                                   |                                                                                                       |
| `teamID`                                                                                              | *string*                                                                                              | :heavy_check_mark:                                                                                    | The team ID                                                                                           | e81bc3e5-e9db-4764-b7dd-e81e39072cbe                                                                  |
| `filter`                                                                                              | [*operations.ListTeamRolesQueryParamFilter](../../models/operations/listteamrolesqueryparamfilter.md) | :heavy_minus_sign:                                                                                    | Filter roles returned in the response.                                                                |                                                                                                       |
| `opts`                                                                                                | [][operations.Option](../../models/operations/option.md)                                              | :heavy_minus_sign:                                                                                    | The options for this request.                                                                         |                                                                                                       |

### Response

**[*operations.ListTeamRolesResponse](../../models/operations/listteamrolesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## TeamsAssignRole

Assigns a role to a team. Returns 409 if role is already assigned.

### Example Usage: Assigned Role Example

<!-- UsageSnippet language="go" operationID="teams-assign-role" method="post" path="/v3/teams/{teamId}/assigned-roles" example="Assigned Role Example" -->
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

    res, err := s.Roles.TeamsAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Assigned Roles Global

<!-- UsageSnippet language="go" operationID="teams-assign-role" method="post" path="/v3/teams/{teamId}/assigned-roles" example="Assigned Roles Global" -->
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

    res, err := s.Roles.TeamsAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameAdmin.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("18ee2573-dec0-4b83-be99-fa7700bcdc61"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionWildcard.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Assigned Roles US

<!-- UsageSnippet language="go" operationID="teams-assign-role" method="post" path="/v3/teams/{teamId}/assigned-roles" example="Assigned Roles US" -->
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

    res, err := s.Roles.TeamsAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("18ee2573-dec0-4b83-be99-fa7700bcdc61"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionUs.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Authentication Settings cannot be all Disabled

<!-- UsageSnippet language="go" operationID="teams-assign-role" method="post" path="/v3/teams/{teamId}/assigned-roles" example="Authentication Settings cannot be all Disabled" -->
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

    res, err := s.Roles.TeamsAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Cannot be Blank

<!-- UsageSnippet language="go" operationID="teams-assign-role" method="post" path="/v3/teams/{teamId}/assigned-roles" example="Cannot be Blank" -->
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

    res, err := s.Roles.TeamsAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: IdP configuration is required

<!-- UsageSnippet language="go" operationID="teams-assign-role" method="post" path="/v3/teams/{teamId}/assigned-roles" example="IdP configuration is required" -->
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

    res, err := s.Roles.TeamsAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Invalid ID format

<!-- UsageSnippet language="go" operationID="teams-assign-role" method="post" path="/v3/teams/{teamId}/assigned-roles" example="Invalid ID format" -->
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

    res, err := s.Roles.TeamsAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Invalid Permissions

<!-- UsageSnippet language="go" operationID="teams-assign-role" method="post" path="/v3/teams/{teamId}/assigned-roles" example="Invalid Permissions" -->
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

    res, err := s.Roles.TeamsAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Must be a valid UUID v4

<!-- UsageSnippet language="go" operationID="teams-assign-role" method="post" path="/v3/teams/{teamId}/assigned-roles" example="Must be a valid UUID v4" -->
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

    res, err := s.Roles.TeamsAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Not Found

<!-- UsageSnippet language="go" operationID="teams-assign-role" method="post" path="/v3/teams/{teamId}/assigned-roles" example="Not Found" -->
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

    res, err := s.Roles.TeamsAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: OIDC needs an IdP configuration

<!-- UsageSnippet language="go" operationID="teams-assign-role" method="post" path="/v3/teams/{teamId}/assigned-roles" example="OIDC needs an IdP configuration" -->
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

    res, err := s.Roles.TeamsAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Password Complexity

<!-- UsageSnippet language="go" operationID="teams-assign-role" method="post" path="/v3/teams/{teamId}/assigned-roles" example="Password Complexity" -->
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

    res, err := s.Roles.TeamsAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Permission Denied

<!-- UsageSnippet language="go" operationID="teams-assign-role" method="post" path="/v3/teams/{teamId}/assigned-roles" example="Permission Denied" -->
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

    res, err := s.Roles.TeamsAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Request Format is Invalid

<!-- UsageSnippet language="go" operationID="teams-assign-role" method="post" path="/v3/teams/{teamId}/assigned-roles" example="Request Format is Invalid" -->
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

    res, err := s.Roles.TeamsAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: System teams cannot be modified

<!-- UsageSnippet language="go" operationID="teams-assign-role" method="post" path="/v3/teams/{teamId}/assigned-roles" example="System teams cannot be modified" -->
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

    res, err := s.Roles.TeamsAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="teams-assign-role" method="post" path="/v3/teams/{teamId}/assigned-roles" example="Unauthorized" -->
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

    res, err := s.Roles.TeamsAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Unsupported filter operation

<!-- UsageSnippet language="go" operationID="teams-assign-role" method="post" path="/v3/teams/{teamId}/assigned-roles" example="Unsupported filter operation" -->
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

    res, err := s.Roles.TeamsAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: User is Already Active

<!-- UsageSnippet language="go" operationID="teams-assign-role" method="post" path="/v3/teams/{teamId}/assigned-roles" example="User is Already Active" -->
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

    res, err := s.Roles.TeamsAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Username or Password invalid

<!-- UsageSnippet language="go" operationID="teams-assign-role" method="post" path="/v3/teams/{teamId}/assigned-roles" example="Username or Password invalid" -->
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

    res, err := s.Roles.TeamsAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                       | Type                                                            | Required                                                        | Description                                                     | Example                                                         |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `ctx`                                                           | [context.Context](https://pkg.go.dev/context#Context)           | :heavy_check_mark:                                              | The context to use for the request.                             |                                                                 |
| `teamID`                                                        | *string*                                                        | :heavy_check_mark:                                              | The team ID                                                     | e81bc3e5-e9db-4764-b7dd-e81e39072cbe                            |
| `assignRole`                                                    | [*components.AssignRole](../../models/components/assignrole.md) | :heavy_minus_sign:                                              | The request schema for assigning a role.                        |                                                                 |
| `opts`                                                          | [][operations.Option](../../models/operations/option.md)        | :heavy_minus_sign:                                              | The options for this request.                                   |                                                                 |

### Response

**[*operations.TeamsAssignRoleResponse](../../models/operations/teamsassignroleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## TeamsRemoveRole

Removes an assigned role from a team. Returns 404 if the requested team or assigned role were not found.

### Example Usage

<!-- UsageSnippet language="go" operationID="teams-remove-role" method="delete" path="/v3/teams/{teamId}/assigned-roles/{roleId}" -->
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

    res, err := s.Roles.TeamsRemoveRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", "8350205f-a305-4e39-abe9-bc082a80091a")
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
| `teamID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The team ID.                                             | e81bc3e5-e9db-4764-b7dd-e81e39072cbe                     |
| `roleID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The role ID.                                             | 8350205f-a305-4e39-abe9-bc082a80091a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.TeamsRemoveRoleResponse](../../models/operations/teamsremoveroleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListUserRoles

Lists the roles assigned to a user.  Returns 400 if any filter parameters are invalid.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-user-roles" method="get" path="/v3/users/{userId}/assigned-roles" example="CollectionExample" -->
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

    res, err := s.Roles.ListUserRoles(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRoleCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           | Example                                                                                               |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                 | :heavy_check_mark:                                                                                    | The context to use for the request.                                                                   |                                                                                                       |
| `userID`                                                                                              | *string*                                                                                              | :heavy_check_mark:                                                                                    | The user ID                                                                                           | e81bc3e5-e9db-4764-b7dd-e81e39072cbe                                                                  |
| `filter`                                                                                              | [*operations.ListUserRolesQueryParamFilter](../../models/operations/listuserrolesqueryparamfilter.md) | :heavy_minus_sign:                                                                                    | Filter roles returned in the response.                                                                |                                                                                                       |
| `opts`                                                                                                | [][operations.Option](../../models/operations/option.md)                                              | :heavy_minus_sign:                                                                                    | The options for this request.                                                                         |                                                                                                       |

### Response

**[*operations.ListUserRolesResponse](../../models/operations/listuserrolesresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## UsersAssignRole

Assigns a role to a user. Returns 409 if role is already assigned.

### Example Usage: Assigned Role Example

<!-- UsageSnippet language="go" operationID="users-assign-role" method="post" path="/v3/users/{userId}/assigned-roles" example="Assigned Role Example" -->
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

    res, err := s.Roles.UsersAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Assigned Roles Global

<!-- UsageSnippet language="go" operationID="users-assign-role" method="post" path="/v3/users/{userId}/assigned-roles" example="Assigned Roles Global" -->
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

    res, err := s.Roles.UsersAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameAdmin.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("18ee2573-dec0-4b83-be99-fa7700bcdc61"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionWildcard.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Assigned Roles US

<!-- UsageSnippet language="go" operationID="users-assign-role" method="post" path="/v3/users/{userId}/assigned-roles" example="Assigned Roles US" -->
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

    res, err := s.Roles.UsersAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("18ee2573-dec0-4b83-be99-fa7700bcdc61"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionUs.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Authentication Settings cannot be all Disabled

<!-- UsageSnippet language="go" operationID="users-assign-role" method="post" path="/v3/users/{userId}/assigned-roles" example="Authentication Settings cannot be all Disabled" -->
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

    res, err := s.Roles.UsersAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Cannot be Blank

<!-- UsageSnippet language="go" operationID="users-assign-role" method="post" path="/v3/users/{userId}/assigned-roles" example="Cannot be Blank" -->
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

    res, err := s.Roles.UsersAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: IdP configuration is required

<!-- UsageSnippet language="go" operationID="users-assign-role" method="post" path="/v3/users/{userId}/assigned-roles" example="IdP configuration is required" -->
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

    res, err := s.Roles.UsersAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Invalid ID format

<!-- UsageSnippet language="go" operationID="users-assign-role" method="post" path="/v3/users/{userId}/assigned-roles" example="Invalid ID format" -->
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

    res, err := s.Roles.UsersAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Invalid Permissions

<!-- UsageSnippet language="go" operationID="users-assign-role" method="post" path="/v3/users/{userId}/assigned-roles" example="Invalid Permissions" -->
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

    res, err := s.Roles.UsersAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Must be a valid UUID v4

<!-- UsageSnippet language="go" operationID="users-assign-role" method="post" path="/v3/users/{userId}/assigned-roles" example="Must be a valid UUID v4" -->
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

    res, err := s.Roles.UsersAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Not Found

<!-- UsageSnippet language="go" operationID="users-assign-role" method="post" path="/v3/users/{userId}/assigned-roles" example="Not Found" -->
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

    res, err := s.Roles.UsersAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: OIDC needs an IdP configuration

<!-- UsageSnippet language="go" operationID="users-assign-role" method="post" path="/v3/users/{userId}/assigned-roles" example="OIDC needs an IdP configuration" -->
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

    res, err := s.Roles.UsersAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Password Complexity

<!-- UsageSnippet language="go" operationID="users-assign-role" method="post" path="/v3/users/{userId}/assigned-roles" example="Password Complexity" -->
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

    res, err := s.Roles.UsersAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Permission Denied

<!-- UsageSnippet language="go" operationID="users-assign-role" method="post" path="/v3/users/{userId}/assigned-roles" example="Permission Denied" -->
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

    res, err := s.Roles.UsersAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Request Format is Invalid

<!-- UsageSnippet language="go" operationID="users-assign-role" method="post" path="/v3/users/{userId}/assigned-roles" example="Request Format is Invalid" -->
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

    res, err := s.Roles.UsersAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: System teams cannot be modified

<!-- UsageSnippet language="go" operationID="users-assign-role" method="post" path="/v3/users/{userId}/assigned-roles" example="System teams cannot be modified" -->
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

    res, err := s.Roles.UsersAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="users-assign-role" method="post" path="/v3/users/{userId}/assigned-roles" example="Unauthorized" -->
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

    res, err := s.Roles.UsersAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Unsupported filter operation

<!-- UsageSnippet language="go" operationID="users-assign-role" method="post" path="/v3/users/{userId}/assigned-roles" example="Unsupported filter operation" -->
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

    res, err := s.Roles.UsersAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: User is Already Active

<!-- UsageSnippet language="go" operationID="users-assign-role" method="post" path="/v3/users/{userId}/assigned-roles" example="User is Already Active" -->
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

    res, err := s.Roles.UsersAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```
### Example Usage: Username or Password invalid

<!-- UsageSnippet language="go" operationID="users-assign-role" method="post" path="/v3/users/{userId}/assigned-roles" example="Username or Password invalid" -->
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

    res, err := s.Roles.UsersAssignRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.Pointer("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.AssignRoleEntityRegionEu.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRole != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                       | Type                                                            | Required                                                        | Description                                                     | Example                                                         |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `ctx`                                                           | [context.Context](https://pkg.go.dev/context#Context)           | :heavy_check_mark:                                              | The context to use for the request.                             |                                                                 |
| `userID`                                                        | *string*                                                        | :heavy_check_mark:                                              | The user ID                                                     | e81bc3e5-e9db-4764-b7dd-e81e39072cbe                            |
| `assignRole`                                                    | [*components.AssignRole](../../models/components/assignrole.md) | :heavy_minus_sign:                                              | The request schema for assigning a role.                        |                                                                 |
| `opts`                                                          | [][operations.Option](../../models/operations/option.md)        | :heavy_minus_sign:                                              | The options for this request.                                   |                                                                 |

### Response

**[*operations.UsersAssignRoleResponse](../../models/operations/usersassignroleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UsersRemoveRole

Removes an assigned role from a user. Returns 404 if the requested user or assigned role were not found.

### Example Usage

<!-- UsageSnippet language="go" operationID="users-remove-role" method="delete" path="/v3/users/{userId}/assigned-roles/{roleId}" -->
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

    res, err := s.Roles.UsersRemoveRole(ctx, "e81bc3e5-e9db-4764-b7dd-e81e39072cbe", "8350205f-a305-4e39-abe9-bc082a80091a")
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
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | ID of the user.                                          | e81bc3e5-e9db-4764-b7dd-e81e39072cbe                     |
| `roleID`                                                 | *string*                                                 | :heavy_check_mark:                                       | ID of the role.                                          | 8350205f-a305-4e39-abe9-bc082a80091a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.UsersRemoveRoleResponse](../../models/operations/usersremoveroleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |