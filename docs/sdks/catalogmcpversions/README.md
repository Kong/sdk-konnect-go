# CatalogMCPVersions

## Overview

Manage an MCP's version.

### Available Operations

* [ListMcpVersions](#listmcpversions) - List MCP versions
* [CreateMcpVersion](#createmcpversion) - Create MCP version
* [GetLatestMcpVersion](#getlatestmcpversion) - Get latest MCP version
* [UpdateLatestMcpVersion](#updatelatestmcpversion) - Update latest MCP version
* [UpsertLatestMcpVersion](#upsertlatestmcpversion) - Replace latest MCP version
* [DeleteLatestMcpVersion](#deletelatestmcpversion) - Delete latest MCP version

## ListMcpVersions

Returns the MCP's versions.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-mcp-versions" method="get" path="/v1/mcp-servers/{mcpId}/versions" -->
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

    res, err := s.CatalogMCPVersions.ListMcpVersions(ctx, operations.ListMcpVersionsRequest{
        McpID: "a0119846-f179-4d9f-a168-d701facce7fb",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCatalogMCPVersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListMcpVersionsRequest](../../models/operations/listmcpversionsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListMcpVersionsResponse](../../models/operations/listmcpversionsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateMcpVersion

Creates the `latest` version. Currently, a MCP only has a single version. Returns `409` if a version already exists.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-mcp-version" method="post" path="/v1/mcp-servers/{mcpId}/versions" -->
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

    res, err := s.CatalogMCPVersions.CreateMcpVersion(ctx, "a0119846-f179-4d9f-a168-d701facce7fb", components.CreateCatalogMCPVersion{
        Version: "1.0.0",
        Tools: []components.CatalogMCPTool{},
        Resources: []components.CatalogMCPResource{},
        Prompts: []components.CatalogMCPPrompt{
            components.CatalogMCPPrompt{
                Name: "<value>",
            },
        },
        Remotes: nil,
        Packages: []components.MCPPackage{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogMCPVersion != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `mcpID`                                                                                  | `string`                                                                                 | :heavy_check_mark:                                                                       | The unique identifier of the MCP.                                                        | a0119846-f179-4d9f-a168-d701facce7fb                                                     |
| `createCatalogMCPVersion`                                                                | [components.CreateCatalogMCPVersion](../../models/components/createcatalogmcpversion.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |                                                                                          |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.CreateMcpVersionResponse](../../models/operations/createmcpversionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetLatestMcpVersion

Returns the MCP's `latest` version.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-latest-mcp-version" method="get" path="/v1/mcp-servers/{mcpId}/versions/latest" -->
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

    res, err := s.CatalogMCPVersions.GetLatestMcpVersion(ctx, "a0119846-f179-4d9f-a168-d701facce7fb")
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogMCPVersion != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `mcpID`                                                  | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the MCP.                        | a0119846-f179-4d9f-a168-d701facce7fb                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetLatestMcpVersionResponse](../../models/operations/getlatestmcpversionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateLatestMcpVersion

Partially updates the `latest` version.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-latest-mcp-version" method="patch" path="/v1/mcp-servers/{mcpId}/versions/latest" -->
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

    res, err := s.CatalogMCPVersions.UpdateLatestMcpVersion(ctx, "a0119846-f179-4d9f-a168-d701facce7fb", components.UpdateCatalogMCPVersion{
        Version: sdkkonnectgo.Pointer("1.0.0"),
        Tools: nil,
        Resources: []components.CatalogMCPResource{
            components.CatalogMCPResource{
                Name: "my-resource",
                URI: "https://graceful-handover.biz",
            },
        },
        Remotes: []components.MCPRemoteTransport{
            components.CreateMCPRemoteTransportMCPStreamableHTTPTransportMCPStreamableHTTPTransport(
                components.MCPStreamableHTTPTransportMCPStreamableHTTPTransport{
                    URL: "https://api.example.com/mcp",
                },
            ),
        },
        Packages: []components.MCPPackage{
            components.MCPPackage{
                Registry: components.Registry{
                    Type: "npm",
                    BaseURL: sdkkonnectgo.Pointer("https://registry.npmjs.org"),
                },
                Identifier: "@modelcontextprotocol/server-brave-search",
                Version: sdkkonnectgo.Pointer("1.0.2"),
                FileSha256: sdkkonnectgo.Pointer("fe333e598595000ae021bd27117db32ec69af6987f507ba7a63c90638ff633ce"),
                Transport: components.CreateMCPTransportMCPStdioTransport(
                    components.MCPStdioTransport{},
                ),
                Runtime: &components.Runtime{
                    Hint: sdkkonnectgo.Pointer("npx"),
                    Arguments: []components.MCPArgument{
                        components.CreateMCPArgumentMCPPositionalArgument(
                            components.CreateMCPPositionalArgumentMCPPositionalArgument1(
                                components.MCPPositionalArgument1{
                                    ValueHint: "file_path",
                                },
                            ),
                        ),
                    },
                },
                PackageArguments: []components.MCPArgument{
                    components.CreateMCPArgumentMCPPositionalArgument(
                        components.CreateMCPPositionalArgumentMCPPositionalArgument1(
                            components.MCPPositionalArgument1{
                                ValueHint: "file_path",
                            },
                        ),
                    ),
                },
                EnvironmentVariables: []components.MCPKeyValueInput{
                    components.MCPKeyValueInput{
                        Choices: []string{},
                        Variables: map[string]components.MCPInput{
                            "key": components.MCPInput{
                                Choices: []string{},
                            },
                        },
                        Name: "SOME_VARIABLE",
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogMCPVersion != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `mcpID`                                                                                  | `string`                                                                                 | :heavy_check_mark:                                                                       | The unique identifier of the MCP.                                                        | a0119846-f179-4d9f-a168-d701facce7fb                                                     |
| `updateCatalogMCPVersion`                                                                | [components.UpdateCatalogMCPVersion](../../models/components/updatecatalogmcpversion.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |                                                                                          |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.UpdateLatestMcpVersionResponse](../../models/operations/updatelatestmcpversionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpsertLatestMcpVersion

Replaces the `latest` version.


### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-latest-mcp-version" method="put" path="/v1/mcp-servers/{mcpId}/versions/latest" -->
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

    res, err := s.CatalogMCPVersions.UpsertLatestMcpVersion(ctx, "a0119846-f179-4d9f-a168-d701facce7fb", components.CreateCatalogMCPVersion{
        Version: "1.0.0",
        Tools: nil,
        Resources: []components.CatalogMCPResource{},
        Prompts: []components.CatalogMCPPrompt{},
        Remotes: []components.MCPRemoteTransport{},
        Packages: []components.MCPPackage{
            components.MCPPackage{
                Registry: components.Registry{
                    Type: "npm",
                    BaseURL: sdkkonnectgo.Pointer("https://registry.npmjs.org"),
                },
                Identifier: "@modelcontextprotocol/server-brave-search",
                Version: sdkkonnectgo.Pointer("1.0.2"),
                FileSha256: sdkkonnectgo.Pointer("fe333e598595000ae021bd27117db32ec69af6987f507ba7a63c90638ff633ce"),
                Transport: components.CreateMCPTransportMCPStdioTransport(
                    components.MCPStdioTransport{},
                ),
                Runtime: &components.Runtime{
                    Hint: sdkkonnectgo.Pointer("npx"),
                    Arguments: []components.MCPArgument{
                        components.CreateMCPArgumentMCPPositionalArgument(
                            components.CreateMCPPositionalArgumentMCPPositionalArgument1(
                                components.MCPPositionalArgument1{
                                    ValueHint: "file_path",
                                },
                            ),
                        ),
                    },
                },
                PackageArguments: []components.MCPArgument{
                    components.CreateMCPArgumentMCPPositionalArgument(
                        components.CreateMCPPositionalArgumentMCPPositionalArgument1(
                            components.MCPPositionalArgument1{
                                ValueHint: "file_path",
                            },
                        ),
                    ),
                },
                EnvironmentVariables: []components.MCPKeyValueInput{
                    components.MCPKeyValueInput{
                        Choices: []string{},
                        Variables: map[string]components.MCPInput{
                            "key": components.MCPInput{
                                Choices: []string{},
                            },
                        },
                        Name: "SOME_VARIABLE",
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogMCPVersion != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `mcpID`                                                                                  | `string`                                                                                 | :heavy_check_mark:                                                                       | The unique identifier of the MCP.                                                        | a0119846-f179-4d9f-a168-d701facce7fb                                                     |
| `createCatalogMCPVersion`                                                                | [components.CreateCatalogMCPVersion](../../models/components/createcatalogmcpversion.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |                                                                                          |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.UpsertLatestMcpVersionResponse](../../models/operations/upsertlatestmcpversionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteLatestMcpVersion

Deletes the MCP's `latest` version. Returns `404` if no version exists yet.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-latest-mcp-version" method="delete" path="/v1/mcp-servers/{mcpId}/versions/latest" -->
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

    res, err := s.CatalogMCPVersions.DeleteLatestMcpVersion(ctx, "a0119846-f179-4d9f-a168-d701facce7fb")
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
| `mcpID`                                                  | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the MCP.                        | a0119846-f179-4d9f-a168-d701facce7fb                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteLatestMcpVersionResponse](../../models/operations/deletelatestmcpversionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |