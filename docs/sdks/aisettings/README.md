# AISettings

## Overview

APIs related to Konnect Developer Portal AI Settings.

### Available Operations

* [GetAiSettings](#getaisettings) - Get AI Settings by Portal
* [ReplaceAiSettings](#replaceaisettings) - Replace AI Settings by Portal
* [PatchAiSettings](#patchaisettings) - Patch AI Settings by Portal

## GetAiSettings

Gets AI settings for a given portal including the configuration of all AI features.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-settings" method="get" path="/v3/portals/{portalId}/ai-settings" -->
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

    res, err := s.AISettings.GetAiSettings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.AISettingsResponseBody != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `portalID`                                               | `string`                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAiSettingsResponse](../../models/operations/getaisettingsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ReplaceAiSettings

Replaces AI settings for a given portal including the configuration of all AI features.

### Example Usage

<!-- UsageSnippet language="go" operationID="replace-ai-settings" method="put" path="/v3/portals/{portalId}/ai-settings" -->
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

    res, err := s.AISettings.ReplaceAiSettings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", components.ReplaceAISettingsRequest{
        Features: components.ReplaceAISettingsRequestFeatures{
            McpServer: components.ReplaceAISettingsRequestFeaturesMcpServer{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AISettingsResponseBody != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `portalID`                                                                                 | `string`                                                                                   | :heavy_check_mark:                                                                         | ID of the portal.                                                                          | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                       |
| `replaceAISettingsRequest`                                                                 | [components.ReplaceAISettingsRequest](../../models/components/replaceaisettingsrequest.md) | :heavy_check_mark:                                                                         | Replaces AI settings for a portal.                                                         |                                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |

### Response

**[*operations.ReplaceAiSettingsResponse](../../models/operations/replaceaisettingsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchAiSettings

Patches AI settings for a given portal including the configuration of all AI features.

### Example Usage

<!-- UsageSnippet language="go" operationID="patch-ai-settings" method="patch" path="/v3/portals/{portalId}/ai-settings" -->
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

    res, err := s.AISettings.PatchAiSettings(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", components.PatchAISettingsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.AISettingsResponseBody != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `portalID`                                                                             | `string`                                                                               | :heavy_check_mark:                                                                     | ID of the portal.                                                                      | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                   |
| `patchAISettingsRequest`                                                               | [components.PatchAISettingsRequest](../../models/components/patchaisettingsrequest.md) | :heavy_check_mark:                                                                     | Patch AI settings for a portal.                                                        |                                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

### Response

**[*operations.PatchAiSettingsResponse](../../models/operations/patchaisettingsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |