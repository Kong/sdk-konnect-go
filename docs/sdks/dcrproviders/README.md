# DCRProviders

## Overview

Dynamic Client Registration Providers are configurations representing an external Identity Provider whose clients (i.e. Applications) Konnect will be authorized to manage.
For instance, they will be able to perform dynamic client registration (DCR) with the provider. 
The DCR provider provides credentials to each DCR-enabled application in Konnect that can be used to access Product Versions that the app is registered for.


### Available Operations

* [CreateDcrProvider](#createdcrprovider) - Create DCR provider
* [ListDcrProviders](#listdcrproviders) - List DCR Providers
* [GetDcrProvider](#getdcrprovider) - Get a DCR provider
* [UpdateDcrProvider](#updatedcrprovider) - Update DCR provider
* [DeleteDcrProvider](#deletedcrprovider) - Delete DCR provider
* [VerifyDcrProvider](#verifydcrprovider) - Verify DCR provider configuration

## CreateDcrProvider

Creates a DCR provider.

### Example Usage: DcrProviderAuth0

<!-- UsageSnippet language="go" operationID="create-dcr-provider" method="post" path="/v2/dcr-providers" example="DcrProviderAuth0" -->
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

    res, err := s.DCRProviders.CreateDcrProvider(ctx, components.CreateCreateDcrProviderRequestAuth0(
        components.CreateDcrProviderRequestAuth0{
            ProviderType: components.ProviderTypeAuth0,
            DcrConfig: components.CreateDcrConfigAuth0InRequest{
                InitialClientID: "<id>",
                InitialClientSecret: "<value>",
            },
            Name: "<value>",
            Issuer: "https://physical-pension.biz",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateDcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: DcrProviderAuth0CreateRequest

<!-- UsageSnippet language="go" operationID="create-dcr-provider" method="post" path="/v2/dcr-providers" example="DcrProviderAuth0CreateRequest" -->
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

    res, err := s.DCRProviders.CreateDcrProvider(ctx, components.CreateCreateDcrProviderRequestAuth0(
        components.CreateDcrProviderRequestAuth0{
            ProviderType: components.ProviderTypeAuth0,
            DcrConfig: components.CreateDcrConfigAuth0InRequest{
                InitialClientID: "abc123",
                InitialClientSecret: "abc123xyz098!",
                InitialClientAudience: sdkkonnectgo.Pointer("https://my-custom-domain.com/api/v2/"),
            },
            Name: "DCR Auth0 1 - Segment A",
            Issuer: "https://my-issuer.auth0.com/api/v2/",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateDcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: DcrProviderAzureAd

<!-- UsageSnippet language="go" operationID="create-dcr-provider" method="post" path="/v2/dcr-providers" example="DcrProviderAzureAd" -->
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

    res, err := s.DCRProviders.CreateDcrProvider(ctx, components.CreateCreateDcrProviderRequestOkta(
        components.CreateDcrProviderRequestOkta{
            ProviderType: components.CreateDcrProviderRequestOktaProviderTypeOkta,
            DcrConfig: components.CreateDcrConfigOktaInRequest{
                DcrToken: "<value>",
            },
            Name: "<value>",
            Issuer: "https://sinful-operating.org",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateDcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: DcrProviderAzureAdCreateRequest

<!-- UsageSnippet language="go" operationID="create-dcr-provider" method="post" path="/v2/dcr-providers" example="DcrProviderAzureAdCreateRequest" -->
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

    res, err := s.DCRProviders.CreateDcrProvider(ctx, components.CreateCreateDcrProviderRequestAzureAd(
        components.CreateDcrProviderRequestAzureAd{
            ProviderType: components.CreateDcrProviderRequestAzureAdProviderTypeAzureAd,
            DcrConfig: components.CreateDcrConfigAzureAdInRequest{
                InitialClientID: "abc123",
                InitialClientSecret: "abc123xyz098!",
            },
            Name: "DCR Azure 1 - Segment A",
            Issuer: "https://sts.windows.net/my-tenant",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateDcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: DcrProviderCurity

<!-- UsageSnippet language="go" operationID="create-dcr-provider" method="post" path="/v2/dcr-providers" example="DcrProviderCurity" -->
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

    res, err := s.DCRProviders.CreateDcrProvider(ctx, components.CreateCreateDcrProviderRequestCurity(
        components.CreateDcrProviderRequestCurity{
            ProviderType: components.CreateDcrProviderRequestCurityProviderTypeCurity,
            DcrConfig: components.CreateDcrConfigCurityInRequest{
                InitialClientID: "<id>",
                InitialClientSecret: "<value>",
            },
            Name: "<value>",
            Issuer: "https://snarling-punctuation.com",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateDcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: DcrProviderCurityCreateRequest

<!-- UsageSnippet language="go" operationID="create-dcr-provider" method="post" path="/v2/dcr-providers" example="DcrProviderCurityCreateRequest" -->
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

    res, err := s.DCRProviders.CreateDcrProvider(ctx, components.CreateCreateDcrProviderRequestCurity(
        components.CreateDcrProviderRequestCurity{
            ProviderType: components.CreateDcrProviderRequestCurityProviderTypeCurity,
            DcrConfig: components.CreateDcrConfigCurityInRequest{
                InitialClientID: "abc123",
                InitialClientSecret: "abc123xyz098!",
            },
            Name: "DCR Curity 1 - Segment A",
            Issuer: "https://my-curity-instance.com/oauth/v2/oauth-anonymous/.well-known/openid-configuration",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateDcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: DcrProviderHttp

<!-- UsageSnippet language="go" operationID="create-dcr-provider" method="post" path="/v2/dcr-providers" example="DcrProviderHttp" -->
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

    res, err := s.DCRProviders.CreateDcrProvider(ctx, components.CreateCreateDcrProviderRequestHTTP(
        components.CreateDcrProviderRequestHTTP{
            ProviderType: components.CreateDcrProviderRequestHTTPProviderTypeHTTP,
            DcrConfig: components.CreateDcrConfigHTTPInRequest{
                DcrBaseURL: "https://wilted-information.info",
                APIKey: "<value>",
            },
            Name: "<value>",
            Issuer: "https://darling-slime.net",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateDcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: DcrProviderHttpCreateRequest

<!-- UsageSnippet language="go" operationID="create-dcr-provider" method="post" path="/v2/dcr-providers" example="DcrProviderHttpCreateRequest" -->
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

    res, err := s.DCRProviders.CreateDcrProvider(ctx, components.CreateCreateDcrProviderRequestHTTP(
        components.CreateDcrProviderRequestHTTP{
            ProviderType: components.CreateDcrProviderRequestHTTPProviderTypeHTTP,
            DcrConfig: components.CreateDcrConfigHTTPInRequest{
                DcrBaseURL: "https://my-http-dcr-server.com/v1/dcr",
                APIKey: "gYmrbDfu_7PTsZWH",
            },
            Name: "DCR HTTP 1 - Segment A",
            Issuer: "https://my-issuer-server.com",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateDcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: DcrProviderOkta

<!-- UsageSnippet language="go" operationID="create-dcr-provider" method="post" path="/v2/dcr-providers" example="DcrProviderOkta" -->
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

    res, err := s.DCRProviders.CreateDcrProvider(ctx, components.CreateCreateDcrProviderRequestAuth0(
        components.CreateDcrProviderRequestAuth0{
            ProviderType: components.ProviderTypeAuth0,
            DcrConfig: components.CreateDcrConfigAuth0InRequest{
                InitialClientID: "<id>",
                InitialClientSecret: "<value>",
            },
            Name: "<value>",
            Issuer: "https://lively-bracelet.info/",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateDcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: DcrProviderOktaCreateRequest

<!-- UsageSnippet language="go" operationID="create-dcr-provider" method="post" path="/v2/dcr-providers" example="DcrProviderOktaCreateRequest" -->
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

    res, err := s.DCRProviders.CreateDcrProvider(ctx, components.CreateCreateDcrProviderRequestOkta(
        components.CreateDcrProviderRequestOkta{
            ProviderType: components.CreateDcrProviderRequestOktaProviderTypeOkta,
            DcrConfig: components.CreateDcrConfigOktaInRequest{
                DcrToken: "abc123xyz098!",
            },
            Name: "DCR Okta 1 - Segment A",
            Issuer: "https://my-issuer.okta.com/default",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateDcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="create-dcr-provider" method="post" path="/v2/dcr-providers" example="Unauthorized" -->
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

    res, err := s.DCRProviders.CreateDcrProvider(ctx, components.CreateCreateDcrProviderRequestHTTP(
        components.CreateDcrProviderRequestHTTP{
            ProviderType: components.CreateDcrProviderRequestHTTPProviderTypeHTTP,
            DcrConfig: components.CreateDcrConfigHTTPInRequest{
                DcrBaseURL: "https://wilted-information.info",
                APIKey: "<value>",
            },
            Name: "<value>",
            Issuer: "https://general-circumference.biz/",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateDcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="create-dcr-provider" method="post" path="/v2/dcr-providers" example="UnauthorizedExample" -->
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

    res, err := s.DCRProviders.CreateDcrProvider(ctx, components.CreateCreateDcrProviderRequestAuth0(
        components.CreateDcrProviderRequestAuth0{
            ProviderType: components.ProviderTypeAuth0,
            DcrConfig: components.CreateDcrConfigAuth0InRequest{
                InitialClientID: "<id>",
                InitialClientSecret: "<value>",
            },
            Name: "<value>",
            Issuer: "https://merry-ownership.name",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateDcrProviderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.CreateDcrProviderRequest](../../models/components/createdcrproviderrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateDcrProviderResponse](../../models/operations/createdcrproviderresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListDcrProviders

Returns a paginated collection of DCR providers.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-dcr-providers" method="get" path="/v2/dcr-providers" example="ListDcrProviders" -->
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

    res, err := s.DCRProviders.ListDcrProviders(ctx, operations.ListDcrProvidersRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListDcrProvidersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListDcrProvidersRequest](../../models/operations/listdcrprovidersrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListDcrProvidersResponse](../../models/operations/listdcrprovidersresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetDcrProvider

Returns a DCR provider.

### Example Usage: DcrProviderAuth0

<!-- UsageSnippet language="go" operationID="get-dcr-provider" method="get" path="/v2/dcr-providers/{dcrProviderId}" example="DcrProviderAuth0" -->
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

    res, err := s.DCRProviders.GetDcrProvider(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
    if err != nil {
        log.Fatal(err)
    }
    if res.DcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: DcrProviderAzureAd

<!-- UsageSnippet language="go" operationID="get-dcr-provider" method="get" path="/v2/dcr-providers/{dcrProviderId}" example="DcrProviderAzureAd" -->
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

    res, err := s.DCRProviders.GetDcrProvider(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
    if err != nil {
        log.Fatal(err)
    }
    if res.DcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: DcrProviderCurity

<!-- UsageSnippet language="go" operationID="get-dcr-provider" method="get" path="/v2/dcr-providers/{dcrProviderId}" example="DcrProviderCurity" -->
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

    res, err := s.DCRProviders.GetDcrProvider(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
    if err != nil {
        log.Fatal(err)
    }
    if res.DcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: DcrProviderHttp

<!-- UsageSnippet language="go" operationID="get-dcr-provider" method="get" path="/v2/dcr-providers/{dcrProviderId}" example="DcrProviderHttp" -->
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

    res, err := s.DCRProviders.GetDcrProvider(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
    if err != nil {
        log.Fatal(err)
    }
    if res.DcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: DcrProviderOkta

<!-- UsageSnippet language="go" operationID="get-dcr-provider" method="get" path="/v2/dcr-providers/{dcrProviderId}" example="DcrProviderOkta" -->
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

    res, err := s.DCRProviders.GetDcrProvider(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
    if err != nil {
        log.Fatal(err)
    }
    if res.DcrProviderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `dcrProviderID`                                          | *string*                                                 | :heavy_check_mark:                                       | DCR provider identifier                                  | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetDcrProviderResponse](../../models/operations/getdcrproviderresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateDcrProvider

Updates a DCR provider.

### Example Usage: DcrProviderAuth0

<!-- UsageSnippet language="go" operationID="update-dcr-provider" method="patch" path="/v2/dcr-providers/{dcrProviderId}" example="DcrProviderAuth0" -->
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

    res, err := s.DCRProviders.UpdateDcrProvider(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.UpdateDcrProviderRequest{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: DcrProviderAzureAd

<!-- UsageSnippet language="go" operationID="update-dcr-provider" method="patch" path="/v2/dcr-providers/{dcrProviderId}" example="DcrProviderAzureAd" -->
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

    res, err := s.DCRProviders.UpdateDcrProvider(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.UpdateDcrProviderRequest{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: DcrProviderCurity

<!-- UsageSnippet language="go" operationID="update-dcr-provider" method="patch" path="/v2/dcr-providers/{dcrProviderId}" example="DcrProviderCurity" -->
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

    res, err := s.DCRProviders.UpdateDcrProvider(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.UpdateDcrProviderRequest{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: DcrProviderHttp

<!-- UsageSnippet language="go" operationID="update-dcr-provider" method="patch" path="/v2/dcr-providers/{dcrProviderId}" example="DcrProviderHttp" -->
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

    res, err := s.DCRProviders.UpdateDcrProvider(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.UpdateDcrProviderRequest{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: DcrProviderOkta

<!-- UsageSnippet language="go" operationID="update-dcr-provider" method="patch" path="/v2/dcr-providers/{dcrProviderId}" example="DcrProviderOkta" -->
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

    res, err := s.DCRProviders.UpdateDcrProvider(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.UpdateDcrProviderRequest{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: DcrProviderRequest

<!-- UsageSnippet language="go" operationID="update-dcr-provider" method="patch" path="/v2/dcr-providers/{dcrProviderId}" example="DcrProviderRequest" -->
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

    res, err := s.DCRProviders.UpdateDcrProvider(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.UpdateDcrProviderRequest{
        Name: sdkkonnectgo.Pointer("DCR Okta 1 - Segment A"),
        Issuer: sdkkonnectgo.Pointer("https://my-issuer.okta.com/default"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: NotFoundExample

<!-- UsageSnippet language="go" operationID="update-dcr-provider" method="patch" path="/v2/dcr-providers/{dcrProviderId}" example="NotFoundExample" -->
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

    res, err := s.DCRProviders.UpdateDcrProvider(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.UpdateDcrProviderRequest{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="update-dcr-provider" method="patch" path="/v2/dcr-providers/{dcrProviderId}" example="Unauthorized" -->
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

    res, err := s.DCRProviders.UpdateDcrProvider(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.UpdateDcrProviderRequest{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="update-dcr-provider" method="patch" path="/v2/dcr-providers/{dcrProviderId}" example="UnauthorizedExample" -->
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

    res, err := s.DCRProviders.UpdateDcrProvider(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.UpdateDcrProviderRequest{
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DcrProviderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `dcrProviderID`                                                                            | *string*                                                                                   | :heavy_check_mark:                                                                         | DCR provider identifier                                                                    | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                       |
| `updateDcrProviderRequest`                                                                 | [components.UpdateDcrProviderRequest](../../models/components/updatedcrproviderrequest.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |

### Response

**[*operations.UpdateDcrProviderResponse](../../models/operations/updatedcrproviderresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteDcrProvider

Deletes a DCR provider.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-dcr-provider" method="delete" path="/v2/dcr-providers/{dcrProviderId}" -->
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

    res, err := s.DCRProviders.DeleteDcrProvider(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
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
| `dcrProviderID`                                          | *string*                                                 | :heavy_check_mark:                                       | DCR provider identifier                                  | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteDcrProviderResponse](../../models/operations/deletedcrproviderresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## VerifyDcrProvider

Verifies if a DCR provider is configured properly. Returns 200 for success, 4xx for failure.

### Example Usage: VerifyDcrProviderFailed

<!-- UsageSnippet language="go" operationID="verify-dcr-provider" method="post" path="/v2/dcr-providers/{dcrProviderId}/verify" example="VerifyDcrProviderFailed" -->
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

    res, err := s.DCRProviders.VerifyDcrProvider(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
    if err != nil {
        log.Fatal(err)
    }
    if res.VerifyDcrProviderResponse != nil {
        // handle response
    }
}
```
### Example Usage: VerifyDcrProviderSuccess

<!-- UsageSnippet language="go" operationID="verify-dcr-provider" method="post" path="/v2/dcr-providers/{dcrProviderId}/verify" example="VerifyDcrProviderSuccess" -->
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

    res, err := s.DCRProviders.VerifyDcrProvider(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
    if err != nil {
        log.Fatal(err)
    }
    if res.VerifyDcrProviderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `dcrProviderID`                                          | *string*                                                 | :heavy_check_mark:                                       | DCR provider identifier                                  | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.VerifyDcrProviderResponse](../../models/operations/verifydcrproviderresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |