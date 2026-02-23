package mcp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"
)

// Client represents an MCP client
type Client struct {
	serverURL    string
	httpClient   *http.Client
	requestID    atomic.Int64
	serverInfo   *ServerInfo
	capabilities *ServerCapabilities
}

// NewClient creates a new MCP client
func NewClient(serverURL string, httpClient ...*http.Client) *Client {
	client := &Client{
		serverURL: serverURL,
	}
	if len(httpClient) > 0 && httpClient[0] != nil {
		client.httpClient = httpClient[0]
	} else {
		client.httpClient = &http.Client{}
	}
	return client
}

// Initialize performs the MCP initialization handshake
func (c *Client) Initialize(ctx context.Context, clientInfo ClientInfo) error {
	initReq := InitializeRequest{
		ProtocolVersion: ProtocolVersion,
		Capabilities: Capabilities{
			Roots: &RootsCapability{
				ListChanged: true,
			},
		},
		ClientInfo: clientInfo,
	}

	var initResp InitializeResponse
	if err := c.call(ctx, "initialize", initReq, &initResp); err != nil {
		return fmt.Errorf("failed to initialize: %w", err)
	}

	c.serverInfo = &initResp.ServerInfo
	c.capabilities = &initResp.Capabilities

	// Send initialized notification
	if err := c.notify(ctx, "notifications/initialized", nil); err != nil {
		return fmt.Errorf("failed to send initialized notification: %w", err)
	}

	return nil
}

// ListTools retrieves the list of available tools from the server
func (c *Client) ListTools(ctx context.Context, cursor *string) (*ListToolsResponse, error) {
	params := make(map[string]interface{})
	if cursor != nil {
		params["cursor"] = *cursor
	}

	var resp ListToolsResponse
	if err := c.call(ctx, "tools/list", params, &resp); err != nil {
		return nil, fmt.Errorf("failed to list tools: %w", err)
	}

	return &resp, nil
}

// CallTool calls a tool on the server
func (c *Client) CallTool(ctx context.Context, req CallToolRequest) (*CallToolResponse, error) {
	var resp CallToolResponse
	if err := c.call(ctx, "tools/call", req, &resp); err != nil {
		return nil, fmt.Errorf("failed to call tool: %w", err)
	}

	return &resp, nil
}

// ListResources retrieves the list of available resources from the server
func (c *Client) ListResources(ctx context.Context, cursor *string) (*ListResourcesResponse, error) {
	params := make(map[string]interface{})
	if cursor != nil {
		params["cursor"] = *cursor
	}

	var resp ListResourcesResponse
	if err := c.call(ctx, "resources/list", params, &resp); err != nil {
		return nil, fmt.Errorf("failed to list resources: %w", err)
	}

	return &resp, nil
}

// ReadResource reads a resource from the server
func (c *Client) ReadResource(ctx context.Context, uri string) (*ReadResourceResponse, error) {
	req := ReadResourceRequest{URI: uri}

	var resp ReadResourceResponse
	if err := c.call(ctx, "resources/read", req, &resp); err != nil {
		return nil, fmt.Errorf("failed to read resource: %w", err)
	}

	return &resp, nil
}

// ServerInfo returns information about the connected server
func (c *Client) ServerInfo() *ServerInfo {
	return c.serverInfo
}

// Capabilities returns the server's capabilities
func (c *Client) Capabilities() *ServerCapabilities {
	return c.capabilities
}

// call makes a JSON-RPC request with a response
func (c *Client) call(ctx context.Context, method string, params interface{}, result interface{}) error {
	reqID := c.requestID.Add(1)

	req := JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      reqID,
		Method:  method,
		Params:  params,
	}

	respBody, err := c.sendRequest(ctx, req)
	if err != nil {
		return err
	}

	var resp JSONRPCResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if resp.Error != nil {
		return fmt.Errorf("JSON-RPC error (code %d): %s", resp.Error.Code, resp.Error.Message)
	}

	if result != nil && resp.Result != nil {
		if err := json.Unmarshal(resp.Result, result); err != nil {
			return fmt.Errorf("failed to unmarshal result: %w", err)
		}
	}

	return nil
}

// notify makes a JSON-RPC notification (no response expected)
func (c *Client) notify(ctx context.Context, method string, params interface{}) error {
	req := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
	}

	_, err := c.sendRequest(ctx, req)
	return err
}

// sendRequest sends a JSON-RPC request via HTTP
func (c *Client) sendRequest(ctx context.Context, req JSONRPCRequest) ([]byte, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.serverURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(httpResp.Body)
		return nil, fmt.Errorf("HTTP error %d: %s", httpResp.StatusCode, string(body))
	}

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return respBody, nil
}
