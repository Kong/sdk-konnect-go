# ResourceIntegrationName


## Supported Types

### StringFieldFilter

```go
resourceIntegrationName := components.CreateResourceIntegrationNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch resourceIntegrationName.Type {
	case components.ResourceIntegrationNameTypeStringFieldFilter:
		// resourceIntegrationName.StringFieldFilter is populated
}
```
