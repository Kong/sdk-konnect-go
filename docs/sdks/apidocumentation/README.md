# APIDocumentation

## Overview

### Available Operations

* [CreateAPIDocument](#createapidocument) - Create API Document
* [ListAPIDocuments](#listapidocuments) - List API Documents
* [FetchAPIDocument](#fetchapidocument) - Get an API Document
* [UpdateAPIDocument](#updateapidocument) - Update API Document
* [DeleteAPIDocument](#deleteapidocument) - Delete API Documentation
* [MoveAPIDocument](#moveapidocument) - Move API Documentation

## CreateAPIDocument

Publish a new document attached to an API.

All configuration options may be provided in the frontmatter section of `content`. If you set values in both the `POST` request _and_ in the frontmatter, the values in the `POST` request will take precedence.


### Example Usage: ApiForbiddenExample

<!-- UsageSnippet language="go" operationID="create-api-document" method="post" path="/v3/apis/{apiId}/documents" example="ApiForbiddenExample" -->
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

    res, err := s.APIDocumentation.CreateAPIDocument(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIDocumentRequest{
        Content: "<value>",
        Title: sdkkonnectgo.Pointer("API Document"),
        Slug: sdkkonnectgo.Pointer("api-document"),
        ParentDocumentID: nil,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiNameBadRequestExample

<!-- UsageSnippet language="go" operationID="create-api-document" method="post" path="/v3/apis/{apiId}/documents" example="ApiNameBadRequestExample" -->
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

    res, err := s.APIDocumentation.CreateAPIDocument(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIDocumentRequest{
        Content: "<value>",
        Title: sdkkonnectgo.Pointer("API Document"),
        Slug: sdkkonnectgo.Pointer("api-document"),
        ParentDocumentID: nil,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiNotFoundExample

<!-- UsageSnippet language="go" operationID="create-api-document" method="post" path="/v3/apis/{apiId}/documents" example="ApiNotFoundExample" -->
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

    res, err := s.APIDocumentation.CreateAPIDocument(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIDocumentRequest{
        Content: "<value>",
        Title: sdkkonnectgo.Pointer("API Document"),
        Slug: sdkkonnectgo.Pointer("api-document"),
        ParentDocumentID: nil,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSlugConflictExample

<!-- UsageSnippet language="go" operationID="create-api-document" method="post" path="/v3/apis/{apiId}/documents" example="ApiSlugConflictExample" -->
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

    res, err := s.APIDocumentation.CreateAPIDocument(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIDocumentRequest{
        Content: "<value>",
        Title: sdkkonnectgo.Pointer("API Document"),
        Slug: sdkkonnectgo.Pointer("api-document"),
        ParentDocumentID: nil,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiUnauthorizedExample

<!-- UsageSnippet language="go" operationID="create-api-document" method="post" path="/v3/apis/{apiId}/documents" example="ApiUnauthorizedExample" -->
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

    res, err := s.APIDocumentation.CreateAPIDocument(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIDocumentRequest{
        Content: "<value>",
        Title: sdkkonnectgo.Pointer("API Document"),
        Slug: sdkkonnectgo.Pointer("api-document"),
        ParentDocumentID: nil,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```
### Example Usage: CreateApiDocumentRequestExample

<!-- UsageSnippet language="go" operationID="create-api-document" method="post" path="/v3/apis/{apiId}/documents" example="CreateApiDocumentRequestExample" -->
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

    res, err := s.APIDocumentation.CreateAPIDocument(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIDocumentRequest{
        Content: "# API Document Header",
        Title: sdkkonnectgo.Pointer("API Document"),
        Slug: sdkkonnectgo.Pointer("api-document"),
        Status: components.APIDocumentStatusPublished.ToPointer(),
        ParentDocumentID: sdkkonnectgo.Pointer("417a2aed-9a75-4eca-a943-fa7d88ed6a73"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```
### Example Usage: CreateApiDocumentRequestMinimalExample

<!-- UsageSnippet language="go" operationID="create-api-document" method="post" path="/v3/apis/{apiId}/documents" example="CreateApiDocumentRequestMinimalExample" -->
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

    res, err := s.APIDocumentation.CreateAPIDocument(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIDocumentRequest{
        Content: "# API Document Header",
        Title: sdkkonnectgo.Pointer("API Document"),
        Slug: sdkkonnectgo.Pointer("api-document"),
        Status: components.APIDocumentStatusPublished.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```
### Example Usage: CreateApiDocumentResponse

<!-- UsageSnippet language="go" operationID="create-api-document" method="post" path="/v3/apis/{apiId}/documents" example="CreateApiDocumentResponse" -->
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

    res, err := s.APIDocumentation.CreateAPIDocument(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIDocumentRequest{
        Content: "<value>",
        Title: sdkkonnectgo.Pointer("API Document"),
        Slug: sdkkonnectgo.Pointer("api-document"),
        ParentDocumentID: nil,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```
### Example Usage: UnsupportedMediaTypeExample

<!-- UsageSnippet language="go" operationID="create-api-document" method="post" path="/v3/apis/{apiId}/documents" example="UnsupportedMediaTypeExample" -->
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

    res, err := s.APIDocumentation.CreateAPIDocument(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIDocumentRequest{
        Content: "<value>",
        Title: sdkkonnectgo.Pointer("API Document"),
        Slug: sdkkonnectgo.Pointer("api-document"),
        ParentDocumentID: nil,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `apiID`                                                                                    | *string*                                                                                   | :heavy_check_mark:                                                                         | The UUID API identifier                                                                    | 9f5061ce-78f6-4452-9108-ad7c02821fd5                                                       |
| `createAPIDocumentRequest`                                                                 | [components.CreateAPIDocumentRequest](../../models/components/createapidocumentrequest.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |

### Response

**[*operations.CreateAPIDocumentResponse](../../models/operations/createapidocumentresponse.md), error**

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

## ListAPIDocuments

Returns a collection of all documents for an API.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-api-documents" method="get" path="/v3/apis/{apiId}/documents" -->
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

    res, err := s.APIDocumentation.ListAPIDocuments(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAPIDocumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       | Example                                                                                           |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |                                                                                                   |
| `apiID`                                                                                           | *string*                                                                                          | :heavy_check_mark:                                                                                | The UUID API identifier                                                                           | 9f5061ce-78f6-4452-9108-ad7c02821fd5                                                              |
| `filter`                                                                                          | [*components.APIDocumentFilterParameters](../../models/components/apidocumentfilterparameters.md) | :heavy_minus_sign:                                                                                | Filters API Documents in the response.                                                            |                                                                                                   |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |                                                                                                   |

### Response

**[*operations.ListAPIDocumentsResponse](../../models/operations/listapidocumentsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchAPIDocument

Returns a document for the API.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-api-document" method="get" path="/v3/apis/{apiId}/documents/{documentId}" example="CreateApiDocumentResponse" -->
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

    res, err := s.APIDocumentation.FetchAPIDocument(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", "de5c9818-be5c-42e6-b514-e3d4bc30ddeb")
    if err != nil {
        log.Fatal(err)
    }
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `apiID`                                                  | *string*                                                 | :heavy_check_mark:                                       | The UUID API identifier                                  | 9f5061ce-78f6-4452-9108-ad7c02821fd5                     |
| `documentID`                                             | *string*                                                 | :heavy_check_mark:                                       | The document identifier related to the API               | de5c9818-be5c-42e6-b514-e3d4bc30ddeb                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchAPIDocumentResponse](../../models/operations/fetchapidocumentresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateAPIDocument

Updates a document for an API.

### Example Usage: ApiForbiddenExample

<!-- UsageSnippet language="go" operationID="update-api-document" method="patch" path="/v3/apis/{apiId}/documents/{documentId}" example="ApiForbiddenExample" -->
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

    res, err := s.APIDocumentation.UpdateAPIDocument(ctx, operations.UpdateAPIDocumentRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
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
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiNameBadRequestExample

<!-- UsageSnippet language="go" operationID="update-api-document" method="patch" path="/v3/apis/{apiId}/documents/{documentId}" example="ApiNameBadRequestExample" -->
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

    res, err := s.APIDocumentation.UpdateAPIDocument(ctx, operations.UpdateAPIDocumentRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
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
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiNotFoundExample

<!-- UsageSnippet language="go" operationID="update-api-document" method="patch" path="/v3/apis/{apiId}/documents/{documentId}" example="ApiNotFoundExample" -->
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

    res, err := s.APIDocumentation.UpdateAPIDocument(ctx, operations.UpdateAPIDocumentRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
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
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSlugConflictExample

<!-- UsageSnippet language="go" operationID="update-api-document" method="patch" path="/v3/apis/{apiId}/documents/{documentId}" example="ApiSlugConflictExample" -->
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

    res, err := s.APIDocumentation.UpdateAPIDocument(ctx, operations.UpdateAPIDocumentRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
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
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiUnauthorizedExample

<!-- UsageSnippet language="go" operationID="update-api-document" method="patch" path="/v3/apis/{apiId}/documents/{documentId}" example="ApiUnauthorizedExample" -->
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

    res, err := s.APIDocumentation.UpdateAPIDocument(ctx, operations.UpdateAPIDocumentRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
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
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```
### Example Usage: CreateApiDocumentResponse

<!-- UsageSnippet language="go" operationID="update-api-document" method="patch" path="/v3/apis/{apiId}/documents/{documentId}" example="CreateApiDocumentResponse" -->
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

    res, err := s.APIDocumentation.UpdateAPIDocument(ctx, operations.UpdateAPIDocumentRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
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
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```
### Example Usage: UnsupportedMediaTypeExample

<!-- UsageSnippet language="go" operationID="update-api-document" method="patch" path="/v3/apis/{apiId}/documents/{documentId}" example="UnsupportedMediaTypeExample" -->
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

    res, err := s.APIDocumentation.UpdateAPIDocument(ctx, operations.UpdateAPIDocumentRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
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
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateApiDocumentAttributes

<!-- UsageSnippet language="go" operationID="update-api-document" method="patch" path="/v3/apis/{apiId}/documents/{documentId}" example="UpdateApiDocumentAttributes" -->
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

    res, err := s.APIDocumentation.UpdateAPIDocument(ctx, operations.UpdateAPIDocumentRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        DocumentID: "de5c9818-be5c-42e6-b514-e3d4bc30ddeb",
        APIDocument: components.APIDocument{
            Title: sdkkonnectgo.Pointer("API Document"),
            Slug: sdkkonnectgo.Pointer("api-document"),
            Status: components.APIDocumentStatusPublished.ToPointer(),
            ParentDocumentID: sdkkonnectgo.Pointer("dd4e1b98-3629-4dd3-acc0-759a726ffee2"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateApiDocumentAttributesAndContent

<!-- UsageSnippet language="go" operationID="update-api-document" method="patch" path="/v3/apis/{apiId}/documents/{documentId}" example="UpdateApiDocumentAttributesAndContent" -->
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

    res, err := s.APIDocumentation.UpdateAPIDocument(ctx, operations.UpdateAPIDocumentRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        DocumentID: "de5c9818-be5c-42e6-b514-e3d4bc30ddeb",
        APIDocument: components.APIDocument{
            Content: sdkkonnectgo.Pointer("# Updated API Document Header"),
            Title: sdkkonnectgo.Pointer("API Document"),
            Slug: sdkkonnectgo.Pointer("api-document"),
            Status: components.APIDocumentStatusPublished.ToPointer(),
            ParentDocumentID: sdkkonnectgo.Pointer("dd4e1b98-3629-4dd3-acc0-759a726ffee2"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateApiDocumentContent

<!-- UsageSnippet language="go" operationID="update-api-document" method="patch" path="/v3/apis/{apiId}/documents/{documentId}" example="UpdateApiDocumentContent" -->
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

    res, err := s.APIDocumentation.UpdateAPIDocument(ctx, operations.UpdateAPIDocumentRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        DocumentID: "de5c9818-be5c-42e6-b514-e3d4bc30ddeb",
        APIDocument: components.APIDocument{
            Content: sdkkonnectgo.Pointer("# Updated API Document Header"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIDocumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateAPIDocumentRequest](../../models/operations/updateapidocumentrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateAPIDocumentResponse](../../models/operations/updateapidocumentresponse.md), error**

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

## DeleteAPIDocument

Removes a document from an API.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-api-document" method="delete" path="/v3/apis/{apiId}/documents/{documentId}" -->
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

    res, err := s.APIDocumentation.DeleteAPIDocument(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", "de5c9818-be5c-42e6-b514-e3d4bc30ddeb")
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
| `apiID`                                                  | *string*                                                 | :heavy_check_mark:                                       | The UUID API identifier                                  | 9f5061ce-78f6-4452-9108-ad7c02821fd5                     |
| `documentID`                                             | *string*                                                 | :heavy_check_mark:                                       | The document identifier related to the API               | de5c9818-be5c-42e6-b514-e3d4bc30ddeb                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAPIDocumentResponse](../../models/operations/deleteapidocumentresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## MoveAPIDocument

This api allows the user to move a document within the document tree using the parameters parent_document_id and index. If parent_document_id is not provided, the document will be placed at the top level of the document tree. index represents a zero-indexed document order relative to its siblings under the same parent. For example, if we want to put the document at top level in first position we would send parent_document_id: null and index: 0. This api also supports using a negative index to count backwards from the end of the document list, which means you can put the document in last position by using index: -1.

### Example Usage: ApiForbiddenExample

<!-- UsageSnippet language="go" operationID="move-api-document" method="post" path="/v3/apis/{apiId}/documents/{documentId}/move" example="ApiForbiddenExample" -->
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

    res, err := s.APIDocumentation.MoveAPIDocument(ctx, operations.MoveAPIDocumentRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
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
### Example Usage: ApiNotFoundExample

<!-- UsageSnippet language="go" operationID="move-api-document" method="post" path="/v3/apis/{apiId}/documents/{documentId}/move" example="ApiNotFoundExample" -->
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

    res, err := s.APIDocumentation.MoveAPIDocument(ctx, operations.MoveAPIDocumentRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
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
### Example Usage: ApiUnauthorizedExample

<!-- UsageSnippet language="go" operationID="move-api-document" method="post" path="/v3/apis/{apiId}/documents/{documentId}/move" example="ApiUnauthorizedExample" -->
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

    res, err := s.APIDocumentation.MoveAPIDocument(ctx, operations.MoveAPIDocumentRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.MoveAPIDocumentRequest](../../models/operations/moveapidocumentrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.MoveAPIDocumentResponse](../../models/operations/moveapidocumentresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |