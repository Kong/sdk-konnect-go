# MCPResourceContents

## Overview

### Available Operations

* [GetMcpResourceContents](#getmcpresourcecontents) - Get MCP Resource Contents

## GetMcpResourceContents

Retrieve the contents for an MCP resource.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mcp-resource-contents" method="get" path="/v1/mcp-resources/{resourceId}/contents" -->
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

    res, err := s.MCPResourceContents.GetMcpResourceContents(ctx, "69b02904-b94c-4718-a1b2-5341c111f432")
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPAPISpec != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `resourceID`                                             | `string`                                                 | :heavy_check_mark:                                       | The ID of the MCP resource.                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetMcpResourceContentsResponse](../../models/operations/getmcpresourcecontentsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |