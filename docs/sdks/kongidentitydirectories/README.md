# KongIdentityDirectories

## Overview

Use a directory to group consumers around an identity, defined by organizational boundaries,
such as a production directory or a development directory. Directories are connected to a
[geographic region](https://docs.konghq.com/konnect/geo/) in Konnect. Centrally managed consumers defined in
directories can be used across multiple control planes.


### Available Operations

* [CreateDirectory](#createdirectory) - Create a new Directory
* [ListKongDirectories](#listkongdirectories) - List all directories
* [UpdateDirectory](#updatedirectory) - Update a Directory
* [ReplaceDirectory](#replacedirectory) - Replace a Directory
* [GetDirectory](#getdirectory) - Get a Directory
* [DeleteDirectory](#deletedirectory) - Delete a Directory

## CreateDirectory

Creates a new directory for grouping principals around an organizational identity boundary.
Directories can be configured with metadata, labels, and control plane access restrictions.


### Example Usage

<!-- UsageSnippet language="go" operationID="createDirectory" method="post" path="/v2/directories" -->
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

    res, err := s.KongIdentityDirectories.CreateDirectory(ctx, components.CreateDirectoryBody{
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
    if res.KongDirectory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.CreateDirectoryBody](../../models/components/createdirectorybody.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.CreateDirectoryResponse](../../models/operations/createdirectoryresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListKongDirectories

Retrieves a paginated list of all available directories in the organization.
Results can be sorted by name, created_at, or updated_at.


### Example Usage

<!-- UsageSnippet language="go" operationID="ListKongDirectories" method="get" path="/v2/directories" -->
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

    res, err := s.KongIdentityDirectories.ListKongDirectories(ctx, &components.CursorPageParameters{
        Size: sdkkonnectgo.Pointer[int64](10),
        After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    }, sdkkonnectgo.Pointer("created_at desc"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListKongDirectories != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                           | Type                                                                                                | Required                                                                                            | Description                                                                                         | Example                                                                                             |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                               | :heavy_check_mark:                                                                                  | The context to use for the request.                                                                 |                                                                                                     |
| `page`                                                                                              | [*components.CursorPageParameters](../../models/components/cursorpageparameters.md)                 | :heavy_minus_sign:                                                                                  | Determines which page of the collection to retrieve.                                                |                                                                                                     |
| `sort`                                                                                              | `*string`                                                                                           | :heavy_minus_sign:                                                                                  | Sorts a collection of directories. Supported sort attributes are:<br/>- name<br/>- created_at<br/>- updated_at<br/> | created_at desc                                                                                     |
| `opts`                                                                                              | [][operations.Option](../../models/operations/option.md)                                            | :heavy_minus_sign:                                                                                  | The options for this request.                                                                       |                                                                                                     |

### Response

**[*operations.ListKongDirectoriesResponse](../../models/operations/listkongdirectoriesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateDirectory

Partially updates an existing directory by modifying only the provided fields.
Unspecified fields will retain their current values.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateDirectory" method="patch" path="/v2/directories/{directoryId}" -->
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

    res, err := s.KongIdentityDirectories.UpdateDirectory(ctx, "a02c929f-dd3e-4154-9012-f54a9eb3a053", components.UpdateDirectoryBody{
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongDirectory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `directoryID`                                                                    | `string`                                                                         | :heavy_check_mark:                                                               | ID of the directory                                                              |
| `updateDirectoryBody`                                                            | [components.UpdateDirectoryBody](../../models/components/updatedirectorybody.md) | :heavy_check_mark:                                                               | Request body for partially updating a directory                                  |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.UpdateDirectoryResponse](../../models/operations/updatedirectoryresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ReplaceDirectory

Completely replaces an existing directory with the provided data.
All fields must be specified; any omitted fields will be reset to their defaults.


### Example Usage

<!-- UsageSnippet language="go" operationID="replaceDirectory" method="put" path="/v2/directories/{directoryId}" -->
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

    res, err := s.KongIdentityDirectories.ReplaceDirectory(ctx, "2c33078f-16f4-4bc9-8eb3-ec133183f8ed", components.ReplaceDirectoryBody{
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
    if res.KongDirectory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `directoryID`                                                                      | `string`                                                                           | :heavy_check_mark:                                                                 | ID of the directory                                                                |
| `replaceDirectoryBody`                                                             | [components.ReplaceDirectoryBody](../../models/components/replacedirectorybody.md) | :heavy_check_mark:                                                                 | Directory to be replaced                                                           |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ReplaceDirectoryResponse](../../models/operations/replacedirectoryresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetDirectory

Retrieves detailed information about a specific directory by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDirectory" method="get" path="/v2/directories/{directoryId}" -->
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

    res, err := s.KongIdentityDirectories.GetDirectory(ctx, "803f01bb-6279-4d04-a298-a1e2a9e4d1b6")
    if err != nil {
        log.Fatal(err)
    }
    if res.KongDirectory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `directoryID`                                            | `string`                                                 | :heavy_check_mark:                                       | ID of the directory                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetDirectoryResponse](../../models/operations/getdirectoryresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteDirectory

Deletes a directory and optionally all its associated resources (principals, keys, etc.).
Use the force parameter to control deletion behavior when dependencies exist.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteDirectory" method="delete" path="/v2/directories/{directoryId}" -->
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

    res, err := s.KongIdentityDirectories.DeleteDirectory(ctx, "12c13150-c0a9-42a1-b2a2-75521cc1c058", operations.DeleteDirectoryForceFalse.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `directoryID`                                                                                                                                                                | `string`                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                           | ID of the directory                                                                                                                                                          |
| `forceDestroy`                                                                                                                                                               | [*operations.DeleteDirectoryForce](../../models/operations/deletedirectoryforce.md)                                                                                          | :heavy_minus_sign:                                                                                                                                                           | If true, delete the specified directory and all its associated resources. If false, only allow deletion if no clients, scopes or claims are associated with the auth server. |
| `opts`                                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                                     | :heavy_minus_sign:                                                                                                                                                           | The options for this request.                                                                                                                                                |

### Response

**[*operations.DeleteDirectoryResponse](../../models/operations/deletedirectoryresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |