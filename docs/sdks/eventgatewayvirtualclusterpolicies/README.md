# EventGatewayVirtualClusterPolicies

## Overview

Policies control how Kafka protocol traffic is modified between the client and the backend cluster.

Cluster policies are transformation and validation policies that can be applied to Kafka messages.


### Available Operations

* [ListEventGatewayVirtualClusterClusterLevelPolicies](#listeventgatewayvirtualclusterclusterlevelpolicies) - List Cluster Policies for Virtual Cluster
* [CreateEventGatewayVirtualClusterClusterLevelPolicy](#createeventgatewayvirtualclusterclusterlevelpolicy) - Create Cluster Policy for Virtual Cluster
* [GetEventGatewayVirtualClusterClusterLevelPolicy](#geteventgatewayvirtualclusterclusterlevelpolicy) - Get a Cluster Policy for Virtual Cluster
* [UpdateEventGatewayVirtualClusterClusterLevelPolicy](#updateeventgatewayvirtualclusterclusterlevelpolicy) - Update Cluster Policy for Virtual Cluster
* [PatchEventGatewayVirtualClusterClusterLevelPolicy](#patcheventgatewayvirtualclusterclusterlevelpolicy) - Patch Cluster Policy for Virtual Cluster
* [DeleteEventGatewayVirtualClusterClusterLevelPolicy](#deleteeventgatewayvirtualclusterclusterlevelpolicy) - Delete Cluster Policy for Virtual Cluster
* [MoveEventGatewayVirtualClusterClusterLevelPolicy](#moveeventgatewayvirtualclusterclusterlevelpolicy) - Move Cluster Policy
* [GetEventGatewayVirtualClusterClusterLevelPolicyChain](#geteventgatewayvirtualclusterclusterlevelpolicychain) - Get Cluster Policy Chain for Virtual Cluster
* [UpdateEventGatewayVirtualClusterClusterLevelPolicyChain](#updateeventgatewayvirtualclusterclusterlevelpolicychain) - Update Cluster Policy Chain

## ListEventGatewayVirtualClusterClusterLevelPolicies

Returns a list of cluster-level policies associated with the virtual cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-event-gateway-virtual-cluster-cluster-level-policies" method="get" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/cluster-policies" -->
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

    res, err := s.EventGatewayVirtualClusterPolicies.ListEventGatewayVirtualClusterClusterLevelPolicies(ctx, operations.ListEventGatewayVirtualClusterClusterLevelPoliciesRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "84528c6f-b191-41cb-aa9a-f0b6dcd04c2a",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListClusterPoliciesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.ListEventGatewayVirtualClusterClusterLevelPoliciesRequest](../../models/operations/listeventgatewayvirtualclusterclusterlevelpoliciesrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                     | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.ListEventGatewayVirtualClusterClusterLevelPoliciesResponse](../../models/operations/listeventgatewayvirtualclusterclusterlevelpoliciesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateEventGatewayVirtualClusterClusterLevelPolicy

Creates a new cluster-level policy associated with the specified Event Gateway virtual cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-event-gateway-virtual-cluster-cluster-level-policy" method="post" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/cluster-policies" -->
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

    res, err := s.EventGatewayVirtualClusterPolicies.CreateEventGatewayVirtualClusterClusterLevelPolicy(ctx, operations.CreateEventGatewayVirtualClusterClusterLevelPolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "3930057d-34a7-48a1-a438-70f200e65244",
        EventGatewayClusterPolicyModify: sdkkonnectgo.Pointer(components.CreateEventGatewayClusterPolicyModifyAcls(
            components.EventGatewayACLsPolicy{
                Config: components.EventGatewayACLPolicyConfig{
                    Rules: []components.EventGatewayACLRule{
                        components.EventGatewayACLRule{
                            ResourceType: components.ResourceTypeCluster,
                            Action: components.ActionAllow,
                            Operations: []components.EventGatewayACLOperation{
                                components.EventGatewayACLOperation{
                                    Name: components.NameRead,
                                },
                            },
                            ResourceNames: []components.EventGatewayACLResourceName{
                                components.EventGatewayACLResourceName{
                                    Match: "<value>",
                                },
                            },
                        },
                    },
                },
                Condition: sdkkonnectgo.Pointer("context.auth.principal.name == \"this-user\""),
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

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.CreateEventGatewayVirtualClusterClusterLevelPolicyRequest](../../models/operations/createeventgatewayvirtualclusterclusterlevelpolicyrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                     | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.CreateEventGatewayVirtualClusterClusterLevelPolicyResponse](../../models/operations/createeventgatewayvirtualclusterclusterlevelpolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetEventGatewayVirtualClusterClusterLevelPolicy

Returns information about a specific cluster-level policy associated with the Event Gateway virtual cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-event-gateway-virtual-cluster-cluster-level-policy" method="get" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/cluster-policies/{policyId}" -->
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

    res, err := s.EventGatewayVirtualClusterPolicies.GetEventGatewayVirtualClusterClusterLevelPolicy(ctx, operations.GetEventGatewayVirtualClusterClusterLevelPolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "48789a72-6de6-434f-b56c-a2ff67b409ff",
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

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.GetEventGatewayVirtualClusterClusterLevelPolicyRequest](../../models/operations/geteventgatewayvirtualclusterclusterlevelpolicyrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |
| `opts`                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                     | The options for this request.                                                                                                                          |

### Response

**[*operations.GetEventGatewayVirtualClusterClusterLevelPolicyResponse](../../models/operations/geteventgatewayvirtualclusterclusterlevelpolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateEventGatewayVirtualClusterClusterLevelPolicy

Updates an existing cluster-level policy associated with the specified Event Gateway virtual cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-event-gateway-virtual-cluster-cluster-level-policy" method="put" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/cluster-policies/{policyId}" -->
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

    res, err := s.EventGatewayVirtualClusterPolicies.UpdateEventGatewayVirtualClusterClusterLevelPolicy(ctx, operations.UpdateEventGatewayVirtualClusterClusterLevelPolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "5fe1fc3b-5d75-4327-a961-90830801fdbf",
        PolicyID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        EventGatewayClusterPolicyModify: sdkkonnectgo.Pointer(components.CreateEventGatewayClusterPolicyModifyAcls(
            components.EventGatewayACLsPolicy{
                Config: components.EventGatewayACLPolicyConfig{
                    Rules: []components.EventGatewayACLRule{},
                },
                Condition: sdkkonnectgo.Pointer("context.auth.principal.name == \"this-user\""),
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

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.UpdateEventGatewayVirtualClusterClusterLevelPolicyRequest](../../models/operations/updateeventgatewayvirtualclusterclusterlevelpolicyrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                     | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.UpdateEventGatewayVirtualClusterClusterLevelPolicyResponse](../../models/operations/updateeventgatewayvirtualclusterclusterlevelpolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchEventGatewayVirtualClusterClusterLevelPolicy

Partially updates an existing cluster-level policy associated with the specified Event Gateway virtual cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="patch-event-gateway-virtual-cluster-cluster-level-policy" method="patch" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/cluster-policies/{policyId}" -->
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

    res, err := s.EventGatewayVirtualClusterPolicies.PatchEventGatewayVirtualClusterClusterLevelPolicy(ctx, operations.PatchEventGatewayVirtualClusterClusterLevelPolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "ba703393-5387-44fa-af76-f575a0c8e5c3",
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

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.PatchEventGatewayVirtualClusterClusterLevelPolicyRequest](../../models/operations/patcheventgatewayvirtualclusterclusterlevelpolicyrequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |
| `opts`                                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                                   | :heavy_minus_sign:                                                                                                                                         | The options for this request.                                                                                                                              |

### Response

**[*operations.PatchEventGatewayVirtualClusterClusterLevelPolicyResponse](../../models/operations/patcheventgatewayvirtualclusterclusterlevelpolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteEventGatewayVirtualClusterClusterLevelPolicy

Deletes a specific cluster-level policy associated with the Event Gateway virtual cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-event-gateway-virtual-cluster-cluster-level-policy" method="delete" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/cluster-policies/{policyId}" -->
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

    res, err := s.EventGatewayVirtualClusterPolicies.DeleteEventGatewayVirtualClusterClusterLevelPolicy(ctx, operations.DeleteEventGatewayVirtualClusterClusterLevelPolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "153e0dee-68a5-46f8-b5dc-6b9a47f769b9",
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

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.DeleteEventGatewayVirtualClusterClusterLevelPolicyRequest](../../models/operations/deleteeventgatewayvirtualclusterclusterlevelpolicyrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                     | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.DeleteEventGatewayVirtualClusterClusterLevelPolicyResponse](../../models/operations/deleteeventgatewayvirtualclusterclusterlevelpolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## MoveEventGatewayVirtualClusterClusterLevelPolicy

Moves the position of a specific cluster-level policy relative to the chain associated with the Event Gateway virtual
cluster.
If a policy is defined under a parent policy, it moves the position relative to the sibling policies under the
same parent.


### Example Usage

<!-- UsageSnippet language="go" operationID="move-event-gateway-virtual-cluster-cluster-level-policy" method="post" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/cluster-policies/{policyId}/move" -->
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

    res, err := s.EventGatewayVirtualClusterPolicies.MoveEventGatewayVirtualClusterClusterLevelPolicy(ctx, operations.MoveEventGatewayVirtualClusterClusterLevelPolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "7a1b94f2-5d9e-4d3c-88ad-555625bb0c58",
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

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.MoveEventGatewayVirtualClusterClusterLevelPolicyRequest](../../models/operations/moveeventgatewayvirtualclusterclusterlevelpolicyrequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.MoveEventGatewayVirtualClusterClusterLevelPolicyResponse](../../models/operations/moveeventgatewayvirtualclusterclusterlevelpolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetEventGatewayVirtualClusterClusterLevelPolicyChain

Get the cluster-level policy chain for a virtual cluster composed of all the ids of the cluster-level policies in order of execution.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-event-gateway-virtual-cluster-cluster-level-policy-chain" method="get" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/cluster-policy-chain" -->
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

    res, err := s.EventGatewayVirtualClusterPolicies.GetEventGatewayVirtualClusterClusterLevelPolicyChain(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "cc7d1778-0cd3-4fa6-8469-55856fb6341b")
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

**[*operations.GetEventGatewayVirtualClusterClusterLevelPolicyChainResponse](../../models/operations/geteventgatewayvirtualclusterclusterlevelpolicychainresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateEventGatewayVirtualClusterClusterLevelPolicyChain

Update the cluster-level policy chain for a virtual cluster by providing an ordered list of cluster-level policy ids.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-event-gateway-virtual-cluster-cluster-level-policy-chain" method="put" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/cluster-policy-chain" -->
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

    res, err := s.EventGatewayVirtualClusterPolicies.UpdateEventGatewayVirtualClusterClusterLevelPolicyChain(ctx, operations.UpdateEventGatewayVirtualClusterClusterLevelPolicyChainRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "e5934101-6a09-4d93-a554-f825eaa5c927",
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

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |
| `request`                                                                                                                                                              | [operations.UpdateEventGatewayVirtualClusterClusterLevelPolicyChainRequest](../../models/operations/updateeventgatewayvirtualclusterclusterlevelpolicychainrequest.md) | :heavy_check_mark:                                                                                                                                                     | The request object to use for the request.                                                                                                                             |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |

### Response

**[*operations.UpdateEventGatewayVirtualClusterClusterLevelPolicyChainResponse](../../models/operations/updateeventgatewayvirtualclusterclusterlevelpolicychainresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |