# CreateCatalogMCPImplementation

Request body for creating an MCP implementation.


## Supported Types

### CreateCatalogMCPGatewayImplementation

```go
createCatalogMCPImplementation := components.CreateCreateCatalogMCPImplementationCreateCatalogMCPGatewayImplementation(components.CreateCatalogMCPGatewayImplementation{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createCatalogMCPImplementation.Type {
	case components.CreateCatalogMCPImplementationTypeCreateCatalogMCPGatewayImplementation:
		// createCatalogMCPImplementation.CreateCatalogMCPGatewayImplementation is populated
}
```
