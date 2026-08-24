# Skills

## Overview

### Available Operations

* [ListMcpServerSkillsByControlPlane](#listmcpserverskillsbycontrolplane) - List Skills by Control Plane
* [GetMcpServerSkillContentsByControlPlane](#getmcpserverskillcontentsbycontrolplane) - Get Skill Contents by Control Plane

## ListMcpServerSkillsByControlPlane

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Separate control-plane-scoped list endpoint for KO to enumerate skills at deploy time.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-mcp-server-skills-by-control-plane" method="get" path="/v1/mcp-cp/{controlPlaneId}/mcp-servers/{mcpServerId}/skills" -->
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

    res, err := s.Skills.ListMcpServerSkillsByControlPlane(ctx, operations.ListMcpServerSkillsByControlPlaneRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        McpServerID: "85bf863b-78d5-4b8f-a114-c9609f990798",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListSkillsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.ListMcpServerSkillsByControlPlaneRequest](../../models/operations/listmcpserverskillsbycontrolplanerequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.ListMcpServerSkillsByControlPlaneResponse](../../models/operations/listmcpserverskillsbycontrolplaneresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetMcpServerSkillContentsByControlPlane

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

KO-only endpoint that redirects to a presigned S3 URL for the skill's SKILL.md file.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mcp-server-skill-contents-by-control-plane" method="get" path="/v1/mcp-cp/{controlPlaneId}/mcp-servers/{mcpServerId}/skills/{skillId}/contents" -->
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

    res, err := s.Skills.GetMcpServerSkillContentsByControlPlane(ctx, operations.GetMcpServerSkillContentsByControlPlaneRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        McpServerID: "5bae9b81-0da6-4ba2-ba78-bb99f822149e",
        SkillID: "8bb561d4-55c6-493f-8f2c-145ac063fea3",
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

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.GetMcpServerSkillContentsByControlPlaneRequest](../../models/operations/getmcpserverskillcontentsbycontrolplanerequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.GetMcpServerSkillContentsByControlPlaneResponse](../../models/operations/getmcpserverskillcontentsbycontrolplaneresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |