# APIVersion

## Overview

### Available Operations

* [CreateAPIVersion](#createapiversion) - Create API Version
* [ListAPIVersions](#listapiversions) - List API Versions
* [FetchAPIVersion](#fetchapiversion) - Get an API Version
* [UpdateAPIVersion](#updateapiversion) - Update API Version
* [DeleteAPIVersion](#deleteapiversion) - Delete API Version

## CreateAPIVersion

Creates a version (OpenAPI or AsyncAPI) for an API.


### Example Usage: ApiForbiddenExample

<!-- UsageSnippet language="go" operationID="create-api-version" method="post" path="/v3/apis/{apiId}/versions" example="ApiForbiddenExample" -->
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

    res, err := s.APIVersion.CreateAPIVersion(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIVersionRequest{
        Version: sdkkonnectgo.Pointer("1.0.0"),
        Spec: components.CreateAPIVersionRequestSpec{
            Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiNotFoundExample

<!-- UsageSnippet language="go" operationID="create-api-version" method="post" path="/v3/apis/{apiId}/versions" example="ApiNotFoundExample" -->
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

    res, err := s.APIVersion.CreateAPIVersion(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIVersionRequest{
        Version: sdkkonnectgo.Pointer("1.0.0"),
        Spec: components.CreateAPIVersionRequestSpec{
            Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSpecConflictExample

<!-- UsageSnippet language="go" operationID="create-api-version" method="post" path="/v3/apis/{apiId}/versions" example="ApiSpecConflictExample" -->
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

    res, err := s.APIVersion.CreateAPIVersion(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIVersionRequest{
        Version: sdkkonnectgo.Pointer("1.0.0"),
        Spec: components.CreateAPIVersionRequestSpec{
            Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSpecResponse

<!-- UsageSnippet language="go" operationID="create-api-version" method="post" path="/v3/apis/{apiId}/versions" example="ApiSpecResponse" -->
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

    res, err := s.APIVersion.CreateAPIVersion(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIVersionRequest{
        Version: sdkkonnectgo.Pointer("1.0.0"),
        Spec: components.CreateAPIVersionRequestSpec{
            Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSpecResponseAsyncApi

<!-- UsageSnippet language="go" operationID="create-api-version" method="post" path="/v3/apis/{apiId}/versions" example="ApiSpecResponseAsyncApi" -->
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

    res, err := s.APIVersion.CreateAPIVersion(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIVersionRequest{
        Version: sdkkonnectgo.Pointer("1.0.0"),
        Spec: components.CreateAPIVersionRequestSpec{
            Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiUnauthorizedExample

<!-- UsageSnippet language="go" operationID="create-api-version" method="post" path="/v3/apis/{apiId}/versions" example="ApiUnauthorizedExample" -->
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

    res, err := s.APIVersion.CreateAPIVersion(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIVersionRequest{
        Version: sdkkonnectgo.Pointer("1.0.0"),
        Spec: components.CreateAPIVersionRequestSpec{
            Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiVersionBadRequestExample

<!-- UsageSnippet language="go" operationID="create-api-version" method="post" path="/v3/apis/{apiId}/versions" example="ApiVersionBadRequestExample" -->
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

    res, err := s.APIVersion.CreateAPIVersion(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIVersionRequest{
        Version: sdkkonnectgo.Pointer("1.0.0"),
        Spec: components.CreateAPIVersionRequestSpec{
            Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: CreateApiVersionRequestExampleAsyncApi

<!-- UsageSnippet language="go" operationID="create-api-version" method="post" path="/v3/apis/{apiId}/versions" example="CreateApiVersionRequestExampleAsyncApi" -->
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

    res, err := s.APIVersion.CreateAPIVersion(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIVersionRequest{
        Version: sdkkonnectgo.Pointer("1.0.0"),
        Spec: components.CreateAPIVersionRequestSpec{
            Content: sdkkonnectgo.Pointer("{\"asyncapi\":\"2.6.0\",\"info\":{\"title\":\"Account Service\",\"version\":\"1.0.0\",\"description\":\"This service is in charge of processing user signups\"},\"channels\":{\"user/signedup\":{\"publish\":{\"summary\":\"Inform about signed up users\",\"operationId\":\"sendUserSignedUp\",\"message\":{\"name\":\"UserSignedUp\",\"payload\":{\"type\":\"object\",\"properties\":{\"displayName\":{\"type\":\"string\",\"description\":\"Name of the user\"},\"email\":{\"type\":\"string\",\"format\":\"email\",\"description\":\"Email of the user\"}}}}}}}}"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: CreateApiVersionRequestExampleJson

<!-- UsageSnippet language="go" operationID="create-api-version" method="post" path="/v3/apis/{apiId}/versions" example="CreateApiVersionRequestExampleJson" -->
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

    res, err := s.APIVersion.CreateAPIVersion(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIVersionRequest{
        Version: sdkkonnectgo.Pointer("1.0.0"),
        Spec: components.CreateAPIVersionRequestSpec{
            Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: CreateApiVersionRequestExampleYaml

<!-- UsageSnippet language="go" operationID="create-api-version" method="post" path="/v3/apis/{apiId}/versions" example="CreateApiVersionRequestExampleYaml" -->
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

    res, err := s.APIVersion.CreateAPIVersion(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIVersionRequest{
        Version: sdkkonnectgo.Pointer("1.0.0"),
        Spec: components.CreateAPIVersionRequestSpec{
            Content: sdkkonnectgo.Pointer("openapi: 3.0.3\\ninfo:\\n  title: Example API\\n  version: 1.0.0\\npaths:\\n  /example:\\n    get:\\n      summary: Example endpoint\\n      responses:\\n        \\\"200\\\":\\n          description: Successful response"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: UnsupportedMediaTypeExample

<!-- UsageSnippet language="go" operationID="create-api-version" method="post" path="/v3/apis/{apiId}/versions" example="UnsupportedMediaTypeExample" -->
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

    res, err := s.APIVersion.CreateAPIVersion(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIVersionRequest{
        Version: sdkkonnectgo.Pointer("1.0.0"),
        Spec: components.CreateAPIVersionRequestSpec{
            Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `apiID`                                                                                  | *string*                                                                                 | :heavy_check_mark:                                                                       | The UUID API identifier                                                                  | 9f5061ce-78f6-4452-9108-ad7c02821fd5                                                     |
| `createAPIVersionRequest`                                                                | [components.CreateAPIVersionRequest](../../models/components/createapiversionrequest.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |                                                                                          |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.CreateAPIVersionResponse](../../models/operations/createapiversionresponse.md), error**

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

## ListAPIVersions

Returns a list of versions for an API. The version can be of type OpenAPI or AsyncAPI.
**Note:** You can only have one version for an API.


### Example Usage: ListApiSpecResponse

<!-- UsageSnippet language="go" operationID="list-api-versions" method="get" path="/v3/apis/{apiId}/versions" example="ListApiSpecResponse" -->
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

    res, err := s.APIVersion.ListAPIVersions(ctx, operations.ListAPIVersionsRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAPIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: ListApiSpecResponseAsyncApi

<!-- UsageSnippet language="go" operationID="list-api-versions" method="get" path="/v3/apis/{apiId}/versions" example="ListApiSpecResponseAsyncApi" -->
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

    res, err := s.APIVersion.ListAPIVersions(ctx, operations.ListAPIVersionsRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAPIVersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListAPIVersionsRequest](../../models/operations/listapiversionsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListAPIVersionsResponse](../../models/operations/listapiversionsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchAPIVersion

Fetches the version (OpenAPI or AsyncAPI) of an API.

### Example Usage: ApiSpecResponse

<!-- UsageSnippet language="go" operationID="fetch-api-version" method="get" path="/v3/apis/{apiId}/versions/{versionId}" example="ApiSpecResponse" -->
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

    res, err := s.APIVersion.FetchAPIVersion(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", "d32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSpecResponseAsyncApi

<!-- UsageSnippet language="go" operationID="fetch-api-version" method="get" path="/v3/apis/{apiId}/versions/{versionId}" example="ApiSpecResponseAsyncApi" -->
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

    res, err := s.APIVersion.FetchAPIVersion(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", "d32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `apiID`                                                  | *string*                                                 | :heavy_check_mark:                                       | The UUID API identifier                                  | 9f5061ce-78f6-4452-9108-ad7c02821fd5                     |
| `versionID`                                              | *string*                                                 | :heavy_check_mark:                                       | The API version identifier                               | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchAPIVersionResponse](../../models/operations/fetchapiversionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateAPIVersion

Updates the version (OpenAPI or AsyncAPI) of an API.

### Example Usage: ApiForbiddenExample

<!-- UsageSnippet language="go" operationID="update-api-version" method="patch" path="/v3/apis/{apiId}/versions/{versionId}" example="ApiForbiddenExample" -->
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

    res, err := s.APIVersion.UpdateAPIVersion(ctx, operations.UpdateAPIVersionRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        VersionID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APIVersion: components.APIVersion{
            Version: sdkkonnectgo.Pointer("1.0.0"),
            Spec: &components.APIVersionSpec{
                Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiNotFoundExample

<!-- UsageSnippet language="go" operationID="update-api-version" method="patch" path="/v3/apis/{apiId}/versions/{versionId}" example="ApiNotFoundExample" -->
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

    res, err := s.APIVersion.UpdateAPIVersion(ctx, operations.UpdateAPIVersionRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        VersionID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APIVersion: components.APIVersion{
            Version: sdkkonnectgo.Pointer("1.0.0"),
            Spec: &components.APIVersionSpec{
                Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSpecConflictExample

<!-- UsageSnippet language="go" operationID="update-api-version" method="patch" path="/v3/apis/{apiId}/versions/{versionId}" example="ApiSpecConflictExample" -->
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

    res, err := s.APIVersion.UpdateAPIVersion(ctx, operations.UpdateAPIVersionRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        VersionID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APIVersion: components.APIVersion{
            Version: sdkkonnectgo.Pointer("1.0.0"),
            Spec: &components.APIVersionSpec{
                Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSpecResponse

<!-- UsageSnippet language="go" operationID="update-api-version" method="patch" path="/v3/apis/{apiId}/versions/{versionId}" example="ApiSpecResponse" -->
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

    res, err := s.APIVersion.UpdateAPIVersion(ctx, operations.UpdateAPIVersionRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        VersionID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APIVersion: components.APIVersion{
            Version: sdkkonnectgo.Pointer("1.0.0"),
            Spec: &components.APIVersionSpec{
                Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiSpecResponseAsyncApi

<!-- UsageSnippet language="go" operationID="update-api-version" method="patch" path="/v3/apis/{apiId}/versions/{versionId}" example="ApiSpecResponseAsyncApi" -->
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

    res, err := s.APIVersion.UpdateAPIVersion(ctx, operations.UpdateAPIVersionRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        VersionID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APIVersion: components.APIVersion{
            Version: sdkkonnectgo.Pointer("1.0.0"),
            Spec: &components.APIVersionSpec{
                Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiUnauthorizedExample

<!-- UsageSnippet language="go" operationID="update-api-version" method="patch" path="/v3/apis/{apiId}/versions/{versionId}" example="ApiUnauthorizedExample" -->
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

    res, err := s.APIVersion.UpdateAPIVersion(ctx, operations.UpdateAPIVersionRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        VersionID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APIVersion: components.APIVersion{
            Version: sdkkonnectgo.Pointer("1.0.0"),
            Spec: &components.APIVersionSpec{
                Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: ApiVersionBadRequestExample

<!-- UsageSnippet language="go" operationID="update-api-version" method="patch" path="/v3/apis/{apiId}/versions/{versionId}" example="ApiVersionBadRequestExample" -->
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

    res, err := s.APIVersion.UpdateAPIVersion(ctx, operations.UpdateAPIVersionRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        VersionID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APIVersion: components.APIVersion{
            Version: sdkkonnectgo.Pointer("1.0.0"),
            Spec: &components.APIVersionSpec{
                Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: UnsupportedMediaTypeExample

<!-- UsageSnippet language="go" operationID="update-api-version" method="patch" path="/v3/apis/{apiId}/versions/{versionId}" example="UnsupportedMediaTypeExample" -->
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

    res, err := s.APIVersion.UpdateAPIVersion(ctx, operations.UpdateAPIVersionRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        VersionID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APIVersion: components.APIVersion{
            Version: sdkkonnectgo.Pointer("1.0.0"),
            Spec: &components.APIVersionSpec{
                Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateApiVersionContentRequest

<!-- UsageSnippet language="go" operationID="update-api-version" method="patch" path="/v3/apis/{apiId}/versions/{versionId}" example="UpdateApiVersionContentRequest" -->
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

    res, err := s.APIVersion.UpdateAPIVersion(ctx, operations.UpdateAPIVersionRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        VersionID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APIVersion: components.APIVersion{
            Version: sdkkonnectgo.Pointer("1.0.1"),
            Spec: &components.APIVersionSpec{
                Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.1\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateApiVersionContentRequestAsyncApi

<!-- UsageSnippet language="go" operationID="update-api-version" method="patch" path="/v3/apis/{apiId}/versions/{versionId}" example="UpdateApiVersionContentRequestAsyncApi" -->
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

    res, err := s.APIVersion.UpdateAPIVersion(ctx, operations.UpdateAPIVersionRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        VersionID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APIVersion: components.APIVersion{
            Version: sdkkonnectgo.Pointer("1.0.2"),
            Spec: &components.APIVersionSpec{
                Content: sdkkonnectgo.Pointer("{\"asyncapi\":\"2.6.0\",\"info\":{\"title\":\"Account Service Updated\",\"version\":\"1.0.1\",\"description\":\"This service is in charge of processing user signups (updated)\"},\"channels\":{\"user/signedup\":{\"publish\":{\"summary\":\"Inform about signed up users\",\"operationId\":\"sendUserSignedUp\",\"message\":{\"name\":\"UserSignedUp\",\"payload\":{\"type\":\"object\",\"properties\":{\"displayName\":{\"type\":\"string\",\"description\":\"Name of the user\"},\"email\":{\"type\":\"string\",\"format\":\"email\",\"description\":\"Email of the user\"}}}}}}}}"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIVersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateAPIVersionRequest](../../models/operations/updateapiversionrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.UpdateAPIVersionResponse](../../models/operations/updateapiversionresponse.md), error**

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

## DeleteAPIVersion

Deletes the version (OpenAPI or AsyncAPI) of an API.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-api-version" method="delete" path="/v3/apis/{apiId}/versions/{versionId}" -->
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

    res, err := s.APIVersion.DeleteAPIVersion(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", "d32d905a-ed33-46a3-a093-d8f536af9a8a")
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
| `versionID`                                              | *string*                                                 | :heavy_check_mark:                                       | The API version identifier                               | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAPIVersionResponse](../../models/operations/deleteapiversionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |