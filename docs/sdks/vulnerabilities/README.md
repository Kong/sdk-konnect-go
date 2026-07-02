# Vulnerabilities

## Overview

Retrieving, analyzing, and managing software vulnerabilities.


### Available Operations

* [ListVulnerabilityScans](#listvulnerabilityscans) - List Vulnerability Scans
* [CreateVulnerabilityScan](#createvulnerabilityscan) - Create Vulnerability Scan
* [FetchVulnerabilityScan](#fetchvulnerabilityscan) - Get a Vulnerability Scan
* [ListVulnerabilityScanVulnerabilities](#listvulnerabilityscanvulnerabilities) - List Vulnerability Scan Vulnerabilities
* [ListServiceVulnerabilityScans](#listservicevulnerabilityscans) - List Service Vulnerability Scans
* [FetchServiceVulnerabilityScan](#fetchservicevulnerabilityscan) - Get a Service Vulnerability Scan
* [ListVulnerabilities](#listvulnerabilities) - List Vulnerabilities
* [FetchVulnerability](#fetchvulnerability) - Get Vulnerability
* [ListVulnerabilityServices](#listvulnerabilityservices) - List Vulnerability Services
* [ListVulnerabilityInstances](#listvulnerabilityinstances) - List Vulnerability Instances
* [ListServiceVulnerabilities](#listservicevulnerabilities) - List Service Vulnerabilities
* [FetchServiceVulnerability](#fetchservicevulnerability) - Get Service Vulnerability
* [ListServiceVulnerabilityInstances](#listservicevulnerabilityinstances) - List Service Vulnerability Instances
* [PutServiceVulnerabilitySeverityOverride](#putservicevulnerabilityseverityoverride) - Put Service Vulnerability Severity Override
* [DeleteServiceVulnerabilitySeverityOverride](#deleteservicevulnerabilityseverityoverride) - Delete Vulnerability Severity Override
* [PutServiceVulnerabilityDismissal](#putservicevulnerabilitydismissal) - Put Service Vulnerability Dismissal
* [DeleteServiceVulnerabilityDismissal](#deleteservicevulnerabilitydismissal) - Delete Service Vulnerability Dismissal
* [ListCatalogVulnerabilityServices](#listcatalogvulnerabilityservices) - List Catalog Vulnerability-Services
* [QueryVulnerabilitiesMetrics](#queryvulnerabilitiesmetrics) - Query Vulnerabilities Metrics

## ListVulnerabilityScans

Returns a paginated collection of vulnerability scans.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-vulnerability-scans" method="get" path="/v1/vulnerability-scans" -->
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

    res, err := s.Vulnerabilities.ListVulnerabilityScans(ctx, operations.ListVulnerabilityScansRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.VulnerabilityScanFilterParameters{
            Ts: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldGTEFilter(
                components.DateTimeFieldGTEFilter{
                    Gte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListVulnerabilityScansResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListVulnerabilityScansRequest](../../models/operations/listvulnerabilityscansrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListVulnerabilityScansResponse](../../models/operations/listvulnerabilityscansresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateVulnerabilityScan

Create a vulnerability scan, to be broken down into a set of vulnerabilities.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-vulnerability-scan" method="post" path="/v1/vulnerability-scans" -->
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

    res, err := s.Vulnerabilities.CreateVulnerabilityScan(ctx, components.CreateVulnerabilityScan{
        ScanKey: "namespaces/foo:pods/bar-api",
        Tool: components.VulnerabilityScanToolTrivy,
        Description: sdkkonnectgo.Pointer("futon unto by as reluctantly hygienic carefully for"),
        RawScanReport: map[string]any{
            "key": "<value>",
            "key1": "<value>",
            "key2": "<value>",
        },
        CatalogReference: components.CreateCreateVulnerabilityScanCatalogReferenceCreateServiceVulnerabilityScanCatalogReference(
            components.CreateServiceVulnerabilityScanCatalogReference{
                Service: "user-svc",
            },
        ),
        Environment: sdkkonnectgo.Pointer("prod"),
        Region: sdkkonnectgo.Pointer("us-east-2"),
        SourceCorrelationKey: sdkkonnectgo.Pointer("kong/repo-id"),
        Attributes: map[string]string{
            "org.kong.enterprise.source.version": "v1",
        },
        Ts: types.MustNewTimeFromString("2025-01-01T00:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VulnerabilityScan != nil {
        switch res.VulnerabilityScan.CatalogReference.Type {
            case components.VulnerabilityScanCatalogReferenceTypeServiceVulnerabilityScanCatalogReference:
                // res.VulnerabilityScan.CatalogReference.ServiceVulnerabilityScanCatalogReference is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.CreateVulnerabilityScan](../../models/components/createvulnerabilityscan.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateVulnerabilityScanResponse](../../models/operations/createvulnerabilityscanresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchVulnerabilityScan

Returns vulnerability scan by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-vulnerability-scan" method="get" path="/v1/vulnerability-scans/{scanId}" -->
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

    res, err := s.Vulnerabilities.FetchVulnerabilityScan(ctx, "5aea1e7d-8d51-465e-bb51-5d86de960833")
    if err != nil {
        log.Fatal(err)
    }
    if res.VulnerabilityScan != nil {
        switch res.VulnerabilityScan.CatalogReference.Type {
            case components.VulnerabilityScanCatalogReferenceTypeServiceVulnerabilityScanCatalogReference:
                // res.VulnerabilityScan.CatalogReference.ServiceVulnerabilityScanCatalogReference is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `scanID`                                                 | `string`                                                 | :heavy_check_mark:                                       | ID of the vulnerability scan.                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.FetchVulnerabilityScanResponse](../../models/operations/fetchvulnerabilityscanresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListVulnerabilityScanVulnerabilities

Returns a paginated collection of vulnerabilities for a given vulnerability scan.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-vulnerability-scan-vulnerabilities" method="get" path="/v1/vulnerability-scans/{scanId}/vulnerabilities" -->
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

    res, err := s.Vulnerabilities.ListVulnerabilityScanVulnerabilities(ctx, operations.ListVulnerabilityScanVulnerabilitiesRequest{
        ScanID: "e6fc4958-1158-46e8-9ad9-12076d89be71",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.VulnerabilityScanVulnerabilityFilterParameters{
            OpenedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
                components.DateTimeFieldLTFilter{
                    Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            ReopenedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldGTEFilter(
                components.DateTimeFieldGTEFilter{
                    Gte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            LastOpenedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldGTFilter(
                components.DateTimeFieldGTFilter{
                    Gt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            DismissedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTEFilter(
                components.DateTimeFieldLTEFilter{
                    Lte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            FixedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
                components.DateTimeFieldLTFilter{
                    Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            InstanceCount: sdkkonnectgo.Pointer(components.CreateNumericFieldFilterNumber(
                21.0,
            )),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListVulnerabilityScanVulnerabilitiesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.ListVulnerabilityScanVulnerabilitiesRequest](../../models/operations/listvulnerabilityscanvulnerabilitiesrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.ListVulnerabilityScanVulnerabilitiesResponse](../../models/operations/listvulnerabilityscanvulnerabilitiesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListServiceVulnerabilityScans

Returns a paginated collection of vulnerability scans for a given service.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-service-vulnerability-scans" method="get" path="/v1/catalog-services/{serviceId}/vulnerability-scans" -->
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

    res, err := s.Vulnerabilities.ListServiceVulnerabilityScans(ctx, operations.ListServiceVulnerabilityScansRequest{
        ServiceID: "<id>",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.ServiceVulnerabilityScanFilterParameters{
            Ts: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldEqualsFilter(
                components.DateTimeFieldEqualsFilter{
                    Eq: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListVulnerabilityScansResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ListServiceVulnerabilityScansRequest](../../models/operations/listservicevulnerabilityscansrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ListServiceVulnerabilityScansResponse](../../models/operations/listservicevulnerabilityscansresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchServiceVulnerabilityScan

Returns vulnerability scan, for a given service, by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-service-vulnerability-scan" method="get" path="/v1/catalog-services/{serviceId}/vulnerability-scans/{scanId}" -->
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

    res, err := s.Vulnerabilities.FetchServiceVulnerabilityScan(ctx, "<id>", "a0d7c071-1a70-45e7-92cd-5dcca87639af")
    if err != nil {
        log.Fatal(err)
    }
    if res.VulnerabilityScan != nil {
        switch res.VulnerabilityScan.CatalogReference.Type {
            case components.VulnerabilityScanCatalogReferenceTypeServiceVulnerabilityScanCatalogReference:
                // res.VulnerabilityScan.CatalogReference.ServiceVulnerabilityScanCatalogReference is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serviceID`                                              | `string`                                                 | :heavy_check_mark:                                       | ID of the service.                                       |
| `scanID`                                                 | `string`                                                 | :heavy_check_mark:                                       | ID of the vulnerability scan.                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.FetchServiceVulnerabilityScanResponse](../../models/operations/fetchservicevulnerabilityscanresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListVulnerabilities

Returns a paginated collection of vulnerabilities.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-vulnerabilities" method="get" path="/v1/vulnerabilities" -->
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

    res, err := s.Vulnerabilities.ListVulnerabilities(ctx, operations.ListVulnerabilitiesRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.CatalogVulnerabilityFilterParameters{
            OpenedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldEqualsFilter(
                components.DateTimeFieldEqualsFilter{
                    Eq: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            ReopenedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldGTEFilter(
                components.DateTimeFieldGTEFilter{
                    Gte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            LastOpenedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
                components.DateTimeFieldLTFilter{
                    Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            DismissedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldGTEFilter(
                components.DateTimeFieldGTEFilter{
                    Gte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            FixedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldEqualsFilter(
                components.DateTimeFieldEqualsFilter{
                    Eq: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            InstanceCount: sdkkonnectgo.Pointer(components.CreateNumericFieldFilterNumber(
                21.0,
            )),
            ServiceCount: sdkkonnectgo.Pointer(components.CreateNumericFieldFilterNumber(
                21.0,
            )),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCatalogVulnerabilitiesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListVulnerabilitiesRequest](../../models/operations/listvulnerabilitiesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListVulnerabilitiesResponse](../../models/operations/listvulnerabilitiesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchVulnerability

Returns vulnerability by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-vulnerability" method="get" path="/v1/vulnerabilities/{vulnerabilityId}" -->
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

    res, err := s.Vulnerabilities.FetchVulnerability(ctx, "d7f3e221-bbec-4853-83d3-3ecbf01dc335")
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogVulnerability != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `vulnerabilityID`                                        | `string`                                                 | :heavy_check_mark:                                       | ID of the vulnerability.                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.FetchVulnerabilityResponse](../../models/operations/fetchvulnerabilityresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListVulnerabilityServices

Returns a paginated collection of services for a given vulnerability.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-vulnerability-services" method="get" path="/v1/vulnerabilities/{vulnerabilityId}/catalog-services" -->
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

    res, err := s.Vulnerabilities.ListVulnerabilityServices(ctx, operations.ListVulnerabilityServicesRequest{
        VulnerabilityID: "e13d0e14-8a41-4872-871a-2e030d6c85fc",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
        Filter: &components.VulnerabilityServiceFilterParameters{
            VulnerabilityOpenedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTEFilter(
                components.DateTimeFieldLTEFilter{
                    Lte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            VulnerabilityReopenedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTEFilter(
                components.DateTimeFieldLTEFilter{
                    Lte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            VulnerabilityLastOpenedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
                components.DateTimeFieldLTFilter{
                    Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            VulnerabilityDismissedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldEqualsFilter(
                components.DateTimeFieldEqualsFilter{
                    Eq: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            VulnerabilityFixedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
                components.DateTimeFieldLTFilter{
                    Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            VulnerabilityInstanceCount: sdkkonnectgo.Pointer(components.CreateNumericFieldFilterNumber(
                21.0,
            )),
            CustomFields: sdkkonnectgo.Pointer(components.CreateVulnerabilityServiceFilterParametersCustomFieldsNumericFieldFilter(
                components.CreateNumericFieldFilterNumber(
                    21.0,
                ),
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListVulnerabilityServicesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListVulnerabilityServicesRequest](../../models/operations/listvulnerabilityservicesrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ListVulnerabilityServicesResponse](../../models/operations/listvulnerabilityservicesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListVulnerabilityInstances

Returns a paginated collection of vulnerability instances for a given vulnerability.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-vulnerability-instances" method="get" path="/v1/vulnerabilities/{vulnerabilityId}/instances" -->
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

    res, err := s.Vulnerabilities.ListVulnerabilityInstances(ctx, operations.ListVulnerabilityInstancesRequest{
        VulnerabilityID: "0f7b6254-949b-4c8f-bfea-cfc9aaa65eb9",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListVulnerabilityInstancesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListVulnerabilityInstancesRequest](../../models/operations/listvulnerabilityinstancesrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListVulnerabilityInstancesResponse](../../models/operations/listvulnerabilityinstancesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListServiceVulnerabilities

Returns a paginated collection of vulnerabilities for a given service.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-service-vulnerabilities" method="get" path="/v1/catalog-services/{serviceId}/vulnerabilities" -->
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

    res, err := s.Vulnerabilities.ListServiceVulnerabilities(ctx, operations.ListServiceVulnerabilitiesRequest{
        ServiceID: "<id>",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.ServiceVulnerabilityFilterParameters{
            OpenedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldEqualsFilter(
                components.DateTimeFieldEqualsFilter{
                    Eq: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            ReopenedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
                components.DateTimeFieldLTFilter{
                    Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            LastOpenedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldGTEFilter(
                components.DateTimeFieldGTEFilter{
                    Gte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            DismissedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldEqualsFilter(
                components.DateTimeFieldEqualsFilter{
                    Eq: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            FixedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTEFilter(
                components.DateTimeFieldLTEFilter{
                    Lte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            InstanceCount: sdkkonnectgo.Pointer(components.CreateNumericFieldFilterNumber(
                21.0,
            )),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListVulnerabilitiesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListServiceVulnerabilitiesRequest](../../models/operations/listservicevulnerabilitiesrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListServiceVulnerabilitiesResponse](../../models/operations/listservicevulnerabilitiesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchServiceVulnerability

Returns vulnerability instance, for a given service, by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-service-vulnerability" method="get" path="/v1/catalog-services/{serviceId}/vulnerabilities/{vulnerabilityId}" -->
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

    res, err := s.Vulnerabilities.FetchServiceVulnerability(ctx, "<id>", "5c4c5260-66d0-4940-b4d3-c8a4ac43807d")
    if err != nil {
        log.Fatal(err)
    }
    if res.Vulnerability != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serviceID`                                              | `string`                                                 | :heavy_check_mark:                                       | ID of the service.                                       |
| `vulnerabilityID`                                        | `string`                                                 | :heavy_check_mark:                                       | ID of the vulnerability.                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.FetchServiceVulnerabilityResponse](../../models/operations/fetchservicevulnerabilityresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListServiceVulnerabilityInstances

Returns a paginated collection of vulnerability instances for a given service.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-service-vulnerability-instances" method="get" path="/v1/catalog-services/{serviceId}/vulnerabilities/{vulnerabilityId}/instances" -->
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

    res, err := s.Vulnerabilities.ListServiceVulnerabilityInstances(ctx, operations.ListServiceVulnerabilityInstancesRequest{
        ServiceID: "<id>",
        VulnerabilityID: "ec2832f3-7527-4be4-8324-c9e4276e7a44",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListVulnerabilityInstancesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.ListServiceVulnerabilityInstancesRequest](../../models/operations/listservicevulnerabilityinstancesrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.ListServiceVulnerabilityInstancesResponse](../../models/operations/listservicevulnerabilityinstancesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PutServiceVulnerabilitySeverityOverride

Overrides a vulnerability severity for a given service.

### Example Usage

<!-- UsageSnippet language="go" operationID="put-service-vulnerability-severity-override" method="put" path="/v1/catalog-services/{serviceId}/vulnerabilities/{vulnerabilityId}/severity" -->
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

    res, err := s.Vulnerabilities.PutServiceVulnerabilitySeverityOverride(ctx, operations.PutServiceVulnerabilitySeverityOverrideRequest{
        ServiceID: "<id>",
        VulnerabilityID: "fe0bb845-4af3-471c-9b0c-f07081d059c8",
        PutVulnerabilitySeverityOverride: components.PutVulnerabilitySeverityOverride{
            Comment: sdkkonnectgo.Pointer("Severity bumped to high due to business impact."),
            Severity: components.VulnerabilitySeverityLow,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Vulnerability != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.PutServiceVulnerabilitySeverityOverrideRequest](../../models/operations/putservicevulnerabilityseverityoverriderequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.PutServiceVulnerabilitySeverityOverrideResponse](../../models/operations/putservicevulnerabilityseverityoverrideresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteServiceVulnerabilitySeverityOverride

Deletes the severity override on a vulnerability for a given service.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-service-vulnerability-severity-override" method="delete" path="/v1/catalog-services/{serviceId}/vulnerabilities/{vulnerabilityId}/severity" -->
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

    res, err := s.Vulnerabilities.DeleteServiceVulnerabilitySeverityOverride(ctx, "<id>", "0cc40a5e-38c1-4d7a-a9d1-5af5a746be3f")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serviceID`                                              | `string`                                                 | :heavy_check_mark:                                       | ID of the service.                                       |
| `vulnerabilityID`                                        | `string`                                                 | :heavy_check_mark:                                       | ID of the vulnerability.                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteServiceVulnerabilitySeverityOverrideResponse](../../models/operations/deleteservicevulnerabilityseverityoverrideresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PutServiceVulnerabilityDismissal

Dismisses a vulnerability for a given service.

### Example Usage

<!-- UsageSnippet language="go" operationID="put-service-vulnerability-dismissal" method="put" path="/v1/catalog-services/{serviceId}/vulnerabilities/{vulnerabilityId}/dismissal" -->
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

    res, err := s.Vulnerabilities.PutServiceVulnerabilityDismissal(ctx, operations.PutServiceVulnerabilityDismissalRequest{
        ServiceID: "<id>",
        VulnerabilityID: "fa638a4b-a2d6-444e-9e5f-08dd47d9caf9",
        PutVulnerabilityDismissal: components.PutVulnerabilityDismissal{
            Comment: sdkkonnectgo.Pointer("Acceptable risk."),
            Reason: sdkkonnectgo.Pointer("acceptable_risk"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Vulnerability != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.PutServiceVulnerabilityDismissalRequest](../../models/operations/putservicevulnerabilitydismissalrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.PutServiceVulnerabilityDismissalResponse](../../models/operations/putservicevulnerabilitydismissalresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteServiceVulnerabilityDismissal

Deletes the dismissal on a vulnerability for a given service.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-service-vulnerability-dismissal" method="delete" path="/v1/catalog-services/{serviceId}/vulnerabilities/{vulnerabilityId}/dismissal" -->
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

    res, err := s.Vulnerabilities.DeleteServiceVulnerabilityDismissal(ctx, "<id>", "00008622-6737-40e4-bd1a-60ec338bfac3")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serviceID`                                              | `string`                                                 | :heavy_check_mark:                                       | ID of the service.                                       |
| `vulnerabilityID`                                        | `string`                                                 | :heavy_check_mark:                                       | ID of the vulnerability.                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteServiceVulnerabilityDismissalResponse](../../models/operations/deleteservicevulnerabilitydismissalresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCatalogVulnerabilityServices

Returns a paginated collection of services with cross-vulnerability context.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-catalog-vulnerability-services" method="get" path="/v1/vulnerability-catalog-services" -->
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

    res, err := s.Vulnerabilities.ListCatalogVulnerabilityServices(ctx, operations.ListCatalogVulnerabilityServicesRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
        Filter: &components.CatalogVulnerabilityServiceFilterParameters{
            VulnerabilityMostRecentVulnerabilityOpenedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldEqualsFilter(
                components.DateTimeFieldEqualsFilter{
                    Eq: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            VulnerabilityInstanceCount: sdkkonnectgo.Pointer(components.CreateNumericFieldFilterNumber(
                21.0,
            )),
            VulnerabilityVulnerabilityCount: sdkkonnectgo.Pointer(components.CreateNumericFieldFilterNumber(
                21.0,
            )),
            CustomFields: sdkkonnectgo.Pointer(components.CreateCatalogVulnerabilityServiceFilterParametersCustomFieldsNumericFieldFilter(
                components.CreateNumericFieldFilterNumber(
                    21.0,
                ),
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCatalogVulnerabilityServicesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.ListCatalogVulnerabilityServicesRequest](../../models/operations/listcatalogvulnerabilityservicesrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.ListCatalogVulnerabilityServicesResponse](../../models/operations/listcatalogvulnerabilityservicesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## QueryVulnerabilitiesMetrics

Post a query to the endpoint to retrieve aggregated vulnerabilities metrics.
Data can be requested to group by any 2 dimensions and can include any number of filter conditions.


### Example Usage

<!-- UsageSnippet language="go" operationID="query-vulnerabilities-metrics" method="post" path="/v1/metrics/vulnerabilities" -->
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

    res, err := s.Vulnerabilities.QueryVulnerabilitiesMetrics(ctx, components.VulnerabilitiesMetricsQuery{
        Metrics: []components.Metrics{},
        Dimensions: []components.Dimensions{
            components.DimensionsService,
            components.DimensionsState,
        },
        Filters: []components.Filters{
            components.CreateFiltersVulnerabilitiesMetricsFilterByScanAttributes(
                components.CreateVulnerabilitiesMetricsFilterByScanAttributesStringValueMetricsFilter(
                    components.StringValueMetricsFilter{
                        Operator: components.MetricsFilterInOperatorIn,
                        Value: []string{
                            "2.0.7",
                        },
                        Field: "scan_attributes.image_version",
                    },
                ),
            ),
        },
        TimeRange: components.CreateVulnerabilitiesMetricsQueryTimeRangeRelative(
            components.VulnerabilityMetricsRelativeTimeRange{
                Tz: sdkkonnectgo.Pointer("America/Chicago"),
                Type: components.VulnerabilityMetricsRelativeTimeRangeTypeRelative,
                TimeRange: "current_month",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VulnerabilitiesMetricsQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [components.VulnerabilitiesMetricsQuery](../../models/components/vulnerabilitiesmetricsquery.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.QueryVulnerabilitiesMetricsResponse](../../models/operations/queryvulnerabilitiesmetricsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |