# CatalogServiceAPIMappings

## Overview

Service API mappings represent the link between Service and API entities.
Once an API is mapped to a Service, a rich view of the linked APIs will be presented on the APIs tab of the Catalog Service.
Similarly, Services mapped to an API will be listed on the API overview page under Catalog.
An API may be mapped to multiple services, but it cannot be mapped to the same service twice.
If a mapped API is unlinked from a Service, the mapping will be deleted.


### Available Operations

* [ListServiceMappingsForAPI](#listservicemappingsforapi) - List Service Mappings for an API
* [ListCatalogServiceAPIMappings](#listcatalogserviceapimappings) - List API Mappings for a Service
* [CreateCatalogServiceAPIMapping](#createcatalogserviceapimapping) - Create API Mapping for a Service
* [GetCatalogServiceAPIMapping](#getcatalogserviceapimapping) - Get API Mapping for a Service
* [DeleteCatalogServiceAPIMapping](#deletecatalogserviceapimapping) - Delete API Mapping for a Service

## ListServiceMappingsForAPI

Returns a paginated collection of Service mappings for the given API.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-service-mappings-for-api" method="get" path="/v1/apis/{apiId}/service-mappings" -->
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

    res, err := s.CatalogServiceAPIMappings.ListServiceMappingsForAPI(ctx, operations.ListServiceMappingsForAPIRequest{
        APIID: "d687f4ea-aa04-4157-b446-34519e5b18a7",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCatalogServiceAPIMappingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListServiceMappingsForAPIRequest](../../models/operations/listservicemappingsforapirequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ListServiceMappingsForAPIResponse](../../models/operations/listservicemappingsforapiresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCatalogServiceAPIMappings

Returns a paginated collection of API mappings for the given service.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-catalog-service-api-mappings" method="get" path="/v1/catalog-services/{serviceId}/api-mappings" -->
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

    res, err := s.CatalogServiceAPIMappings.ListCatalogServiceAPIMappings(ctx, operations.ListCatalogServiceAPIMappingsRequest{
        ServiceID: "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCatalogServiceAPIMappingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ListCatalogServiceAPIMappingsRequest](../../models/operations/listcatalogserviceapimappingsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ListCatalogServiceAPIMappingsResponse](../../models/operations/listcatalogserviceapimappingsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateCatalogServiceAPIMapping

Creates a new API mapping for the given service.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-catalog-service-api-mapping" method="post" path="/v1/catalog-services/{serviceId}/api-mappings" -->
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

    res, err := s.CatalogServiceAPIMappings.CreateCatalogServiceAPIMapping(ctx, "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.CreateCatalogServiceAPIMappingBody{
        APIID: "545a8470-5401-4840-b5e1-55b305aa2440",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogServiceAPIMapping != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    | Example                                                                                                        |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |                                                                                                                |
| `serviceID`                                                                                                    | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The `id` of the service.                                                                                       | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                                           |
| `createCatalogServiceAPIMappingBody`                                                                           | [components.CreateCatalogServiceAPIMappingBody](../../models/components/createcatalogserviceapimappingbody.md) | :heavy_check_mark:                                                                                             | Request body schema for creating a new API mapping for catalog service.                                        |                                                                                                                |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |                                                                                                                |

### Response

**[*operations.CreateCatalogServiceAPIMappingResponse](../../models/operations/createcatalogserviceapimappingresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCatalogServiceAPIMapping

Returns information about a specific API mapping for the given service.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-catalog-service-api-mapping" method="get" path="/v1/catalog-services/{serviceId}/api-mappings/{mappingId}" -->
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

    res, err := s.CatalogServiceAPIMappings.GetCatalogServiceAPIMapping(ctx, "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7", "d277faad-ed4e-4c56-a0fb-acce065dee34")
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogServiceAPIMapping != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `serviceID`                                              | *string*                                                 | :heavy_check_mark:                                       | The `id` of the service.                                 | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `mappingID`                                              | *string*                                                 | :heavy_check_mark:                                       | ID of the catalog service API mapping.                   | d277faad-ed4e-4c56-a0fb-acce065dee34                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetCatalogServiceAPIMappingResponse](../../models/operations/getcatalogserviceapimappingresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteCatalogServiceAPIMapping

Deletes a specific API mapping for the given service.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-catalog-service-api-mapping" method="delete" path="/v1/catalog-services/{serviceId}/api-mappings/{mappingId}" -->
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

    res, err := s.CatalogServiceAPIMappings.DeleteCatalogServiceAPIMapping(ctx, "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7", "d277faad-ed4e-4c56-a0fb-acce065dee34")
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
| `serviceID`                                              | *string*                                                 | :heavy_check_mark:                                       | The `id` of the service.                                 | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `mappingID`                                              | *string*                                                 | :heavy_check_mark:                                       | ID of the catalog service API mapping.                   | d277faad-ed4e-4c56-a0fb-acce065dee34                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteCatalogServiceAPIMappingResponse](../../models/operations/deletecatalogserviceapimappingresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |