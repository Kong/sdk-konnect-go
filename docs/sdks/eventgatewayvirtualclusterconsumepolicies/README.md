# EventGatewayVirtualClusterConsumePolicies

## Overview

Consume policies operate on Kafka messages as they are read from a Kafka cluster.

Transformations may be applied at consume time, but they are applied once per Consumer. Where possible, transofmrations should be applied as a Produce policy


### Available Operations

* [ListEventGatewayVirtualClusterConsumePolicies](#listeventgatewayvirtualclusterconsumepolicies) - List Consume Policies for Virtual Cluster
* [CreateEventGatewayVirtualClusterConsumePolicy](#createeventgatewayvirtualclusterconsumepolicy) - Create Consume Policy for Virtual Cluster
* [GetEventGatewayVirtualClusterConsumePolicy](#geteventgatewayvirtualclusterconsumepolicy) - Get a Consume Policy for Virtual Cluster
* [UpdateEventGatewayVirtualClusterConsumePolicy](#updateeventgatewayvirtualclusterconsumepolicy) - Update Consume Policy for Virtual Cluster
* [PatchEventGatewayVirtualClusterConsumePolicy](#patcheventgatewayvirtualclusterconsumepolicy) - Patch Consume Policy for Virtual Cluster
* [DeleteEventGatewayVirtualClusterConsumePolicy](#deleteeventgatewayvirtualclusterconsumepolicy) - Delete Consume Policy for Virtual Cluster
* [MoveEventGatewayVirtualClusterConsumePolicy](#moveeventgatewayvirtualclusterconsumepolicy) - Move Consume Policy
* [GetEventGatewayVirtualClusterConsumePolicyChain](#geteventgatewayvirtualclusterconsumepolicychain) - Get Consume Policy Chain
* [UpdateEventGatewayVirtualClusterConsumePolicyChain](#updateeventgatewayvirtualclusterconsumepolicychain) - Update Consume Policy Chain

## ListEventGatewayVirtualClusterConsumePolicies

Returns a list of consume policies associated with the specified Event Gateway virtual cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-event-gateway-virtual-cluster-consume-policies" method="get" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/consume-policies" -->
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

    res, err := s.EventGatewayVirtualClusterConsumePolicies.ListEventGatewayVirtualClusterConsumePolicies(ctx, operations.ListEventGatewayVirtualClusterConsumePoliciesRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "ff140e55-40f7-4276-b03f-7af85865dbb7",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListConsumePoliciesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.ListEventGatewayVirtualClusterConsumePoliciesRequest](../../models/operations/listeventgatewayvirtualclusterconsumepoliciesrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.ListEventGatewayVirtualClusterConsumePoliciesResponse](../../models/operations/listeventgatewayvirtualclusterconsumepoliciesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateEventGatewayVirtualClusterConsumePolicy

Creates a new consume policy associated with the specified Event Gateway virtual cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-event-gateway-virtual-cluster-consume-policy" method="post" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/consume-policies" -->
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

    res, err := s.EventGatewayVirtualClusterConsumePolicies.CreateEventGatewayVirtualClusterConsumePolicy(ctx, operations.CreateEventGatewayVirtualClusterConsumePolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "78114908-b42c-4c6b-a88b-3866f4ad4fb1",
        EventGatewayConsumePolicyCreate: sdkkonnectgo.Pointer(components.CreateEventGatewayConsumePolicyCreateDecrypt(
            components.EventGatewayDecryptPolicy{
                Config: components.EventGatewayDecryptPolicyConfig{
                    FailureMode: components.EncryptionFailureModeError,
                    KeySources: []components.EventGatewayKeySource{
                        components.CreateEventGatewayKeySourceStatic(
                            components.EventGatewayStaticKeySource{},
                        ),
                    },
                    PartOfRecord: []components.DecryptionRecordPart{},
                },
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
| `request`                                                                                                                                          | [operations.CreateEventGatewayVirtualClusterConsumePolicyRequest](../../models/operations/createeventgatewayvirtualclusterconsumepolicyrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.CreateEventGatewayVirtualClusterConsumePolicyResponse](../../models/operations/createeventgatewayvirtualclusterconsumepolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetEventGatewayVirtualClusterConsumePolicy

Returns information about a specific consume policy associated with the Event Gateway virtual cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-event-gateway-virtual-cluster-consume-policy" method="get" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/consume-policies/{policyId}" -->
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

    res, err := s.EventGatewayVirtualClusterConsumePolicies.GetEventGatewayVirtualClusterConsumePolicy(ctx, operations.GetEventGatewayVirtualClusterConsumePolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "5da4d9b4-64b5-408a-b061-92a9726c768a",
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
| `request`                                                                                                                                    | [operations.GetEventGatewayVirtualClusterConsumePolicyRequest](../../models/operations/geteventgatewayvirtualclusterconsumepolicyrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.GetEventGatewayVirtualClusterConsumePolicyResponse](../../models/operations/geteventgatewayvirtualclusterconsumepolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateEventGatewayVirtualClusterConsumePolicy

Updates an existing consume policy associated with the specified Event Gateway virtual cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-event-gateway-virtual-cluster-consume-policy" method="put" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/consume-policies/{policyId}" -->
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

    res, err := s.EventGatewayVirtualClusterConsumePolicies.UpdateEventGatewayVirtualClusterConsumePolicy(ctx, operations.UpdateEventGatewayVirtualClusterConsumePolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "e376f81d-2145-4c9e-8f23-d4b50c9e2f30",
        PolicyID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        EventGatewayConsumePolicyUpdate: sdkkonnectgo.Pointer(components.CreateEventGatewayConsumePolicyUpdateDecrypt(
            components.EventGatewayDecryptPolicy{
                Config: components.EventGatewayDecryptPolicyConfig{
                    FailureMode: components.EncryptionFailureModeError,
                    KeySources: []components.EventGatewayKeySource{},
                    PartOfRecord: []components.DecryptionRecordPart{},
                },
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
| `request`                                                                                                                                          | [operations.UpdateEventGatewayVirtualClusterConsumePolicyRequest](../../models/operations/updateeventgatewayvirtualclusterconsumepolicyrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.UpdateEventGatewayVirtualClusterConsumePolicyResponse](../../models/operations/updateeventgatewayvirtualclusterconsumepolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchEventGatewayVirtualClusterConsumePolicy

Updates an existing consume policy associated with the specified Event Gateway virtual cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="patch-event-gateway-virtual-cluster-consume-policy" method="patch" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/consume-policies/{policyId}" -->
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

    res, err := s.EventGatewayVirtualClusterConsumePolicies.PatchEventGatewayVirtualClusterConsumePolicy(ctx, operations.PatchEventGatewayVirtualClusterConsumePolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "8e94d8b9-530b-40fc-be08-8260acbd02cf",
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
| `request`                                                                                                                                        | [operations.PatchEventGatewayVirtualClusterConsumePolicyRequest](../../models/operations/patcheventgatewayvirtualclusterconsumepolicyrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |
| `opts`                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                         | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |

### Response

**[*operations.PatchEventGatewayVirtualClusterConsumePolicyResponse](../../models/operations/patcheventgatewayvirtualclusterconsumepolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteEventGatewayVirtualClusterConsumePolicy

Deletes a specific consume policy associated with the Event Gateway virtual cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-event-gateway-virtual-cluster-consume-policy" method="delete" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/consume-policies/{policyId}" -->
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

    res, err := s.EventGatewayVirtualClusterConsumePolicies.DeleteEventGatewayVirtualClusterConsumePolicy(ctx, operations.DeleteEventGatewayVirtualClusterConsumePolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "5008f907-b4b3-4d59-953c-cb782d7ef21d",
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
| `request`                                                                                                                                          | [operations.DeleteEventGatewayVirtualClusterConsumePolicyRequest](../../models/operations/deleteeventgatewayvirtualclusterconsumepolicyrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.DeleteEventGatewayVirtualClusterConsumePolicyResponse](../../models/operations/deleteeventgatewayvirtualclusterconsumepolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## MoveEventGatewayVirtualClusterConsumePolicy

Moves the position of a specific consume policy relative to the chain associated with the Event Gateway virtual
cluster.
If a policy is defined under a parent policy, it moves the position relative to the sibling policies under the
same parent.


### Example Usage

<!-- UsageSnippet language="go" operationID="move-event-gateway-virtual-cluster-consume-policy" method="post" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/consume-policies/{policyId}/move" -->
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

    res, err := s.EventGatewayVirtualClusterConsumePolicies.MoveEventGatewayVirtualClusterConsumePolicy(ctx, operations.MoveEventGatewayVirtualClusterConsumePolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "d24472cf-1d02-4226-b0b7-05bade0d004d",
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
| `request`                                                                                                                                      | [operations.MoveEventGatewayVirtualClusterConsumePolicyRequest](../../models/operations/moveeventgatewayvirtualclusterconsumepolicyrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.MoveEventGatewayVirtualClusterConsumePolicyResponse](../../models/operations/moveeventgatewayvirtualclusterconsumepolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetEventGatewayVirtualClusterConsumePolicyChain

Get the consume policy chain for a virtual cluster composed of all the ids of the consume policies in order of execution.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-event-gateway-virtual-cluster-consume-policy-chain" method="get" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/consume-policy-chain" -->
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

    res, err := s.EventGatewayVirtualClusterConsumePolicies.GetEventGatewayVirtualClusterConsumePolicyChain(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "d46c4526-76db-41e1-8ead-82c19cfcf3ee")
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

**[*operations.GetEventGatewayVirtualClusterConsumePolicyChainResponse](../../models/operations/geteventgatewayvirtualclusterconsumepolicychainresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateEventGatewayVirtualClusterConsumePolicyChain

Update the consume policy chain for a virtual cluster by providing an ordered list of consume policy ids.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-event-gateway-virtual-cluster-consume-policy-chain" method="put" path="/v1/event-gateways/{gatewayId}/virtual-clusters/{virtualClusterId}/consume-policy-chain" -->
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

    res, err := s.EventGatewayVirtualClusterConsumePolicies.UpdateEventGatewayVirtualClusterConsumePolicyChain(ctx, operations.UpdateEventGatewayVirtualClusterConsumePolicyChainRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VirtualClusterID: "bd59783d-ecbb-49ab-8ead-2aeabd75f81b",
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
| `request`                                                                                                                                                    | [operations.UpdateEventGatewayVirtualClusterConsumePolicyChainRequest](../../models/operations/updateeventgatewayvirtualclusterconsumepolicychainrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                     | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.UpdateEventGatewayVirtualClusterConsumePolicyChainResponse](../../models/operations/updateeventgatewayvirtualclusterconsumepolicychainresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |