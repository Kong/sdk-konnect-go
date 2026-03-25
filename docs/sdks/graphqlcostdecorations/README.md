# GraphQLCostDecorations

## Overview

### Available Operations

* [ListGraphqlRateLimitingAdvancedCostWithService](#listgraphqlratelimitingadvancedcostwithservice) - List all GraphQL Cost Decorations associated with a Service
* [CreateGraphqlRateLimitingAdvancedCostWithService](#creategraphqlratelimitingadvancedcostwithservice) - Create a new GraphQL Cost Decoration associated with a Service

## ListGraphqlRateLimitingAdvancedCostWithService

List all GraphQL Cost Decorations associated with a Service

### Example Usage

<!-- UsageSnippet language="go" operationID="list-graphql-rate-limiting-advanced-cost-with-service" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/services/{ServiceId}/graphql-rate-limiting-advanced/costs" -->
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

    res, err := s.GraphQLCostDecorations.ListGraphqlRateLimitingAdvancedCostWithService(ctx, operations.ListGraphqlRateLimitingAdvancedCostWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        Tags: sdkkonnectgo.Pointer("tag1,tag2"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.ListGraphqlRateLimitingAdvancedCostWithServiceRequest](../../models/operations/listgraphqlratelimitingadvancedcostwithservicerequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |
| `opts`                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |

### Response

**[*operations.ListGraphqlRateLimitingAdvancedCostWithServiceResponse](../../models/operations/listgraphqlratelimitingadvancedcostwithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateGraphqlRateLimitingAdvancedCostWithService

Create a new GraphQL Cost Decoration associated with a Service

### Example Usage

<!-- UsageSnippet language="go" operationID="create-graphql-rate-limiting-advanced-cost-with-service" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/services/{ServiceId}/graphql-rate-limiting-advanced/costs" -->
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

    res, err := s.GraphQLCostDecorations.CreateGraphqlRateLimitingAdvancedCostWithService(ctx, operations.CreateGraphqlRateLimitingAdvancedCostWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        GraphQLCostDecorationWithoutParents: components.GraphQLCostDecorationWithoutParents{
            TypePath: "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GraphQLCostDecoration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.CreateGraphqlRateLimitingAdvancedCostWithServiceRequest](../../models/operations/creategraphqlratelimitingadvancedcostwithservicerequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.CreateGraphqlRateLimitingAdvancedCostWithServiceResponse](../../models/operations/creategraphqlratelimitingadvancedcostwithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |