# ContextSourceContents

## Overview

### Available Operations

* [GetContextSourceContents](#getcontextsourcecontents) - Get Context Source Contents

## GetContextSourceContents

Retrieve the contents for a Context Source.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-context-source-contents" method="get" path="/v1/context-sources/{sourceId}/contents" -->
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

    res, err := s.ContextSourceContents.GetContextSourceContents(ctx, "f683b838-bbc6-4927-aca4-3c2be3741b7a")
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
| `sourceID`                                               | `string`                                                 | :heavy_check_mark:                                       | The ID of the MCP resource.                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetContextSourceContentsResponse](../../models/operations/getcontextsourcecontentsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |