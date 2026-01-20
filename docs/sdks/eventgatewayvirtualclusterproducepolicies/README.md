# EventGatewayVirtualClusterProducePolicies

## Overview

Produce policies operate on Kafka messages before they are written to the Kafka cluster.

Where possible, apply transformations to the data using produce policies rather than consume policies for maximum efficiency.


### Available Operations

* [ListEventGatewayVirtualClusterProducePolicies](#listeventgatewayvirtualclusterproducepolicies) - List Produce Policies for Virtual Cluster
* [CreateEventGatewayVirtualClusterProducePolicy](#createeventgatewayvirtualclusterproducepolicy) - Create Produce Policy for Virtual Cluster
* [GetEventGatewayVirtualClusterProducePolicy](#geteventgatewayvirtualclusterproducepolicy) - Get a Produce Policy for Virtual Cluster
* [UpdateEventGatewayVirtualClusterProducePolicy](#updateeventgatewayvirtualclusterproducepolicy) - Update Produce Policy for Virtual Cluster
* [PatchEventGatewayVirtualClusterProducePolicy](#patcheventgatewayvirtualclusterproducepolicy) - Patch Produce Policy for Virtual Cluster
* [DeleteEventGatewayVirtualClusterProducePolicy](#deleteeventgatewayvirtualclusterproducepolicy) - Delete Produce Policy for Virtual Cluster
* [MoveEventGatewayVirtualClusterProducePolicy](#moveeventgatewayvirtualclusterproducepolicy) - Move Produce Policy
* [GetEventGatewayVirtualClusterProducePolicyChain](#geteventgatewayvirtualclusterproducepolicychain) - Get Produce Policy Chain for Virtual Cluster
* [UpdateEventGatewayVirtualClusterProducePolicyChain](#updateeventgatewayvirtualclusterproducepolicychain) - Update Produce Policy Chain

## ListEventGatewayVirtualClusterProducePolicies

Returns a list of produce policies associated with the specified Event Gateway virtual cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-event-gateway-virtual-cluster-produce-policies" method="get" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/produce-policies" -->
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

    res, err := s.EventGatewayVirtualClusterProducePolicies.ListEventGatewayVirtualClusterProducePolicies(ctx, operations.ListEventGatewayVirtualClusterProducePoliciesRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "bdca6ea3-78bf-47d5-b6cb-0141a2dfa3ed",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListProducePoliciesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.ListEventGatewayVirtualClusterProducePoliciesRequest](../../models/operations/listeventgatewayvirtualclusterproducepoliciesrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.ListEventGatewayVirtualClusterProducePoliciesResponse](../../models/operations/listeventgatewayvirtualclusterproducepoliciesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateEventGatewayVirtualClusterProducePolicy

Creates a new produce policy associated with the specified Event Gateway virtual cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-event-gateway-virtual-cluster-produce-policy" method="post" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/produce-policies" -->
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

    res, err := s.EventGatewayVirtualClusterProducePolicies.CreateEventGatewayVirtualClusterProducePolicy(ctx, operations.CreateEventGatewayVirtualClusterProducePolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "cb5c6cc6-ed6d-4ae5-8b28-fd349b94ea6e",
        EventGatewayProducePolicyCreate: sdkkonnectgo.Pointer(components.CreateEventGatewayProducePolicyCreateSchemaValidation(
            components.EventGatewayProduceSchemaValidationPolicy{
                Config: components.CreateEventGatewayProduceSchemaValidationPolicyConfigJSON(
                    components.EventGatewayProduceSchemaValidationPolicyJSONConfig{},
                ),
                Condition: sdkkonnectgo.Pointer("context.topic.name.endsWith(\"my_suffix\") && records.headers[\"x-flag\"] == \"a-value\""),
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.CreateEventGatewayVirtualClusterProducePolicyRequest](../../models/operations/createeventgatewayvirtualclusterproducepolicyrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.CreateEventGatewayVirtualClusterProducePolicyResponse](../../models/operations/createeventgatewayvirtualclusterproducepolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetEventGatewayVirtualClusterProducePolicy

Returns information about a specific produce policy associated with the Event Gateway virtual cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-event-gateway-virtual-cluster-produce-policy" method="get" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/produce-policies/{policyId}" -->
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

    res, err := s.EventGatewayVirtualClusterProducePolicies.GetEventGatewayVirtualClusterProducePolicy(ctx, operations.GetEventGatewayVirtualClusterProducePolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "12000971-86dc-4d28-8c98-cf9085de2c25",
        PolicyID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.GetEventGatewayVirtualClusterProducePolicyRequest](../../models/operations/geteventgatewayvirtualclusterproducepolicyrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.GetEventGatewayVirtualClusterProducePolicyResponse](../../models/operations/geteventgatewayvirtualclusterproducepolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateEventGatewayVirtualClusterProducePolicy

Updates an existing produce policy associated with the specified Event Gateway virtual cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-event-gateway-virtual-cluster-produce-policy" method="put" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/produce-policies/{policyId}" -->
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

    res, err := s.EventGatewayVirtualClusterProducePolicies.UpdateEventGatewayVirtualClusterProducePolicy(ctx, operations.UpdateEventGatewayVirtualClusterProducePolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "1fbc34c2-c5b0-464f-a084-203480bf8d4c",
        PolicyID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        EventGatewayProducePolicyUpdate: sdkkonnectgo.Pointer(components.CreateEventGatewayProducePolicyUpdateSchemaValidation(
            components.EventGatewayProduceSchemaValidationPolicy{
                Config: components.CreateEventGatewayProduceSchemaValidationPolicyConfigConfluentSchemaRegistry(
                    components.EventGatewayProduceSchemaValidationPolicySchemaRegistryConfig{},
                ),
                Condition: sdkkonnectgo.Pointer("context.topic.name.endsWith(\"my_suffix\") && records.headers[\"x-flag\"] == \"a-value\""),
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.UpdateEventGatewayVirtualClusterProducePolicyRequest](../../models/operations/updateeventgatewayvirtualclusterproducepolicyrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.UpdateEventGatewayVirtualClusterProducePolicyResponse](../../models/operations/updateeventgatewayvirtualclusterproducepolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchEventGatewayVirtualClusterProducePolicy

Partially updates an existing produce policy associated with the specified Event Gateway virtual cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="patch-event-gateway-virtual-cluster-produce-policy" method="patch" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/produce-policies/{policyId}" -->
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

    res, err := s.EventGatewayVirtualClusterProducePolicies.PatchEventGatewayVirtualClusterProducePolicy(ctx, operations.PatchEventGatewayVirtualClusterProducePolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "5a14b7d8-84d2-42fe-ae35-6ce8a06f4033",
        PolicyID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        EventGatewayPolicyPatch: &components.EventGatewayPolicyPatch{
            Condition: sdkkonnectgo.Pointer("context.topic.name.endsWith('my_suffix')"),
            Labels: map[string]string{
                "env": "test",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.PatchEventGatewayVirtualClusterProducePolicyRequest](../../models/operations/patcheventgatewayvirtualclusterproducepolicyrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |
| `opts`                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                         | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |

### Response

**[*operations.PatchEventGatewayVirtualClusterProducePolicyResponse](../../models/operations/patcheventgatewayvirtualclusterproducepolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteEventGatewayVirtualClusterProducePolicy

Deletes a specific produce policy associated with the Event Gateway virtual cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-event-gateway-virtual-cluster-produce-policy" method="delete" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/produce-policies/{policyId}" -->
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

    res, err := s.EventGatewayVirtualClusterProducePolicies.DeleteEventGatewayVirtualClusterProducePolicy(ctx, operations.DeleteEventGatewayVirtualClusterProducePolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "b11f9cef-8525-4040-b5eb-48289887c507",
        PolicyID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
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

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.DeleteEventGatewayVirtualClusterProducePolicyRequest](../../models/operations/deleteeventgatewayvirtualclusterproducepolicyrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.DeleteEventGatewayVirtualClusterProducePolicyResponse](../../models/operations/deleteeventgatewayvirtualclusterproducepolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## MoveEventGatewayVirtualClusterProducePolicy

Moves the position of a specific produce policy relative to the chain associated with the Event Gateway virtual
cluster.
If a policy is defined under a parent policy, it moves the position relative to the sibling policies under the
same parent.


### Example Usage

<!-- UsageSnippet language="go" operationID="move-event-gateway-virtual-cluster-produce-policy" method="post" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/produce-policies/{policyId}/move" -->
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

    res, err := s.EventGatewayVirtualClusterProducePolicies.MoveEventGatewayVirtualClusterProducePolicy(ctx, operations.MoveEventGatewayVirtualClusterProducePolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "e5c38528-657f-4bb2-9e6c-54c0180c223b",
        PolicyID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        MoveEventGatewayPolicy: &components.MoveEventGatewayPolicy{
            Index: 2,
        },
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

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.MoveEventGatewayVirtualClusterProducePolicyRequest](../../models/operations/moveeventgatewayvirtualclusterproducepolicyrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.MoveEventGatewayVirtualClusterProducePolicyResponse](../../models/operations/moveeventgatewayvirtualclusterproducepolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetEventGatewayVirtualClusterProducePolicyChain

Get the produce policy chain for a virtual cluster composed of all the ids of the produce policies in order of execution.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-event-gateway-virtual-cluster-produce-policy-chain" method="get" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/produce-policy-chain" -->
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

    res, err := s.EventGatewayVirtualClusterProducePolicies.GetEventGatewayVirtualClusterProducePolicyChain(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "dd270061-c47e-401d-b520-af9afadccafe")
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayPolicyChainResponse != nil {
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

**[*operations.GetEventGatewayVirtualClusterProducePolicyChainResponse](../../models/operations/geteventgatewayvirtualclusterproducepolicychainresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateEventGatewayVirtualClusterProducePolicyChain

Update the produce policy chain for a virtual cluster by providing an ordered list of produce policy ids.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-event-gateway-virtual-cluster-produce-policy-chain" method="put" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/produce-policy-chain" -->
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

    res, err := s.EventGatewayVirtualClusterProducePolicies.UpdateEventGatewayVirtualClusterProducePolicyChain(ctx, operations.UpdateEventGatewayVirtualClusterProducePolicyChainRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "4fe0289b-22a8-4a1f-b215-3028b02c82fd",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayPolicyChainResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.UpdateEventGatewayVirtualClusterProducePolicyChainRequest](../../models/operations/updateeventgatewayvirtualclusterproducepolicychainrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                     | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.UpdateEventGatewayVirtualClusterProducePolicyChainResponse](../../models/operations/updateeventgatewayvirtualclusterproducepolicychainresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |