# APIResourceSourcePayload

The source of the MCP resource used in create and update requests.


## Supported Types

### MCPResourceSourceAPICatalog

```go
apiResourceSourcePayload := components.CreateAPIResourceSourcePayloadAPICatalog(components.MCPResourceSourceAPICatalog{/* values here */})
```

### MCPResourceSourceRawInput

```go
apiResourceSourcePayload := components.CreateAPIResourceSourcePayloadRaw(components.MCPResourceSourceRawInput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch apiResourceSourcePayload.Type {
	case components.APIResourceSourcePayloadTypeAPICatalog:
		// apiResourceSourcePayload.MCPResourceSourceAPICatalog is populated
	case components.APIResourceSourcePayloadTypeRaw:
		// apiResourceSourcePayload.MCPResourceSourceRawInput is populated
}
```
