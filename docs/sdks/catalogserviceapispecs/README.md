# CatalogServiceAPISpecs

## Overview

Attach API specifications which describe the interface and behavior of the service.
These specifications describe service API constracts for internal or external consumption.
You may attach an API specification by uploading a file, providing a link to the specification, or my mapping a resource from a API spec compliant integration.


### Available Operations

* [CreateCatalogServiceAPISpec](#createcatalogserviceapispec) - Create API spec
* [ListCatalogServiceAPISpecs](#listcatalogserviceapispecs) - List Catalog Service API Specs
* [PreviewCatalogServiceAPISpec](#previewcatalogserviceapispec) - Preview API Spec
* [FetchCatalogServiceAPISpec](#fetchcatalogserviceapispec) - Get an API spec
* [UpdateCatalogServiceAPISpec](#updatecatalogserviceapispec) - Update API Spec
* [DeleteCatalogServiceAPISpec](#deletecatalogserviceapispec) - Delete API Spec
* [FetchCatalogServiceAPISpecContents](#fetchcatalogserviceapispeccontents) - Get an API Spec Contents

## CreateCatalogServiceAPISpec

Creates an API spec associated to the given service.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-catalog-service-api-spec" method="post" path="/v1/catalog-services/{id}/api-specs" -->
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

    res, err := s.CatalogServiceAPISpecs.CreateCatalogServiceAPISpec(ctx, "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.CreateCatalogServiceAPISpec{
        Name: "Pet Store",
        Description: sdkkonnectgo.Pointer("A sample API that uses a pet store as an example to demonstrate features in the OpenAPI specification\n"),
        Provider: components.CreateCreateAPISpecProviderResourceBoundIntegrationAPISpecProviderPayload(
            components.ResourceBoundIntegrationAPISpecProviderPayload{
                Type: "<value>",
                Config: components.ResourceBoundIntegrationAPISpecProviderPayloadConfig{
                    ResourceID: "IqkHvMdyHukxcwAs",
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogServiceAPISpec != nil {
        switch res.CatalogServiceAPISpec.Provider.Type {
            case components.APISpecProviderTypeURLAPISpecProvider:
                // res.CatalogServiceAPISpec.Provider.URLAPISpecProvider is populated
            case components.APISpecProviderTypeRawAPISpecProvider:
                // res.CatalogServiceAPISpec.Provider.RawAPISpecProvider is populated
            case components.APISpecProviderTypeIntegrationAPISpecProvider:
                // res.CatalogServiceAPISpec.Provider.IntegrationAPISpecProvider is populated
            case components.APISpecProviderTypeResourceBoundIntegrationAPISpecProvider:
                // res.CatalogServiceAPISpec.Provider.ResourceBoundIntegrationAPISpecProvider is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `id`                                                                                             | `string`                                                                                         | :heavy_check_mark:                                                                               | The `id` of the service.                                                                         | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                             |
| `createCatalogServiceAPISpec`                                                                    | [components.CreateCatalogServiceAPISpec](../../models/components/createcatalogserviceapispec.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*operations.CreateCatalogServiceAPISpecResponse](../../models/operations/createcatalogserviceapispecresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCatalogServiceAPISpecs

Returns a paginated collection of API Specs for the given service.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-catalog-service-api-specs" method="get" path="/v1/catalog-services/{id}/api-specs" -->
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

    res, err := s.CatalogServiceAPISpecs.ListCatalogServiceAPISpecs(ctx, operations.ListCatalogServiceAPISpecsRequest{
        ID: "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCatalogServiceAPISpecsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListCatalogServiceAPISpecsRequest](../../models/operations/listcatalogserviceapispecsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListCatalogServiceAPISpecsResponse](../../models/operations/listcatalogserviceapispecsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PreviewCatalogServiceAPISpec

Previews the contents of an API spec as JSON or YAML.
Specify `Accept: application/json` to retrieve the contents as JSON.
Specify `Accept: application/yaml` to retrieve the contents as YAML.
Defaults to JSON format.


### Example Usage

<!-- UsageSnippet language="go" operationID="preview-catalog-service-api-spec" method="post" path="/v1/catalog-services/{serviceId}/api-specs/preview" -->
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

    res, err := s.CatalogServiceAPISpecs.PreviewCatalogServiceAPISpec(ctx, "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.PreviewCatalogServiceAPISpec{
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
    if res.APISpecContents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |                                                                                                    |
| `serviceID`                                                                                        | `string`                                                                                           | :heavy_check_mark:                                                                                 | The `id` of the service.                                                                           | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                               |
| `previewCatalogServiceAPISpec`                                                                     | [components.PreviewCatalogServiceAPISpec](../../models/components/previewcatalogserviceapispec.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |                                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |                                                                                                    |

### Response

**[*operations.PreviewCatalogServiceAPISpecResponse](../../models/operations/previewcatalogserviceapispecresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchCatalogServiceAPISpec

Returns the API spec for the given service.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-catalog-service-api-spec" method="get" path="/v1/catalog-services/{serviceId}/api-specs/{apiSpecId}" -->
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

    res, err := s.CatalogServiceAPISpecs.FetchCatalogServiceAPISpec(ctx, "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7", "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogServiceAPISpec != nil {
        switch res.CatalogServiceAPISpec.Provider.Type {
            case components.APISpecProviderTypeURLAPISpecProvider:
                // res.CatalogServiceAPISpec.Provider.URLAPISpecProvider is populated
            case components.APISpecProviderTypeRawAPISpecProvider:
                // res.CatalogServiceAPISpec.Provider.RawAPISpecProvider is populated
            case components.APISpecProviderTypeIntegrationAPISpecProvider:
                // res.CatalogServiceAPISpec.Provider.IntegrationAPISpecProvider is populated
            case components.APISpecProviderTypeResourceBoundIntegrationAPISpecProvider:
                // res.CatalogServiceAPISpec.Provider.ResourceBoundIntegrationAPISpecProvider is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `serviceID`                                              | `string`                                                 | :heavy_check_mark:                                       | The `id` of the service.                                 | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `apiSpecID`                                              | `string`                                                 | :heavy_check_mark:                                       | The `id` of the service.                                 | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchCatalogServiceAPISpecResponse](../../models/operations/fetchcatalogserviceapispecresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateCatalogServiceAPISpec

Updates the API spec for the given service.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-catalog-service-api-spec" method="patch" path="/v1/catalog-services/{serviceId}/api-specs/{apiSpecId}" -->
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

    res, err := s.CatalogServiceAPISpecs.UpdateCatalogServiceAPISpec(ctx, operations.UpdateCatalogServiceAPISpecRequest{
        ServiceID: "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7",
        APISpecID: "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7",
        UpdateCatalogServiceAPISpec: components.UpdateCatalogServiceAPISpec{
            Name: sdkkonnectgo.Pointer("Pet Store"),
            Description: sdkkonnectgo.Pointer("A sample API that uses a pet store as an example to demonstrate features in the OpenAPI specification\n"),
            Provider: sdkkonnectgo.Pointer(components.CreateUpdateCatalogServiceAPISpecCreateAPISpecProviderResourceBoundIntegrationAPISpecProviderPayload(
                components.ResourceBoundIntegrationAPISpecProviderPayload{
                    Type: "<value>",
                    Config: components.ResourceBoundIntegrationAPISpecProviderPayloadConfig{
                        ResourceID: "IqkHvMdyHukxcwAs",
                    },
                },
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogServiceAPISpec != nil {
        switch res.CatalogServiceAPISpec.Provider.Type {
            case components.APISpecProviderTypeURLAPISpecProvider:
                // res.CatalogServiceAPISpec.Provider.URLAPISpecProvider is populated
            case components.APISpecProviderTypeRawAPISpecProvider:
                // res.CatalogServiceAPISpec.Provider.RawAPISpecProvider is populated
            case components.APISpecProviderTypeIntegrationAPISpecProvider:
                // res.CatalogServiceAPISpec.Provider.IntegrationAPISpecProvider is populated
            case components.APISpecProviderTypeResourceBoundIntegrationAPISpecProvider:
                // res.CatalogServiceAPISpec.Provider.ResourceBoundIntegrationAPISpecProvider is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.UpdateCatalogServiceAPISpecRequest](../../models/operations/updatecatalogserviceapispecrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.UpdateCatalogServiceAPISpecResponse](../../models/operations/updatecatalogserviceapispecresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteCatalogServiceAPISpec

Deletes the API spec for the given service.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-catalog-service-api-spec" method="delete" path="/v1/catalog-services/{serviceId}/api-specs/{apiSpecId}" -->
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

    res, err := s.CatalogServiceAPISpecs.DeleteCatalogServiceAPISpec(ctx, "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7", "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
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
| `serviceID`                                              | `string`                                                 | :heavy_check_mark:                                       | The `id` of the service.                                 | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `apiSpecID`                                              | `string`                                                 | :heavy_check_mark:                                       | The `id` of the service.                                 | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteCatalogServiceAPISpecResponse](../../models/operations/deletecatalogserviceapispecresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchCatalogServiceAPISpecContents

Returns the API spec for the given service.
Specify `Accept: application/json` to retrieve the contents as JSON.
Specify `Accept: application/yaml` to retrieve the contents as YAML.
Defaults to JSON format.


### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-catalog-service-api-spec-contents" method="get" path="/v1/catalog-services/{serviceId}/api-specs/{apiSpecId}/contents" -->
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

    res, err := s.CatalogServiceAPISpecs.FetchCatalogServiceAPISpecContents(ctx, "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7", "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecContents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `serviceID`                                              | `string`                                                 | :heavy_check_mark:                                       | The `id` of the service.                                 | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `apiSpecID`                                              | `string`                                                 | :heavy_check_mark:                                       | The `id` of the service.                                 | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchCatalogServiceAPISpecContentsResponse](../../models/operations/fetchcatalogserviceapispeccontentsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |