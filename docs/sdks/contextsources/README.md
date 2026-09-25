# ContextSources

## Overview

### Available Operations

* [ListContextSources](#listcontextsources) - List Context Sources
* [CreateContextSource](#createcontextsource) - Create a Context Source
* [GetContextSource](#getcontextsource) - Get a Context Source
* [DeleteContextSource](#deletecontextsource) - Delete a Context Source
* [UpdateContextSource](#updatecontextsource) - Update a Context Source

## ListContextSources

Returns a list of Context Source objects.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-context-sources" method="get" path="/v1/context-sources" -->
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

    res, err := s.ContextSources.ListContextSources(ctx, operations.ListContextSourcesRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListMCPResourcesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListContextSourcesRequest](../../models/operations/listcontextsourcesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListContextSourcesResponse](../../models/operations/listcontextsourcesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateContextSource

Create a Context Source in the Konnect Organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-context-source" method="post" path="/v1/context-sources" -->
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

    res, err := s.ContextSources.CreateContextSource(ctx, operations.CreateCreateContextSourceRequestBodyAPIResourcePayload(
        components.APIResourcePayload{
            Name: "my-api-resource",
            DisplayName: "Hershel.Bechtelar",
            Description: "wherever via animated via beyond and brr lest",
            Labels: map[string]string{
                "env": "test",
            },
            Source: components.CreateAPIResourceSourcePayloadAPICatalog(
                components.MCPResourceSourceAPICatalogPayload{
                    Type: components.MCPResourceSourceAPICatalogPayloadTypeAPICatalog,
                    Config: components.MCPResourceSourceAPICatalogConfigPayload{
                        APIID: "158a0fc7-479a-4928-8a64-7656c1824f7f",
                        APIVersionID: "bc0b9df1-67ba-475f-9670-37de3cc2e4f6",
                    },
                },
            ),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPResourceInfo != nil {
        switch res.MCPResourceInfo.Type {
            case components.MCPResourceInfoTypeAPI:
                // res.MCPResourceInfo.APIResource is populated
            case components.MCPResourceInfoTypeMcpServer:
                // res.MCPResourceInfo.McpServerResource is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateContextSourceRequestBody](../../models/operations/createcontextsourcerequestbody.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateContextSourceResponse](../../models/operations/createcontextsourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetContextSource

Retrieve an Context Source by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-context-source" method="get" path="/v1/context-sources/{sourceId}" -->
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

    res, err := s.ContextSources.GetContextSource(ctx, "565f7f07-49c7-4591-accd-38f5cc4879d6")
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPResourceInfo != nil {
        switch res.MCPResourceInfo.Type {
            case components.MCPResourceInfoTypeAPI:
                // res.MCPResourceInfo.APIResource is populated
            case components.MCPResourceInfoTypeMcpServer:
                // res.MCPResourceInfo.McpServerResource is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `sourceID`                                               | `string`                                                 | :heavy_check_mark:                                       | The ID of the Context Source.                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetContextSourceResponse](../../models/operations/getcontextsourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteContextSource

Delete a Context Source by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-context-source" method="delete" path="/v1/context-sources/{sourceId}" -->
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

    res, err := s.ContextSources.DeleteContextSource(ctx, "6977c918-824d-4366-b25b-34a19d6f5b4f")
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
| `sourceID`                                               | `string`                                                 | :heavy_check_mark:                                       | The ID of the Context Source.                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteContextSourceResponse](../../models/operations/deletecontextsourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateContextSource

Update a Context Source by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-context-source" method="put" path="/v1/context-sources/{sourceId}" -->
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

    res, err := s.ContextSources.UpdateContextSource(ctx, "90fc8931-cf90-4890-bb3b-07cbb228526c", operations.CreateUpdateContextSourceRequestBodyMcpServerResourcePayload(
        components.McpServerResourcePayload{
            Name: "my-api-resource",
            DisplayName: "Thora9",
            Description: "save developing deliberately heartache which huzzah",
            Labels: map[string]string{
                "env": "test",
            },
            Source: components.CreateMcpServerResourceSourcePayloadRemoteMcpServer(
                components.MCPResourceRemoteMCPServerPayload{
                    Type: components.MCPResourceRemoteMCPServerPayloadTypeRemoteMcpServer,
                    Config: components.MCPResourceRemoteMCPServerPayloadConfig{
                        URL: "https://parallel-kettledrum.org/",
                        Headers: map[string]string{
                            "key": "<value>",
                        },
                    },
                },
            ),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPResourceInfo != nil {
        switch res.MCPResourceInfo.Type {
            case components.MCPResourceInfoTypeAPI:
                // res.MCPResourceInfo.APIResource is populated
            case components.MCPResourceInfoTypeMcpServer:
                // res.MCPResourceInfo.McpServerResource is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `sourceID`                                                                                             | `string`                                                                                               | :heavy_check_mark:                                                                                     | The ID of the Context Source.                                                                          |
| `requestBody`                                                                                          | [operations.UpdateContextSourceRequestBody](../../models/operations/updatecontextsourcerequestbody.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateContextSourceResponse](../../models/operations/updatecontextsourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |