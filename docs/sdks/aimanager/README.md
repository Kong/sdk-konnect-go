# AIManager

## Overview

### Available Operations

* [ListVirtualKeys](#listvirtualkeys) - Get Control Plane Virtual Keys

## ListVirtualKeys

Returns the Virtual Keys defined on a control plane.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-virtual-keys" method="get" path="/v2/control-planes/{controlPlaneId}/ai-manager/virtual-keys" -->
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

    res, err := s.AIManager.ListVirtualKeys(ctx, operations.ListVirtualKeysRequest{
        ControlPlaneID: "1279e60b-cbe9-415c-acc4-085ecd051652",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIManagerVirtualKeysResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListVirtualKeysRequest](../../models/operations/listvirtualkeysrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListVirtualKeysResponse](../../models/operations/listvirtualkeysresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequestError     | 400                           | application/problem+json      |
| sdkerrors.UnauthorizedError   | 401                           | application/problem+json      |
| sdkerrors.ForbiddenError      | 403                           | application/problem+json      |
| sdkerrors.NotFoundError       | 404                           | application/problem+json      |
| sdkerrors.BaseError           | 500                           | application/problem+json      |
| sdkerrors.NotImplementedError | 501                           | application/problem+json      |
| sdkerrors.ServiceUnavailable  | 503                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |