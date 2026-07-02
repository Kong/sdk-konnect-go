# Nodes

## Overview

### Available Operations

* [UpsertNode](#upsertnode) - Create or update a node

## UpsertNode

Create or update a node.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-node" method="put" path="/v2/control-planes/{controlPlaneId}/nodes/{nodeId}" -->
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

    res, err := s.Nodes.UpsertNode(ctx, operations.UpsertNodeRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        NodeID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        UpsertNode: components.UpsertNode{
            Hostname: sdkkonnectgo.Pointer("node1.my.domain.com"),
            Type: sdkkonnectgo.Pointer("knep"),
            Version: sdkkonnectgo.Pointer("1.0"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetNode != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.UpsertNodeRequest](../../models/operations/upsertnoderequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.UpsertNodeResponse](../../models/operations/upsertnoderesponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |