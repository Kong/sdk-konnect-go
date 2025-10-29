# APIGatewayDataPlaneCertificates
(*APIGatewayDataPlaneCertificates*)

## Overview

### Available Operations

* [ListAPIGatewayDataPlaneCertificates](#listapigatewaydataplanecertificates) - List API Gateway DataPlane Certificates
* [CreateAPIGatewayDataPlaneCertificate](#createapigatewaydataplanecertificate) - Create a New DataPlane Certificate

## ListAPIGatewayDataPlaneCertificates

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns a list of dataplane certificates that are associated to this API gateway. A dataplane certificate allows dataplanes configured with the certificate and corresponding private key to establish connection with this API gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-api-gateway-data-plane-certificates" method="get" path="/v1/api-gateways/{gatewayId}/data-plane-certificates" -->
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

    res, err := s.APIGatewayDataPlaneCertificates.ListAPIGatewayDataPlaneCertificates(ctx, operations.ListAPIGatewayDataPlaneCertificatesRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAPIGatewayDataPlaneCertificatesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.ListAPIGatewayDataPlaneCertificatesRequest](../../models/operations/listapigatewaydataplanecertificatesrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.ListAPIGatewayDataPlaneCertificatesResponse](../../models/operations/listapigatewaydataplanecertificatesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateAPIGatewayDataPlaneCertificate

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Create new dataplane certificate to this API gateway. A dataplane certificate allows dataplanes configured with the certificate and corresponding private key to establish connection with this API gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-api-gateway-data-plane-certificate" method="post" path="/v1/api-gateways/{gatewayId}/data-plane-certificates" -->
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

    res, err := s.APIGatewayDataPlaneCertificates.CreateAPIGatewayDataPlaneCertificate(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.APIGatewayDataPlaneCertificate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                         | Type                                                                                                                              | Required                                                                                                                          | Description                                                                                                                       | Example                                                                                                                           |
| --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                             | :heavy_check_mark:                                                                                                                | The context to use for the request.                                                                                               |                                                                                                                                   |
| `gatewayID`                                                                                                                       | *string*                                                                                                                          | :heavy_check_mark:                                                                                                                | The UUID of your Gateway.                                                                                                         | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                                                              |
| `createAPIGatewayDataPlaneCertificateRequest`                                                                                     | [*components.CreateAPIGatewayDataPlaneCertificateRequest](../../models/components/createapigatewaydataplanecertificaterequest.md) | :heavy_minus_sign:                                                                                                                | Request body for creating a certificate.                                                                                          |                                                                                                                                   |
| `opts`                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                          | :heavy_minus_sign:                                                                                                                | The options for this request.                                                                                                     |                                                                                                                                   |

### Response

**[*operations.CreateAPIGatewayDataPlaneCertificateResponse](../../models/operations/createapigatewaydataplanecertificateresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |