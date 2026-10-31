# PrincipalGroup

## Overview

Groups organize principals, or other groups, within a directory for authorization purposes.


### Available Operations

* [CreateGroup](#creategroup) - Create a Group in a Directory
* [ListGroups](#listgroups) - List all the Groups of a Directory
* [GetGroup](#getgroup) - Get a Group
* [ReplaceGroup](#replacegroup) - Replace a Group
* [DeleteGroup](#deletegroup) - Delete a Group
* [CreateGroupKey](#creategroupkey) - Create a group API key
* [ListGroupKeys](#listgroupkeys) - List all API keys for a group
* [UpdateGroupKey](#updategroupkey) - Update a group API key
* [GetGroupKey](#getgroupkey) - Get a group API key
* [DeleteGroupKey](#deletegroupkey) - Delete an API key
* [ListGroupMembership](#listgroupmembership) - List a Group's membership
* [CreateGroupMembershipEntry](#creategroupmembershipentry) - Add a member to a Group
* [GetGroupMembershipEntry](#getgroupmembershipentry) - Get a Group membership entry
* [DeleteGroupMembershipEntry](#deletegroupmembershipentry) - Remove a member from a Group
* [BatchCreateGroupMembershipEntries](#batchcreategroupmembershipentries) - Add multiple members to a Group

## CreateGroup

Creates a new group within the specified directory. Groups organize principals, or other groups,
for authorization purposes.

`name` must be unique within the directory; creating a group
with a `name` already in use returns 409.


### Example Usage

<!-- UsageSnippet language="go" operationID="createGroup" method="post" path="/v2/directories/{directoryId}/groups" -->
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

    res, err := s.PrincipalGroup.CreateGroup(ctx, "1fae497b-f627-4ef0-b7fd-49eb43c4a7af", components.CreateGroupBody{
        Name: "<value>",
        Labels: map[string]string{
            "env": "test",
        },
        ManagedBy: map[string]string{
            "owner": "terraform",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `directoryID`                                                            | `string`                                                                 | :heavy_check_mark:                                                       | ID of the directory                                                      |
| `createGroupBody`                                                        | [components.CreateGroupBody](../../models/components/creategroupbody.md) | :heavy_check_mark:                                                       | Group to be created                                                      |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.CreateGroupResponse](../../models/operations/creategroupresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListGroups

Retrieves a paginated list of all groups within the specified directory.
Results can be sorted by name, created_at, or updated_at, and filtered by name, description, id, labels, or type.


### Example Usage

<!-- UsageSnippet language="go" operationID="listGroups" method="get" path="/v2/directories/{directoryId}/groups" -->
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

    res, err := s.PrincipalGroup.ListGroups(ctx, operations.ListGroupsRequest{
        DirectoryID: "a291d21b-6ee9-4e96-b330-459bb024840d",
        Page: &components.CursorPageParameters{
            Size: sdkkonnectgo.Pointer[int64](10),
            After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
            Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListKongGroups != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListGroupsRequest](../../models/operations/listgroupsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.ListGroupsResponse](../../models/operations/listgroupsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetGroup

Retrieves detailed information about a specific group by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getGroup" method="get" path="/v2/directories/{directoryId}/groups/{groupId}" -->
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

    res, err := s.PrincipalGroup.GetGroup(ctx, "5fb844cc-afa5-41ff-aaca-820029554124", "25a6da92-84da-492f-95e8-e70d01b3b4fa")
    if err != nil {
        log.Fatal(err)
    }
    if res.KongGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `directoryID`                                            | `string`                                                 | :heavy_check_mark:                                       | ID of the directory                                      |
| `groupID`                                                | `string`                                                 | :heavy_check_mark:                                       | Unique identifier of the group (UUID format)             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetGroupResponse](../../models/operations/getgroupresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ReplaceGroup

Completely replaces an existing group with the provided data.
All fields must be specified; any omitted fields will be reset to their defaults.

`name` must be unique within the directory; replacing a group
with a `name` already used by another group returns 409.


### Example Usage

<!-- UsageSnippet language="go" operationID="replaceGroup" method="put" path="/v2/directories/{directoryId}/groups/{groupId}" -->
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

    res, err := s.PrincipalGroup.ReplaceGroup(ctx, operations.ReplaceGroupRequest{
        DirectoryID: "49b47b20-13bb-49ee-8d3e-dfdf7ab5ec76",
        GroupID: "922607ce-e795-45de-8b94-a1d6fd667dbe",
        ReplaceGroupBody: components.ReplaceGroupBody{
            Name: "<value>",
            Labels: map[string]string{
                "env": "test",
            },
            ManagedBy: map[string]string{
                "owner": "terraform",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ReplaceGroupRequest](../../models/operations/replacegrouprequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ReplaceGroupResponse](../../models/operations/replacegroupresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteGroup

Permanently deletes a group from the directory.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteGroup" method="delete" path="/v2/directories/{directoryId}/groups/{groupId}" -->
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

    res, err := s.PrincipalGroup.DeleteGroup(ctx, "c2c7eec9-6fc6-4056-ad97-e9e567996cc5", "d8527e29-40fb-498e-8f77-490604bdc214")
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
| `directoryID`                                            | `string`                                                 | :heavy_check_mark:                                       | ID of the directory                                      |
| `groupID`                                                | `string`                                                 | :heavy_check_mark:                                       | Unique identifier of the group (UUID format)             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteGroupResponse](../../models/operations/deletegroupresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateGroupKey

Creates a new API key for the specified group.
Keys can be system-generated (v1) or imported with a custom secret, and may have an optional expiration date.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-group-key" method="post" path="/v2/directories/{directoryId}/groups/{groupId}/api-keys" -->
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

    res, err := s.PrincipalGroup.CreateGroupKey(ctx, operations.CreateGroupKeyRequest{
        DirectoryID: "f6cd2cb2-efde-4ac3-801f-c24253cac0df",
        GroupID: "8fe53b37-b65e-4ab7-b8f7-d71fafc853d8",
        CreatePrincipalKeyBody: components.CreateCreatePrincipalKeyBodyV1(
            components.CreatePrincipalKeyBodyv1{
                Type: components.CreatePrincipalKeyBodyv1TypeV1,
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongPrincipalKeyWithSecret != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateGroupKeyRequest](../../models/operations/creategroupkeyrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateGroupKeyResponse](../../models/operations/creategroupkeyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListGroupKeys

Retrieves a paginated list of all API keys associated with the specified group.
Results can be sorted by created_at, updated_at, or name.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-group-keys" method="get" path="/v2/directories/{directoryId}/groups/{groupId}/api-keys" -->
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

    res, err := s.PrincipalGroup.ListGroupKeys(ctx, operations.ListGroupKeysRequest{
        DirectoryID: "acdccb2e-19ec-4eb8-827f-c82fabe48d22",
        GroupID: "24e07a14-84dc-4955-863c-e87f1800cc9a",
        Page: &components.CursorPageParameters{
            Size: sdkkonnectgo.Pointer[int64](10),
            After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
            Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListKongPrincipalKeys != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListGroupKeysRequest](../../models/operations/listgroupkeysrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ListGroupKeysResponse](../../models/operations/listgroupkeysresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateGroupKey

Partially updates an existing API key's metadata.
Note: The key secret cannot be modified after creation.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-group-key" method="patch" path="/v2/directories/{directoryId}/groups/{groupId}/api-keys/{apiKeyId}" -->
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

    res, err := s.PrincipalGroup.UpdateGroupKey(ctx, operations.UpdateGroupKeyRequest{
        DirectoryID: "3252ac7c-01bd-49f4-8e40-024b357f0aa2",
        GroupID: "dfe79694-0010-43f0-bff6-6fc3fe8581c3",
        APIKeyID: "ef777c9a-e52d-4c42-95d9-a8833d75cfdd",
        UpdatePrincipalKeyBody: components.UpdatePrincipalKeyBody{
            ManagedBy: map[string]*string{
                "owner": sdkkonnectgo.Pointer("terraform"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongPrincipalKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateGroupKeyRequest](../../models/operations/updategroupkeyrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateGroupKeyResponse](../../models/operations/updategroupkeyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetGroupKey

Retrieves detailed information about a specific API key.
Note: The key secret is not returned; it is only available when the key is first created.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-group-key" method="get" path="/v2/directories/{directoryId}/groups/{groupId}/api-keys/{apiKeyId}" -->
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

    res, err := s.PrincipalGroup.GetGroupKey(ctx, operations.GetGroupKeyRequest{
        DirectoryID: "0c858fc6-395c-467c-9786-5c6df5a6c88b",
        GroupID: "61263f0b-3377-4a9c-83e9-9d6afa5b07ef",
        APIKeyID: "8435da50-4bbd-4d6a-ac7f-86533e84277f",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongPrincipalKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetGroupKeyRequest](../../models/operations/getgroupkeyrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.GetGroupKeyResponse](../../models/operations/getgroupkeyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteGroupKey

Permanently deletes an API key, immediately revoking its access.
This operation cannot be undone.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-group-key" method="delete" path="/v2/directories/{directoryId}/groups/{groupId}/api-keys/{apiKeyId}" -->
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

    res, err := s.PrincipalGroup.DeleteGroupKey(ctx, operations.DeleteGroupKeyRequest{
        DirectoryID: "54501208-2c78-4bf1-96b0-62ee61025afc",
        GroupID: "02f9c7c3-e47f-4df5-a228-bf1f8a761892",
        APIKeyID: "aaafbdbe-d602-4fcb-96b8-0c9b769d5710",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteGroupKeyRequest](../../models/operations/deletegroupkeyrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.DeleteGroupKeyResponse](../../models/operations/deletegroupkeyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListGroupMembership

Retrieves a paginated list of members (principals or groups) belonging to a group.
Results can be filtered by membership entry id or by the member's display name.


### Example Usage

<!-- UsageSnippet language="go" operationID="listGroupMembership" method="get" path="/v2/directories/{directoryId}/groups/{groupId}/membership" -->
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

    res, err := s.PrincipalGroup.ListGroupMembership(ctx, operations.ListGroupMembershipRequest{
        DirectoryID: "01ebc09a-23b7-4795-bdd3-9e0ec59b3089",
        GroupID: "af0e7099-9d3c-461f-8124-f7e674625ad3",
        Page: &components.CursorPageParameters{
            Size: sdkkonnectgo.Pointer[int64](10),
            After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
            Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListKongGroupMembership != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListGroupMembershipRequest](../../models/operations/listgroupmembershiprequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListGroupMembershipResponse](../../models/operations/listgroupmembershipresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateGroupMembershipEntry

Adds a single member, either a principal or another group, to a group's membership.


### Example Usage

<!-- UsageSnippet language="go" operationID="createGroupMembershipEntry" method="post" path="/v2/directories/{directoryId}/groups/{groupId}/membership" -->
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

    res, err := s.PrincipalGroup.CreateGroupMembershipEntry(ctx, operations.CreateGroupMembershipEntryRequest{
        DirectoryID: "abdaa2cb-9f3e-46ba-a3b6-89a1b8f72bf5",
        GroupID: "075ca53b-b4a0-4bef-b40c-8ff0db8cbcef",
        CreateGroupMembershipEntryBody: components.CreateGroupMembershipEntryBody{
            Member: components.KongGroupMemberInput{
                ID: "15bd0e5c-2a2d-4cdb-bd62-364d8dca780b",
                Type: components.KongGroupMembershipTypeGroup,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongGroupMembershipEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.CreateGroupMembershipEntryRequest](../../models/operations/creategroupmembershipentryrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.CreateGroupMembershipEntryResponse](../../models/operations/creategroupmembershipentryresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetGroupMembershipEntry

Retrieves a specific group membership entry by the member's ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getGroupMembershipEntry" method="get" path="/v2/directories/{directoryId}/groups/{groupId}/membership/{groupMembershipId}" -->
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

    res, err := s.PrincipalGroup.GetGroupMembershipEntry(ctx, operations.GetGroupMembershipEntryRequest{
        DirectoryID: "79d363ed-8811-4c51-a0f2-653d940aefad",
        GroupID: "598187e3-9e72-4159-963e-fa44d3c9a936",
        GroupMembershipID: "9716bc2e-781f-442a-b0e6-5d601345bc8a",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongGroupMembershipEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetGroupMembershipEntryRequest](../../models/operations/getgroupmembershipentryrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.GetGroupMembershipEntryResponse](../../models/operations/getgroupmembershipentryresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteGroupMembershipEntry

Removes a single member, either a principal or another group, from a group's membership.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteGroupMembershipEntry" method="delete" path="/v2/directories/{directoryId}/groups/{groupId}/membership/{groupMembershipId}" -->
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

    res, err := s.PrincipalGroup.DeleteGroupMembershipEntry(ctx, operations.DeleteGroupMembershipEntryRequest{
        DirectoryID: "83ad75ff-0914-4bab-b75d-d025007fcd5e",
        GroupID: "9eddd515-0199-4661-af12-e7d8e98feec4",
        GroupMembershipID: "59d2ae60-c871-41fb-a4d1-1df8921aabfb",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.DeleteGroupMembershipEntryRequest](../../models/operations/deletegroupmembershipentryrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.DeleteGroupMembershipEntryResponse](../../models/operations/deletegroupmembershipentryresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## BatchCreateGroupMembershipEntries

Adds multiple members (principals or groups) to a group's membership in a single request.
Either the entire batch is added or none of it is; body validation failures are reported per item.


### Example Usage

<!-- UsageSnippet language="go" operationID="batchCreateGroupMembershipEntries" method="post" path="/v2/directories/{directoryId}/groups/{groupId}/membership/batch" -->
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

    res, err := s.PrincipalGroup.BatchCreateGroupMembershipEntries(ctx, operations.BatchCreateGroupMembershipEntriesRequest{
        DirectoryID: "cab7822c-f085-40b1-afd8-5d11010c2ef7",
        GroupID: "ac829613-3a6b-478f-be6f-7d6044e3700d",
        BatchCreateGroupMembershipEntriesBody: components.BatchCreateGroupMembershipEntriesBody{
            Data: []components.CreateGroupMembershipEntryBody{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BatchKongGroupMembership != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.BatchCreateGroupMembershipEntriesRequest](../../models/operations/batchcreategroupmembershipentriesrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.BatchCreateGroupMembershipEntriesResponse](../../models/operations/batchcreategroupmembershipentriesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |