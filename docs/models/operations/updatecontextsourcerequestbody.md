# UpdateContextSourceRequestBody


## Supported Types

### APIResourcePayload

```go
updateContextSourceRequestBody := operations.CreateUpdateContextSourceRequestBodyAPIResourcePayload(components.APIResourcePayload{/* values here */})
```

### McpServerResourcePayload

```go
updateContextSourceRequestBody := operations.CreateUpdateContextSourceRequestBodyMcpServerResourcePayload(components.McpServerResourcePayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateContextSourceRequestBody.Type {
	case operations.UpdateContextSourceRequestBodyTypeAPIResourcePayload:
		// updateContextSourceRequestBody.APIResourcePayload is populated
	case operations.UpdateContextSourceRequestBodyTypeMcpServerResourcePayload:
		// updateContextSourceRequestBody.McpServerResourcePayload is populated
}
```
