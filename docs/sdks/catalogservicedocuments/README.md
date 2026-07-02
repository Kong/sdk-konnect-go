# CatalogServiceDocuments

## Overview

Attach documentation for your services to communicate it's architecture, support details, release notes, and usage examples, to help others understand and integrate with the service.


### Available Operations

* [CreateCatalogServiceDocument](#createcatalogservicedocument) - Create Service Document
* [ListCatalogServiceDocuments](#listcatalogservicedocuments) - List Catalog Service Documents
* [FetchCatalogServiceDocument](#fetchcatalogservicedocument) - Get a Catalog Service Document
* [UpdateCatalogServiceDocument](#updatecatalogservicedocument) - Update Catalog Service Document
* [DeleteCatalogServiceDocument](#deletecatalogservicedocument) - Delete Catalog Service Document
* [MoveCatalogServiceDocument](#movecatalogservicedocument) - Move Catalog Service Document

## CreateCatalogServiceDocument

Creates a service document

### Example Usage

<!-- UsageSnippet language="go" operationID="create-catalog-service-document" method="post" path="/v1/catalog-services/{serviceId}/documents" -->
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

    res, err := s.CatalogServiceDocuments.CreateCatalogServiceDocument(ctx, "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.CreateServiceDocumentPayload{
        Slug: "path-for-seo",
        Title: "<value>",
        ParentDocumentID: sdkkonnectgo.Pointer("dd4e1b98-3629-4dd3-acc0-759a726ffee2"),
        Content: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ServiceDocumentContentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |                                                                                                    |
| `serviceID`                                                                                        | `string`                                                                                           | :heavy_check_mark:                                                                                 | The `id` of the service.                                                                           | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                               |
| `createServiceDocumentPayload`                                                                     | [components.CreateServiceDocumentPayload](../../models/components/createservicedocumentpayload.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |                                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |                                                                                                    |

### Response

**[*operations.CreateCatalogServiceDocumentResponse](../../models/operations/createcatalogservicedocumentresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.BaseError         | 500                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCatalogServiceDocuments

Returns a paginated collection of documents attached to the given service.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-catalog-service-documents" method="get" path="/v1/catalog-services/{serviceId}/documents" -->
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

    res, err := s.CatalogServiceDocuments.ListCatalogServiceDocuments(ctx, operations.ListCatalogServiceDocumentsRequest{
        ServiceID: "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListCatalogServiceDocumentsRequest](../../models/operations/listcatalogservicedocumentsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ListCatalogServiceDocumentsResponse](../../models/operations/listcatalogservicedocumentsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.BaseError         | 500                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchCatalogServiceDocument

Fetched a service document.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-catalog-service-document" method="get" path="/v1/catalog-services/{serviceId}/documents/{documentId}" -->
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

    res, err := s.CatalogServiceDocuments.FetchCatalogServiceDocument(ctx, "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7", "7b453ae4-66f5-4530-88b0-0b6e34072325")
    if err != nil {
        log.Fatal(err)
    }
    if res.ServiceDocumentContentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `serviceID`                                              | `string`                                                 | :heavy_check_mark:                                       | The `id` of the service.                                 | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `documentID`                                             | `string`                                                 | :heavy_check_mark:                                       | The `id` of the document.                                | 7b453ae4-66f5-4530-88b0-0b6e34072325                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchCatalogServiceDocumentResponse](../../models/operations/fetchcatalogservicedocumentresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.BaseError         | 500                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateCatalogServiceDocument

Updates a service document.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-catalog-service-document" method="patch" path="/v1/catalog-services/{serviceId}/documents/{documentId}" -->
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

    res, err := s.CatalogServiceDocuments.UpdateCatalogServiceDocument(ctx, operations.UpdateCatalogServiceDocumentRequest{
        ServiceID: "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7",
        DocumentID: "7b453ae4-66f5-4530-88b0-0b6e34072325",
        UpdateServiceDocumentPayload: components.UpdateServiceDocumentPayload{
            Slug: sdkkonnectgo.Pointer("path-for-seo"),
            ParentDocumentID: sdkkonnectgo.Pointer("dd4e1b98-3629-4dd3-acc0-759a726ffee2"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ServiceDocumentContentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.UpdateCatalogServiceDocumentRequest](../../models/operations/updatecatalogservicedocumentrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.UpdateCatalogServiceDocumentResponse](../../models/operations/updatecatalogservicedocumentresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.BaseError         | 500                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteCatalogServiceDocument

Deletes a service document.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-catalog-service-document" method="delete" path="/v1/catalog-services/{serviceId}/documents/{documentId}" -->
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

    res, err := s.CatalogServiceDocuments.DeleteCatalogServiceDocument(ctx, "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7", "7b453ae4-66f5-4530-88b0-0b6e34072325")
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
| `documentID`                                             | `string`                                                 | :heavy_check_mark:                                       | The `id` of the document.                                | 7b453ae4-66f5-4530-88b0-0b6e34072325                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteCatalogServiceDocumentResponse](../../models/operations/deletecatalogservicedocumentresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.BaseError         | 500                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## MoveCatalogServiceDocument

Moves a service document.

### Example Usage

<!-- UsageSnippet language="go" operationID="move-catalog-service-document" method="post" path="/v1/catalog-services/{serviceId}/documents/{documentId}/move" -->
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

    res, err := s.CatalogServiceDocuments.MoveCatalogServiceDocument(ctx, operations.MoveCatalogServiceDocumentRequest{
        ServiceID: "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7",
        DocumentID: "7b453ae4-66f5-4530-88b0-0b6e34072325",
        MoveDocumentRequestPayload: components.MoveDocumentRequestPayload{
            ParentDocumentID: sdkkonnectgo.Pointer("dd4e1b98-3629-4dd3-acc0-759a726ffee2"),
            Index: sdkkonnectgo.Pointer[int64](1),
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.MoveCatalogServiceDocumentRequest](../../models/operations/movecatalogservicedocumentrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.MoveCatalogServiceDocumentResponse](../../models/operations/movecatalogservicedocumentresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.BaseError         | 500                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |