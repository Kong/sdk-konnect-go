# Resource


## Supported Types

### 

```go
resource := components.CreateResourceStr(string{/* values here */})
```

### CreateCatalogResourceMappingResourceByConfig

```go
resource := components.CreateResourceCreateCatalogResourceMappingResourceByConfig(components.CreateCatalogResourceMappingResourceByConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch resource.Type {
	case components.ResourceUnionTypeStr:
		// resource.Str is populated
	case components.ResourceUnionTypeCreateCatalogResourceMappingResourceByConfig:
		// resource.CreateCatalogResourceMappingResourceByConfig is populated
}
```
