# EventGatewayTLSTrustBundles

## Overview

A TLS trust bundle defines a set of trusted certificate authorities (CAs) used for client certificate
verification during mutual TLS (mTLS). Trust bundles are referenced by TLS listener policies to
determine which client certificates are accepted.


### Available Operations

* [ListEventGatewayTLSTrustBundles](#listeventgatewaytlstrustbundles) - List TLS Trust Bundles
* [CreateEventGatewayTLSTrustBundle](#createeventgatewaytlstrustbundle) - Create TLS Trust Bundle
* [GetEventGatewayTLSTrustBundle](#geteventgatewaytlstrustbundle) - Get a TLS Trust Bundle
* [UpdateEventGatewayTLSTrustBundle](#updateeventgatewaytlstrustbundle) - Update TLS Trust Bundle
* [DeleteEventGatewayTLSTrustBundle](#deleteeventgatewaytlstrustbundle) - Delete TLS Trust Bundle

## ListEventGatewayTLSTrustBundles

Returns a list of TLS trust bundles associated with the specified Event Gateway.

**Requires a minimum runtime version of `1.1`**.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-event-gateway-tls-trust-bundles" method="get" path="/v1/event-gateways/{gatewayId}/tls-trust-bundles" -->
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

    res, err := s.EventGatewayTLSTrustBundles.ListEventGatewayTLSTrustBundles(ctx, operations.ListEventGatewayTLSTrustBundlesRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListTLSTrustBundlesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.ListEventGatewayTLSTrustBundlesRequest](../../models/operations/listeventgatewaytlstrustbundlesrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.ListEventGatewayTLSTrustBundlesResponse](../../models/operations/listeventgatewaytlstrustbundlesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateEventGatewayTLSTrustBundle

Creates a new TLS trust bundle containing trusted CA certificates for client certificate verification.

**Requires a minimum runtime version of `1.1`**.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-event-gateway-tls-trust-bundle" method="post" path="/v1/event-gateways/{gatewayId}/tls-trust-bundles" -->
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

    res, err := s.EventGatewayTLSTrustBundles.CreateEventGatewayTLSTrustBundle(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.CreateTLSTrustBundleRequest{
        Name: "<value>",
        Config: components.TLSTrustBundleConfig{
            TrustedCa: "<value>",
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TLSTrustBundle != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `gatewayID`                                                                                      | `string`                                                                                         | :heavy_check_mark:                                                                               | The UUID of your Gateway.                                                                        | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                             |
| `createTLSTrustBundleRequest`                                                                    | [components.CreateTLSTrustBundleRequest](../../models/components/createtlstrustbundlerequest.md) | :heavy_check_mark:                                                                               | The request schema for creating a TLS trust bundle.                                              |                                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*operations.CreateEventGatewayTLSTrustBundleResponse](../../models/operations/createeventgatewaytlstrustbundleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetEventGatewayTLSTrustBundle

Returns information about a specific TLS trust bundle.

**Requires a minimum runtime version of `1.1`**.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-event-gateway-tls-trust-bundle" method="get" path="/v1/event-gateways/{gatewayId}/tls-trust-bundles/{tlsTrustBundleId}" -->
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

    res, err := s.EventGatewayTLSTrustBundles.GetEventGatewayTLSTrustBundle(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "316e1e3c-30ff-4cce-9fd6-0748eed00ad3")
    if err != nil {
        log.Fatal(err)
    }
    if res.TLSTrustBundle != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The UUID of your Gateway.                                | 9524ec7d-36d9-465d-a8c5-83a3c9390458                     |
| `tlsTrustBundleID`                                       | `string`                                                 | :heavy_check_mark:                                       | The ID of the TLS trust bundle.                          |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetEventGatewayTLSTrustBundleResponse](../../models/operations/geteventgatewaytlstrustbundleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateEventGatewayTLSTrustBundle

Updates an existing TLS trust bundle.

**Requires a minimum runtime version of `1.1`**.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-event-gateway-tls-trust-bundle" method="put" path="/v1/event-gateways/{gatewayId}/tls-trust-bundles/{tlsTrustBundleId}" -->
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

    res, err := s.EventGatewayTLSTrustBundles.UpdateEventGatewayTLSTrustBundle(ctx, operations.UpdateEventGatewayTLSTrustBundleRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        TLSTrustBundleID: "c4750cd9-2ea6-48fa-baa7-41e9deedd842",
        UpdateTLSTrustBundleRequest: components.UpdateTLSTrustBundleRequest{
            Labels: map[string]string{
                "env": "test",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TLSTrustBundle != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.UpdateEventGatewayTLSTrustBundleRequest](../../models/operations/updateeventgatewaytlstrustbundlerequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.UpdateEventGatewayTLSTrustBundleResponse](../../models/operations/updateeventgatewaytlstrustbundleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteEventGatewayTLSTrustBundle

Deletes a specific TLS trust bundle associated with the Event Gateway.

**Requires a minimum runtime version of `1.1`**.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-event-gateway-tls-trust-bundle" method="delete" path="/v1/event-gateways/{gatewayId}/tls-trust-bundles/{tlsTrustBundleId}" -->
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

    res, err := s.EventGatewayTLSTrustBundles.DeleteEventGatewayTLSTrustBundle(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "20c08d4b-f5a3-49ec-bf2c-2a93e6bd536a")
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
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The UUID of your Gateway.                                | 9524ec7d-36d9-465d-a8c5-83a3c9390458                     |
| `tlsTrustBundleID`                                       | `string`                                                 | :heavy_check_mark:                                       | The ID of the TLS trust bundle.                          |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteEventGatewayTLSTrustBundleResponse](../../models/operations/deleteeventgatewaytlstrustbundleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |