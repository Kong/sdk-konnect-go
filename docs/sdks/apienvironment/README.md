# APIEnvironment

## Overview

### Available Operations

* [ListAPIEnvironments](#listapienvironments) - List API Environments
* [CreateAPIEnvironment](#createapienvironment) - Add an API Environment
* [FetchAPIEnvironment](#fetchapienvironment) - Get an API Environment
* [UpdateAPIEnvironment](#updateapienvironment) - Update an API Environment
* [DeleteAPIEnvironment](#deleteapienvironment) - Remove an API Environment

## ListAPIEnvironments

Returns the environments an API is associated with. Each record includes the
associated environment and the API version pinned to that environment.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-api-environments" method="get" path="/v3/apis/{apiId}/environments" -->
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

    res, err := s.APIEnvironment.ListAPIEnvironments(ctx, operations.ListAPIEnvironmentsRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAPIEnvironmentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListAPIEnvironmentsRequest](../../models/operations/listapienvironmentsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListAPIEnvironmentsResponse](../../models/operations/listapienvironmentsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateAPIEnvironment

Associates an API with an environment. Adding a second environment moves the
API into multi-environment mode. APIs implemented by a gateway service
(service-granularity) cannot be associated with more than one environment.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-api-environment" method="post" path="/v3/apis/{apiId}/environments" -->
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

    res, err := s.APIEnvironment.CreateAPIEnvironment(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPIEnvironment{
        EnvironmentID: "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7",
        VersionID: sdkkonnectgo.Pointer("7710d5c4-d902-410b-992f-18b814155b53"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIEnvironment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `apiID`                                                                            | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID API identifier                                                            | 9f5061ce-78f6-4452-9108-ad7c02821fd5                                               |
| `createAPIEnvironment`                                                             | [components.CreateAPIEnvironment](../../models/components/createapienvironment.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.CreateAPIEnvironmentResponse](../../models/operations/createapienvironmentresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchAPIEnvironment

Retrieve a single API-environment association by environment ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-api-environment" method="get" path="/v3/apis/{apiId}/environments/{environmentId}" -->
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

    res, err := s.APIEnvironment.FetchAPIEnvironment(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", "9ec7615a-288b-4a0a-b6a1-e0c49be23854")
    if err != nil {
        log.Fatal(err)
    }
    if res.APIEnvironment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                     | Type                                                                                                                                                          | Required                                                                                                                                                      | Description                                                                                                                                                   | Example                                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                                         | :heavy_check_mark:                                                                                                                                            | The context to use for the request.                                                                                                                           |                                                                                                                                                               |
| `apiID`                                                                                                                                                       | `string`                                                                                                                                                      | :heavy_check_mark:                                                                                                                                            | The UUID API identifier                                                                                                                                       | 9f5061ce-78f6-4452-9108-ad7c02821fd5                                                                                                                          |
| `environmentID`                                                                                                                                               | `string`                                                                                                                                                      | :heavy_check_mark:                                                                                                                                            | The environment's UUID. An environment is addressed by ID only; its name is a<br/>display label and is never resolved here, so a name in this position is a 404.<br/> | 9ec7615a-288b-4a0a-b6a1-e0c49be23854                                                                                                                          |
| `opts`                                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                                      | :heavy_minus_sign:                                                                                                                                            | The options for this request.                                                                                                                                 |                                                                                                                                                               |

### Response

**[*operations.FetchAPIEnvironmentResponse](../../models/operations/fetchapienvironmentresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateAPIEnvironment

Updates the API version pinned to this environment. Only the version pin can be
changed; the associated environment is immutable.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-api-environment" method="patch" path="/v3/apis/{apiId}/environments/{environmentId}" -->
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

    res, err := s.APIEnvironment.UpdateAPIEnvironment(ctx, operations.UpdateAPIEnvironmentRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        EnvironmentID: "9ec7615a-288b-4a0a-b6a1-e0c49be23854",
        UpdateAPIEnvironment: components.UpdateAPIEnvironment{
            VersionID: sdkkonnectgo.Pointer("7710d5c4-d902-410b-992f-18b814155b53"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIEnvironment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateAPIEnvironmentRequest](../../models/operations/updateapienvironmentrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateAPIEnvironmentResponse](../../models/operations/updateapienvironmentresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteAPIEnvironment

Removes an environment association from the API and hard-deletes the
implementations, publications, and application registrations scoped to it. An API
must always retain at least one environment; removing its last environment is
rejected.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-api-environment" method="delete" path="/v3/apis/{apiId}/environments/{environmentId}" -->
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

    res, err := s.APIEnvironment.DeleteAPIEnvironment(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", "9ec7615a-288b-4a0a-b6a1-e0c49be23854")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                     | Type                                                                                                                                                          | Required                                                                                                                                                      | Description                                                                                                                                                   | Example                                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                                         | :heavy_check_mark:                                                                                                                                            | The context to use for the request.                                                                                                                           |                                                                                                                                                               |
| `apiID`                                                                                                                                                       | `string`                                                                                                                                                      | :heavy_check_mark:                                                                                                                                            | The UUID API identifier                                                                                                                                       | 9f5061ce-78f6-4452-9108-ad7c02821fd5                                                                                                                          |
| `environmentID`                                                                                                                                               | `string`                                                                                                                                                      | :heavy_check_mark:                                                                                                                                            | The environment's UUID. An environment is addressed by ID only; its name is a<br/>display label and is never resolved here, so a name in this position is a 404.<br/> | 9ec7615a-288b-4a0a-b6a1-e0c49be23854                                                                                                                          |
| `opts`                                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                                      | :heavy_minus_sign:                                                                                                                                            | The options for this request.                                                                                                                                 |                                                                                                                                                               |

### Response

**[*operations.DeleteAPIEnvironmentResponse](../../models/operations/deleteapienvironmentresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |