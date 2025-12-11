# Plugins

## Overview

A plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. Plugins let you add functionality to services that run behind a Kong Gateway instance, like authentication or rate limiting.
You can find more information about available plugins and which values each plugin accepts at the [Plugin Hub](https://developer.konghq.com/plugins/).
<br><br>
When adding a plugin configuration to a service, the plugin will run on every request made by a client to that service. If a plugin needs to be tuned to different values for some specific consumers, you can do so by creating a separate plugin instance that specifies both the service and the consumer, through the service and consumer fields.

### Available Operations

* [ListPluginWithConsumerGroup](#listpluginwithconsumergroup) - List all Plugins associated with a Consumer Group
* [CreatePluginWithConsumerGroup](#createpluginwithconsumergroup) - Create a new Plugin associated with a Consumer Group
* [DeletePluginWithConsumerGroup](#deletepluginwithconsumergroup) - Delete a a Plugin associated with a Consumer Group
* [GetPluginWithConsumerGroup](#getpluginwithconsumergroup) - Get a Plugin associated with a Consumer Group
* [UpsertPluginWithConsumerGroup](#upsertpluginwithconsumergroup) - Upsert a Plugin associated with a Consumer Group
* [ListPluginWithConsumer](#listpluginwithconsumer) - List all Plugins associated with a Consumer
* [CreatePluginWithConsumer](#createpluginwithconsumer) - Create a new Plugin associated with a Consumer
* [DeletePluginWithConsumer](#deletepluginwithconsumer) - Delete a a Plugin associated with a Consumer
* [GetPluginWithConsumer](#getpluginwithconsumer) - Get a Plugin associated with a Consumer
* [UpsertPluginWithConsumer](#upsertpluginwithconsumer) - Upsert a Plugin associated with a Consumer
* [ListPlugin](#listplugin) - List all Plugins
* [CreatePlugin](#createplugin) - Create a new Plugin
* [DeletePlugin](#deleteplugin) - Delete a Plugin
* [GetPlugin](#getplugin) - Get a Plugin
* [UpsertPlugin](#upsertplugin) - Upsert a Plugin
* [ListPluginWithRoute](#listpluginwithroute) - List all Plugins associated with a Route
* [CreatePluginWithRoute](#createpluginwithroute) - Create a new Plugin associated with a Route
* [DeletePluginWithRoute](#deletepluginwithroute) - Delete a a Plugin associated with a Route
* [GetPluginWithRoute](#getpluginwithroute) - Get a Plugin associated with a Route
* [UpsertPluginWithRoute](#upsertpluginwithroute) - Upsert a Plugin associated with a Route
* [FetchPluginSchema](#fetchpluginschema) - Get plugin schema
* [ListPluginWithService](#listpluginwithservice) - List all Plugins associated with a Service
* [CreatePluginWithService](#createpluginwithservice) - Create a new Plugin associated with a Service
* [DeletePluginWithService](#deletepluginwithservice) - Delete a a Plugin associated with a Service
* [GetPluginWithService](#getpluginwithservice) - Get a Plugin associated with a Service
* [UpsertPluginWithService](#upsertpluginwithservice) - Upsert a Plugin associated with a Service

## ListPluginWithConsumerGroup

List all Plugins associated with a Consumer Group

### Example Usage

<!-- UsageSnippet language="go" operationID="list-plugin-with-consumer_group" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/consumer_groups/{ConsumerGroupId}/plugins" -->
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

    res, err := s.Plugins.ListPluginWithConsumerGroup(ctx, operations.ListPluginWithConsumerGroupRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerGroupID: "",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListPluginWithConsumerGroupRequest](../../models/operations/listpluginwithconsumergrouprequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ListPluginWithConsumerGroupResponse](../../models/operations/listpluginwithconsumergroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreatePluginWithConsumerGroup

Create a new Plugin associated with a Consumer Group

### Example Usage

<!-- UsageSnippet language="go" operationID="create-plugin-with-consumer_group" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/consumer_groups/{ConsumerGroupId}/plugins" -->
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

    res, err := s.Plugins.CreatePluginWithConsumerGroup(ctx, operations.CreatePluginWithConsumerGroupRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerGroupID: "",
        PluginWithoutParents: components.PluginWithoutParents{
            Config: map[string]any{
                "anonymous": "<value>",
                "hide_credentials": false,
                "key_in_body": false,
                "key_in_header": true,
                "key_in_query": true,
                "key_names": []any{
                    "apikey",
                },
                "run_on_preflight": true,
            },
            ID: sdkkonnectgo.Pointer("3fd1eea1-885a-4011-b986-289943ff8177"),
            Name: "key-auth",
            Partials: []components.PluginWithoutParentsPartials{
                components.PluginWithoutParentsPartials{
                    ID: sdkkonnectgo.Pointer("cff1230a-00f7-4ae8-b376-c370f0eb4dae"),
                    Name: sdkkonnectgo.Pointer("foo-partial"),
                    Path: sdkkonnectgo.Pointer("config.redis"),
                },
                components.PluginWithoutParentsPartials{
                    ID: sdkkonnectgo.Pointer("129ee345-cba8-4e55-9d6d-93c223ff91ae"),
                    Name: sdkkonnectgo.Pointer("bar-partial"),
                    Path: sdkkonnectgo.Pointer("config.redis"),
                },
            },
            Protocols: []components.PluginWithoutParentsProtocols{
                components.PluginWithoutParentsProtocolsGrpc,
                components.PluginWithoutParentsProtocolsGrpcs,
                components.PluginWithoutParentsProtocolsHTTP,
                components.PluginWithoutParentsProtocolsHTTPS,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Plugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.CreatePluginWithConsumerGroupRequest](../../models/operations/createpluginwithconsumergrouprequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.CreatePluginWithConsumerGroupResponse](../../models/operations/createpluginwithconsumergroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeletePluginWithConsumerGroup

Delete a a Plugin associated with a Consumer Group using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-plugin-with-consumer_group" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/consumer_groups/{ConsumerGroupId}/plugins/{PluginId}" -->
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

    res, err := s.Plugins.DeletePluginWithConsumerGroup(ctx, operations.DeletePluginWithConsumerGroupRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerGroupID: "",
        PluginID: "3473c251-5b6c-4f45-b1ff-7ede735a366d",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.DeletePluginWithConsumerGroupRequest](../../models/operations/deletepluginwithconsumergrouprequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.DeletePluginWithConsumerGroupResponse](../../models/operations/deletepluginwithconsumergroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPluginWithConsumerGroup

Get a Plugin associated with a Consumer Group using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-plugin-with-consumer_group" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/consumer_groups/{ConsumerGroupId}/plugins/{PluginId}" -->
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

    res, err := s.Plugins.GetPluginWithConsumerGroup(ctx, operations.GetPluginWithConsumerGroupRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerGroupID: "",
        PluginID: "3473c251-5b6c-4f45-b1ff-7ede735a366d",
        ExpandPartials: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Plugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetPluginWithConsumerGroupRequest](../../models/operations/getpluginwithconsumergrouprequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.GetPluginWithConsumerGroupResponse](../../models/operations/getpluginwithconsumergroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpsertPluginWithConsumerGroup

Create or Update a Plugin associated with a Consumer Group using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-plugin-with-consumer_group" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/consumer_groups/{ConsumerGroupId}/plugins/{PluginId}" -->
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

    res, err := s.Plugins.UpsertPluginWithConsumerGroup(ctx, operations.UpsertPluginWithConsumerGroupRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerGroupID: "",
        PluginID: "3473c251-5b6c-4f45-b1ff-7ede735a366d",
        PluginWithoutParents: components.PluginWithoutParents{
            Config: map[string]any{
                "anonymous": "<value>",
                "hide_credentials": false,
                "key_in_body": false,
                "key_in_header": true,
                "key_in_query": true,
                "key_names": []any{
                    "apikey",
                },
                "run_on_preflight": true,
            },
            ID: sdkkonnectgo.Pointer("3fd1eea1-885a-4011-b986-289943ff8177"),
            Name: "key-auth",
            Partials: []components.PluginWithoutParentsPartials{
                components.PluginWithoutParentsPartials{
                    ID: sdkkonnectgo.Pointer("cff1230a-00f7-4ae8-b376-c370f0eb4dae"),
                    Name: sdkkonnectgo.Pointer("foo-partial"),
                    Path: sdkkonnectgo.Pointer("config.redis"),
                },
                components.PluginWithoutParentsPartials{
                    ID: sdkkonnectgo.Pointer("129ee345-cba8-4e55-9d6d-93c223ff91ae"),
                    Name: sdkkonnectgo.Pointer("bar-partial"),
                    Path: sdkkonnectgo.Pointer("config.redis"),
                },
            },
            Protocols: []components.PluginWithoutParentsProtocols{
                components.PluginWithoutParentsProtocolsGrpc,
                components.PluginWithoutParentsProtocolsGrpcs,
                components.PluginWithoutParentsProtocolsHTTP,
                components.PluginWithoutParentsProtocolsHTTPS,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Plugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.UpsertPluginWithConsumerGroupRequest](../../models/operations/upsertpluginwithconsumergrouprequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.UpsertPluginWithConsumerGroupResponse](../../models/operations/upsertpluginwithconsumergroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPluginWithConsumer

List all Plugins associated with a Consumer

### Example Usage

<!-- UsageSnippet language="go" operationID="list-plugin-with-consumer" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/plugins" -->
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

    res, err := s.Plugins.ListPluginWithConsumer(ctx, operations.ListPluginWithConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "f28acbfa-c866-4587-b688-0208ac24df21",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListPluginWithConsumerRequest](../../models/operations/listpluginwithconsumerrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListPluginWithConsumerResponse](../../models/operations/listpluginwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreatePluginWithConsumer

Create a new Plugin associated with a Consumer

### Example Usage

<!-- UsageSnippet language="go" operationID="create-plugin-with-consumer" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/plugins" -->
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

    res, err := s.Plugins.CreatePluginWithConsumer(ctx, operations.CreatePluginWithConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "f28acbfa-c866-4587-b688-0208ac24df21",
        PluginWithoutParents: components.PluginWithoutParents{
            Config: map[string]any{
                "anonymous": "<value>",
                "hide_credentials": false,
                "key_in_body": false,
                "key_in_header": true,
                "key_in_query": true,
                "key_names": []any{
                    "apikey",
                },
                "run_on_preflight": true,
            },
            ID: sdkkonnectgo.Pointer("3fd1eea1-885a-4011-b986-289943ff8177"),
            Name: "key-auth",
            Partials: []components.PluginWithoutParentsPartials{
                components.PluginWithoutParentsPartials{
                    ID: sdkkonnectgo.Pointer("cff1230a-00f7-4ae8-b376-c370f0eb4dae"),
                    Name: sdkkonnectgo.Pointer("foo-partial"),
                    Path: sdkkonnectgo.Pointer("config.redis"),
                },
                components.PluginWithoutParentsPartials{
                    ID: sdkkonnectgo.Pointer("129ee345-cba8-4e55-9d6d-93c223ff91ae"),
                    Name: sdkkonnectgo.Pointer("bar-partial"),
                    Path: sdkkonnectgo.Pointer("config.redis"),
                },
            },
            Protocols: []components.PluginWithoutParentsProtocols{
                components.PluginWithoutParentsProtocolsGrpc,
                components.PluginWithoutParentsProtocolsGrpcs,
                components.PluginWithoutParentsProtocolsHTTP,
                components.PluginWithoutParentsProtocolsHTTPS,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Plugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreatePluginWithConsumerRequest](../../models/operations/createpluginwithconsumerrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.CreatePluginWithConsumerResponse](../../models/operations/createpluginwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeletePluginWithConsumer

Delete a a Plugin associated with a Consumer using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-plugin-with-consumer" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/plugins/{PluginId}" -->
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

    res, err := s.Plugins.DeletePluginWithConsumer(ctx, operations.DeletePluginWithConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "f28acbfa-c866-4587-b688-0208ac24df21",
        PluginID: "3473c251-5b6c-4f45-b1ff-7ede735a366d",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.DeletePluginWithConsumerRequest](../../models/operations/deletepluginwithconsumerrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.DeletePluginWithConsumerResponse](../../models/operations/deletepluginwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPluginWithConsumer

Get a Plugin associated with a Consumer using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-plugin-with-consumer" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/plugins/{PluginId}" -->
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

    res, err := s.Plugins.GetPluginWithConsumer(ctx, operations.GetPluginWithConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "f28acbfa-c866-4587-b688-0208ac24df21",
        PluginID: "3473c251-5b6c-4f45-b1ff-7ede735a366d",
        ExpandPartials: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Plugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetPluginWithConsumerRequest](../../models/operations/getpluginwithconsumerrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetPluginWithConsumerResponse](../../models/operations/getpluginwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpsertPluginWithConsumer

Create or Update a Plugin associated with a Consumer using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-plugin-with-consumer" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerIdForNestedEntities}/plugins/{PluginId}" -->
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

    res, err := s.Plugins.UpsertPluginWithConsumer(ctx, operations.UpsertPluginWithConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "f28acbfa-c866-4587-b688-0208ac24df21",
        PluginID: "3473c251-5b6c-4f45-b1ff-7ede735a366d",
        PluginWithoutParents: components.PluginWithoutParents{
            Config: map[string]any{
                "anonymous": "<value>",
                "hide_credentials": false,
                "key_in_body": false,
                "key_in_header": true,
                "key_in_query": true,
                "key_names": []any{
                    "apikey",
                },
                "run_on_preflight": true,
            },
            ID: sdkkonnectgo.Pointer("3fd1eea1-885a-4011-b986-289943ff8177"),
            Name: "key-auth",
            Partials: []components.PluginWithoutParentsPartials{
                components.PluginWithoutParentsPartials{
                    ID: sdkkonnectgo.Pointer("cff1230a-00f7-4ae8-b376-c370f0eb4dae"),
                    Name: sdkkonnectgo.Pointer("foo-partial"),
                    Path: sdkkonnectgo.Pointer("config.redis"),
                },
                components.PluginWithoutParentsPartials{
                    ID: sdkkonnectgo.Pointer("129ee345-cba8-4e55-9d6d-93c223ff91ae"),
                    Name: sdkkonnectgo.Pointer("bar-partial"),
                    Path: sdkkonnectgo.Pointer("config.redis"),
                },
            },
            Protocols: []components.PluginWithoutParentsProtocols{
                components.PluginWithoutParentsProtocolsGrpc,
                components.PluginWithoutParentsProtocolsGrpcs,
                components.PluginWithoutParentsProtocolsHTTP,
                components.PluginWithoutParentsProtocolsHTTPS,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Plugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpsertPluginWithConsumerRequest](../../models/operations/upsertpluginwithconsumerrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpsertPluginWithConsumerResponse](../../models/operations/upsertpluginwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPlugin

List all Plugins

### Example Usage

<!-- UsageSnippet language="go" operationID="list-plugin" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/plugins" -->
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

    res, err := s.Plugins.ListPlugin(ctx, operations.ListPluginRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Tags: sdkkonnectgo.Pointer("tag1,tag2"),
        FilterNameContains: sdkkonnectgo.Pointer("john"),
        FilterNameEq: sdkkonnectgo.Pointer("john"),
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListPluginRequest](../../models/operations/listpluginrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.ListPluginResponse](../../models/operations/listpluginresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## CreatePlugin

Create a new Plugin

### Example Usage

<!-- UsageSnippet language="go" operationID="create-plugin" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/plugins" -->
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

    res, err := s.Plugins.CreatePlugin(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.Plugin{
        Config: map[string]any{
            "anonymous": "<value>",
            "hide_credentials": false,
            "key_in_body": false,
            "key_in_header": true,
            "key_in_query": true,
            "key_names": []any{
                "apikey",
            },
            "run_on_preflight": true,
        },
        ID: sdkkonnectgo.Pointer("3fd1eea1-885a-4011-b986-289943ff8177"),
        Name: "key-auth",
        Partials: []components.Partials{
            components.Partials{
                ID: sdkkonnectgo.Pointer("cff1230a-00f7-4ae8-b376-c370f0eb4dae"),
                Name: sdkkonnectgo.Pointer("foo-partial"),
                Path: sdkkonnectgo.Pointer("config.redis"),
            },
            components.Partials{
                ID: sdkkonnectgo.Pointer("129ee345-cba8-4e55-9d6d-93c223ff91ae"),
                Name: sdkkonnectgo.Pointer("bar-partial"),
                Path: sdkkonnectgo.Pointer("config.redis"),
            },
        },
        Protocols: []components.Protocols{
            components.ProtocolsGrpc,
            components.ProtocolsGrpcs,
            components.ProtocolsHTTP,
            components.ProtocolsHTTPS,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Plugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | Example                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| `controlPlaneID`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | The UUID of your control plane. This variable is available in the Konnect manager.                                                                                                                                                                                                                                                                                                                                                                                                                                                             | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| `plugin`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | [components.Plugin](../../models/components/plugin.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Description of the new Plugin for creation                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | {<br/>"config": {<br/>"anonymous": null,<br/>"hide_credentials": false,<br/>"key_in_body": false,<br/>"key_in_header": true,<br/>"key_in_query": true,<br/>"key_names": [<br/>"apikey"<br/>],<br/>"run_on_preflight": true<br/>},<br/>"enabled": true,<br/>"id": "3fd1eea1-885a-4011-b986-289943ff8177",<br/>"name": "key-auth",<br/>"partials": [<br/>{<br/>"id": "cff1230a-00f7-4ae8-b376-c370f0eb4dae",<br/>"name": "foo-partial",<br/>"path": "config.redis"<br/>},<br/>{<br/>"id": "129ee345-cba8-4e55-9d6d-93c223ff91ae",<br/>"name": "bar-partial",<br/>"path": "config.redis"<br/>}<br/>],<br/>"protocols": [<br/>"grpc",<br/>"grpcs",<br/>"http",<br/>"https"<br/>]<br/>} |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |

### Response

**[*operations.CreatePluginResponse](../../models/operations/createpluginresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeletePlugin

Delete a Plugin

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-plugin" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/plugins/{PluginId}" -->
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

    res, err := s.Plugins.DeletePlugin(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "3473c251-5b6c-4f45-b1ff-7ede735a366d")
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
| `pluginID`                                                                         | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Plugin to lookup                                                         | 3473c251-5b6c-4f45-b1ff-7ede735a366d                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.DeletePluginResponse](../../models/operations/deletepluginresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetPlugin

Get a Plugin using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-plugin" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/plugins/{PluginId}" -->
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

    res, err := s.Plugins.GetPlugin(ctx, operations.GetPluginRequest{
        PluginID: "3473c251-5b6c-4f45-b1ff-7ede735a366d",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ExpandPartials: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Plugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetPluginRequest](../../models/operations/getpluginrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.GetPluginResponse](../../models/operations/getpluginresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertPlugin

Create or Update Plugin using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-plugin" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/plugins/{PluginId}" -->
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

    res, err := s.Plugins.UpsertPlugin(ctx, operations.UpsertPluginRequest{
        PluginID: "3473c251-5b6c-4f45-b1ff-7ede735a366d",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Plugin: components.Plugin{
            Config: map[string]any{
                "anonymous": "<value>",
                "hide_credentials": false,
                "key_in_body": false,
                "key_in_header": true,
                "key_in_query": true,
                "key_names": []any{
                    "apikey",
                },
                "run_on_preflight": true,
            },
            ID: sdkkonnectgo.Pointer("3fd1eea1-885a-4011-b986-289943ff8177"),
            Name: "key-auth",
            Partials: []components.Partials{
                components.Partials{
                    ID: sdkkonnectgo.Pointer("cff1230a-00f7-4ae8-b376-c370f0eb4dae"),
                    Name: sdkkonnectgo.Pointer("foo-partial"),
                    Path: sdkkonnectgo.Pointer("config.redis"),
                },
                components.Partials{
                    ID: sdkkonnectgo.Pointer("129ee345-cba8-4e55-9d6d-93c223ff91ae"),
                    Name: sdkkonnectgo.Pointer("bar-partial"),
                    Path: sdkkonnectgo.Pointer("config.redis"),
                },
            },
            Protocols: []components.Protocols{
                components.ProtocolsGrpc,
                components.ProtocolsGrpcs,
                components.ProtocolsHTTP,
                components.ProtocolsHTTPS,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Plugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.UpsertPluginRequest](../../models/operations/upsertpluginrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.UpsertPluginResponse](../../models/operations/upsertpluginresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## ListPluginWithRoute

List all Plugins associated with a Route

### Example Usage

<!-- UsageSnippet language="go" operationID="list-plugin-with-route" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/routes/{RouteId}/plugins" -->
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

    res, err := s.Plugins.ListPluginWithRoute(ctx, operations.ListPluginWithRouteRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        RouteID: "a4326a41-aa12-44e3-93e4-6b6e58bfb9d7",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListPluginWithRouteRequest](../../models/operations/listpluginwithrouterequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListPluginWithRouteResponse](../../models/operations/listpluginwithrouteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreatePluginWithRoute

Create a new Plugin associated with a Route

### Example Usage

<!-- UsageSnippet language="go" operationID="create-plugin-with-route" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/routes/{RouteId}/plugins" -->
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

    res, err := s.Plugins.CreatePluginWithRoute(ctx, operations.CreatePluginWithRouteRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        RouteID: "a4326a41-aa12-44e3-93e4-6b6e58bfb9d7",
        PluginWithoutParents: components.PluginWithoutParents{
            Config: map[string]any{
                "anonymous": "<value>",
                "hide_credentials": false,
                "key_in_body": false,
                "key_in_header": true,
                "key_in_query": true,
                "key_names": []any{
                    "apikey",
                },
                "run_on_preflight": true,
            },
            ID: sdkkonnectgo.Pointer("3fd1eea1-885a-4011-b986-289943ff8177"),
            Name: "key-auth",
            Partials: []components.PluginWithoutParentsPartials{
                components.PluginWithoutParentsPartials{
                    ID: sdkkonnectgo.Pointer("cff1230a-00f7-4ae8-b376-c370f0eb4dae"),
                    Name: sdkkonnectgo.Pointer("foo-partial"),
                    Path: sdkkonnectgo.Pointer("config.redis"),
                },
                components.PluginWithoutParentsPartials{
                    ID: sdkkonnectgo.Pointer("129ee345-cba8-4e55-9d6d-93c223ff91ae"),
                    Name: sdkkonnectgo.Pointer("bar-partial"),
                    Path: sdkkonnectgo.Pointer("config.redis"),
                },
            },
            Protocols: []components.PluginWithoutParentsProtocols{
                components.PluginWithoutParentsProtocolsGrpc,
                components.PluginWithoutParentsProtocolsGrpcs,
                components.PluginWithoutParentsProtocolsHTTP,
                components.PluginWithoutParentsProtocolsHTTPS,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Plugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreatePluginWithRouteRequest](../../models/operations/createpluginwithrouterequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreatePluginWithRouteResponse](../../models/operations/createpluginwithrouteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeletePluginWithRoute

Delete a a Plugin associated with a Route using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-plugin-with-route" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/routes/{RouteId}/plugins/{PluginId}" -->
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

    res, err := s.Plugins.DeletePluginWithRoute(ctx, operations.DeletePluginWithRouteRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        RouteID: "a4326a41-aa12-44e3-93e4-6b6e58bfb9d7",
        PluginID: "3473c251-5b6c-4f45-b1ff-7ede735a366d",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.DeletePluginWithRouteRequest](../../models/operations/deletepluginwithrouterequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.DeletePluginWithRouteResponse](../../models/operations/deletepluginwithrouteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPluginWithRoute

Get a Plugin associated with a Route using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-plugin-with-route" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/routes/{RouteId}/plugins/{PluginId}" -->
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

    res, err := s.Plugins.GetPluginWithRoute(ctx, operations.GetPluginWithRouteRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        RouteID: "a4326a41-aa12-44e3-93e4-6b6e58bfb9d7",
        PluginID: "3473c251-5b6c-4f45-b1ff-7ede735a366d",
        ExpandPartials: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Plugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetPluginWithRouteRequest](../../models/operations/getpluginwithrouterequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetPluginWithRouteResponse](../../models/operations/getpluginwithrouteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpsertPluginWithRoute

Create or Update a Plugin associated with a Route using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-plugin-with-route" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/routes/{RouteId}/plugins/{PluginId}" -->
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

    res, err := s.Plugins.UpsertPluginWithRoute(ctx, operations.UpsertPluginWithRouteRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        RouteID: "a4326a41-aa12-44e3-93e4-6b6e58bfb9d7",
        PluginID: "3473c251-5b6c-4f45-b1ff-7ede735a366d",
        PluginWithoutParents: components.PluginWithoutParents{
            Config: map[string]any{
                "anonymous": "<value>",
                "hide_credentials": false,
                "key_in_body": false,
                "key_in_header": true,
                "key_in_query": true,
                "key_names": []any{
                    "apikey",
                },
                "run_on_preflight": true,
            },
            ID: sdkkonnectgo.Pointer("3fd1eea1-885a-4011-b986-289943ff8177"),
            Name: "key-auth",
            Partials: []components.PluginWithoutParentsPartials{
                components.PluginWithoutParentsPartials{
                    ID: sdkkonnectgo.Pointer("cff1230a-00f7-4ae8-b376-c370f0eb4dae"),
                    Name: sdkkonnectgo.Pointer("foo-partial"),
                    Path: sdkkonnectgo.Pointer("config.redis"),
                },
                components.PluginWithoutParentsPartials{
                    ID: sdkkonnectgo.Pointer("129ee345-cba8-4e55-9d6d-93c223ff91ae"),
                    Name: sdkkonnectgo.Pointer("bar-partial"),
                    Path: sdkkonnectgo.Pointer("config.redis"),
                },
            },
            Protocols: []components.PluginWithoutParentsProtocols{
                components.PluginWithoutParentsProtocolsGrpc,
                components.PluginWithoutParentsProtocolsGrpcs,
                components.PluginWithoutParentsProtocolsHTTP,
                components.PluginWithoutParentsProtocolsHTTPS,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Plugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpsertPluginWithRouteRequest](../../models/operations/upsertpluginwithrouterequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpsertPluginWithRouteResponse](../../models/operations/upsertpluginwithrouteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## FetchPluginSchema

Get the schema for a plugin

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-plugin-schema" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/schemas/plugins/{pluginName}" -->
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

    res, err := s.Plugins.FetchPluginSchema(ctx, "<value>", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetPluginSchemaResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `pluginName`                                                                       | *string*                                                                           | :heavy_check_mark:                                                                 | The name of the plugin                                                             |                                                                                    |
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.FetchPluginSchemaResponse](../../models/operations/fetchpluginschemaresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPluginWithService

List all Plugins associated with a Service

### Example Usage

<!-- UsageSnippet language="go" operationID="list-plugin-with-service" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/services/{ServiceId}/plugins" -->
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

    res, err := s.Plugins.ListPluginWithService(ctx, operations.ListPluginWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListPluginWithServiceRequest](../../models/operations/listpluginwithservicerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListPluginWithServiceResponse](../../models/operations/listpluginwithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreatePluginWithService

Create a new Plugin associated with a Service

### Example Usage

<!-- UsageSnippet language="go" operationID="create-plugin-with-service" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/services/{ServiceId}/plugins" -->
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

    res, err := s.Plugins.CreatePluginWithService(ctx, operations.CreatePluginWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        PluginWithoutParents: components.PluginWithoutParents{
            Config: map[string]any{
                "anonymous": "<value>",
                "hide_credentials": false,
                "key_in_body": false,
                "key_in_header": true,
                "key_in_query": true,
                "key_names": []any{
                    "apikey",
                },
                "run_on_preflight": true,
            },
            ID: sdkkonnectgo.Pointer("3fd1eea1-885a-4011-b986-289943ff8177"),
            Name: "key-auth",
            Partials: []components.PluginWithoutParentsPartials{
                components.PluginWithoutParentsPartials{
                    ID: sdkkonnectgo.Pointer("cff1230a-00f7-4ae8-b376-c370f0eb4dae"),
                    Name: sdkkonnectgo.Pointer("foo-partial"),
                    Path: sdkkonnectgo.Pointer("config.redis"),
                },
                components.PluginWithoutParentsPartials{
                    ID: sdkkonnectgo.Pointer("129ee345-cba8-4e55-9d6d-93c223ff91ae"),
                    Name: sdkkonnectgo.Pointer("bar-partial"),
                    Path: sdkkonnectgo.Pointer("config.redis"),
                },
            },
            Protocols: []components.PluginWithoutParentsProtocols{
                components.PluginWithoutParentsProtocolsGrpc,
                components.PluginWithoutParentsProtocolsGrpcs,
                components.PluginWithoutParentsProtocolsHTTP,
                components.PluginWithoutParentsProtocolsHTTPS,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Plugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreatePluginWithServiceRequest](../../models/operations/createpluginwithservicerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreatePluginWithServiceResponse](../../models/operations/createpluginwithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeletePluginWithService

Delete a a Plugin associated with a Service using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-plugin-with-service" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/services/{ServiceId}/plugins/{PluginId}" -->
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

    res, err := s.Plugins.DeletePluginWithService(ctx, operations.DeletePluginWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        PluginID: "3473c251-5b6c-4f45-b1ff-7ede735a366d",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.DeletePluginWithServiceRequest](../../models/operations/deletepluginwithservicerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.DeletePluginWithServiceResponse](../../models/operations/deletepluginwithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPluginWithService

Get a Plugin associated with a Service using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-plugin-with-service" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/services/{ServiceId}/plugins/{PluginId}" -->
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

    res, err := s.Plugins.GetPluginWithService(ctx, operations.GetPluginWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        PluginID: "3473c251-5b6c-4f45-b1ff-7ede735a366d",
        ExpandPartials: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Plugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetPluginWithServiceRequest](../../models/operations/getpluginwithservicerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetPluginWithServiceResponse](../../models/operations/getpluginwithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpsertPluginWithService

Create or Update a Plugin associated with a Service using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-plugin-with-service" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/services/{ServiceId}/plugins/{PluginId}" -->
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

    res, err := s.Plugins.UpsertPluginWithService(ctx, operations.UpsertPluginWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        PluginID: "3473c251-5b6c-4f45-b1ff-7ede735a366d",
        PluginWithoutParents: components.PluginWithoutParents{
            Config: map[string]any{
                "anonymous": "<value>",
                "hide_credentials": false,
                "key_in_body": false,
                "key_in_header": true,
                "key_in_query": true,
                "key_names": []any{
                    "apikey",
                },
                "run_on_preflight": true,
            },
            ID: sdkkonnectgo.Pointer("3fd1eea1-885a-4011-b986-289943ff8177"),
            Name: "key-auth",
            Partials: []components.PluginWithoutParentsPartials{
                components.PluginWithoutParentsPartials{
                    ID: sdkkonnectgo.Pointer("cff1230a-00f7-4ae8-b376-c370f0eb4dae"),
                    Name: sdkkonnectgo.Pointer("foo-partial"),
                    Path: sdkkonnectgo.Pointer("config.redis"),
                },
                components.PluginWithoutParentsPartials{
                    ID: sdkkonnectgo.Pointer("129ee345-cba8-4e55-9d6d-93c223ff91ae"),
                    Name: sdkkonnectgo.Pointer("bar-partial"),
                    Path: sdkkonnectgo.Pointer("config.redis"),
                },
            },
            Protocols: []components.PluginWithoutParentsProtocols{
                components.PluginWithoutParentsProtocolsGrpc,
                components.PluginWithoutParentsProtocolsGrpcs,
                components.PluginWithoutParentsProtocolsHTTP,
                components.PluginWithoutParentsProtocolsHTTPS,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Plugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpsertPluginWithServiceRequest](../../models/operations/upsertpluginwithservicerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpsertPluginWithServiceResponse](../../models/operations/upsertpluginwithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |