# AIGatewayCustomPolicies

## Overview

Custom policies that let you bring your own Lua plugin schema and handler to the AI Gateway.

### Available Operations

* [ListAiGatewayCustomPolicies](#listaigatewaycustompolicies) - List AI Gateway Custom Policies
* [CreateAiGatewayCustomPolicy](#createaigatewaycustompolicy) - Create an AI Gateway Custom Policy
* [GetAiGatewayCustomPolicy](#getaigatewaycustompolicy) - Get an AI Gateway Custom Policy
* [UpdateAiGatewayCustomPolicy](#updateaigatewaycustompolicy) - Update an AI Gateway Custom Policy
* [DeleteAiGatewayCustomPolicy](#deleteaigatewaycustompolicy) - Delete an AI Gateway Custom Policy

## ListAiGatewayCustomPolicies

Returns a list of custom policies registered for the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-gateway-custom-policies" method="get" path="/v1/ai-gateways/{gatewayId}/custom-policies" -->
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

    res, err := s.AIGatewayCustomPolicies.ListAiGatewayCustomPolicies(ctx, operations.ListAiGatewayCustomPoliciesRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAIGatewayCustomPoliciesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListAiGatewayCustomPoliciesRequest](../../models/operations/listaigatewaycustompoliciesrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ListAiGatewayCustomPoliciesResponse](../../models/operations/listaigatewaycustompoliciesresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## CreateAiGatewayCustomPolicy

Registers a new custom policy for the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-gateway-custom-policy" method="post" path="/v1/ai-gateways/{gatewayId}/custom-policies" -->
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

    res, err := s.AIGatewayCustomPolicies.CreateAiGatewayCustomPolicy(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", components.CreateCreateAIGatewayCustomPolicyRequestInstalled(
        components.CreateAIGatewayCustomPolicyInstalledRequest{
            Name: "my-installed-custom-policy",
            Type: components.CreateAIGatewayCustomPolicyInstalledRequestTypeInstalled,
            DisplayName: "Custom Policy - Installed plugin",
            Schema: "<lua_schema>",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayCustomPolicy != nil {
        switch res.AIGatewayCustomPolicy.Type {
            case components.AIGatewayCustomPolicyTypeInstalled:
                // res.AIGatewayCustomPolicy.AIGatewayCustomPolicyInstalled is populated
            case components.AIGatewayCustomPolicyTypeStreaming:
                // res.AIGatewayCustomPolicy.AIGatewayCustomPolicyStreaming is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    | Example                                                                                                        |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |                                                                                                                |
| `gatewayID`                                                                                                    | `string`                                                                                                       | :heavy_check_mark:                                                                                             | The unique ID of the AI Gateway.                                                                               | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                                           |
| `createAIGatewayCustomPolicyRequest`                                                                           | [components.CreateAIGatewayCustomPolicyRequest](../../models/components/createaigatewaycustompolicyrequest.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |                                                                                                                |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |                                                                                                                |

### Response

**[*operations.CreateAiGatewayCustomPolicyResponse](../../models/operations/createaigatewaycustompolicyresponse.md), error**

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

## GetAiGatewayCustomPolicy

Returns the details of a specific AI Gateway custom policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-gateway-custom-policy" method="get" path="/v1/ai-gateways/{gatewayId}/custom-policies/{customPolicyIdOrName}" -->
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

    res, err := s.AIGatewayCustomPolicies.GetAiGatewayCustomPolicy(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayCustomPolicy != nil {
        switch res.AIGatewayCustomPolicy.Type {
            case components.AIGatewayCustomPolicyTypeInstalled:
                // res.AIGatewayCustomPolicy.AIGatewayCustomPolicyInstalled is populated
            case components.AIGatewayCustomPolicyTypeStreaming:
                // res.AIGatewayCustomPolicy.AIGatewayCustomPolicyStreaming is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique ID of the AI Gateway.                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `customPolicyIDOrName`                                   | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway custom policy.   | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAiGatewayCustomPolicyResponse](../../models/operations/getaigatewaycustompolicyresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## UpdateAiGatewayCustomPolicy

Updates the configuration of an existing AI Gateway custom policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-ai-gateway-custom-policy" method="put" path="/v1/ai-gateways/{gatewayId}/custom-policies/{customPolicyIdOrName}" -->
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

    res, err := s.AIGatewayCustomPolicies.UpdateAiGatewayCustomPolicy(ctx, operations.UpdateAiGatewayCustomPolicyRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        CustomPolicyIDOrName: "my-entity-name",
        UpdateAIGatewayCustomPolicyRequest: components.CreateUpdateAIGatewayCustomPolicyRequestStreaming(
            components.UpdateAIGatewayCustomPolicyStreamingRequest{
                Name: "my-streaming-custom-policy",
                Type: components.UpdateAIGatewayCustomPolicyStreamingRequestTypeStreaming,
                DisplayName: "Custom Policy Streaming Plugin",
                Schema: "<lua_schema>",
                Handler: "<lua_handler>",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayCustomPolicy != nil {
        switch res.AIGatewayCustomPolicy.Type {
            case components.AIGatewayCustomPolicyTypeInstalled:
                // res.AIGatewayCustomPolicy.AIGatewayCustomPolicyInstalled is populated
            case components.AIGatewayCustomPolicyTypeStreaming:
                // res.AIGatewayCustomPolicy.AIGatewayCustomPolicyStreaming is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.UpdateAiGatewayCustomPolicyRequest](../../models/operations/updateaigatewaycustompolicyrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.UpdateAiGatewayCustomPolicyResponse](../../models/operations/updateaigatewaycustompolicyresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## DeleteAiGatewayCustomPolicy

Removes a specific AI Gateway custom policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-gateway-custom-policy" method="delete" path="/v1/ai-gateways/{gatewayId}/custom-policies/{customPolicyIdOrName}" -->
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

    res, err := s.AIGatewayCustomPolicies.DeleteAiGatewayCustomPolicy(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
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
| `customPolicyIDOrName`                                   | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway custom policy.   | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAiGatewayCustomPolicyResponse](../../models/operations/deleteaigatewaycustompolicyresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |