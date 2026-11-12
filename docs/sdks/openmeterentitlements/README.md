# OpenMeterEntitlements

## Overview

Entitlements are used to control access to features for customers.

### Available Operations

* [ListCustomerEntitlementAccess](#listcustomerentitlementaccess) - List customer entitlement access
* [GetCustomerEntitlementAccess](#getcustomerentitlementaccess) - Get customer entitlement access
* [GetCustomerEntitlementValueByFeatureKey](#getcustomerentitlementvaluebyfeaturekey) - Get customer entitlement value by feature key
* [CreateCustomerEntitlement](#createcustomerentitlement) - Create customer entitlement
* [ListCustomerEntitlements](#listcustomerentitlements) - List customer entitlements
* [GetCustomerEntitlement](#getcustomerentitlement) - Get customer entitlement
* [DeleteCustomerEntitlement](#deletecustomerentitlement) - Delete customer entitlement
* [CreateCustomerEntitlementGrant](#createcustomerentitlementgrant) - Create customer entitlement grant
* [ListCustomerEntitlementGrants](#listcustomerentitlementgrants) - List customer entitlement grants
* [GetCustomerEntitlementHistory](#getcustomerentitlementhistory) - Get customer entitlement history
* [OverrideCustomerEntitlement](#overridecustomerentitlement) - Override customer entitlement
* [ResetCustomerEntitlementUsage](#resetcustomerentitlementusage) - Reset customer entitlement usage
* [GetCustomerEntitlementValue](#getcustomerentitlementvalue) - Get customer entitlement value
* [QueryEntitlementAccess](#queryentitlementaccess) - Query entitlement access
* [ListEntitlements](#listentitlements) - List entitlements
* [GetEntitlement](#getentitlement) - Get entitlement
* [ListGrants](#listgrants) - List grants
* [VoidGrant](#voidgrant) - Void grant

## ListCustomerEntitlementAccess

List customer entitlement access

### Example Usage

<!-- UsageSnippet language="go" operationID="list-customer-entitlement-access" method="get" path="/v3/openmeter/customers/{customerId}/entitlement-access" -->
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

    res, err := s.OpenMeterEntitlements.ListCustomerEntitlementAccess(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCustomerEntitlementAccessResponseData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `customerID`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListCustomerEntitlementAccessResponse](../../models/operations/listcustomerentitlementaccessresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCustomerEntitlementAccess

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Get the customer's access to a single feature.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-customer-entitlement-access" method="get" path="/v3/openmeter/customers/{customerId}/entitlement-access/features/{featureKey}" -->
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

    res, err := s.OpenMeterEntitlements.GetCustomerEntitlementAccess(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", "resource_key")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingEntitlementAccessCheckResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `customerID`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `featureKey`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | resource_key                                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetCustomerEntitlementAccessResponse](../../models/operations/getcustomerentitlementaccessresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCustomerEntitlementValueByFeatureKey

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Get the customer's entitlement value for a feature at a point in time. Without
an active entitlement, the result denies access and omits the type.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-customer-entitlement-value-by-feature-key" method="get" path="/v3/openmeter/customers/{customerId}/entitlement-access/features/{featureKey}/value" -->
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

    res, err := s.OpenMeterEntitlements.GetCustomerEntitlementValueByFeatureKey(ctx, operations.GetCustomerEntitlementValueByFeatureKeyRequest{
        CustomerID: "01G65Z755AFWAKHE12NY0CQ9FH",
        FeatureKey: "resource_key",
        At: types.MustNewTimeFromString("2023-01-01T01:01:01.001Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingEntitlementFeatureValueResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.GetCustomerEntitlementValueByFeatureKeyRequest](../../models/operations/getcustomerentitlementvaluebyfeaturekeyrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.GetCustomerEntitlementValueByFeatureKeyResponse](../../models/operations/getcustomerentitlementvaluebyfeaturekeyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateCustomerEntitlement

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Create an entitlement for the customer.

A customer can have only one active entitlement per feature. The feature must be
compatible with the entitlement type. Entitlements cannot be modified after
creation, only deleted.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-customer-entitlement" method="post" path="/v3/openmeter/customers/{customerId}/entitlements" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.OpenMeterEntitlements.CreateCustomerEntitlement(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.CreateCreateEntitlementRequestMetered(
        components.CreateEntitlementMeteredRequest{
            Type: components.CreateEntitlementMeteredRequestTypeMetered,
            Feature: components.CreateEntitlementMeteredRequestFeature{
                ID: "01G65Z755AFWAKHE12NY0CQ9FH",
            },
            UsagePeriod: components.UsagePeriod{
                Interval: "P1M",
                Anchor: types.MustNewTimeFromString("2023-01-01T01:01:01.001Z"),
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingEntitlement != nil {
        switch res.BillingEntitlement.Type {
            case components.BillingEntitlementTypeMetered:
                // res.BillingEntitlement.BillingEntitlementMetered is populated
            case components.BillingEntitlementTypeStatic:
                // res.BillingEntitlement.BillingEntitlementStatic is populated
            case components.BillingEntitlementTypeBoolean:
                // res.BillingEntitlement.BillingEntitlementBoolean is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `customerID`                                                                               | `string`                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        | 01G65Z755AFWAKHE12NY0CQ9FH                                                                 |
| `createEntitlementRequest`                                                                 | [components.CreateEntitlementRequest](../../models/components/createentitlementrequest.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |

### Response

**[*operations.CreateCustomerEntitlementResponse](../../models/operations/createcustomerentitlementresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCustomerEntitlements

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

List the entitlements of the customer that are active at the time of the
request. For checking entitlement access, use the entitlement access endpoints
instead.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-customer-entitlements" method="get" path="/v3/openmeter/customers/{customerId}/entitlements" -->
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

    res, err := s.OpenMeterEntitlements.ListCustomerEntitlements(ctx, operations.ListCustomerEntitlementsRequest{
        CustomerID: "01G65Z755AFWAKHE12NY0CQ9FH",
        Sort: sdkkonnectgo.Pointer("created_at desc"),
        Filter: &components.ListCustomerEntitlementsParamsFilter{
            FeatureID: sdkkonnectgo.Pointer(components.CreateListCustomerEntitlementsParamsFilterULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntitlementPagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListCustomerEntitlementsRequest](../../models/operations/listcustomerentitlementsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListCustomerEntitlementsResponse](../../models/operations/listcustomerentitlementsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCustomerEntitlement

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Get an entitlement of the customer by ID. For checking entitlement access, use
the entitlement access endpoints instead.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-customer-entitlement" method="get" path="/v3/openmeter/customers/{customerId}/entitlements/{entitlementId}" -->
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

    res, err := s.OpenMeterEntitlements.GetCustomerEntitlement(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingEntitlement != nil {
        switch res.BillingEntitlement.Type {
            case components.BillingEntitlementTypeMetered:
                // res.BillingEntitlement.BillingEntitlementMetered is populated
            case components.BillingEntitlementTypeStatic:
                // res.BillingEntitlement.BillingEntitlementStatic is populated
            case components.BillingEntitlementTypeBoolean:
                // res.BillingEntitlement.BillingEntitlementBoolean is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `customerID`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `entitlementID`                                          | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetCustomerEntitlementResponse](../../models/operations/getcustomerentitlementresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteCustomerEntitlement

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Deletes the entitlement and revokes access to its feature. A customer can hold
only one active entitlement per feature, so migrating a feature requires
deleting the previous entitlement first.

Deletion sets the `deleted_at` timestamp instead of removing history. Access and
status queries for earlier points in time still treat the entitlement as active,
so access changes are never retroactive.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-customer-entitlement" method="delete" path="/v3/openmeter/customers/{customerId}/entitlements/{entitlementId}" -->
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

    res, err := s.OpenMeterEntitlements.DeleteCustomerEntitlement(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", "01G65Z755AFWAKHE12NY0CQ9FH")
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
| `customerID`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `entitlementID`                                          | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteCustomerEntitlementResponse](../../models/operations/deletecustomerentitlementresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateCustomerEntitlementGrant

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Issue a grant for a metered entitlement of the customer. Boolean and static
entitlements cannot have grants, so the request is rejected for them.

Grants are immutable. The amount is added to the balance from `effective_at`,
which cannot be earlier than the start of the current usage period.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-customer-entitlement-grant" method="post" path="/v3/openmeter/customers/{customerId}/entitlements/{entitlementId}/grants" -->
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

    res, err := s.OpenMeterEntitlements.CreateCustomerEntitlementGrant(ctx, operations.CreateCustomerEntitlementGrantRequest{
        CustomerID: "01G65Z755AFWAKHE12NY0CQ9FH",
        EntitlementID: "01G65Z755AFWAKHE12NY0CQ9FH",
        BillingEntitlementGrantCreateRequest: components.BillingEntitlementGrantCreateRequest{
            Amount: "100",
            Priority: sdkkonnectgo.Pointer[int64](1),
            EffectiveAt: types.MustTimeFromString("2023-01-01T01:01:01.001Z"),
            ExpiresAfter: sdkkonnectgo.Pointer("P1M"),
            MaxRolloverAmount: sdkkonnectgo.Pointer("100"),
            MinRolloverAmount: sdkkonnectgo.Pointer("0"),
            Labels: map[string]string{
                "env": "test",
            },
            Recurrence: &components.Recurrence{
                Interval: "P1M",
                Anchor: types.MustNewTimeFromString("2023-01-01T01:01:01.001Z"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingEntitlementGrant != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.CreateCustomerEntitlementGrantRequest](../../models/operations/createcustomerentitlementgrantrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.CreateCustomerEntitlementGrantResponse](../../models/operations/createcustomerentitlementgrantresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCustomerEntitlementGrants

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

List the grants issued for an entitlement of the customer. Grants only exist for
metered entitlements, so the list is empty for boolean and static entitlements.

Deleted grants are excluded unless `include_deleted` is set. Voided and expired
grants are always included, as they are part of the balance history.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-customer-entitlement-grants" method="get" path="/v3/openmeter/customers/{customerId}/entitlements/{entitlementId}/grants" -->
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

    res, err := s.OpenMeterEntitlements.ListCustomerEntitlementGrants(ctx, operations.ListCustomerEntitlementGrantsRequest{
        CustomerID: "01G65Z755AFWAKHE12NY0CQ9FH",
        EntitlementID: "01G65Z755AFWAKHE12NY0CQ9FH",
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntitlementGrantPagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ListCustomerEntitlementGrantsRequest](../../models/operations/listcustomerentitlementgrantsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ListCustomerEntitlementGrantsResponse](../../models/operations/listcustomerentitlementgrantsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCustomerEntitlementHistory

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Get the balance and usage history of a metered entitlement. The queried range
may span multiple usage periods.

`windowed_history` groups usage into windows of the requested size and reports
the balance at the start of each window. `burndown_history` lists the periods in
which grants were consumed in a fixed order, together with the usage taken from
each grant.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-customer-entitlement-history" method="get" path="/v3/openmeter/customers/{customerId}/entitlements/{entitlementId}/history" -->
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

    res, err := s.OpenMeterEntitlements.GetCustomerEntitlementHistory(ctx, operations.GetCustomerEntitlementHistoryRequest{
        CustomerID: "01G65Z755AFWAKHE12NY0CQ9FH",
        EntitlementID: "01G65Z755AFWAKHE12NY0CQ9FH",
        From: types.MustNewTimeFromString("2023-01-01T01:01:01.001Z"),
        To: types.MustNewTimeFromString("2023-01-01T01:01:01.001Z"),
        WindowSize: components.BillingEntitlementHistoryWindowSizePt1H,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingEntitlementHistory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetCustomerEntitlementHistoryRequest](../../models/operations/getcustomerentitlementhistoryrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.GetCustomerEntitlementHistoryResponse](../../models/operations/getcustomerentitlementhistoryresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## OverrideCustomerEntitlement

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Override an entitlement of the customer with a new one.

The referenced entitlement ends and the new one starts at the same instant, so
access continues without a gap. Both must belong to the same feature. Use this
for upgrades and downgrades.

Fails if the referenced entitlement does not exist, is deleted, or is no longer
active.

### Example Usage

<!-- UsageSnippet language="go" operationID="override-customer-entitlement" method="put" path="/v3/openmeter/customers/{customerId}/entitlements/{entitlementId}/override" -->
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

    res, err := s.OpenMeterEntitlements.OverrideCustomerEntitlement(ctx, operations.OverrideCustomerEntitlementRequest{
        CustomerID: "01G65Z755AFWAKHE12NY0CQ9FH",
        EntitlementID: "01G65Z755AFWAKHE12NY0CQ9FH",
        CreateEntitlementRequest: components.CreateCreateEntitlementRequestBoolean(
            components.CreateEntitlementBooleanRequest{
                Type: components.CreateEntitlementBooleanRequestTypeBoolean,
                Feature: components.CreateEntitlementBooleanRequestFeature{
                    ID: "01G65Z755AFWAKHE12NY0CQ9FH",
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingEntitlement != nil {
        switch res.BillingEntitlement.Type {
            case components.BillingEntitlementTypeMetered:
                // res.BillingEntitlement.BillingEntitlementMetered is populated
            case components.BillingEntitlementTypeStatic:
                // res.BillingEntitlement.BillingEntitlementStatic is populated
            case components.BillingEntitlementTypeBoolean:
                // res.BillingEntitlement.BillingEntitlementBoolean is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.OverrideCustomerEntitlementRequest](../../models/operations/overridecustomerentitlementrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.OverrideCustomerEntitlementResponse](../../models/operations/overridecustomerentitlementresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ResetCustomerEntitlementUsage

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Reset the usage of a metered entitlement. The reset starts a new usage period:
usage is zeroed and grants roll over according to their rollover settings.

Usage is reset automatically at the end of each usage period. Use this operation
to reset it earlier, for example to align the entitlement with the customer's
billing period. The usage period anchor can be moved at the same time.

### Example Usage

<!-- UsageSnippet language="go" operationID="reset-customer-entitlement-usage" method="post" path="/v3/openmeter/customers/{customerId}/entitlements/{entitlementId}/reset" -->
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

    res, err := s.OpenMeterEntitlements.ResetCustomerEntitlementUsage(ctx, operations.ResetCustomerEntitlementUsageRequest{
        CustomerID: "01G65Z755AFWAKHE12NY0CQ9FH",
        EntitlementID: "01G65Z755AFWAKHE12NY0CQ9FH",
        ResetCustomerEntitlementUsageRequest: &components.ResetCustomerEntitlementUsageRequest{
            EffectiveAt: types.MustNewTimeFromString("2023-01-01T01:01:01.001Z"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ResetCustomerEntitlementUsageRequest](../../models/operations/resetcustomerentitlementusagerequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ResetCustomerEntitlementUsageResponse](../../models/operations/resetcustomerentitlementusageresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCustomerEntitlementValue

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Get the customer's access through a single entitlement, optionally evaluated at
a point in time.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-customer-entitlement-value" method="get" path="/v3/openmeter/customers/{customerId}/entitlements/{entitlementId}/value" -->
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

    res, err := s.OpenMeterEntitlements.GetCustomerEntitlementValue(ctx, operations.GetCustomerEntitlementValueRequest{
        CustomerID: "01G65Z755AFWAKHE12NY0CQ9FH",
        EntitlementID: "01G65Z755AFWAKHE12NY0CQ9FH",
        At: types.MustNewTimeFromString("2023-01-01T01:01:01.001Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingEntitlementValueResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetCustomerEntitlementValueRequest](../../models/operations/getcustomerentitlementvaluerequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.GetCustomerEntitlementValueResponse](../../models/operations/getcustomerentitlementvalueresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## QueryEntitlementAccess

Query feature access for a list of customers.

The endpoint resolves each provided identifier to a customer and returns the
access status for the requested features, plus optional credit balance
availability.

_Designed to be called on a fixed refresh interval and the query response is
intended to be cached._

### Example Usage

<!-- UsageSnippet language="go" operationID="query-entitlement-access" method="post" path="/v3/openmeter/entitlement-access/query" -->
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

    res, err := s.OpenMeterEntitlements.QueryEntitlementAccess(ctx, components.EntitlementAccessQueryRequest{
        Customer: components.EntitlementAccessQueryRequestCustomer{
            Keys: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.EntitlementAccessQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `entitlementAccessQueryRequest`                                                                      | [components.EntitlementAccessQueryRequest](../../models/components/entitlementaccessqueryrequest.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `page`                                                                                               | [*components.CursorPaginationQueryPage](../../models/components/cursorpaginationquerypage.md)        | :heavy_minus_sign:                                                                                   | Determines which page of the collection to retrieve.                                                 |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.QueryEntitlementAccessResponse](../../models/operations/queryentitlementaccessresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListEntitlements

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

List the active entitlements of all customers. Intended for administrative use.
To list the entitlements of a single customer, use the customer entitlements
endpoints; to check entitlement access, use the entitlement access endpoints.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-entitlements" method="get" path="/v3/openmeter/entitlements" -->
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

    res, err := s.OpenMeterEntitlements.ListEntitlements(ctx, operations.ListEntitlementsRequest{
        Sort: sdkkonnectgo.Pointer("created_at desc"),
        Filter: &components.ListEntitlementsParamsFilter{
            FeatureID: sdkkonnectgo.Pointer(components.CreateListEntitlementsParamsFilterULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
            CustomerID: sdkkonnectgo.Pointer(components.CreateListEntitlementsParamsFilterCustomerIDULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntitlementPagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListEntitlementsRequest](../../models/operations/listentitlementsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListEntitlementsResponse](../../models/operations/listentitlementsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetEntitlement

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Get an entitlement by ID. To check entitlement access, use the entitlement
access endpoints instead.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-entitlement" method="get" path="/v3/openmeter/entitlements/{entitlementId}" -->
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

    res, err := s.OpenMeterEntitlements.GetEntitlement(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingEntitlement != nil {
        switch res.BillingEntitlement.Type {
            case components.BillingEntitlementTypeMetered:
                // res.BillingEntitlement.BillingEntitlementMetered is populated
            case components.BillingEntitlementTypeStatic:
                // res.BillingEntitlement.BillingEntitlementStatic is populated
            case components.BillingEntitlementTypeBoolean:
                // res.BillingEntitlement.BillingEntitlementBoolean is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `entitlementID`                                          | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetEntitlementResponse](../../models/operations/getentitlementresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListGrants

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

List the grants of all customers and entitlements. To list the grants of a
single entitlement, use the customer entitlement grants endpoint.

Deleted grants are excluded unless `include_deleted` is set. Voided and expired
grants are always included, as they are part of the balance history.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-grants" method="get" path="/v3/openmeter/grants" -->
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

    res, err := s.OpenMeterEntitlements.ListGrants(ctx, operations.ListGrantsRequest{
        Sort: sdkkonnectgo.Pointer("created_at desc"),
        Filter: &components.ListGrantsParamsFilter{
            CustomerID: sdkkonnectgo.Pointer(components.CreateListGrantsParamsFilterULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
            FeatureID: sdkkonnectgo.Pointer(components.CreateListGrantsParamsFilterFeatureIDULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntitlementGrantPagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListGrantsRequest](../../models/operations/listgrantsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.ListGrantsResponse](../../models/operations/listgrantsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## VoidGrant

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Void a grant so it no longer adds to the balance. Usage already deducted from
the grant is kept.

### Example Usage

<!-- UsageSnippet language="go" operationID="void-grant" method="delete" path="/v3/openmeter/grants/{grantId}" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.OpenMeterEntitlements.VoidGrant(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", types.MustNewTimeFromString("2023-01-01T01:01:01.001Z"))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                       | Type                                                                                                                            | Required                                                                                                                        | Description                                                                                                                     | Example                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                           | :heavy_check_mark:                                                                                                              | The context to use for the request.                                                                                             |                                                                                                                                 |
| `grantID`                                                                                                                       | `string`                                                                                                                        | :heavy_check_mark:                                                                                                              | N/A                                                                                                                             | 01G65Z755AFWAKHE12NY0CQ9FH                                                                                                      |
| `voidedAt`                                                                                                                      | [*time.Time](https://pkg.go.dev/time#Time)                                                                                      | :heavy_minus_sign:                                                                                                              | The time the grant is voided from. Defaults to now; it cannot be in the future<br/>or before the start of the current usage period. | 2023-01-01T01:01:01.001Z                                                                                                        |
| `opts`                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                        | :heavy_minus_sign:                                                                                                              | The options for this request.                                                                                                   |                                                                                                                                 |

### Response

**[*operations.VoidGrantResponse](../../models/operations/voidgrantresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |