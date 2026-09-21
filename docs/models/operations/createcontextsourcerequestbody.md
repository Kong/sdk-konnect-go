# CreateContextSourceRequestBody


## Supported Types

### APIResourcePayload

```go
createContextSourceRequestBody := operations.CreateCreateContextSourceRequestBodyAPIResourcePayload(components.APIResourcePayload{/* values here */})
```

### McpServerResourcePayload

```go
createContextSourceRequestBody := operations.CreateCreateContextSourceRequestBodyMcpServerResourcePayload(components.McpServerResourcePayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createContextSourceRequestBody.Type {
	case operations.CreateContextSourceRequestBodyTypeAPIResourcePayload:
		// createContextSourceRequestBody.APIResourcePayload is populated
	case operations.CreateContextSourceRequestBodyTypeMcpServerResourcePayload:
		// createContextSourceRequestBody.McpServerResourcePayload is populated
}
```
