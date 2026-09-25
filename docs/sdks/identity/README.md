# Identity

## Overview

API for connected apps authorization.

### Available Operations

* [AuthorizeConnectedApp](#authorizeconnectedapp) - Authorize Connected App

## AuthorizeConnectedApp

This endpoint allows a developer to authorize or deny a connected app.

### Example Usage

<!-- UsageSnippet language="go" operationID="authorize-connected-app" method="post" path="/v3/auth/consent" -->
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
            KonnectAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.Identity.AuthorizeConnectedApp(ctx, components.AuthorizeConnectedAppRequest{
        Consent: true,
        Capps: components.ConnectedAppContext{
            State: "xyz789statetoken",
            Source: "connected-app-123",
            Subdomain: "acme",
            ConsentedScopeNames: sdkkonnectgo.Pointer("konnect:read konnect:write"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConsentConnectedAppResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [components.AuthorizeConnectedAppRequest](../../models/components/authorizeconnectedapprequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.AuthorizeConnectedAppResponse](../../models/operations/authorizeconnectedappresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |