# MCP Client for Kong Konnect SDK

This package provides a Model Context Protocol (MCP) client implementation for the Kong Konnect Go SDK.

## What is MCP?

Model Context Protocol (MCP) is an open protocol that enables AI applications to securely connect to external data sources and tools. It provides a standardized way for LLMs and AI systems to interact with various services using JSON-RPC 2.0.

## Features

- ✅ Full JSON-RPC 2.0 implementation
- ✅ Tool discovery and invocation
- ✅ Resource listing and reading
- ✅ Context-aware requests with cancellation support
- ✅ Comprehensive error handling
- ✅ Type-safe API
- ✅ Production-ready with tests

## Installation

```bash
go get github.com/Kong/sdk-konnect-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/Kong/sdk-konnect-go/pkg/mcp"
)

func main() {
    // Create a new MCP client
    client := mcp.NewClient("http://localhost:3000/mcp")

    // Initialize the connection
    err := client.Initialize(context.Background(), mcp.ClientInfo{
        Name:    "my-app",
        Version: "1.0.0",
    })
    if err != nil {
        log.Fatal(err)
    }

    // List available tools
    tools, err := client.ListTools(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d tools\n", len(tools.Tools))

    // Call a tool
    result, err := client.CallTool(context.Background(), mcp.CallToolRequest{
        Name: "calculator",
        Arguments: map[string]interface{}{
            "operation": "add",
            "a": 5,
            "b": 3,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Result: %s\n", result.Content[0].Text)
}
```

## API Reference

### Client

#### `NewClient(serverURL string, httpClient ...*http.Client) *Client`

Creates a new MCP client instance.

**Parameters:**
- `serverURL`: The URL of the MCP server
- `httpClient`: Optional custom HTTP client (uses default if not provided)

**Returns:** A new `Client` instance

#### `Initialize(ctx context.Context, clientInfo ClientInfo) error`

Initializes the MCP connection with the server. This must be called before any other operations.

**Parameters:**
- `ctx`: Context for the request
- `clientInfo`: Information about the client (name and version)

**Returns:** Error if initialization fails

#### `ListTools(ctx context.Context, cursor *string) (*ListToolsResponse, error)`

Lists all available tools from the MCP server.

**Parameters:**
- `ctx`: Context for the request
- `cursor`: Optional pagination cursor for large result sets

**Returns:** List of tools or error

#### `CallTool(ctx context.Context, req CallToolRequest) (*CallToolResponse, error)`

Calls a tool on the MCP server.

**Parameters:**
- `ctx`: Context for the request
- `req`: Tool call request with tool name and arguments

**Returns:** Tool result or error

#### `ListResources(ctx context.Context, cursor *string) (*ListResourcesResponse, error)`

Lists all available resources from the MCP server.

**Parameters:**
- `ctx`: Context for the request
- `cursor`: Optional pagination cursor

**Returns:** List of resources or error

#### `ReadResource(ctx context.Context, uri string) (*ReadResourceResponse, error)`

Reads a resource from the MCP server.

**Parameters:**
- `ctx`: Context for the request
- `uri`: URI of the resource to read

**Returns:** Resource content or error

#### `ServerInfo() *ServerInfo`

Returns information about the connected MCP server.

#### `Capabilities() *ServerCapabilities`

Returns the capabilities supported by the MCP server.

## Types

### Tool

Represents a tool exposed by the MCP server:

```go
type Tool struct {
    Name        string      `json:"name"`
    Description string      `json:"description,omitempty"`
    InputSchema InputSchema `json:"inputSchema"`
}
```

### Resource

Represents a resource exposed by the MCP server:

```go
type Resource struct {
    URI         string `json:"uri"`
    Name        string `json:"name"`
    Description string `json:"description,omitempty"`
    MimeType    string `json:"mimeType,omitempty"`
}
```

### CallToolRequest

Request for calling a tool:

```go
type CallToolRequest struct {
    Name      string                 `json:"name"`
    Arguments map[string]interface{} `json:"arguments,omitempty"`
}
```

### CallToolResponse

Response from calling a tool:

```go
type CallToolResponse struct {
    Content []Content `json:"content"`
    IsError bool      `json:"isError,omitempty"`
}
```

## Error Handling

The MCP client returns descriptive errors for various failure scenarios:

```go
client := mcp.NewClient("http://localhost:3000/mcp")

result, err := client.CallTool(ctx, mcp.CallToolRequest{
    Name: "calculator",
    Arguments: map[string]interface{}{
        "operation": "divide",
        "a": 10,
        "b": 0,
    },
})
if err != nil {
    if strings.Contains(err.Error(), "JSON-RPC error") {
        // Handle JSON-RPC specific errors
        log.Printf("Tool execution failed: %v", err)
    } else {
        // Handle network or other errors
        log.Printf("Request failed: %v", err)
    }
    return
}
```

## Context Support

All client methods support context for timeouts and cancellation:

```go
// Create a context with timeout
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

// Use the context in client calls
tools, err := client.ListTools(ctx, nil)
if err != nil {
    if ctx.Err() == context.DeadlineExceeded {
        log.Println("Request timed out")
    }
}
```

## Examples

See the [examples](./examples/) directory for more usage examples:

- [basic_usage.go](./examples/basic_usage.go) - Basic MCP client usage

## Protocol Support

This client implements the Model Context Protocol version `2024-11-05`.

Supported features:
- ✅ Tool listing and invocation
- ✅ Resource listing and reading
- ✅ JSON-RPC 2.0 over HTTP
- ✅ Pagination support
- [Planned] Prompts
- [Planned] Sampling
- [Planned] WebSocket transport

## Testing

Run the tests:

```bash
go test ./pkg/mcp/...
```

Run tests with coverage:

```bash
go test -cover ./pkg/mcp/...
```

## Contributing

Contributions are welcome! Please ensure:

1. All tests pass
2. Code follows Go conventions
3. New features include tests
4. Documentation is updated

## License

See the [LICENSE](../../LICENSE) file in the root of the repository.

## Resources

- [Model Context Protocol Specification](https://modelcontextprotocol.io/)
- [MCP Documentation](https://modelcontextprotocol.io/docs/)
- [Kong Konnect Documentation](https://docs.konghq.com/konnect/)
