# ContextInterfaceConfigs

## Overview

### Available Operations

* [GetContextInterfacesSourceMappingCapabilityControls](#getcontextinterfacessourcemappingcapabilitycontrols) - Get Interface-Source Capability Controls
* [CreateContextInterfacesSourceMappingCapabilityControls](#createcontextinterfacessourcemappingcapabilitycontrols) - Create Interface-Source Capability Controls
* [PatchContextInterfacesSourceMappingCapabilityControls](#patchcontextinterfacessourcemappingcapabilitycontrols) - Update Interface-Source Capability Controls

## GetContextInterfacesSourceMappingCapabilityControls

Retrieve the capability controls configured for a context interface-source mapping.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-context-interfaces-source-mapping-capability-controls" method="get" path="/v1/context-interfaces/{interfaceId}/context-source-mappings/{mappingId}/capability-controls" -->
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

    res, err := s.ContextInterfaceConfigs.GetContextInterfacesSourceMappingCapabilityControls(ctx, "0e706248-92bd-45b8-8f4d-6b4478af6416", "7ce40fe5-875f-4d52-bc9b-9cfb9c6aa0f9")
    if err != nil {
        log.Fatal(err)
    }
    if res.CapabilityControls != nil {
        switch res.CapabilityControls.Type {
            case components.CapabilityControlsTypeAPI:
                // res.CapabilityControls.APICapabilityControls is populated
            case components.CapabilityControlsTypeMcpServer:
                // res.CapabilityControls.MCPCapabilityControls is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `interfaceID`                                            | `string`                                                 | :heavy_check_mark:                                       | The ID of the Context Interface.                         |
| `mappingID`                                              | `string`                                                 | :heavy_check_mark:                                       | The ID of the Context Interface Source Mapping.          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetContextInterfacesSourceMappingCapabilityControlsResponse](../../models/operations/getcontextinterfacessourcemappingcapabilitycontrolsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateContextInterfacesSourceMappingCapabilityControls

Create the capability controls for a context interface-source mapping.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-context-interfaces-source-mapping-capability-controls" method="post" path="/v1/context-interfaces/{interfaceId}/context-source-mappings/{mappingId}/capability-controls" -->
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

    res, err := s.ContextInterfaceConfigs.CreateContextInterfacesSourceMappingCapabilityControls(ctx, operations.CreateContextInterfacesSourceMappingCapabilityControlsRequest{
        InterfaceID: "7274382c-6b40-4251-9416-632b50d56a31",
        MappingID: "53123bc8-087e-4b95-b7d6-73bd48285229",
        CapabilityControls: components.CreateCapabilityControlsAPI(
            components.APICapabilityControls{
                Type: components.APICapabilityControlsTypeAPI,
                Deny: components.APICapabilityControlsDeny{
                    Operations: []components.APICapabilityControlsOperation{
                        components.APICapabilityControlsOperation{
                            Path: "/library/{id}/books",
                            Methods: []string{
                                "GET",
                            },
                        },
                    },
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CapabilityControls != nil {
        switch res.CapabilityControls.Type {
            case components.CapabilityControlsTypeAPI:
                // res.CapabilityControls.APICapabilityControls is populated
            case components.CapabilityControlsTypeMcpServer:
                // res.CapabilityControls.MCPCapabilityControls is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `request`                                                                                                                                                            | [operations.CreateContextInterfacesSourceMappingCapabilityControlsRequest](../../models/operations/createcontextinterfacessourcemappingcapabilitycontrolsrequest.md) | :heavy_check_mark:                                                                                                                                                   | The request object to use for the request.                                                                                                                           |
| `opts`                                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                                             | :heavy_minus_sign:                                                                                                                                                   | The options for this request.                                                                                                                                        |

### Response

**[*operations.CreateContextInterfacesSourceMappingCapabilityControlsResponse](../../models/operations/createcontextinterfacessourcemappingcapabilitycontrolsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchContextInterfacesSourceMappingCapabilityControls

Partially update the capability controls configured for a context interface-source mapping. Omitted deny keys preserve their existing value, an empty array clears a capability, and a non-empty array replaces.


### Example Usage

<!-- UsageSnippet language="go" operationID="patch-context-interfaces-source-mapping-capability-controls" method="patch" path="/v1/context-interfaces/{interfaceId}/context-source-mappings/{mappingId}/capability-controls" -->
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

    res, err := s.ContextInterfaceConfigs.PatchContextInterfacesSourceMappingCapabilityControls(ctx, operations.PatchContextInterfacesSourceMappingCapabilityControlsRequest{
        InterfaceID: "0430575f-04b7-4e81-878e-bcec778af9ff",
        MappingID: "aca0d828-07b2-4237-8031-3c32f5a907f5",
        PatchCapabilityControls: components.CreatePatchCapabilityControlsAPI(
            components.PatchAPICapabilityControls{
                Type: components.PatchAPICapabilityControlsTypeAPI,
                Deny: components.PatchAPICapabilityControlsDeny{
                    Operations: []components.APICapabilityControlsOperation{
                        components.APICapabilityControlsOperation{
                            Path: "/library/{id}/books",
                            Methods: []string{
                                "GET",
                            },
                        },
                    },
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CapabilityControls != nil {
        switch res.CapabilityControls.Type {
            case components.CapabilityControlsTypeAPI:
                // res.CapabilityControls.APICapabilityControls is populated
            case components.CapabilityControlsTypeMcpServer:
                // res.CapabilityControls.MCPCapabilityControls is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `request`                                                                                                                                                          | [operations.PatchContextInterfacesSourceMappingCapabilityControlsRequest](../../models/operations/patchcontextinterfacessourcemappingcapabilitycontrolsrequest.md) | :heavy_check_mark:                                                                                                                                                 | The request object to use for the request.                                                                                                                         |
| `opts`                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                           | :heavy_minus_sign:                                                                                                                                                 | The options for this request.                                                                                                                                      |

### Response

**[*operations.PatchContextInterfacesSourceMappingCapabilityControlsResponse](../../models/operations/patchcontextinterfacessourcemappingcapabilitycontrolsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |