# AIGatewayMCPServers

## Overview

MCP Servers that expose tools for AI Gateway integrations.

### Available Operations

* [ListAiGatewayMcpServers](#listaigatewaymcpservers) - List MCP Servers
* [CreateAiGatewayMcpServer](#createaigatewaymcpserver) - Create an MCP Server
* [GetAiGatewayMcpServer](#getaigatewaymcpserver) - Get an MCP Server
* [UpdateAiGatewayMcpServer](#updateaigatewaymcpserver) - Update an MCP Server
* [DeleteAiGatewayMcpServer](#deleteaigatewaymcpserver) - Delete an MCP Server

## ListAiGatewayMcpServers

Returns a list of MCP Servers configured for the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-gateway-mcp-servers" method="get" path="/v1/ai-gateways/{gatewayId}/mcp-servers" -->
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

    res, err := s.AIGatewayMCPServers.ListAiGatewayMcpServers(ctx, operations.ListAiGatewayMcpServersRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAIGatewayMCPServersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListAiGatewayMcpServersRequest](../../models/operations/listaigatewaymcpserversrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListAiGatewayMcpServersResponse](../../models/operations/listaigatewaymcpserversresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## CreateAiGatewayMcpServer

Registers a new MCP Server for the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-gateway-mcp-server" method="post" path="/v1/ai-gateways/{gatewayId}/mcp-servers" -->
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

    res, err := s.AIGatewayMCPServers.CreateAiGatewayMcpServer(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", components.CreateCreateAIGatewayMCPServerRequestUpstreamServer(
        components.CreateAIGatewayMCPServerUpstreamServerOauthAccessToken(
            components.AIGatewayMCPServerBaseACLPropertiesOauthAIGatewayMCPServerUpstreamServerAIGatewayMCPServerBaseACLPropertiesOauth{
                ACLAttributeType: components.AIGatewayMCPServerBaseACLPropertiesOauthAIGatewayMCPServerUpstreamServerACLAttributeTypeOauthAccessToken,
                AccessTokenClaimField: "<value>",
                Type: components.AIGatewayMCPServerBaseACLPropertiesOauthAIGatewayMCPServerUpstreamServerTypeUpstreamServer,
                Config: components.AIGatewayMCPServerUpstreamServerConfig{
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
                    URL: "https://mcp.internal.kongair.com",
                    ToolsCacheTTLSeconds: 55656,
                },
                DisplayName: "Kong Air Flights",
                Name: "kongair-flights",
            },
        ),
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayMCPServer != nil {
        switch res.AIGatewayMCPServer.Type {
            case components.AIGatewayMCPServerTypeConversionOnly:
                // res.AIGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerConversionOnly is populated
            case components.AIGatewayMCPServerTypeConversionListener:
                // res.AIGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerConversionListener is populated
            case components.AIGatewayMCPServerTypeListener:
                // res.AIGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerListener is populated
            case components.AIGatewayMCPServerTypePassthroughListener:
                // res.AIGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerPassthroughListener is populated
            case components.AIGatewayMCPServerTypeUpstreamServer:
                // res.AIGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerUpstreamServer is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              | Example                                                                                                  |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |                                                                                                          |
| `gatewayID`                                                                                              | `string`                                                                                                 | :heavy_check_mark:                                                                                       | The unique ID of the AI Gateway.                                                                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                                     |
| `createAIGatewayMCPServerRequest`                                                                        | [components.CreateAIGatewayMCPServerRequest](../../models/components/createaigatewaymcpserverrequest.md) | :heavy_check_mark:                                                                                       | N/A                                                                                                      |                                                                                                          |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |                                                                                                          |

### Response

**[*operations.CreateAiGatewayMcpServerResponse](../../models/operations/createaigatewaymcpserverresponse.md), error**

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

## GetAiGatewayMcpServer

Returns the details of a specific MCP Server.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-gateway-mcp-server" method="get" path="/v1/ai-gateways/{gatewayId}/mcp-servers/{mcpServerIdOrName}" -->
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

    res, err := s.AIGatewayMCPServers.GetAiGatewayMcpServer(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayMCPServer != nil {
        switch res.AIGatewayMCPServer.Type {
            case components.AIGatewayMCPServerTypeConversionOnly:
                // res.AIGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerConversionOnly is populated
            case components.AIGatewayMCPServerTypeConversionListener:
                // res.AIGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerConversionListener is populated
            case components.AIGatewayMCPServerTypeListener:
                // res.AIGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerListener is populated
            case components.AIGatewayMCPServerTypePassthroughListener:
                // res.AIGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerPassthroughListener is populated
            case components.AIGatewayMCPServerTypeUpstreamServer:
                // res.AIGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerUpstreamServer is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique ID of the AI Gateway.                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `mcpServerIDOrName`                                      | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the MCP Server.                 | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAiGatewayMcpServerResponse](../../models/operations/getaigatewaymcpserverresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## UpdateAiGatewayMcpServer

Updates the configuration of an existing MCP Server.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-ai-gateway-mcp-server" method="put" path="/v1/ai-gateways/{gatewayId}/mcp-servers/{mcpServerIdOrName}" -->
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

    res, err := s.AIGatewayMCPServers.UpdateAiGatewayMcpServer(ctx, operations.UpdateAiGatewayMcpServerRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        McpServerIDOrName: "my-entity-name",
        UpdateAIGatewayMCPServerRequest: components.CreateUpdateAIGatewayMCPServerRequestUpstreamServer(
            components.CreateAIGatewayMCPServerUpstreamServerOauthAccessToken(
                components.AIGatewayMCPServerBaseACLPropertiesOauthAIGatewayMCPServerUpstreamServerAIGatewayMCPServerBaseACLPropertiesOauth{
                    ACLAttributeType: components.AIGatewayMCPServerBaseACLPropertiesOauthAIGatewayMCPServerUpstreamServerACLAttributeTypeOauthAccessToken,
                    AccessTokenClaimField: "<value>",
                    Type: components.AIGatewayMCPServerBaseACLPropertiesOauthAIGatewayMCPServerUpstreamServerTypeUpstreamServer,
                    Config: components.AIGatewayMCPServerUpstreamServerConfig{
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
                        URL: "https://mcp.internal.kongair.com",
                        ToolsCacheTTLSeconds: 793952,
                    },
                    DisplayName: "Kong Air Flights",
                    Name: "kongair-flights",
                },
            ),
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayMCPServer != nil {
        switch res.AIGatewayMCPServer.Type {
            case components.AIGatewayMCPServerTypeConversionOnly:
                // res.AIGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerConversionOnly is populated
            case components.AIGatewayMCPServerTypeConversionListener:
                // res.AIGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerConversionListener is populated
            case components.AIGatewayMCPServerTypeListener:
                // res.AIGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerListener is populated
            case components.AIGatewayMCPServerTypePassthroughListener:
                // res.AIGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerPassthroughListener is populated
            case components.AIGatewayMCPServerTypeUpstreamServer:
                // res.AIGatewayMCPServer.AIGatewayMCPServerAIGatewayMCPServerUpstreamServer is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdateAiGatewayMcpServerRequest](../../models/operations/updateaigatewaymcpserverrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpdateAiGatewayMcpServerResponse](../../models/operations/updateaigatewaymcpserverresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## DeleteAiGatewayMcpServer

Removes a specific MCP Server from the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-gateway-mcp-server" method="delete" path="/v1/ai-gateways/{gatewayId}/mcp-servers/{mcpServerIdOrName}" -->
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

    res, err := s.AIGatewayMCPServers.DeleteAiGatewayMcpServer(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
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
| `mcpServerIDOrName`                                      | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the MCP Server.                 | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAiGatewayMcpServerResponse](../../models/operations/deleteaigatewaymcpserverresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |