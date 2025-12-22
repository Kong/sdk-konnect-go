# Scorecards

## Overview

A scorecard helps you evaluate services based on its criteria.
Scorecards help you detect issues, like whether there are services in the catalog that don't have an on-call engineer assigned, or if you have GitHub repositories with stale pull requests that aren't getting reviewed or closed.
From the scorecard, you can view details on either a per-service or per-criteria basis.
Learn more about scorecards by visiting our [documentation](https://developer.konghq.com/service-catalog/scorecards/).


### Available Operations

* [ListCatalogServiceScorecards](#listcatalogservicescorecards) - List Catalog Service Scorecards
* [FetchCatalogServiceScorecard](#fetchcatalogservicescorecard) - Get a Catalog Service Scorecard
* [ListScorecardTemplates](#listscorecardtemplates) - List Scorecard Templates
* [CreateScorecard](#createscorecard) - Create Scorecard
* [ListScorecards](#listscorecards) - List Scorecards
* [FetchScorecard](#fetchscorecard) - Get a Scorecard
* [UpdateScorecard](#updatescorecard) - Update Scorecard
* [DeleteScorecard](#deletescorecard) - Delete Scorecard
* [ListScorecardServices](#listscorecardservices) - List Scorecard Services
* [ListScorecardCriteria](#listscorecardcriteria) - List Scorecard Criteria
* [ListScorecardCriteriaServices](#listscorecardcriteriaservices) - List Scorecard Criteria Services

## ListCatalogServiceScorecards

Returns a paginated collection of scorecards targeting the given service.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-catalog-service-scorecards" method="get" path="/v1/catalog-services/{id}/scorecards" -->
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

    res, err := s.Scorecards.ListCatalogServiceScorecards(ctx, operations.ListCatalogServiceScorecardsRequest{
        ID: "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.ScorecardFilterParameters{
            ScoreRawValue: sdkkonnectgo.Pointer(components.CreateScoreRawValueNumericFieldFilter(
                components.CreateNumericFieldFilterNumber(
                    21,
                ),
            )),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCatalogServiceScorecardsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListCatalogServiceScorecardsRequest](../../models/operations/listcatalogservicescorecardsrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ListCatalogServiceScorecardsResponse](../../models/operations/listcatalogservicescorecardsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchCatalogServiceScorecard

Fetches the given scorecard targeting the service.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-catalog-service-scorecard" method="get" path="/v1/catalog-services/{serviceId}/scorecards/{scorecardId}" -->
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

    res, err := s.Scorecards.FetchCatalogServiceScorecard(ctx, "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7", "f3704e4c-104d-4f21-998a-20d4364c893f")
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogServiceScorecard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `serviceID`                                              | *string*                                                 | :heavy_check_mark:                                       | The `id` of the service.                                 | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `scorecardID`                                            | *string*                                                 | :heavy_check_mark:                                       | The `id` of the scorecard.                               | f3704e4c-104d-4f21-998a-20d4364c893f                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchCatalogServiceScorecardResponse](../../models/operations/fetchcatalogservicescorecardresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListScorecardTemplates

Returns a paginated collection of scorecard templates.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-scorecard-templates" method="get" path="/v1/scorecard-templates" -->
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

    res, err := s.Scorecards.ListScorecardTemplates(ctx, operations.ListScorecardTemplatesRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListScorecardTemplatesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListScorecardTemplatesRequest](../../models/operations/listscorecardtemplatesrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListScorecardTemplatesResponse](../../models/operations/listscorecardtemplatesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateScorecard

Creates a scorecard.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-scorecard" method="post" path="/v1/scorecards" -->
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

    res, err := s.Scorecards.CreateScorecard(ctx, components.CreateScorecard{
        Name: "Incident Response",
        Description: sdkkonnectgo.Pointer("Governs key metrics pertaining to teams' incident response practices."),
        EntitySelector: sdkkonnectgo.Pointer(components.CreateScorecardEntitySelectorServiceSelector(
            components.CreateServiceSelectorByLabelSelector(
                components.ByLabelSelector{
                    Selector: components.ByLabelSelectorServiceSelectorSelectorLabel,
                    SelectorParameters: components.LabelSelectorParams{
                        Operator: components.StringSelectorOperatorEq,
                        Value: "cloud_platform",
                        LabelKey: "product_area",
                    },
                },
            ),
        )),
        Criteria: []components.CreateScorecardCriteria{
            components.CreateScorecardCriteria{
                Name: sdkkonnectgo.Pointer("Median time to merge opened PRs is less than 6 hours over the last 3 months."),
                Enabled: true,
                TemplateName: "time_to_merge",
                TemplateParameters: map[string]any{
                    "measure": "median",
                    "threshold": map[string]any{
                        "unit": "hours",
                        "value": 6,
                    },
                    "window": map[string]any{
                        "unit": "months",
                        "value": 3,
                    },
                },
                SectionName: sdkkonnectgo.Pointer("Source Code Management"),
            },
        },
        ScorecardTemplate: components.CreateScorecardScorecardTemplateKongBestPractices.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ScorecardWithCriteria != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [components.CreateScorecard](../../models/components/createscorecard.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.CreateScorecardResponse](../../models/operations/createscorecardresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListScorecards

Returns a paginated collection of scorecards.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-scorecards" method="get" path="/v1/scorecards" -->
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

    res, err := s.Scorecards.ListScorecards(ctx, operations.ListScorecardsRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.ScorecardFilterParameters{
            ScoreRawValue: sdkkonnectgo.Pointer(components.CreateScoreRawValueNumericFieldFilter(
                components.CreateNumericFieldFilterNumber(
                    21,
                ),
            )),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListScorecardsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListScorecardsRequest](../../models/operations/listscorecardsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListScorecardsResponse](../../models/operations/listscorecardsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchScorecard

Fetches a scorecard.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-scorecard" method="get" path="/v1/scorecards/{id}" -->
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

    res, err := s.Scorecards.FetchScorecard(ctx, "f3704e4c-104d-4f21-998a-20d4364c893f")
    if err != nil {
        log.Fatal(err)
    }
    if res.ScorecardWithCriteria != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The `id` of the scorecard.                               | f3704e4c-104d-4f21-998a-20d4364c893f                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchScorecardResponse](../../models/operations/fetchscorecardresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateScorecard

Updates a scorecard.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-scorecard" method="put" path="/v1/scorecards/{id}" -->
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

    res, err := s.Scorecards.UpdateScorecard(ctx, "f3704e4c-104d-4f21-998a-20d4364c893f", components.UpdateScorecard{
        Name: "Incident Response",
        Description: sdkkonnectgo.Pointer("Governs key metrics pertaining to teams' incident response practices."),
        EntitySelector: sdkkonnectgo.Pointer(components.CreateScorecardEntitySelectorServiceSelector(
            components.CreateServiceSelectorByLabelSelector(
                components.ByLabelSelector{
                    Selector: components.ByLabelSelectorServiceSelectorSelectorLabel,
                    SelectorParameters: components.LabelSelectorParams{
                        Operator: components.StringSelectorOperatorEq,
                        Value: "cloud_platform",
                        LabelKey: "product_area",
                    },
                },
            ),
        )),
        Criteria: []components.Criteria{},
        ScorecardTemplate: components.UpdateScorecardScorecardTemplateKongBestPractices.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ScorecardWithCriteria != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |                                                                          |
| `id`                                                                     | *string*                                                                 | :heavy_check_mark:                                                       | The `id` of the scorecard.                                               | f3704e4c-104d-4f21-998a-20d4364c893f                                     |
| `updateScorecard`                                                        | [components.UpdateScorecard](../../models/components/updatescorecard.md) | :heavy_check_mark:                                                       | N/A                                                                      |                                                                          |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |                                                                          |

### Response

**[*operations.UpdateScorecardResponse](../../models/operations/updatescorecardresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteScorecard

Deletes a scorecard.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-scorecard" method="delete" path="/v1/scorecards/{id}" -->
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

    res, err := s.Scorecards.DeleteScorecard(ctx, "f3704e4c-104d-4f21-998a-20d4364c893f")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The `id` of the scorecard.                               | f3704e4c-104d-4f21-998a-20d4364c893f                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteScorecardResponse](../../models/operations/deletescorecardresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListScorecardServices

Lists services targeted by a scorecard.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-scorecard-services" method="get" path="/v1/scorecards/{id}/catalog-services" -->
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

    res, err := s.Scorecards.ListScorecardServices(ctx, operations.ListScorecardServicesRequest{
        ID: "f3704e4c-104d-4f21-998a-20d4364c893f",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.ScorecardServiceFilterParameters{
            CustomFields: sdkkonnectgo.Pointer(components.CreateScorecardServiceFilterParametersCustomFieldsNumericFieldFilter(
                components.CreateNumericFieldFilterNumber(
                    21,
                ),
            )),
            CreatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
                components.DateTimeFieldLTFilter{
                    Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            UpdatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
                components.DateTimeFieldLTFilter{
                    Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            ScoreRawValue: sdkkonnectgo.Pointer(components.CreateScorecardServiceFilterParametersScoreRawValueNumericFieldFilter(
                components.CreateNumericFieldFilterNumber(
                    21,
                ),
            )),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListScorecardServicesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListScorecardServicesRequest](../../models/operations/listscorecardservicesrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListScorecardServicesResponse](../../models/operations/listscorecardservicesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListScorecardCriteria

Lists criteria, including passing service counts, belonging to a scorecard.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-scorecard-criteria" method="get" path="/v1/scorecards/{id}/criteria" -->
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

    res, err := s.Scorecards.ListScorecardCriteria(ctx, operations.ListScorecardCriteriaRequest{
        ID: "f3704e4c-104d-4f21-998a-20d4364c893f",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.ScorecardCriteriaFilterParameters{
            EvaluationPassingServicesCount: sdkkonnectgo.Pointer(components.CreateEvaluationPassingServicesCountNumericFieldFilter(
                components.CreateNumericFieldFilterNumber(
                    21,
                ),
            )),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListScorecardCriteriaResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListScorecardCriteriaRequest](../../models/operations/listscorecardcriteriarequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListScorecardCriteriaResponse](../../models/operations/listscorecardcriteriaresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListScorecardCriteriaServices

Lists services targeted by a scorecard criteria with evaluation results per-service.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-scorecard-criteria-services" method="get" path="/v1/scorecards/{scorecardId}/criteria/{criteriaId}/catalog-services" -->
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

    res, err := s.Scorecards.ListScorecardCriteriaServices(ctx, operations.ListScorecardCriteriaServicesRequest{
        ScorecardID: "f3704e4c-104d-4f21-998a-20d4364c893f",
        CriteriaID: "5c1121f9-3f3a-47c7-9bb6-c81a51128714",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.ScorecardCriteriaServiceFilterParameters{
            CustomFields: sdkkonnectgo.Pointer(components.CreateScorecardCriteriaServiceFilterParametersCustomFieldsStringFieldFilter(
                components.StringFieldFilter{},
            )),
            CreatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldEqualsFilter(
                components.DateTimeFieldEqualsFilter{
                    Eq: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            UpdatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTEFilter(
                components.DateTimeFieldLTEFilter{
                    Lte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            EvaluationSuccessfullyEvaluatedAt: sdkkonnectgo.Pointer(components.CreateEvaluationSuccessfullyEvaluatedAtDateTimeFieldFilter(
                components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
                    components.DateTimeFieldLTFilter{
                        Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                    },
                ),
            )),
            EvaluationAttemptedAt: sdkkonnectgo.Pointer(components.CreateEvaluationAttemptedAtDateTimeFieldFilter(
                components.CreateDateTimeFieldFilterDateTimeFieldLTEFilter(
                    components.DateTimeFieldLTEFilter{
                        Lte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                    },
                ),
            )),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListScorecardCriteriaServicesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ListScorecardCriteriaServicesRequest](../../models/operations/listscorecardcriteriaservicesrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ListScorecardCriteriaServicesResponse](../../models/operations/listscorecardcriteriaservicesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |