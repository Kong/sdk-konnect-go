# ControlPlaneResourceQuotas

## Overview

### Available Operations

* [ListControlPlaneDefaultResourceQuotas](#listcontrolplanedefaultresourcequotas) - List Default Quotas
* [ListControlPlaneResourceQuota](#listcontrolplaneresourcequota) - List Quota Overrides
* [CreateControlPlaneResourceQuota](#createcontrolplaneresourcequota) - Create a control plane resource quota
* [GetControlPlaneResourceQuota](#getcontrolplaneresourcequota) - Get Control Plane Resource Quota
* [UpdateControlPlaneResourceQuota](#updatecontrolplaneresourcequota) - Update Resource Quota
* [DeleteControlPlaneResourceQuota](#deletecontrolplaneresourcequota) - Delete the control plane resource quota

## ListControlPlaneDefaultResourceQuotas

Returns a paginated collection of default resource quotas for control planes, along with
organizationally-defined overrides for those resource quotas.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-control-plane-default-resource-quotas" method="get" path="/v2/control-planes/default-resource-quotas" -->
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

    res, err := s.ControlPlaneResourceQuotas.ListControlPlaneDefaultResourceQuotas(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListControlPlaneDefaultResourceQuotasResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListControlPlaneDefaultResourceQuotasResponse](../../models/operations/listcontrolplanedefaultresourcequotasresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.BaseError         | 500, 503                    | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListControlPlaneResourceQuota

Returns a list of control plane quota override values created by the user.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-control-plane-resource-quota" method="get" path="/v2/control-planes/resource-quotas" -->
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

    res, err := s.ControlPlaneResourceQuotas.ListControlPlaneResourceQuota(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListControlPlaneResourceQuotasResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListControlPlaneResourceQuotaResponse](../../models/operations/listcontrolplaneresourcequotaresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.BaseError         | 500, 503                    | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateControlPlaneResourceQuota

Creates a control plane resource quota scoped to a given resource, for an organization.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-control-plane-resource-quota" method="post" path="/v2/control-planes/resource-quotas" -->
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

    res, err := s.ControlPlaneResourceQuotas.CreateControlPlaneResourceQuota(ctx, components.ControlPlaneResourceQuotaInput{
        Name: sdkkonnectgo.Pointer("Consumer"),
        Description: sdkkonnectgo.Pointer("Across the organization, the aggregate number of consumers cannot exceed this value."),
        Resource: "count/consumer",
        Value: 5000,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlaneResourceQuota != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [components.ControlPlaneResourceQuotaInput](../../models/components/controlplaneresourcequotainput.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateControlPlaneResourceQuotaResponse](../../models/operations/createcontrolplaneresourcequotaresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.BaseError         | 500, 503                    | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetControlPlaneResourceQuota

Returns the control plane resource quota for the ID.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-control-plane-resource-quota" method="get" path="/v2/control-planes/resource-quotas/{resourceQuotaId}" -->
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

    res, err := s.ControlPlaneResourceQuotas.GetControlPlaneResourceQuota(ctx, "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlaneResourceQuota != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `resourceQuotaID`                                        | `string`                                                 | :heavy_check_mark:                                       | The UUID for the control plane resource quota.           | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetControlPlaneResourceQuotaResponse](../../models/operations/getcontrolplaneresourcequotaresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.BaseError         | 500, 503                    | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateControlPlaneResourceQuota

Update the control plane resource quota for the corresponding ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-control-plane-resource-quota" method="patch" path="/v2/control-planes/resource-quotas/{resourceQuotaId}" -->
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

    res, err := s.ControlPlaneResourceQuotas.UpdateControlPlaneResourceQuota(ctx, "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7", components.PatchControlPlaneResourceQuota{
        Value: 5000,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlaneResourceQuota != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `resourceQuotaID`                                                                                      | `string`                                                                                               | :heavy_check_mark:                                                                                     | The UUID for the control plane resource quota.                                                         | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                                   |
| `patchControlPlaneResourceQuota`                                                                       | [components.PatchControlPlaneResourceQuota](../../models/components/patchcontrolplaneresourcequota.md) | :heavy_check_mark:                                                                                     | Request body for updating the control plane override resource quota value.                             |                                                                                                        |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |                                                                                                        |

### Response

**[*operations.UpdateControlPlaneResourceQuotaResponse](../../models/operations/updatecontrolplaneresourcequotaresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.BaseError         | 500, 503                    | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteControlPlaneResourceQuota

Deletes the created control plane resource quota which matches the ID provided.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-control-plane-resource-quota" method="delete" path="/v2/control-planes/resource-quotas/{resourceQuotaId}" -->
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

    res, err := s.ControlPlaneResourceQuotas.DeleteControlPlaneResourceQuota(ctx, "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
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
| `resourceQuotaID`                                        | `string`                                                 | :heavy_check_mark:                                       | The UUID for the control plane resource quota.           | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteControlPlaneResourceQuotaResponse](../../models/operations/deletecontrolplaneresourcequotaresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.BaseError         | 500, 503                    | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |