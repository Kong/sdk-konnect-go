# CatalogResources

## Overview

Resources are entities discovered from integration instances and are intended to be mapped to the relevant services in the catalog.
Once a resource has been mapped to a service, a rich view of this resource will be displayed when viewing your service.

Any resources which you would like to ignore and hide can be archived.
Note that archiving a resource will **remove** any mappings it has to services in the catalog.

You can create Resources using only the properties required to identify the resource in the third-party system.
Additional data will be hydrated if the integration that the Resource is linked to has been provided with authentication credentials.
Resources which have not yet been hydrated may still be mapped to a Catalog Service.


### Available Operations

* [ListResources](#listresources) - List Resources
* [FetchResource](#fetchresource) - Get a Resource
* [InitializeResource](#initializeresource) - Initialize Resource
* [UpsertResources](#upsertresources) - Upsert Resources
* [DeleteResources](#deleteresources) - Delete Resources
* [UpdateResource](#updateresource) - Update Resource

## ListResources

Returns a paginated collection of resources.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-resources" method="get" path="/v1/resources" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/types"
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

    res, err := s.CatalogResources.ListResources(ctx, operations.ListResourcesRequest{
        Page: &components.CursorPageParameters{
            Size: sdkkonnectgo.Pointer[int64](10),
            After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
            Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        },
        Filter: &components.CatalogResourceFilterParameters{
            ServiceAssociations: sdkkonnectgo.Pointer(components.CreateNumericFieldFilterNumber(
                21.0,
            )),
            CreatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
                components.DateTimeFieldLTFilter{
                    Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            UpdatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
                components.DateTimeFieldLTFilter{
                    Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCatalogResourcesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListResourcesRequest](../../models/operations/listresourcesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ListResourcesResponse](../../models/operations/listresourcesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchResource

Fetches a resource by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-resource" method="get" path="/v1/resources/{id}" -->
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

    res, err := s.CatalogResources.FetchResource(ctx, "AflTNLY0tTQhv2my")
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogResource != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The `id` of the resource.                                | AflTNLY0tTQhv2my                                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchResourceResponse](../../models/operations/fetchresourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## InitializeResource

Initializes a placeholder resource. This operation is typically used in GitOps workflows or other infrastructure-as-code setups where a resource is declared before it is fully discovered.
The initialized resource acts as a stub — it contains minimal identifying metadata and does not yet include the full set of attributes typically populated by integration-based discovery.
After initialization, the resource will be automatically hydrated with additional data by catalog integrations responsible for discovering and maintaining its up-to-date state.
This API is not intended for direct resource management, but rather for signaling the intent to track a given resource, allowing integrations to take over and populate it asynchronously.


### Example Usage

<!-- UsageSnippet language="go" operationID="initialize-resource" method="post" path="/v1/integration-instances/{integrationInstanceId}/resources" -->
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

    res, err := s.CatalogResources.InitializeResource(ctx, "3f51fa25-310a-421d-bd1a-007f859021a3", components.InitializeCatalogResource{
        Type: "gateway_svc",
        Config: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogResource != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `integrationInstanceID`                                                                      | `string`                                                                                     | :heavy_check_mark:                                                                           | The `id` of the integration instance.                                                        | 3f51fa25-310a-421d-bd1a-007f859021a3                                                         |
| `initializeCatalogResource`                                                                  | [components.InitializeCatalogResource](../../models/components/initializecatalogresource.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

### Response

**[*operations.InitializeResourceResponse](../../models/operations/initializeresourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpsertResources

Upserts resources. If a resource with the same `externalId` already exists for the given integration instance, it will be updated; otherwise, a new resource will be created.
When a resource's backing definition in the integration has `integration_data_schema` populated, the `integration_data` value provided for that resource will be validated against `integration_data_schema` during upsert operations.
When a resource's backing definition in the integration has no `integration_data_schema` populated (i.e. null), no validation will be performed on the `integration_data` value provided for that resource.
This API is typically used by service catalog integration applications to report discovered resources back to the catalog.


### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-resources" method="put" path="/v1/integration-instances/{integrationInstanceId}/resources" -->
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

    res, err := s.CatalogResources.UpsertResources(ctx, "3f51fa25-310a-421d-bd1a-007f859021a3", components.UpsertCatalogResourcesRequest{
        Data: []components.UpsertCatalogResource{
            components.UpsertCatalogResource{
                Name: "<value>",
                Type: "gateway_svc",
                Config: "<value>",
                IntegrationData: map[string]any{

                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `integrationInstanceID`                                                                              | `string`                                                                                             | :heavy_check_mark:                                                                                   | The `id` of the integration instance.                                                                | 3f51fa25-310a-421d-bd1a-007f859021a3                                                                 |
| `upsertCatalogResourcesRequest`                                                                      | [components.UpsertCatalogResourcesRequest](../../models/components/upsertcatalogresourcesrequest.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.UpsertResourcesResponse](../../models/operations/upsertresourcesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteResources

Delete up to 100 resources which belong to the related integration instance.
This API is intended only for service catalog integration applications.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-resources" method="post" path="/v1/integration-instances/{id}/resources/bulk-delete" -->
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

    res, err := s.CatalogResources.DeleteResources(ctx, "3f51fa25-310a-421d-bd1a-007f859021a3", components.DeleteCatalogResources{
        Data: []components.DeleteCatalogResourcesData{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `id`                                                                                   | `string`                                                                               | :heavy_check_mark:                                                                     | The `id` of the integration instance.                                                  | 3f51fa25-310a-421d-bd1a-007f859021a3                                                   |
| `deleteCatalogResources`                                                               | [components.DeleteCatalogResources](../../models/components/deletecatalogresources.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

### Response

**[*operations.DeleteResourcesResponse](../../models/operations/deleteresourcesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateResource

Updates a resource.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-resource" method="patch" path="/v1/integration-instances/{integrationInstanceId}/resources/{resourceId}" -->
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

    res, err := s.CatalogResources.UpdateResource(ctx, operations.UpdateResourceRequest{
        IntegrationInstanceID: "3f51fa25-310a-421d-bd1a-007f859021a3",
        ResourceID: "AflTNLY0tTQhv2my",
        UpdateCatalogResource: components.UpdateCatalogResource{
            Archived: sdkkonnectgo.Pointer(false),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogResource != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateResourceRequest](../../models/operations/updateresourcerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateResourceResponse](../../models/operations/updateresourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |