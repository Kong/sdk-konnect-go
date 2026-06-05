# CatalogIntegrations

## Overview

Integrations are applications, either Konnect-internal or external, which extend the functionality of the Service Catalog.
Install and authorize an integration to discover the resources across your organization which support your services.
Map relevant resources to your services to provide a rich view of cataloged services.
To set up and view a list of all the integrations we support please view our [documentation](https://developer.konghq.com/service-catalog/integrations/).


### Available Operations

* [CreateCatalogIntegration](#createcatalogintegration) - Create Integration
* [ListCatalogIntegrations](#listcatalogintegrations) - List Integrations
* [GetCatalogIntegration](#getcatalogintegration) - Get Integration
* [UpdateCatalogIntegration](#updatecatalogintegration) - Update Integration
* [DeleteCatalogIntegration](#deletecatalogintegration) - Delete Integration

## CreateCatalogIntegration

Creates a catalog integration.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-catalog-integration" method="post" path="/v1/integrations" -->
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

    res, err := s.CatalogIntegrations.CreateCatalogIntegration(ctx, components.CreateCreateCatalogIntegrationCreatePrivateCatalogIntegration(
        components.CreatePrivateCatalogIntegration{
            Visibility: components.CreatePrivateCatalogIntegrationVisibilityPrivate,
            ResourceTypes: map[string]components.CatalogIntegrationResourceTypes{
                "gateway_svc": components.CatalogIntegrationResourceTypes{
                    DisplayName: sdkkonnectgo.Pointer("Gateway Service"),
                    ResourceIDTemplate: "{{control_plane_id}}:{{gateway_service_id}}",
                    Schema: components.SimpleSchema{
                        Type: components.SimpleSchemaTypeSimple,
                        Definition: map[string]components.Definition{
                            "control_plane_id": components.DefinitionString,
                            "gatway_service_id": components.DefinitionString,
                        },
                    },
                    IntegrationDataSchema: nil,
                },
                "analytics_dashboard": components.CatalogIntegrationResourceTypes{
                    DisplayName: sdkkonnectgo.Pointer("Dashboard"),
                    ResourceIDTemplate: "{{dashboard_id}}",
                    Schema: components.SimpleSchema{
                        Type: components.SimpleSchemaTypeSimple,
                        Definition: map[string]components.Definition{
                            "dashboard_id": components.DefinitionString,
                        },
                    },
                    IntegrationDataSchema: nil,
                },
            },
            Description: "pinstripe onto zowie amid interviewer alongside clueless",
            DisplayName: "My Integration",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogIntegration != nil {
        switch res.CatalogIntegration.Authorization.Type {
            case components.CatalogIntegrationAuthorizationTypeOne:
                // res.CatalogIntegration.Authorization.One is populated
            case components.CatalogIntegrationAuthorizationTypeOAuth:
                // res.CatalogIntegration.Authorization.OAuth is populated
            case components.CatalogIntegrationAuthorizationTypeMultiKeyAuth:
                // res.CatalogIntegration.Authorization.MultiKeyAuth is populated
            case components.CatalogIntegrationAuthorizationTypeGitHubAppInstallationAuth:
                // res.CatalogIntegration.Authorization.GitHubAppInstallationAuth is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.CreateCatalogIntegration](../../models/components/createcatalogintegration.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateCatalogIntegrationResponse](../../models/operations/createcatalogintegrationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCatalogIntegrations

Returns a paginated collection of all catalog integrations available to connect.
Each integration represents a built-in connector that extends the platform's capabilities;
enabling discovery, resource management, event ingestion, and more.
Integrations expose metadata that describes how they authenticate, what configuration they
require, and how they interact with catalog entities like Resources and Services.

Currently, integrations are platform-defined and cannot be extend or registered by customers.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-catalog-integrations" method="get" path="/v1/integrations" -->
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

    res, err := s.CatalogIntegrations.ListCatalogIntegrations(ctx, operations.ListCatalogIntegrationsRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCatalogIntegrationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListCatalogIntegrationsRequest](../../models/operations/listcatalogintegrationsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListCatalogIntegrationsResponse](../../models/operations/listcatalogintegrationsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCatalogIntegration

Gets a catalog integration.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-catalog-integration" method="get" path="/v1/integrations/{integration}" -->
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

    res, err := s.CatalogIntegrations.GetCatalogIntegration(ctx, "jira")
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogIntegration != nil {
        switch res.CatalogIntegration.Authorization.Type {
            case components.CatalogIntegrationAuthorizationTypeOne:
                // res.CatalogIntegration.Authorization.One is populated
            case components.CatalogIntegrationAuthorizationTypeOAuth:
                // res.CatalogIntegration.Authorization.OAuth is populated
            case components.CatalogIntegrationAuthorizationTypeMultiKeyAuth:
                // res.CatalogIntegration.Authorization.MultiKeyAuth is populated
            case components.CatalogIntegrationAuthorizationTypeGitHubAppInstallationAuth:
                // res.CatalogIntegration.Authorization.GitHubAppInstallationAuth is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `integration`                                            | `string`                                                 | :heavy_check_mark:                                       | The name of the integration.                             | jira                                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetCatalogIntegrationResponse](../../models/operations/getcatalogintegrationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateCatalogIntegration

Updates a catalog integration.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-catalog-integration" method="patch" path="/v1/integrations/{integration}" -->
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

    res, err := s.CatalogIntegrations.UpdateCatalogIntegration(ctx, "jira", components.UpdateCatalogIntegration{
        DisplayName: sdkkonnectgo.Pointer("Updated Catalog Integration"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogIntegration != nil {
        switch res.CatalogIntegration.Authorization.Type {
            case components.CatalogIntegrationAuthorizationTypeOne:
                // res.CatalogIntegration.Authorization.One is populated
            case components.CatalogIntegrationAuthorizationTypeOAuth:
                // res.CatalogIntegration.Authorization.OAuth is populated
            case components.CatalogIntegrationAuthorizationTypeMultiKeyAuth:
                // res.CatalogIntegration.Authorization.MultiKeyAuth is populated
            case components.CatalogIntegrationAuthorizationTypeGitHubAppInstallationAuth:
                // res.CatalogIntegration.Authorization.GitHubAppInstallationAuth is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `integration`                                                                              | `string`                                                                                   | :heavy_check_mark:                                                                         | The name of the integration.                                                               | jira                                                                                       |
| `updateCatalogIntegration`                                                                 | [components.UpdateCatalogIntegration](../../models/components/updatecatalogintegration.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |

### Response

**[*operations.UpdateCatalogIntegrationResponse](../../models/operations/updatecatalogintegrationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteCatalogIntegration

Deletes a catalog integration.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-catalog-integration" method="delete" path="/v1/integrations/{integration}" -->
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

    res, err := s.CatalogIntegrations.DeleteCatalogIntegration(ctx, "jira")
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
| `integration`                                            | `string`                                                 | :heavy_check_mark:                                       | The name of the integration.                             | jira                                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteCatalogIntegrationResponse](../../models/operations/deletecatalogintegrationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |