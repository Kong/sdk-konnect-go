# APISpecsPreview

## Overview

### Available Operations

* [PreviewAPISpec](#previewapispec) - Preview API Spec

## PreviewAPISpec

Previews the contents of an API spec as JSON or YAML string.
Specify `Accept: application/json` to retrieve the contents as JSON string.
Specify `Accept: application/yaml` to retrieve the contents as YAML string.
Defaults to JSON string format.


### Example Usage

<!-- UsageSnippet language="go" operationID="preview-api-spec" method="post" path="/v1/api-specs/preview" -->
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

    res, err := s.APISpecsPreview.PreviewAPISpec(ctx, components.PreviewCatalogServiceAPISpec{
        Provider: components.CreatePreviewCatalogServiceAPISpecCreateAPISpecProviderURLAPISpecProvider(
            components.URLAPISpecProvider{
                Type: components.TypeURL,
                Config: components.URLAPISpecProviderConfig{
                    URL: "https://api.petstore.com/v3/openapi.json",
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecContentsPreview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [components.PreviewCatalogServiceAPISpec](../../models/components/previewcatalogserviceapispec.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PreviewAPISpecResponse](../../models/operations/previewapispecresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |