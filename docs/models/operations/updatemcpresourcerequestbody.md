# UpdateMcpResourceRequestBody


## Supported Types

### APIResourcePayload

```go
updateMcpResourceRequestBody := operations.CreateUpdateMcpResourceRequestBodyAPIResourcePayload(components.APIResourcePayload{/* values here */})
```

### McpServerResourcePayload

```go
updateMcpResourceRequestBody := operations.CreateUpdateMcpResourceRequestBodyMcpServerResourcePayload(components.McpServerResourcePayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateMcpResourceRequestBody.Type {
	case operations.UpdateMcpResourceRequestBodyTypeAPIResourcePayload:
		// updateMcpResourceRequestBody.APIResourcePayload is populated
	case operations.UpdateMcpResourceRequestBodyTypeMcpServerResourcePayload:
		// updateMcpResourceRequestBody.McpServerResourcePayload is populated
}
```
