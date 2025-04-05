# Services
(*Services*)

## Overview

Service entities are abstractions of your microservice interfaces or formal APIs. For example, a service could be a data transformation microservice or a billing API.
<br><br>
The main attribute of a service is the destination URL for proxying traffic. This URL can be set as a single string or by specifying its protocol, host, port and path individually.
<br><br>
Services are associated to routes, and a single service can have many routes associated with it. Routes are entrypoints in Kong Gateway which define rules to match client requests. Once a route is matched, Kong Gateway proxies the request to its associated service. See the [Proxy Reference](https://docs.konghq.com/gateway/latest/how-kong-works/routing-traffic/) for a detailed explanation of how Kong proxies traffic.
<br><br>
Services can be both [tagged and filtered by tags](https://docs.konghq.com/gateway/latest/admin-api/#tags).


### Available Operations

* [ListService](#listservice) - List all Services
* [CreateService](#createservice) - Create a new Service
* [DeleteService](#deleteservice) - Delete a Service
* [GetService](#getservice) - Fetch a Service
* [UpsertService](#upsertservice) - Upsert a Service

## ListService

List all Services

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.Services.ListService(ctx, operations.ListServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Tags: sdkkonnectgo.String("tag1,tag2"),
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListServiceRequest](../../models/operations/listservicerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.ListServiceResponse](../../models/operations/listserviceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## CreateService

Create a new Service

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

    res, err := s.Services.CreateService(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.Service{
        Host: "example.internal",
        ID: sdkkonnectgo.String("49fd316e-c457-481c-9fc7-8079153e4f3c"),
        Name: sdkkonnectgo.String("example-service"),
        Path: sdkkonnectgo.String("/"),
        Port: sdkkonnectgo.Int64(80),
        Protocol: components.ProtocolHTTP.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Service != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          | Example                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |                                                                                                                                                      |
| `controlPlaneID`                                                                                                                                     | *string*                                                                                                                                             | :heavy_check_mark:                                                                                                                                   | The UUID of your control plane. This variable is available in the Konnect manager.                                                                   | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                                                                                 |
| `service`                                                                                                                                            | [components.Service](../../models/components/service.md)                                                                                             | :heavy_check_mark:                                                                                                                                   | Description of the new Service for creation                                                                                                          | {<br/>"host": "example.internal",<br/>"id": "49fd316e-c457-481c-9fc7-8079153e4f3c",<br/>"name": "example-service",<br/>"path": "/",<br/>"port": 80,<br/>"protocol": "http"<br/>} |
| `opts`                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |                                                                                                                                                      |

### Response

**[*operations.CreateServiceResponse](../../models/operations/createserviceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteService

Delete a Service

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

    res, err := s.Services.DeleteService(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "7fca84d6-7d37-4a74-a7b0-93e576089a41")
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
| `serviceID`                                                                        | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Service to lookup                                                        | 7fca84d6-7d37-4a74-a7b0-93e576089a41                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.DeleteServiceResponse](../../models/operations/deleteserviceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetService

Get a Service using ID or name.

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

    res, err := s.Services.GetService(ctx, "7fca84d6-7d37-4a74-a7b0-93e576089a41", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.Service != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `serviceID`                                                                        | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Service to lookup                                                        | 7fca84d6-7d37-4a74-a7b0-93e576089a41                                               |
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetServiceResponse](../../models/operations/getserviceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertService

Create or Update Service using ID or name.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.Services.UpsertService(ctx, operations.UpsertServiceRequest{
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Service: components.Service{
            Host: "example.internal",
            ID: sdkkonnectgo.String("49fd316e-c457-481c-9fc7-8079153e4f3c"),
            Name: sdkkonnectgo.String("example-service"),
            Path: sdkkonnectgo.String("/"),
            Port: sdkkonnectgo.Int64(80),
            Protocol: components.ProtocolHTTP.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Service != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpsertServiceRequest](../../models/operations/upsertservicerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.UpsertServiceResponse](../../models/operations/upsertserviceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |