# APIPackageDocumentation

## Overview

### Available Operations

* [CreateAPIPackageDocument](#createapipackagedocument) - Create API Package Document
* [ListAPIPackageDocuments](#listapipackagedocuments) - List API Package Documents
* [FetchAPIPackageDocument](#fetchapipackagedocument) - Get an API Package Document
* [UpdateAPIPackageDocument](#updateapipackagedocument) - Update API Package Document
* [DeleteAPIPackageDocument](#deleteapipackagedocument) - Delete API Package Documentation
* [MoveAPIPackageDocument](#moveapipackagedocument) - Move API Package Documentation

## CreateAPIPackageDocument

Publish a new document attached to an API Package.

All configuration options may be provided in the frontmatter section of `content`. If you set values in both the `POST` request _and_ in the frontmatter, the values in the `POST` request will take precedence.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-api-package-document" method="post" path="/v3/api-packages/{packageId}/documents" -->
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

    res, err := s.APIPackageDocumentation.CreateAPIPackageDocument(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIDocumentRequest{
        Content: "<value>",
        Title: sdkkonnectgo.Pointer("API Document"),
        Slug: sdkkonnectgo.Pointer("api-document"),
        ParentDocumentID: nil,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIPackageDocumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `packageID`                                                                                | *string*                                                                                   | :heavy_check_mark:                                                                         | The UUID API Package identifier                                                            | 9f5061ce-78f6-4452-9108-ad7c02821fd5                                                       |
| `createAPIDocumentRequest`                                                                 | [components.CreateAPIDocumentRequest](../../models/components/createapidocumentrequest.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |

### Response

**[*operations.CreateAPIPackageDocumentResponse](../../models/operations/createapipackagedocumentresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.BadRequestError           | 400                                 | application/problem+json            |
| sdkerrors.UnauthorizedError         | 401                                 | application/problem+json            |
| sdkerrors.ForbiddenError            | 403                                 | application/problem+json            |
| sdkerrors.NotFoundError             | 404                                 | application/problem+json            |
| sdkerrors.APISlugConflict           | 409                                 | application/problem+json            |
| sdkerrors.UnsupportedMediaTypeError | 415                                 | application/problem+json            |
| sdkerrors.SDKError                  | 4XX, 5XX                            | \*/\*                               |

## ListAPIPackageDocuments

Returns a collection of all documents for an API package.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-api-package-documents" method="get" path="/v3/api-packages/{packageId}/documents" -->
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

    res, err := s.APIPackageDocumentation.ListAPIPackageDocuments(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAPIPackageDocumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       | Example                                                                                           |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |                                                                                                   |
| `packageID`                                                                                       | *string*                                                                                          | :heavy_check_mark:                                                                                | The UUID API Package identifier                                                                   | 9f5061ce-78f6-4452-9108-ad7c02821fd5                                                              |
| `filter`                                                                                          | [*components.APIDocumentFilterParameters](../../models/components/apidocumentfilterparameters.md) | :heavy_minus_sign:                                                                                | Filters API Documents in the response.                                                            |                                                                                                   |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |                                                                                                   |

### Response

**[*operations.ListAPIPackageDocumentsResponse](../../models/operations/listapipackagedocumentsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchAPIPackageDocument

Returns a document for the API Package.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-api-package-document" method="get" path="/v3/api-packages/{packageId}/documents/{documentId}" -->
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

    res, err := s.APIPackageDocumentation.FetchAPIPackageDocument(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", "de5c9818-be5c-42e6-b514-e3d4bc30ddeb")
    if err != nil {
        log.Fatal(err)
    }
    if res.APIPackageDocumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `packageID`                                              | *string*                                                 | :heavy_check_mark:                                       | The UUID API Package identifier                          | 9f5061ce-78f6-4452-9108-ad7c02821fd5                     |
| `documentID`                                             | *string*                                                 | :heavy_check_mark:                                       | The document identifier related to the API               | de5c9818-be5c-42e6-b514-e3d4bc30ddeb                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchAPIPackageDocumentResponse](../../models/operations/fetchapipackagedocumentresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateAPIPackageDocument

Updates a document for an API Package.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-api-package-document" method="patch" path="/v3/api-packages/{packageId}/documents/{documentId}" -->
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

    res, err := s.APIPackageDocumentation.UpdateAPIPackageDocument(ctx, operations.UpdateAPIPackageDocumentRequest{
        PackageID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        DocumentID: "de5c9818-be5c-42e6-b514-e3d4bc30ddeb",
        APIDocument: components.APIDocument{
            Title: sdkkonnectgo.Pointer("API Document"),
            Slug: sdkkonnectgo.Pointer("api-document"),
            ParentDocumentID: nil,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIPackageDocumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdateAPIPackageDocumentRequest](../../models/operations/updateapipackagedocumentrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpdateAPIPackageDocumentResponse](../../models/operations/updateapipackagedocumentresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.BadRequestError           | 400                                 | application/problem+json            |
| sdkerrors.UnauthorizedError         | 401                                 | application/problem+json            |
| sdkerrors.ForbiddenError            | 403                                 | application/problem+json            |
| sdkerrors.NotFoundError             | 404                                 | application/problem+json            |
| sdkerrors.APISlugConflict           | 409                                 | application/problem+json            |
| sdkerrors.UnsupportedMediaTypeError | 415                                 | application/problem+json            |
| sdkerrors.SDKError                  | 4XX, 5XX                            | \*/\*                               |

## DeleteAPIPackageDocument

Removes a document from an API Package.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-api-package-document" method="delete" path="/v3/api-packages/{packageId}/documents/{documentId}" -->
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

    res, err := s.APIPackageDocumentation.DeleteAPIPackageDocument(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", "de5c9818-be5c-42e6-b514-e3d4bc30ddeb")
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
| `packageID`                                              | *string*                                                 | :heavy_check_mark:                                       | The UUID API Package identifier                          | 9f5061ce-78f6-4452-9108-ad7c02821fd5                     |
| `documentID`                                             | *string*                                                 | :heavy_check_mark:                                       | The document identifier related to the API               | de5c9818-be5c-42e6-b514-e3d4bc30ddeb                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAPIPackageDocumentResponse](../../models/operations/deleteapipackagedocumentresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## MoveAPIPackageDocument

This api allows the user to move a document within the document tree using the parameters parent_document_id and index. If parent_document_id is not provided, the document will be placed at the top level of the document tree. index represents a zero-indexed document order relative to its siblings under the same parent. For example, if we want to put the document at top level in first position we would send parent_document_id: null and index: 0. This api also supports using a negative index to count backwards from the end of the document list, which means you can put the document in last position by using index: -1.

### Example Usage

<!-- UsageSnippet language="go" operationID="move-api-package-document" method="post" path="/v3/api-packages/{packageId}/documents/{documentId}/move" -->
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

    res, err := s.APIPackageDocumentation.MoveAPIPackageDocument(ctx, operations.MoveAPIPackageDocumentRequest{
        PackageID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        DocumentID: "de5c9818-be5c-42e6-b514-e3d4bc30ddeb",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.MoveAPIPackageDocumentRequest](../../models/operations/moveapipackagedocumentrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.MoveAPIPackageDocumentResponse](../../models/operations/moveapipackagedocumentresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |