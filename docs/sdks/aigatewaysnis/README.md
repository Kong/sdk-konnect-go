# AIGatewaySNIs

## Overview

### Available Operations

* [ListAiGatewaySnisForCertificate](#listaigatewaysnisforcertificate) - List AI Gateway SNIs for a Certificate
* [CreateAiGatewaySniForCertificate](#createaigatewaysniforcertificate) - Create an AI Gateway SNI for a Certificate
* [GetAiGatewaySniForCertificate](#getaigatewaysniforcertificate) - Get an AI Gateway SNI for a Certificate
* [UpdateAiGatewaySniForCertificate](#updateaigatewaysniforcertificate) - Update an AI Gateway SNI for a Certificate
* [DeleteAiGatewaySniForCertificate](#deleteaigatewaysniforcertificate) - Delete an AI Gateway SNI for a Certificate
* [ListAiGatewaySnis](#listaigatewaysnis) - List AI Gateway SNIs
* [CreateAiGatewaySni](#createaigatewaysni) - Create an AI Gateway SNI
* [GetAiGatewaySni](#getaigatewaysni) - Get an AI Gateway SNI
* [UpdateAiGatewaySni](#updateaigatewaysni) - Update an AI Gateway SNI
* [DeleteAiGatewaySni](#deleteaigatewaysni) - Delete an AI Gateway SNI

## ListAiGatewaySnisForCertificate

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns a list of SNIs associated with a specific AI Gateway certificate.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-gateway-snis-for-certificate" method="get" path="/v1/ai-gateways/{gatewayId}/certificates/{certificateIdOrName}/snis" -->
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

    res, err := s.AIGatewaySNIs.ListAiGatewaySnisForCertificate(ctx, operations.ListAiGatewaySnisForCertificateRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        CertificateIDOrName: "my-entity-name",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAIGatewaySNIsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.ListAiGatewaySnisForCertificateRequest](../../models/operations/listaigatewaysnisforcertificaterequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.ListAiGatewaySnisForCertificateResponse](../../models/operations/listaigatewaysnisforcertificateresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## CreateAiGatewaySniForCertificate

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Creates a new SNI associated with a specific AI Gateway certificate.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-gateway-sni-for-certificate" method="post" path="/v1/ai-gateways/{gatewayId}/certificates/{certificateIdOrName}/snis" -->
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

    res, err := s.AIGatewaySNIs.CreateAiGatewaySniForCertificate(ctx, operations.CreateAiGatewaySniForCertificateRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        CertificateIDOrName: "my-entity-name",
        CreateAIGatewaySNIForCertificateRequest: components.CreateAIGatewaySNIForCertificateRequest{
            Name: "my-sni",
            DisplayName: "My SNI",
            Hostname: "example.org",
            Labels: map[string]string{
                "category": "finance",
            },
            ManagedBy: map[string]string{
                "owner": "terraform",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewaySNI != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.CreateAiGatewaySniForCertificateRequest](../../models/operations/createaigatewaysniforcertificaterequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.CreateAiGatewaySniForCertificateResponse](../../models/operations/createaigatewaysniforcertificateresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.ConflictError        | 409                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## GetAiGatewaySniForCertificate

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns the details of a specific SNI associated with an AI Gateway certificate.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-gateway-sni-for-certificate" method="get" path="/v1/ai-gateways/{gatewayId}/certificates/{certificateIdOrName}/snis/{sniIdOrName}" -->
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

    res, err := s.AIGatewaySNIs.GetAiGatewaySniForCertificate(ctx, operations.GetAiGatewaySniForCertificateRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        CertificateIDOrName: "my-entity-name",
        SniIDOrName: "my-entity-name",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewaySNI != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetAiGatewaySniForCertificateRequest](../../models/operations/getaigatewaysniforcertificaterequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.GetAiGatewaySniForCertificateResponse](../../models/operations/getaigatewaysniforcertificateresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## UpdateAiGatewaySniForCertificate

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Updates a specific SNI associated with an AI Gateway certificate.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-ai-gateway-sni-for-certificate" method="put" path="/v1/ai-gateways/{gatewayId}/certificates/{certificateIdOrName}/snis/{sniIdOrName}" -->
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

    res, err := s.AIGatewaySNIs.UpdateAiGatewaySniForCertificate(ctx, operations.UpdateAiGatewaySniForCertificateRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        CertificateIDOrName: "my-entity-name",
        SniIDOrName: "my-entity-name",
        UpdateAIGatewaySNIForCertificateRequest: components.UpdateAIGatewaySNIForCertificateRequest{
            Name: "my-sni",
            DisplayName: "My SNI",
            Hostname: "example.org",
            Labels: map[string]string{
                "category": "finance",
            },
            ManagedBy: map[string]string{
                "owner": "terraform",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewaySNI != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.UpdateAiGatewaySniForCertificateRequest](../../models/operations/updateaigatewaysniforcertificaterequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.UpdateAiGatewaySniForCertificateResponse](../../models/operations/updateaigatewaysniforcertificateresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## DeleteAiGatewaySniForCertificate

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Removes a specific SNI associated with an AI Gateway certificate.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-gateway-sni-for-certificate" method="delete" path="/v1/ai-gateways/{gatewayId}/certificates/{certificateIdOrName}/snis/{sniIdOrName}" -->
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

    res, err := s.AIGatewaySNIs.DeleteAiGatewaySniForCertificate(ctx, operations.DeleteAiGatewaySniForCertificateRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        CertificateIDOrName: "my-entity-name",
        SniIDOrName: "my-entity-name",
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
| `request`                                                                                                                | [operations.DeleteAiGatewaySniForCertificateRequest](../../models/operations/deleteaigatewaysniforcertificaterequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.DeleteAiGatewaySniForCertificateResponse](../../models/operations/deleteaigatewaysniforcertificateresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## ListAiGatewaySnis

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns a list of SNIs associated with the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-gateway-snis" method="get" path="/v1/ai-gateways/{gatewayId}/snis" -->
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

    res, err := s.AIGatewaySNIs.ListAiGatewaySnis(ctx, operations.ListAiGatewaySnisRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAIGatewaySNIsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListAiGatewaySnisRequest](../../models/operations/listaigatewaysnisrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListAiGatewaySnisResponse](../../models/operations/listaigatewaysnisresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## CreateAiGatewaySni

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Creates a new SNI for the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-gateway-sni" method="post" path="/v1/ai-gateways/{gatewayId}/snis" -->
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

    res, err := s.AIGatewaySNIs.CreateAiGatewaySni(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", components.CreateAIGatewaySNIRequest{
        Name: "my-sni",
        DisplayName: "My SNI",
        Hostname: "example.org",
        Labels: map[string]string{
            "category": "finance",
        },
        ManagedBy: map[string]string{
            "owner": "terraform",
        },
        Certificate: "my-certificate",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewaySNI != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `gatewayID`                                                                                  | `string`                                                                                     | :heavy_check_mark:                                                                           | The unique ID of the AI Gateway.                                                             | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                         |
| `createAIGatewaySNIRequest`                                                                  | [components.CreateAIGatewaySNIRequest](../../models/components/createaigatewaysnirequest.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

### Response

**[*operations.CreateAiGatewaySniResponse](../../models/operations/createaigatewaysniresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.ConflictError        | 409                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## GetAiGatewaySni

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns the details of a specific AI Gateway SNI.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-gateway-sni" method="get" path="/v1/ai-gateways/{gatewayId}/snis/{sniIdOrName}" -->
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

    res, err := s.AIGatewaySNIs.GetAiGatewaySni(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewaySNI != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique ID of the AI Gateway.                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `sniIDOrName`                                            | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway SNI.             | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAiGatewaySniResponse](../../models/operations/getaigatewaysniresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## UpdateAiGatewaySni

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Updates an existing AI Gateway SNI.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-ai-gateway-sni" method="put" path="/v1/ai-gateways/{gatewayId}/snis/{sniIdOrName}" -->
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

    res, err := s.AIGatewaySNIs.UpdateAiGatewaySni(ctx, operations.UpdateAiGatewaySniRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        SniIDOrName: "my-entity-name",
        UpdateAIGatewaySNIRequest: components.UpdateAIGatewaySNIRequest{
            Name: "my-sni",
            DisplayName: "My SNI",
            Hostname: "example.org",
            Labels: map[string]string{
                "category": "finance",
            },
            ManagedBy: map[string]string{
                "owner": "terraform",
            },
            Certificate: "my-certificate",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewaySNI != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateAiGatewaySniRequest](../../models/operations/updateaigatewaysnirequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateAiGatewaySniResponse](../../models/operations/updateaigatewaysniresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## DeleteAiGatewaySni

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Removes a specific AI Gateway SNI.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-gateway-sni" method="delete" path="/v1/ai-gateways/{gatewayId}/snis/{sniIdOrName}" -->
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

    res, err := s.AIGatewaySNIs.DeleteAiGatewaySni(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
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
| `sniIDOrName`                                            | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway SNI.             | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAiGatewaySniResponse](../../models/operations/deleteaigatewaysniresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |