# EventGatewayStaticKeys
(*EventGatewayStaticKeys*)

## Overview

Static Keys are used by the Encrypt and Decrypt policies to encrypt data at rest


### Available Operations

* [ListEventGatewayStaticKeys](#listeventgatewaystatickeys) - List Event Gateway Static Keys
* [CreateEventGatewayStaticKey](#createeventgatewaystatickey) - Create a New Static Key
* [GetEventGatewayStaticKey](#geteventgatewaystatickey) - Get a Static Key
* [DeleteEventGatewayStaticKey](#deleteeventgatewaystatickey) - Delete Event Gateway Static Key

## ListEventGatewayStaticKeys

Returns a list of static keys associated to this event gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-event-gateway-static-keys" method="get" path="/v1/event-gateways/{gatewayId}/static-keys" -->
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

    res, err := s.EventGatewayStaticKeys.ListEventGatewayStaticKeys(ctx, operations.ListEventGatewayStaticKeysRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListEventGatewayStaticKeysResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListEventGatewayStaticKeysRequest](../../models/operations/listeventgatewaystatickeysrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListEventGatewayStaticKeysResponse](../../models/operations/listeventgatewaystatickeysresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateEventGatewayStaticKey

Create new static key to this event gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-event-gateway-static-key" method="post" path="/v1/event-gateways/{gatewayId}/static-keys" -->
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

    res, err := s.EventGatewayStaticKeys.CreateEventGatewayStaticKey(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", &components.EventGatewayStaticKeyCreate{
        Name: "<value>",
        Labels: map[string]string{
            "env": "test",
        },
        Value: "${vault.env['MY_ENV_VAR']}",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayStaticKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       | Example                                                                                           |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |                                                                                                   |
| `gatewayID`                                                                                       | *string*                                                                                          | :heavy_check_mark:                                                                                | The UUID of your Gateway.                                                                         | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                              |
| `eventGatewayStaticKeyCreate`                                                                     | [*components.EventGatewayStaticKeyCreate](../../models/components/eventgatewaystatickeycreate.md) | :heavy_minus_sign:                                                                                | The request schema for creating a static key.                                                     |                                                                                                   |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |                                                                                                   |

### Response

**[*operations.CreateEventGatewayStaticKeyResponse](../../models/operations/createeventgatewaystatickeyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetEventGatewayStaticKey

Returns information about an individual static key.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-event-gateway-static-key" method="get" path="/v1/event-gateways/{gatewayId}/static-keys/{staticKeyId}" -->
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

    res, err := s.EventGatewayStaticKeys.GetEventGatewayStaticKey(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "8d254be7-bd8f-484b-82e5-613cd779035d")
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayStaticKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | *string*                                                 | :heavy_check_mark:                                       | The UUID of your Gateway.                                | 9524ec7d-36d9-465d-a8c5-83a3c9390458                     |
| `staticKeyID`                                            | *string*                                                 | :heavy_check_mark:                                       | The ID of the static key.                                |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetEventGatewayStaticKeyResponse](../../models/operations/geteventgatewaystatickeyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteEventGatewayStaticKey

Deletes a specific static key associated with the Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-event-gateway-static-key" method="delete" path="/v1/event-gateways/{gatewayId}/static-keys/{staticKeyId}" -->
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

    res, err := s.EventGatewayStaticKeys.DeleteEventGatewayStaticKey(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "ca8184fe-dd21-42a1-be12-30b80192392f")
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
| `staticKeyID`                                            | *string*                                                 | :heavy_check_mark:                                       | The ID of the static key.                                |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteEventGatewayStaticKeyResponse](../../models/operations/deleteeventgatewaystatickeyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |