# CreateMcpResourceRequestBody


## Supported Types

### APIResourcePayload

```go
createMcpResourceRequestBody := operations.CreateCreateMcpResourceRequestBodyAPIResourcePayload(components.APIResourcePayload{/* values here */})
```

### McpServerResourcePayload

```go
createMcpResourceRequestBody := operations.CreateCreateMcpResourceRequestBodyMcpServerResourcePayload(components.McpServerResourcePayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createMcpResourceRequestBody.Type {
	case operations.CreateMcpResourceRequestBodyTypeAPIResourcePayload:
		// createMcpResourceRequestBody.APIResourcePayload is populated
	case operations.CreateMcpResourceRequestBodyTypeMcpServerResourcePayload:
		// createMcpResourceRequestBody.McpServerResourcePayload is populated
}
```
