# EventGatewayVirtualClusters
(*EventGatewayVirtualClusters*)

## Overview

Virtual clusters are the primary way clients interact with the Event Gateway proxy. They allow you to isolate clients from each other when connecting to the same backend cluster, and provide each client with modified view while still appearing as a standard Kafka cluster.


### Available Operations

* [ListEventGatewayVirtualClusters](#listeventgatewayvirtualclusters) - List all virtual clusters
* [CreateEventGatewayVirtualCluster](#createeventgatewayvirtualcluster) - Create Virtual Cluster
* [GetEventGatewayVirtualCluster](#geteventgatewayvirtualcluster) - Get a Virtual Cluster
* [UpdateEventGatewayVirtualCluster](#updateeventgatewayvirtualcluster) - Update Virtual Cluster
* [DeleteEventGatewayVirtualCluster](#deleteeventgatewayvirtualcluster) - Delete Virtual Cluster

## ListEventGatewayVirtualClusters

Returns a paginated list of virtual clusters associated with the specified Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-event-gateway-virtual-clusters" method="get" path="/v1/event-gateways/{gatewayId}/virtual-clusters" -->
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

    res, err := s.EventGatewayVirtualClusters.ListEventGatewayVirtualClusters(ctx, operations.ListEventGatewayVirtualClustersRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListVirtualClustersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.ListEventGatewayVirtualClustersRequest](../../models/operations/listeventgatewayvirtualclustersrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.ListEventGatewayVirtualClustersResponse](../../models/operations/listeventgatewayvirtualclustersresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateEventGatewayVirtualCluster

Creates a new virtual cluster associated with the specified Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-event-gateway-virtual-cluster" method="post" path="/v1/event-gateways/{gatewayId}/virtual-clusters" -->
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

    res, err := s.EventGatewayVirtualClusters.CreateEventGatewayVirtualCluster(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", &components.CreateVirtualClusterRequest{
        Name: "<value>",
        Destination: components.CreateBackendClusterReferenceModifyBackendClusterReferenceByName(
            components.BackendClusterReferenceByName{
                Name: "<value>",
            },
        ),
        Authentication: []components.VirtualClusterAuthenticationScheme{},
        ACLMode: components.VirtualClusterACLModeEnforceOnGateway,
        DNSLabel: "vcluster-1",
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualCluster != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       | Example                                                                                           |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |                                                                                                   |
| `gatewayID`                                                                                       | *string*                                                                                          | :heavy_check_mark:                                                                                | The UUID of your Gateway.                                                                         | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                              |
| `createVirtualClusterRequest`                                                                     | [*components.CreateVirtualClusterRequest](../../models/components/createvirtualclusterrequest.md) | :heavy_minus_sign:                                                                                | The request schema for creating a virtual cluster.                                                |                                                                                                   |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |                                                                                                   |

### Response

**[*operations.CreateEventGatewayVirtualClusterResponse](../../models/operations/createeventgatewayvirtualclusterresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetEventGatewayVirtualCluster

Returns information about a specific virtual cluster associated with the Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-event-gateway-virtual-cluster" method="get" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}" -->
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

    res, err := s.EventGatewayVirtualClusters.GetEventGatewayVirtualCluster(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "0b2723a8-8eff-4698-8224-0c823cd13fbf")
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualCluster != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | *string*                                                 | :heavy_check_mark:                                       | The UUID of your Gateway.                                | 9524ec7d-36d9-465d-a8c5-83a3c9390458                     |
| `virtualClusterID`                                       | *string*                                                 | :heavy_check_mark:                                       | The ID of the Virtual Cluster.                           |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetEventGatewayVirtualClusterResponse](../../models/operations/geteventgatewayvirtualclusterresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateEventGatewayVirtualCluster

Updates an existing virtual cluster associated with the specified Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-event-gateway-virtual-cluster" method="put" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}" -->
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

    res, err := s.EventGatewayVirtualClusters.UpdateEventGatewayVirtualCluster(ctx, operations.UpdateEventGatewayVirtualClusterRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "612098cd-49b2-41cd-8c09-b16296a1322a",
        UpdateVirtualClusterRequest: &components.UpdateVirtualClusterRequest{
            Name: "<value>",
            Destination: components.CreateBackendClusterReferenceModifyBackendClusterReferenceByID(
                components.BackendClusterReferenceByID{
                    ID: "fe909fd6-277e-424b-b080-aba0bbb79170",
                },
            ),
            Authentication: []components.VirtualClusterAuthenticationSensitiveDataAwareScheme{
                components.CreateVirtualClusterAuthenticationSensitiveDataAwareSchemeSaslPlain(
                    components.VirtualClusterAuthenticationSaslPlainSensitiveDataAware{
                        Mediation: components.MediationPassthrough,
                    },
                ),
            },
            ACLMode: components.VirtualClusterACLModePassthrough,
            DNSLabel: "vcluster-1",
            Labels: map[string]string{
                "env": "test",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualCluster != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.UpdateEventGatewayVirtualClusterRequest](../../models/operations/updateeventgatewayvirtualclusterrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.UpdateEventGatewayVirtualClusterResponse](../../models/operations/updateeventgatewayvirtualclusterresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteEventGatewayVirtualCluster

Deletes a specific virtual cluster associated with the Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-event-gateway-virtual-cluster" method="delete" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}" -->
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

    res, err := s.EventGatewayVirtualClusters.DeleteEventGatewayVirtualCluster(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "bc89fca5-c6f5-4572-8d11-4848dfbb66ad")
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
| `virtualClusterID`                                       | *string*                                                 | :heavy_check_mark:                                       | The ID of the Virtual Cluster.                           |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteEventGatewayVirtualClusterResponse](../../models/operations/deleteeventgatewayvirtualclusterresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |