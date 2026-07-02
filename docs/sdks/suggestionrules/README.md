# SuggestionRules

## Overview

Suggestion rules allow for automated management of resources.
Suggestion rule can be configured to:
- Map resources to existing catalog service
- Map resources to a new catalog service
- Archive resources

System suggestion rules are built-in rules which are configured and managed by Kong.


### Available Operations

* [ListSystemIntegrationSuggestionRule](#listsystemintegrationsuggestionrule) - List System Suggestion Rules
* [CreateIntegrationSuggestionRule](#createintegrationsuggestionrule) - Create Suggestion Rule
* [ListIntegrationSuggestionRule](#listintegrationsuggestionrule) - List Suggestion Rules
* [GetIntegrationSuggestionRule](#getintegrationsuggestionrule) - Get a Suggestion Rule
* [UpdateIntegrationSuggestionRule](#updateintegrationsuggestionrule) - Update Suggestion Rule
* [DeleteIntegrationSuggestionRule](#deleteintegrationsuggestionrule) - Delete Suggestion Rule
* [TestSuggestionRule](#testsuggestionrule) - Test a Suggestion Rule Configuration

## ListSystemIntegrationSuggestionRule

Returns a paginated collection of a system defined suggestion rules.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-system-integration-suggestion-rule" method="get" path="/v1/integrations/{integration}/system-suggestion-rules" -->
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

    res, err := s.SuggestionRules.ListSystemIntegrationSuggestionRule(ctx, operations.ListSystemIntegrationSuggestionRuleRequest{
        Integration: "gateway-manager",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListSystemSuggestionRuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.ListSystemIntegrationSuggestionRuleRequest](../../models/operations/listsystemintegrationsuggestionrulerequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.ListSystemIntegrationSuggestionRuleResponse](../../models/operations/listsystemintegrationsuggestionruleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateIntegrationSuggestionRule

Creates a suggestion rule for the given integration.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-integration-suggestion-rule" method="post" path="/v1/integrations/{integration}/suggestion-rules" -->
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

    res, err := s.SuggestionRules.CreateIntegrationSuggestionRule(ctx, "gateway-manager", components.SuggestionRulePayload{
        Name: "<value>",
        Description: sdkkonnectgo.Pointer("Suggestion rule description"),
        Binding: "pagerduty_service",
        Active: false,
        AutoAccept: false,
        DiscoverySelector: ".gateway_service.tags | contains([\"_KonnectDeployment:\"])",
        Action: components.CreateSuggestionRulePayloadActionMapActionPayload(
            components.MapActionPayload{
                Type: components.MapActionPayloadTypeMap,
                Data: components.MapActionPayloadData{
                    Service: components.CreateMapActionPayloadServiceMapByName(
                        components.MapByName{
                            Name: "{{ .gateway_service.name }}-{{ .runtime_group.labels.env }}",
                        },
                    ),
                    Resource: components.MapActionPayloadResource{
                        Labels: map[string]string{
                            "env": "test",
                        },
                    },
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SuggestionRule != nil {
        switch res.SuggestionRule.Action.Type {
            case components.SuggestedRuleActionTypeArchiveActionPayload:
                // res.SuggestionRule.Action.ArchiveActionPayload is populated
            case components.SuggestedRuleActionTypeIgnoreActionPayload:
                // res.SuggestionRule.Action.IgnoreActionPayload is populated
            case components.SuggestedRuleActionTypeMapActionPayload:
                // res.SuggestionRule.Action.MapActionPayload is populated
            case components.SuggestedRuleActionTypeCreateOrMapAction:
                // res.SuggestionRule.Action.CreateOrMapAction is populated
        }

    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          | Example                                                                              |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |                                                                                      |
| `integration`                                                                        | `string`                                                                             | :heavy_check_mark:                                                                   | Machine name of the integration.                                                     | gateway-manager                                                                      |
| `suggestionRulePayload`                                                              | [components.SuggestionRulePayload](../../models/components/suggestionrulepayload.md) | :heavy_check_mark:                                                                   | Request body schema for creating a suggestion rule for an integration.               |                                                                                      |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |                                                                                      |

### Response

**[*operations.CreateIntegrationSuggestionRuleResponse](../../models/operations/createintegrationsuggestionruleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListIntegrationSuggestionRule

Returns a paginated collection of a suggestion rules.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-integration-suggestion-rule" method="get" path="/v1/integrations/{integration}/suggestion-rules" -->
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

    res, err := s.SuggestionRules.ListIntegrationSuggestionRule(ctx, operations.ListIntegrationSuggestionRuleRequest{
        Integration: "gateway-manager",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.SuggestionRuleFilterParameters{
            CreatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTEFilter(
                components.DateTimeFieldLTEFilter{
                    Lte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            UpdatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTEFilter(
                components.DateTimeFieldLTEFilter{
                    Lte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListSuggestionRuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ListIntegrationSuggestionRuleRequest](../../models/operations/listintegrationsuggestionrulerequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ListIntegrationSuggestionRuleResponse](../../models/operations/listintegrationsuggestionruleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetIntegrationSuggestionRule

Returns information about a suggestion rule from a given suggestion rule ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-integration-suggestion-rule" method="get" path="/v1/integrations/{integration}/suggestion-rules/{suggestionRuleId}" -->
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

    res, err := s.SuggestionRules.GetIntegrationSuggestionRule(ctx, "gateway-manager", "22f72bd1-2897-473e-9471-302de3ccf38b")
    if err != nil {
        log.Fatal(err)
    }
    if res.SuggestionRule != nil {
        switch res.SuggestionRule.Action.Type {
            case components.SuggestedRuleActionTypeArchiveActionPayload:
                // res.SuggestionRule.Action.ArchiveActionPayload is populated
            case components.SuggestedRuleActionTypeIgnoreActionPayload:
                // res.SuggestionRule.Action.IgnoreActionPayload is populated
            case components.SuggestedRuleActionTypeMapActionPayload:
                // res.SuggestionRule.Action.MapActionPayload is populated
            case components.SuggestedRuleActionTypeCreateOrMapAction:
                // res.SuggestionRule.Action.CreateOrMapAction is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `integration`                                            | `string`                                                 | :heavy_check_mark:                                       | Machine name of the integration.                         | gateway-manager                                          |
| `suggestionRuleID`                                       | `string`                                                 | :heavy_check_mark:                                       | ID of the suggestion rule.                               | 22f72bd1-2897-473e-9471-302de3ccf38b                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetIntegrationSuggestionRuleResponse](../../models/operations/getintegrationsuggestionruleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateIntegrationSuggestionRule

Updates the given suggestion rule.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-integration-suggestion-rule" method="put" path="/v1/integrations/{integration}/suggestion-rules/{suggestionRuleId}" -->
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

    res, err := s.SuggestionRules.UpdateIntegrationSuggestionRule(ctx, operations.UpdateIntegrationSuggestionRuleRequest{
        Integration: "gateway-manager",
        SuggestionRuleID: "22f72bd1-2897-473e-9471-302de3ccf38b",
        SuggestionRulePayload: components.SuggestionRulePayload{
            Name: "<value>",
            Description: sdkkonnectgo.Pointer("Suggestion rule description"),
            Binding: "pagerduty_service",
            Active: true,
            AutoAccept: false,
            DiscoverySelector: ".gateway_service.tags | contains([\"_KonnectDeployment:\"])",
            Action: components.CreateSuggestionRulePayloadActionCreateOrMapAction(
                components.CreateOrMapAction{
                    Type: components.CreateOrMapActionTypeCreateOrMap,
                    Data: components.CreateOrMapActionData{
                        Service: components.CreateOrMapActionService{
                            Name: "<value>",
                            DisplayName: sdkkonnectgo.Pointer("Jeffry74"),
                            Description: "oh whose secularize deny fuzzy after blah fooey saw pro",
                            Labels: map[string]string{
                                "env": "test",
                            },
                            Metadata: map[string]*components.MetadataTemplate{

                            },
                        },
                        Resource: components.CreateOrMapActionResource{
                            Labels: map[string]string{
                                "env": "test",
                            },
                        },
                    },
                },
            ),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SuggestionRule != nil {
        switch res.SuggestionRule.Action.Type {
            case components.SuggestedRuleActionTypeArchiveActionPayload:
                // res.SuggestionRule.Action.ArchiveActionPayload is populated
            case components.SuggestedRuleActionTypeIgnoreActionPayload:
                // res.SuggestionRule.Action.IgnoreActionPayload is populated
            case components.SuggestedRuleActionTypeMapActionPayload:
                // res.SuggestionRule.Action.MapActionPayload is populated
            case components.SuggestedRuleActionTypeCreateOrMapAction:
                // res.SuggestionRule.Action.CreateOrMapAction is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.UpdateIntegrationSuggestionRuleRequest](../../models/operations/updateintegrationsuggestionrulerequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.UpdateIntegrationSuggestionRuleResponse](../../models/operations/updateintegrationsuggestionruleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteIntegrationSuggestionRule

Deletes the given suggestion rule.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-integration-suggestion-rule" method="delete" path="/v1/integrations/{integration}/suggestion-rules/{suggestionRuleId}" -->
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

    res, err := s.SuggestionRules.DeleteIntegrationSuggestionRule(ctx, "gateway-manager", "22f72bd1-2897-473e-9471-302de3ccf38b")
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
| `integration`                                            | `string`                                                 | :heavy_check_mark:                                       | Machine name of the integration.                         | gateway-manager                                          |
| `suggestionRuleID`                                       | `string`                                                 | :heavy_check_mark:                                       | ID of the suggestion rule.                               | 22f72bd1-2897-473e-9471-302de3ccf38b                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteIntegrationSuggestionRuleResponse](../../models/operations/deleteintegrationsuggestionruleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## TestSuggestionRule

Test suggestion rule configuration against an integration record.

### Example Usage

<!-- UsageSnippet language="go" operationID="test-suggestion-rule" method="post" path="/v1/integrations/{integration}/suggestion-rules/test" -->
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

    res, err := s.SuggestionRules.TestSuggestionRule(ctx, "gateway-manager", components.TestSuggestionRulePayload{
        Binding: "pagerduty_service",
        DiscoverySelector: ".gateway_service.tags | contains([\"_KonnectDeployment:\"])",
        Action: components.CreateSuggestedRuleActionArchiveActionPayload(
            components.ArchiveActionPayload{
                Type: components.ArchiveActionPayloadTypeArchive,
            },
        ),
        IntegrationRecord: map[string]any{
            "control_plane": map[string]any{
                "id": "0a7647d1-483a-4817-a5ac-28937b511563",
                "name": "control_plane_name",
            },
            "gateway_service": map[string]any{
                "id": "3fd1bff6-b60d-4c22-bb74-e4cac63196fb",
                "name": "gateway_service_name",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TestSuggestionRule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `integration`                                                                                | `string`                                                                                     | :heavy_check_mark:                                                                           | Machine name of the integration.                                                             | gateway-manager                                                                              |
| `testSuggestionRulePayload`                                                                  | [components.TestSuggestionRulePayload](../../models/components/testsuggestionrulepayload.md) | :heavy_check_mark:                                                                           | Request body schema for testing a suggestion rule configuration against an integration data. |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

### Response

**[*operations.TestSuggestionRuleResponse](../../models/operations/testsuggestionruleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |