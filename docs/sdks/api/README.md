# Api

## Overview

### Available Operations

* [CreateAPI](#createapi) - Create API
* [ListApis](#listapis) - List APIs
* [FetchAPI](#fetchapi) - Get an API
* [UpdateAPI](#updateapi) - Update API
* [DeleteAPI](#deleteapi) - Delete API

## CreateAPI

Creates an API.

### Example Usage: ApiForbiddenExample

<!-- UsageSnippet language="go" operationID="create-api" method="post" path="/v3/apis" example="ApiForbiddenExample" -->
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

    res, err := s.API.CreateAPI(ctx, components.CreateAPIRequest{
        Name: "MyAPI",
        Slug: sdkkonnectgo.Pointer("my-api-v1"),
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponseSchema != nil {
        // handle response
    }
}
```
### Example Usage: ApiNameBadRequestExample

<!-- UsageSnippet language="go" operationID="create-api" method="post" path="/v3/apis" example="ApiNameBadRequestExample" -->
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

    res, err := s.API.CreateAPI(ctx, components.CreateAPIRequest{
        Name: "MyAPI",
        Slug: sdkkonnectgo.Pointer("my-api-v1"),
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponseSchema != nil {
        // handle response
    }
}
```
### Example Usage: ApiNotFoundExample

<!-- UsageSnippet language="go" operationID="create-api" method="post" path="/v3/apis" example="ApiNotFoundExample" -->
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

    res, err := s.API.CreateAPI(ctx, components.CreateAPIRequest{
        Name: "MyAPI",
        Slug: sdkkonnectgo.Pointer("my-api-v1"),
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponseSchema != nil {
        // handle response
    }
}
```
### Example Usage: ApiUnauthorizedExample

<!-- UsageSnippet language="go" operationID="create-api" method="post" path="/v3/apis" example="ApiUnauthorizedExample" -->
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

    res, err := s.API.CreateAPI(ctx, components.CreateAPIRequest{
        Name: "MyAPI",
        Slug: sdkkonnectgo.Pointer("my-api-v1"),
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponseSchema != nil {
        // handle response
    }
}
```
### Example Usage: CreateApiFull

<!-- UsageSnippet language="go" operationID="create-api" method="post" path="/v3/apis" example="CreateApiFull" -->
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

    res, err := s.API.CreateAPI(ctx, components.CreateAPIRequest{
        Name: "myAPI",
        Version: sdkkonnectgo.Pointer("v1"),
        Slug: sdkkonnectgo.Pointer("my-api-v1"),
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponseSchema != nil {
        // handle response
    }
}
```
### Example Usage: CreateApiMinimal

<!-- UsageSnippet language="go" operationID="create-api" method="post" path="/v3/apis" example="CreateApiMinimal" -->
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

    res, err := s.API.CreateAPI(ctx, components.CreateAPIRequest{
        Name: "myAPI",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponseSchema != nil {
        // handle response
    }
}
```
### Example Usage: CreateApiResponse

<!-- UsageSnippet language="go" operationID="create-api" method="post" path="/v3/apis" example="CreateApiResponse" -->
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

    res, err := s.API.CreateAPI(ctx, components.CreateAPIRequest{
        Name: "MyAPI",
        Slug: sdkkonnectgo.Pointer("my-api-v1"),
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponseSchema != nil {
        // handle response
    }
}
```
### Example Usage: UnsupportedMediaTypeExample

<!-- UsageSnippet language="go" operationID="create-api" method="post" path="/v3/apis" example="UnsupportedMediaTypeExample" -->
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

    res, err := s.API.CreateAPI(ctx, components.CreateAPIRequest{
        Name: "MyAPI",
        Slug: sdkkonnectgo.Pointer("my-api-v1"),
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.CreateAPIRequest](../../models/components/createapirequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.CreateAPIResponse](../../models/operations/createapiresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.BadRequestError           | 400                                 | application/problem+json            |
| sdkerrors.UnauthorizedError         | 401                                 | application/problem+json            |
| sdkerrors.ForbiddenError            | 403                                 | application/problem+json            |
| sdkerrors.NotFoundError             | 404                                 | application/problem+json            |
| sdkerrors.UnsupportedMediaTypeError | 415                                 | application/problem+json            |
| sdkerrors.SDKError                  | 4XX, 5XX                            | \*/\*                               |

## ListApis

Returns a collection of all APIs.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-apis" method="get" path="/v3/apis" example="ApiCollection" -->
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

    res, err := s.API.ListApis(ctx, operations.ListApisRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.APIFilterParameters{
            CreatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldGTEFilter(
                components.DateTimeFieldGTEFilter{
                    Lte: types.MustNewTimeFromString("2022-03-30T07:20:50Z"),
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
    if res.ListAPIResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.ListApisRequest](../../models/operations/listapisrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.ListApisResponse](../../models/operations/listapisresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchAPI

Get an API.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-api" method="get" path="/v3/apis/{apiId}" example="CreateApiResponse" -->
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

    res, err := s.API.FetchAPI(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5")
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `apiID`                                                  | *string*                                                 | :heavy_check_mark:                                       | The UUID API identifier                                  | 9f5061ce-78f6-4452-9108-ad7c02821fd5                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchAPIResponse](../../models/operations/fetchapiresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateAPI

Updates an API.

### Example Usage: ApiForbiddenExample

<!-- UsageSnippet language="go" operationID="update-api" method="patch" path="/v3/apis/{apiId}" example="ApiForbiddenExample" -->
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

    res, err := s.API.UpdateAPI(ctx, operations.UpdateAPIRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        UpdateAPIRequest: components.UpdateAPIRequest{
            Name: sdkkonnectgo.Pointer("MyAPI"),
            Slug: sdkkonnectgo.Pointer("my-api-v1"),
            Labels: map[string]*string{
                "env": sdkkonnectgo.Pointer("test"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponseSchema != nil {
        // handle response
    }
}
```
### Example Usage: ApiNameBadRequestExample

<!-- UsageSnippet language="go" operationID="update-api" method="patch" path="/v3/apis/{apiId}" example="ApiNameBadRequestExample" -->
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

    res, err := s.API.UpdateAPI(ctx, operations.UpdateAPIRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        UpdateAPIRequest: components.UpdateAPIRequest{
            Name: sdkkonnectgo.Pointer("MyAPI"),
            Slug: sdkkonnectgo.Pointer("my-api-v1"),
            Labels: map[string]*string{
                "env": sdkkonnectgo.Pointer("test"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponseSchema != nil {
        // handle response
    }
}
```
### Example Usage: ApiNotFoundExample

<!-- UsageSnippet language="go" operationID="update-api" method="patch" path="/v3/apis/{apiId}" example="ApiNotFoundExample" -->
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

    res, err := s.API.UpdateAPI(ctx, operations.UpdateAPIRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        UpdateAPIRequest: components.UpdateAPIRequest{
            Name: sdkkonnectgo.Pointer("MyAPI"),
            Slug: sdkkonnectgo.Pointer("my-api-v1"),
            Labels: map[string]*string{
                "env": sdkkonnectgo.Pointer("test"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponseSchema != nil {
        // handle response
    }
}
```
### Example Usage: ApiUnauthorizedExample

<!-- UsageSnippet language="go" operationID="update-api" method="patch" path="/v3/apis/{apiId}" example="ApiUnauthorizedExample" -->
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

    res, err := s.API.UpdateAPI(ctx, operations.UpdateAPIRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        UpdateAPIRequest: components.UpdateAPIRequest{
            Name: sdkkonnectgo.Pointer("MyAPI"),
            Slug: sdkkonnectgo.Pointer("my-api-v1"),
            Labels: map[string]*string{
                "env": sdkkonnectgo.Pointer("test"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponseSchema != nil {
        // handle response
    }
}
```
### Example Usage: CreateApiResponse

<!-- UsageSnippet language="go" operationID="update-api" method="patch" path="/v3/apis/{apiId}" example="CreateApiResponse" -->
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

    res, err := s.API.UpdateAPI(ctx, operations.UpdateAPIRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        UpdateAPIRequest: components.UpdateAPIRequest{
            Name: sdkkonnectgo.Pointer("MyAPI"),
            Slug: sdkkonnectgo.Pointer("my-api-v1"),
            Labels: map[string]*string{
                "env": sdkkonnectgo.Pointer("test"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponseSchema != nil {
        // handle response
    }
}
```
### Example Usage: UnsupportedMediaTypeExample

<!-- UsageSnippet language="go" operationID="update-api" method="patch" path="/v3/apis/{apiId}" example="UnsupportedMediaTypeExample" -->
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

    res, err := s.API.UpdateAPI(ctx, operations.UpdateAPIRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        UpdateAPIRequest: components.UpdateAPIRequest{
            Name: sdkkonnectgo.Pointer("MyAPI"),
            Slug: sdkkonnectgo.Pointer("my-api-v1"),
            Labels: map[string]*string{
                "env": sdkkonnectgo.Pointer("test"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponseSchema != nil {
        // handle response
    }
}
```
### Example Usage: UpdateApiName

<!-- UsageSnippet language="go" operationID="update-api" method="patch" path="/v3/apis/{apiId}" example="UpdateApiName" -->
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

    res, err := s.API.UpdateAPI(ctx, operations.UpdateAPIRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        UpdateAPIRequest: components.UpdateAPIRequest{
            Name: sdkkonnectgo.Pointer("new API name"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponseSchema != nil {
        // handle response
    }
}
```
### Example Usage: UpdateApiVersion

<!-- UsageSnippet language="go" operationID="update-api" method="patch" path="/v3/apis/{apiId}" example="UpdateApiVersion" -->
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

    res, err := s.API.UpdateAPI(ctx, operations.UpdateAPIRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        UpdateAPIRequest: components.UpdateAPIRequest{
            Version: sdkkonnectgo.Pointer("v2"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.UpdateAPIRequest](../../models/operations/updateapirequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.UpdateAPIResponse](../../models/operations/updateapiresponse.md), error**

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

## DeleteAPI

Deletes an API.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-api" method="delete" path="/v3/apis/{apiId}" -->
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

    res, err := s.API.DeleteAPI(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5")
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAPIResponse](../../models/operations/deleteapiresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |