# EventGatewaySchemaRegistries

## Overview

Configure a schema registry that can be used to validate payloads when producing/consuming messages


### Available Operations

* [ListEventGatewaySchemaRegistries](#listeventgatewayschemaregistries) - List Schema Registries
* [CreateEventGatewaySchemaRegistry](#createeventgatewayschemaregistry) - Create Schema Registry
* [GetEventGatewaySchemaRegistry](#geteventgatewayschemaregistry) - Get a Schema Registry
* [UpdateEventGatewaySchemaRegistry](#updateeventgatewayschemaregistry) - Update Schema Registry
* [DeleteEventGatewaySchemaRegistry](#deleteeventgatewayschemaregistry) - Delete Schema Registry

## ListEventGatewaySchemaRegistries

Returns a list of schema registries associated with the specified Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-event-gateway-schema-registries" method="get" path="/v1/event-gateways/{gatewayId}/schema-registries" -->
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

    res, err := s.EventGatewaySchemaRegistries.ListEventGatewaySchemaRegistries(ctx, operations.ListEventGatewaySchemaRegistriesRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListSchemaRegistriesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.ListEventGatewaySchemaRegistriesRequest](../../models/operations/listeventgatewayschemaregistriesrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.ListEventGatewaySchemaRegistriesResponse](../../models/operations/listeventgatewayschemaregistriesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateEventGatewaySchemaRegistry

Creates a new schema registry associated with the specified Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-event-gateway-schema-registry" method="post" path="/v1/event-gateways/{gatewayId}/schema-registries" -->
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

    res, err := s.EventGatewaySchemaRegistries.CreateEventGatewaySchemaRegistry(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", sdkkonnectgo.Pointer(components.CreateSchemaRegistryCreateConfluent(
        components.SchemaRegistryConfluent{
            Name: "<value>",
            Config: components.SchemaRegistryConfluentConfig{
                SchemaType: components.SchemaTypeAvro,
                Endpoint: "https://broken-tuber.net/",
                Authentication: sdkkonnectgo.Pointer(components.CreateSchemaRegistryAuthenticationSchemeBasic(
                    components.SchemaRegistryAuthenticationBasic{
                        Username: "Stanley71",
                        Password: "${vault.env['MY_ENV_VAR']}",
                    },
                )),
            },
        },
    )))
    if err != nil {
        log.Fatal(err)
    }
    if res.SchemaRegistry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         | Example                                                                             |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ctx`                                                                               | [context.Context](https://pkg.go.dev/context#Context)                               | :heavy_check_mark:                                                                  | The context to use for the request.                                                 |                                                                                     |
| `gatewayID`                                                                         | *string*                                                                            | :heavy_check_mark:                                                                  | The UUID of your Gateway.                                                           | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                |
| `schemaRegistryCreate`                                                              | [*components.SchemaRegistryCreate](../../models/components/schemaregistrycreate.md) | :heavy_minus_sign:                                                                  | The request schema for creating a schema registry.                                  |                                                                                     |
| `opts`                                                                              | [][operations.Option](../../models/operations/option.md)                            | :heavy_minus_sign:                                                                  | The options for this request.                                                       |                                                                                     |

### Response

**[*operations.CreateEventGatewaySchemaRegistryResponse](../../models/operations/createeventgatewayschemaregistryresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetEventGatewaySchemaRegistry

Returns information about a specific schema registry associated with the Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-event-gateway-schema-registry" method="get" path="/v1/event-gateways/{gatewayId}/schema-registries/{schemaRegistryId}" -->
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

    res, err := s.EventGatewaySchemaRegistries.GetEventGatewaySchemaRegistry(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "cde6a4b4-3ebf-407e-8879-4fb69cbecb40")
    if err != nil {
        log.Fatal(err)
    }
    if res.SchemaRegistry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | *string*                                                 | :heavy_check_mark:                                       | The UUID of your Gateway.                                | 9524ec7d-36d9-465d-a8c5-83a3c9390458                     |
| `schemaRegistryID`                                       | *string*                                                 | :heavy_check_mark:                                       | The ID of the Schema Registry.                           |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetEventGatewaySchemaRegistryResponse](../../models/operations/geteventgatewayschemaregistryresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateEventGatewaySchemaRegistry

Updates an existing schema registry associated with the specified Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-event-gateway-schema-registry" method="put" path="/v1/event-gateways/{gatewayId}/schema-registries/{schemaRegistryId}" -->
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

    res, err := s.EventGatewaySchemaRegistries.UpdateEventGatewaySchemaRegistry(ctx, operations.UpdateEventGatewaySchemaRegistryRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        SchemaRegistryID: "703ba762-1708-4a8e-b805-03811b6dbf55",
        SchemaRegistryUpdate: sdkkonnectgo.Pointer(components.CreateSchemaRegistryUpdateConfluent(
            components.SchemaRegistryConfluentSensitiveDataAware{
                Name: "<value>",
                Config: components.SchemaRegistryConfluentConfigSensitiveDataAware{
                    SchemaType: components.SchemaRegistryConfluentConfigSensitiveDataAwareSchemaTypeAvro,
                    Endpoint: "https://polished-caption.info/",
                    Authentication: sdkkonnectgo.Pointer(components.CreateSchemaRegistryAuthenticationSensitiveDataAwareSchemeBasic(
                        components.SchemaRegistryAuthenticationBasicSensitiveDataAware{
                            Username: "Kelton_Lind-Spencer42",
                        },
                    )),
                },
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SchemaRegistry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.UpdateEventGatewaySchemaRegistryRequest](../../models/operations/updateeventgatewayschemaregistryrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.UpdateEventGatewaySchemaRegistryResponse](../../models/operations/updateeventgatewayschemaregistryresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteEventGatewaySchemaRegistry

Deletes a specific schema registry associated with the Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-event-gateway-schema-registry" method="delete" path="/v1/event-gateways/{gatewayId}/schema-registries/{schemaRegistryId}" -->
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

    res, err := s.EventGatewaySchemaRegistries.DeleteEventGatewaySchemaRegistry(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "6b0a69e8-8fc1-4c90-88c7-b9e13fb55c6a")
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
| `gatewayID`                                              | *string*                                                 | :heavy_check_mark:                                       | The UUID of your Gateway.                                | 9524ec7d-36d9-465d-a8c5-83a3c9390458                     |
| `schemaRegistryID`                                       | *string*                                                 | :heavy_check_mark:                                       | The ID of the Schema Registry.                           |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteEventGatewaySchemaRegistryResponse](../../models/operations/deleteeventgatewayschemaregistryresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |