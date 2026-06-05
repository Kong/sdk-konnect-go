# CustomPluginSchemas

## Overview

Custom Plugin Schemas

### Available Operations

* [ListPluginSchemas](#listpluginschemas) - List Custom Plugin Schemas
* [CreatePluginSchemas](#createpluginschemas) - Upload custom plugin schema
* [GetPluginSchema](#getpluginschema) - Get a custom plugin schema
* [DeletePluginSchemas](#deletepluginschemas) - Delete custom plugin schema
* [UpdatePluginSchemas](#updatepluginschemas) - Create or update a custom plugin schema
* [ListPluginSchemasInWorkspace](#listpluginschemasinworkspace) - List Custom Plugin Schemas in a workspace
* [CreatePluginSchemasInWorkspace](#createpluginschemasinworkspace) - Upload custom plugin schema in a workspace
* [GetPluginSchemaInWorkspace](#getpluginschemainworkspace) - Get a custom plugin schema in a workspace
* [DeletePluginSchemasInWorkspace](#deletepluginschemasinworkspace) - Delete custom plugin schema in a workspace
* [UpdatePluginSchemasInWorkspace](#updatepluginschemasinworkspace) - Create or update a custom plugin schema in a workspace

## ListPluginSchemas

Returns an array of custom plugins schemas associated with a control plane.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-plugin-schemas" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/plugin-schemas" -->
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

    res, err := s.CustomPluginSchemas.ListPluginSchemas(ctx, operations.ListPluginSchemasRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPluginSchemas != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListPluginSchemasRequest](../../models/operations/listpluginschemasrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListPluginSchemasResponse](../../models/operations/listpluginschemasresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreatePluginSchemas

Upload a custom plugin schema associated with a control plane.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-plugin-schemas" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/plugin-schemas" -->
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

    res, err := s.CustomPluginSchemas.CreatePluginSchemas(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", &components.CreatePluginSchemas{
        LuaSchema: "return { name = \"myplugin\", fields = { { config = { type = \"record\", fields = { } } } } }",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PluginSchemas != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `createPluginSchemas`                                                              | [*components.CreatePluginSchemas](../../models/components/createpluginschemas.md)  | :heavy_minus_sign:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.CreatePluginSchemasResponse](../../models/operations/createpluginschemasresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.KonnectCPLegacyBadRequestError   | 400                                        | application/json                           |
| sdkerrors.KonnectCPLegacyUnauthorizedError | 401                                        | application/json                           |
| sdkerrors.KonnectCPLegacyForbiddenError    | 403                                        | application/json                           |
| sdkerrors.KonnectCPLegacyConflictError     | 409                                        | application/json                           |
| sdkerrors.SDKError                         | 4XX, 5XX                                   | \*/\*                                      |

## GetPluginSchema

Returns information about a custom plugin from a given name.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-plugin-schema" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/plugin-schemas/{name}" -->
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

    res, err := s.CustomPluginSchemas.GetPluginSchema(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "myplugin")
    if err != nil {
        log.Fatal(err)
    }
    if res.PluginSchemas != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `name`                                                                             | `string`                                                                           | :heavy_check_mark:                                                                 | The custom plugin name                                                             | myplugin                                                                           |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetPluginSchemaResponse](../../models/operations/getpluginschemaresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.KonnectCPLegacyUnauthorizedError | 401                                        | application/json                           |
| sdkerrors.KonnectCPLegacyForbiddenError    | 403                                        | application/json                           |
| sdkerrors.KonnectCPLegacyNotFoundError     | 404                                        | application/json                           |
| sdkerrors.SDKError                         | 4XX, 5XX                                   | \*/\*                                      |

## DeletePluginSchemas

Delete an individual custom plugin schema.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-plugin-schemas" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/plugin-schemas/{name}" -->
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

    res, err := s.CustomPluginSchemas.DeletePluginSchemas(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "myplugin")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `name`                                                                             | `string`                                                                           | :heavy_check_mark:                                                                 | The custom plugin name                                                             | myplugin                                                                           |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.DeletePluginSchemasResponse](../../models/operations/deletepluginschemasresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.KonnectCPLegacyUnauthorizedError | 401                                        | application/json                           |
| sdkerrors.KonnectCPLegacyForbiddenError    | 403                                        | application/json                           |
| sdkerrors.KonnectCPLegacyNotFoundError     | 404                                        | application/json                           |
| sdkerrors.SDKError                         | 4XX, 5XX                                   | \*/\*                                      |

## UpdatePluginSchemas

Create or update an individual custom plugin schema.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-plugin-schemas" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/plugin-schemas/{name}" -->
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

    res, err := s.CustomPluginSchemas.UpdatePluginSchemas(ctx, operations.UpdatePluginSchemasRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Name: "myplugin",
        CreatePluginSchemas: &components.CreatePluginSchemas{
            LuaSchema: "return { name = \"myplugin\", fields = { { config = { type = \"record\", fields = { } } } } }",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PluginSchemas != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdatePluginSchemasRequest](../../models/operations/updatepluginschemasrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdatePluginSchemasResponse](../../models/operations/updatepluginschemasresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.KonnectCPLegacyBadRequestError   | 400                                        | application/json                           |
| sdkerrors.KonnectCPLegacyUnauthorizedError | 401                                        | application/json                           |
| sdkerrors.KonnectCPLegacyForbiddenError    | 403                                        | application/json                           |
| sdkerrors.KonnectCPLegacyNotFoundError     | 404                                        | application/json                           |
| sdkerrors.SDKError                         | 4XX, 5XX                                   | \*/\*                                      |

## ListPluginSchemasInWorkspace

Returns an array of custom plugin schemas associated with a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-plugin-schemas-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/plugin-schemas" -->
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

    res, err := s.CustomPluginSchemas.ListPluginSchemasInWorkspace(ctx, operations.ListPluginSchemasInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPluginSchemas != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListPluginSchemasInWorkspaceRequest](../../models/operations/listpluginschemasinworkspacerequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ListPluginSchemasInWorkspaceResponse](../../models/operations/listpluginschemasinworkspaceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreatePluginSchemasInWorkspace

Upload a custom plugin schema associated with a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-plugin-schemas-in-workspace" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/plugin-schemas" -->
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

    res, err := s.CustomPluginSchemas.CreatePluginSchemasInWorkspace(ctx, operations.CreatePluginSchemasInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        CreatePluginSchemas: &components.CreatePluginSchemas{
            LuaSchema: "return { name = \"myplugin\", fields = { { config = { type = \"record\", fields = { } } } } }",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PluginSchemas != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.CreatePluginSchemasInWorkspaceRequest](../../models/operations/createpluginschemasinworkspacerequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.CreatePluginSchemasInWorkspaceResponse](../../models/operations/createpluginschemasinworkspaceresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.KonnectCPLegacyBadRequestError   | 400                                        | application/json                           |
| sdkerrors.KonnectCPLegacyUnauthorizedError | 401                                        | application/json                           |
| sdkerrors.KonnectCPLegacyForbiddenError    | 403                                        | application/json                           |
| sdkerrors.KonnectCPLegacyConflictError     | 409                                        | application/json                           |
| sdkerrors.SDKError                         | 4XX, 5XX                                   | \*/\*                                      |

## GetPluginSchemaInWorkspace

Returns information about a custom plugin schema from a given name in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-plugin-schema-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/plugin-schemas/{name}" -->
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

    res, err := s.CustomPluginSchemas.GetPluginSchemaInWorkspace(ctx, operations.GetPluginSchemaInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Name: "myplugin",
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PluginSchemas != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetPluginSchemaInWorkspaceRequest](../../models/operations/getpluginschemainworkspacerequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.GetPluginSchemaInWorkspaceResponse](../../models/operations/getpluginschemainworkspaceresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.KonnectCPLegacyUnauthorizedError | 401                                        | application/json                           |
| sdkerrors.KonnectCPLegacyForbiddenError    | 403                                        | application/json                           |
| sdkerrors.KonnectCPLegacyNotFoundError     | 404                                        | application/json                           |
| sdkerrors.SDKError                         | 4XX, 5XX                                   | \*/\*                                      |

## DeletePluginSchemasInWorkspace

Delete an individual custom plugin schema in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-plugin-schemas-in-workspace" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/plugin-schemas/{name}" -->
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

    res, err := s.CustomPluginSchemas.DeletePluginSchemasInWorkspace(ctx, operations.DeletePluginSchemasInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Name: "myplugin",
        Workspace: "team-payments",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.DeletePluginSchemasInWorkspaceRequest](../../models/operations/deletepluginschemasinworkspacerequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.DeletePluginSchemasInWorkspaceResponse](../../models/operations/deletepluginschemasinworkspaceresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.KonnectCPLegacyUnauthorizedError | 401                                        | application/json                           |
| sdkerrors.KonnectCPLegacyForbiddenError    | 403                                        | application/json                           |
| sdkerrors.KonnectCPLegacyNotFoundError     | 404                                        | application/json                           |
| sdkerrors.SDKError                         | 4XX, 5XX                                   | \*/\*                                      |

## UpdatePluginSchemasInWorkspace

Create or update an individual custom plugin schema in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-plugin-schemas-in-workspace" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/plugin-schemas/{name}" -->
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

    res, err := s.CustomPluginSchemas.UpdatePluginSchemasInWorkspace(ctx, operations.UpdatePluginSchemasInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Name: "myplugin",
        Workspace: "team-payments",
        CreatePluginSchemas: &components.CreatePluginSchemas{
            LuaSchema: "return { name = \"myplugin\", fields = { { config = { type = \"record\", fields = { } } } } }",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PluginSchemas != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.UpdatePluginSchemasInWorkspaceRequest](../../models/operations/updatepluginschemasinworkspacerequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.UpdatePluginSchemasInWorkspaceResponse](../../models/operations/updatepluginschemasinworkspaceresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.KonnectCPLegacyBadRequestError   | 400                                        | application/json                           |
| sdkerrors.KonnectCPLegacyUnauthorizedError | 401                                        | application/json                           |
| sdkerrors.KonnectCPLegacyForbiddenError    | 403                                        | application/json                           |
| sdkerrors.KonnectCPLegacyNotFoundError     | 404                                        | application/json                           |
| sdkerrors.SDKError                         | 4XX, 5XX                                   | \*/\*                                      |