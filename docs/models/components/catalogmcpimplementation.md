# CatalogMCPImplementation

An MCP implementation resource.


## Supported Types

### CatalogMCPGatewayImplementation

```go
catalogMCPImplementation := components.CreateCatalogMCPImplementationCatalogMCPGatewayImplementation(components.CatalogMCPGatewayImplementation{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogMCPImplementation.Type {
	case components.CatalogMCPImplementationTypeCatalogMCPGatewayImplementation:
		// catalogMCPImplementation.CatalogMCPGatewayImplementation is populated
}
```
