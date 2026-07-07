# MCPResourceInfo


## Supported Types

### RawAPIResource

```go
mcpResourceInfo := components.CreateMCPResourceInfoRawAPIResource(components.RawAPIResource{/* values here */})
```

### CatalogAPIResource

```go
mcpResourceInfo := components.CreateMCPResourceInfoCatalogAPIResource(components.CatalogAPIResource{/* values here */})
```

### RemoteMcpServerResource

```go
mcpResourceInfo := components.CreateMCPResourceInfoRemoteMcpServerResource(components.RemoteMcpServerResource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch mcpResourceInfo.Type {
	case components.MCPResourceInfoTypeRawAPIResource:
		// mcpResourceInfo.RawAPIResource is populated
	case components.MCPResourceInfoTypeCatalogAPIResource:
		// mcpResourceInfo.CatalogAPIResource is populated
	case components.MCPResourceInfoTypeRemoteMcpServerResource:
		// mcpResourceInfo.RemoteMcpServerResource is populated
}
```
