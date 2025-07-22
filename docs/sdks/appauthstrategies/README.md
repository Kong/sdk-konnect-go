# AppAuthStrategies
(*AppAuthStrategies*)

## Overview

Application Auth Strategies are sets of plugin configurations that represent how the gateway will perform authentication and authorization for a Product Version. 
Called “Auth Strategy” for short in the context of portals/applications. 
The plugins are synced to any Gateway Service that is currently linked or becomes linked to the Product Version.


### Available Operations

* [CreateAppAuthStrategy](#createappauthstrategy) - Create App Auth Strategy
* [ListAppAuthStrategies](#listappauthstrategies) - List App Auth Strategies
* [GetAppAuthStrategy](#getappauthstrategy) - Get App Auth Strategy
* [ReplaceAppAuthStrategy](#replaceappauthstrategy) - Replace App Auth Strategy
* [UpdateAppAuthStrategy](#updateappauthstrategy) - Update App Auth Strategy
* [DeleteAppAuthStrategy](#deleteappauthstrategy) - Delete App Auth Strategy

## CreateAppAuthStrategy

Creates an application auth strategy.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.AppAuthStrategies.CreateAppAuthStrategy(ctx, components.CreateCreateAppAuthStrategyRequestKeyAuth(
        components.AppAuthStrategyKeyAuthRequest{
            Name: "auth strategy 1",
            DisplayName: "API Key Auth",
            StrategyType: components.StrategyTypeKeyAuth,
            Configs: components.AppAuthStrategyKeyAuthRequestConfigs{
                KeyAuth: components.AppAuthStrategyConfigKeyAuth{
                    KeyNames: []string{
                        "apikey",
                        "api-key",
                        "x-api-key",
                    },
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppAuthStrategyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [components.CreateAppAuthStrategyRequest](../../models/components/createappauthstrategyrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateAppAuthStrategyResponse](../../models/operations/createappauthstrategyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListAppAuthStrategies

Returns a paginated collection of application auth strategies.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.AppAuthStrategies.ListAppAuthStrategies(ctx, operations.ListAppAuthStrategiesRequest{
        PageSize: sdkkonnectgo.Int64(10),
        PageNumber: sdkkonnectgo.Int64(1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAppAuthStrategiesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListAppAuthStrategiesRequest](../../models/operations/listappauthstrategiesrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListAppAuthStrategiesResponse](../../models/operations/listappauthstrategiesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetAppAuthStrategy

Returns an application auth strategy.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.AppAuthStrategies.GetAppAuthStrategy(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppAuthStrategyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `authStrategyID`                                         | *string*                                                 | :heavy_check_mark:                                       | Application auth strategy identifier                     | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAppAuthStrategyResponse](../../models/operations/getappauthstrategyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ReplaceAppAuthStrategy

Replaces an application auth strategy.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.AppAuthStrategies.ReplaceAppAuthStrategy(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.CreateCreateAppAuthStrategyRequestKeyAuth(
        components.AppAuthStrategyKeyAuthRequest{
            Name: "auth strategy 1",
            DisplayName: "API Key Auth",
            StrategyType: components.StrategyTypeKeyAuth,
            Configs: components.AppAuthStrategyKeyAuthRequestConfigs{
                KeyAuth: components.AppAuthStrategyConfigKeyAuth{
                    KeyNames: []string{
                        "apikey",
                        "api-key",
                        "x-api-key",
                    },
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppAuthStrategyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |                                                                                                    |
| `authStrategyID`                                                                                   | *string*                                                                                           | :heavy_check_mark:                                                                                 | Application auth strategy identifier                                                               | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                               |
| `createAppAuthStrategyRequest`                                                                     | [components.CreateAppAuthStrategyRequest](../../models/components/createappauthstrategyrequest.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |                                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |                                                                                                    |

### Response

**[*operations.ReplaceAppAuthStrategyResponse](../../models/operations/replaceappauthstrategyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateAppAuthStrategy

Updates an application auth strategy.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.AppAuthStrategies.UpdateAppAuthStrategy(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.UpdateAppAuthStrategyRequest{
        DisplayName: sdkkonnectgo.String("API Key"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppAuthStrategyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |                                                                                                    |
| `authStrategyID`                                                                                   | *string*                                                                                           | :heavy_check_mark:                                                                                 | Application auth strategy identifier                                                               | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                               |
| `updateAppAuthStrategyRequest`                                                                     | [components.UpdateAppAuthStrategyRequest](../../models/components/updateappauthstrategyrequest.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |                                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |                                                                                                    |

### Response

**[*operations.UpdateAppAuthStrategyResponse](../../models/operations/updateappauthstrategyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteAppAuthStrategy

Deletes an application auth strategy. An application auth strategy can be deleted ONLY if it's not used by any product version within any portal regardless of their publication statuses. If an application auth strategy is still in use the request will result in an HTTP 409 CONFLICT.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.AppAuthStrategies.DeleteAppAuthStrategy(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
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
| `authStrategyID`                                         | *string*                                                 | :heavy_check_mark:                                       | Application auth strategy identifier                     | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAppAuthStrategyResponse](../../models/operations/deleteappauthstrategyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |