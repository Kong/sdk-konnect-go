# TransitGateways
(*TransitGateways*)

## Overview

### Available Operations

* [ListTransitGateways](#listtransitgateways) - List Transit Gateways
* [CreateTransitGateway](#createtransitgateway) - Create Transit Gateway
* [GetTransitGateway](#gettransitgateway) - Get Transit Gateway
* [DeleteTransitGateway](#deletetransitgateway) - Delete Transit Gateway

## ListTransitGateways

Returns a paginated collection of transit gateways for a given network.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.TransitGateways.ListTransitGateways(ctx, operations.ListTransitGatewaysRequest{
        NetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
        PageSize: sdkkonnectgo.Int64(10),
        PageNumber: sdkkonnectgo.Int64(1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListTransitGatewaysResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListTransitGatewaysRequest](../../models/operations/listtransitgatewaysrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListTransitGatewaysResponse](../../models/operations/listtransitgatewaysresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateTransitGateway

Creates a new transit gateway for a given network.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.TransitGateways.CreateTransitGateway(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreateCreateTransitGatewayRequestAWSTransitGateway(
        components.AWSTransitGateway{
            Name: "us-east-2 transit gateway",
            DNSConfig: []components.TransitGatewayDNSConfig{
                components.TransitGatewayDNSConfig{
                    RemoteDNSServerIPAddresses: []string{
                        "10.0.0.2",
                    },
                    DomainProxyList: []string{
                        "foobar.com",
                    },
                },
                components.TransitGatewayDNSConfig{
                    RemoteDNSServerIPAddresses: []string{
                        "10.0.0.2",
                    },
                    DomainProxyList: []string{
                        "foobar.com",
                    },
                },
            },
            CidrBlocks: []string{
                "10.0.0.0/8",
                "100.64.0.0/10",
                "172.16.0.0/12",
            },
            TransitGatewayAttachmentConfig: components.AwsTransitGatewayAttachmentConfig{
                Kind: components.AWSTransitGatewayAttachmentTypeAwsTransitGatewayAttachment,
                TransitGatewayID: "<id>",
                RAMShareArn: "<value>",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.TransitGatewayResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `networkID`                                                                                      | *string*                                                                                         | :heavy_check_mark:                                                                               | The network to operate on.                                                                       | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                                                             |
| `createTransitGatewayRequest`                                                                    | [components.CreateTransitGatewayRequest](../../models/components/createtransitgatewayrequest.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*operations.CreateTransitGatewayResponse](../../models/operations/createtransitgatewayresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetTransitGateway

Retrieves a transit gateway by ID for a given network.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.TransitGateways.GetTransitGateway(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", "0850820b-d153-4a2a-b9be-7d2204779139")
    if err != nil {
        log.Fatal(err)
    }
    if res.TransitGatewayResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `networkID`                                              | *string*                                                 | :heavy_check_mark:                                       | The network to operate on.                               | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                     |
| `transitGatewayID`                                       | *string*                                                 | :heavy_check_mark:                                       | The ID of the transit gateway to operate on.             | 0850820b-d153-4a2a-b9be-7d2204779139                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetTransitGatewayResponse](../../models/operations/gettransitgatewayresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteTransitGateway

Deletes a transit gateway by ID for a given network.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.TransitGateways.DeleteTransitGateway(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", "0850820b-d153-4a2a-b9be-7d2204779139")
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
| `networkID`                                              | *string*                                                 | :heavy_check_mark:                                       | The network to operate on.                               | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                     |
| `transitGatewayID`                                       | *string*                                                 | :heavy_check_mark:                                       | The ID of the transit gateway to operate on.             | 0850820b-d153-4a2a-b9be-7d2204779139                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteTransitGatewayResponse](../../models/operations/deletetransitgatewayresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |