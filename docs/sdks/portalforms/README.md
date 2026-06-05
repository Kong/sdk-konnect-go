# PortalForms

## Overview

### Available Operations

* [CreatePortalForm](#createportalform) - Create Form
* [ListPortalForms](#listportalforms) - List Forms
* [GetPortalForm](#getportalform) - Get Form
* [ReplacePortalForm](#replaceportalform) - Replace Form
* [DeletePortalForm](#deleteportalform) - Delete Form

## CreatePortalForm

Creates a custom form for a portal. The form's `type` determines its consumer: type `developer_registration` is served by the portal's developer signup flow; `api_registration` is linked to API publications via `form_id`. The `fields` array must contain exactly one `submit` field — required for every form type. For type `developer_registration`, the `name` is reserved as `developer-registration` and the built-in `full_name` and `email` fields are required in the array; the server returns 400 if either is missing.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-portal-form" method="post" path="/v3/portals/{portalId}/forms" -->
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

    res, err := s.PortalForms.CreatePortalForm(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", components.CreateCreatePortalFormRequestAPIRegistration(
        components.CreateAPIRegistrationFormRequest{
            Type: components.CreateAPIRegistrationFormRequestTypeAPIRegistration,
            Name: "flights-api-registration",
            Fields: []components.CustomFormFieldInput{
                components.CreateCustomFormFieldInputSubmit(
                    components.CustomFormSubmitFieldInput{
                        Name: components.CustomFormSubmitFieldInputNameSubmit,
                        Type: components.CustomFormSubmitFieldInputTypeSubmit,
                        Value: "Create account",
                    },
                ),
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomFormResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `portalID`                                                                               | `string`                                                                                 | :heavy_check_mark:                                                                       | ID of the portal.                                                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                     |
| `createPortalFormRequest`                                                                | [components.CreatePortalFormRequest](../../models/components/createportalformrequest.md) | :heavy_check_mark:                                                                       | Create a custom form for a portal.                                                       |                                                                                          |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.CreatePortalFormResponse](../../models/operations/createportalformresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListPortalForms

Lists custom forms for a portal. Default `developer_registration` forms are not returned virtually — only forms persisted in the portal are listed.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-portal-forms" method="get" path="/v3/portals/{portalId}/forms" -->
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

    res, err := s.PortalForms.ListPortalForms(ctx, operations.ListPortalFormsRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPortalFormsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListPortalFormsRequest](../../models/operations/listportalformsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListPortalFormsResponse](../../models/operations/listportalformsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPortalForm

Retrieves a single custom form by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-portal-form" method="get" path="/v3/portals/{portalId}/forms/{formId}" -->
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

    res, err := s.PortalForms.GetPortalForm(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomFormResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `portalID`                                               | `string`                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `formID`                                                 | `string`                                                 | :heavy_check_mark:                                       | ID of the portal form.                                   | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPortalFormResponse](../../models/operations/getportalformresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ReplacePortalForm

Replaces a custom form with the supplied representation. Field removal is achieved by omitting the field from the `fields` array; field renames are achieved by editing `label` (the field `name` slug is immutable). The form `type` cannot be changed; the form `name` is immutable for `developer_registration` and mutable for `api_registration`. For type `developer_registration`, the built-in `full_name` and `email` fields are required in the replacement array; the server returns 400 if either is missing.

### Example Usage

<!-- UsageSnippet language="go" operationID="replace-portal-form" method="put" path="/v3/portals/{portalId}/forms/{formId}" -->
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

    res, err := s.PortalForms.ReplacePortalForm(ctx, operations.ReplacePortalFormRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        FormID: "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7",
        UpdatePortalFormRequest: components.CreateUpdatePortalFormRequestAPIRegistration(
            components.UpdateAPIRegistrationFormRequest{
                Type: components.UpdateAPIRegistrationFormRequestTypeAPIRegistration,
                Fields: []components.CustomFormFieldInput{},
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomFormResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ReplacePortalFormRequest](../../models/operations/replaceportalformrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ReplacePortalFormResponse](../../models/operations/replaceportalformresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeletePortalForm

Deletes a custom form. Any API publications previously linked to this form via `form_id` have the link cleared. Stored form responses collected before deletion are retained and remain viewable via the response detail view.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-portal-form" method="delete" path="/v3/portals/{portalId}/forms/{formId}" -->
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

    res, err := s.PortalForms.DeletePortalForm(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
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
| `portalID`                                               | `string`                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `formID`                                                 | `string`                                                 | :heavy_check_mark:                                       | ID of the portal form.                                   | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeletePortalFormResponse](../../models/operations/deleteportalformresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |