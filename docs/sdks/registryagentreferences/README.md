# RegistryAgentReferences

## Overview

### Available Operations

* [ListAgents](#listagents) - List agents
* [ListAgentPrincipals](#listagentprincipals) - List agent principals

## ListAgents

Returns a paginated list of agents. Default order is most recently created first (`created_at DESC, id DESC`).


### Example Usage

<!-- UsageSnippet language="go" operationID="list-agents" method="get" path="/v1/agents" -->
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
            KonnectAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.RegistryAgentReferences.ListAgents(ctx, operations.ListAgentsRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.RegistryReferenceAgentFilterParameters{
            CreatedAt: sdkkonnectgo.Pointer(components.CreateRegistryReferenceAgentFilterParametersDateTimeFieldFilterDateTime(
                types.MustTimeFromString("2022-03-30T07:20:50Z"),
            )),
            UpdatedAt: sdkkonnectgo.Pointer(components.CreateCreatedAtDateTime(
                types.MustTimeFromString("2022-03-30T07:20:50Z"),
            )),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegistryReferenceAgentList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListAgentsRequest](../../models/operations/listagentsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.ListAgentsResponse](../../models/operations/listagentsresponse.md), error**

### Errors

| Error Type                                               | Status Code                                              | Content Type                                             |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| sdkerrors.ListAgentsError                                | 400                                                      | application/problem+json                                 |
| sdkerrors.ListAgentsRegistryAgentReferencesError         | 401                                                      | application/problem+json                                 |
| sdkerrors.ListAgentsRegistryAgentReferencesResponseError | 403                                                      | application/problem+json                                 |
| sdkerrors.SDKError                                       | 4XX, 5XX                                                 | \*/\*                                                    |

## ListAgentPrincipals

Returns the agent's identity principals (0 or 1). An agent has at most one linked principal.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-agent-principals" method="get" path="/v1/agents/{agentId}/principals" -->
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
            KonnectAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.RegistryAgentReferences.ListAgentPrincipals(ctx, operations.ListAgentPrincipalsRequest{
        AgentID: "123e4567-e89b-12d3-a456-426614174000",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegistryReferenceAgentPrincipalList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListAgentPrincipalsRequest](../../models/operations/listagentprincipalsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListAgentPrincipalsResponse](../../models/operations/listagentprincipalsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.RegistryReferenceReadAnchor4 | 401                                    | application/problem+json               |
| sdkerrors.RegistryReferenceReadAnchor5 | 403                                    | application/problem+json               |
| sdkerrors.ListAgentPrincipalsError     | 404                                    | application/problem+json               |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |