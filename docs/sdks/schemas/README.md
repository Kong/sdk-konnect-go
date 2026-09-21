# Schemas

## Overview

### Available Operations

* [FetchCoreEntityJSONSchema](#fetchcoreentityjsonschema) - Get core entity JSON schema
* [ValidateEntitySchema](#validateentityschema) - Validate entity schema
* [FetchPartialSchema](#fetchpartialschema) - Get partial schema

## FetchCoreEntityJSONSchema

Returns the JSON schema for a Kong Gateway core entity.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-core-entity-json-schema" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/v1/schemas/json/{entityName}" -->
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

    res, err := s.Schemas.FetchCoreEntityJSONSchema(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "upstream")
    if err != nil {
        log.Fatal(err)
    }
    if res.CoreEntityJSONSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `entityName`                                                                       | `string`                                                                           | :heavy_check_mark:                                                                 | The name of the core entity                                                        | upstream                                                                           |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.FetchCoreEntityJSONSchemaResponse](../../models/operations/fetchcoreentityjsonschemaresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.KonnectCPLegacyUnauthorizedError | 401                                        | application/json                           |
| sdkerrors.KonnectCPLegacyForbiddenError    | 403                                        | application/json                           |
| sdkerrors.KonnectCPLegacyNotFoundError     | 404                                        | application/json                           |
| sdkerrors.SDKError                         | 4XX, 5XX                                   | \*/\*                                      |

## ValidateEntitySchema

Validate schema for an entity

### Example Usage

<!-- UsageSnippet language="go" operationID="validate-entity-schema" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/schemas/{entityName}/validate" -->
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

    res, err := s.Schemas.ValidateEntitySchema(ctx, operations.ValidateEntitySchemaRequest{
        EntityName: "<value>",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ValidateEntityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ValidateEntitySchemaRequest](../../models/operations/validateentityschemarequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ValidateEntitySchemaResponse](../../models/operations/validateentityschemaresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## FetchPartialSchema

Get the schema for a partial

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-partial-schema" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/schemas/partials/{partialType}" -->
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

    res, err := s.Schemas.FetchPartialSchema(ctx, "<value>", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetPartialSchemaResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `partialType`                                                                      | `string`                                                                           | :heavy_check_mark:                                                                 | The type of a partial                                                              |                                                                                    |
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.FetchPartialSchemaResponse](../../models/operations/fetchpartialschemaresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |