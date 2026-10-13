# SignalingAPIs

## Overview

### Available Operations

* [ListDeployedContextInterfaceSkillArtifacts](#listdeployedcontextinterfaceskillartifacts) - List Skill Artifacts By Control Plane

## ListDeployedContextInterfaceSkillArtifacts

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Retrieve list of skill artifacts from deployed context interface

### Example Usage

<!-- UsageSnippet language="go" operationID="list-deployed-context-interface-skill-artifacts" method="get" path="/v1/mcp-cp/{controlPlaneId}/context-interfaces/{interfaceId}/skills/{skillId}/artifacts" -->
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

    res, err := s.SignalingAPIs.ListDeployedContextInterfaceSkillArtifacts(ctx, operations.ListDeployedContextInterfaceSkillArtifactsRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        InterfaceID: "531f3c71-c16a-4ff1-a19c-7ab05bef775e",
        SkillID: "b2688c99-5c59-40f4-88b3-1aa04cd15f42",
        Page: &components.CursorPageParameters{
            Size: sdkkonnectgo.Pointer[int64](10),
            After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
            Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListSkillArtifactsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.ListDeployedContextInterfaceSkillArtifactsRequest](../../models/operations/listdeployedcontextinterfaceskillartifactsrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.ListDeployedContextInterfaceSkillArtifactsResponse](../../models/operations/listdeployedcontextinterfaceskillartifactsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |