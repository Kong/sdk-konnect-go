# UpdateContextSourceRequestBody


## Supported Types

### APIResourcePayload

```go
updateContextSourceRequestBody := operations.CreateUpdateContextSourceRequestBodyAPI(components.APIResourcePayload{/* values here */})
```

### McpServerResourcePayload

```go
updateContextSourceRequestBody := operations.CreateUpdateContextSourceRequestBodyMcpServer(components.McpServerResourcePayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateContextSourceRequestBody.Type {
	case operations.UpdateContextSourceRequestBodyTypeAPI:
		// updateContextSourceRequestBody.APIResourcePayload is populated
	case operations.UpdateContextSourceRequestBodyTypeMcpServer:
		// updateContextSourceRequestBody.McpServerResourcePayload is populated
}
```
