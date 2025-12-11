# EventGatewayListeners

## Overview

A listener represents hostname-port or IP-port combinations that connect to TCP sockets. Listeners need at least as many ports as backend brokers if you use port mapping in a Forward to Virtual Cluster policy. For SNI routing, you can route all brokers using a listener with only one port. Ports can be expressed as a single port or range. Addresses can be IPv4, IPv6, or hostnames.

A listener can have policies that enforce TLS certificates and perform SNI routing. The listener runs at Layer 4 of the network stack. In Kong Event Gateway, listeners first take in the connection and then route the TCP connection to a virtual cluster based on conditions defined in listener policies.


### Available Operations

* [ListEventGatewayListeners](#listeventgatewaylisteners) - List Event Gateway Listeners
* [CreateEventGatewayListener](#createeventgatewaylistener) - Create Event Gateway Listener
* [GetEventGatewayListener](#geteventgatewaylistener) - Get an Event Gateway Listener
* [UpdateEventGatewayListener](#updateeventgatewaylistener) - Update Event Gateway Listener
* [DeleteEventGatewayListener](#deleteeventgatewaylistener) - Delete Event Gateway Listener

## ListEventGatewayListeners

Returns a list of listeners associated with the specified Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-event-gateway-listeners" method="get" path="/v1/event-gateways/{gatewayId}/listeners" -->
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

    res, err := s.EventGatewayListeners.ListEventGatewayListeners(ctx, operations.ListEventGatewayListenersRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListEventGatewayListenersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListEventGatewayListenersRequest](../../models/operations/listeventgatewaylistenersrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ListEventGatewayListenersResponse](../../models/operations/listeventgatewaylistenersresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateEventGatewayListener

Creates a new listener associated with the specified Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-event-gateway-listener" method="post" path="/v1/event-gateways/{gatewayId}/listeners" -->
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

    res, err := s.EventGatewayListeners.CreateEventGatewayListener(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", &components.CreateEventGatewayListenerRequest{
        Name: "<value>",
        Addresses: []string{
            "<value 1>",
        },
        Ports: []components.EventGatewayListenerPort{},
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayListener != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                     | Type                                                                                                          | Required                                                                                                      | Description                                                                                                   | Example                                                                                                       |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                         | :heavy_check_mark:                                                                                            | The context to use for the request.                                                                           |                                                                                                               |
| `gatewayID`                                                                                                   | *string*                                                                                                      | :heavy_check_mark:                                                                                            | The UUID of your Gateway.                                                                                     | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                                          |
| `createEventGatewayListenerRequest`                                                                           | [*components.CreateEventGatewayListenerRequest](../../models/components/createeventgatewaylistenerrequest.md) | :heavy_minus_sign:                                                                                            | The request schema for creating a listener.                                                                   |                                                                                                               |
| `opts`                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                      | :heavy_minus_sign:                                                                                            | The options for this request.                                                                                 |                                                                                                               |

### Response

**[*operations.CreateEventGatewayListenerResponse](../../models/operations/createeventgatewaylistenerresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetEventGatewayListener

Returns information about a specific listener associated with the Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-event-gateway-listener" method="get" path="/v1/event-gateways/{gatewayId}/listeners/{eventGatewayListenerId}" -->
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

    res, err := s.EventGatewayListeners.GetEventGatewayListener(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "31f35308-f007-42b7-9c56-aadf0a2e6fee")
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayListener != nil {
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

**[*operations.GetEventGatewayListenerResponse](../../models/operations/geteventgatewaylistenerresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateEventGatewayListener

Updates an existing listener associated with the specified Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-event-gateway-listener" method="put" path="/v1/event-gateways/{gatewayId}/listeners/{eventGatewayListenerId}" -->
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

    res, err := s.EventGatewayListeners.UpdateEventGatewayListener(ctx, operations.UpdateEventGatewayListenerRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        EventGatewayListenerID: "927df604-cc92-426a-b160-a94aa6e5de7d",
        UpdateEventGatewayListenerRequest: &components.UpdateEventGatewayListenerRequest{
            Name: "<value>",
            Addresses: []string{
                "<value 1>",
                "<value 2>",
            },
            Ports: []components.EventGatewayListenerPort{},
            Labels: map[string]string{
                "env": "test",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayListener != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.UpdateEventGatewayListenerRequest](../../models/operations/updateeventgatewaylistenerrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.UpdateEventGatewayListenerResponse](../../models/operations/updateeventgatewaylistenerresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteEventGatewayListener

Deletes a specific listener associated with the Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-event-gateway-listener" method="delete" path="/v1/event-gateways/{gatewayId}/listeners/{eventGatewayListenerId}" -->
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

    res, err := s.EventGatewayListeners.DeleteEventGatewayListener(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "aafd826e-b37d-4b92-b5d2-2c34bf000e4a")
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
| `eventGatewayListenerID`                                 | *string*                                                 | :heavy_check_mark:                                       | The ID of the Event Gateway Listener.                    |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteEventGatewayListenerResponse](../../models/operations/deleteeventgatewaylistenerresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |