# AIGatewayAgents

## Overview

AI Agents registered with the AI Gateway.

### Available Operations

* [ListAiGatewayAgents](#listaigatewayagents) - List AI Gateway Agents
* [CreateAiGatewayAgent](#createaigatewayagent) - Create an AI Gateway Agent
* [GetAiGatewayAgent](#getaigatewayagent) - Get an AI Gateway Agent
* [UpdateAiGatewayAgent](#updateaigatewayagent) - Update an AI Gateway Agent
* [DeleteAiGatewayAgent](#deleteaigatewayagent) - Delete an AI Gateway Agent

## ListAiGatewayAgents

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns a list of all agents registered in the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-gateway-agents" method="get" path="/v1/ai-gateways/{gatewayId}/agents" -->
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

    res, err := s.AIGatewayAgents.ListAiGatewayAgents(ctx, operations.ListAiGatewayAgentsRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAIGatewayAgentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListAiGatewayAgentsRequest](../../models/operations/listaigatewayagentsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListAiGatewayAgentsResponse](../../models/operations/listaigatewayagentsresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## CreateAiGatewayAgent

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Creates a new agent for the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-gateway-agent" method="post" path="/v1/ai-gateways/{gatewayId}/agents" -->
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

    res, err := s.AIGatewayAgents.CreateAiGatewayAgent(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", components.CreateAIGatewayAgentRequest{
        DisplayName: "Kong Air Flight Booking Agent",
        Name: "kongair-flight-booking-agent",
        Enabled: sdkkonnectgo.Pointer(true),
        Type: components.CreateAIGatewayAgentRequestTypeA2a,
        Access: &components.AIGatewayAgentAccess{
            Acls: sdkkonnectgo.Pointer(components.CreateAIGatewayACLSAIGatewayAllowACL(
                components.AIGatewayAllowACL{
                    Allow: []string{
                        "consumer-group-1",
                    },
                },
            )),
        },
        Config: components.CreateAIGatewayAgentRequestConfig{
            URL: "https://booking-agent.internal.kongair.com",
            Route: &components.AIGatewayRouteConfig{
                Headers: map[string]any{
                    "version": []any{
                        "v1",
                        "v2",
                    },
                },
                Hosts: []string{
                    "foo.example.com",
                },
            },
            Logging: &components.CreateAIGatewayAgentRequestLogging{
                MaxPayloadSize: sdkkonnectgo.Pointer[int64](524288),
            },
        },
        Labels: map[string]string{
            "category": "finance",
        },
        ManagedBy: map[string]string{
            "owner": "terraform",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayAgent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `gatewayID`                                                                                      | `string`                                                                                         | :heavy_check_mark:                                                                               | The unique ID of the AI Gateway.                                                                 | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                             |
| `createAIGatewayAgentRequest`                                                                    | [components.CreateAIGatewayAgentRequest](../../models/components/createaigatewayagentrequest.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*operations.CreateAiGatewayAgentResponse](../../models/operations/createaigatewayagentresponse.md), error**

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

## GetAiGatewayAgent

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns the details of a specific AI Gateway agent.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-gateway-agent" method="get" path="/v1/ai-gateways/{gatewayId}/agents/{agentIdOrName}" -->
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

    res, err := s.AIGatewayAgents.GetAiGatewayAgent(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayAgent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique ID of the AI Gateway.                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `agentIDOrName`                                          | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway agent.           | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAiGatewayAgentResponse](../../models/operations/getaigatewayagentresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## UpdateAiGatewayAgent

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Updates the configuration of an existing AI Gateway agent.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-ai-gateway-agent" method="put" path="/v1/ai-gateways/{gatewayId}/agents/{agentIdOrName}" -->
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

    res, err := s.AIGatewayAgents.UpdateAiGatewayAgent(ctx, operations.UpdateAiGatewayAgentRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        AgentIDOrName: "my-entity-name",
        UpdateAIGatewayAgentRequest: components.UpdateAIGatewayAgentRequest{
            DisplayName: "Kong Air Flight Booking Agent",
            Name: "kongair-flight-booking-agent",
            Enabled: sdkkonnectgo.Pointer(true),
            Type: components.UpdateAIGatewayAgentRequestTypeA2a,
            Access: &components.AIGatewayAgentAccess{
                Acls: sdkkonnectgo.Pointer(components.CreateAIGatewayACLSAIGatewayAllowACL(
                    components.AIGatewayAllowACL{
                        Allow: []string{
                            "consumer-group-1",
                        },
                    },
                )),
            },
            Config: components.UpdateAIGatewayAgentRequestConfig{
                URL: "https://booking-agent.internal.kongair.com",
                Route: &components.AIGatewayRouteConfig{
                    Headers: map[string]any{
                        "version": []any{
                            "v1",
                            "v2",
                        },
                    },
                    Hosts: []string{
                        "foo.example.com",
                    },
                },
                Logging: &components.UpdateAIGatewayAgentRequestLogging{
                    MaxPayloadSize: sdkkonnectgo.Pointer[int64](524288),
                },
            },
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
    if res.AIGatewayAgent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateAiGatewayAgentRequest](../../models/operations/updateaigatewayagentrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateAiGatewayAgentResponse](../../models/operations/updateaigatewayagentresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## DeleteAiGatewayAgent

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Removes a specific AI Gateway agent.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-gateway-agent" method="delete" path="/v1/ai-gateways/{gatewayId}/agents/{agentIdOrName}" -->
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

    res, err := s.AIGatewayAgents.DeleteAiGatewayAgent(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
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
| `agentIDOrName`                                          | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway agent.           | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAiGatewayAgentResponse](../../models/operations/deleteaigatewayagentresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |