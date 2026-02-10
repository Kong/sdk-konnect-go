# AppAuthStrategies

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

### Example Usage: AppAuthStrategyResponseExample1

<!-- UsageSnippet language="go" operationID="create-app-auth-strategy" method="post" path="/v2/application-auth-strategies" example="AppAuthStrategyResponseExample1" -->
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

    res, err := s.AppAuthStrategies.CreateAppAuthStrategy(ctx, components.CreateCreateAppAuthStrategyRequestOpenidConnect(
        components.AppAuthStrategyOpenIDConnectRequest{
            Name: "<value>",
            DisplayName: "Jalen_Cummerata42",
            StrategyType: components.AppAuthStrategyOpenIDConnectRequestStrategyTypeOpenidConnect,
            Configs: components.AppAuthStrategyOpenIDConnectRequestConfigs{
                OpenidConnect: components.AppAuthStrategyConfigOpenIDConnect{
                    Issuer: "https://tangible-diver.com",
                    CredentialClaim: []string{
                        "<value 1>",
                    },
                    Scopes: []string{
                        "<value 1>",
                        "<value 2>",
                    },
                    AuthMethods: []string{
                        "<value 1>",
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
### Example Usage: CreateAppAuthStrategyKeyAuthRequest

<!-- UsageSnippet language="go" operationID="create-app-auth-strategy" method="post" path="/v2/application-auth-strategies" example="CreateAppAuthStrategyKeyAuthRequest" -->
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
### Example Usage: CreateAppAuthStrategyOpenIDConnectDCRRequest

<!-- UsageSnippet language="go" operationID="create-app-auth-strategy" method="post" path="/v2/application-auth-strategies" example="CreateAppAuthStrategyOpenIDConnectDCRRequest" -->
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

    res, err := s.AppAuthStrategies.CreateAppAuthStrategy(ctx, components.CreateCreateAppAuthStrategyRequestOpenidConnect(
        components.AppAuthStrategyOpenIDConnectRequest{
            Name: "auth strategy 3",
            DisplayName: "Enterprise Auth with DCR",
            StrategyType: components.AppAuthStrategyOpenIDConnectRequestStrategyTypeOpenidConnect,
            Configs: components.AppAuthStrategyOpenIDConnectRequestConfigs{
                OpenidConnect: components.AppAuthStrategyConfigOpenIDConnect{
                    Issuer: "https://oidc-dcr.example.com",
                    CredentialClaim: []string{
                        "client_id",
                    },
                    Scopes: []string{
                        "openid",
                        "email",
                    },
                    AuthMethods: []string{
                        "client_credentials",
                        "bearer",
                    },
                },
            },
            DcrProviderID: sdkkonnectgo.Pointer("73f8380e-7798-4566-99e3-2edf2b57d281"),
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
### Example Usage: CreateAppAuthStrategyOpenIDConnectRequest

<!-- UsageSnippet language="go" operationID="create-app-auth-strategy" method="post" path="/v2/application-auth-strategies" example="CreateAppAuthStrategyOpenIDConnectRequest" -->
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

    res, err := s.AppAuthStrategies.CreateAppAuthStrategy(ctx, components.CreateCreateAppAuthStrategyRequestOpenidConnect(
        components.AppAuthStrategyOpenIDConnectRequest{
            Name: "auth strategy 2",
            DisplayName: "Enterprise Auth",
            StrategyType: components.AppAuthStrategyOpenIDConnectRequestStrategyTypeOpenidConnect,
            Configs: components.AppAuthStrategyOpenIDConnectRequestConfigs{
                OpenidConnect: components.AppAuthStrategyConfigOpenIDConnect{
                    Issuer: "https://oidc.example.com",
                    CredentialClaim: []string{
                        "client_id",
                    },
                    Scopes: []string{
                        "openid",
                        "email",
                    },
                    AuthMethods: []string{
                        "client_credentials",
                        "bearer",
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
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="create-app-auth-strategy" method="post" path="/v2/application-auth-strategies" example="Unauthorized" -->
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

    res, err := s.AppAuthStrategies.CreateAppAuthStrategy(ctx, components.CreateCreateAppAuthStrategyRequestKeyAuth(
        components.AppAuthStrategyKeyAuthRequest{
            Name: "<value>",
            DisplayName: "Carlee_Kozey70",
            StrategyType: components.StrategyTypeKeyAuth,
            Configs: components.AppAuthStrategyKeyAuthRequestConfigs{
                KeyAuth: components.AppAuthStrategyConfigKeyAuth{},
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
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="create-app-auth-strategy" method="post" path="/v2/application-auth-strategies" example="UnauthorizedExample" -->
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

    res, err := s.AppAuthStrategies.CreateAppAuthStrategy(ctx, components.CreateCreateAppAuthStrategyRequestKeyAuth(
        components.AppAuthStrategyKeyAuthRequest{
            Name: "<value>",
            DisplayName: "Wyatt.Schmitt",
            StrategyType: components.StrategyTypeKeyAuth,
            Configs: components.AppAuthStrategyKeyAuthRequestConfigs{
                KeyAuth: components.AppAuthStrategyConfigKeyAuth{},
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

<!-- UsageSnippet language="go" operationID="list-app-auth-strategies" method="get" path="/v2/application-auth-strategies" example="AppAuthStrategyResponseExample1" -->
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

    res, err := s.AppAuthStrategies.ListAppAuthStrategies(ctx, operations.ListAppAuthStrategiesRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
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

<!-- UsageSnippet language="go" operationID="get-app-auth-strategy" method="get" path="/v2/application-auth-strategies/{authStrategyId}" example="AppAuthStrategyResponseExample1" -->
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

### Example Usage: AppAuthStrategyResponseExample1

<!-- UsageSnippet language="go" operationID="replace-app-auth-strategy" method="put" path="/v2/application-auth-strategies/{authStrategyId}" example="AppAuthStrategyResponseExample1" -->
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

    res, err := s.AppAuthStrategies.ReplaceAppAuthStrategy(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.CreateCreateAppAuthStrategyRequestKeyAuth(
        components.AppAuthStrategyKeyAuthRequest{
            Name: "<value>",
            DisplayName: "Camille45",
            StrategyType: components.StrategyTypeKeyAuth,
            Configs: components.AppAuthStrategyKeyAuthRequestConfigs{
                KeyAuth: components.AppAuthStrategyConfigKeyAuth{},
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
### Example Usage: CreateAppAuthStrategyKeyAuthRequest

<!-- UsageSnippet language="go" operationID="replace-app-auth-strategy" method="put" path="/v2/application-auth-strategies/{authStrategyId}" example="CreateAppAuthStrategyKeyAuthRequest" -->
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
### Example Usage: CreateAppAuthStrategyOpenIDConnectDCRRequest

<!-- UsageSnippet language="go" operationID="replace-app-auth-strategy" method="put" path="/v2/application-auth-strategies/{authStrategyId}" example="CreateAppAuthStrategyOpenIDConnectDCRRequest" -->
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

    res, err := s.AppAuthStrategies.ReplaceAppAuthStrategy(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.CreateCreateAppAuthStrategyRequestOpenidConnect(
        components.AppAuthStrategyOpenIDConnectRequest{
            Name: "auth strategy 3",
            DisplayName: "Enterprise Auth with DCR",
            StrategyType: components.AppAuthStrategyOpenIDConnectRequestStrategyTypeOpenidConnect,
            Configs: components.AppAuthStrategyOpenIDConnectRequestConfigs{
                OpenidConnect: components.AppAuthStrategyConfigOpenIDConnect{
                    Issuer: "https://oidc-dcr.example.com",
                    CredentialClaim: []string{
                        "client_id",
                    },
                    Scopes: []string{
                        "openid",
                        "email",
                    },
                    AuthMethods: []string{
                        "client_credentials",
                        "bearer",
                    },
                },
            },
            DcrProviderID: sdkkonnectgo.Pointer("73f8380e-7798-4566-99e3-2edf2b57d281"),
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
### Example Usage: CreateAppAuthStrategyOpenIDConnectRequest

<!-- UsageSnippet language="go" operationID="replace-app-auth-strategy" method="put" path="/v2/application-auth-strategies/{authStrategyId}" example="CreateAppAuthStrategyOpenIDConnectRequest" -->
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

    res, err := s.AppAuthStrategies.ReplaceAppAuthStrategy(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.CreateCreateAppAuthStrategyRequestOpenidConnect(
        components.AppAuthStrategyOpenIDConnectRequest{
            Name: "auth strategy 2",
            DisplayName: "Enterprise Auth",
            StrategyType: components.AppAuthStrategyOpenIDConnectRequestStrategyTypeOpenidConnect,
            Configs: components.AppAuthStrategyOpenIDConnectRequestConfigs{
                OpenidConnect: components.AppAuthStrategyConfigOpenIDConnect{
                    Issuer: "https://oidc.example.com",
                    CredentialClaim: []string{
                        "client_id",
                    },
                    Scopes: []string{
                        "openid",
                        "email",
                    },
                    AuthMethods: []string{
                        "client_credentials",
                        "bearer",
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
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="replace-app-auth-strategy" method="put" path="/v2/application-auth-strategies/{authStrategyId}" example="Unauthorized" -->
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

    res, err := s.AppAuthStrategies.ReplaceAppAuthStrategy(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.CreateCreateAppAuthStrategyRequestOpenidConnect(
        components.AppAuthStrategyOpenIDConnectRequest{
            Name: "<value>",
            DisplayName: "Aliyah_Fadel",
            StrategyType: components.AppAuthStrategyOpenIDConnectRequestStrategyTypeOpenidConnect,
            Configs: components.AppAuthStrategyOpenIDConnectRequestConfigs{
                OpenidConnect: components.AppAuthStrategyConfigOpenIDConnect{
                    Issuer: "https://tedious-tinderbox.org/",
                    CredentialClaim: []string{
                        "<value 1>",
                        "<value 2>",
                        "<value 3>",
                    },
                    Scopes: []string{},
                    AuthMethods: []string{},
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
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="replace-app-auth-strategy" method="put" path="/v2/application-auth-strategies/{authStrategyId}" example="UnauthorizedExample" -->
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

    res, err := s.AppAuthStrategies.ReplaceAppAuthStrategy(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.CreateCreateAppAuthStrategyRequestKeyAuth(
        components.AppAuthStrategyKeyAuthRequest{
            Name: "<value>",
            DisplayName: "Leone_Veum48",
            StrategyType: components.StrategyTypeKeyAuth,
            Configs: components.AppAuthStrategyKeyAuthRequestConfigs{
                KeyAuth: components.AppAuthStrategyConfigKeyAuth{},
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

### Example Usage: AppAuthStrategyResponseExample1

<!-- UsageSnippet language="go" operationID="update-app-auth-strategy" method="patch" path="/v2/application-auth-strategies/{authStrategyId}" example="AppAuthStrategyResponseExample1" -->
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

    res, err := s.AppAuthStrategies.UpdateAppAuthStrategy(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.UpdateAppAuthStrategyRequest{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppAuthStrategyResponse != nil {
        // handle response
    }
}
```
### Example Usage: NotFoundExample

<!-- UsageSnippet language="go" operationID="update-app-auth-strategy" method="patch" path="/v2/application-auth-strategies/{authStrategyId}" example="NotFoundExample" -->
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

    res, err := s.AppAuthStrategies.UpdateAppAuthStrategy(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.UpdateAppAuthStrategyRequest{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppAuthStrategyResponse != nil {
        // handle response
    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="update-app-auth-strategy" method="patch" path="/v2/application-auth-strategies/{authStrategyId}" example="Unauthorized" -->
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

    res, err := s.AppAuthStrategies.UpdateAppAuthStrategy(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.UpdateAppAuthStrategyRequest{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppAuthStrategyResponse != nil {
        // handle response
    }
}
```
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="update-app-auth-strategy" method="patch" path="/v2/application-auth-strategies/{authStrategyId}" example="UnauthorizedExample" -->
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

    res, err := s.AppAuthStrategies.UpdateAppAuthStrategy(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.UpdateAppAuthStrategyRequest{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppAuthStrategyResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateAppAuthStrategyKeyAuthRequest

<!-- UsageSnippet language="go" operationID="update-app-auth-strategy" method="patch" path="/v2/application-auth-strategies/{authStrategyId}" example="UpdateAppAuthStrategyKeyAuthRequest" -->
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

    res, err := s.AppAuthStrategies.UpdateAppAuthStrategy(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.UpdateAppAuthStrategyRequest{
        DisplayName: sdkkonnectgo.Pointer("API Key"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppAuthStrategyResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateAppAuthStrategyOpenIDConnectRequest

<!-- UsageSnippet language="go" operationID="update-app-auth-strategy" method="patch" path="/v2/application-auth-strategies/{authStrategyId}" example="UpdateAppAuthStrategyOpenIDConnectRequest" -->
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

    res, err := s.AppAuthStrategies.UpdateAppAuthStrategy(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.UpdateAppAuthStrategyRequest{
        DisplayName: sdkkonnectgo.Pointer("Client Credentials"),
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

<!-- UsageSnippet language="go" operationID="delete-app-auth-strategy" method="delete" path="/v2/application-auth-strategies/{authStrategyId}" -->
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