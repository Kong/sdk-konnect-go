# Skills

## Overview

### Available Operations

* [ListMcpServerSkillsByControlPlane](#listmcpserverskillsbycontrolplane) - List Skills by Control Plane
* [GetMcpServerSkillContentsByControlPlane](#getmcpserverskillcontentsbycontrolplane) - Get Skill Contents by Control Plane
* [GetContextInterfaceSkill](#getcontextinterfaceskill) - Get a Skill
* [PatchContextInterfaceSkill](#patchcontextinterfaceskill) - Update a Skill
* [DeleteContextInterfaceSkill](#deletecontextinterfaceskill) - Delete a Skill
* [GetContextInterfaceSkillContents](#getcontextinterfaceskillcontents) - Get Skill Contents
* [ListContextInterfaceSkills](#listcontextinterfaceskills) - List Skills
* [CreateContextInterfaceSkill](#createcontextinterfaceskill) - Create a Skill

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

## GetContextInterfaceSkill

Retrieve a single skill by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-context-interface-skill" method="get" path="/v1/context-interfaces/{interfaceId}/skills/{skillId}" -->
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

    res, err := s.Skills.GetContextInterfaceSkill(ctx, "90858168-eee7-4e67-b207-88eff5fab482", "4dd1b9bb-40aa-4f9d-9a11-a70344737178")
    if err != nil {
        log.Fatal(err)
    }
    if res.Skill != nil {
        switch res.Skill.Source.Type {
            case components.SkillSourceTypeRaw:
                // res.Skill.Source.RawSkillSource is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `interfaceID`                                            | `string`                                                 | :heavy_check_mark:                                       | The ID of the Context Interface.                         |
| `skillID`                                                | `string`                                                 | :heavy_check_mark:                                       | The ID of the skill.                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetContextInterfaceSkillResponse](../../models/operations/getcontextinterfaceskillresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchContextInterfaceSkill

Partially update a skill. Supplying `source` replaces the skill's source in its entirety and re-runs validation against the new content.


### Example Usage

<!-- UsageSnippet language="go" operationID="patch-context-interface-skill" method="patch" path="/v1/context-interfaces/{interfaceId}/skills/{skillId}" -->
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

    res, err := s.Skills.PatchContextInterfaceSkill(ctx, operations.PatchContextInterfaceSkillRequest{
        InterfaceID: "9cb81f58-8e7e-4d0e-9f56-26153eac497d",
        SkillID: "12917b31-c65d-47ce-9861-19c8a52b252b",
        PatchSkillRequest: components.PatchSkillRequest{
            Source: sdkkonnectgo.Pointer(components.CreatePatchSkillRequestSourceRaw(
                components.RawSkillSourcePayload{
                    Type: components.RawSkillSourcePayloadTypeRaw,
                    Config: components.RawSkillSourceConfigPayload{
                        Contents: "---\nname: pdf-processing\ndescription: Extract and summarize content from PDF documents\n---\n# PDF Processing\n...",
                    },
                },
            )),
            Labels: map[string]string{
                "env": "test",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Skill != nil {
        switch res.Skill.Source.Type {
            case components.SkillSourceTypeRaw:
                // res.Skill.Source.RawSkillSource is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PatchContextInterfaceSkillRequest](../../models/operations/patchcontextinterfaceskillrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.PatchContextInterfaceSkillResponse](../../models/operations/patchcontextinterfaceskillresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteContextInterfaceSkill

Delete a skill and its associated content.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-context-interface-skill" method="delete" path="/v1/context-interfaces/{interfaceId}/skills/{skillId}" -->
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

    res, err := s.Skills.DeleteContextInterfaceSkill(ctx, "981cf816-3aa2-4512-8a55-605428b2c21a", "a32f588f-2443-46f7-b0d7-c51dfecb8676")
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
| `interfaceID`                                            | `string`                                                 | :heavy_check_mark:                                       | The ID of the Context Interface.                         |
| `skillID`                                                | `string`                                                 | :heavy_check_mark:                                       | The ID of the skill.                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteContextInterfaceSkillResponse](../../models/operations/deletecontextinterfaceskillresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetContextInterfaceSkillContents

Returns a presigned S3 URL for the skill's SKILL.md file. Used by the Konnect UI to let a user view a skill's content.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-context-interface-skill-contents" method="get" path="/v1/context-interfaces/{interfaceId}/skills/{skillId}/contents" -->
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

    res, err := s.Skills.GetContextInterfaceSkillContents(ctx, "5d5b3442-02ce-49e6-bdd3-ee92354b3ae6", "68feb49a-7718-4b17-add4-9a3cfd51748c")
    if err != nil {
        log.Fatal(err)
    }
    if res.SkillContentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `interfaceID`                                            | `string`                                                 | :heavy_check_mark:                                       | The ID of the Context Interface.                         |
| `skillID`                                                | `string`                                                 | :heavy_check_mark:                                       | The ID of the skill.                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetContextInterfaceSkillContentsResponse](../../models/operations/getcontextinterfaceskillcontentsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListContextInterfaceSkills

Returns a list of skills for the specified Context Interface.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-context-interface-skills" method="get" path="/v1/context-interfaces/{interfaceId}/skills" -->
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

    res, err := s.Skills.ListContextInterfaceSkills(ctx, operations.ListContextInterfaceSkillsRequest{
        InterfaceID: "72fe3c27-ab46-4ad5-9698-e8870e4b55af",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListContextInterfaceSkillsRequest](../../models/operations/listcontextinterfaceskillsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListContextInterfaceSkillsResponse](../../models/operations/listcontextinterfaceskillsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateContextInterfaceSkill

Create a skill for the specified Context Interface. The `source` determines where the skill's content comes from.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-context-interface-skill" method="post" path="/v1/context-interfaces/{interfaceId}/skills" -->
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

    res, err := s.Skills.CreateContextInterfaceSkill(ctx, "cc5067fd-3514-4974-bc56-7c504711542c", components.CreateSkillRequest{
        Name: "pdf-processing",
        DisplayName: "PDF Processing",
        Description: "Extract and summarize content from PDF documents",
        Source: components.CreateSkillSourcePayloadRaw(
            components.RawSkillSourcePayload{
                Type: components.RawSkillSourcePayloadTypeRaw,
                Config: components.RawSkillSourceConfigPayload{
                    Contents: "---\nname: pdf-processing\ndescription: Extract and summarize content from PDF documents\n---\n# PDF Processing\n...",
                },
            },
        ),
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Skill != nil {
        switch res.Skill.Source.Type {
            case components.SkillSourceTypeRaw:
                // res.Skill.Source.RawSkillSource is populated
        }

    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `interfaceID`                                                                  | `string`                                                                       | :heavy_check_mark:                                                             | The ID of the Context Interface.                                               |
| `createSkillRequest`                                                           | [components.CreateSkillRequest](../../models/components/createskillrequest.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CreateContextInterfaceSkillResponse](../../models/operations/createcontextinterfaceskillresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |