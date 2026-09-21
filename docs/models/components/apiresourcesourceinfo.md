# APIResourceSourceInfo

The source of an API MCP resource.


## Supported Types

### MCPResourceSourceAPICatalog

```go
apiResourceSourceInfo := components.CreateAPIResourceSourceInfoAPICatalog(components.MCPResourceSourceAPICatalog{/* values here */})
```

### MCPResourceSourceRaw

```go
apiResourceSourceInfo := components.CreateAPIResourceSourceInfoRaw(components.MCPResourceSourceRaw{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch apiResourceSourceInfo.Type {
	case components.APIResourceSourceInfoTypeAPICatalog:
		// apiResourceSourceInfo.MCPResourceSourceAPICatalog is populated
	case components.APIResourceSourceInfoTypeRaw:
		// apiResourceSourceInfo.MCPResourceSourceRaw is populated
}
```
