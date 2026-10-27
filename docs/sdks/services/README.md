# Services

## Overview

Service entities are abstractions of your microservice interfaces or formal APIs. For example, a service could be a data transformation microservice or a billing API.
<br><br>
The main attribute of a service is the destination URL for proxying traffic. This URL can be set as a single string or by specifying its protocol, host, port and path individually.
<br><br>
Services are associated to routes, and a single service can have many routes associated with it. Routes are entrypoints in Kong Gateway which define rules to match client requests. Once a route is matched, Kong Gateway proxies the request to its associated service. See the [Route documentation](https://developer.konghq.com/gateway/entities/route/) for a detailed explanation of how Kong proxies traffic.
<br><br>
Services can be both [tagged and filtered by tags](https://developer.konghq.com/admin-api/).


### Available Operations

* [ListService](#listservice) - List all Services in a workspace
* [CreateService](#createservice) - Create a new Service in a workspace
* [DeleteService](#deleteservice) - Delete a Service in a workspace
* [GetService](#getservice) - Get a Service in a workspace
* [UpsertService](#upsertservice) - Upsert a Service in a workspace

## ListService

List all Services in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-service" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/services" -->
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

    res, err := s.Services.ListService(ctx, operations.ListServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Tags: sdkkonnectgo.Pointer("tag1,tag2"),
        Workspace: "team-payments",
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

Create a new Service in a workspace

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="create-service" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/services" example="DuplicateApiKey" -->
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

    res, err := s.Services.CreateService(ctx, operations.CreateServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        ServiceRequest: components.CreateServiceRequestServiceWithHost(
            components.ServiceWithHost{
                Host: "example.internal",
                ID: sdkkonnectgo.Pointer("49fd316e-c457-481c-9fc7-8079153e4f3c"),
                Name: sdkkonnectgo.Pointer("example-service"),
                Path: sdkkonnectgo.Pointer("/"),
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Service != nil {
        // handle response
    }
}
```
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="create-service" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/services" example="InvalidAuthCred" -->
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

    res, err := s.Services.CreateService(ctx, operations.CreateServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        ServiceRequest: components.CreateServiceRequestServiceWithHost(
            components.ServiceWithHost{
                Host: "example.internal",
                ID: sdkkonnectgo.Pointer("49fd316e-c457-481c-9fc7-8079153e4f3c"),
                Name: sdkkonnectgo.Pointer("example-service"),
                Path: sdkkonnectgo.Pointer("/"),
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Service != nil {
        // handle response
    }
}
```
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="create-service" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/services" example="NoAPIKey" -->
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

    res, err := s.Services.CreateService(ctx, operations.CreateServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        ServiceRequest: components.CreateServiceRequestServiceWithHost(
            components.ServiceWithHost{
                Host: "example.internal",
                ID: sdkkonnectgo.Pointer("49fd316e-c457-481c-9fc7-8079153e4f3c"),
                Name: sdkkonnectgo.Pointer("example-service"),
                Path: sdkkonnectgo.Pointer("/"),
            },
        ),
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
| `request`                                                                          | [operations.CreateServiceRequest](../../models/operations/createservicerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.CreateServiceResponse](../../models/operations/createserviceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteService

Delete a Service in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-service" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/services/{ServiceId}" -->
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

    res, err := s.Services.DeleteService(ctx, operations.DeleteServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteServiceRequest](../../models/operations/deleteservicerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.DeleteServiceResponse](../../models/operations/deleteserviceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetService

Get a Service using ID or name in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-service" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/services/{ServiceId}" -->
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

    res, err := s.Services.GetService(ctx, operations.GetServiceRequest{
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetServiceRequest](../../models/operations/getservicerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.GetServiceResponse](../../models/operations/getserviceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertService

Create or Update Service using ID or name in a workspace.

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="upsert-service" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/services/{ServiceId}" example="DuplicateApiKey" -->
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

    res, err := s.Services.UpsertService(ctx, operations.UpsertServiceRequest{
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        ServiceRequest: components.CreateServiceRequestServiceWithHost(
            components.ServiceWithHost{
                Host: "example.internal",
                ID: sdkkonnectgo.Pointer("49fd316e-c457-481c-9fc7-8079153e4f3c"),
                Name: sdkkonnectgo.Pointer("example-service"),
                Path: sdkkonnectgo.Pointer("/"),
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Service != nil {
        // handle response
    }
}
```
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="upsert-service" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/services/{ServiceId}" example="InvalidAuthCred" -->
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

    res, err := s.Services.UpsertService(ctx, operations.UpsertServiceRequest{
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        ServiceRequest: components.CreateServiceRequestServiceWithHost(
            components.ServiceWithHost{
                Host: "example.internal",
                ID: sdkkonnectgo.Pointer("49fd316e-c457-481c-9fc7-8079153e4f3c"),
                Name: sdkkonnectgo.Pointer("example-service"),
                Path: sdkkonnectgo.Pointer("/"),
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Service != nil {
        // handle response
    }
}
```
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="upsert-service" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/services/{ServiceId}" example="NoAPIKey" -->
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

    res, err := s.Services.UpsertService(ctx, operations.UpsertServiceRequest{
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        ServiceRequest: components.CreateServiceRequestServiceWithHost(
            components.ServiceWithHost{
                Host: "example.internal",
                ID: sdkkonnectgo.Pointer("49fd316e-c457-481c-9fc7-8079153e4f3c"),
                Name: sdkkonnectgo.Pointer("example-service"),
                Path: sdkkonnectgo.Pointer("/"),
            },
        ),
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