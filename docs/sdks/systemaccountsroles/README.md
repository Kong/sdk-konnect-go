# SystemAccountsRoles
(*SystemAccountsRoles*)

## Overview

### Available Operations

* [GetSystemAccountsAccountIDAssignedRoles](#getsystemaccountsaccountidassignedroles) - Fetch Assigned Roles for System Account
* [PostSystemAccountsAccountIDAssignedRoles](#postsystemaccountsaccountidassignedroles) - Create Assigned Role for System Account
* [DeleteSystemAccountsAccountIDAssignedRolesRoleID](#deletesystemaccountsaccountidassignedrolesroleid) - Delete Assigned Role from System Account

## GetSystemAccountsAccountIDAssignedRoles

Lists the roles belonging to a system account. Returns 400 if any filter parameters are invalid.

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

    res, err := s.SystemAccountsRoles.GetSystemAccountsAccountIDAssignedRoles(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignedRoleCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                 | Type                                                                                                                                                      | Required                                                                                                                                                  | Description                                                                                                                                               |
| --------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                                                     | :heavy_check_mark:                                                                                                                                        | The context to use for the request.                                                                                                                       |
| `accountID`                                                                                                                                               | *string*                                                                                                                                                  | :heavy_check_mark:                                                                                                                                        | ID of the system account.                                                                                                                                 |
| `filter`                                                                                                                                                  | [*operations.GetSystemAccountsAccountIDAssignedRolesQueryParamFilter](../../models/operations/getsystemaccountsaccountidassignedrolesqueryparamfilter.md) | :heavy_minus_sign:                                                                                                                                        | Filter roles returned in the response.                                                                                                                    |
| `opts`                                                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                                                  | :heavy_minus_sign:                                                                                                                                        | The options for this request.                                                                                                                             |

### Response

**[*operations.GetSystemAccountsAccountIDAssignedRolesResponse](../../models/operations/getsystemaccountsaccountidassignedrolesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PostSystemAccountsAccountIDAssignedRoles

Assigns a role to a system account. Returns 409 if role is already assigned.

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

    res, err := s.SystemAccountsRoles.PostSystemAccountsAccountIDAssignedRoles(ctx, "<id>", &components.AssignRole{
        RoleName: components.RoleNameViewer.ToPointer(),
        EntityID: sdkkonnectgo.String("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: components.EntityTypeNameControlPlanes.ToPointer(),
        EntityRegion: components.EntityRegionEu.ToPointer(),
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

| Parameter                                                       | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `ctx`                                                           | [context.Context](https://pkg.go.dev/context#Context)           | :heavy_check_mark:                                              | The context to use for the request.                             |
| `accountID`                                                     | *string*                                                        | :heavy_check_mark:                                              | ID of the system account.                                       |
| `assignRole`                                                    | [*components.AssignRole](../../models/components/assignrole.md) | :heavy_minus_sign:                                              | The request schema for assigning a role.                        |
| `opts`                                                          | [][operations.Option](../../models/operations/option.md)        | :heavy_minus_sign:                                              | The options for this request.                                   |

### Response

**[*operations.PostSystemAccountsAccountIDAssignedRolesResponse](../../models/operations/postsystemaccountsaccountidassignedrolesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteSystemAccountsAccountIDAssignedRolesRoleID

Removes an assigned role from a system account. Returns 404 if the system account or assigned role were not found.

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

    res, err := s.SystemAccountsRoles.DeleteSystemAccountsAccountIDAssignedRolesRoleID(ctx, "<id>", "<id>")
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
| `roleID`                                                 | *string*                                                 | :heavy_check_mark:                                       | ID of the role.                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteSystemAccountsAccountIDAssignedRolesRoleIDResponse](../../models/operations/deletesystemaccountsaccountidassignedrolesroleidresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |