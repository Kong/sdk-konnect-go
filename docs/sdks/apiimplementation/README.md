# APIImplementation

## Overview

### Available Operations

* [ListAPIImplementations](#listapiimplementations) - List API Implementations
* [CreateAPIImplementation](#createapiimplementation) - Create API Implementation
* [FetchAPIImplementation](#fetchapiimplementation) - Get an API Implementation
* [DeleteAPIImplementation](#deleteapiimplementation) - Delete API Implementation

## ListAPIImplementations

List gateway implementations for this API

### Example Usage

<!-- UsageSnippet language="go" operationID="list-api-implementations" method="get" path="/v3/api-implementations" -->
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

    res, err := s.APIImplementation.ListAPIImplementations(ctx, operations.ListAPIImplementationsRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.APIImplementationFilterParameters{
            CreatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTEFilter(
                components.DateTimeFieldLTEFilter{
                    Lte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
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
    if res.ListAPIImplementationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListAPIImplementationsRequest](../../models/operations/listapiimplementationsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListAPIImplementationsResponse](../../models/operations/listapiimplementationsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateAPIImplementation

Creates an implementation for an API.
If all operations in an API are implemented by a single gateway service and the service
has no routes that are not part of the API, then the API can be linked to the service.
For cases where an API is implemented by multiple gateway services, only a subset of
routes in one or more gateway services, or API operations need to be made available for API packages,
then the API should be linked to the control plane that defines the routes that overlap with the API.


### Example Usage: ApiForbiddenExample

<!-- UsageSnippet language="go" operationID="create-api-implementation" method="post" path="/v3/apis/{apiId}/implementations" example="ApiForbiddenExample" -->
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

    res, err := s.APIImplementation.CreateAPIImplementation(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIImplementationServiceReference(
        components.ServiceReference{
            Service: &components.APIImplementationService{
                ControlPlaneID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
                ID: "7710d5c4-d902-410b-992f-18b814155b53",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.APIImplementationResponse != nil {
        switch res.APIImplementationResponse.Type {
            case components.APIImplementationResponseTypeAPIImplementationResponseServiceReference:
                // res.APIImplementationResponse.APIImplementationResponseServiceReference is populated
            case components.APIImplementationResponseTypeAPIImplementationResponseControlPlaneReference:
                // res.APIImplementationResponse.APIImplementationResponseControlPlaneReference is populated
        }

    }
}
```
### Example Usage: ApiImplementationBadRequestExample

<!-- UsageSnippet language="go" operationID="create-api-implementation" method="post" path="/v3/apis/{apiId}/implementations" example="ApiImplementationBadRequestExample" -->
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

    res, err := s.APIImplementation.CreateAPIImplementation(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIImplementationServiceReference(
        components.ServiceReference{
            Service: &components.APIImplementationService{
                ControlPlaneID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
                ID: "7710d5c4-d902-410b-992f-18b814155b53",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.APIImplementationResponse != nil {
        switch res.APIImplementationResponse.Type {
            case components.APIImplementationResponseTypeAPIImplementationResponseServiceReference:
                // res.APIImplementationResponse.APIImplementationResponseServiceReference is populated
            case components.APIImplementationResponseTypeAPIImplementationResponseControlPlaneReference:
                // res.APIImplementationResponse.APIImplementationResponseControlPlaneReference is populated
        }

    }
}
```
### Example Usage: ApiImplementationCountConflictExample

<!-- UsageSnippet language="go" operationID="create-api-implementation" method="post" path="/v3/apis/{apiId}/implementations" example="ApiImplementationCountConflictExample" -->
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

    res, err := s.APIImplementation.CreateAPIImplementation(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIImplementationServiceReference(
        components.ServiceReference{
            Service: &components.APIImplementationService{
                ControlPlaneID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
                ID: "7710d5c4-d902-410b-992f-18b814155b53",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.APIImplementationResponse != nil {
        switch res.APIImplementationResponse.Type {
            case components.APIImplementationResponseTypeAPIImplementationResponseServiceReference:
                // res.APIImplementationResponse.APIImplementationResponseServiceReference is populated
            case components.APIImplementationResponseTypeAPIImplementationResponseControlPlaneReference:
                // res.APIImplementationResponse.APIImplementationResponseControlPlaneReference is populated
        }

    }
}
```
### Example Usage: ApiImplementationServiceConflictExample

<!-- UsageSnippet language="go" operationID="create-api-implementation" method="post" path="/v3/apis/{apiId}/implementations" example="ApiImplementationServiceConflictExample" -->
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

    res, err := s.APIImplementation.CreateAPIImplementation(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIImplementationServiceReference(
        components.ServiceReference{
            Service: &components.APIImplementationService{
                ControlPlaneID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
                ID: "7710d5c4-d902-410b-992f-18b814155b53",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.APIImplementationResponse != nil {
        switch res.APIImplementationResponse.Type {
            case components.APIImplementationResponseTypeAPIImplementationResponseServiceReference:
                // res.APIImplementationResponse.APIImplementationResponseServiceReference is populated
            case components.APIImplementationResponseTypeAPIImplementationResponseControlPlaneReference:
                // res.APIImplementationResponse.APIImplementationResponseControlPlaneReference is populated
        }

    }
}
```
### Example Usage: ApiNotFoundExample

<!-- UsageSnippet language="go" operationID="create-api-implementation" method="post" path="/v3/apis/{apiId}/implementations" example="ApiNotFoundExample" -->
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

    res, err := s.APIImplementation.CreateAPIImplementation(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIImplementationServiceReference(
        components.ServiceReference{
            Service: &components.APIImplementationService{
                ControlPlaneID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
                ID: "7710d5c4-d902-410b-992f-18b814155b53",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.APIImplementationResponse != nil {
        switch res.APIImplementationResponse.Type {
            case components.APIImplementationResponseTypeAPIImplementationResponseServiceReference:
                // res.APIImplementationResponse.APIImplementationResponseServiceReference is populated
            case components.APIImplementationResponseTypeAPIImplementationResponseControlPlaneReference:
                // res.APIImplementationResponse.APIImplementationResponseControlPlaneReference is populated
        }

    }
}
```
### Example Usage: ApiUnauthorizedExample

<!-- UsageSnippet language="go" operationID="create-api-implementation" method="post" path="/v3/apis/{apiId}/implementations" example="ApiUnauthorizedExample" -->
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

    res, err := s.APIImplementation.CreateAPIImplementation(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIImplementationServiceReference(
        components.ServiceReference{
            Service: &components.APIImplementationService{
                ControlPlaneID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
                ID: "7710d5c4-d902-410b-992f-18b814155b53",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.APIImplementationResponse != nil {
        switch res.APIImplementationResponse.Type {
            case components.APIImplementationResponseTypeAPIImplementationResponseServiceReference:
                // res.APIImplementationResponse.APIImplementationResponseServiceReference is populated
            case components.APIImplementationResponseTypeAPIImplementationResponseControlPlaneReference:
                // res.APIImplementationResponse.APIImplementationResponseControlPlaneReference is populated
        }

    }
}
```
### Example Usage: CreateApiImplementationRequest

<!-- UsageSnippet language="go" operationID="create-api-implementation" method="post" path="/v3/apis/{apiId}/implementations" example="CreateApiImplementationRequest" -->
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

    res, err := s.APIImplementation.CreateAPIImplementation(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIImplementationServiceReference(
        components.ServiceReference{
            Service: &components.APIImplementationService{
                ControlPlaneID: "fd4e1b98-3629-4dd3-acc0-759a726ffee2",
                ID: "dd4e1b98-3629-4dd3-acc0-759a726ffee2",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.APIImplementationResponse != nil {
        switch res.APIImplementationResponse.Type {
            case components.APIImplementationResponseTypeAPIImplementationResponseServiceReference:
                // res.APIImplementationResponse.APIImplementationResponseServiceReference is populated
            case components.APIImplementationResponseTypeAPIImplementationResponseControlPlaneReference:
                // res.APIImplementationResponse.APIImplementationResponseControlPlaneReference is populated
        }

    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |                                                                              |
| `apiID`                                                                      | *string*                                                                     | :heavy_check_mark:                                                           | The UUID API identifier                                                      | 9f5061ce-78f6-4452-9108-ad7c02821fd5                                         |
| `apiImplementation`                                                          | [components.APIImplementation](../../models/components/apiimplementation.md) | :heavy_check_mark:                                                           | N/A                                                                          |                                                                              |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |                                                                              |

### Response

**[*operations.CreateAPIImplementationResponse](../../models/operations/createapiimplementationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchAPIImplementation

Retrieve a gateway implementation for this API

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-api-implementation" method="get" path="/v3/apis/{apiId}/implementations/{implementationId}" -->
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

    res, err := s.APIImplementation.FetchAPIImplementation(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", "032d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.APIImplementationResponse != nil {
        switch res.APIImplementationResponse.Type {
            case components.APIImplementationResponseTypeAPIImplementationResponseServiceReference:
                // res.APIImplementationResponse.APIImplementationResponseServiceReference is populated
            case components.APIImplementationResponseTypeAPIImplementationResponseControlPlaneReference:
                // res.APIImplementationResponse.APIImplementationResponseControlPlaneReference is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `apiID`                                                  | *string*                                                 | :heavy_check_mark:                                       | The UUID API identifier                                  | 9f5061ce-78f6-4452-9108-ad7c02821fd5                     |
| `implementationID`                                       | *string*                                                 | :heavy_check_mark:                                       | The Portal identifier                                    | 032d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchAPIImplementationResponse](../../models/operations/fetchapiimplementationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteAPIImplementation

Unlink a gateway implementation from this API

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-api-implementation" method="delete" path="/v3/apis/{apiId}/implementations/{implementationId}" -->
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

    res, err := s.APIImplementation.DeleteAPIImplementation(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", "032d905a-ed33-46a3-a093-d8f536af9a8a")
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
| `implementationID`                                       | *string*                                                 | :heavy_check_mark:                                       | The Portal identifier                                    | 032d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAPIImplementationResponse](../../models/operations/deleteapiimplementationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |