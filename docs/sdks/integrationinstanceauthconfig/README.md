# IntegrationInstanceAuthConfig

## Overview

A integration instance may need to be provided with an auth config before authorizing the instance.
Typically an auth config will be required when authorizing against a integration which is hosted within your organization.
The integration instance's auth config will inform how the Service Catalog will authorize the integration instance.
Note that updating the auth config for an instance which is already authorized will **remove the existing credential**, requiring you to re-authorize the instance.


### Available Operations

* [GetIntegrationInstanceAuthConfig](#getintegrationinstanceauthconfig) - Get Integration Instance Auth Config
* [UpsertIntegrationInstanceAuthConfig](#upsertintegrationinstanceauthconfig) - Upsert Integration Instance Auth Config
* [DeleteIntegrationInstanceAuthConfig](#deleteintegrationinstanceauthconfig) - Delete Integration Instance Auth Config

## GetIntegrationInstanceAuthConfig

Fetches auth config scoped to the given integration instance.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-integration-instance-auth-config" method="get" path="/v1/integration-instances/{integrationInstanceId}/auth-config" -->
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

    res, err := s.IntegrationInstanceAuthConfig.GetIntegrationInstanceAuthConfig(ctx, "3f51fa25-310a-421d-bd1a-007f859021a3")
    if err != nil {
        log.Fatal(err)
    }
    if res.IntegrationInstanceAuthConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `integrationInstanceID`                                  | *string*                                                 | :heavy_check_mark:                                       | The `id` of the integration instance.                    | 3f51fa25-310a-421d-bd1a-007f859021a3                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetIntegrationInstanceAuthConfigResponse](../../models/operations/getintegrationinstanceauthconfigresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpsertIntegrationInstanceAuthConfig

Upserts auth config scoped to the given integration instance.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-integration-instance-auth-config" method="put" path="/v1/integration-instances/{integrationInstanceId}/auth-config" -->
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

    res, err := s.IntegrationInstanceAuthConfig.UpsertIntegrationInstanceAuthConfig(ctx, "3f51fa25-310a-421d-bd1a-007f859021a3", components.CreateUpsertIntegrationInstanceAuthConfigOauthConfig(
        components.OauthConfig{
            ClientID: "d745213a-b7e8-4998-abe3-41f164001970",
            ClientSecret: "s3cr3t4p1cl13ntt0k3n1234567890abcdef",
            AuthorizationEndpoint: "https://identity.service.com/oauth/authorize",
            TokenEndpoint: "https://identity.service.com/oauth/token",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.IntegrationInstanceAuthConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      | Example                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |                                                                                                                  |
| `integrationInstanceID`                                                                                          | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The `id` of the integration instance.                                                                            | 3f51fa25-310a-421d-bd1a-007f859021a3                                                                             |
| `upsertIntegrationInstanceAuthConfig`                                                                            | [components.UpsertIntegrationInstanceAuthConfig](../../models/components/upsertintegrationinstanceauthconfig.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |                                                                                                                  |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |                                                                                                                  |

### Response

**[*operations.UpsertIntegrationInstanceAuthConfigResponse](../../models/operations/upsertintegrationinstanceauthconfigresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteIntegrationInstanceAuthConfig

Deletes the auth config scoped to the given integration instance.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-integration-instance-auth-config" method="delete" path="/v1/integration-instances/{integrationInstanceId}/auth-config" -->
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

    res, err := s.IntegrationInstanceAuthConfig.DeleteIntegrationInstanceAuthConfig(ctx, "3f51fa25-310a-421d-bd1a-007f859021a3")
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
| `integrationInstanceID`                                  | *string*                                                 | :heavy_check_mark:                                       | The `id` of the integration instance.                    | 3f51fa25-310a-421d-bd1a-007f859021a3                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteIntegrationInstanceAuthConfigResponse](../../models/operations/deleteintegrationinstanceauthconfigresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |