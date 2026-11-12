# CreateContextSourceRequestBody


## Supported Types

### APIResourcePayload

```go
createContextSourceRequestBody := operations.CreateCreateContextSourceRequestBodyAPI(components.APIResourcePayload{/* values here */})
```

### McpServerResourcePayload

```go
createContextSourceRequestBody := operations.CreateCreateContextSourceRequestBodyMcpServer(components.McpServerResourcePayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createContextSourceRequestBody.Type {
	case operations.CreateContextSourceRequestBodyTypeAPI:
		// createContextSourceRequestBody.APIResourcePayload is populated
	case operations.CreateContextSourceRequestBodyTypeMcpServer:
		// createContextSourceRequestBody.McpServerResourcePayload is populated
}
```
