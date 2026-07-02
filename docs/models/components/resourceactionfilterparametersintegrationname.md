# ResourceActionFilterParametersIntegrationName


## Supported Types

### StringFieldFilter

```go
resourceActionFilterParametersIntegrationName := components.CreateResourceActionFilterParametersIntegrationNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch resourceActionFilterParametersIntegrationName.Type {
	case components.ResourceActionFilterParametersIntegrationNameTypeStringFieldFilter:
		// resourceActionFilterParametersIntegrationName.StringFieldFilter is populated
}
```
