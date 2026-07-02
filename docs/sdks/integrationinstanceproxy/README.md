# IntegrationInstanceProxy

## Overview

The integration instance proxy allows requests to be made to the authorized integration instance.
The APIs which are allowed to be proxied to the integration instances are defined in the `proxies` section of the Integration.


### Available Operations

* [IntegrationInstanceProxyRequest](#integrationinstanceproxyrequest) - Integration Instance Proxy Request

## IntegrationInstanceProxyRequest

Proxies a request to an external API for the integration instance.

### Example Usage

<!-- UsageSnippet language="go" operationID="integration-instance-proxy-request" method="post" path="/v1/integration-instances/{integrationInstance}/proxy" -->
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

    res, err := s.IntegrationInstanceProxy.IntegrationInstanceProxyRequest(ctx, "3321067c-407f-417b-8aac-cd2c5ec92d85", components.IntegrationInstanceProxyBody{
        Method: components.IntegrationInstanceProxyBodyMethodGet,
        BaseURL: "https://api.pagerduty.com",
        Path: "/incidents",
        Body: map[string]any{
            "name": "Foo",
        },
        Query: map[string]any{
            "limit": 5,
            "offset": 10,
        },
        Headers: map[string]any{
            "content-type": "application/json",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoXXApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |                                                                                                    |
| `integrationInstance`                                                                              | `string`                                                                                           | :heavy_check_mark:                                                                                 | The `id` of the integration instance.                                                              | 3321067c-407f-417b-8aac-cd2c5ec92d85                                                               |
| `integrationInstanceProxyBody`                                                                     | [components.IntegrationInstanceProxyBody](../../models/components/integrationinstanceproxybody.md) | :heavy_check_mark:                                                                                 | Request body schema for an integration instance proxy request.                                     |                                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |                                                                                                    |

### Response

**[*operations.IntegrationInstanceProxyRequestResponse](../../models/operations/integrationinstanceproxyrequestresponse.md), error**

### Errors

| Error Type                                                                    | Status Code                                                                   | Content Type                                                                  |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| sdkerrors.IntegrationInstanceProxyRequestResponseBody                         | 400                                                                           | application/problem+json                                                      |
| sdkerrors.UnauthorizedError                                                   | 401                                                                           | application/problem+json                                                      |
| sdkerrors.ForbiddenError                                                      | 403                                                                           | application/problem+json                                                      |
| sdkerrors.IntegrationInstanceProxyRequestIntegrationInstanceProxyResponseBody | 404                                                                           | application/problem+json                                                      |
| sdkerrors.SDKError                                                            | 4XX, 5XX                                                                      | \*/\*                                                                         |