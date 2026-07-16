# AIGatewayDataPlaneCertificates

## Overview

API related to the management of AI Gateway DataPlane Certificates.

### Available Operations

* [ListAiGatewayDataPlaneCertificates](#listaigatewaydataplanecertificates) - List AI Gateway DataPlane Certificates
* [CreateAiGatewayDataPlaneCertificate](#createaigatewaydataplanecertificate) - Create New AI Gateway DataPlane Certificate
* [GetAiGatewayDataPlaneCertificate](#getaigatewaydataplanecertificate) - Get a DataPlane Certificate
* [DeleteAiGatewayDataPlaneCertificate](#deleteaigatewaydataplanecertificate) - Delete AI Gateway DataPlane Certificate

## ListAiGatewayDataPlaneCertificates

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns a list of DataPlane certificates that are associated to this AI Gateway. A DataPlane certificate allows DataPlanes configured with the certificate and corresponding private key to establish connection with this AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-gateway-data-plane-certificates" method="get" path="/v1/ai-gateways/{gatewayId}/data-plane-certificates" -->
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

    res, err := s.AIGatewayDataPlaneCertificates.ListAiGatewayDataPlaneCertificates(ctx, operations.ListAiGatewayDataPlaneCertificatesRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAIGatewayDataPlaneCertificatesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.ListAiGatewayDataPlaneCertificatesRequest](../../models/operations/listaigatewaydataplanecertificatesrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.ListAiGatewayDataPlaneCertificatesResponse](../../models/operations/listaigatewaydataplanecertificatesresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## CreateAiGatewayDataPlaneCertificate

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Create a new DataPlane Certificate for this AI Gateway. A DataPlane certificate allows DataPlanes configured with the certificate and corresponding private key to establish connection with this AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-gateway-data-plane-certificate" method="post" path="/v1/ai-gateways/{gatewayId}/data-plane-certificates" -->
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

    res, err := s.AIGatewayDataPlaneCertificates.CreateAiGatewayDataPlaneCertificate(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayDataPlaneClientCertificate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                       | Type                                                                                                                            | Required                                                                                                                        | Description                                                                                                                     | Example                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                           | :heavy_check_mark:                                                                                                              | The context to use for the request.                                                                                             |                                                                                                                                 |
| `gatewayID`                                                                                                                     | `string`                                                                                                                        | :heavy_check_mark:                                                                                                              | The unique ID of the AI Gateway.                                                                                                | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                                                            |
| `createAIGatewayDataPlaneCertificateRequest`                                                                                    | [*components.CreateAIGatewayDataPlaneCertificateRequest](../../models/components/createaigatewaydataplanecertificaterequest.md) | :heavy_minus_sign:                                                                                                              | N/A                                                                                                                             |                                                                                                                                 |
| `opts`                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                        | :heavy_minus_sign:                                                                                                              | The options for this request.                                                                                                   |                                                                                                                                 |

### Response

**[*operations.CreateAiGatewayDataPlaneCertificateResponse](../../models/operations/createaigatewaydataplanecertificateresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## GetAiGatewayDataPlaneCertificate

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Retrieve a DataPlane certificate associated to this AI Gateway. A DataPlane certificate allows DataPlanes configured with the certificate and corresponding private key to establish connection with this AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-gateway-data-plane-certificate" method="get" path="/v1/ai-gateways/{gatewayId}/data-plane-certificates/{certificateId}" -->
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

    res, err := s.AIGatewayDataPlaneCertificates.GetAiGatewayDataPlaneCertificate(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "bf138ba2-c9b1-4229-b268-04d9d8a6410b")
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayDataPlaneClientCertificate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique ID of the AI Gateway.                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `certificateID`                                          | `string`                                                 | :heavy_check_mark:                                       | The unique ID of the DataPlane Certificate.              | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAiGatewayDataPlaneCertificateResponse](../../models/operations/getaigatewaydataplanecertificateresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## DeleteAiGatewayDataPlaneCertificate

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Remove a DataPlane client certificate associated to this AI Gateway. Removing a DataPlane certificate would invalidate any DataPlanes currently connected to this AI Gateway using this certificate.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-gateway-data-plane-certificate" method="delete" path="/v1/ai-gateways/{gatewayId}/data-plane-certificates/{certificateId}" -->
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

    res, err := s.AIGatewayDataPlaneCertificates.DeleteAiGatewayDataPlaneCertificate(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "bf138ba2-c9b1-4229-b268-04d9d8a6410b")
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
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique ID of the AI Gateway.                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `certificateID`                                          | `string`                                                 | :heavy_check_mark:                                       | The unique ID of the DataPlane Certificate.              | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAiGatewayDataPlaneCertificateResponse](../../models/operations/deleteaigatewaydataplanecertificateresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |