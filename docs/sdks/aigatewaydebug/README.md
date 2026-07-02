# AIGatewayDebug

## Overview

AI Gateway debug endpoints.

### Available Operations

* [GetAiGatewayDebugCpOutput](#getaigatewaydebugcpoutput) - Get the CP config output for an AI Gateway

## GetAiGatewayDebugCpOutput

Internal / privileged endpoint. Returns the control-plane configuration Koko would deliver to the given AI
Gateway's data planes, as JSON by default or as YAML when the Accept header requests application/yaml.

Requires a privileged internal service-client token; the organization is derived from the gateway.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-gateway-debug-cp-output" method="get" path="/v1/ai-gateways/{gatewayId}/debug-cp-output" -->
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

    res, err := s.AIGatewayDebug.GetAiGatewayDebugCpOutput(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique ID of the AI Gateway.                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAiGatewayDebugCpOutputResponse](../../models/operations/getaigatewaydebugcpoutputresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |