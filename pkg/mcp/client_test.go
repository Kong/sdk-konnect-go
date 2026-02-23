package mcp

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	client := NewClient("http://localhost:8080")
	assert.NotNil(t, client)
	assert.Equal(t, "http://localhost:8080", client.serverURL)
	assert.NotNil(t, client.httpClient)
}

func TestInitialize(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req JSONRPCRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		require.NoError(t, err)

		if req.Method == "initialize" {
			resp := JSONRPCResponse{
				JSONRPC: "2.0",
				ID:      req.ID,
				Result: json.RawMessage(`{
					"protocolVersion": "2024-11-05",
					"capabilities": {
						"tools": {"listChanged": true},
						"resources": {"subscribe": true}
					},
					"serverInfo": {
						"name": "test-server",
						"version": "1.0.0"
					}
				}`),
			}
			json.NewEncoder(w).Encode(resp)
		} else if req.Method == "notifications/initialized" {
			// No response for notifications
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer server.Close()

	client := NewClient(server.URL)
	err := client.Initialize(context.Background(), ClientInfo{
		Name:    "test-client",
		Version: "1.0.0",
	})

	require.NoError(t, err)
	assert.NotNil(t, client.serverInfo)
	assert.Equal(t, "test-server", client.serverInfo.Name)
	assert.Equal(t, "1.0.0", client.serverInfo.Version)
	assert.NotNil(t, client.capabilities)
	assert.NotNil(t, client.capabilities.Tools)
	assert.True(t, client.capabilities.Tools.ListChanged)
}

func TestListTools(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req JSONRPCRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		require.NoError(t, err)

		assert.Equal(t, "tools/list", req.Method)

		resp := JSONRPCResponse{
			JSONRPC: "2.0",
			ID:      req.ID,
			Result: json.RawMessage(`{
				"tools": [
					{
						"name": "get_weather",
						"description": "Get weather information",
						"inputSchema": {
							"type": "object",
							"properties": {
								"location": {
									"type": "string",
									"description": "City name"
								}
							},
							"required": ["location"]
						}
					}
				]
			}`),
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	tools, err := client.ListTools(context.Background(), nil)

	require.NoError(t, err)
	assert.NotNil(t, tools)
	assert.Len(t, tools.Tools, 1)
	assert.Equal(t, "get_weather", tools.Tools[0].Name)
	assert.Equal(t, "Get weather information", tools.Tools[0].Description)
	assert.Equal(t, "object", tools.Tools[0].InputSchema.Type)
	assert.Len(t, tools.Tools[0].InputSchema.Required, 1)
	assert.Equal(t, "location", tools.Tools[0].InputSchema.Required[0])
}

func TestCallTool(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req JSONRPCRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		require.NoError(t, err)

		assert.Equal(t, "tools/call", req.Method)

		resp := JSONRPCResponse{
			JSONRPC: "2.0",
			ID:      req.ID,
			Result: json.RawMessage(`{
				"content": [
					{
						"type": "text",
						"text": "The weather in London is sunny and 72°F"
					}
				]
			}`),
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	result, err := client.CallTool(context.Background(), CallToolRequest{
		Name: "get_weather",
		Arguments: map[string]interface{}{
			"location": "London",
		},
	})

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Content, 1)
	assert.Equal(t, "text", result.Content[0].Type)
	assert.Equal(t, "The weather in London is sunny and 72°F", result.Content[0].Text)
	assert.False(t, result.IsError)
}

func TestListResources(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req JSONRPCRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		require.NoError(t, err)

		assert.Equal(t, "resources/list", req.Method)

		resp := JSONRPCResponse{
			JSONRPC: "2.0",
			ID:      req.ID,
			Result: json.RawMessage(`{
				"resources": [
					{
						"uri": "file:///data/readme.txt",
						"name": "README",
						"description": "Project documentation",
						"mimeType": "text/plain"
					}
				]
			}`),
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	resources, err := client.ListResources(context.Background(), nil)

	require.NoError(t, err)
	assert.NotNil(t, resources)
	assert.Len(t, resources.Resources, 1)
	assert.Equal(t, "file:///data/readme.txt", resources.Resources[0].URI)
	assert.Equal(t, "README", resources.Resources[0].Name)
	assert.Equal(t, "Project documentation", resources.Resources[0].Description)
	assert.Equal(t, "text/plain", resources.Resources[0].MimeType)
}

func TestReadResource(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req JSONRPCRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		require.NoError(t, err)

		assert.Equal(t, "resources/read", req.Method)

		resp := JSONRPCResponse{
			JSONRPC: "2.0",
			ID:      req.ID,
			Result: json.RawMessage(`{
				"contents": [
					{
						"uri": "file:///data/readme.txt",
						"mimeType": "text/plain",
						"text": "# Hello World\n\nThis is a test document."
					}
				]
			}`),
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	content, err := client.ReadResource(context.Background(), "file:///data/readme.txt")

	require.NoError(t, err)
	assert.NotNil(t, content)
	assert.Len(t, content.Contents, 1)
	assert.Equal(t, "file:///data/readme.txt", content.Contents[0].URI)
	assert.Equal(t, "text/plain", content.Contents[0].MimeType)
	assert.Equal(t, "# Hello World\n\nThis is a test document.", content.Contents[0].Text)
}

func TestJSONRPCError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := JSONRPCResponse{
			JSONRPC: "2.0",
			ID:      1,
			Error: &JSONRPCError{
				Code:    -32601,
				Message: "Method not found",
			},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	_, err := client.ListTools(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "JSON-RPC error")
	assert.Contains(t, err.Error(), "-32601")
	assert.Contains(t, err.Error(), "Method not found")
}

func TestHTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}))
	defer server.Close()

	client := NewClient(server.URL)
	_, err := client.ListTools(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "HTTP error 500")
}

func TestContextCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Never respond
		<-r.Context().Done()
	}))
	defer server.Close()

	client := NewClient(server.URL)
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	_, err := client.ListTools(ctx, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "context canceled")
}
