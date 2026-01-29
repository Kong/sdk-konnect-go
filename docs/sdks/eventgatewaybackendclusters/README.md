# EventGatewayBackendClusters

## Overview

A backend cluster is an abstraction of a real Kafka cluster. It stores the connection and configuration details required for Kong Event Gateway to proxy traffic to Kafka.

Multiple Kafka clusters can be proxied through a single Kong Event Gateway.


### Available Operations

* [ListEventGatewayBackendClusters](#listeventgatewaybackendclusters) - List Backend Clusters
* [CreateEventGatewayBackendCluster](#createeventgatewaybackendcluster) - Create Backend Cluster
* [GetEventGatewayBackendCluster](#geteventgatewaybackendcluster) - Get a Backend Cluster
* [UpdateEventGatewayBackendCluster](#updateeventgatewaybackendcluster) - Update Backend Cluster
* [DeleteEventGatewayBackendCluster](#deleteeventgatewaybackendcluster) - Delete Backend Cluster

## ListEventGatewayBackendClusters

Returns a list of backend clusters associated with the specified Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-event-gateway-backend-clusters" method="get" path="/v1/event-gateways/{gatewayId}/backend-clusters" -->
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

    res, err := s.EventGatewayBackendClusters.ListEventGatewayBackendClusters(ctx, operations.ListEventGatewayBackendClustersRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListBackendClustersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.ListEventGatewayBackendClustersRequest](../../models/operations/listeventgatewaybackendclustersrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.ListEventGatewayBackendClustersResponse](../../models/operations/listeventgatewaybackendclustersresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateEventGatewayBackendCluster

Creates a new backend cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-event-gateway-backend-cluster" method="post" path="/v1/event-gateways/{gatewayId}/backend-clusters" -->
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

    res, err := s.EventGatewayBackendClusters.CreateEventGatewayBackendCluster(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", &components.CreateBackendClusterRequest{
        Name: "<value>",
        Authentication: components.CreateBackendClusterAuthenticationSchemeAnonymous(
            components.BackendClusterAuthenticationAnonymous{},
        ),
        BootstrapServers: []string{
            "<value 1>",
            "<value 2>",
        },
        TLS: components.BackendClusterTLS{
            Enabled: false,
            ClientIdentity: &components.ClientIdentity{
                Certificate: "<value>",
                Key: "${vault.env['MY_ENV_VAR']}",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BackendCluster != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       | Example                                                                                           |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |                                                                                                   |
| `gatewayID`                                                                                       | *string*                                                                                          | :heavy_check_mark:                                                                                | The UUID of your Gateway.                                                                         | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                              |
| `createBackendClusterRequest`                                                                     | [*components.CreateBackendClusterRequest](../../models/components/createbackendclusterrequest.md) | :heavy_minus_sign:                                                                                | The request schema for creating a backend cluster.                                                |                                                                                                   |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |                                                                                                   |

### Response

**[*operations.CreateEventGatewayBackendClusterResponse](../../models/operations/createeventgatewaybackendclusterresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetEventGatewayBackendCluster

Returns information about a specific backend cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-event-gateway-backend-cluster" method="get" path="/v1/event-gateways/{gatewayId}/backend-clusters/{backendClusterId}" -->
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

    res, err := s.EventGatewayBackendClusters.GetEventGatewayBackendCluster(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "5ef7f382-5bc5-4c2e-a215-300f045de3da")
    if err != nil {
        log.Fatal(err)
    }
    if res.BackendCluster != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | *string*                                                 | :heavy_check_mark:                                       | The UUID of your Gateway.                                | 9524ec7d-36d9-465d-a8c5-83a3c9390458                     |
| `backendClusterID`                                       | *string*                                                 | :heavy_check_mark:                                       | The ID of the Backend Cluster.                           |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetEventGatewayBackendClusterResponse](../../models/operations/geteventgatewaybackendclusterresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateEventGatewayBackendCluster

Updates an existing backend cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-event-gateway-backend-cluster" method="put" path="/v1/event-gateways/{gatewayId}/backend-clusters/{backendClusterId}" -->
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

    res, err := s.EventGatewayBackendClusters.UpdateEventGatewayBackendCluster(ctx, operations.UpdateEventGatewayBackendClusterRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        BackendClusterID: "daa3274b-d11d-4c55-a802-1603cfec6a6c",
        UpdateBackendClusterRequest: &components.UpdateBackendClusterRequest{
            Name: "<value>",
            Authentication: components.CreateBackendClusterAuthenticationSensitiveDataAwareSchemeSaslPlain(
                components.BackendClusterAuthenticationSaslPlainSensitiveDataAware{
                    Username: "Yasmin66",
                },
            ),
            BootstrapServers: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            TLS: components.BackendClusterTLS{
                Enabled: false,
                ClientIdentity: &components.ClientIdentity{
                    Certificate: "<value>",
                    Key: "${vault.env['MY_ENV_VAR']}",
                },
            },
            Labels: map[string]string{
                "env": "test",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BackendCluster != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.UpdateEventGatewayBackendClusterRequest](../../models/operations/updateeventgatewaybackendclusterrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.UpdateEventGatewayBackendClusterResponse](../../models/operations/updateeventgatewaybackendclusterresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteEventGatewayBackendCluster

Deletes a specific backend cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-event-gateway-backend-cluster" method="delete" path="/v1/event-gateways/{gatewayId}/backend-clusters/{backendClusterId}" -->
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

    res, err := s.EventGatewayBackendClusters.DeleteEventGatewayBackendCluster(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "61a9ef4a-7f93-412e-8e7a-e7fd047ffd14")
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
| `backendClusterID`                                       | *string*                                                 | :heavy_check_mark:                                       | The ID of the Backend Cluster.                           |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteEventGatewayBackendClusterResponse](../../models/operations/deleteeventgatewaybackendclusterresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |