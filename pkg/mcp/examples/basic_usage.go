package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Kong/sdk-konnect-go/pkg/mcp"
)

func main() {
	// Create a new MCP client pointing to your MCP server
	client := mcp.NewClient("http://localhost:3000/mcp")

	// Initialize the connection
	err := client.Initialize(context.Background(), mcp.ClientInfo{
		Name:    "konnect-mcp-client",
		Version: "1.0.0",
	})
	if err != nil {
		log.Fatalf("Failed to initialize MCP client: %v", err)
	}

	fmt.Printf("Connected to MCP server: %s v%s\n",
		client.ServerInfo().Name,
		client.ServerInfo().Version)

	// List available tools
	fmt.Println("\n=== Available Tools ===")
	tools, err := client.ListTools(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to list tools: %v", err)
	}

	for _, tool := range tools.Tools {
		fmt.Printf("- %s: %s\n", tool.Name, tool.Description)
		if len(tool.InputSchema.Required) > 0 {
			fmt.Printf("  Required parameters: %v\n", tool.InputSchema.Required)
		}
	}

	// Call a tool (example: weather tool)
	fmt.Println("\n=== Calling Tool ===")
	result, err := client.CallTool(context.Background(), mcp.CallToolRequest{
		Name: "get_weather",
		Arguments: map[string]interface{}{
			"location": "San Francisco",
		},
	})
	if err != nil {
		log.Fatalf("Failed to call tool: %v", err)
	}

	for _, content := range result.Content {
		fmt.Printf("Result: %s\n", content.Text)
	}

	// List available resources
	fmt.Println("\n=== Available Resources ===")
	resources, err := client.ListResources(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to list resources: %v", err)
	}

	for _, resource := range resources.Resources {
		fmt.Printf("- %s (%s): %s\n", resource.Name, resource.URI, resource.Description)
	}

	// Read a resource (example)
	if len(resources.Resources) > 0 {
		fmt.Println("\n=== Reading Resource ===")
		uri := resources.Resources[0].URI
		content, err := client.ReadResource(context.Background(), uri)
		if err != nil {
			log.Fatalf("Failed to read resource: %v", err)
		}

		for _, c := range content.Contents {
			fmt.Printf("Content from %s:\n%s\n", c.URI, c.Text)
		}
	}
}
