# CustomPlugins
(*CustomPlugins*)

## Overview

### Available Operations

* [ListCustomPlugin](#listcustomplugin) - List all CustomPlugins
* [CreateCustomPlugin](#createcustomplugin) - Create a new CustomPlugin
* [DeleteCustomPlugin](#deletecustomplugin) - Delete a CustomPlugin
* [GetCustomPlugin](#getcustomplugin) - Get a CustomPlugin
* [UpsertCustomPlugin](#upsertcustomplugin) - Upsert a CustomPlugin

## ListCustomPlugin

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

List all CustomPlugins

### Example Usage

<!-- UsageSnippet language="go" operationID="list-custom-plugin" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/custom-plugins" -->
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

    res, err := s.CustomPlugins.ListCustomPlugin(ctx, operations.ListCustomPluginRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Tags: sdkkonnectgo.Pointer("tag1,tag2"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListCustomPluginRequest](../../models/operations/listcustompluginrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListCustomPluginResponse](../../models/operations/listcustompluginresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## CreateCustomPlugin

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Create a new CustomPlugin

### Example Usage

<!-- UsageSnippet language="go" operationID="create-custom-plugin" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/custom-plugins" -->
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

    res, err := s.CustomPlugins.CreateCustomPlugin(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.CustomPlugin{
        Handler: "return { VERSION = '1.0,0', PRIORITY = 500, access = function(self, config) kong.service.request.set_header(config.name, config.value) end }",
        ID: sdkkonnectgo.Pointer("868346aa-1105-4b77-8346-aa1105fb77c4"),
        Name: "set-header",
        Schema: "return { name = 'set-header', fields = { { protocols = require('kong.db.schema.typedefs').protocols_http }, { config = { type = 'record', fields = { { name = { description = 'The name of the header to set.', type = 'string', required = true } }, { value = { description = 'The value for the header.', type = 'string', required = true } } } } } } }",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomPlugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | Example                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| `controlPlaneID`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | The UUID of your control plane. This variable is available in the Konnect manager.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| `customPlugin`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [components.CustomPlugin](../../models/components/customplugin.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | Description of the new CustomPlugin for creation                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | {<br/>"handler": "return { VERSION = '1.0,0', PRIORITY = 500, access = function(self, config) kong.service.request.set_header(config.name, config.value) end }",<br/>"id": "868346aa-1105-4b77-8346-aa1105fb77c4",<br/>"name": "set-header",<br/>"schema": "return { name = 'set-header', fields = { { protocols = require('kong.db.schema.typedefs').protocols_http }, { config = { type = 'record', fields = { { name = { description = 'The name of the header to set.', type = 'string', required = true } }, { value = { description = 'The value for the header.', type = 'string', required = true } } } } } } }"<br/>} |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |

### Response

**[*operations.CreateCustomPluginResponse](../../models/operations/createcustompluginresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteCustomPlugin

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Delete a CustomPlugin

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-custom-plugin" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/custom-plugins/{CustomPluginId}" -->
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

    res, err := s.CustomPlugins.DeleteCustomPlugin(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "")
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
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `customPluginID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the CustomPlugin to lookup                                                   |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.DeleteCustomPluginResponse](../../models/operations/deletecustompluginresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetCustomPlugin

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Get a CustomPlugin using ID or name.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-custom-plugin" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/custom-plugins/{CustomPluginId}" -->
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

    res, err := s.CustomPlugins.GetCustomPlugin(ctx, "", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomPlugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `customPluginID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the CustomPlugin to lookup                                                   |                                                                                    |
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetCustomPluginResponse](../../models/operations/getcustompluginresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertCustomPlugin

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Create or Update CustomPlugin using ID or name.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-custom-plugin" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/custom-plugins/{CustomPluginId}" -->
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

    res, err := s.CustomPlugins.UpsertCustomPlugin(ctx, operations.UpsertCustomPluginRequest{
        CustomPluginID: "",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        CustomPlugin: components.CustomPlugin{
            Handler: "return { VERSION = '1.0,0', PRIORITY = 500, access = function(self, config) kong.service.request.set_header(config.name, config.value) end }",
            ID: sdkkonnectgo.Pointer("868346aa-1105-4b77-8346-aa1105fb77c4"),
            Name: "set-header",
            Schema: "return { name = 'set-header', fields = { { protocols = require('kong.db.schema.typedefs').protocols_http }, { config = { type = 'record', fields = { { name = { description = 'The name of the header to set.', type = 'string', required = true } }, { value = { description = 'The value for the header.', type = 'string', required = true } } } } } } }",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomPlugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpsertCustomPluginRequest](../../models/operations/upsertcustompluginrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpsertCustomPluginResponse](../../models/operations/upsertcustompluginresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |