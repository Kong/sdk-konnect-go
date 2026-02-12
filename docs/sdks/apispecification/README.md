# APISpecification

## Overview

### Available Operations

* [~~CreateAPISpec~~](#createapispec) - Create API Specification :warning: **Deprecated**
* [~~ListAPISpecs~~](#listapispecs) - List API Specifications :warning: **Deprecated**
* [~~FetchAPISpec~~](#fetchapispec) - Get API Specification :warning: **Deprecated**
* [~~UpdateAPISpec~~](#updateapispec) - Update API Specification :warning: **Deprecated**
* [~~DeleteAPISpec~~](#deleteapispec) - Delete API Specification :warning: **Deprecated**
* [ValidateSpecification](#validatespecification) - Validate API Specification

## ~~CreateAPISpec~~

Creates a specification (OpenAPI or AsyncAPI) for an API.
**Note:** You can only have one specification for an API. This endpoint is deprecated and will be removed: use /versions instead.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage: ApiForbiddenExample

<!-- UsageSnippet language="go" operationID="create-api-spec" method="post" path="/v3/apis/{apiId}/specifications" example="ApiForbiddenExample" -->
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

    res, err := s.APISpecification.CreateAPISpec(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPISpecRequest{
        Content: "{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}",
        Type: components.CreateAPISpecRequestAPISpecTypeOas3.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiNotFoundExample

<!-- UsageSnippet language="go" operationID="create-api-spec" method="post" path="/v3/apis/{apiId}/specifications" example="ApiNotFoundExample" -->
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

    res, err := s.APISpecification.CreateAPISpec(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPISpecRequest{
        Content: "{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}",
        Type: components.CreateAPISpecRequestAPISpecTypeOas3.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSpecConflictExample

<!-- UsageSnippet language="go" operationID="create-api-spec" method="post" path="/v3/apis/{apiId}/specifications" example="ApiSpecConflictExample" -->
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

    res, err := s.APISpecification.CreateAPISpec(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPISpecRequest{
        Content: "{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}",
        Type: components.CreateAPISpecRequestAPISpecTypeOas3.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSpecContentBadRequestExample

<!-- UsageSnippet language="go" operationID="create-api-spec" method="post" path="/v3/apis/{apiId}/specifications" example="ApiSpecContentBadRequestExample" -->
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

    res, err := s.APISpecification.CreateAPISpec(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPISpecRequest{
        Content: "{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}",
        Type: components.CreateAPISpecRequestAPISpecTypeOas3.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSpecResponse

<!-- UsageSnippet language="go" operationID="create-api-spec" method="post" path="/v3/apis/{apiId}/specifications" example="ApiSpecResponse" -->
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

    res, err := s.APISpecification.CreateAPISpec(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPISpecRequest{
        Content: "{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}",
        Type: components.CreateAPISpecRequestAPISpecTypeOas3.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSpecResponseAsyncApi

<!-- UsageSnippet language="go" operationID="create-api-spec" method="post" path="/v3/apis/{apiId}/specifications" example="ApiSpecResponseAsyncApi" -->
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

    res, err := s.APISpecification.CreateAPISpec(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPISpecRequest{
        Content: "{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}",
        Type: components.CreateAPISpecRequestAPISpecTypeOas3.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiUnauthorizedExample

<!-- UsageSnippet language="go" operationID="create-api-spec" method="post" path="/v3/apis/{apiId}/specifications" example="ApiUnauthorizedExample" -->
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

    res, err := s.APISpecification.CreateAPISpec(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPISpecRequest{
        Content: "{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}",
        Type: components.CreateAPISpecRequestAPISpecTypeOas3.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: CreateApiSpecRequestExampleAsyncApi

<!-- UsageSnippet language="go" operationID="create-api-spec" method="post" path="/v3/apis/{apiId}/specifications" example="CreateApiSpecRequestExampleAsyncApi" -->
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

    res, err := s.APISpecification.CreateAPISpec(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPISpecRequest{
        Content: "{\"asyncapi\":\"2.6.0\",\"info\":{\"title\":\"Account Service\",\"version\":\"1.0.0\",\"description\":\"This service is in charge of processing user signups\"},\"channels\":{\"user/signedup\":{\"publish\":{\"summary\":\"Inform about signed up users\",\"operationId\":\"sendUserSignedUp\",\"message\":{\"name\":\"UserSignedUp\",\"payload\":{\"type\":\"object\",\"properties\":{\"displayName\":{\"type\":\"string\",\"description\":\"Name of the user\"},\"email\":{\"type\":\"string\",\"format\":\"email\",\"description\":\"Email of the user\"}}}}}}}}",
        Type: components.CreateAPISpecRequestAPISpecTypeAsyncapi.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: CreateApiSpecRequestExampleJson

<!-- UsageSnippet language="go" operationID="create-api-spec" method="post" path="/v3/apis/{apiId}/specifications" example="CreateApiSpecRequestExampleJson" -->
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

    res, err := s.APISpecification.CreateAPISpec(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPISpecRequest{
        Content: "{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}",
        Type: components.CreateAPISpecRequestAPISpecTypeOas3.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: CreateApiSpecRequestExampleYaml

<!-- UsageSnippet language="go" operationID="create-api-spec" method="post" path="/v3/apis/{apiId}/specifications" example="CreateApiSpecRequestExampleYaml" -->
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

    res, err := s.APISpecification.CreateAPISpec(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPISpecRequest{
        Content: "openapi: 3.0.3\\ninfo:\\n  title: Example API\\n  version: 1.0.0\\npaths:\\n  /example:\\n    get:\\n      summary: Example endpoint\\n      responses:\\n        \\\"200\\\":\\n          description: Successful response",
        Type: components.CreateAPISpecRequestAPISpecTypeOas3.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: UnsupportedMediaTypeExample

<!-- UsageSnippet language="go" operationID="create-api-spec" method="post" path="/v3/apis/{apiId}/specifications" example="UnsupportedMediaTypeExample" -->
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

    res, err := s.APISpecification.CreateAPISpec(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPISpecRequest{
        Content: "{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}",
        Type: components.CreateAPISpecRequestAPISpecTypeOas3.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `apiID`                                                                            | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID API identifier                                                            | 9f5061ce-78f6-4452-9108-ad7c02821fd5                                               |
| `createAPISpecRequest`                                                             | [components.CreateAPISpecRequest](../../models/components/createapispecrequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.CreateAPISpecResponse](../../models/operations/createapispecresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.BadRequestError           | 400                                 | application/problem+json            |
| sdkerrors.UnauthorizedError         | 401                                 | application/problem+json            |
| sdkerrors.ForbiddenError            | 403                                 | application/problem+json            |
| sdkerrors.NotFoundError             | 404                                 | application/problem+json            |
| sdkerrors.ConflictError             | 409                                 | application/problem+json            |
| sdkerrors.UnsupportedMediaTypeError | 415                                 | application/problem+json            |
| sdkerrors.SDKError                  | 4XX, 5XX                            | \*/\*                               |

## ~~ListAPISpecs~~

Returns a list of specifications for an API. The specification can be of type OpenAPI or AsyncAPI.
**Note:** You can only have one specification for an API. This endpoint is deprecated and will be removed: use /versions instead.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage: ListApiSpecResponse

<!-- UsageSnippet language="go" operationID="list-api-specs" method="get" path="/v3/apis/{apiId}/specifications" example="ListApiSpecResponse" -->
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

    res, err := s.APISpecification.ListAPISpecs(ctx, operations.ListAPISpecsRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAPISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: ListApiSpecResponseAsyncApi

<!-- UsageSnippet language="go" operationID="list-api-specs" method="get" path="/v3/apis/{apiId}/specifications" example="ListApiSpecResponseAsyncApi" -->
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

    res, err := s.APISpecification.ListAPISpecs(ctx, operations.ListAPISpecsRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAPISpecResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListAPISpecsRequest](../../models/operations/listapispecsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListAPISpecsResponse](../../models/operations/listapispecsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ~~FetchAPISpec~~

Fetches the specification (OpenAPI or AsyncAPI) of an API.
**Note:** This endpoint is deprecated and will be removed: use /versions instead.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage: ApiSpecResponse

<!-- UsageSnippet language="go" operationID="fetch-api-spec" method="get" path="/v3/apis/{apiId}/specifications/{specId}" example="ApiSpecResponse" -->
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

    res, err := s.APISpecification.FetchAPISpec(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", "d32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSpecResponseAsyncApi

<!-- UsageSnippet language="go" operationID="fetch-api-spec" method="get" path="/v3/apis/{apiId}/specifications/{specId}" example="ApiSpecResponseAsyncApi" -->
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

    res, err := s.APISpecification.FetchAPISpec(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", "d32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `apiID`                                                  | *string*                                                 | :heavy_check_mark:                                       | The UUID API identifier                                  | 9f5061ce-78f6-4452-9108-ad7c02821fd5                     |
| `specID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The API specification identifier                         | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchAPISpecResponse](../../models/operations/fetchapispecresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ~~UpdateAPISpec~~

Updates the specification (OpenAPI or AsyncAPI) of an API.
**Note:** This endpoint is deprecated and will be removed: use /versions instead.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage: ApiForbiddenExample

<!-- UsageSnippet language="go" operationID="update-api-spec" method="patch" path="/v3/apis/{apiId}/specifications/{specId}" example="ApiForbiddenExample" -->
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

    res, err := s.APISpecification.UpdateAPISpec(ctx, operations.UpdateAPISpecRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        SpecID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APISpec: components.APISpec{
            Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
            Type: components.APISpecAPISpecTypeOas3.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiNotFoundExample

<!-- UsageSnippet language="go" operationID="update-api-spec" method="patch" path="/v3/apis/{apiId}/specifications/{specId}" example="ApiNotFoundExample" -->
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

    res, err := s.APISpecification.UpdateAPISpec(ctx, operations.UpdateAPISpecRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        SpecID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APISpec: components.APISpec{
            Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
            Type: components.APISpecAPISpecTypeOas3.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSpecConflictExample

<!-- UsageSnippet language="go" operationID="update-api-spec" method="patch" path="/v3/apis/{apiId}/specifications/{specId}" example="ApiSpecConflictExample" -->
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

    res, err := s.APISpecification.UpdateAPISpec(ctx, operations.UpdateAPISpecRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        SpecID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APISpec: components.APISpec{
            Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
            Type: components.APISpecAPISpecTypeOas3.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSpecContentBadRequestExample

<!-- UsageSnippet language="go" operationID="update-api-spec" method="patch" path="/v3/apis/{apiId}/specifications/{specId}" example="ApiSpecContentBadRequestExample" -->
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

    res, err := s.APISpecification.UpdateAPISpec(ctx, operations.UpdateAPISpecRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        SpecID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APISpec: components.APISpec{
            Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
            Type: components.APISpecAPISpecTypeOas3.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSpecResponse

<!-- UsageSnippet language="go" operationID="update-api-spec" method="patch" path="/v3/apis/{apiId}/specifications/{specId}" example="ApiSpecResponse" -->
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

    res, err := s.APISpecification.UpdateAPISpec(ctx, operations.UpdateAPISpecRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        SpecID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APISpec: components.APISpec{
            Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
            Type: components.APISpecAPISpecTypeOas3.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSpecResponseAsyncApi

<!-- UsageSnippet language="go" operationID="update-api-spec" method="patch" path="/v3/apis/{apiId}/specifications/{specId}" example="ApiSpecResponseAsyncApi" -->
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

    res, err := s.APISpecification.UpdateAPISpec(ctx, operations.UpdateAPISpecRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        SpecID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APISpec: components.APISpec{
            Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
            Type: components.APISpecAPISpecTypeOas3.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiUnauthorizedExample

<!-- UsageSnippet language="go" operationID="update-api-spec" method="patch" path="/v3/apis/{apiId}/specifications/{specId}" example="ApiUnauthorizedExample" -->
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

    res, err := s.APISpecification.UpdateAPISpec(ctx, operations.UpdateAPISpecRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        SpecID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APISpec: components.APISpec{
            Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
            Type: components.APISpecAPISpecTypeOas3.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: UnsupportedMediaTypeExample

<!-- UsageSnippet language="go" operationID="update-api-spec" method="patch" path="/v3/apis/{apiId}/specifications/{specId}" example="UnsupportedMediaTypeExample" -->
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

    res, err := s.APISpecification.UpdateAPISpec(ctx, operations.UpdateAPISpecRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        SpecID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APISpec: components.APISpec{
            Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
            Type: components.APISpecAPISpecTypeOas3.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateApiSpecContentRequest

<!-- UsageSnippet language="go" operationID="update-api-spec" method="patch" path="/v3/apis/{apiId}/specifications/{specId}" example="UpdateApiSpecContentRequest" -->
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

    res, err := s.APISpecification.UpdateAPISpec(ctx, operations.UpdateAPISpecRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        SpecID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APISpec: components.APISpec{
            Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.1\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateApiSpecContentRequestAsyncApi

<!-- UsageSnippet language="go" operationID="update-api-spec" method="patch" path="/v3/apis/{apiId}/specifications/{specId}" example="UpdateApiSpecContentRequestAsyncApi" -->
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

    res, err := s.APISpecification.UpdateAPISpec(ctx, operations.UpdateAPISpecRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        SpecID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APISpec: components.APISpec{
            Content: sdkkonnectgo.Pointer("{\"asyncapi\":\"2.6.0\",\"info\":{\"title\":\"Account Service Updated\",\"version\":\"1.0.1\",\"description\":\"This service is in charge of processing user signups (updated)\"},\"channels\":{\"user/signedup\":{\"publish\":{\"summary\":\"Inform about signed up users\",\"operationId\":\"sendUserSignedUp\",\"message\":{\"name\":\"UserSignedUp\",\"payload\":{\"type\":\"object\",\"properties\":{\"displayName\":{\"type\":\"string\",\"description\":\"Name of the user\"},\"email\":{\"type\":\"string\",\"format\":\"email\",\"description\":\"Email of the user\"}}}}}}}}"),
            Type: components.APISpecAPISpecTypeAsyncapi.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateAPISpecRequest](../../models/operations/updateapispecrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.UpdateAPISpecResponse](../../models/operations/updateapispecresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.BadRequestError           | 400                                 | application/problem+json            |
| sdkerrors.UnauthorizedError         | 401                                 | application/problem+json            |
| sdkerrors.ForbiddenError            | 403                                 | application/problem+json            |
| sdkerrors.NotFoundError             | 404                                 | application/problem+json            |
| sdkerrors.ConflictError             | 409                                 | application/problem+json            |
| sdkerrors.UnsupportedMediaTypeError | 415                                 | application/problem+json            |
| sdkerrors.SDKError                  | 4XX, 5XX                            | \*/\*                               |

## ~~DeleteAPISpec~~

Deletes the specification (OpenAPI or AsyncAPI) of an API.
**Note:** This endpoint is deprecated and will be removed: use /versions instead.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-api-spec" method="delete" path="/v3/apis/{apiId}/specifications/{specId}" -->
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

    res, err := s.APISpecification.DeleteAPISpec(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", "d32d905a-ed33-46a3-a093-d8f536af9a8a")
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
| `specID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The API specification identifier                         | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAPISpecResponse](../../models/operations/deleteapispecresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ValidateSpecification

Validates the content and type (OpenAPI or AsyncAPI) of a potential API specification without associating it with a specific API.

### Example Usage: ApiForbiddenExample

<!-- UsageSnippet language="go" operationID="validate-specification" method="post" path="/v3/apis/validate-specification" example="ApiForbiddenExample" -->
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

    res, err := s.APISpecification.ValidateSpecification(ctx, components.ValidateAPISpecRequestPayload{
        Content: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ValidateAPISpecSuccessResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSpecContentBadRequestExample

<!-- UsageSnippet language="go" operationID="validate-specification" method="post" path="/v3/apis/validate-specification" example="ApiSpecContentBadRequestExample" -->
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

    res, err := s.APISpecification.ValidateSpecification(ctx, components.ValidateAPISpecRequestPayload{
        Content: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ValidateAPISpecSuccessResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSpecValidationFailedBadRequestExample

<!-- UsageSnippet language="go" operationID="validate-specification" method="post" path="/v3/apis/validate-specification" example="ApiSpecValidationFailedBadRequestExample" -->
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

    res, err := s.APISpecification.ValidateSpecification(ctx, components.ValidateAPISpecRequestPayload{
        Content: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ValidateAPISpecSuccessResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiUnauthorizedExample

<!-- UsageSnippet language="go" operationID="validate-specification" method="post" path="/v3/apis/validate-specification" example="ApiUnauthorizedExample" -->
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

    res, err := s.APISpecification.ValidateSpecification(ctx, components.ValidateAPISpecRequestPayload{
        Content: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ValidateAPISpecSuccessResponse != nil {
        // handle response
    }
}
```
### Example Usage: UnsupportedMediaTypeExample

<!-- UsageSnippet language="go" operationID="validate-specification" method="post" path="/v3/apis/validate-specification" example="UnsupportedMediaTypeExample" -->
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

    res, err := s.APISpecification.ValidateSpecification(ctx, components.ValidateAPISpecRequestPayload{
        Content: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ValidateAPISpecSuccessResponse != nil {
        // handle response
    }
}
```
### Example Usage: ValidateApiSpecRequestExampleAsyncApi

<!-- UsageSnippet language="go" operationID="validate-specification" method="post" path="/v3/apis/validate-specification" example="ValidateApiSpecRequestExampleAsyncApi" -->
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

    res, err := s.APISpecification.ValidateSpecification(ctx, components.ValidateAPISpecRequestPayload{
        Content: "{\"asyncapi\":\"2.6.0\",\"info\":{\"title\":\"Account Service\",\"version\":\"1.0.0\",\"description\":\"This service is in charge of processing user signups\"},\"channels\":{\"user/signedup\":{\"publish\":{\"summary\":\"Inform about signed up users\",\"operationId\":\"sendUserSignedUp\",\"message\":{\"name\":\"UserSignedUp\",\"payload\":{\"type\":\"object\",\"properties\":{\"displayName\":{\"type\":\"string\",\"description\":\"Name of the user\"},\"email\":{\"type\":\"string\",\"format\":\"email\",\"description\":\"Email of the user\"}}}}}}}}",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ValidateAPISpecSuccessResponse != nil {
        // handle response
    }
}
```
### Example Usage: ValidateApiSpecRequestExampleJson

<!-- UsageSnippet language="go" operationID="validate-specification" method="post" path="/v3/apis/validate-specification" example="ValidateApiSpecRequestExampleJson" -->
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

    res, err := s.APISpecification.ValidateSpecification(ctx, components.ValidateAPISpecRequestPayload{
        Content: "{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Valid API\",\"version\":\"1.0.0\"},\"paths\":{\"/health\":{\"get\":{\"summary\":\"Health check\",\"responses\":{\"200\":{\"description\":\"OK\"}}}}}}",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ValidateAPISpecSuccessResponse != nil {
        // handle response
    }
}
```
### Example Usage: ValidateApiSpecRequestExampleYaml

<!-- UsageSnippet language="go" operationID="validate-specification" method="post" path="/v3/apis/validate-specification" example="ValidateApiSpecRequestExampleYaml" -->
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

    res, err := s.APISpecification.ValidateSpecification(ctx, components.ValidateAPISpecRequestPayload{
        Content: "openapi: 3.0.3\\ninfo:\\n  title: Valid API\\n  version: 1.0.0\\npaths:\\n  /health:\\n    get:\\n      summary: Health check\\n      responses:\\n        '200':\\n          description: OK",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ValidateAPISpecSuccessResponse != nil {
        // handle response
    }
}
```
### Example Usage: ValidateApiSpecSuccessResponseExample

<!-- UsageSnippet language="go" operationID="validate-specification" method="post" path="/v3/apis/validate-specification" example="ValidateApiSpecSuccessResponseExample" -->
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

    res, err := s.APISpecification.ValidateSpecification(ctx, components.ValidateAPISpecRequestPayload{
        Content: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ValidateAPISpecSuccessResponse != nil {
        // handle response
    }
}
```
### Example Usage: ValidateApiSpecSuccessResponseWithWarningExample

<!-- UsageSnippet language="go" operationID="validate-specification" method="post" path="/v3/apis/validate-specification" example="ValidateApiSpecSuccessResponseWithWarningExample" -->
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

    res, err := s.APISpecification.ValidateSpecification(ctx, components.ValidateAPISpecRequestPayload{
        Content: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ValidateAPISpecSuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [components.ValidateAPISpecRequestPayload](../../models/components/validateapispecrequestpayload.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ValidateSpecificationResponse](../../models/operations/validatespecificationresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.BadRequestError           | 400                                 | application/problem+json            |
| sdkerrors.UnauthorizedError         | 401                                 | application/problem+json            |
| sdkerrors.ForbiddenError            | 403                                 | application/problem+json            |
| sdkerrors.UnsupportedMediaTypeError | 415                                 | application/problem+json            |
| sdkerrors.SDKError                  | 4XX, 5XX                            | \*/\*                               |