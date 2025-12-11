# EventGatewayDataPlaneCertificates

## Overview

DataPlane certificates control how your running Event Gateway instances connect to the Control Plane


### Available Operations

* [ListEventGatewayDataPlaneCertificates](#listeventgatewaydataplanecertificates) - List Event Gateway DataPlane Certificates
* [CreateEventGatewayDataPlaneCertificate](#createeventgatewaydataplanecertificate) - Create a New DataPlane Certificate
* [GetEventGatewayDataPlaneCertificate](#geteventgatewaydataplanecertificate) - Get a DataPlane Certificate
* [UpdateEventGatewayDataPlaneCertificate](#updateeventgatewaydataplanecertificate) - Update Event Gateway DataPlane Certificate
* [DeleteEventGatewayDataPlaneCertificate](#deleteeventgatewaydataplanecertificate) - Delete Event Gateway DataPlane Certificate

## ListEventGatewayDataPlaneCertificates

Returns a list of dataplane certificates that are associated to this event gateway. A dataplane certificate allows dataplanes configured with the certificate and corresponding private key to establish connection with this event gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-event-gateway-data-plane-certificates" method="get" path="/v1/event-gateways/{gatewayId}/data-plane-certificates" -->
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

    res, err := s.EventGatewayDataPlaneCertificates.ListEventGatewayDataPlaneCertificates(ctx, operations.ListEventGatewayDataPlaneCertificatesRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListEventGatewayDataPlaneCertificatesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.ListEventGatewayDataPlaneCertificatesRequest](../../models/operations/listeventgatewaydataplanecertificatesrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.ListEventGatewayDataPlaneCertificatesResponse](../../models/operations/listeventgatewaydataplanecertificatesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateEventGatewayDataPlaneCertificate

Create new dataplane certificate to this event gateway. A dataplane certificate allows dataplanes configured with the certificate and corresponding private key to establish connection with this event gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-event-gateway-data-plane-certificate" method="post" path="/v1/event-gateways/{gatewayId}/data-plane-certificates" -->
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

    res, err := s.EventGatewayDataPlaneCertificates.CreateEventGatewayDataPlaneCertificate(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayDataPlaneCertificate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                             | Type                                                                                                                                  | Required                                                                                                                              | Description                                                                                                                           | Example                                                                                                                               |
| ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                 | :heavy_check_mark:                                                                                                                    | The context to use for the request.                                                                                                   |                                                                                                                                       |
| `gatewayID`                                                                                                                           | *string*                                                                                                                              | :heavy_check_mark:                                                                                                                    | The UUID of your Gateway.                                                                                                             | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                                                                  |
| `createEventGatewayDataPlaneCertificateRequest`                                                                                       | [*components.CreateEventGatewayDataPlaneCertificateRequest](../../models/components/createeventgatewaydataplanecertificaterequest.md) | :heavy_minus_sign:                                                                                                                    | Request body for creating a certificate.                                                                                              |                                                                                                                                       |
| `opts`                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                              | :heavy_minus_sign:                                                                                                                    | The options for this request.                                                                                                         |                                                                                                                                       |

### Response

**[*operations.CreateEventGatewayDataPlaneCertificateResponse](../../models/operations/createeventgatewaydataplanecertificateresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetEventGatewayDataPlaneCertificate

Returns information about an individual dataplane certificate.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-event-gateway-data-plane-certificate" method="get" path="/v1/event-gateways/{gatewayId}/data-plane-certificates/{certificateId}" -->
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

    res, err := s.EventGatewayDataPlaneCertificates.GetEventGatewayDataPlaneCertificate(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "4df4520b-0ab6-4f5e-87b1-ac5150f3968a")
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayDataPlaneCertificate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | *string*                                                 | :heavy_check_mark:                                       | The UUID of your Gateway.                                | 9524ec7d-36d9-465d-a8c5-83a3c9390458                     |
| `certificateID`                                          | *string*                                                 | :heavy_check_mark:                                       | The ID of the dataplane certificate.                     |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetEventGatewayDataPlaneCertificateResponse](../../models/operations/geteventgatewaydataplanecertificateresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateEventGatewayDataPlaneCertificate

Updates an existing dataplane certificate associated with the specified Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-event-gateway-data-plane-certificate" method="put" path="/v1/event-gateways/{gatewayId}/data-plane-certificates/{certificateId}" -->
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

    res, err := s.EventGatewayDataPlaneCertificates.UpdateEventGatewayDataPlaneCertificate(ctx, operations.UpdateEventGatewayDataPlaneCertificateRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        CertificateID: "8a789778-d873-4a5b-ad36-40cd2f78c306",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayDataPlaneCertificate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.UpdateEventGatewayDataPlaneCertificateRequest](../../models/operations/updateeventgatewaydataplanecertificaterequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.UpdateEventGatewayDataPlaneCertificateResponse](../../models/operations/updateeventgatewaydataplanecertificateresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteEventGatewayDataPlaneCertificate

Deletes a specific dataplane certificate associated with the Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-event-gateway-data-plane-certificate" method="delete" path="/v1/event-gateways/{gatewayId}/data-plane-certificates/{certificateId}" -->
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

    res, err := s.EventGatewayDataPlaneCertificates.DeleteEventGatewayDataPlaneCertificate(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "0ae293bb-20e2-4a8f-aaeb-70cec1594054")
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
| `certificateID`                                          | *string*                                                 | :heavy_check_mark:                                       | The ID of the dataplane certificate.                     |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteEventGatewayDataPlaneCertificateResponse](../../models/operations/deleteeventgatewaydataplanecertificateresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |