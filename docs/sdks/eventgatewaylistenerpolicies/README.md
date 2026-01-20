# EventGatewayListenerPolicies

## Overview

Policies control how Kafka protocol traffic is modified between the client and the backend cluster.

Listener policies are routing policies that pass traffic to the virtual cluster.


### Available Operations

* [ListEventGatewayListenerPolicies](#listeventgatewaylistenerpolicies) - List Policies for Listener
* [CreateEventGatewayListenerPolicy](#createeventgatewaylistenerpolicy) - Create Policy for Listener
* [GetEventGatewayListenerPolicy](#geteventgatewaylistenerpolicy) - Get a Policy for Listener
* [UpdateEventGatewayListenerPolicy](#updateeventgatewaylistenerpolicy) - Update Policy for Listener
* [PatchEventGatewayListenerPolicy](#patcheventgatewaylistenerpolicy) - Partially Update Policy for Listener
* [DeleteEventGatewayListenerPolicy](#deleteeventgatewaylistenerpolicy) - Delete Policy for Listener
* [MoveEventGatewayListenerPolicy](#moveeventgatewaylistenerpolicy) - Move Policy
* [GetEventGatewayListenerPolicyChain](#geteventgatewaylistenerpolicychain) - Get Policy Chain for Listener
* [UpdateEventGatewayListenerPolicyChain](#updateeventgatewaylistenerpolicychain) - Update Policy Chain for Listener

## ListEventGatewayListenerPolicies

Returns a list of policies associated with the specified Event Gateway listener.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-event-gateway-listener-policies" method="get" path="/v1/event-gateways/{gatewayId}/listeners/{eventGatewayListenerId}/policies" -->
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

    res, err := s.EventGatewayListenerPolicies.ListEventGatewayListenerPolicies(ctx, operations.ListEventGatewayListenerPoliciesRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        EventGatewayListenerID: "6c709b1e-c2a0-4793-9e98-d1d902fe8b6f",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListEventGatewayListenerPoliciesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.ListEventGatewayListenerPoliciesRequest](../../models/operations/listeventgatewaylistenerpoliciesrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.ListEventGatewayListenerPoliciesResponse](../../models/operations/listeventgatewaylistenerpoliciesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateEventGatewayListenerPolicy

Creates a new policy associated with the specified Event Gateway listener.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-event-gateway-listener-policy" method="post" path="/v1/event-gateways/{gatewayId}/listeners/{eventGatewayListenerId}/policies" -->
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

    res, err := s.EventGatewayListenerPolicies.CreateEventGatewayListenerPolicy(ctx, operations.CreateEventGatewayListenerPolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        EventGatewayListenerID: "3407fe24-3af9-43d2-92df-f7396e794e62",
        EventGatewayListenerPolicyCreate: components.CreateEventGatewayListenerPolicyCreateTLSServer(
            components.EventGatewayTLSListenerPolicy{
                Config: components.EventGatewayTLSListenerPolicyConfig{
                    Certificates: []components.TLSCertificate{},
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayListenerPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.CreateEventGatewayListenerPolicyRequest](../../models/operations/createeventgatewaylistenerpolicyrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.CreateEventGatewayListenerPolicyResponse](../../models/operations/createeventgatewaylistenerpolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetEventGatewayListenerPolicy

Returns information about a specific policy associated with the Event Gateway listener.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-event-gateway-listener-policy" method="get" path="/v1/event-gateways/{gatewayId}/listeners/{eventGatewayListenerId}/policies/{policyId}" -->
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

    res, err := s.EventGatewayListenerPolicies.GetEventGatewayListenerPolicy(ctx, operations.GetEventGatewayListenerPolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        EventGatewayListenerID: "028b6c72-e081-4f74-a7c0-60c1911a2615",
        PolicyID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayListenerPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetEventGatewayListenerPolicyRequest](../../models/operations/geteventgatewaylistenerpolicyrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.GetEventGatewayListenerPolicyResponse](../../models/operations/geteventgatewaylistenerpolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateEventGatewayListenerPolicy

Updates an existing policy associated with the specified Event Gateway listener.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-event-gateway-listener-policy" method="put" path="/v1/event-gateways/{gatewayId}/listeners/{eventGatewayListenerId}/policies/{policyId}" -->
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

    res, err := s.EventGatewayListenerPolicies.UpdateEventGatewayListenerPolicy(ctx, operations.UpdateEventGatewayListenerPolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        EventGatewayListenerID: "ba676a00-d85e-467e-b0b8-728645f766fb",
        PolicyID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        EventGatewayListenerPolicyUpdate: components.CreateEventGatewayListenerPolicyUpdateForwardToVirtualCluster(
            components.ForwardToVirtualClusterPolicy{
                Config: components.CreateForwardToVirtualClusterPolicyConfigPortMapping(
                    components.ForwardToClusterByPortMappingConfig{
                        Destination: components.CreateVirtualClusterReferenceVirtualClusterReferenceByName(
                            components.VirtualClusterReferenceByName{
                                Name: "<value>",
                            },
                        ),
                        AdvertisedHost: "<value>",
                    },
                ),
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayListenerPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.UpdateEventGatewayListenerPolicyRequest](../../models/operations/updateeventgatewaylistenerpolicyrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.UpdateEventGatewayListenerPolicyResponse](../../models/operations/updateeventgatewaylistenerpolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchEventGatewayListenerPolicy

Partially updates an existing policy associated with the specified Event Gateway listener.

### Example Usage

<!-- UsageSnippet language="go" operationID="patch-event-gateway-listener-policy" method="patch" path="/v1/event-gateways/{gatewayId}/listeners/{eventGatewayListenerId}/policies/{policyId}" -->
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

    res, err := s.EventGatewayListenerPolicies.PatchEventGatewayListenerPolicy(ctx, operations.PatchEventGatewayListenerPolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        EventGatewayListenerID: "d34c0b72-33a5-4a32-8aba-d9c0867897f8",
        PolicyID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        EventGatewayListenerPolicyPatch: &components.EventGatewayListenerPolicyPatch{
            Labels: map[string]string{
                "env": "test",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayListenerPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.PatchEventGatewayListenerPolicyRequest](../../models/operations/patcheventgatewaylistenerpolicyrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.PatchEventGatewayListenerPolicyResponse](../../models/operations/patcheventgatewaylistenerpolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteEventGatewayListenerPolicy

Deletes a specific policy associated with the Event Gateway listener.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-event-gateway-listener-policy" method="delete" path="/v1/event-gateways/{gatewayId}/listeners/{eventGatewayListenerId}/policies/{policyId}" -->
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

    res, err := s.EventGatewayListenerPolicies.DeleteEventGatewayListenerPolicy(ctx, operations.DeleteEventGatewayListenerPolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        EventGatewayListenerID: "bba98afd-6e26-46d2-b5c2-ae9d8d47cecc",
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.DeleteEventGatewayListenerPolicyRequest](../../models/operations/deleteeventgatewaylistenerpolicyrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.DeleteEventGatewayListenerPolicyResponse](../../models/operations/deleteeventgatewaylistenerpolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## MoveEventGatewayListenerPolicy

Moves the position of a specific policy relative to the chain associated with the Event Gateway listener.


### Example Usage

<!-- UsageSnippet language="go" operationID="move-event-gateway-listener-policy" method="post" path="/v1/event-gateways/{gatewayId}/listeners/{eventGatewayListenerId}/policies/{policyId}/move" -->
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

    res, err := s.EventGatewayListenerPolicies.MoveEventGatewayListenerPolicy(ctx, operations.MoveEventGatewayListenerPolicyRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        EventGatewayListenerID: "49affc92-6aff-4fa8-ab92-e31bc124708f",
        PolicyID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        MoveEventGatewayPolicy: components.MoveEventGatewayPolicy{
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.MoveEventGatewayListenerPolicyRequest](../../models/operations/moveeventgatewaylistenerpolicyrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.MoveEventGatewayListenerPolicyResponse](../../models/operations/moveeventgatewaylistenerpolicyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetEventGatewayListenerPolicyChain

Get the policy chain for a listener composed of all the ids of the policies in order of execution.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-event-gateway-listener-policy-chain" method="get" path="/v1/event-gateways/{gatewayId}/listeners/{eventGatewayListenerId}/policy-chain" -->
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

    res, err := s.EventGatewayListenerPolicies.GetEventGatewayListenerPolicyChain(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "26c421b6-e915-437b-8975-05f5925a07a5")
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
| `eventGatewayListenerID`                                 | *string*                                                 | :heavy_check_mark:                                       | The ID of the Event Gateway Listener.                    |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetEventGatewayListenerPolicyChainResponse](../../models/operations/geteventgatewaylistenerpolicychainresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateEventGatewayListenerPolicyChain

Update the policy chain for a listener by providing an ordered list of policy ids.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-event-gateway-listener-policy-chain" method="put" path="/v1/event-gateways/{gatewayId}/listeners/{eventGatewayListenerId}/policy-chain" -->
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

    res, err := s.EventGatewayListenerPolicies.UpdateEventGatewayListenerPolicyChain(ctx, operations.UpdateEventGatewayListenerPolicyChainRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        EventGatewayListenerID: "fe9f5f21-fcaa-45c3-ae26-22ad7df78b60",
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

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.UpdateEventGatewayListenerPolicyChainRequest](../../models/operations/updateeventgatewaylistenerpolicychainrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.UpdateEventGatewayListenerPolicyChainResponse](../../models/operations/updateeventgatewaylistenerpolicychainresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |