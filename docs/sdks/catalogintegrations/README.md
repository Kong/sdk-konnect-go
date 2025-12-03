# CatalogIntegrations
(*CatalogIntegrations*)

## Overview

Integrations are applications, either Konnect-internal or external, which extend the functionality of the Service Catalog.
Install and authorize an integration to discover the resources across your organization which support your services.
Map relevant resources to your services to provide a rich view of cataloged services.
To set up and view a list of all the integrations we support please view our [documentation](https://developer.konghq.com/service-catalog/integrations/).


### Available Operations

* [ListCatalogIntegrations](#listcatalogintegrations) - List Integrations

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