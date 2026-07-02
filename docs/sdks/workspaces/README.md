# Workspaces

## Overview

### Available Operations

* [ListWorkspaces](#listworkspaces) - List all Workspaces
* [CreateWorkspace](#createworkspace) - Create a Workspace
* [GetWorkspace](#getworkspace) - Get a Workspace
* [DeleteWorkspace](#deleteworkspace) - Delete a Workspace
* [UpsertWorkspace](#upsertworkspace) - Upsert a Workspace

## ListWorkspaces

List all Workspaces

### Example Usage

<!-- UsageSnippet language="go" operationID="list-workspaces" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/workspaces" -->
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

    res, err := s.Workspaces.ListWorkspaces(ctx, operations.ListWorkspacesRequest{
        ControlPlaneID: "a35c6ae6-4928-4940-8352-eede7b82fd6e",
        FilterNameContains: sdkkonnectgo.Pointer("test"),
        FilterNameEq: sdkkonnectgo.Pointer("test"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListWorkspacesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListWorkspacesRequest](../../models/operations/listworkspacesrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListWorkspacesResponse](../../models/operations/listworkspacesresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequestError     | 400                           | application/problem+json      |
| sdkerrors.UnauthorizedError   | 401                           | application/problem+json      |
| sdkerrors.ForbiddenError      | 403                           | application/problem+json      |
| sdkerrors.NotFoundError       | 404                           | application/problem+json      |
| sdkerrors.BaseError           | 500                           | application/problem+json      |
| sdkerrors.NotImplementedError | 501                           | application/problem+json      |
| sdkerrors.ServiceUnavailable  | 503                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## CreateWorkspace

Create a Workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="create-workspace" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/workspaces" -->
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

    res, err := s.Workspaces.CreateWorkspace(ctx, "78cd1b93-9c1c-4201-b39f-396c6c316a74", components.CreateWorkspaceRequest{
        Name: "team-1",
        Comment: sdkkonnectgo.Pointer("A test workspace for team 1"),
        Description: sdkkonnectgo.Pointer("A test workspace for team 1"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Workspace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `controlPlaneID`                                                                       | `string`                                                                               | :heavy_check_mark:                                                                     | ID of a control plane                                                                  |
| `createWorkspaceRequest`                                                               | [components.CreateWorkspaceRequest](../../models/components/createworkspacerequest.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateWorkspaceResponse](../../models/operations/createworkspaceresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequestError     | 400                           | application/problem+json      |
| sdkerrors.UnauthorizedError   | 401                           | application/problem+json      |
| sdkerrors.ForbiddenError      | 403                           | application/problem+json      |
| sdkerrors.NotFoundError       | 404                           | application/problem+json      |
| sdkerrors.BaseError           | 500                           | application/problem+json      |
| sdkerrors.NotImplementedError | 501                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## GetWorkspace

Get a Workspace using ID or name.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/workspaces/{workspaceIdOrName}" -->
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

    res, err := s.Workspaces.GetWorkspace(ctx, "6275b909-34d6-4d87-80b9-5a7f467ca626", "2747d1e5-8246-4f65-a939-b392f1ee17f8")
    if err != nil {
        log.Fatal(err)
    }
    if res.Workspace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `controlPlaneID`                                         | `string`                                                 | :heavy_check_mark:                                       | ID of a control plane                                    |                                                          |
| `workspaceIDOrName`                                      | `string`                                                 | :heavy_check_mark:                                       | ID or name of the workspace to retreive                  | 2747d1e5-8246-4f65-a939-b392f1ee17f8                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetWorkspaceResponse](../../models/operations/getworkspaceresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequestError     | 400                           | application/problem+json      |
| sdkerrors.UnauthorizedError   | 401                           | application/problem+json      |
| sdkerrors.ForbiddenError      | 403                           | application/problem+json      |
| sdkerrors.NotFoundError       | 404                           | application/problem+json      |
| sdkerrors.BaseError           | 500                           | application/problem+json      |
| sdkerrors.NotImplementedError | 501                           | application/problem+json      |
| sdkerrors.ServiceUnavailable  | 503                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## DeleteWorkspace

Delete an individual workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-workspace" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/workspaces/{workspaceIdOrName}" -->
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

    res, err := s.Workspaces.DeleteWorkspace(ctx, "e3bd38a9-fbcc-4592-895a-ef1c2a5941b2", "2747d1e5-8246-4f65-a939-b392f1ee17f8")
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
| `controlPlaneID`                                         | `string`                                                 | :heavy_check_mark:                                       | ID of a control plane                                    |                                                          |
| `workspaceIDOrName`                                      | `string`                                                 | :heavy_check_mark:                                       | ID or name of the workspace to retreive                  | 2747d1e5-8246-4f65-a939-b392f1ee17f8                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteWorkspaceResponse](../../models/operations/deleteworkspaceresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequestError     | 400                           | application/problem+json      |
| sdkerrors.UnauthorizedError   | 401                           | application/problem+json      |
| sdkerrors.ForbiddenError      | 403                           | application/problem+json      |
| sdkerrors.NotFoundError       | 404                           | application/problem+json      |
| sdkerrors.BaseError           | 500                           | application/problem+json      |
| sdkerrors.NotImplementedError | 501                           | application/problem+json      |
| sdkerrors.ServiceUnavailable  | 503                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## UpsertWorkspace

Create or Update Workspace using ID or name.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-workspace" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/workspaces/{workspaceIdOrName}" -->
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

    res, err := s.Workspaces.UpsertWorkspace(ctx, operations.UpsertWorkspaceRequest{
        ControlPlaneID: "801690c4-cc1e-4dea-a405-bbe6924e39b4",
        WorkspaceIDOrName: "2747d1e5-8246-4f65-a939-b392f1ee17f8",
        Workspace: components.WorkspaceInput{
            Name: "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Workspace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpsertWorkspaceRequest](../../models/operations/upsertworkspacerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpsertWorkspaceResponse](../../models/operations/upsertworkspaceresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequestError     | 400                           | application/problem+json      |
| sdkerrors.UnauthorizedError   | 401                           | application/problem+json      |
| sdkerrors.ForbiddenError      | 403                           | application/problem+json      |
| sdkerrors.NotFoundError       | 404                           | application/problem+json      |
| sdkerrors.BaseError           | 500                           | application/problem+json      |
| sdkerrors.NotImplementedError | 501                           | application/problem+json      |
| sdkerrors.ServiceUnavailable  | 503                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |